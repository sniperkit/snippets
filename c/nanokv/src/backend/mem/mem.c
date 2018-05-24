#include <backend/mem.h>

#ifdef NANOKV_DEBUG
#include <ctype.h>
#endif

struct _nanokv {
	uint64_t magic;
	struct {
		uint32_t used;
		uint32_t free;
	} allocator;
} __attribute__((packed));

struct _nanokv_kv {
	uint32_t khash; /**< Key hash */
	uint8_t  ksize; /**< Key size */
	uint8_t  vsize; /**< Value size */
	uint8_t  ktype; /**< Key type */
	uint8_t  vtype; /**< Value type */
	uint32_t next;  /**< Next kv item offset */
} __attribute__((packed));

struct _nanokv_kv_data {
	uint8_t *kdata; /**< Key data address */
	uint8_t *vdata; /**< Value data address */
};

/* MurmurHash2 */
static uint32_t nanokv_mem_hash(const uint8_t *data, size_t len)
{
	uint32_t  h, k;

	h = 0 ^ len;

	while (len >= 4) {
		k  = data[0];
		k |= data[1] << 8;
		k |= data[2] << 16;
		k |= data[3] << 24;

		k *= 0x5bd1e995;
		k ^= k >> 24;
		k *= 0x5bd1e995;

		h *= 0x5bd1e995;
		h ^= k;

		data += 4;
		len -= 4;
	}

	switch (len) {
	case 3:
		h ^= data[2] << 16;
	case 2:
		h ^= data[1] << 8;
	case 1:
		h ^= data[0];
		h *= 0x5bd1e995;
	}

	h ^= h >> 13;
	h *= 0x5bd1e995;
	h ^= h >> 15;

	return h;
}

enum nanokv_ret nanokv_init(nanokv_t *store, uint8_t *data, size_t len)
{
	struct _nanokv *ks = (void *)data;

	store->data = data;
	store->len  = len;

	memset(data, 0, len);

	nanokv_u64((uint8_t *)&ks->magic, NANOKV_MAGIC);

	ks->allocator.free = len;
	store->free = data + sizeof(struct _nanokv);

	return NANOKV_OK;
}

void nanokv_flush(nanokv_t *store)
{
	((struct _nanokv *)store->data)->allocator.free = store->len;
}

enum nanokv_ret _nanokv_set(nanokv_t *store, enum nanokv_type vtype, const char *key, uint8_t *val, uint8_t len)
{
	struct _nanokv_kv_data kvd;
	struct _nanokv_kv *kv = store->free;

	// first we should search if the key already exists
	// when it exists and the value fits exact we reuse the kv item
	// when it wont fit in the searched kv item we mark it for gc and allocate a new free slot

	kv->ksize = strlen(key); 
	kv->khash = nanokv_hash((const uint8_t *)key, kv->ksize);

	kv->ktype = NANOKV_TYPE_STRING;
	kv->vtype = vtype;

	if (kv->ktype == NANOKV_TYPE_STRING)
		kv->vsize = strlen((const char *)val);
	else
		kv->vsize = len;

	kvd.kdata = (uint8_t *)kv + sizeof(*kv);
	kvd.vdata = kvd.kdata + kv->ksize;

	store->free = kvd.vdata + kv->vsize;

	memcpy(kvd.kdata, key, kv->ksize);
	memcpy(kvd.vdata, val, kv->vsize);

	return NANOKV_OK;
}

enum nanokv_ret nanokv_incr(nanokv_t *store, const char *key)
{
	(void)store;
	(void)key;

	return NANOKV_OK;
}

#ifdef NANOKV_DEBUG

void nanokv_dump(nanokv_t *store, unsigned int columns)
{
	unsigned int i, j;
	void *mem = store->data;
	unsigned int len = store->len;

	for(i = 0; i < len + ((len % columns) ? (columns - len % columns) : 0); i++)
        {
                /* print offset */
                if(i % columns == 0)
                {
                        printf("0x%06x: ", i);
                }
 
                /* print hex data */
                if(i < len)
                {
                        printf("%02x ", 0xFF & ((char*)mem)[i]);
                }
                else /* end of block, just aligning for ASCII dump */
                {
                        printf("   ");
                }
                
                /* print ASCII dump */
                if(i % columns == (columns - 1))
                {
                        for(j = i - (columns - 1); j <= i; j++)
                        {
                                if(j >= len) /* end of block, not really printing */
                                {
                                        putchar(' ');
                                }
                                else if(isprint(((char*)mem)[j])) /* printable char */
                                {
                                        putchar(0xFF & ((char*)mem)[j]);        
                                }
                                else /* other char */
                                {
                                        putchar('.');
                                }
                        }
                        putchar('\n');
                }
        }
}

#endif
