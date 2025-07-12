package clientpackets

import (
	"errors"
	"fmt"
	"l2gogameserver/data/logger"
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
	"l2gogameserver/utils"
	"os"
	"strconv"
	"strings"
)

const (
	// Команды
	CmdTeleport = "teleport"
	CmdAdmin    = "admin"

	// Константы для телепортации
	DefaultZ = 300
	DefaultH = 0

	// Пути к файлам
	AdminMainHTMLPath = "./datapack/html/admin/main.html"
)

var (
	ErrInvalidCommandFormat = errors.New("неверный формат команды")
	ErrInvalidCoordinates   = errors.New("неверные координаты")
	ErrFileNotFound         = errors.New("файл не найден")
)

// CommandHandler обрабатывает команды, отправленные клиентом
type CommandHandler struct {
	character interfaces.CharacterI
	client    interfaces.NewClientCtxInterface
	gs        GsInterfNew
}

// NewCommandHandler создает новый обработчик команд
func NewCommandHandler(character interfaces.CharacterI, client interfaces.NewClientCtxInterface, gs GsInterfNew) *CommandHandler {
	return &CommandHandler{
		character: character,
		client:    client,
		gs:        gs,
	}
}

// SendBypassBuildCmd обрабатывает команды, отправленные клиентом
// TODO: Добавить проверку на админа
// Когда персонаж кликает по карте CTRL+SHIFT+CLICK передаются координаты X Y и персонажа телепортируем.
func SendBypassBuildCmd(character interfaces.CharacterI, data []byte, gs GsInterfNew, client interfaces.NewClientCtxInterface) {
	logger.Info.Println("=== SendBypassBuildCmd вызвана ===")

	handler := NewCommandHandler(character, client, gs)

	if err := handler.handleCommand(data); err != nil {
		logger.Info.Printf("Ошибка обработки команды: %v", err)
	}
}

// handleCommand обрабатывает команду из пакета данных
func (h *CommandHandler) handleCommand(data []byte) error {
	packet := packets.NewReader(data)
	command := packet.ReadString()

	logger.Info.Printf("Получена команда: %s", command)

	// Проверяем команду телепортации отдельно, так как она содержит координаты
	if strings.HasPrefix(command, CmdTeleport) {
		return h.handleTeleportCommand(command)
	}

	switch command {
	case CmdAdmin:
		return h.handleAdminCommand()
	default:
		logger.Info.Printf("Неизвестная команда: %s", command)
		return nil
	}
}

// handleTeleportCommand обрабатывает команду телепортации
func (h *CommandHandler) handleTeleportCommand(command string) error {
	parts := strings.Fields(command)
	if len(parts) != 3 || parts[0] != CmdTeleport {
		return fmt.Errorf("%w: %s", ErrInvalidCommandFormat, command)
	}

	locX, err := strconv.Atoi(parts[1])
	if err != nil {
		return fmt.Errorf("%w: X координата: %v", ErrInvalidCoordinates, err)
	}

	locY, err := strconv.Atoi(parts[2])
	if err != nil {
		return fmt.Errorf("%w: Y координата: %v", ErrInvalidCoordinates, err)
	}

	h.teleport(locX, locY)
	return nil
}

// handleAdminCommand обрабатывает команду админа
func (h *CommandHandler) handleAdminCommand() error {
	data, err := os.ReadFile(AdminMainHTMLPath)
	if err != nil {
		return fmt.Errorf("%w: %s: %v", ErrFileNotFound, AdminMainHTMLPath, err)
	}

	htmlContent := utils.B2s(data)
	accountLogin := h.character.GetAccountLogin()
	client := h.gs.GetClientByLogin(accountLogin)

	if client == nil {
		return errors.New("клиент не найден")
	}

	return client.SendBuf(serverpackets.NpcHtmlMessage2(0, htmlContent, 0))
}

// teleport телепортирует персонажа в указанные координаты
func (h *CommandHandler) teleport(locX, locY int) {
	x, y, z, heading := locX, locY, DefaultZ, DefaultH
	pkg := serverpackets.TeleportToLocation(h.character, x, y, z, heading)

	if err := h.client.EncryptAndSend(pkg); err != nil {
		logger.Info.Printf("Ошибка отправки пакета телепортации: %v", err)
	}
}

// teleport телепортирует персонажа в указанные координаты (устаревшая функция)
// Deprecated: используйте CommandHandler.teleport
func teleport(character interfaces.CharacterI, client interfaces.NewClientCtxInterface, locX int, locY int) {
	handler := &CommandHandler{
		character: character,
		client:    client,
	}
	handler.teleport(locX, locY)
}
