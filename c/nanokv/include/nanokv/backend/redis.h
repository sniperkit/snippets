#ifndef NANOKV_BACKEND_REDIS_H_
#define NANOKV_BACKEND_REDIS_H_

#include <nanokv.h>

#ifdef __cplusplus
extern "C"
{
#endif

enum nanokv_ret nanokv_backend_redis_init(nanokv_t *store, const char *ip, int port);

#ifdef __cplusplus
}
#endif

#endif /* NANOKV_BACKEND_REDIS_H_ */
