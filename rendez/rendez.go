/*
* Saúl Ibáñez Cerro
* Grado en Ingeniería Telemática
 */

package rendezvous

import (
	"sync"
)

type AddRndz struct {
	wg    sync.WaitGroup
	value interface{}
}

var map_rndz = make(map[int]*AddRndz)
var mutex = &sync.Mutex{}

func Rendezvous(tag int, val interface{}) interface{} {
	mutex.Lock()
	aux, found := map_rndz[tag]

	if found {
		map_rndz[tag].wg.Done()
		map_rndz[tag].value = val
		delete(map_rndz, tag)
		mutex.Unlock()
	} else {
		map_rndz[tag] = new(AddRndz)
		map_rndz[tag].wg.Add(1)
		map_rndz[tag].value = val
		aux = map_rndz[tag]
		mutex.Unlock()
		map_rndz[tag].wg.Wait()
	}
	return aux.value
}
