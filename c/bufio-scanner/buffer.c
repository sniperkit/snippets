#include "buffer.h"

#define BUFFER_SIZE 4096

struct buffer {
	size_t cap;
	size_t len;
	size_t rpos;
	uint8_t *buf;
};

buffer_t buffer_new(void)
{
	struct buffer *buf;

	buf = malloc(sizeof(*buf));
	if (!buf)
		return NULL;

	buf->buf = malloc(BUFFER_SIZE);
	if (!buf->buf) {
		free(buf);
		return NULL;
	}

	buf->cap = BUFFER_SIZE;
	buf->len = 0;
	buf->rpos = 0;

	return buf;
}

void buffer_free(struct buffer **b)
{
	struct buffer *buf = *b;

	free(buf->buf);
	free(*b);
	*b = NULL;
}

const size_t buffer_len(buffer_t b)
{
	return b->len;
}

const size_t buffer_cap(buffer_t b)
{
	return b->cap;
}

void buffer_grow(buffer_t b, size_t n)
{
	if (!n)
		return;

	size_t cap = b->cap + n;

	if (cap < n)
		cap = SIZE_MAX;

	b->buf = realloc(b->buf, cap);
	if (b->buf)
		b->cap = cap;
	else
		b->cap = 0;
}

void buffer_reset(buffer_t b)
{
	b->len = 0;
}

bool buffer_write_byte(buffer_t b, uint8_t c)
{
	if (b->cap == b->len)
		buffer_grow(b, BUFFER_SIZE);

	b->buf[b->len] = c;
	b->len++;

	return true;
}

bool buffer_read_byte(buffer_t b, uint8_t *c)
{
	if (!b->len)
		return false;

	*c = b->buf[b->rpos];
	b->len--;

	if (b->len) {
		b->rpos++;
	} else {
		b->rpos = 0;
	}

	return true;
}

bool buffer_discard(buffer_t b, size_t n)
{
	if (b->len < n)
		return false;

	b->len -= n;

	return true;
}

const uint8_t *buffer_bytes(buffer_t b)
{
	return b->buf;
}

const char *buffer_string(buffer_t b)
{
	if (b->buf[b->len] != '\0')
		b->buf[b->len] = '\0';

	return (const char *)&b->buf[b->rpos];
}

void buffer_copy(buffer_t d, buffer_t s)
{
	uint8_t b;
	const size_t len = buffer_len(s);

	if (!len)
		return;

	for (size_t n = 0; n < len; n++) {
		if (!buffer_read_byte(s, &b))
			break;
		buffer_write_byte(d, b);
	}
}
