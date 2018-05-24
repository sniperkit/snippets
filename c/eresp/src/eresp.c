#include <eresp.h>
#include <string.h>

/** Reader FSM states */
enum eresp_states {
	ERESP_S_START  = 'S', /** Start */
	ERESP_S_INLINE = 'I', /** Inline commands */
	ERESP_S_LEN    = 'L', /** Length processing */
	ERESP_S_LF     = 'F', /** Waiting for linefeed '\r' */
	ERESP_S_BUF    = 'B'  /** Buffering item content */
};

/**
 * Buffer a single character
 */
static void eresp_ctx_bufc(eresp_ctx_t *ctx, const char c) {
	if (ctx->bufc >= (ctx->buf + ctx->buflen)) {
		ctx->err = ERESP_ENOMEM;
		return;
	}
	*ctx->bufc = c;
	ctx->bufc++;
}

/**
 * Get current item
 */
static eresp_item_t *eresp_item_get(eresp_ctx_t *ctx)
{
	return &ctx->items[ctx->itemslen];
}

/**
 * Get next item
 */
static eresp_item_t *eresp_item_next(eresp_ctx_t *ctx)
{
	struct eresp_item_t *item = eresp_item_get(ctx);

	ctx->itemslen++;
	if (ctx->itemslen == ERESP_ARRAY_SIZE(ctx->items)) {
		ctx->err = ERESP_EPROTO;
		ctx->itemslen = 0;
		return NULL;
	}
	return item;
}

/**
 * Run item processor
 */
static void eresp_item_proc(eresp_ctx_t *ctx, eresp_item_t *i)
{
	if (!ctx->proc)
		return;

	const bool reset = ctx->proc(ctx, i);

	if (reset)
		eresp_reset(ctx);
}

/**
 * Get previous item
 */
static const eresp_item_t *eresp_item_get_prev(const eresp_ctx_t *ctx)
{
	if (ctx->itemslen > 0)
		return &ctx->items[ctx->itemslen - 1];
	return NULL;
}

/**
 * Run item finisher if buffered
 */
static void eresp_item_finish(eresp_ctx_t *ctx, eresp_item_t *item)
{
	// Current buffered item is finished
	// FIXME limited to prev item with data, else the data pointer is incorrect calculated
	const eresp_item_t *itemp = eresp_item_get_prev(ctx);

	if (itemp && itemp->data) {
		item->len = (size_t)(ctx->bufc - (itemp->data + itemp->len));
	} else {
		item->len = (size_t)(ctx->bufc - ctx->buf);
	}
	item->data = ctx->bufc - item->len;

	eresp_item_proc(ctx, item);
}

/** Reader buffering state processing */
static void eresp_s_buf(eresp_ctx_t *ctx, const char c)
{
	eresp_item_t *item = eresp_item_get(ctx);

	switch (item->type) {
	case ERESP_T_SIMPLESTRING:
	case ERESP_T_ERROR:
		if (c != '\r') {
			eresp_ctx_bufc(ctx, c);
			return;
		}
		eresp_ctx_bufc(ctx, '\0');
		break;
	case ERESP_T_BULKSTRING:
		eresp_ctx_bufc(ctx, c);
		item->len--;
		if (item->len > 0)
			return;
		break;
	default:
		return;
	}

	eresp_item_finish(ctx, item);

	// check if array is finished
	// FIXME itterate backwards from current item to see if it was part of an array
	//       instead of assuming the first item is an array
	// XXX document the fire of item processor when the array is finished
	if (ctx->items[0].type == ERESP_T_ARRAY) {
		if (ctx->items[0].len == ctx->itemslen)
			eresp_item_proc(ctx, &ctx->items[0]);
	}

	// Prepare next item if eresp_item_proc did not reset the context
	if (ctx->state == ERESP_S_START)
		return;

	eresp_item_next(ctx);
	ctx->state = ERESP_S_START;
}

/** Reader linefeed state processor */
static void eresp_s_lf(eresp_ctx_t *ctx, const char c)
{
	if (c != '\n') {
		ctx->state = ERESP_S_START;
		return;
	}

	eresp_item_t *item = eresp_item_get(ctx);

	if (item->type == ERESP_T_ARRAY) {
		eresp_item_proc(ctx, item);
		eresp_item_next(ctx);
		ctx->state = ERESP_S_START;
	} else {
		ctx->state = ERESP_S_BUF;
	}
}

/** Reader len state processor */
static void eresp_s_len(eresp_ctx_t *ctx, const char c)
{
	eresp_item_t *item = eresp_item_get(ctx);

	if (c == '\r') {
		ctx->state = ERESP_S_LF;
		return;
	}

	// TODO add sign

	if (c < '0' || c > '9') {
		ctx->state = ERESP_S_START;
		return;
	}

	item->len = (item->len * 10) + (size_t)(c - '0');
}

