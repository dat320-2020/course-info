#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int main(int argc, char *argv[]) {
    // copy string from source pointer to destination pointer
    char *src = "hello";  // allocated on the stack
    char *dst = (char *) malloc(strlen(src));  // allocated on the heap
    strcpy(dst, src);
    free(src);
    printf("Hey: %s\n", dst);
    free(dst);
}
