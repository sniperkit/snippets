#ifndef SCANNER_H_
#define SCANNER_H_

#include <stdlib.h>
#include <stdbool.h>
#include "reader.h"

typedef struct scanner *scanner_t;

typedef bool (*scanner_split_t)(buffer_t data, buffer_t token, bool eof, size_t *advance);

bool scanner_words(buffer_t data, buffer_t token, bool eof, size_t *advance);
bool scanner_deadbeef(buffer_t data, buffer_t token, bool eof, size_t *advance);

scanner_t scanner_new(reader_t r, scanner_split_t split);
void scanner_free(scanner_t *s);

bool scanner_scan(scanner_t s);
const char *scanner_text(scanner_t s);

#endif /* SCANNER_H_ */
