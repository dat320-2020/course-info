int buffer[MAX];
int fill = 0;
int use = 0;

void put(int value) {
    buffer[fill] = value;       // Line F1
    fill = (fill + 1) % MAX;    // Line F2
}

int get() {
    int tmp = buffer[use];  // Line G1
    use = (use + 1) % MAX;  // Line G2
    return tmp;
}

sem_t empty;
sem_t full;

void * producer(void *arg) {
    int i;
    for (i = 0; i < loops; i++) {
        sem_wait(&empty);           // Line P1
        sem_wait(&mutex);           // Lock (P0)
        put(i);                     // Line P2
        sem_post(&mutex);           // Unlock
        sem_post(&full);            // Line P3
    }
}

void * consumer(void *arg) {
    int i, tmp = 0;
    while (tmp != -1) {
        sem_wait(&full);            // Line C1
        sem_wait(&mutex);           // Lock (C0)
        tmp = get();                // Line C2
        sem_post(&mutex);           // Unlock
        sem_post(&empty);           // Line C3
        printf("%d\n", tmp);
    }
}

int main() {
    sem_init(&mutex, 0, 1);         // mutex=1
    sem_init(&empty, 0, MAX);       // MAX buffers are empty to begin with
    sem_init(&full, 0, 0);          // there are 0 full buffers
}
