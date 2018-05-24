#include <stdio.h>
#include <stdlib.h>
#include <stddef.h>
#include <string.h>

#include "nanokv.h"

enum nanokv_ret nanokv_init(nanokv_t *store, nanokv_backend_t *backend)
{
	enum nanokv_ret ret = NANOKV_OK;

	memset(store, 0, sizeof(*store));
	store->backend = backend;

	if (backend)
		ret = backend->init(store->private_data);

	return ret;
}

void nanokv_set_private(nanokv_t *store, void *private_data)
{
	store->private_data = private_data;
}

void *nanokv_get_private(nanokv_t *store)
{
	return store->private_data;
}

enum nanokv_ret nanokv_destroy(nanokv_t *store)
{
	enum nanokv_ret ret = NANOKV_EOPNOTSUPP;

	if (store->backend && store->backend->destroy)
		ret = store->backend->destroy(store->private_data);

	memset(store, 0, sizeof(*store));

	return ret;
}

enum nanokv_ret nanokv_flush(nanokv_t *store)
{
	enum nanokv_ret ret = NANOKV_EOPNOTSUPP;

	if (store->backend && store->backend->flush)
		ret = store->backend->flush(store->private_data);

	return ret;
}

enum nanokv_ret _nanokv_set(nanokv_t *store, enum nanokv_type vtype, const char *key, uint8_t *val, uint8_t len)
{
	enum nanokv_ret ret = NANOKV_EOPNOTSUPP;

	if (store->backend && store->backend->set){
		if (vtype == NANOKV_TYPE_STRING)
			len = strlen((const char *)val);

		ret = store->backend->set(store->private_data, vtype, key, val, len);
	}
	return ret;
}

enum nanokv_ret nanokv_incr(nanokv_t *store, const char *key)
{
	enum nanokv_ret ret = NANOKV_EOPNOTSUPP;

	if (store->backend && store->backend->incr)
		ret = store->backend->incr(store->private_data, key);

	return ret;
}

enum nanokv_ret _nanokv_get(nanokv_t *store, enum nanokv_type type, const char *key, uint8_t *val)
{
	enum nanokv_ret ret = NANOKV_EOPNOTSUPP;

	if (store->backend && store->backend->get)
		ret = store->backend->get(store->private_data, type, key, val);
	return ret;
}

enum nanokv_ret nanokv_del(nanokv_t *store, const char *key)
{
	enum nanokv_ret ret = NANOKV_EOPNOTSUPP;

	if (store->backend && store->backend->del)
		ret = store->backend->del(store->private_data, key);
	return ret;
}

enum nanokv_ret nanokv_incrby(nanokv_t *store, const char *key, uint64_t increment)
{
	enum nanokv_ret ret =  NANOKV_EOPNOTSUPP;

	if (store->backend && store->backend->incrby)
		ret = store->backend->incrby(store->private_data, key, increment);
	return ret;
}

enum nanokv_ret nanokv_decr(nanokv_t *store, const char *key)
{
	enum nanokv_ret ret = NANOKV_EOPNOTSUPP;

	if (store->backend && store->backend->decr)
		ret = store->backend->decr(store->private_data, key);
	return ret;
}

enum nanokv_ret nanokv_decrby(nanokv_t *store, const char *key, uint64_t decrement)
{
	enum nanokv_ret ret = NANOKV_EOPNOTSUPP;

	if (store->backend && store->backend->decrby)
		ret = store->backend->decrby(store->private_data, key, decrement);
	return ret;
}
