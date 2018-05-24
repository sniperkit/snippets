#include <eresp.h>
#include <stdio.h>
#include <unistd.h>

static void pingCommand(int argc, const eresp_item_t **argv)
{
	(void)argc;
	(void)argv;

	printf("+PONG\r\n");
}

static void setCommand(int argc, const eresp_item_t **argv)
{
	(void)argc;

	printf("SET %s %s", argv[0]->data, argv[1]->data);
}

static const struct eresp_cmd_t commandTable[] = {
	{"PING",pingCommand,0},
	{"SET",setCommand,2}
};
static const size_t commandTableSize = ERESP_ARRAY_SIZE(commandTable);

int main(void) {
	char rbuf[1024];
	char stdinbuf[128];

	eresp_reader_t r;
	eresp_reader_init(&r, rbuf, sizeof(rbuf));

	while (1) {
		ssize_t n = read(0, stdinbuf, sizeof(stdinbuf));
		if (n < 1) {
			return 0;
		}
		eresp_read_data(&r, stdinbuf, (size_t)n);
	}

	return 0;
}
