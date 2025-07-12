package handlers

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"l2gogameserver/config"
	"l2gogameserver/data/logger"
	"l2gogameserver/gameserver/broadcast"
	"l2gogameserver/gameserver/clientpackets"
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/listeners"
	"l2gogameserver/gameserver/models/clientStates"
	"l2gogameserver/packets"
	"sync"
)

// Константы для опкодов
const (
	// Основные опкоды
	OpcodeProtocolVersion             = 0x0e
	OpcodeAuthLogin                   = 0x2b
	OpcodeLogout                      = 0x00
	OpcodeCharacterCreate             = 0x0c
	OpcodeCharacterDelete             = 0x0d
	OpcodeCharSelected                = 0x12
	OpcodeRequestNewCharacter         = 0x13
	OpcodeRequestEnterWorld           = 0x11
	OpcodeAttack                      = 0x01
	OpcodeTradeRequest                = 0x1a
	OpcodeAnswerTradeRequest          = 0x55
	OpcodeAddTradeItem                = 0x1b
	OpcodeTradeDone                   = 0x1c
	OpcodeDropItem                    = 0x17
	OpcodeRequestItemList             = 0x14
	OpcodeBypassToServer              = 0x23
	OpcodeUseItem                     = 0x19
	OpcodeSetPrivateStoreListSell     = 0x31
	OpcodeRequestMagicSkillUse        = 0x39
	OpcodeRequestShortCutReg          = 0x3d
	OpcodeRequestShortCutDel          = 0x3f
	OpcodeRequestRestart              = 0x57
	OpcodeDestroyItem                 = 0x60
	OpcodeRequestObserverEnd          = 0xc1
	OpcodeFinishRotating              = 0x5c
	OpcodeRequestShowMiniMap          = 0x6c
	OpcodeRequestSkillCoolTime        = 0xa6
	OpcodeMoveBackwardToLocation      = 0x0f
	OpcodeRequestJoinParty            = 0x42
	OpcodeRequestAnswerJoinParty      = 0x43
	OpcodeRequestWithDrawalParty      = 0x44
	OpcodeSay                         = 0x49
	OpcodeValidationPosition          = 0x59
	OpcodeRequestSkillList            = 0x50
	OpcodeAction                      = 0x1f
	OpcodeRequestTargetCancel         = 0x48
	OpcodeRequestMakeMacro            = 0xcd
	OpcodeRequestDeleteMacro          = 0xce
	OpcodeRequestActionUse            = 0x56
	OpcodeSendBypassBuildCmd          = 0x74
	OpcodeRequestPrivateStoreBuy      = 0x83
	OpcodeRequestPrivateStoreQuitSell = 0x96
	OpcodeSetPrivateStoreMsgSell      = 0x97

	// Вложенные опкоды (для 0xd0)
	OpcodeRequestGoToLobby          = 0x36
	OpcodeRequestManorList          = 0x01
	OpcodeRequestSaveInventoryOrder = 0x24
	OpcodeRequestAutoSoulShot       = 0x0d
	OpcodeAnswerCoupleAction        = 0x7a
)

// Вложенный опкод для 0xd0
const (
	OpcodeD0 = 0xd0
)

var (
	ErrUnknownOpcode      = errors.New("неизвестный опкод")
	ErrInvalidState       = errors.New("неверное состояние клиента")
	ErrClientDisconnected = errors.New("клиент отключен")
)

// GameServerInterface определяет интерфейс для игрового сервера
type GameServerInterface interface {
	AddClient(string, interfaces.NewClientCtxInterface) bool
	AddWaitClient(string, uint32, uint32, uint32, uint32)
	RemoveClient(string)
	SendLogout(string)
	GetDbConn() *sql.DB
	CharOffline(ctxInterface interfaces.NewClientCtxInterface)
	AddOnlineChar(character interfaces.CharacterI)
	GetChar(string) (interfaces.CharacterI, bool)
	GetNetConnByCharObjectId(objectId int32) interfaces.CharacterI
	GetClientByLogin(login string) interfaces.NewClientCtxInterface
}

