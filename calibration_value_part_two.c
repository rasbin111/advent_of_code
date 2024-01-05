#include <stdio.h>
#include <string.h>
#include <stdlib.h>

/* 
Question: The newly-improved calibration document consists of lines of text;
 each line originally contained a specific calibration value that the Elves 
 now need to recover. On each line, the calibration value can be found by 
 combining the first digit and the last digit (in that order) to form a 
 single two-digit number.Consider your entire calibration document. 
 What is the sum of all of the calibration values?
*/

int two_digit_calibrated_values(char *s);
int get_number(char *str);
int is_a_number(char ch);

int main(void){
    FILE *input_file = fopen("calibration_input.txt", "r");
    if (input_file == NULL){
        return 1;
    }
    int calibartion_value;
    int total_calibration_values = 0;
    char data[100];
    while(fgets(data, 200, input_file) != NULL){
        calibartion_value = two_digit_calibrated_values(data);
        total_calibration_values += calibartion_value;
    }
    printf("Sum of all of the calibration values: %d \n", total_calibration_values);
    
    fclose(input_file);

    return 0;
}

int is_a_number(char ch){
    int is_number = 0; 

    if (ch=='0' || ch=='1' || ch=='2' || ch=='3' || ch=='4' 
    || ch=='5' || ch=='6' || ch=='7' || ch=='8' || ch=='9'){
        is_number = 1;
        }

    return is_number;     // 0 represents false and 1 represents true

}

typedef struct number{
    int num;
    int position;
};


int two_digit_calibrated_values(char *str){
    int first_number;
    int last_number;
    int counter = 0;
    int number_position=0;

    if (strcasestr(str, "one") != NULL){
        
    }

    if (strcasestr(str, "two") != NULL){

    }

    if (strcasestr(str, "three") != NULL){ 

    }

    if (strcasestr(str, "four") != NULL){

    }

    if (strcasestr(str, "five") != NULL){ 

    }

    if (strcasestr(str, "six") != NULL){

    }

    if (strcasestr(str, "seven") != NULL){ 

    }

    if (strcasestr(str, "eight") != NULL){

    }

    if (strcasestr(str, "nine") != NULL){ 

    }

    for (int i=0;i<strlen(str);i++){
         is_a_number(str[i]);
        if (is_a_number(str[i]) == 1){
            if (counter == 0){
                first_number = str[i] - '0'; // converts char to int
            }
            last_number = str[i] - '0';
            counter ++;
        }
    }

    return first_number * 10 + last_number;

}
    
   
