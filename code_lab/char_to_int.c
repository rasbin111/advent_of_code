#include <stdio.h>

int main(){
    char ch;
    printf("Please enter a char (0-9): ");
    scanf("%c", &ch);

    int ch_num = ch - '0';
    printf("char %c -> %d \n", ch, ch_num);
}