// PacketHandler обрабатывает пакеты от клиентов
type PacketHandler struct {
	gs GameServerInterface
	mu sync.RWMutex
}

// NewPacketHandler создает новый обработчик пакетов
func NewPacketHandler(gs GameServerInterface) *PacketHandler {
	return &PacketHandler{
		gs: gs,
	}
}

// Handler обрабатывает пакеты от клиента
func Handler(client interfaces.NewClientCtxInterface, gs GameServerInterface) {
	handler := NewPacketHandler(gs)
	handler.handleClient(client)
}

// handleClient обрабатывает пакеты от клиента в бесконечном цикле
func (h *PacketHandler) handleClient(client interfaces.NewClientCtxInterface) {
	defer func() {
		if r := recover(); r != nil {
			logger.Info.Printf("Паника в обработчике пакетов: %v", r)
		}
		h.gs.CharOffline(client)
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			return
		default:
			if err := h.processPacket(client); err != nil {
				if errors.Is(err, ErrClientDisconnected) {
					logger.Info.Println("Клиент отключен")
					return
				}
				logger.LogError("Ошибка обработки пакета: %v", err)
			}
		}
	}
}

// processPacket обрабатывает один пакет от клиента
func (h *PacketHandler) processPacket(client interfaces.NewClientCtxInterface) error {
	opcode, data, err := client.Receive()
	if err != nil {
		return fmt.Errorf("%w: %v", ErrClientDisconnected, err)
	}

	if config.GetDebug().IsShowPackets() {
		logger.LogClientToServer("Client->Server: #%02x %s", opcode, packets.GetNamePacket(opcode))
	}

	state := client.GetState()
	return h.handlePacketByState(client, opcode, data, state)
}

// handlePacketByState обрабатывает пакет в зависимости от состояния клиента
func (h *PacketHandler) handlePacketByState(client interfaces.NewClientCtxInterface, opcode byte, data []byte, state clientStates.State) error {
	switch state {
	case clientStates.Connected:
		return h.handleConnectedState(client, opcode, data)
	case clientStates.Authed:
		return h.handleAuthedState(client, opcode, data)
	case clientStates.Joining:
		return h.handleJoiningState(client, opcode, data)
	case clientStates.InGame:
		return h.handleInGameState(client, opcode, data)
	default:
		return fmt.Errorf("%w: состояние %v", ErrInvalidState, state)
	}
}

// handleConnectedState обрабатывает пакеты в состоянии Connected
func (h *PacketHandler) handleConnectedState(client interfaces.NewClientCtxInterface, opcode byte, data []byte) error {
	switch opcode {
	case OpcodeProtocolVersion:
		clientpackets.ProtocolVersion(client, data)
	case OpcodeAuthLogin:
		clientpackets.AuthLogin(data, client, h.gs)
	default:
		logger.Info.Printf("Неизвестный опкод 0x%02x в состоянии Connected", opcode)
		return fmt.Errorf("%w: 0x%02x в состоянии Connected", ErrUnknownOpcode, opcode)
	}
	return nil
}

// handleAuthedState обрабатывает пакеты в состоянии Authed
func (h *PacketHandler) handleAuthedState(client interfaces.NewClientCtxInterface, opcode byte, data []byte) error {
	switch opcode {
	case OpcodeLogout:
		clientpackets.Logout(client, clientStates.Authed, h.gs)
	case OpcodeCharacterCreate:
		clientpackets.CharacterCreate(client, data, h.gs.GetDbConn())
	case OpcodeCharacterDelete:
		clientpackets.CharacterDelete(client, data, h.gs.GetDbConn())
	case OpcodeCharSelected:
		clientpackets.CharSelected(data, client, h.gs.GetDbConn())
		h.gs.AddOnlineChar(client.GetAccount().GetCurrentChar())
	case OpcodeRequestNewCharacter:
		clientpackets.RequestNewCharacter(client)
	case OpcodeD0:
		return h.handleD0Packet(client, data, clientStates.Authed)
	default:
		logger.Info.Printf("Неизвестный опкод 0x%02x в состоянии Authed", opcode)
		return fmt.Errorf("%w: 0x%02x в состоянии Authed", ErrUnknownOpcode, opcode)
	}
	return nil
}

