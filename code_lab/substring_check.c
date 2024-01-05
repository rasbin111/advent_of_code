#include <stdio.h>
#include <string.h>

int main(void){
    char *sentence = "Hello from the other side";
    char *word = "ello";
    char *pch = strstr(sentence, word); // return pointer to start of sentence if contains, else null
    int position = pch - sentence; // position of substing
    printf("%d", position);
    return 0;
}