/** Reader inline command state */
static void eresp_s_inline(eresp_ctx_t *ctx, const char c)
{
	eresp_item_t *item = NULL;

	switch (c) {
	case '\r':
	case '\n':
		ctx->state = ERESP_S_START;
	case ' ':
		item = eresp_item_get(ctx);
		eresp_item_finish(ctx, item);
		item = eresp_item_next(ctx);
		item->type = ERESP_T_INLINE;
		break;
	default:
		eresp_ctx_bufc(ctx, c);
		return;
	}

	// When finished state, we need to trigger the application callback
	if (ctx->state == ERESP_S_START)
		eresp_item_proc(ctx, &ctx->items[0]);
}

/** Reader start state */
static void eresp_s_start(eresp_ctx_t *ctx, const char c)
{
	eresp_item_t *item = eresp_item_get(ctx);

	switch (c) {
	case ERESP_T_SIMPLESTRING:
		break;
	case ERESP_T_ARRAY:
	case ERESP_T_BULKSTRING:
		ctx->state = ERESP_S_LEN;
		break;
	case ERESP_T_INTEGER:
		// TODO
		break;
	case '\r':
	case '\n':
		// Consume trailing \r and/or \n characters
		break;
	default:
		ctx->state = ERESP_S_INLINE;
		item->type = ERESP_T_INLINE;
		eresp_ctx_bufc(ctx, c);
		return;
	}

	item->type = (enum eresp_item_type)c;
}

void eresp_init(eresp_ctx_t *ctx, char *buf, size_t len)
{
	ctx->proc    = NULL;
	ctx->context = NULL;
	ctx->buf     = buf;
	ctx->buflen  = len;
	ctx->writer  = NULL;
	eresp_reset(ctx);
}

void eresp_set_context(eresp_ctx_t *ctx, void *context)
{
	ctx->context = context;
}

void eresp_set_writer(eresp_ctx_t *ctx, eresp_writer_t *writer)
{
	ctx->writer = writer;
}

void eresp_set_item_proc(eresp_ctx_t *ctx, eresp_item_proc_t *proc)
{
	ctx->proc = proc;
}

void eresp_set_commands(eresp_ctx_t *ctx, const eresp_cmd_t *commands)
{
	ctx->commands = commands;
	eresp_set_item_proc(ctx, eresp_cmd_item_proc);
}


void eresp_item_dump(const eresp_item_t *i)
{
#ifndef ERESP_HAVE_STDIO
	(void)i;
#else
	printf("(item at %p) item->len: %zu, item->type: %c, item->data: %p\n", i, i->len, i->type, i->data);
	if (i->data == NULL || i->type == ERESP_T_ARRAY)
		return;
	printf("\t");
	if (i->type == ERESP_T_SIMPLESTRING) {
		printf("%s", i->data);
	} else {
		for (size_t n = 0; n < i->len; n++)
			printf("%c", i->data[n]);
	}
	printf("\n");
#endif
}

void eresp_reset(eresp_ctx_t *ctx)
{
	ctx->state = ERESP_S_START;
	ctx->err   = ERESP_OK;
	ctx->bufc  = ctx->buf;

	ctx->itemslen = 0;

	for (size_t n = 0; n < ERESP_ARRAY_SIZE(ctx->items); n++) {
		ctx->items[n].type = ERESP_T_UNKNOWN;
		ctx->items[n].data = NULL;
		ctx->items[n].len  = 0;
	}
}

void eresp_read_char(eresp_ctx_t *ctx, const char c)
{
	switch (ctx->state) {
	case ERESP_S_START:
		eresp_s_start(ctx, c);
		break;
	case ERESP_S_INLINE:
		eresp_s_inline(ctx, c);
		break;
	case ERESP_S_LEN:
		eresp_s_len(ctx, c);
		break;
	case ERESP_S_LF:
		eresp_s_lf(ctx, c);
		break;
	case ERESP_S_BUF:
		eresp_s_buf(ctx, c);
		break;
	default:
		break;
	}
}

void eresp_read_data(eresp_ctx_t *ctx, const char *data, const size_t len)
{
	for (size_t n = 0; n < len; n++)
		eresp_read_char(ctx, data[n]);
}

const struct eresp_cmd_t *eresp_cmd_search(const struct eresp_cmd_t *list, const struct eresp_item_t *i)
{
	if (!i->data)
		return NULL;
	if (i->len == 0)
		return NULL;

	for (const struct eresp_cmd_t *cmd = list; cmd->name != NULL; cmd++) {
		if (i->len != cmd->namelen)
			continue;
		if (memcmp(i->data, cmd->name, cmd->namelen) != 0)
			continue;
		return cmd;
	}
	return NULL;
}

bool eresp_cmd_item_proc(const eresp_ctx_t *ctx, const eresp_item_t *item)
{
	if (ctx->itemslen == 0)
		return false;

	int argc;
	int cmd_offset;

	switch (item->type) {
	case ERESP_T_ARRAY:
		// When type is a RESP array the command is at offset 1
		//  the first item is the array itself
		cmd_offset = 1;
		argc = (int)ctx->itemslen - 1;
		break;
	case ERESP_T_INLINE:
		cmd_offset = 0;
		argc = (int)ctx->itemslen - 1;
		break;
	default:
		return false;
	}

	const struct eresp_cmd_t *cmd = eresp_cmd_search(ctx->commands, &ctx->items[cmd_offset]);

	if (!cmd)
		return true;

	int cmdargc;

	if (argc > 0)
		cmdargc = argc;
	else
		cmdargc = 0;

	// FIXME item 3 can be empty
	if (cmdargc > 0)
		cmd->proc(ctx, cmdargc, &ctx->items[cmd_offset + 1]);
	else
		cmd->proc(ctx, 0, NULL);

	return true;
}

