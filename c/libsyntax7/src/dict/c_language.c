#include "s7.h"

static const char *keywords[] = {
	/* default types */
	"signed", "unsigned",
	"union", "struct",
	"enum", "char", "short", "long", "int", "float", "double",
	/* builtins */
	"typedef", "sizeof", "void", "inline", "const",
	/* scope */
	"static", "extern", "register", "volatile", "auto",
	/* branching */
	"break", "continue", "return", "goto",
	/* conditionals */
	"if", "else", "for", "while", "do", "switch", "break", "default",
	/* stdbool.h */
	"bool", "_Bool", "true", "false",
	/* stdint.h */
	"int8_t", "int16_t", "int32_t", "int64_t",
	"uint8_t", "uint16_t", "uint32_t", "uint64_t",
	/* stdint.h limits */
	"INT8_MIN", "INT16_MIN", "INT32_MIN", "INT64_MIN",
	"INT8_MAX", "INT16_MAX", "INT32_MAX", "INT64_MAX",
	"UINT8_MAX", "UINT16_MAX", "UINT32_MAX", "UINT64_MAX",
	/* stdatomic.h */
	"_Atomic",
	"atomic_bool"
};

static const char *builtins[] = {
	"printf",
};

const struct s7_dictionary s7_dict_c_language = {
	.name         = "c",
	.keywords     = keywords,
	.keywords_len = sizeof(keywords)/sizeof(const char *),
	.builtins     = builtins,
	.builtins_len = sizeof(builtins)/sizeof(const char *)
};
