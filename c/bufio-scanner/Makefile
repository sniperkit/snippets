main:
	cc -I. -o main main.c buffer.c scanner.c reader.c

test: main test_word test_deadbeef
	@true

test_word:
	printf "0   1   1   2   3   5   8  13  21  34  55  89 144 233" | ./main words
test_deadbeef: main
	printf "\xde\xad\xbe\xef0101\xde\xad\xbe\xef0202\xde\xad\xbe\xef0303\xde\xad\xbe\xef0404" | ./main deadbeef

.PHONY: main test_word test_deadbeef
