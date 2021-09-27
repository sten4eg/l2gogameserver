package gameserver

import (
	"fmt"
	"l2gogameserver/gameserver/clientpackets"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/models/attack"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
	"l2gogameserver/utils"
	"log"
	"net"
	"sync"
	"time"
)

type GameServer struct {
	clientsListener  net.Listener
	clients          sync.Map
	Socket           net.Conn
	OnlineCharacters *models.OnlineCharacters
}

func (g *GameServer) AddClient(c *models.Client) {
	g.clients.Store(c.Account.Login, c)
}
func New() *GameServer {
	return &GameServer{}
}

func (g *GameServer) Start() {
	var err error

	/* #nosec */
	g.clientsListener, err = net.Listen("tcp4", ":7777")
	if err != nil {
		panic(err.Error())
	}

	var onlineChars models.OnlineCharacters

	onlineChars.Char = make(map[int32]*models.Character)
	g.OnlineCharacters = &onlineChars

	//go g.Tick()
	defer g.clientsListener.Close()
	for {
		client := models.NewClient()
		client.Socket, err = g.clientsListener.Accept()

		if err != nil {
			fmt.Println("Couldn't accept the incoming connection.", err)
			continue
		} else {
			g.AddClient(client)
			go g.handler(client)
		}
	}
}

func (g *GameServer) ChannelListener(client *models.Client) {
	for q := range client.CurrentChar.F {
		pkg := serverpackets.ItemUpdate(client, q.UpdateType, q.ObjId)
		i := client.CryptAndReturnPackageReadyToShip(pkg)
		client.SSend(i)
		if q.UpdateType == models.UpdateTypeRemove {
			g.BroadCastUserInfoInRadius(client, 2000)
		}
	}
}

func (g *GameServer) NpcListener(client *models.Client) {
	for q := range client.CurrentChar.NpcInfo {
		buff := packets.Get()
		for _, v := range q {
			pkg := serverpackets.NpcInfo(v)
			buff.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg))
		}
		client.SSend(buff.Bytes())
		packets.Put(buff)
	}
}
func (g *GameServer) MoveListener(client *models.Client) {

	for q := range client.CurrentChar.CharInfoTo {
		var pkg utils.PacketByte
		pkg.SetB(serverpackets.CharInfo(client.CurrentChar))
		for _, v := range q {
			g.OnlineCharacters.Mu.Lock()
			g.OnlineCharacters.Char[v].Conn.Send(pkg.GetB(), true)
			g.OnlineCharacters.Mu.Unlock()
		}
	}

	for q := range client.CurrentChar.DeleteObjectTo {
		var pkg utils.PacketByte
		pkg.SetB(serverpackets.DeleteObject(client.CurrentChar))
		for _, v := range q {
			g.OnlineCharacters.Mu.Lock()
			g.OnlineCharacters.Char[v].Conn.Send(pkg.GetB(), true)
			g.OnlineCharacters.Mu.Unlock()
		}
	}

}

func (g *GameServer) charOffline(client *models.Client) {
	g.OnlineCharacters.Mu.Lock()
	delete(g.OnlineCharacters.Char, client.CurrentChar.ObjectId)
	g.OnlineCharacters.Mu.Unlock()
	client.CurrentChar.CurrentRegion.DeleteVisibleChar(client.CurrentChar)

	client.CurrentChar.F = nil
	client.CurrentChar.NpcInfo = nil
	client.CurrentChar.CharInfoTo = nil
	client.CurrentChar.DeleteObjectTo = nil

	//todo close all character goroutine, save info in DB
	log.Println("Socket Close For: ", client.CurrentChar.CharName)
}

func (g *GameServer) addOnlineChar(character *models.Character) {
	g.OnlineCharacters.Mu.Lock()
	g.OnlineCharacters.Char[character.ObjectId] = character
	g.OnlineCharacters.Mu.Unlock()
}

