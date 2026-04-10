package cmd

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type Logger struct {
	Text   string
	TimeCr time.Time
	Err    error
}

type Logs struct {
	logList []Logger
}

var (
	NilNameErr = errors.New("имя файла пустое")
	mutex      sync.Mutex
)

func (l *Logs) Log(message string, err error) {
	mutex.Lock()
	l.logList = append(l.logList, Logger{
		Text:   message,
		TimeCr: time.Now(),
		Err:    err,
	})
	defer mutex.Unlock()
}

func (l Logs) ListLogger() {
	fmt.Println("Список логов:")
	for _, i := range l.logList {
		switch i.Err {
		case nil:
			fmt.Println(i.Text, "-", i.TimeCr.Format("2001-01-02"))
		default:
			fmt.Println(i.Text, "-", i.Err, i.TimeCr.Format("2001-01-02"))
		}
	}
}
