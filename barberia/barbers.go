/*
* Saúl Ibáñez Cerro
* Grado en telemática
 */

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numOfBarbers = 2
	numOfSeats   = 5
	numOfClients = 20
)

var wg sync.WaitGroup

func barber(barbershop <-chan int, id int) {
	for {
		select {
		case cl := <-barbershop:
			fmt.Printf("Barbero %d: empiezo a cortar el pelo\nCliente %d: me corto el pelo\n", id, cl)
			time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
			fmt.Printf("Barbero %d: termino de cortar el pelo\nCliente %d: termino de cortarme el pelo\n", id, cl)
			wg.Done()
		default:
			fmt.Printf("Barbero %d: me duermo esperando clientes\n", id)
			time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
		}

	}
}

func client(barbershop chan<- int, id int) {
	select {
	case barbershop <- id:
		fmt.Printf("Cliente %d: me siento en la sala de espera\n", id)
		wg.Add(1)
	default:
		fmt.Printf("Cliente %d: me voy de la barberia, esta llena\n", id)
	}
}

func main() {
	barbershop := make(chan int, numOfSeats)

	fmt.Println("\nLa barbería abre sus puertas al público\n")
	for i := 1; i <= numOfBarbers; i++ {
		go barber(barbershop, i)
	}

	for i := 1; i <= numOfClients; i++ {
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
		go client(barbershop, i)
	}

	wg.Wait()
	fmt.Println("\nLa barbería cierra sus puertas al público\n")
}
