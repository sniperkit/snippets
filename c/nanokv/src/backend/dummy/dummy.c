#include <stdio.h>

#include "nanokv.h"

enum nanokv_ret dummy_init(void *arg)
{
	(void)arg;
	printf("%s:%d\n", __func__, __LINE__);
	return NANOKV_OK;
}

enum nanokv_ret dummy_destroy(void *arg)
{
	(void)arg;
	printf("%s:%d\n", __func__, __LINE__);
	return NANOKV_OK;
}

enum nanokv_ret dummy_flush(void *arg)
{
	(void)arg;
	printf("%s:%d\n", __func__, __LINE__);
	return NANOKV_OK;
}

enum nanokv_ret dummy_get(void *arg, enum nanokv_type type, const char *key, uint8_t *val)
{
	(void)arg;
	(void)type;
	(void)key;
	(void)val;
	printf("%s:%d\n", __func__, __LINE__);
	return NANOKV_OK;
}

enum nanokv_ret dummy_set(void *arg, enum nanokv_type vtype, const char *key, uint8_t *val, uint8_t len)
{
	(void)arg;
	(void)vtype;
	(void)key;
	(void)val;
	printf("%s:%d\n", __func__, __LINE__);
	return NANOKV_OK;
}

enum nanokv_ret dummy_del(void *arg, const char *key)
{
	(void)arg;
	(void)key;
	printf("%s:%d\n", __func__, __LINE__);
	return NANOKV_OK;
}

enum nanokv_ret dummy_incr(void *arg, const char *key)
{
	(void)arg;
	(void)key;
	printf("%s:%d\n", __func__, __LINE__);
	return NANOKV_OK;
}

enum nanokv_ret dummy_incrby(void *arg, const char *key, uint64_t increment)
{
	(void)arg;
	(void)key;
	(void)increment;
	printf("%s:%d\n", __func__, __LINE__);
	return NANOKV_OK;
}

enum nanokv_ret dummy_decr(void *arg, const char *key)
{
	(void)arg;
	(void)key;
	printf("%s:%d\n", __func__, __LINE__);
	return NANOKV_OK;
}

enum nanokv_ret dummy_decrby(void *arg, const char *key, uint64_t decrement)
{
	(void)arg;
	(void)key;
	(void)decrement;
	printf("%s:%d\n", __func__, __LINE__);
	return NANOKV_OK;
}

nanokv_backend_t nanokv_backend_dummy = {
	.name    = "dummy",
	.init    = dummy_init,
	.destroy = dummy_destroy,
	.flush   = dummy_flush,
	.get     = dummy_get,
	.set     = dummy_set,
	.del     = dummy_del,
	.incr    = dummy_incr,
	.incrby  = dummy_incrby,
	.decr    = dummy_decr,
	.decrby  = dummy_decrby
};
