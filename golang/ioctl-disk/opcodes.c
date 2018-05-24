#include <stdio.h>
#include <stdint.h>
#include <linux/hdreg.h>
#include <linux/fs.h>

int main(void) {
	printf("const (\n");
	printf("\tHDIO_GETGEO  0x%04x\n", HDIO_GETGEO);
	printf("\tBLKGETSIZE64 0x%04x\n", BLKGETSIZE64);
	printf("}");

	struct hd_geometry geo;

	printf("//hd_geometry: %zu\n", sizeof(geo));
	printf("//hd_geometry.heads: %zu\n", sizeof(geo.heads));
	printf("//hd_geometry.sectors: %zu\n", sizeof(geo.sectors));
	printf("//hd_geometry.cylinders: %zu\n", sizeof(geo.cylinders));
	printf("//hd_geometry.start: %zu\n", sizeof(geo.start));
}