// handleJoiningState обрабатывает пакеты в состоянии Joining
func (h *PacketHandler) handleJoiningState(client interfaces.NewClientCtxInterface, opcode byte, data []byte) error {
	switch opcode {
	case OpcodeRequestEnterWorld:
		clientpackets.RequestEnterWorld(client, data, h.gs.GetDbConn())
		go listeners.StartClientListener(client, h.gs)
	case OpcodeD0:
		return h.handleD0Packet(client, data, clientStates.Joining)
	default:
		logger.Info.Printf("Неизвестный опкод 0x%02x в состоянии Joining", opcode)
		return fmt.Errorf("%w: 0x%02x в состоянии Joining", ErrUnknownOpcode, opcode)
	}
	return nil
}

// handleInGameState обрабатывает пакеты в состоянии InGame
func (h *PacketHandler) handleInGameState(client interfaces.NewClientCtxInterface, opcode byte, data []byte) error {
	character := client.GetAccount().GetCurrentChar()

	switch opcode {
	case OpcodeLogout:
		clientpackets.Logout(client, clientStates.InGame, h.gs)
	case OpcodeAttack:
		clientpackets.Attack(data, client)
	case OpcodeTradeRequest:
		clientpackets.TradeRequest(data, client, h.gs)
	case OpcodeAnswerTradeRequest:
		clientpackets.AnswerTradeRequest(data, client, h.gs)
	case OpcodeAddTradeItem:
		clientpackets.AddTradeItem(data, client, h.gs)
	case OpcodeTradeDone:
		clientpackets.TradeDone(data, client, h.gs.GetDbConn(), h.gs)
	case OpcodeDropItem:
		clientpackets.DropItem(client, data, h.gs.GetDbConn())
	case OpcodeRequestItemList:
		clientpackets.RequestItemList(data, client)
	case OpcodeBypassToServer:
		clientpackets.BypassToServer(data, client, h.gs.GetDbConn())
	case OpcodeUseItem:
		clientpackets.UseItem(character, data, h.gs.GetDbConn(), h.gs)
	case OpcodeSetPrivateStoreListSell:
		clientpackets.SetPrivateStoreListSell(client, data, h.gs)
	case OpcodeRequestMagicSkillUse:
		clientpackets.RequestMagicSkillUse(data, client)
	case OpcodeRequestShortCutReg:
		clientpackets.RequestShortCutReg(data, client, h.gs.GetDbConn())
	case OpcodeRequestShortCutDel:
		clientpackets.RequestShortCutDel(data, client, h.gs.GetDbConn())
	case OpcodeRequestRestart:
		clientpackets.RequestRestart(client, h.gs, h.gs.GetDbConn())
	case OpcodeDestroyItem:
		clientpackets.DestroyItem(data, client, h.gs.GetDbConn())
	case OpcodeRequestObserverEnd:
		clientpackets.RequestObserverEnd(client, data)
	case OpcodeFinishRotating:
		clientpackets.FinishRotating(client, data)
	case OpcodeRequestShowMiniMap:
		clientpackets.RequestShowMiniMap(client, data)
	case OpcodeRequestSkillCoolTime:
		clientpackets.RequestSkillCoolTime(client, data)
	case OpcodeMoveBackwardToLocation:
		pkg := clientpackets.MoveBackwardToLocation(client, data)
		broadcast.Checkaem(client, pkg)
	case OpcodeRequestJoinParty:
		clientpackets.RequestJoinParty(client, data, h.gs)
	case OpcodeRequestAnswerJoinParty:
		clientpackets.RequestAnswerJoinParty(client, data, h.gs)
	case OpcodeRequestWithDrawalParty:
		clientpackets.RequestWithDrawalParty(client, h.gs)
	case OpcodeSay:
		say := clientpackets.Say(client, data)
		broadcast.BroadCastChat(client, say, h.gs)
	case OpcodeValidationPosition:
		clientpackets.ValidationPosition(data, character)
	case OpcodeRequestSkillList:
		clientpackets.RequestSkillList(client, data)
	case OpcodeAction:
		pkg := clientpackets.Action(data, client, broadcast.Checkaem, h.gs.GetDbConn())
		if pkg != nil {
			broadcast.Checkaem(client, *pkg)
		}
	case OpcodeRequestTargetCancel:
		clientpackets.RequestTargetCancel(data, client)
	case OpcodeRequestMakeMacro:
		clientpackets.RequestMakeMacro(client, data)
	case OpcodeRequestDeleteMacro:
		clientpackets.RequestDeleteMacro(client, data, h.gs.GetDbConn())
	case OpcodeRequestActionUse:
		clientpackets.RequestActionUse(client, data, h.gs)
	case OpcodeSendBypassBuildCmd:
		logger.Info.Println("Получен пакет SendBypassBuildCmd (0x74)")
		clientpackets.SendBypassBuildCmd(character, data, h.gs, client)
	case OpcodeRequestPrivateStoreBuy:
		clientpackets.RequestPrivateStoreBuy(client, data, h.gs.GetDbConn())
	case OpcodeRequestPrivateStoreQuitSell:
		clientpackets.RequestPrivateStoreQuitSell(client)
	case OpcodeSetPrivateStoreMsgSell:
		clientpackets.SetPrivateStoreMsgSell(client, data)
	case OpcodeD0:
		return h.handleD0Packet(client, data, clientStates.InGame)
	default:
		logger.LogInfo("Неизвестный опкод 0x%02x в состоянии InGame", opcode)
		return fmt.Errorf("%w: 0x%02x в состоянии InGame", ErrUnknownOpcode, opcode)
	}
	return nil
}

