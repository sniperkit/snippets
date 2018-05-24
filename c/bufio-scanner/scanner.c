#include "scanner.h"

struct scanner {
	reader_t r;
	buffer_t buf;
	buffer_t token;
	scanner_split_t split;
};

scanner_t scanner_new(reader_t r, scanner_split_t split)
{
	struct scanner *s;

	s = malloc(sizeof(*s));
	if (!s)
		return NULL;

	s->r = r;
	s->buf = buffer_new();
	if (!s->buf) {
		free(s);
		return NULL;
	}

	s->split = split;
	s->token = buffer_new();

	return s;
}

void scanner_free(scanner_t *s)
{
	struct scanner *_s = *s;

	free(*s);
	*s = NULL;
}

bool scanner_scan(scanner_t s)
{
	return reader_read(s->r, s->buf);
}

bool scanner_words(buffer_t data, buffer_t token, bool eof, size_t *advance)
{
	uint8_t b;

	for (size_t n = 0; n < buffer_len(data); n++) {
		if (buffer_read_byte(data, &b)) {
			if (b == ' ')
				return true;
			buffer_write_byte(token, b);
		}
	}

	return false;
}

bool scanner_deadbeef(buffer_t data, buffer_t token, bool eof, size_t *advance)
{
	if (eof) {
		buffer_copy(token, data);
		return false;
	}

	if (buffer_len(data) < 4)
		return false;

//	const uint8_t *d = buffer_bytes(data);
//	printf("d[0..3] = %02x %02x %02x %02x\n", d[0], d[1], d[2], d[3]);

	if (memcmp("\xde\xad\xbe\xef", buffer_bytes(data), 4) == 0) {
		buffer_discard(data, 4);
		if (buffer_len(token))
			return true;
	}

	buffer_copy(token, data);

	return false;
}

const char *scanner_text(scanner_t s)
{
	bool eof = false;
	size_t advance = 0;

	buffer_reset(s->token);

	do {

		if (s->split(s->buf, s->token, eof, &advance) &&
		    buffer_len(s->token))
			break;
		if (!scanner_scan(s))
			eof = true;
	} while (!eof);

	return buffer_string(s->token);
}
