Implemente un paquete en Go llamado rendez que implemente la primitiva rendezvous. El paquete debe proporcionar la operación Rendezvous con la siguiente cabecera:
func Rendezvous(tag int, val interface{}) interface{}
El paquete sólo puede usar sync.Mutex y sync.WaitGroup. 
El paquete debe incluir tests unitarios de prueba. 