// handleD0Packet обрабатывает вложенные пакеты с опкодом 0xd0
func (h *PacketHandler) handleD0Packet(client interfaces.NewClientCtxInterface, data []byte, state clientStates.State) error {
	if len(data) < 2 {
		return errors.New("недостаточно данных для обработки вложенного пакета")
	}

	subOpcode := data[0]
	switch state {
	case clientStates.Authed:
		switch subOpcode {
		case OpcodeRequestGoToLobby:
			clientpackets.RequestGoToLobby(client, h.gs.GetDbConn())
		default:
			logger.Info.Printf("Неизвестный вложенный опкод 0x%02x в состоянии Authed", subOpcode)
			return fmt.Errorf("%w: вложенный 0x%02x в состоянии Authed", ErrUnknownOpcode, subOpcode)
		}
	case clientStates.Joining:
		switch subOpcode {
		case OpcodeRequestManorList:
			clientpackets.RequestManorList(client, data)
		default:
			logger.LogError("Неизвестный вложенный опкод 0x%02x в состоянии Joining", subOpcode)
			return fmt.Errorf("%w: вложенный 0x%02x в состоянии Joining", ErrUnknownOpcode, subOpcode)
		}
	case clientStates.InGame:
		switch subOpcode {
		case OpcodeRequestSaveInventoryOrder:
			clientpackets.RequestSaveInventoryOrder(client, data)
		case OpcodeRequestAutoSoulShot:
			clientpackets.RequestAutoSoulShot(data, client)
		case OpcodeAnswerCoupleAction:
			clientpackets.AnswerCoupleAction(client, data)
		default:
			logger.Info.Printf("Неизвестный вложенный опкод 0x%02x в состоянии InGame", subOpcode)
			return fmt.Errorf("%w: вложенный 0x%02x в состоянии InGame", ErrUnknownOpcode, subOpcode)
		}
	default:
		return fmt.Errorf("%w: состояние %v", ErrInvalidState, state)
	}
	return nil
}
