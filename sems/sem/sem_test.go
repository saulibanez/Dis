/*
* Saúl Ibáñez Cerro
* Grado en Ingeniería Telemática
 */

package sem

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"testing"
	"time"
)

const (
	MaxCount = 50
)

var wg *sync.WaitGroup

func TestSem1(t *testing.T) {
	semaphore := NewSem(0)
	wg = &sync.WaitGroup{}

	wg.Add(1)
	var count_test = 0

	go func() {
		for i := 0; i < MaxCount; i++ {
			semaphore.Down()
			count_test++
		}
		wg.Done()

	}()

	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

	for i := 0; i < MaxCount; i++ {
		semaphore.Up()
	}

	wg.Wait()
	if count_test != MaxCount {
		fmt.Fprintf(os.Stderr, "The values of MaxCount: %d and count_test: %d, don't match.\n", MaxCount, count_test)
		t.Error()
	}
}

func TestSem(t *testing.T) {
	semaphore := NewSem(1)
	var count_test = 0

	for i := 0; i < MaxCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			semaphore.Down()
			count_test++
			semaphore.Up()
		}()
	}

	wg.Wait()
	if count_test != MaxCount {
		fmt.Fprintf(os.Stderr, "The values of MaxCount: %d and count_test: %d, don't match.\n", MaxCount, count_test)
		t.Error()
	}
}
