#include <s7.h>
#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
#include <string.h>
#include <ctype.h>

#define ANSI_COLOR_RED     "\x1b[31m"
#define ANSI_COLOR_GREEN   "\x1b[32m"
#define ANSI_COLOR_YELLOW  "\x1b[33m"
#define ANSI_COLOR_BLUE    "\x1b[34m"
#define ANSI_COLOR_MAGENTA "\x1b[35m"
#define ANSI_COLOR_CYAN    "\x1b[36m"
#define ANSI_COLOR_RESET   "\x1b[0m"

enum parser_type {
	TYPE_KEYWORD,
	TYPE_BUILTIN
};

struct parser_obj {
	enum parser_type type;
};

static char text[] = "int main(void) {\n\
	uint16_t testvar = UINT16_MAX;\n\
	float testf1 = 1.0;\n\
	float testf2 = 1f;\n\
	float testf3 = 5e9;\n\
	const char text[1234] = \"Hello world\";\n\
	if (true) {\n\
		for (int i = 0; i < 10; i++) {\n\
			printf(\"Hello w%crld\", \'o\');\n\
		}\n\
	}\n\
}\n";

void s7_parse_dump(struct s7_token *head)
{
	const struct s7_token *cur = head;

	while (cur) {
		if (cur->type == S7_TT_LITERAL_SINGLE)
			printf(ANSI_COLOR_CYAN "\'%s\'" ANSI_COLOR_RESET, cur->value);
		else if (cur->type == S7_TT_LITERAL_DOUBLE)
			printf(ANSI_COLOR_CYAN "\"%s\"" ANSI_COLOR_RESET, cur->value);
		else {
			if (s7_token_is_builtin(cur, &s7_dict_c_language))
				printf(ANSI_COLOR_GREEN "%s" ANSI_COLOR_RESET, cur->value);
			else if (s7_token_is_keyword(cur, &s7_dict_c_language))
				printf(ANSI_COLOR_MAGENTA "%s" ANSI_COLOR_RESET, cur->value);
			else if (cur->type == S7_TT_INTEGER || cur->type == S7_TT_FLOAT)
				printf(ANSI_COLOR_YELLOW "%s" ANSI_COLOR_RESET, cur->value);
			else
				printf("%s", cur->value);
		}
		cur = cur->next;
	}
}

int main(void)
{
	struct s7_token *tlist;
	tlist = s7_token_scan(text);
	s7_token_dump(tlist);
	//s7_parse_keywords(tlist);
	//s7_parse_builtins(tlist);

	s7_parse_dump(tlist);
	s7_token_free(&tlist);

	return 0;
}
