//Пакет для чтения, сохранения HTML диалогов.
//Сохраняет в массиве только те файлы, к которым обращались.
//TODO: В будущем сделать автоочистку массива от HTML диалогов, к которым не обращались раз в N минут.

package htm

import (
	"errors"
	"l2gogameserver/config"
	"os"
	"regexp"
	"strings"
	"time"
)

// Общий массив с файлами
var htmcodes []htmlcode

type htmlcode struct {
	filename string    //Название файла
	code     *string   //HTML код
	time     time.Time //Время добавления
}

/*
Чтение файла и возвращаем как текст
В htmlcode сразу откинем лишнее.
Результат закэшим.
*/
func Open(filename string) (*string, error) {
	if !strings.HasSuffix(filename, ".htm") {
		filename += "/index.htm"
	}
	htm, ok := getHTMLCache(filename)
	if ok {
		return htm, nil
	}

	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	if len(content) == 0 {
		return nil, errors.New("html file is empty")
	}

	r := strings.NewReplacer(
		"\t", "",
		"\r", "",
		"\n", "",
	)
	contents := r.Replace(string(content))
	contents = regexp.MustCompile("\\s+").ReplaceAllString(contents, " ")
	c := &contents
	addHTMLCache(filename, c)
	return c, nil
}

// Поиск ранее сохраненного в массив html файла
func getHTMLCache(filename string) (*string, bool) {
	//Если выкл. сохранение диалогов, тогда пропускаем поиск
	if !config.IsEnableCachedHtml() {
		return nil, false
	}
	for _, htm := range htmcodes {
		if htm.filename == filename {
			return htm.code, true
		}
	}
	return nil, false
}

// Добавить HTML код в массив
func addHTMLCache(filename string, filecode *string) {
	//Если выкл. сохранение диалогов, тогда пропускаем сохранение в массиве
	if !config.IsEnableCachedHtml() {
		return
	}
	htmcodes = append(htmcodes, htmlcode{
		filename: filename,
		code:     filecode,
		time:     time.Now(),
	})
}
