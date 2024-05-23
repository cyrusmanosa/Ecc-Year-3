#include <dlfcn.h>

int main() {
    void *handle;
    void (*goprint)();
    char *error;

    handle = dlopen("share.so", RTLD_NOW);
    goprint = dlsym(handle, "goprint");
    goprint();
    dlclose(handle);
}
