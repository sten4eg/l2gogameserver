package data

import (
	"io/fs"
	"l2gogameserver/data/logger"
	"path/filepath"
	"strings"
)

//Find Возвращает массив файлов с ext расширением
func Find(root, extension string) []string {
	var res []string
	if extension[0] != '.' {
		extension = "." + extension
	}

	err := filepath.WalkDir(root, func(s string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if filepath.Ext(d.Name()) == extension {
			res = append(res, s)
		}
		return nil
	})

	if err != nil {
		logger.Error.Panicln("data Find logger.Error.Panicln")
	}
	return res
}

//FileNameWithoutExtension Возвращает имя файла без расширения и пути
func FileNameWithoutExtension(fileName string) string {
	return strings.TrimSuffix(filepath.Base(fileName), filepath.Ext(fileName))
}
