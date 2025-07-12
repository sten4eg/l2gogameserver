package clientpackets

import (
	"fmt"
	"l2gogameserver/data/logger"
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
	"strconv"
	"strings"
)

// TODO: Нужно добавить проверку на админа
// Когда персонаж кликает по карте CLTR+SHIFT+CLICK передаются координаты X Y и персонажа телепортируем.
func SendBypassBuildCmd(character interfaces.CharacterI, data []byte, client interfaces.NewClientCtxInterface) {
	fmt.Println("=== SendBypassBuildCmd вызвана ===")
	logger.Info.Println("=== SendBypassBuildCmd вызвана ===")

	var packet = packets.NewReader(data)

	// Читаем строку команды
	command := packet.ReadString()
	logger.Info.Println("Получена команда:", command)

	// Парсим команду "teleport X Y"
	parts := strings.Fields(command)
	if len(parts) != 3 || parts[0] != "teleport" {
		logger.Info.Println("Неверный формат команды:", command)
		return
	}

	// Парсим координаты
	locX, err := strconv.Atoi(parts[1])
	if err != nil {
		logger.Info.Println("Ошибка парсинга X координаты:", err)
		return
	}

	locY, err := strconv.Atoi(parts[2])
	if err != nil {
		logger.Info.Println("Ошибка парсинга Y координаты:", err)
		return
	}

	x, y, z, h := locX, locY, 300, 0
	pkg := serverpackets.TeleportToLocation(character, int(x), int(y), z, h)
	client.EncryptAndSend(pkg)
}
