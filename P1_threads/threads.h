/*
* Saul Ibañez Cerro
* Grado en Telematica
*/

#include <ucontext.h>
#include <sys/time.h>

typedef enum {UNUSED, RUNNING, EXITED, WAITING} threadStatus;

typedef struct Threadlist {
	ucontext_t context;
	int id;
	void *stack;
	threadStatus state;
	time_t th_msec;
}Threadlist;



void initthreads(void);
int createthread(void (*mainf)(void*), void *arg, int stacksize);
void exitsthread(void);
void yieldthread(void);
int curidthread(void);
