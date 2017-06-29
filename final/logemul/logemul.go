package main

import (
	"final/logiclog"
	"fmt"
	"strconv"
)

func intermediario(in chan string, numLectores int, out chan bool) {
	canalesIn := make([]chan string, numLectores)
	canalesOut := make([]chan bool, numLectores)
	for i, _ := range canalesIn {
		canalesIn[i] = make(chan string)
		canalesOut[i] = make(chan bool)
		go lector(canalesIn[i], i, canalesOut[i])
	}
	for texto, abierto := <-in; abierto; texto, abierto = <-in {
		for _, c := range canalesIn {
			c <- texto
		}
	}
	for _, c := range canalesIn {
		close(c)
	}
	for _, c := range canalesOut {
		<-c
	}
	out <- true
}

func lector(in chan string, id int, out chan bool) {
	texto, canalAbierto := <-in
	log := logiclog.NewLog("logemul", "logemul.txt", nil)

	if log == nil {
		panic(log)
	}

	for ; canalAbierto; texto, canalAbierto = <-in {
		fmt.Println(id, texto)
		log.Update("th " + strconv.Itoa(id) + ": " + texto)
	}
	out <- true
}

func main() {
	saludos := make(chan string)
	despedida := make(chan bool)

	go intermediario(saludos, 3, despedida)

	saludos <- "Hola, pequeño Padawan"
	saludos <- "La estrella de la muerte está operativa"
	close(saludos)
	<-despedida
}
