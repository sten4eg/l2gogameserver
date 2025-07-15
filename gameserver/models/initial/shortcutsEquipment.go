package initial

import (
	"encoding/json"
	"io/ioutil"
	"l2gogameserver/data/logger"
	"log"
	"os"
)

// Глобальные переменные (временно, без защиты мьютексами)
var (
	allShortcutsByClassID map[int][]PageEntry
	globalShortcuts       []PageEntry
)

// === СТРУКТУРЫ ===

type Slot struct {
	SlotID        int    `json:"slotId"`
	ShortcutType  string `json:"shortcutType"`
	ShortcutID    int    `json:"shortcutId"`
	ShortcutLevel int    `json:"shortcutLevel,omitempty"`
}

type PageEntry struct {
	PageID int    `json:"pageId"`
	Slots  []Slot `json:"slots"`
}

type ClassShortcuts struct {
	ClassID int         `json:"classId"`
	Pages   []PageEntry `json:"pages"`
}

type shortcutsFileFormat struct {
	Shortcuts struct {
		Pages   []PageEntry      `json:"pages"`
		Classes []ClassShortcuts `json:"classes"`
	} `json:"shortcuts"`
}

// === ЗАГРУЗКА ===

func LoadShortcuts() {
	logger.Info.Println("Загрузка shortcuts из нового JSON формата")

	file, err := os.Open("./datapack/data/stats/char/initial/shortcuts.json")
	if err != nil {
		logger.Error.Panicln("Не удалось открыть shortcuts.json:", err)
	}
	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		logger.Error.Panicln("Ошибка чтения файла shortcuts.json:", err)
	}

	var data shortcutsFileFormat
	if err := json.Unmarshal(byteValue, &data); err != nil {
		logger.Error.Panicln("Ошибка при разборе JSON:", err)
	}

	// Инициализация хранилищ
	allShortcutsByClassID = make(map[int][]PageEntry)
	globalShortcuts = data.Shortcuts.Pages

	for _, classShortcut := range data.Shortcuts.Classes {
		allShortcutsByClassID[classShortcut.ClassID] = classShortcut.Pages
	}

	log.Println("✅ Shortcuts успешно загружены")
}

// === ДОСТУП ===

func GetShortcutsByClassID(classID int) (PageOrClassData, bool) {
	if classID == 0 {
		// Глобальные ярлыки
		return PageOrClassData{ClassID: 0, Pages: globalShortcuts}, true
	}
	pages, ok := allShortcutsByClassID[classID]
	if !ok {
		return PageOrClassData{}, false
	}
	return PageOrClassData{ClassID: classID, Pages: pages}, true
}

// === ОБЁРТКА ДЛЯ УНИФИКАЦИИ ===

type PageOrClassData struct {
	ClassID int
	Pages   []PageEntry
}
