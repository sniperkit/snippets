# Runtime reflection in C

Output:

```
decl_get_id(c): 6494383
decl_get_name(c): coord
decl_get_body(c): "size_t monkey; size_t x; size_t y;"
decl_get_type(c, "monkey"): "size_t"
decl_get_type(c, "blaat"): "unknown"
--- member "monkey"
monkey offset: 0
monkey len: 8
monkey type(1): "size_t"
monkey: 1337
--- member "x"
x offset: 8
x len: 8
x type(1): "size_t"
x: 666
--- member "y"
y offset: 16
y len: 8
y type(1): "size_t"
y: 777
--- member "blaat" -> unknown
```
