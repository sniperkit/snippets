/**
 * Simple example with two commands reading data from stdin
 * usage: `printf "PING HELLO\r\nSET HELLO WORLD\r\nCALC 10 10\r\n" | ./build/Debug/simple`
 */
#include <eresp.h>
#include <stdio.h>
#include <string.h>
#include <inttypes.h>

static void pingCommand(const eresp_ctx_t *ctx, int argc, const eresp_item_t *argv)
{
	(void)argv;
	printf("[%s] PING\n", __func__);

	if (argc == 0)
		eresp_write_string(ctx, "PONG");
	else
		eresp_write_bulkstring(ctx, argv[0].data);
}

static void setCommand(const eresp_ctx_t *ctx, int argc, const eresp_item_t *argv)
{
	(void)ctx;
	(void)argc;

	printf("[%s] SET %s %s\n", __func__, argv[0].data, argv[1].data);
}

// CALC <number u8> <number u8>
// prints the sum of both args
static void calcCommand(const eresp_ctx_t *ctx, int argc, const eresp_item_t *argv)
{
	(void)ctx;
	(void)argc;

	uint8_t arg1 = 0;
	uint8_t arg2 = 0;

	// TODO check errors on strconv
	eresp_item_strconv_u8(&arg1, &argv[0]);
	eresp_item_strconv_u8(&arg2, &argv[1]);

	printf("[%s] CALC %" PRIu8 " + %"PRIu8 " = %u\n", __func__, arg1, arg2, arg1 + arg2);
}

static const struct eresp_cmd_t commandTable[] = {
	ERESP_CMD("PING",pingCommand,0),
	ERESP_CMD("SET",setCommand,2),
	ERESP_CMD("CALC",calcCommand,2),
	ERESP_CMD_LAST
};

static void outputWriter(const eresp_ctx_t *ctx, const char c)
{
	(void)ctx;
	printf("%c", c);
}

int main(void) {
	char rbuf[1024];

	eresp_ctx_t r;
	eresp_init(&r, rbuf, sizeof(rbuf));
	eresp_set_commands(&r, commandTable);
	eresp_set_writer(&r, outputWriter);

	while (1) {
		const int c = fgetc(stdin);
		if (c == EOF)
			break;
		eresp_read_char(&r, (const char)c);
	}

	return 0;
}
