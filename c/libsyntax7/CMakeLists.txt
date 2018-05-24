cmake_minimum_required(VERSION 2.8)
project(libsyntax7 C)

include(cmake/CheckCCompilerFlag.txt)
include(cmake/Lint.txt)

add_cflag_if_supported("-std=c99")

if(${CMAKE_BUILD_TYPE} MATCHES "Debug")
	add_cflag_if_supported("-O0")
	add_cflag_if_supported("-ggdb")
	add_cflag_if_supported("-fprofile-arcs")
	add_cflag_if_supported("-fsanitize")
	add_cflag_if_supported("-ftest-coverage")

	add_definitions("-DDEBUG")

	include(CTest)
elseif()
	add_cflag_if_supported("-O2")

	add_definitions("-DRELEASE")
endif()

set(SRC src/scanner.c src/dict/c_language.c)
include_directories(include)

if(OCA_LINT AND BUILD_TESTING)
	add_lint(${SRC})
endif()

add_library(${PROJECT_NAME}-static STATIC ${SRC})
set_target_properties(${PROJECT_NAME}-static PROPERTIES OUTPUT_NAME ${PROJECT_NAME} PREFIX "")

add_library(${PROJECT_NAME}-shared SHARED ${SRC})
set_target_properties(${PROJECT_NAME}-shared PROPERTIES OUTPUT_NAME ${PROJECT_NAME} PREFIX "")

install(TARGETS ${PROJECT_NAME}-static ${PROJECT_NAME}-shared
	LIBRARY DESTINATION lib
	ARCHIVE DESTINATION lib
)
install(DIRECTORY include/ DESTINATION include/syntax7
        FILES_MATCHING PATTERN "*.h")

add_subdirectory("tests")
