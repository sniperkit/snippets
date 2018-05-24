#include <stdio.h>
#include <math.h>
#include <stdint.h>

void build_recursive(int level, const int counts[16], const int remainders[16])
{
	if (level == -1) {
		printf(". ");
	} else if (level == -2) {
		printf("X ");
	} else {
		for (size_t n = 0; n < counts[level]; n++)
			build_recursive(level - 1, counts, remainders);
		if (remainders[level] != 0)
			build_recursive(level - 2, counts, remainders);
	}
}

void euclidian(uint8_t steps, uint8_t pulses)
{
	static int counts[16];
	static int remainders[16];

	int divisor = steps - pulses;
	int level = 0;

	remainders[0] = pulses;

	while (1) {
		counts[level]     = divisor / remainders[level];
		remainders[level+1] = divisor % remainders[level];
		divisor = remainders[level];
		level++;
		if (remainders[level] <= 1)
			break;
	}

	counts[level] = divisor;

	printf("steps %d, pulses %d\n", steps, pulses);

	printf("counts: ");
	for (size_t n = 0; n < steps; n++)
		printf("%d ", counts[n]);
	printf("\n");

	printf("remainders: ");
	for (size_t n = 0; n < steps; n++)
		printf("%d ", remainders[n]);
	printf("\n");

	printf("pattern: ");
	build_recursive(level, counts, remainders);
	printf("\n");
}

int main(void)
{
	euclidian(16, 14);
}
