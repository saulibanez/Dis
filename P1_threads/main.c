/*
* Saul Ibañez Cerro
* Grado en Telematica
*/

#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include "threads.h"

void
f4(void *p)
{
	printf("This is thread %i\n", curidthread());
	usleep(50000);
	printf("Back to thread %i\n", curidthread());
	exitsthread();
}

void
f2(void *p)
{
	int i;
	int arg = 1;
	int stacksize = 16*1024;
	
	printf("This is thread %i\n", curidthread());
	for (i = 0; i<2; i++){
		if (createthread(f4, &arg, stacksize) < 0){
			fprintf(stderr, "failed to create thread\n");
		}
	}
	usleep(50000);
	printf("Leaving from thread %i\n", curidthread());
	yieldthread();
	printf("Back to thread %i\n", curidthread());
	exitsthread();
}

void
f1(void *p)
{ 
	int i;
	
	printf("This is thread %i\n", curidthread());
	for(i=0; i < 4; i++){
		usleep(50000);
		printf("Leaving from thread %i\n", curidthread());
		yieldthread();
	}
	printf("Back to thread %i\n", curidthread());
	exitsthread();

}

void
f3(void *p)
{
	int i;
	int arg = 1;
	int stacksize = 16*1024;
	
	printf("This is thread %i\n", curidthread());
	for (i=0; i<32; i++){
		if (createthread(f1, &arg, stacksize) < 0){
			fprintf(stderr, "failed to create thread\n");
		}
	}

	usleep(50000);
	printf("Leaving from thread %i\n", curidthread());
	yieldthread();
	printf("Back to thread %i\n", curidthread());
	exitsthread();
}

int
main(int argc, char *argv[])
{
	int arg = 1;
	int stacksize = 16*1024;	
	initthreads();

	if (createthread(f3, &arg , stacksize) < 0){
		fprintf(stderr, "failed to create thread\n");
	}

	if (createthread(f2, &arg , stacksize) < 0){
		fprintf(stderr, "failed to create thread\n");
	}

	if (createthread(f1, &arg , stacksize) < 0){
		fprintf(stderr, "failed to create thread\n");
	}
	
	printf("This is thread %i\n", curidthread());
	for(;;){
		sleep(1);
		printf("Leaving from thread %i\n", curidthread());
		yieldthread();
		printf("Back to thread %i\n", curidthread());
		exitsthread();
	}

	exit(EXIT_SUCCESS);
}
