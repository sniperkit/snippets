#include <gtest/gtest.h>
#include <eresp.h>

#define ASSERT_MEMCMP(s1, s2, n) \
	ASSERT_EQ(0, memcmp(s1, s2, n))

static void ASSERT_ERESP_ITEM_EQ(const eresp_item_t *expItem, const eresp_item_t *valItem) {
	ASSERT_EQ(expItem->type, valItem->type);
	ASSERT_EQ(expItem->len,  valItem->len);
	if (expItem->data)
		ASSERT_NE(nullptr, valItem->data);
	if (expItem->data && valItem->data)
		ASSERT_MEMCMP(expItem->data, valItem->data, valItem->len);
}

void eRESPTestBufferedWriter(const eresp_ctx_t *ctx, const char c);

class eRESPTest : public ::testing::Test {
public:
	void Bufc(const char c) {
		obuf[obufn] = c;
		obufn++;
		// TODO protect overflow
	}
	void BufReset() {
		memset(obuf, 0, sizeof(obuf));
		obufn = 0;
	}
protected:
	virtual void SetUp() {
		BufReset();
		eresp_init(&ctx, buf, sizeof(buf));
		eresp_set_context(&ctx, this);
		eresp_set_writer(&ctx, eRESPTestBufferedWriter);
	}
	virtual void TearDown() {}

	char buf[4096];
	char obuf[4096];
	size_t obufn;
	eresp_ctx_t ctx;
};

void eRESPTestBufferedWriter(const eresp_ctx_t *ctx, const char c)
{
	(void)ctx;
	(void)c;

	auto i = static_cast<eRESPTest *>(ctx->context);
	i->Bufc(c);
}

/** Test if the context is correctly initialized with a read buffer */
TEST_F(eRESPTest, InitialWithBuffer) {
	ASSERT_EQ(buf, ctx.buf);
	ASSERT_EQ(buf, ctx.bufc);
	ASSERT_EQ(sizeof(buf), ctx.buflen);
	ASSERT_EQ(0U, ctx.itemslen);
}

/** Test reading a array with a single item */
TEST_F(eRESPTest, ArrayWithSingleItem) {
	const char tv[] = "*1\r\n$4\r\nPING\r\n";
	const struct eresp_item_t tvExp[] = {
		{ ERESP_T_ARRAY, nullptr, 1U},
		{ ERESP_T_BULKSTRING, "PING", 4U}
	};

	eresp_read_data(&ctx, tv, sizeof(tv) - 1);

	ASSERT_EQ(2U, ctx.itemslen);
	ASSERT_ERESP_ITEM_EQ(&tvExp[0], &ctx.items[0]);
	ASSERT_ERESP_ITEM_EQ(&tvExp[1], &ctx.items[1]);
}

/** Test item proc callback */
TEST_F(eRESPTest, ItemProcCb) {
	bool ItemProcCbTest(const eresp_ctx_t *r, const eresp_item_t *item);
	int triggers = 0;

	eresp_set_context(&ctx, &triggers);
	eresp_set_item_proc(&ctx, ItemProcCbTest);

	const char tv[] = "*1\r\n$4\r\nPING\r\n";
	const struct eresp_item_t tvExp[] = {
		{ ERESP_T_ARRAY, nullptr, 1U},
		{ ERESP_T_BULKSTRING, "PING", 4U}
	};

	eresp_read_data(&ctx, tv, sizeof(tv) - 1);

	// Expect 3 callbacks triggers and 2 items
	// * Array start
	// * Bulkstring item finish buffering
	// * Array finish
	ASSERT_EQ(3, triggers);
	ASSERT_EQ(2U, ctx.itemslen);
	ASSERT_ERESP_ITEM_EQ(&tvExp[0], &ctx.items[0]);
	ASSERT_ERESP_ITEM_EQ(&tvExp[1], &ctx.items[1]);
}

bool ItemProcCbTest(const eresp_ctx_t *r, const eresp_item_t *item) {
	(void)r;
	(void)item;
	int *triggers = static_cast<int *>(r->context);
	*triggers += 1;
	return false;
}

