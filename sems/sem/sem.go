/*
* Saúl Ibáñez Cerro
* Grado en Ingeniería Telemática
 */

package sem

import (
	"sync"
)

type UpDowner interface {
	NewSem(ntok int) *Sem
	Up()
	Down()
}

type Sem struct {
	value int
	sync.Mutex
	c *sync.Cond
}

func NewSem(ntok int) *Sem {
	if ntok < 0 {
		panic("No se puede crear el semaforo")
	} else {
		semaphore := Sem{}
		semaphore.value = ntok
		semaphore.c = sync.NewCond(&semaphore)
		return &semaphore
	}
	return nil
}

// Equivale a hacer un Unlock
func (s *Sem) Up() {
	s.Lock()
	s.value++
	s.c.Signal()
	s.Unlock()
}

// Equivale a hacer un Lock
func (s *Sem) Down() {
	s.Lock()
	if s.value > 0 {
		s.value--
	} else {
		s.c.Wait()
	}
	s.Unlock()
}
