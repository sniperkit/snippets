cmake_minimum_required(VERSION 2.6)
enable_language(C CXX)
set(CMAKE_C_STANDARD 99)
set(CMAKE_CXX_STANDARD 11)

project(eresp)

set(CMAKE_C_FLAGS "${CMAKE_C_FLAGS} -Wall -Wextra -Werror -Wunused-parameter -Wmissing-declarations -Wstrict-prototypes -Wmissing-prototypes -Wsign-compare -Wconversion -Wshadow -Wcast-align -Wparentheses -Wsequence-point -Wundef -Wunused -Wpointer-arith -Wnested-externs -Wredundant-decls -Wdisabled-optimization -funit-at-a-time -g -O0 --coverage")
set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -Wall -Wextra -Werror")


include_directories(src)
add_executable(simple examples/simple.c src/eresp.c)
add_library(eresp_static STATIC src/eresp.c)

set(CMAKE_EXE_LINKER_FLAGS "${CMAKE_EXE_LINKER_FLAGS} --coverage")

if (${CMAKE_BUILD_TYPE} STREQUAL "Debug")
	enable_testing()
	add_subdirectory(tests)
endif()
