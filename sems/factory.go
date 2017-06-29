/*
* Saúl Ibáñez Cerro
* Grado en Ingeniería Telemática
 */

package main

import (
	"fmt"
	"sems/sem"
	"strconv"
	"strings"
	"time"
)

type Pieces int

const (
	Cables Pieces = iota
	Screen
	Case
	Board
)

const (
	AssamblyLines = 4
	NumRobots     = 3
	MaxPieces     = 200
)

const (
	CablesNec   = 5
	ScreensNec  = 1
	CaseNec     = 1
	BoardNec    = 1
	ObjTotalNec = 8
)

type ProductorGlobal struct {
	ticketsem *sem.Sem
	holesem   *sem.Sem
	buf       []int
	contador  int
}

var pieces = make([]*ProductorGlobal, AssamblyLines)

func initAssamblyLine() {
	for i := 0; i < AssamblyLines; i++ {
		pieces[i] = &ProductorGlobal{sem.NewSem(0), sem.NewSem(MaxPieces), make([]int, MaxPieces), 0}
	}
}

func (p *ProductorGlobal) productorLocal() {
	i := 0
	for id := 0; ; id++ {
		p.holesem.Down()
		p.buf[i] = id
		i = (i + 1) % MaxPieces
		p.ticketsem.Up()
	}
}

func (p *ProductorGlobal) getAssamblyLine() int {
	p.ticketsem.Down()
	defer p.holesem.Up()
	id := p.buf[p.contador]
	p.contador = (p.contador + 1) % MaxPieces
	return id
}

func robot(idRobot int) {
	id_pieces := make([]int, ObjTotalNec)
	pos_cables := CablesNec
	pos_screen := pos_cables + ScreensNec
	pos_case := pos_screen + CaseNec
	pos_board := pos_case + BoardNec
	for pos := 0; ; pos = (pos + 1) % ObjTotalNec {
		switch {
		case pos < pos_cables:
			id_pieces[pos] = pieces[Cables].getAssamblyLine()
		case pos < pos_screen:
			id_pieces[pos] = pieces[Screen].getAssamblyLine()
		case pos < pos_case:
			id_pieces[pos] = pieces[Case].getAssamblyLine()
		case pos < pos_board:
			id_pieces[pos] = pieces[Board].getAssamblyLine()
			fmt.Println(printProcess(idRobot, id_pieces), "Comenzando")
			time.Sleep(200 * time.Millisecond)
			fmt.Println(printProcess(idRobot, id_pieces), "Terminado")
		}
	}
}

func printProcess(id_robot int, id_pieces []int) string {
	nprint := AssamblyLines*2 + ObjTotalNec*2
	posObj := CablesNec
	print_window := make([]string, nprint)
	i := 0
	print_window[i] = "robot " + strconv.Itoa(id_robot) + ", cables "
	i++

	for pos := 0; pos < CablesNec; pos++ {
		if pos == CablesNec-1 {
			print_window[i] = strconv.Itoa(id_pieces[pos])
			i++
		} else {
			print_window[i] = strconv.Itoa(id_pieces[pos]) + " "
			i++
		}
	}

	print_window[i] = ", pantalla " + strconv.Itoa(id_pieces[posObj]) + ", carcasa "
	i++
	posObj += CaseNec
	print_window[i] = strconv.Itoa(id_pieces[posObj]) + ", placa "
	i++
	posObj += BoardNec
	print_window[i] = strconv.Itoa(id_pieces[posObj])

	return strings.Join(print_window, "")
}

func main() {
	initAssamblyLine()
	go pieces[Cables].productorLocal()
	go pieces[Case].productorLocal()
	go pieces[Screen].productorLocal()
	go pieces[Board].productorLocal()

	for id := 0; id < NumRobots; id++ {
		go robot(id)
	}
	semaphore := sem.NewSem(0)
	semaphore.Down()
}
