#include <iostream>
#include <random>
#include <pthread.h>
#include <sys/time.h>

#define ALLOTTED_TIME 10000
#define THREADS 24
#define BATCH_SIZE 512

using namespace std;

struct WorkerInput {
    long circlePoints;
    long totalPoints;
};

//From StackOverflow
int64_t currentTimeMillis() {
  struct timeval time;
  gettimeofday(&time, NULL);
  int64_t s1 = (int64_t)(time.tv_sec) * 1000;
  int64_t s2 = (time.tv_usec / 1000);
  return s1 + s2;
}

//Checks if point is in the circle
bool inCircle(double x, double y) {
    if (x*x + y*y <= 1.0) {
        return true;
    }
    return false;
}

//Drawing the points
void *estimate(void *inputPtr) {
    WorkerInput *input = (WorkerInput *) inputPtr;
    
    random_device rd;
    default_random_engine generator(rd());
    uniform_real_distribution<double> distribution(-1.0, 1.0); 
    
    long start = currentTimeMillis();
    while (currentTimeMillis() - start < ALLOTTED_TIME) {
        input->totalPoints += BATCH_SIZE;
#pragma unroll BATCH_SIZE
        for (long i = 0; i < BATCH_SIZE; i++) {
            double x = distribution(generator);
            double y = distribution(generator);
            if (inCircle(x, y)) {
                input->circlePoints++;
            }
        }
    }

    return NULL;
}

int main() {
    //Creating threads
    pthread_t tids[THREADS];
    WorkerInput *inputs[THREADS];
    for (int i = 0; i < THREADS; i++) {
        inputs[i] = (WorkerInput *) calloc(sizeof(WorkerInput), 1);
        pthread_create(&(tids[i]), NULL, estimate, inputs[i]);
    }
    
    long circle = 0;
    long total = 0;
    for (int i = 0; i < THREADS; i++) {
        pthread_join(tids[i], NULL);
        circle += inputs[i]->circlePoints;
        total += inputs[i]->totalPoints;
        free(inputs[i]);
    }
    
    double result = 4.0 * (double) circle / (double) total;
    printf("Pi ~ %.12f\nTotal = %ld\n", result, total);
    return 0;
}
