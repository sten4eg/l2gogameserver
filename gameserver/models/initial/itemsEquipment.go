package initial

import (
	"encoding/json"
	"l2gogameserver/config"
	"l2gogameserver/data/logger"
	"log"
	"os"
)

type ItemsEquipment struct {
	ClassId int32  `json:"classId"`
	Name    string `json:"name"`
	Items   []struct {
		Id       int    `json:"id"`
		Count    int    `json:"count"`
		Equipped bool   `json:"equipped"`
		Name     string `json:"name"`
	} `json:"items"`
}

var equipmentData []ItemsEquipment

func LoadItemsEquipment() {
	if !config.IsEnableItems() {
		return
	}
	logger.Info.Println("Загрузка предметов")
	byteValue, err := os.ReadFile("./datapack/data/stats/char/initial/equipment.json")
	if err != nil {
		if os.IsNotExist(err) {
			logger.Error.Panicln("Файл equipment.json не найден:", err)
		} else {
			logger.Error.Panicln("Не удалось прочитать equipment.json:", err)
		}
	}
	if err := json.Unmarshal(byteValue, &equipmentData); err != nil {
		logger.Error.Panicln("Ошибка при разборе JSON equipment.json:", err)
	}
	eq := GetEquipmentByClass(0)
	if eq == nil {
		log.Println("Предметы для класса не найдены")
	}
}

// GetEquipmentByClass возвращает полную структуру itemsEquipment для указанного classId
func GetEquipmentByClass(classId int32) *ItemsEquipment {
	for _, eq := range equipmentData {
		if eq.ClassId == classId {
			return &eq
		}
	}
	return nil
}

// HasEquipmentForClass проверяет, существует начальный класс ID.
func HasEquipmentForClass(classId int32) bool {
	for _, eq := range equipmentData {
		if eq.ClassId == classId {
			return true
		}
	}
	return false
}
