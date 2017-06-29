/*
* Saúl Ibáñez Cerro
* Grado en Ingeniería Telemática
 */

package rendezvous

import (
	"fmt"
	"os"
	"sync"
	"testing"
	"time"
)

var wg = &sync.WaitGroup{}

func sleepRendez(t *testing.T, value string) {

	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
	}

	Rendezvous(0, value)
}

func goWg(name string, tag int, t *testing.T) {

	fmt.Printf("Begin: %s%d, tag: %d\n", name, 1, tag)
	request := Rendezvous(tag, name)
	fmt.Printf("End: %s%d, tag: %d\n", name, 2, tag)
	wg.Done()

	if name != request {
		fmt.Fprintln(os.Stderr, "No coincidence")
		t.Error()
	}
}

func TestRendezvous(t *testing.T) {
	n_max := 10
	aux := 4

	for i := 0; i < n_max; i++ {
		wg.Add(1)
		if i > 4 {
			name := string('a' + aux)
			go goWg(name, aux, t)
			aux--
			continue
		}

		name := string('a' + i)
		go goWg(name, i, t)
	}

	wg.Wait()
}

func TestRendezvousTime(t *testing.T) {
	value := "continue..."

	go sleepRendez(t, value)
	couple_value := Rendezvous(0, "wait routine...")

	if couple_value != value {
		t.Error()
	}
}
