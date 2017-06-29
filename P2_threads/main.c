/*
* Saul Ibañez Cerro
* Grado en Telematica
*/

#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include "threads.h"

void
threadSleep(void *p)
{
	fprintf(stderr, "This is thread %i, this thread is going to sleep \n", curidthread());
	sleepthread(10000);
	fprintf(stderr, "Back to thread that was sleeping with id %i\n", curidthread());
	exitsthread();
}

void
f1(void *p)
{ 
	int i;
	
	fprintf(stderr, "This is thread %i\n", curidthread());

	for(i=0; i < 4; i++){
		usleep(50000);
		fprintf(stderr, "Leaving from thread %i\n", curidthread());
		yieldthread();
	}
	fprintf(stderr, "Back to thread %i\n", curidthread());
	exitsthread();

}

int 
printList(int* list, int num){

	int i;
	int th_id;
	if (num > 0){
		fprintf(stderr, "The id of thread suspended is:");
		for (i = 0; i < num; i++){
			if(i == (num-1)){
				fprintf(stderr, " %i.", list[i]);
			}else{
				fprintf(stderr, " %i,", list[i]);
				th_id = list[i];
			}
		}
		fprintf(stderr, "\n");
	}else{
		fprintf(stderr, "The list of thread suspended is void\n");
	}

	return th_id;
}

void
threadSuspended(void *p)
{
	int resume;
	int *list;

	fprintf(stderr, "This is thread %i, this thread is going to be suspended\n", curidthread());
	suspendthread();

	fprintf(stderr, "The thread with id: %i, suspended now may return to execute \n", curidthread());
	resume = suspendedthreads(&list);
	printList(list, resume);

	fprintf(stderr, "Back to thread that was suspended with id: %i\n", curidthread());
	exitsthread();
}

void
morethreads(void *p)
{
	int i;
	int arg = 1;
	int stacksize = 16*1024;

	fprintf(stderr, "This is thread %i\n", curidthread());
	for (i=0; i<30; i++){
		if(i==18){
			if (createthread(threadSleep, &arg, stacksize) < 0){
				fprintf(stderr, "failed to create thread\n");
			}
		}else if((i==9) || (i==13) || (i==26)){
			if (createthread(threadSuspended, &arg, stacksize) < 0){
				fprintf(stderr, "failed to create thread\n");
			}
		}else if (createthread(f1, &arg, stacksize) < 0){
			fprintf(stderr, "failed to create thread\n");
		}

	}

	fprintf(stderr, "I'm gonna kill the thread with id: %i\n", 11);
	if(killthread(11) < 0){
		fprintf(stderr, "failed to create thread with id: %i\n", 11);
	}else{
		fprintf(stderr, "Back to the killed thread with id: %i\n", 11);
	}
	exitsthread();

	usleep(50000);
	fprintf(stderr, "Leaving from thread %i\n", curidthread());
	yieldthread();
	fprintf(stderr, "Back to thread %i\n", curidthread());
	exitsthread();
}

int
main(int argc, char *argv[])
{
	int arg = 1;
	int stacksize = 16*1024;	
	initthreads();
	int resume;
	int *list;
	int th_id_susp;

	if (createthread(morethreads, &arg , stacksize) < 0){
		fprintf(stderr, "failed to create thread\n");
	}

	if (createthread(f1, &arg , stacksize) < 0){
		fprintf(stderr, "failed to create thread\n");
	}

	if (createthread(threadSleep, &arg , stacksize) < 0){
		fprintf(stderr, "failed to create thread\n");
	}
	
	fprintf(stderr, "This is thread %i\n", curidthread());
	for(;;){
		sleep(1);
		fprintf(stderr, "Leaving from thread %i\n", curidthread());
		yieldthread();

		resume = suspendedthreads(&list);
		th_id_susp = printList(list, resume);
		resumethread(th_id_susp);
		fprintf(stderr, "Back to thread %i\n", curidthread());
		exitsthread();
	}

	exit(EXIT_SUCCESS);
}
