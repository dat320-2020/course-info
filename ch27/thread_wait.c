#include <pthread.h>
#include <stdio.h>
#include <stdlib.h>

typedef struct __myarg_t
{
    int a;
    int b;
} myarg_t;

typedef struct __myret_t
{
    int x;
    int y;
} myret_t;

void *mythread(void *arg)
{
    myarg_t *m = (myarg_t *)arg;
    printf("a=%d b=%d\n", m->a, m->b);

    myret_t *r = malloc(sizeof(myret_t));
    r->x = 1;
    r->y = 2;
    return (void *) r;
}

int main(int argc, char *argv[])
{
    pthread_t p;
    int rc;
    myarg_t args;
    args.a = 10;
    args.b = 20;
    myret_t *retval;
    rc = pthread_create(&p, NULL, mythread, &args);
    pthread_join(p, (void **) &retval);
    printf("x=%d y=%d\n", retval->x, retval->y);
    free(retval);
    return 0;
}