func (g *GameServer) brdsct(me *models.Client, pkg utils.PacketByte) {
	charsIds := models.GetAroundPlayer(me.CurrentChar)
	for _, v := range charsIds {
		v.Conn.Send(pkg.GetB(), true)
		//me.SSend(me.CryptAndReturnPackageReadyToShip(serverpackets.CharInfo(v)))
	}
}

//Действие при атаки
func (g *GameServer) IsAttack(data []byte, client *models.Client) {
	npc, npcx, npcy, npcz, targetObjectId, actionId, reAppeal, err := attack.IsAttack(data, client)
	if err != nil {
		return
	}
	_ = actionId
	if reAppeal == true {
		distance := attack.DistanceToNpc(client, npcx, npcy, npcz)
		if distance <= 60 { //Если дистанция соблюдена
			attack.ReAppeal(npc, npcx, npcy, npcz, targetObjectId, data, client)
		} else { //Если необходимо добежать до цели
			pkg2 := clientpackets.MoveToLocation(client, npcx, npcy, npcz)
			g.Checkaem(client, pkg2)
			go func() {
				for {
					/**
					Увы, пока в коде формулы подсчета месторасположения игрока
					Предпологается что во время бега, мы будем постоянно обновлять игрока XYZ
					Игрок нажимает на цель (берет в таргет), и потом нажимает "атаковать" NPC
					персонаж будет бежать до NPC, и когда останется N расстояния до NPC нужно
					открыть диалоговое окно с NPC, с мобом точно так же только начинаем атаковать.

					Сейчас же работает так, берем в таргет ,жмем ещё раз атаковать и персонаж подбегает
					к NPC,и необходимо сновать нажать атаку, для открытия диалога либо атаки.
					*/
					distance = attack.DistanceToNpc(client, npcx, npcy, npcz)
					if client.CurrentChar.LastOpcode == 72 { //отменяем необходимость бежать к NPC либо атаковать, либо снимаем каст
						log.Println("cancel target")
						return
					}
					if distance <= 60 {
						attack.ReAppeal(npc, npcx, npcy, npcz, targetObjectId, data, client)
						return
					}
					time.Sleep(250 * time.Millisecond)
				}
			}()
		}

	}
}

//Действие при атаки
func (g *GameServer) IsAttack(data []byte, client *models.Client) {
	npc, npcx, npcy, npcz, targetObjectId, actionId, reAppeal, err := attack.IsAttack(data, client)
	if err != nil {
		return
	}
	_ = actionId
	if reAppeal == true {
		distance := attack.DistanceToNpc(client, npcx, npcy, npcz)
		if distance <= 60 { //Если дистанция соблюдена
			attack.ReAppeal(npc, npcx, npcy, npcz, targetObjectId, data, client)
		} else { //Если необходимо добежать до цели
			pkg2 := clientpackets.MoveToLocation(client, npcx, npcy, npcz)
			g.Checkaem(client, pkg2)
			go func() {
				for {
					/**
					Увы, пока в коде формулы подсчета месторасположения игрока
					Предпологается что во время бега, мы будем постоянно обновлять игрока XYZ
					Игрок нажимает на цель (берет в таргет), и потом нажимает "атаковать" NPC
					персонаж будет бежать до NPC, и когда останется N расстояния до NPC нужно
					открыть диалоговое окно с NPC, с мобом точно так же только начинаем атаковать.

					Сейчас же работает так, берем в таргет ,жмем ещё раз атаковать и персонаж подбегает
					к NPC,и необходимо сновать нажать атаку, для открытия диалога либо атаки.
					*/
					distance = attack.DistanceToNpc(client, npcx, npcy, npcz)
					if client.CurrentChar.LastOpcode == 72 { //отменяем необходимость бежать к NPC либо атаковать, либо снимаем каст
						log.Println("cancel target")
						return
					}
					if distance <= 60 {
						attack.ReAppeal(npc, npcx, npcy, npcz, targetObjectId, data, client)
						return
					}
					time.Sleep(250 * time.Millisecond)
				}
			}()
		}

	}
}
