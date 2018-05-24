#include <stdio.h>

#include <nanokv.h>
#include <nanokv/backend/redis.h>
#include <hiredis/hiredis.h>

nanokv_backend_t nanokv_backend_redis;

enum nanokv_ret nanokv_backend_redis_init(nanokv_t *store, const char *ip, int port)
{
	redisContext *c = redisConnect(ip, port);
	if (!c)
		return NANOKV_ENOMEM;
	if (c->err) {
		redisFree(c);
		return NANOKV_EINVAL;
	}

	nanokv_init(store, &nanokv_backend_redis);
	nanokv_set_private(store, c);

	return NANOKV_OK;
}

enum nanokv_ret redis_init(void *arg)
{
	(void)arg;
	printf("%s:%d\n", __func__, __LINE__);
	return NANOKV_OK;
}

enum nanokv_ret redis_destroy(void *arg)
{
	printf("%s:%d\n", __func__, __LINE__);
	if (arg)
		redisFree(arg);
	return NANOKV_OK;
}

enum nanokv_ret redis_flush(void *arg)
{
	(void)arg;
	printf("%s:%d\n", __func__, __LINE__);
	return NANOKV_OK;
}

enum nanokv_ret redis_get(void *arg, enum nanokv_type type, const char *key, uint8_t *val)
{
	(void)arg;
	(void)type;
	(void)key;
	(void)val;
	printf("%s:%d\n", __func__, __LINE__);
	return NANOKV_OK;
}

enum nanokv_ret redis_set(void *arg, enum nanokv_type vtype, const char *key, uint8_t *val, uint8_t len)
{
	(void)vtype;

	redisContext *c = arg;
	printf("%s:%d\n", __func__, __LINE__);
	redisCommand(c, "SET %s %b", key, val, (size_t)len);

	return NANOKV_OK;
}

enum nanokv_ret redis_del(void *arg, const char *key)
{
	(void)arg;
	(void)key;
	printf("%s:%d\n", __func__, __LINE__);
	return NANOKV_OK;
}

enum nanokv_ret redis_incr(void *arg, const char *key)
{
	(void)arg;
	(void)key;
	printf("%s:%d\n", __func__, __LINE__);
	return NANOKV_OK;
}

enum nanokv_ret redis_incrby(void *arg, const char *key, uint64_t increment)
{
	(void)arg;
	(void)key;
	(void)increment;
	printf("%s:%d\n", __func__, __LINE__);
	return NANOKV_OK;
}

enum nanokv_ret redis_decr(void *arg, const char *key)
{
	(void)arg;
	(void)key;
	printf("%s:%d\n", __func__, __LINE__);
	return NANOKV_OK;
}

enum nanokv_ret redis_decrby(void *arg, const char *key, uint64_t decrement)
{
	(void)arg;
	(void)key;
	(void)decrement;
	printf("%s:%d\n", __func__, __LINE__);
	return NANOKV_OK;
}

nanokv_backend_t nanokv_backend_redis = {
	.name    = "dummy",
	.init    = redis_init,
	.destroy = redis_destroy,
	.flush   = redis_flush,
	.get     = redis_get,
	.set     = redis_set,
	.del     = redis_del,
	.incr    = redis_incr,
	.incrby  = redis_incrby,
	.decr    = redis_decr,
	.decrby  = redis_decrby
};
