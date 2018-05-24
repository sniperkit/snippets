#include <stdio.h>
#include <stdint.h>
#include <inttypes.h>

#define set(ctx, uid, val) \
	_Generic(val, \
		uint8_t : _set(ctx, uid, (const uint8_t *)&val, sizeof(uint8_t), 1), \
		float : _set(ctx, uid, (const uint8_t *)&val, sizeof(float), 2), \
		char * : _set(ctx, uid, (const uint8_t *)val, sizeof(float), 3) \
	)

int _set(void *ctx, unsigned int uid, const uint8_t *val, const size_t size, unsigned int type)
{
	(void)ctx;
	(void)uid;
	(void)size;
	(void)val;

	switch (type) {
	case 1:
		printf("uint8_t: %" PRIu8 "\n", *val);
		break;
	case 2:
		printf("float: %f\n", *(const float *)val);
		break;
	case 3:
		printf("char *: %s\n", (const char *)val);
		break;
	default:
		break;
	}
	return 0;
}

struct boem {
	float a;
};

int main(void)
{
	uint8_t u8val = 123;
	float floatval = 1.0;
	struct boem aap;
	aap.a = 133.7;
	char *boemm = "1337";

	set(NULL, 0, u8val);
	set(NULL, 0, aap.a);
	set(NULL, 0, boemm);

	return 0;
}
