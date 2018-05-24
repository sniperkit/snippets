#ifndef BUFFER_H_
#define BUFFER_H_

#include <stdlib.h>
#include <stdbool.h>

typedef struct buffer *buffer_t;

buffer_t buffer_new(void);
void buffer_free(buffer_t *b);

const size_t buffer_len(buffer_t b);
const size_t buffer_cap(buffer_t b);
void buffer_grow(buffer_t b, size_t n);
void buffer_reset(buffer_t b);
bool buffer_discard(buffer_t b, size_t n);
bool buffer_write_byte(buffer_t b, uint8_t c);
bool buffer_read_byte(buffer_t b, uint8_t *c);
const uint8_t *buffer_bytes(buffer_t b);
const char *buffer_string(buffer_t b);
void buffer_copy(buffer_t d, buffer_t s);

#endif /* BUFFER_H_ */
