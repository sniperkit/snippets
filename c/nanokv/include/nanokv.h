#ifndef NANOKV_H_
#define NANOKV_H_

#include <stdlib.h>
#include <stdint.h>

#define NANOKV_MAGIC 0x5f4e414e4f4b565f

/** Types */
enum nanokv_type {
	NANOKV_TYPE_UINT8  = 0x01,
	NANOKV_TYPE_UINT16 = 0x02,
	NANOKV_TYPE_UINT32 = 0x03,
	NANOKV_TYPE_UINT64 = 0x04,
	NANOKV_TYPE_INT8   = 0x05,
	NANOKV_TYPE_INT16  = 0x06,
	NANOKV_TYPE_INT32  = 0x07,
	NANOKV_TYPE_INT64  = 0x08,
	NANOKV_TYPE_FLOAT  = 0x09,
	NANOKV_TYPE_DOUBLE = 0x0a,
	NANOKV_TYPE_STRING = 0x0b
};

/** Return codes */
enum nanokv_ret {
	NANOKV_OK            =  0, /** Transaction succesfull */
	NANOKV_NEW           =  1, /** Transaction resulted in new allocation */
	NANOKV_ENOTFOUND  = -1, /** Transaction failed due to key not found in index */
	NANOKV_ETYPE      = -2, /** Transaction failed due to value not matching with index */
	NANOKV_ENOMEM     = -3, /** Transaction failed due to out of memory while allocating */
	NANOKV_EOPNOTSUPP = -4, /** Operation not supported on plugged backend */
	NANOKV_EINVAL     = -5, /** Invalid argument(s) */
};

/** Pluggable backend */
typedef struct {
	const char *name;
	enum nanokv_ret (*init)(void *arg);
	enum nanokv_ret (*destroy)(void *arg);
	enum nanokv_ret (*flush)(void *arg);
	enum nanokv_ret (*get)(void *arg, enum nanokv_type type, const char *key, uint8_t *val);
	enum nanokv_ret (*set)(void *arg, enum nanokv_type vtype, const char *key, uint8_t *val, uint8_t len);
	enum nanokv_ret (*del)(void *arg, const char *key);
	enum nanokv_ret (*incr)(void *arg, const char *key);
	enum nanokv_ret (*incrby)(void *arg, const char *key, uint64_t increment);
	enum nanokv_ret (*decr)(void *arg, const char *key);
	enum nanokv_ret (*decrby)(void *arg, const char *key, uint64_t decrement);
} nanokv_backend_t;

/** Context */
typedef struct {
	nanokv_backend_t *backend;
	void *private_data;
} nanokv_t;

/* Initialize the datastore with a backend */
enum nanokv_ret nanokv_init(nanokv_t *store, nanokv_backend_t *backend);

/** Set private data for backend */
void nanokv_set_private(nanokv_t *store, void *private_data);

/** Get private data for backend */
void *nanokv_get_private(nanokv_t *store);

/** Destroy the instance */
enum nanokv_ret nanokv_destroy(nanokv_t *store);

/* Delete all the keys of the currently selected DB. This command never fails.
The time-complexity for this operation is O(N), N being the number of keys in the database. */
enum nanokv_ret nanokv_flush(nanokv_t *store);

#define nanokv_set(store, key, val) \
	_Generic((val),\
		   uint8_t *: _nanokv_set(store, NANOKV_TYPE_UINT8,  key, (uint8_t *)val, sizeof(uint8_t)),  \
		  uint16_t *: _nanokv_set(store, NANOKV_TYPE_UINT16, key, (uint8_t *)val, sizeof(uint16_t)), \
		  uint32_t *: _nanokv_set(store, NANOKV_TYPE_UINT16, key, (uint8_t *)val, sizeof(uint32_t)), \
		  uint64_t *: _nanokv_set(store, NANOKV_TYPE_UINT16, key, (uint8_t *)val, sizeof(uint64_t)), \
		    int8_t *: _nanokv_set(store, NANOKV_TYPE_INT8,   key, (uint8_t *)val, sizeof(int8_t)),   \
		   int16_t *: _nanokv_set(store, NANOKV_TYPE_INT16,  key, (uint8_t *)val, sizeof(int16_t)),  \
		   int32_t *: _nanokv_set(store, NANOKV_TYPE_INT32,  key, (uint8_t *)val, sizeof(int32_t)),  \
		   int64_t *: _nanokv_set(store, NANOKV_TYPE_INT64,  key, (uint8_t *)val, sizeof(int64_t)),  \
		     float *: _nanokv_set(store, NANOKV_TYPE_FLOAT,  key, (uint8_t *)val, sizeof(float)),    \
		    double *: _nanokv_set(store, NANOKV_TYPE_DOUBLE, key, (uint8_t *)val, sizeof(double)),   \
		      char *: _nanokv_set(store, NANOKV_TYPE_STRING, key, (uint8_t *)val, 0),                \
		char[sizeof(val)]: _nanokv_set(store, NANOKV_TYPE_STRING, key, (uint8_t *)val, 0),           \
		const char *: _nanokv_set(store, NANOKV_TYPE_STRING, key, (uint8_t *)val, 0),                \
		const char[sizeof(val)]: _nanokv_set(store, NANOKV_TYPE_STRING, key, (uint8_t *)val, 0)      \
	)
enum nanokv_ret _nanokv_set(nanokv_t *store, enum nanokv_type vtype, const char *key, uint8_t *val, uint8_t len);

#define nanokv_get(store, key, val) \
	_Generic((val),\
		 uint8_t *: _nanokv_get(store, NANOKV_TYPE_UINT8,  key, (uint8_t *)val), \
		uint16_t *: _nanokv_get(store, NANOKV_TYPE_UINT16, key, (uint8_t *)val), \
		uint32_t *: _nanokv_get(store, NANOKV_TYPE_UINT32, key, (uint8_t *)val), \
		uint64_t *: _nanokv_get(store, NANOKV_TYPE_UINT64, key, (uint8_t *)val), \
		  int8_t *: _nanokv_get(store, NANOKV_TYPE_INT8,   key, (uint8_t *)val), \
		 int16_t *: _nanokv_get(store, NANOKV_TYPE_INT16,  key, (uint8_t *)val), \
		 int32_t *: _nanokv_get(store, NANOKV_TYPE_INT32,  key, (uint8_t *)val), \
		 int64_t *: _nanokv_get(store, NANOKV_TYPE_INT64,  key, (uint8_t *)val), \
		   float *: _nanokv_get(store, NANOKV_TYPE_FLOAT,  key, (uint8_t *)val), \
		  double *: _nanokv_get(store, NANOKV_TYPE_DOUBLE, key, (uint8_t *)val), \
		    char *: _nanokv_get(store, NANOKV_TYPE_STRING, key, (uint8_t *)val)  \
	)
enum nanokv_ret _nanokv_get(nanokv_t *store, enum nanokv_type type, const char *key, uint8_t *val);

enum nanokv_ret nanokv_del(nanokv_t *store, const char *key);
enum nanokv_ret nanokv_incr(nanokv_t *store, const char *key);
enum nanokv_ret nanokv_incrby(nanokv_t *store, const char *key, uint64_t increment);
enum nanokv_ret nanokv_decr(nanokv_t *store, const char *key);
enum nanokv_ret nanokv_decrby(nanokv_t *store, const char *key, uint64_t decrement);

#endif /* NANOKV_H_ */
