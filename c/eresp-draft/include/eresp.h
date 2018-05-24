#ifndef ERESP_H__
#define ERESP_H__

#include <stdlib.h>
#include <stdint.h>

#define ERESP_ARRAY_SIZE(x) (sizeof(x)/sizeof(x[0]))

typedef struct eresp_reader_t eresp_reader_t;
typedef struct eresp_item_t eresp_item_t;

enum eresp_t {
	ERESP_T_UNKNOWN      = '\0',
	ERESP_T_SIMPLESTRING = '+',
	ERESP_T_ERROR        = '-',
	ERESP_T_INTEGER      = ':',
	ERESP_T_BULKSTRING   = '$',
	ERESP_T_ARRAY        = '*'
};

struct eresp_item_t {
	enum eresp_t type;
	char *data;
	size_t len;
};

struct eresp_reader_t {
	void *context;
	int state;
	char *buf;
	char *bufc;
	char *bufp;
	size_t buflen;
	struct eresp_item_t items[8];
	size_t itemslen;
};

void eresp_reader_init(eresp_reader_t *r, char *buf, size_t len);
void eresp_read_data(eresp_reader_t *r, const char *data, const size_t len);
void eresp_read_char(eresp_reader_t *r, const char c);
void eresp_reader_set_context(eresp_reader_t *r, void *context);

typedef void eresp_item_proc_t (const eresp_reader_t *r, eresp_item_t *i);

typedef void eresp_cmd_proc_t (int argc, const eresp_item_t **argv);
struct eresp_cmd_t {
	char *name;
	eresp_cmd_proc_t *proc;
	int args;
};

#endif /* ERESP_H__ */
