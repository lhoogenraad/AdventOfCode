#include <stdlib.h>
#include <stdio.h>
#include <getopt.h>
#include <string.h>
#include <ctype.h>



char * getMaxMin(char str[]){
	char max = '0', min = '0';
	int i;
	int foundMin = 0;
	printf("%s\n", str);
	for(i = 0; i < strlen(str); i++){
		char currChar = str[i];
		if(isdigit(currChar)){
			if(!foundMin){
				min = currChar ;
				foundMin = 1;
			}
			max = currChar;
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
		maxMin = getMaxMin(line);
		printf("Min: %c, Max: %c\n", maxMin[0], maxMin[1]);
		int value = (maxMin[0] - '0') * 10 + (maxMin[1] - '0');	
		printf("Value: %d\n", value);
		total += value;
	}
	fclose(fptr);
	if(line){
		free(line);
	}


	printf("\nTotal: %d\n", total);


	exit(EXIT_SUCCESS);
}