/** Test a item against a command list and get the valid one */
TEST_F(eRESPTest, CmdSearchValid) {
	const struct eresp_cmd_t tvCmdList[] = {
		{"PING",4U,NULL,0}, 
		ERESP_CMD_LAST
	};
	const struct eresp_item_t tv = {
		ERESP_T_BULKSTRING,
		(char *)"PING",
		4,
	};

	auto cmd = eresp_cmd_search(tvCmdList, &tv);

	ASSERT_EQ(4U,        cmd->namelen);
	ASSERT_STREQ("PING", cmd->name);
	ASSERT_EQ(nullptr,   cmd->proc);
	ASSERT_EQ(0,         cmd->argc);
}

/** Test two array of items parsing and executing of two commands with valid args */
struct CmdExecutePipelineContext {
	bool SetProcExecuted;
	bool PingProcExecuted;
};

TEST_F(eRESPTest, CmdExecutePipeline) {
	void TestCmdExecuteSetProc(const eresp_ctx_t *ctx, int argc, const eresp_item_t *argv);
	void TestCmdExecutePingProc(const eresp_ctx_t *ctx, int argc, const eresp_item_t *argv);

	const char tvSetCmd[] = "*3\r\n$3\r\nSET\r\n$4\r\nBOEM\r\n$4\r\nBATS\r\n*3\r\n$4\r\nPING\r\n$5\r\nHELLO\r\n$5\r\nWORLD\r\n";
	const struct eresp_cmd_t tvCmdList[] = {
		{"SET",3U,TestCmdExecuteSetProc,2},
		{"PING",4U,TestCmdExecutePingProc,0},
		ERESP_CMD_LAST
	};

	struct CmdExecutePipelineContext context;

	context.SetProcExecuted = false;
	context.PingProcExecuted = false;

	eresp_set_context(&ctx, &context);
	ASSERT_EQ(nullptr, ctx.proc);
	eresp_set_commands(&ctx, tvCmdList);
	// Verify eresp_cmd_item_proc callback function is attached after eresp_set_commands
	ASSERT_EQ(&eresp_cmd_item_proc, ctx.proc);
	eresp_read_data(&ctx, tvSetCmd, sizeof(tvSetCmd) - 1);

	ASSERT_TRUE(context.SetProcExecuted);
	ASSERT_TRUE(context.PingProcExecuted);
}

void TestCmdExecuteSetProc(const eresp_ctx_t *ctx, int argc, const eresp_item_t *argv)
{
	(void)argc;
	(void)argv;
	struct CmdExecutePipelineContext *c = static_cast<struct CmdExecutePipelineContext *>(ctx->context);
	c->SetProcExecuted = true;
}

void TestCmdExecutePingProc(const eresp_ctx_t *ctx, int argc, const eresp_item_t *argv)
{
	(void)argv;
	// The test vector passes two arguments to the PING command
	if (argc != 2)
		return;
	// FIXME check correct arguments length and data
	struct CmdExecutePipelineContext *c = static_cast<struct CmdExecutePipelineContext *>(ctx->context);
	c->PingProcExecuted = true;
}

/** Inline command with two args */
TEST_F(eRESPTest, InlineCmdWithTwoArgs) {
	void TestInlineCmdWithTwoArgsProc(const eresp_ctx_t *ctx, int argc, const eresp_item_t *argv);
	const struct eresp_cmd_t tvCmdList[] = {
		{"SET",3U,TestInlineCmdWithTwoArgsProc,2},
		ERESP_CMD_LAST
	};

	bool triggered = false;

	eresp_set_context(&ctx, &triggered);
	eresp_set_commands(&ctx, tvCmdList);

	const char tv[] = "SET HELLO WORLD\n";
	eresp_read_data(&ctx, tv, sizeof(tv) - 1);

	// Command proc triggered?
	ASSERT_TRUE(triggered);

	// NOTE: we are unable to check the inline buffered items because after the command is executed eresp_reset is executed
}

void TestInlineCmdWithTwoArgsProc(const eresp_ctx_t *ctx, int argc, const eresp_item_t *argv)
{
	(void)argc;
	(void)argv;
	if (argc != 2)
		return;
	bool *triggered = static_cast<bool *>(ctx->context);
	*triggered = true;
}

