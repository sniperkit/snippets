# bufio-reader

Inspired by Golang `bufio.Scanner`,`io.Reader` and `bufio.ScanWords`.

Compile and run:

```
$ make
$ make test
```

Output:

```
printf "0   1   1   2   3   5   8  13  21  34  55  89 144 233" | ./main
"0"
"1"
"1"
"2"
"3"
"5"
"8"
"13"
"21"
"34"
"55"
"89"
"144"
"233"
```
