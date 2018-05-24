cmake_minimum_required(VERSION 2.6)
enable_language(C)
set(CMAKE_C_STANDARD 99)

project(eresp)

set(CMAKE_C_FLAGS "${CMAKE_C_FLAGS} -Wall -Wextra -Werror -Wunused-parameter -Wmissing-declarations -Wstrict-prototypes -Wmissing-prototypes -Wsign-compare -Wconversion -Wshadow -Wcast-align -Wparentheses -Wsequence-point -Wundef -Wunused -Wpointer-arith -Wnested-externs -Wredundant-decls -Wdisabled-optimization -funit-at-a-time")

include_directories(include)
add_executable(simple examples/simple.c src/reader.c)
