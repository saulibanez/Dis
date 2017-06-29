Implementa tu propia librería de semáforos con variables condición en un paquete llamado sem.

Cada puntero a semáforo tiene que tener dos métodos, up y down, para las dos operaciones de semáforos y una función para crearlo.

func NewSem(ntok int) *Sem
func (s *Sem) Up()
func (s *Sem) Down()

La librería tiene que ofrecer un interfaz UpDowner con los métodos del semáforo y tests.
Utiliza tu librería de semáforos para implementar un programa llamado factory.go que resuelva el siguiente problema en un programa main:

Hay cuatro líneas de montajes, que en tu programa se simularán con sus correspondientes gorutinas. Una línea de montaje trae pantallas, otra carcasas, cables y  placas madre. Antes de que el robot soldador monte cada caja, se tiene que asegurar de que hay 5 cables, una pantalla una carcasa y una placa madre. Hay tres robots soldadores que sueldan los móviles. Los robots soldadores tardan 200ms en soldar el móvil y pueden soldar los móviles en paralelo (cada uno el suyo). Tu programa es el encargado de dar las órdenes a los robots para soldar. Cada pieza tiene un identificador (un contador) y tu programa tiene que escribir por pantalla los identificadores del robot,  de las piezas y la acción.

Por ejemplo:

robot 0, cables 1 0 3 4 5, pantalla 3, carcasa 1, placa 25 Comenzando
robot 0, cables 1 0 3 4 5, pantalla 3, carcasa 1, placa 25 Terminado

Es importante que los robots se gasten por igual, no se puede tener uno todo el tiempo activo y el otro no.
