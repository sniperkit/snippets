#include <stdio.h>

#include <nanokv.h>
#include <nanokv/backend/redis.h>

static nanokv_t s;

int main(void)
{
	nanokv_backend_redis_init(&s, "dev01", 6660);

	nanokv_set(&s, "foo", "bar");
	nanokv_del(&s, "foo");
	nanokv_del(&s, "foo");
	nanokv_set(&s, "coca", "cola");

	nanokv_incrby(&s, "the number of the beast", 111);

	//nanokv_flush(&s);
	nanokv_destroy(&s);

	return 0;
}
