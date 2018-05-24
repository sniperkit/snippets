# eRESP - embedded REdis Serialization Protocol

## Test

```
make
printf '*3\r\n$3\r\nSET\r\n$4\r\nBOEM\r\n$4\r\nBATS\r\n' | ./build/Debug/simple
```

## License

MIT
