
// POSIX API: pthreads
#include <pthread.h>
int
pthread_create(
    pthread_t *thread,
    const pthread_attr_t *attr,
    void * (*start_routine)(void*),
    void * arg,
)

// thread: pointer to structure pthread_t (initialized by pthread_create())
// attr: specicy attributes of this thread
// - stack size
// - scheduling priority
// - initialized by a separate call pthread_attr_init()
// - most cases: defaults are fine
//   - simply pass: NULL
// - see manual pages for details
// function pointer - which function should this thread start running
// - a function name (start_routine)
// - a single argument of type void * (as indicated in the parens after start_routine)
// - returns a value of type void *
// arg: the argument to be passed to the function

// Void pointer (void *):
// - allow passing any type as argument and
// - return any type as a result.

int pthread_join(
    pthread_t thread,
    void **value_ptr
);

// - pthread_t: the thread to wait for
// - this variable is initialized by pthread_create()
// - pointer to the return value you expect to get back

