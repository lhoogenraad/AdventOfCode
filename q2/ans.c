#include <stdlib.h>
#include <stdio.h>
#include <getopt.h>
#include <string.h>
#include <ctype.h>
#include <assert.h>


const BMAX = 14;
const RMAX = 12;
const GMAX = 13;


int split (const char *str, char c, char ***arr)
{
    int count = 1;
    int token_len = 1;
    int i = 0;
    char *p;
    char *t;

    p = str;
    while (*p != '\0')
    {
        if (*p == c)
            count++;
        p++;
    }

    *arr = (char**) malloc(sizeof(char*) * count);
    if (*arr == NULL)
        exit(1);

    p = str;
    while (*p != '\0')
    {
        if (*p == c)
        {
            (*arr)[i] = (char*) malloc( sizeof(char) * token_len );
            if ((*arr)[i] == NULL)
                exit(1);

            token_len = 0;
            i++;
        }
        p++;
        token_len++;
    }
    (*arr)[i] = (char*) malloc( sizeof(char) * token_len );
    if ((*arr)[i] == NULL)
        exit(1);

    i = 0;
    p = str;
    t = ((*arr)[i]);
    while (*p != '\0')
    {
        if (*p != c && *p != '\0')
        {
            *t = *p;
            t++;
        }
        else
        {
            *t = '\0';
            i++;
            t = ((*arr)[i]);
        }
        p++;
    }

    return count;
}



int validate_line (char* game){
	char **rounds;
	int num_games = split(game, ';', &rounds);
	printf("num_games: %d\n", num_games);
	int i;
	for (i = 0; i < num_games; i++){
		int r = 0, g = 0, b = 0;
		char **cube_counts;
		int num_colours = split(rounds[i], ',' , &cube_counts);
		printf("num_colours for game %d: %d\n", i + 1, num_colours);
		int x;
		for (x = 0; x < num_colours; x++){
			printf("cube counts for cube colour number %d is: %s\n", x, cube_counts[x]);
			char **cube_info;
			int num_strings = split(cube_counts[x], ' ', &cube_info);
			// This is the number of cubes for this type of cube in this game.
			int cubes = atoi(cube_info[0]);
			printf("cube count: %s, colour: %s\n", cube_info[0], cube_info[1]);
			if(strcmp("red", cube_info[1]) == 0){
				if (cubes > RMAX){
					return 0;
				}
			}
			if(strcmp("green", cube_info[1]) == 0){
				if(cubes > GMAX){
					return 0;
				}
			}
			if(strcmp("blue", cube_info[1]) == 0){
				if(cubes > BMAX){
					return 0;
				}
			}
		}
	}

	free(rounds);
	return 1;
}



int main (int argc, char *argv[]){

	int valid = validate_line("3 blue, 2 green, 6 red; 17 green, 4 red, 8 blue; 2 red, 1 green, 10 blue; 1 blue, 5 green");
	printf("valid?: %s", valid);

	exit(EXIT_SUCCESS);
}
