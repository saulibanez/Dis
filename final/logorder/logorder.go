package main

import (
	"bufio"
	"final/logicclock"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
)

type Events struct {
	mc []logicclock.MessageClock
}

func (ev *Events) Len() int {
	return len(ev.mc)
}

func (ev *Events) Swap(i, j int) {
	ev.mc[i], ev.mc[j] = ev.mc[j], ev.mc[i]
}

func goFirst(pos int, ev *Events) float64 {
	var result float64 = 0

	for _, value := range ev.mc[pos].Map {
		result += math.Pow(float64(value), 2)
	}

	return math.Sqrt(result)
}

func (ev *Events) Less(i, j int) bool {
	range_i := goFirst(i, ev)
	range_j := goFirst(j, ev)

	return range_i < range_j
}

func readFile(file string) {
	f, err := os.Open(file)

	if err != nil {
		log.Fatal(err)
	}

	ev := new(Events)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		deserialize := logicclock.Deserialize(scanner.Text())
		ev.mc = append(ev.mc, deserialize)
	}

	sort.Sort(ev)
	for _, event := range ev.mc {
		fmt.Println("Log: " + event.Log + ", Id: " + event.Id)
	}

	f.Close()
}

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		log.Fatal("You should introduced name file")
	}

	for i := 0; i < len(args); i++ {
		fmt.Println(args[i])
		readFile(args[i])
	}
}
