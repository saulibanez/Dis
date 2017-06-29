package logiclog

import (
	"final/logicclock"
	"fmt"
	"log"
	"os"
)

type Log struct {
	mc   *logicclock.MessageClock
	file *os.File
}

func NewLog(log_name, path string, mc *logicclock.MessageClock) *Log {
	l := &Log{}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		file, err := os.Create(path)

		if err != nil {
			log.Fatal(err)
		}

		l.file = file
	} else {
		file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

		if err != nil {
			log.Fatal(err)
		}

		l.file = file
	}

	if mc != nil {
		l.mc = mc
	} else {
		l.mc = logicclock.NewMessageClock(log_name)
	}

	return l
}

func (l *Log) Add() string {
	l.mc.Update()

	str := l.mc.Serialize()
	l.file.WriteString(str + "\n")

	return str
}

func (l *Log) Update(str string) {
	l.mc.Update()
	l.mc.Log = str
	fmt.Println(str)
	_, err := l.file.WriteString(l.mc.Serialize() + "\n")
	if err != nil {
		log.Fatal(err)
	}
}

func (l *Log) Msg(json string) {
	msg := logicclock.Deserialize(json)
	msg.Init(l.mc.Get())
	msg.Update()
	l.mc.Join(msg)
}

func (l *Log) closeFile() {

	if err := l.file.Close(); err != nil {
		log.Fatal(err)
	}
}
