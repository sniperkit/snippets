#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "reader.h"
#include "scanner.h"

int main(int argc, const char **argv)
{
	if (argc != 2)
		return EXIT_FAILURE;

	reader_t r = reader_new(stdin);
	scanner_t s;

	if (strcmp("words", argv[1]) == 0) {
		s = scanner_new(r, scanner_words);
	} else if (strcmp("deadbeef", argv[1]) == 0) {
		s = scanner_new(r, scanner_deadbeef);
	} else {
		reader_free(&r);
		return EXIT_FAILURE;
	}

	while (scanner_scan(s))
		printf("\"%s\"\n", scanner_text(s));

	scanner_free(&s);
	reader_free(&r);

	return EXIT_SUCCESS;
}
