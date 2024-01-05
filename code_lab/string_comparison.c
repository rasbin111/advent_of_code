#include <stdio.h>
#include <string.h>

int main(){
    char *name = "Rasbin";
    char *another_name = "asbin";
    int result = strcmp(name, another_name);
    printf("%d", result);
    return 0;
}