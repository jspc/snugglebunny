snugglebunny: snugglebunny.c
	$(CC) -o snugglebunny syntax.c keys.c term.c buffer.c snugglebunny.c -Wall -Wextra -pedantic -std=c99
