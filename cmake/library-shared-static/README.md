# Cmake static and shared library building

When building both static and shared libraries when the project name is prefixed with `lib`,
 then the targets need to be configured with `set_target_properties`. The real work is done with:

* `set_target_properties(target PROPERTIES PREFIX "")`
* `set_target_properties(target PROPERTIES SUFFIX "")`
* `set_target_properties(target PROPERTIES OUTPUT_NAME libtarget)`

## Output (mac osx)

```
Scanning dependencies of target libfoo-shared
[ 25%] Building C object CMakeFiles/libfoo-shared.dir/foo.c.o
[ 50%] Linking C shared library libfoo.dylib
[ 50%] Built target libfoo-shared
Scanning dependencies of target libfoo-static
[ 75%] Building C object CMakeFiles/libfoo-static.dir/foo.c.o
[100%] Linking C static library libfoo.a
[100%] Built target libfoo-static
```
