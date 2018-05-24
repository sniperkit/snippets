/*
 * Copyright 2016 Jerry Jacobs. All rights reserved.
 * Use of this source code is governed by the MIT
 * license that can be found in the LICENSE file.
 */
#include <stdio.h>
#include <stdlib.h>

enum ihex_offset {
	IHEX_OFFSET_SIZE  = 1,
	IHEX_OFFSET_ADDR  = 3,
	IHEX_OFFSET_RTYPE = 7,
	IHEX_OFFSET_DATA  = 9
};

enum ihex_rtype {
	IHEX_RTYPE_DATA              = 0,
	IHEX_RTYPE_EOF               = 1,
	IHEX_RTYPE_EXTENDED_SEG_ADDR = 2,
	IHEX_RTYPE_EXTENDED_LIN_ADDR = 4
};

struct ihex_rec {
	size_t size;
	uint32_t addr;
	enum ihex_rtype rtype;
	uint8_t *data;
	uint8_t checksum;
	struct ihex_rec *next;
};

void ihex_rec_dump(struct ihex_rec *rec)
{
	printf(" size: %zu\n", rec->size);
	printf(" addr: 0x%04x\n", rec->addr);
	printf("rtype: %d\n", rec->rtype);
	printf(" data: ");
	for (size_t n = 0; n < rec->size; n++)
		printf("%02x ", rec->data[n]);
	printf("\n");
	printf("checksum: %02x\n", rec->checksum);
}

uint8_t ihex_rec_checksum(const struct ihex_rec *rec)
{
	uint8_t checksum = 0;

	checksum += rec->size;
	checksum += rec->addr;
	checksum += rec->rtype;

	for (size_t n = 0; n < rec->size; n++)
		checksum += rec->data[n];

	return 1 + (~checksum);
}

/**
 * hex2int
 * take a hex string and convert it to a 32bit number (max 8 hex digits)
 */
uint32_t ihex_hex2bin(const char *hex, const size_t size) {
	uint32_t val = 0;
	const size_t _size = size * 2;
	for (size_t n = 0; n < _size; n++) {
		char c = hex[n];
		if (c == '\0')
			break;

		if (c >= '0' && c <= '9')
			c = c - '0';
		else if (c >= 'a' && c <='f')
			c = c - 'a' + 10;
		else if (c >= 'A' && c <='F')
			c = c - 'A' + 10;

		val = (val << 4) | (c & 0xf);
	}

	return val;
}

void ihex_readline(struct ihex_rec *rec, const char *line)
{
	rec->size = 0;

	if (line[0] != ':')
		return;

	rec->size  = ihex_hex2bin(line + IHEX_OFFSET_SIZE,  1);
	rec->addr  = ihex_hex2bin(line + IHEX_OFFSET_ADDR,  2);
	rec->rtype = ihex_hex2bin(line + IHEX_OFFSET_RTYPE, 1);

	rec->data = malloc(rec->size);
	if (!rec->data)
		return;

	for (size_t n = 0; n < rec->size; n++)
		rec->data[n] = ihex_hex2bin(line + IHEX_OFFSET_DATA + n, 1);

	rec->checksum = ihex_hex2bin(line + IHEX_OFFSET_DATA + (rec->size * 2), 1);
	if (rec->checksum != ihex_rec_checksum(rec)) {
		free(rec->data);
		rec->data = NULL;
		rec->size = 0;
		rec->addr = 0;
		rec->rtype = 0;
	}
}

int main(void) {
	struct ihex_rec rec;
	const char *ihex = ":10001000FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF0\r\n";

	ihex_readline(&rec, ihex);
	ihex_rec_dump(&rec);
}
