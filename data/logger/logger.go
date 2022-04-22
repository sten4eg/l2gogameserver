package logger

//Логирование всех данных
//TODO: в будущем ERROR сохранять в файл.

import (
	"log"
	"os"
)

var (
	Warning *log.Logger
	Info    *log.Logger
	Error   *log.Logger
)

func init() {
	Info = log.New(os.Stdout, "INFO: ", log.Ltime|log.Lshortfile)
	Warning = log.New(os.Stdout, "WARNING: ", log.Ltime|log.Lshortfile)
	Error = log.New(os.Stdout, "ERROR: ", log.Ltime|log.Lshortfile)
}
