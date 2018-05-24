# NanoKV

Research project to create Embedded (memory constrained, for microcontrollers) key-value store which will not use malloc/free all over the place. With pluggable backends (in-mem, cow, redis, sophia, riak, etcetera).

[![experimental](http://badges.github.io/stability-badges/dist/experimental.svg)](http://github.com/badges/stability-badges)
[![Build Status](https://travis-ci.org/xor-gate/nanokv.svg?branch=master)](https://travis-ci.org/xor-gate/nanokv) [![Coverage Status](https://coveralls.io/repos/github/xor-gate/nanokv/badge.svg?branch=master)](https://coveralls.io/github/xor-gate/nanokv?branch=master) [![Docs](https://readthedocs.org/projects/nanokv/badge/?version=latest)](http://nanokv.readthedocs.org/en/latest)

**Features (API)**

- [X] Pluggable backends, for researching new/existing storage engines
- [X] Redis-like commands API
- [X] Set/get all cstdint types with type-generics (uint8_t, etc)

# Build

- CMake 2.8
- Clang 3.4
- Gcov report (included)

```
$ make debug
$ make releas
$ make test
```

# License

[MIT](LICENSE)
