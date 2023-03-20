#include <unistd.h>

static int fork_wrapper() {
    return fork();
}