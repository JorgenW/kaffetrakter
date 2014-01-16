#include <pthread.h>
#include <stdio.h>
#include <semaphore.h>
int i = 0;

sem_t mutex;
void* adder(){
	
    for(int x = 0; x < 1000000; x++){
		sem_wait(&mutex); 
        i++;
		printf("1");
		sem_post(&mutex);
    }

    return NULL;
}
void* adder2(){
	 
    for(int x = 0; x < 999999; x++){
		sem_wait(&mutex);
        i--;
		printf("2");
		sem_post(&mutex);
    }
	
    return NULL;
}


int main(){
    pthread_t adder_thr;
	pthread_t adder_thr2;

	sem_init(&mutex,0,1);

    pthread_create(&adder_thr, NULL, adder, NULL);
    pthread_create(&adder_thr2, NULL, adder2,NULL);
	
    pthread_join(adder_thr, NULL);
    pthread_join(adder_thr2,NULL);

	sem_destroy(&mutex);

    printf("Done: %i\n", i);
    return 0;
    
}