/**
 * Test the eresp_write_* functions with the buffered output writer
 */
TEST_F(eRESPTest, OutputWriters) {
	// Simplestring
	eresp_write_string(&ctx, "OK");
	ASSERT_STREQ("+OK\r\n", obuf);
	BufReset();

	// Bulkstring data
	const char bindata[] = "BOEMBATS";
	eresp_write_data(&ctx, bindata, sizeof(bindata) - 1);
	ASSERT_STREQ("$8\r\nBOEMBATS\r\n", obuf);
	BufReset();

	// Itemlist of array with 3 bulkstrings
	const struct eresp_item_t itemsTv[] = {
		{ ERESP_T_ARRAY, NULL, 3U},
		{ ERESP_T_BULKSTRING, "FOO", 3U},
		{ ERESP_T_BULKSTRING, "BAR", 3U},
		{ ERESP_T_BULKSTRING, "BAZ", 3U}
	};

	eresp_write_items(&ctx, itemsTv, ERESP_ARRAY_SIZE(itemsTv));
	ASSERT_STREQ("*3\r\n$3\r\nFOO\r\n$3\r\nBAR\r\n$3\r\nBAZ\r\n", obuf);
	BufReset();

	// Item bulkstring, expect "$3\r\nBAZ\r\n"
	eresp_write_item(&ctx, &itemsTv[3]);
	ASSERT_STREQ("$3\r\nBAZ\r\n", obuf);
	BufReset();

	// Error, expect "-ERR unknown command 'foobar'\r\n"
	eresp_write_error(&ctx, "ERR unknown command 'foobar'");
	ASSERT_STREQ("-ERR unknown command 'foobar'\r\n", obuf);
	BufReset();
}

TEST_F(eRESPTest, ItemStrconvUINT) {
	const struct eresp_item_t tv[] = {
		{ERESP_T_BULKSTRING, "-1",   2U},
		{ERESP_T_BULKSTRING, "0",    1U},
		{ERESP_T_BULKSTRING, "1",    1U},
		{ERESP_T_BULKSTRING, "254",  3U},
		{ERESP_T_BULKSTRING, "255",  3U},
		{ERESP_T_BULKSTRING, "256",  3U}
	};

	// u8
	uint8_t u8;

	ASSERT_EQ(ERESP_ETYPE, eresp_item_strconv_u8(&u8, &tv[0]));
	// TODO corner case of zero value
	//ASSERT_EQ(ERESP_OK,    eresp_item_strconv_u8(&u8, &tv[1]));
	//ASSERT_EQ(0, u8);
	ASSERT_EQ(ERESP_OK,    eresp_item_strconv_u8(&u8, &tv[2]));
	ASSERT_EQ(1, u8);
	ASSERT_EQ(ERESP_OK,    eresp_item_strconv_u8(&u8, &tv[3]));
	ASSERT_EQ(254, u8);
	ASSERT_EQ(ERESP_OK,    eresp_item_strconv_u8(&u8, &tv[4]));
	ASSERT_EQ(255, u8);
	ASSERT_EQ(ERESP_ETYPE, eresp_item_strconv_u8(&u8, &tv[5]));

	// u16
	uint16_t u16;

	ASSERT_EQ(ERESP_ETYPE, eresp_item_strconv_u16(&u16, &tv[0]));
	ASSERT_EQ(ERESP_OK, eresp_item_strconv_u16(&u16, &tv[5]));
	ASSERT_EQ(256, u16);
}

// String to bool converter
TEST_F(eRESPTest, ItemStrconvBool) {
	const struct eresp_item_t tv[] = {
		{ERESP_T_BULKSTRING, "1", 1U},
		{ERESP_T_BULKSTRING, "0", 1U}
	};

	bool b = false;

	ASSERT_EQ(ERESP_OK, eresp_item_strconv_bool(&b, &tv[0]));
	ASSERT_TRUE(b);

	b = true;
	ASSERT_EQ(ERESP_OK, eresp_item_strconv_bool(&b, &tv[1]));
	ASSERT_FALSE(b);
}