static void eresp_write_item_str(const eresp_ctx_t *ctx, const enum eresp_item_type type, const char *str)
{
	eresp_item_t item;

	item.type = type;
	item.data = str;
	item.len  = strlen(str);

	eresp_write_item(ctx, &item);
}

void eresp_write_error(const eresp_ctx_t *ctx, const char *err)
{
	eresp_write_item_str(ctx, ERESP_T_ERROR, err);
}

void eresp_write_bulkstring(const eresp_ctx_t *ctx, const char *str)
{
	eresp_write_item_str(ctx, ERESP_T_BULKSTRING, str);
}

void eresp_write_string(const eresp_ctx_t *ctx, const char *str)
{
	eresp_write_item_str(ctx, ERESP_T_SIMPLESTRING, str);
}

void eresp_write_data(const eresp_ctx_t *ctx, const char *data, const size_t len)
{
	eresp_item_t item;

	item.type = ERESP_T_BULKSTRING;
	item.data = data;
	item.len  = len;

	eresp_write_item(ctx, &item);
}

void eresp_write_item(const eresp_ctx_t *ctx, const eresp_item_t *item)
{
	(void)ctx;
	(void)item;

	if (!ctx->writer)
		return;

	ctx->writer(ctx, item->type);

	const size_t len = item->len;

	switch (item->type) {
	case ERESP_T_ARRAY:
	case ERESP_T_BULKSTRING:
	{
		size_t divisor = 1;

		while (len > 9)
			divisor *= 10;

		do {
			ctx->writer(ctx, (const char)('0' + (len / divisor % 10)));
			divisor /= 10;
		} while (divisor > 0);

		ctx->writer(ctx, '\r');
		ctx->writer(ctx, '\n');
		break;
	}
	default:
		break;
	}

	if (!item->data)
		return;

	for (size_t n = 0; n < len; n++)
		ctx->writer(ctx, item->data[n]);

	ctx->writer(ctx, '\r');
	ctx->writer(ctx, '\n');
}

void eresp_write_items(const eresp_ctx_t *ctx, const eresp_item_t *items, const size_t len)
{
	for (size_t n = 0; n < len; n++)
		eresp_write_item(ctx, &items[n]);
}

void eresp_write_array(const eresp_ctx_t *ctx, const size_t len)
{
	eresp_item_t item;

	item.type = ERESP_T_ARRAY;
	item.data = NULL;
	item.len  = len;

	eresp_write_item(ctx, &item);
}

static int8_t eresp_item_strconv_digit(const char digit)
{
	if (digit <= '9')
		return (int8_t)(digit - '0');
	if (digit <= 'Z')
		return (int8_t)(digit - 'A' + 10);
	if (digit <= 'z')
		return (int8_t)(digit - 'a' + 10);
	return INT8_MIN;
}

// TODO check overflows for the ctype instead of accumulating always in uint64_t
#define ERESP_ITEM_STRCONV_UINT_FUNC(name, ctype, min, max) \
	enum eresp_err eresp_item_strconv_##name(ctype *v, const eresp_item_t *i) \
	{ \
		if (!v) \
			return ERESP_EPARAM; \
		if (!i) \
			return ERESP_EPARAM; \
		if (!i->data) \
			return ERESP_EPARAM; \
\
		uint64_t accum = 0; \
		const char *data = i->data; \
		const size_t len = i->len; \
\
		for (size_t n = 0; n < len; n++) { \
			const int8_t digit = eresp_item_strconv_digit(data[n]); \
			if (digit == INT8_MIN) \
				return ERESP_ETYPE; \
			accum = (accum * 10) + (uint64_t)digit; \
		} \
\
		if (accum <= min || \
			accum > max) \
			return ERESP_ETYPE; \
\
		*v = (ctype)accum;\
		return ERESP_OK;\
	}

ERESP_ITEM_STRCONV_UINT_FUNC(u8,  uint8_t,  0, UINT8_MAX)
ERESP_ITEM_STRCONV_UINT_FUNC(u16, uint16_t, 0, UINT16_MAX)
ERESP_ITEM_STRCONV_UINT_FUNC(u32, uint32_t, 0, UINT32_MAX)
ERESP_ITEM_STRCONV_UINT_FUNC(u64, uint64_t, 0, UINT64_MAX)

#undef ERESP_ITEM_STRCONV_FUNC

enum eresp_err eresp_item_strconv_bool(bool *v, const eresp_item_t *i)
{
	const char *data = i->data;
	const size_t len = i->len;

	if (len != 1)
		return ERESP_ETYPE;

	switch (*data) {
	case '0':
		*v = false;
		return ERESP_OK;
	case '1':
		*v = true;
		return ERESP_OK;
	default:
		break;
	}

	return ERESP_ETYPE;
}
