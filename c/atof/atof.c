// https://stackoverflow.com/questions/28292633/implemented-atof-that-named-atof-beta-but-the-result-isnt-the-same-as-atof
#include <stdio.h>
#include <stdbool.h>

float myatof(const char *s)
{
	int sign = 1;

	if (*s == '-') {
		sign = -1;
		s++;
	} else if (*s == '+') {
		sign = 1;
		s++;
	}

	float integer = 0;
	float frac = 0;
	float fraction = 0;
	float divisor = 1;
	bool integer_flag = true;

	while (*s) {
		const char c = *s;

		if (c >= '0' && c <= '9') {
			if (integer_flag) {
				integer = integer * 10 + (c - '0');
			} else {
				fraction = fraction * 10 + (c - '0');
				divisor *= 10;
			}
		} else if (c == '.') {
			integer_flag = false;
		} else {
			break;
		}

		s++;
	}


	return sign * (integer + (fraction / divisor));
}

int main(void)
{
	printf("%f", myatof("1.337"));
	return 0;
}
