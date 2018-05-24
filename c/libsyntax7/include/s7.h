#ifndef S7_H_
#define S7_H_

#include <stdlib.h>
#include <stdbool.h>

#ifdef __cplusplus
extern "C" {
#endif

#define S7_TOKEN_DECL(enumerator, token, string)
#define S7_TOKEN_DECL_LIST                                                \
	S7_TOKEN_DECL(S7_TT_UNKNOWN,           '\0', "unknown")           \
	S7_TOKEN_DECL(S7_TT_TEXT,              '\0', "text")              \
	S7_TOKEN_DECL(S7_TT_INTEGER,           '\0', "integer number")    \
	S7_TOKEN_DECL(S7_TT_FLOAT,             '\0', "float number")      \
	S7_TOKEN_DECL(S7_TT_TAB,               '\t', "tab")               \
	S7_TOKEN_DECL(S7_TT_SPACE,             ' ',  "space")             \
	S7_TOKEN_DECL(S7_TT_SEMICOLON,         ';',  "semicolon")         \
	S7_TOKEN_DECL(S7_TT_NEWLINE,           '\n', "newline")           \
	S7_TOKEN_DECL(S7_TT_DOLLAR,            '$',  "dollar")            \
	S7_TOKEN_DECL(S7_TT_COMMA,             ',',  "comma")             \
	S7_TOKEN_DECL(S7_TT_LITERAL_SINGLE,    '\'', "literal single")    \
	S7_TOKEN_DECL(S7_TT_LITERAL_DOUBLE,    '"',  "literal double")    \
	S7_TOKEN_DECL(S7_TT_BRACKET_OPEN,      '[',  "bracket open")      \
	S7_TOKEN_DECL(S7_TT_BRACKET_CLOSE,     ']',  "bracket close")     \
	S7_TOKEN_DECL(S7_TT_BRACE_OPEN,        '{',  "brace open")        \
	S7_TOKEN_DECL(S7_TT_BRACE_CLOSE,       '}',  "brace close")       \
	S7_TOKEN_DECL(S7_TT_PARENTHESES_OPEN,  '(',  "parenthesis open")  \
	S7_TOKEN_DECL(S7_TT_PARENTHESES_CLOSE, ')',  "parenthesis close")
#undef S7_TOKEN_DECL

#define S7_TOKEN_DECL(enumerator, token, string) enumerator,
enum s7_token_types {
S7_TOKEN_DECL_LIST
};
#undef S7_TOKEN_DECL

#define S7_TOKEN_DECL(enumerator, token, string) string,
static const char *s7_token_names[] = {
S7_TOKEN_DECL_LIST
};
#undef S7_TOKEN_DECL

struct s7_token {
	enum s7_token_types type;
	char *value;
	size_t len;
	size_t cap;
	void *private_data; /** Private data for the scanner */
	struct s7_token *next;
};

struct s7_dictionary {
	const char *name;
	const char **keywords;
	const size_t keywords_len;
	const char **builtins;
	const size_t builtins_len;
};

extern const struct s7_dictionary s7_dict_c_language;

struct s7_token *s7_token_scan(const char *text);
struct s7_token *s7_token_new(struct s7_token **prev, enum s7_token_types type);
void s7_token_free(struct s7_token **head);
void s7_token_dump(struct s7_token *head);
void s7_token_append(struct s7_token *t, char c);
bool s7_token_is_keyword(const struct s7_token *token, const struct s7_dictionary *dict);
bool s7_token_is_builtin(const struct s7_token *token, const struct s7_dictionary *dict);

#ifdef __cplusplus
}
#endif

#endif /* S7_H_ */
