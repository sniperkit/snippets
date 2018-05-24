#include <stdio.h>

#define __FILE_REL__ (__FILE__+CMAKE_SOURCE_DIR_LENGTH)

int main(void)
{
	printf("path to file: %s\n", __FILE__);
	printf("path fixed: %s\n", __FILE_REL__);
}
