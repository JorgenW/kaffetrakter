#include <pthread.h>
#include <stdio.h>

int i = 0;


void* adder(){
    for(int x = 0; x < 1000000; x++){
        i++;
    }
    return NULL;
}
void* adder2(){
    for(int x = 0; x < 1000000; x++){
        i--;
    }
    return NULL;
}


int main(){
    pthread_t adder_thr;
    pthread_create(&adder_thr, NULL, adder, NULL);
    
    pthread_t adder_thr2;
    pthread_create(&adder_thr2, NULL, adder2,NULL);

    
    pthread_join(adder_thr, NULL);
    pthread_join(adder_thr2,NULL);
    printf("Done: %i\n", i);
    return 0;
    
}
