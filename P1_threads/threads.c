/*
* Saul Ibañez Cerro
* Grado en Telematica
*/

#include <stdio.h>
#include <stdlib.h>
#include "threads.h"
#include <err.h>

#define MAXTHREADS 32
#define MILISEC 200
static int flag = 0;
static int th_current = 0;
Threadlist ttable[MAXTHREADS];
int state_th[MAXTHREADS];

static void
thExit(int pos)
{
	if(ttable[pos].state == EXITED){
		free(ttable[pos].stack);
		ttable[pos].state = UNUSED;
	}
}


// return long ms
static long
getTimeMs()
{
	struct timeval t_aux;
	long now;
	gettimeofday(&t_aux, NULL);
	now = t_aux.tv_sec*1000 + t_aux.tv_usec/1000;
	return now;
}

void 
initthreads(void)
{	
	int i;

	if(flag){
		//err(1, "You can only be executed once the function initthreads");
		fprintf(stderr , "%s\n", "Threads are just initialized...");
        return;
	}

	ttable[0].id = 0;
	ttable[0].stack = NULL;
	ttable[0].state = RUNNING;
	getcontext(&ttable[0].context);

	for(i=1;i<MAXTHREADS;i++){
		ttable[i].state = UNUSED;
	}
	ttable[th_current].th_msec = getTimeMs();

	flag = 1;
}

int 
createthread(void (*mainf)(void*), void *arg, int stacksize)
{
	int i;
	int fail = -1;
	void *stack = NULL;


	for (i=0; i<MAXTHREADS;i++){
		thExit(i);
		if(ttable[i].state == UNUSED){
			fail = i;
			break;
		}		
	}

	//fprintf(stderr, "state -> %i\n", ttable[i].state);


	if(fail < 0){
		return fail;
	}

	if (getcontext(&ttable[i].context) < 0){
		err(1, "failure createthread, getcontext");
	}
	
	ttable[i].id = i;
	stack = malloc(stacksize);
	ttable[i].context.uc_stack.ss_sp = stack;
	ttable[i].context.uc_stack.ss_size = stacksize;
	ttable[i].stack = stack;
	ttable[i].state = WAITING;
	makecontext(&ttable[i].context, (void(*)(void))mainf, 1, arg);

	return i;
}

void 
exitsthread(void)
{
	ttable[th_current].state = EXITED;
	yieldthread();

}

static int
planificador(void)
{
	int next = (th_current+1)%MAXTHREADS;

	while(next >= 0){

		if(th_current != next){
			thExit(next);
		}

		state_th[next]=ttable[next].state;
		
		if(ttable[next].state == WAITING){
			return next;
		}
		
		if(ttable[next].state == RUNNING){
			return next;
		}

		if(th_current == next){
			printf("The last thread is exited, the program has ended.\n");
			exit(0);
		}

		next = (next+1)%MAXTHREADS;
	}

	return next;
}

static int
changeTh(void)
{
	long dif;

	dif = getTimeMs() - ttable[th_current].th_msec;

	if(dif < MILISEC){
		return 0;
	}
	return 1;
}

void 
yieldthread(void)
{
	int next = 0;
	int last_th_current;
	
	next = planificador();
	
	if(changeTh() || ttable[th_current].state == EXITED){
		ttable[next].state = RUNNING;
		if(ttable[th_current].state == RUNNING){
			ttable[th_current].state = WAITING;
		}
		last_th_current = th_current;
		th_current = next;
		ttable[next].th_msec = getTimeMs();
		swapcontext(&ttable[last_th_current].context, &ttable[next].context);
	}
}

int 
curidthread(void)
{
	return th_current;
}
