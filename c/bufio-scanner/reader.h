#ifndef READER_H_
#define READER_H_

#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

#include "buffer.h"

typedef struct reader *reader_t;

reader_t reader_new(FILE *fp);
void reader_free(reader_t *r);

bool reader_read(reader_t r, buffer_t buf);

#endif /* READER_H_ */
