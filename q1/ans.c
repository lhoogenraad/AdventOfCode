#include <stdlib.h>
#include <stdio.h>
#include <getopt.h>
#include <string.h>
#include <ctype.h>

const int stringNumsLen = 9;
const char stringNums[9][6]= {
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine"
};

// Checks if any values in arr are 1.
int checkAnyValid(int *arr, int len){
	int i;
	for(i = 0; i < len; i++){
		if(arr[i] == 1){
			return 1;
		}
	}
	return 0;
}

char * getMaxMin(char str[], int partTwo){
	char max = '0', min = '0';
	int i;
	int foundMin = 0;
	for(i = 0; i < strlen(str); i++){
		char currChar = str[i];
		if(isdigit(currChar)){
			if(!foundMin){
				min = currChar ;
				foundMin = 1;
			}
			max = currChar;
		}
		// If we've still got space in the file line to find the shortest strung num
		else if(i < strlen(str) - 3 && partTwo){
			int s_num_array_i, s_num_i = 0, second_i, foundValue = 0, value = 0;
			// The indeces of this array represents which strung numbers are still valid for checking;
			int valid[] = {1, 1, 1, 1, 1, 1, 1, 1, 1};
			// Go forward each letter in 
			for(second_i = i; second_i < strlen(str) - 1 && s_num_i <= 5; second_i++){
				int currSecondChar = str[second_i];
				// For each string in stringNums
				for(s_num_array_i = 0; s_num_array_i < stringNumsLen; s_num_array_i++){
					// If we end up not matching, set this num as invalid 
					if(stringNums[s_num_array_i][s_num_i] != currSecondChar){
						valid[s_num_array_i] = 0;
					}else if (s_num_i == strlen(stringNums[s_num_array_i]) - 1){
						// If up until this point this string num is valid, set everything as found value.
						if(valid[s_num_array_i]){
							foundValue = 1;
							value = s_num_array_i + 1;
						}
					}
				}

				// If it turns out none are valid, break out of loop asap rocky
				if(!checkAnyValid(valid, 9) || foundValue){
					break;
				}

				s_num_i++;


			}

			if(foundValue){
				if(!foundMin){
					min = value + '0';
					foundMin = 1;
				}
				max = value + '0';
				// Skip to next char after this found strung num.
				i = second_i - 1;
			}
		}
	}

	char *ret = malloc(2);
	ret[0] = min;
	ret[1] = max;
	
	return ret;
}




int main (int argc, char *argv[]){
	FILE *fptr;

	int total = 0;

	fptr = fopen("data.txt", "r");
	char * line = NULL;
	size_t line_len = 0;
	ssize_t read;
	while((read = getline(&line, &line_len, fptr)) != -1){
		char *maxMin = malloc(2);
		maxMin = getMaxMin(line, 1);
		printf("line: %s", line);
		printf("min: %c, max: %c\n", maxMin[0], maxMin[1]);
		int value = (maxMin[0] - '0') * 10 + (maxMin[1] - '0');	
		printf("value: %d\n\n", value);
		total += value;
	}
	fclose(fptr);
	if(line){
		free(line);
	}


	printf("\nTotal: %d\n", total);


	exit(EXIT_SUCCESS);
}
