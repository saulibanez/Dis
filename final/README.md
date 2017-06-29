#Práctica distribuidos: sistema de logging.

Escriba en el lenguaje de programación go, un sistema de logging que permita ordenar los logs de un sistema distribuido.

Cada vez que se genere un log, se considerará un evento. En los mensajes del sistema
distribuido, que se simularán mediante gorutinas comunicándose por canales, en el mensaje se
incluirá una marca lógica (puede usar el sistema que considere de los vistos en clase, aunque
se recomienda que implemente alguna variante de relojes vectoriales).
Cada gorutina escribirá al principio de cada línea del log la marca.

Debe escribir:

Un paquete logicclock que implemente el sistema de marcas lógicas y su serialización a texto o binaria.
Un paquete logiclog que implemente el sistema de logs, con un tipo de datos opaco para incluir en
los mensajes y la obtención de una ejecución causal a partir de un conjunto de logs.

Un programa logemul que simule el uso del paquete de logs con gorutinas comunicándose por canales.
Un programa logorder que dados como parámetros los ficheros de log de los diferentes sistema los ordene de forma causal.

Todos los paquetes deben tener sus tests.
