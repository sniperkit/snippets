#include <eresp.h>
#include <stdio.h>

/** Reader states */
enum eresp_reader_states {
	ERESP_S_START = 'S', /** Start */
	ERESP_S_LEN   = 'L', /** Length processing */
	ERESP_S_LF    = 'F', /** Waiting for linefeed '\r' */
	ERESP_S_BUF   = 'B'  /** Buffering item content */
};

static void eresp_item_dump(const eresp_item_t *i)
{
	printf("(item at %p) item->len: %zu, item->type: %c\n", i, i->len, i->type);
	if (i->data == NULL || i->type == ERESP_T_ARRAY)
		return;
	printf("\t");
	for (size_t n = 0; n < i->len; n++)
		printf("%c", i->data[n]);
	printf("\n");
}

static void eresp_reader_bufc(eresp_reader_t *r, const char c) {
	if (r->bufc >= (r->buf + r->buflen)) {
		// TODO err
		return;
	}
	*r->bufc = c;
	r->bufc++;
}

static void eresp_reader_buf_reset(eresp_reader_t *r) {
	r->bufc = r->buf;
	r->bufp = r->buf;
}

static eresp_item_t *eresp_reader_item_get(eresp_reader_t *r) {
	return &r->items[r->itemslen];
}

static eresp_item_t *eresp_reader_item_get_next(eresp_reader_t *r) {
	r->itemslen++;
	if (r->itemslen == (sizeof(r->items)/sizeof(r->items[0])))
		r->itemslen--; // FIXME
	return eresp_reader_item_get(r);
}

static void eresp_reader_array_item_buf_finished(eresp_reader_t *r)
{
	// When first buffered item is a array we make sure to decrement every added item
	eresp_item_t *fitem = &r->items[0];

	if (fitem->type != ERESP_T_ARRAY)
		return;
	if (fitem->len == 0)
		return;

	fitem->len--;
	if (fitem->len != 0)
		return;

	fitem->len  = (size_t)fitem->data;
	fitem->data = NULL;

	eresp_item_dump(fitem);
	printf("\tfinished processing array with len: %zu\n", fitem->len);

	r->itemslen = 0;
	eresp_reader_buf_reset(r);
}

/** Reader buffering state processing */
static void eresp_reader_s_buf(eresp_reader_t *r, const char c)
{
	eresp_item_t *item = eresp_reader_item_get(r);
	eresp_reader_bufc(r, c);

	item->len--;
	if (item->len > 0)
		return;

	// Current buffered item is finished
	item->data = r->bufp;
	item->len  = (size_t)(r->bufc - r->bufp);
	eresp_item_dump(item);

	// Create next item data cut point
	r->bufp = r->bufc;

	eresp_reader_array_item_buf_finished(r);
	eresp_reader_item_get_next(r);
	r->state = ERESP_S_START;
}

/** Reader linefeed state processor */
static void eresp_reader_s_lf(eresp_reader_t *r)
{
	eresp_item_t *item = eresp_reader_item_get(r);

	if (item->type == ERESP_T_ARRAY) {
		eresp_item_dump(item);
		eresp_reader_item_get_next(r);
		r->state = ERESP_S_START;
	} else {
		r->state = ERESP_S_BUF;
	}
}

/** Reader len state processor */
static void eresp_reader_s_len(eresp_reader_t *r, const char c)
{
	eresp_item_t *item = eresp_reader_item_get(r);

	if (c == '\r') {
		if (item->type == ERESP_T_ARRAY)
			item->data = (void *)item->len;
		r->state = ERESP_S_LF;
		return;
	}

	if (c < '0' || c > '9') {
		r->state = ERESP_S_START;
		return;
	}

	item->len = (item->len * 10) + (size_t)(c - '0');
}

/** Reader start state */
static void eresp_reader_s_start(eresp_reader_t *r, const char c)
{
	eresp_item_t *item = eresp_reader_item_get(r);

	switch (c) {
	case ERESP_T_SIMPLESTRING:
	case ERESP_T_ERROR:
		r->state = ERESP_S_BUF;
		break;
	case ERESP_T_ARRAY:
	case ERESP_T_BULKSTRING:
		r->state = ERESP_S_LEN;
		break;
	case ERESP_T_INTEGER:
		break;
	default:
		return;
	}

	item->type = c;

}

void eresp_reader_init(eresp_reader_t *r, char *buf, size_t len)
{
	r->state  = ERESP_S_START;
	// buffer
	r->buf    = buf;
	r->bufc   = buf;
	r->bufp   = buf;
	r->buflen = len;
	// items
	r->itemslen = 0;
}

void eresp_reader_set_context(eresp_reader_t *r, void *context)
{
	r->context = context;
}

void eresp_read_char(eresp_reader_t *r, const char c)
{
	switch (r->state) {
	case ERESP_S_START:
		eresp_reader_s_start(r, c);
		break;
	case ERESP_S_LEN:
		eresp_reader_s_len(r, c);
		break;
	case ERESP_S_LF:
		eresp_reader_s_lf(r);
		break;
	case ERESP_S_BUF:
		eresp_reader_s_buf(r, c);
		break;
	default:
		break;
	}
}

void eresp_read_data(eresp_reader_t *r, const char *data, const size_t len)
{
	for (size_t n = 0; n < len; n++)
		eresp_read_char(r, data[n]);
}
