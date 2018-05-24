#include "reader.h"
#include "buffer.h"

#include <stdio.h>
#include <stdlib.h>

struct reader {
	FILE *fp;
	buffer_t buf;
};

reader_t reader_new(FILE *fp)
{
	struct reader *r;

	r = malloc(sizeof(*r));
	if (!r)
		return NULL;

	r->buf = buffer_new();
	if (!r->buf) {
		free(r);
		return NULL;
	}

	r->fp = fp;

	return r;
}

void reader_free(reader_t *r)
{
	struct reader *_r = *r;

	buffer_free(&_r->buf);
	free(_r);
	*r = NULL;
}

bool reader_read(reader_t r, buffer_t buf)
{
	uint8_t b;

	if (feof(r->fp))
		return false;

	size_t read = fread(&b, 1, 1, r->fp);
	if (read == 0)
		return false;

	buffer_write_byte(buf, b);

	return true;
}
