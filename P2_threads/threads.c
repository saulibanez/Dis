/*
* Saul Ibañez Cerro
* Grado en Telematica
*/

#include <stdio.h>
#include <stdlib.h>
#include "threads.h"
#include <err.h>

enum {
	MAXTHREADS = 32,
	MILISEC = 200
};

enum {
	UNUSED = 0, 
	RUNNING = 1, 
	EXITED = 2, 
	WAITING = 3, 
	SUSPEND = 4, 
	SLEEP = 5
};
typedef struct Threadlist Threadlist;

struct Threadlist {
	ucontext_t context;
	int id;
	void *stack;
	int state;
	long th_msec;
	long fall_sleep;
};

Threadlist ttable[MAXTHREADS];
int state_th[MAXTHREADS];

static int flag = 0;
static int th_current = 0;
static int flag_sleep = 0;

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
		err(1, "You can only be executed once the function initthreads");
	}

	ttable[0].id = 0;
	ttable[0].stack = NULL;
	ttable[0].state = RUNNING;
	if(getcontext(&ttable[0].context) < 0){
		err(1, "failure initthreads, getcontext");
	}

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
wakeUp(int pos)
{
	long now;
	now = getTimeMs();
	int is_sleep = 0;

	if(ttable[pos].state == SLEEP && now > ttable[pos].fall_sleep){
		ttable[pos].state = WAITING;
		is_sleep = 1;
	}

	return is_sleep;
}

static int
exitLibrary(int next)
{
	int fail = -1;

	switch((int)ttable[next].state) {
	case RUNNING:
		return next;
	case SUSPEND:
		errx(1, "Fail, the last thread is suspended, the program can not continue.\n");
	case WAITING:
		errx(1, "Fail, the last thread is waiting, the program can not continue.\n");
	case SLEEP:
		if(!flag_sleep){
			errx(1, "Fail, the last thread wants to sleep and threads aren't asleep, the program can not continue.\n");
		}
	case EXITED:
		if(flag_sleep){
			return fail;
		}

		printf("The last thread is exited, the program has ended.\n");
		exit(0);
	}
	return next;
}

static int
planificador(void)
{
	int next = (th_current+1)%MAXTHREADS;
	flag_sleep = 0;

	while(next >= 0){
		if(wakeUp(next) && flag_sleep > 0){
			flag_sleep--;
		}

		if(th_current != next){
			thExit(next);
		}
		state_th[next]=ttable[next].state;
		if(ttable[next].state == SLEEP){
			flag_sleep = 1;
		}else if(ttable[next].state == WAITING){
			return next;
		}else if(ttable[next].state == SUSPEND){
			if(state_th[next] == SUSPEND && th_current == next && !flag_sleep){
				exitLibrary(next);
			}
			else{
				next = (next+1)%MAXTHREADS;
				continue;
			}
		}
		if(th_current == next){
			exitLibrary(next);
		}

		if(ttable[next].state == RUNNING){
			return next;
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


	if(changeTh() || ttable[th_current].state == EXITED || ttable[th_current].state == SUSPEND || ttable[th_current].state == SLEEP){
		ttable[next].state = RUNNING;
		if(ttable[th_current].state == RUNNING){
			ttable[th_current].state = WAITING;
		}
		last_th_current = th_current;
		th_current = next;
		ttable[next].th_msec = getTimeMs();


		if(next==MAXTHREADS){
			next=0;
		}
		swapcontext(&ttable[last_th_current].context, &ttable[next].context);
	}
}

int 
curidthread(void)
{
	return th_current;
}

void 
suspendthread(void)
{
	ttable[th_current].state = SUSPEND;
	yieldthread();
}

int 
resumethread(int id)
{
	int fail = -1;
	if(ttable[id].state == SUSPEND){
		ttable[id].state = WAITING;
		return 0;
	}

	return fail;
}

int 
suspendedthreads(int **list)
{
	int count = 0;
	int i;
	int *id_list_th;

	id_list_th = (int*)malloc(MAXTHREADS*sizeof(int));
	for (i = 0; i < MAXTHREADS; i++)
	{
		if(ttable[i].state == SUSPEND){
			id_list_th[count] = ttable[i].id;
			count++;
		}
	}

	*list = id_list_th;

	return count;
}

int 
killthread(int id)
{
	int fail = -1;

	if(ttable[id].state == EXITED || ttable[id].state == UNUSED || ttable[id].state == RUNNING){
		return fail;
	}
	
	ttable[id].state = EXITED;
	return id;
}

void 
sleepthread(int msec)
{
	long sum;
	sum = getTimeMs() + msec;

	ttable[th_current].fall_sleep = sum;
	ttable[th_current].state = SLEEP;
	yieldthread();
}
