#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <ctype.h>

#include "s7.h"

struct s7_token *s7_token_scan(const char *text)
{
	const char *c = text;
	struct s7_token *tlist = NULL;
	struct s7_token *tcur  = NULL;

	while (*c) {
		switch (*c) {
		case '{':
			tcur = s7_token_new(&tlist, S7_TT_BRACE_OPEN);
			s7_token_append(tcur, *c);
			tcur = NULL;
		break;
		case '}':
			tcur = s7_token_new(&tlist, S7_TT_BRACE_CLOSE);
			s7_token_append(tcur, *c);
			tcur = NULL;
		break;
		case '[':
			tcur = s7_token_new(&tlist, S7_TT_BRACKET_OPEN);
			s7_token_append(tcur, *c);
			tcur = NULL;
		break;
		case ']':
			tcur = s7_token_new(&tlist, S7_TT_BRACKET_CLOSE);
			s7_token_append(tcur, *c);
			tcur = NULL;
		break;
		case '(':
			tcur = s7_token_new(&tlist, S7_TT_PARENTHESES_OPEN);
			s7_token_append(tcur, *c);
			tcur = NULL;
		break;
		case ')':
			tcur = s7_token_new(&tlist, S7_TT_PARENTHESES_CLOSE);
			s7_token_append(tcur, *c);
			tcur = NULL;
		break;
		case ',':
			if (!tcur || (tcur && tcur->type == S7_TT_TEXT))
				tcur = s7_token_new(&tlist, S7_TT_COMMA);
			s7_token_append(tcur, *c);
		break;
		case ';':
			if (!tcur || (tcur && (tcur->type == S7_TT_TEXT || tcur->type == S7_TT_LITERAL_SINGLE || tcur->type == S7_TT_LITERAL_DOUBLE || tcur->type == S7_TT_INTEGER || tcur->type == S7_TT_FLOAT)))
				tcur = s7_token_new(&tlist, S7_TT_SEMICOLON);
			s7_token_append(tcur, *c);
		break;
		case '\t':
			if (!tcur || (tcur && tcur->type == S7_TT_TEXT))
				tcur = s7_token_new(&tlist, S7_TT_TAB);
			s7_token_append(tcur, *c);
		break;
		case ' ':
			if (!tcur || (tcur && (tcur->type == S7_TT_TEXT || tcur->type == S7_TT_INTEGER)))
				tcur = s7_token_new(&tlist, S7_TT_SPACE);
			s7_token_append(tcur, *c);
		break;
		case '\n':
			if (!tcur || (tcur && tcur->type == S7_TT_TEXT))
				tcur = s7_token_new(&tlist, S7_TT_NEWLINE);
			s7_token_append(tcur, *c);
		break;
		case '"':
			if (tcur && tcur->type == S7_TT_LITERAL_DOUBLE) {
				tcur = NULL;
				break;
			}
			tcur = s7_token_new(&tlist, S7_TT_LITERAL_DOUBLE);
		break;
		case '\'':
			if (tcur && tcur->type == S7_TT_LITERAL_SINGLE) {
				tcur = NULL;
				break;
			}
			tcur = s7_token_new(&tlist, S7_TT_LITERAL_SINGLE);
		break;
		default:
			if (isdigit(*c)) {
				if (!tcur || (tcur && tcur->type != S7_TT_TEXT && tcur->type != S7_TT_LITERAL_SINGLE && tcur->type != S7_TT_LITERAL_DOUBLE && tcur->type != S7_TT_FLOAT && tcur->type != S7_TT_INTEGER))
					tcur = s7_token_new(&tlist, S7_TT_INTEGER);
			} else if (tcur && tcur->type == S7_TT_INTEGER && (*c == 'e' || *c == 'E' || *c == 'f' || *c == 'F' || *c == '.')) {
				tcur->type = S7_TT_FLOAT;

				if (*c == 'f' || *c == 'F') {
					s7_token_append(tcur, *c);
					tcur = NULL;
					break;
				}
			} else if (!tcur || (tcur && tcur->type != S7_TT_TEXT && tcur->type != S7_TT_LITERAL_SINGLE && tcur->type != S7_TT_LITERAL_DOUBLE && tcur->type != S7_TT_FLOAT))
				tcur = s7_token_new(&tlist, S7_TT_TEXT);

			s7_token_append(tcur, *c);
		break;
		}

		c++;
	}

	return tlist;
}

struct s7_token *s7_token_new(struct s7_token **prev, enum s7_token_types type)
{
	struct s7_token *new;
	struct s7_token *last = NULL;

	if (*prev) {
		struct s7_token *cur = *prev;

		last = cur;
		while (cur) {
			if (cur->next)
				last = cur->next;
			cur = cur->next;
		}
	}

	new = calloc(sizeof(struct s7_token), 1);
	if (new) {
		new->type  = type;
		new->value = NULL;

		if (last)
			last->next = new;
		else
			*prev = new;
	}

	return new;
}

void s7_token_free(struct s7_token **head)
{
	struct s7_token *cur = *head;
	struct s7_token *next;

	while (cur) {
		next = cur->next;
		free(cur->value);
		free(cur);
		cur  = next;
	}

	*head = NULL;
}

void s7_token_dump(struct s7_token *head)
{
	struct s7_token *cur = head;

	while (cur) {
		printf("%p, type(%u): %20s", cur, cur->type, s7_token_names[cur->type]);
		if (cur->type == S7_TT_TEXT || cur->type == S7_TT_LITERAL_DOUBLE)
			printf("[\"%s\"]", cur->value);
		if (cur->type == S7_TT_LITERAL_SINGLE)
			printf("[\'%s\']", cur->value);
		if (cur->type == S7_TT_SPACE || cur->type == S7_TT_TAB)
			printf("[len: %zu]", cur->len);
		if (cur->type == S7_TT_INTEGER || cur->type == S7_TT_FLOAT)
			printf("[%s]", cur->value);
		printf("\n");
		cur = cur->next;
	}
}

void s7_token_append(struct s7_token *t, char c)
{
	if (t->cap == 0) {
		t->value = calloc(1024, 1);
		t->cap   = 1024;
	} else if (t->cap == t->len) {
		t->value = realloc(t->value, t->cap + 1024);
		t->cap  += 1024;
	}

	if (t->value) {
		t->value[t->len] = c;
		t->len++;
	}
}

bool s7_token_is_keyword(const struct s7_token *token, const struct s7_dictionary *dict)
{
	if (token->type == S7_TT_TEXT) {
		const char *kw = dict->keywords[0];

		for (size_t n = 0; n < dict->keywords_len; n++, kw = dict->keywords[n]) {
			if (strcmp(kw, token->value) == 0)
				return true;
		}
	}

	return false;
}

bool s7_token_is_builtin(const struct s7_token *token, const struct s7_dictionary *dict)
{
	if (token->type == S7_TT_TEXT) {
		const char *kw = dict->builtins[0];

		for (size_t n = 0; n < dict->builtins_len; n++, kw = dict->builtins[n]) {
			if (strcmp(kw, token->value) == 0)
				return true;
		}
	}

	return false;
}
