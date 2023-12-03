#include <stdlib.h>
#include <stdio.h>
#include <getopt.h>
#include <string.h>
#include <ctype.h>
#include <assert.h>


const int BMAX = 14;
const int RMAX = 12;
const int GMAX = 13;

char *trim(char *str)
{
  char *end;

  // Trim leading space
  while(isspace((unsigned char)*str)) str++;

  if(*str == 0)  // All spaces?
    return str;

  // Trim trailing space
  end = str + strlen(str) - 1;
  while(end > str && isspace((unsigned char)*end)) end--;

  // Write new null terminator character
  end[1] = '\0';

  return str;
}


int validate_line (char* game){
	char * cube_split = strtok(game, ",");
	// 3 colours max
	char ** cube_entries = malloc(3);
	int i = 0;

	//Build cube entries
	while(cube_split != NULL){
		cube_entries[i] = trim(cube_split);
		printf("%s\n", cube_entries[i]);
		i++;
		cube_split = strtok(NULL, ",");
	}

	int x;
	for(x = 0; x <= i; x++){
		char * cube_info = strtok(cube_entries[x], " ");
		int num_cubes = atoi(cube_info);
		char * cube_colour = strtok(NULL, " ");

		printf("num_cubes: %d\ncube_colour: %s\n\n", num_cubes, cube_colour);
	}

	return 1;
}



int main (int argc, char *argv[]){

	char str[100] = "Game 1: 3 blue, 2 green, 6 red; 17 green, 4 red, 8 blue; 2 red, 1 green, 10 blue; 1 blue, 5 green";
	char * game_id = strtok(str, ":");
	char * split = strtok(NULL, ";");
	char * game_id_split = strtok(trim(game_id), " ");

	game_id_split = strtok(NULL, " ");
	int parsed_id = atoi(game_id_split);
	printf("Game id parsed: %d\n\n", parsed_id);
	while(split != NULL){
		char * trimmed = trim(split);
		printf("Game found: %s\n", trimmed);
		validate_line(trimmed);
		split = strtok(NULL, ";");
	}

	exit(EXIT_SUCCESS);
}
