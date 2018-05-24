#ifndef ERESP_H__
#define ERESP_H__

#include <stdlib.h>
#include <stdint.h>
#include <stdbool.h>

#ifdef __cplusplus
extern "C" {
#endif

#define ERESP_ARRAY_SIZE(x) (sizeof(x)/sizeof(x[0]))

typedef struct eresp_ctx_t eresp_ctx_t;
typedef struct eresp_item_t eresp_item_t;
typedef struct eresp_cmd_t eresp_cmd_t;

#define ERESP_CMD(cmd, func, argc) \
	{ cmd, (sizeof(cmd) - 1), func, argc }
#define ERESP_CMD_LAST \
	{ NULL, 0, NULL, 0 }

enum eresp_item_type {
	ERESP_T_UNKNOWN      = 0,
	ERESP_T_SIMPLESTRING = '+',
	ERESP_T_ERROR        = '-',
	ERESP_T_INTEGER      = ':',
	ERESP_T_BULKSTRING   = '$',
	ERESP_T_ARRAY        = '*',
	ERESP_T_INLINE       = 'I'
};

enum eresp_err {
	ERESP_OK,
	ERESP_EPARAM,
	ERESP_ENOMEM,
	ERESP_EPROTO,
	ERESP_ETYPE,
};

struct eresp_item_t {
	enum eresp_item_type type;
	const char *data;
	size_t len;
};

/** Command processor */
typedef void eresp_cmd_proc_t(const eresp_ctx_t *ctx, int argc, const eresp_item_t *argv);

struct eresp_cmd_t {
	const char *name;
	size_t namelen;
	eresp_cmd_proc_t *proc;
	int argc;
};

/**
 * Item processor
 * @retval true when the application is finished processing the all the buffered items
 */
typedef bool eresp_item_proc_t(const eresp_ctx_t *ctx, const eresp_item_t *i);

/**
 * Output writer
 */
typedef void eresp_writer_t(const eresp_ctx_t *ctx, const char c);

struct eresp_ctx_t {
	void *context;
	enum eresp_err err;
	int state;
	eresp_item_proc_t *proc;
	const eresp_cmd_t *commands;
	eresp_writer_t *writer;
	char *buf;
	char *bufc;
	size_t buflen;
	struct eresp_item_t items[8];
	size_t itemslen;
};

void eresp_init(eresp_ctx_t *e, char *buf, size_t len);
void eresp_reset(eresp_ctx_t *e);
void eresp_read_data(eresp_ctx_t *e, const char *data, const size_t len);
void eresp_read_char(eresp_ctx_t *e, const char c);

void eresp_set_context(eresp_ctx_t *ctx, void *context);
void eresp_set_writer(eresp_ctx_t *ctx, eresp_writer_t *writer);
void eresp_set_item_proc(eresp_ctx_t *ctx, eresp_item_proc_t *proc);
void eresp_set_commands(eresp_ctx_t *ctx, const eresp_cmd_t *commands);

void eresp_write_error(const eresp_ctx_t *ctx, const char *err);
void eresp_write_bulkstring(const eresp_ctx_t *ctx, const char *str);
void eresp_write_string(const eresp_ctx_t *ctx, const char *str);
void eresp_write_data(const eresp_ctx_t *ctx, const char *data, const size_t len);
void eresp_write_item(const eresp_ctx_t *ctx, const eresp_item_t *item);
void eresp_write_items(const eresp_ctx_t *ctx, const eresp_item_t *items, const size_t len);
void eresp_write_array(const eresp_ctx_t *ctx, const size_t len);

void eresp_item_dump(const eresp_item_t *i);

enum eresp_err eresp_item_strconv_bool(bool *v, const eresp_item_t *i);
enum eresp_err eresp_item_strconv_u8(uint8_t *v, const eresp_item_t *i);
enum eresp_err eresp_item_strconv_u16(uint16_t *v, const eresp_item_t *i);
enum eresp_err eresp_item_strconv_u32(uint32_t *v, const eresp_item_t *i);
enum eresp_err eresp_item_strconv_u64(uint64_t *v, const eresp_item_t *i);


const struct eresp_cmd_t *eresp_cmd_search(const struct eresp_cmd_t *list, const struct eresp_item_t *i); 
bool eresp_cmd_item_proc(const eresp_ctx_t *ctx, const eresp_item_t *item);


#ifdef __cplusplus
}
#endif

#endif /* ERESP_H__ */
