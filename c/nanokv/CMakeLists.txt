cmake_minimum_required(VERSION 2.8)
project(nanokv C)

set(PROJECT_MAJOR 0)
set(PROJECT_MINOR 0)
set(PROJECT_PATCH 0)

if (${CMAKE_BUILD_TYPE} MATCHES "Debug")
	include(CTest)
	set(CMAKE_C_FLAGS "-g -O0 -fprofile-arcs -ftest-coverage")
	add_custom_target(coverage
		COMMAND ${PROJECT_SOURCE_DIR}/scripts/gcovr -x -r ${CMAKE_SOURCE_DIR} -o cobertura.xml
		WORKING_DIRECTORY ${CMAKE_BINARY_DIR}
	)
endif()

set(NANOKV_SRC
	src/nanokv.c
	src/backend/dummy/dummy.c
#	src/backend/redis/redis.c
)

set(CMAKE_MODULE_PATH   "${CMAKE_MODULE_PATH};${CMAKE_SOURCE_DIR}/cmake")
#find_package(Hiredis REQUIRED)

include_directories(${PROJECT_SOURCE_DIR}/include)
# TODO FIXME XXX -> FindHiredis.cmake
#message(STATUS ${HIREDIS_INCLUDE_DIR})
include_directories(${HIREDIS_INCLUDE_DIR})

add_library(${PROJECT_NAME}-static STATIC ${NANOKV_SRC})
set_target_properties(${PROJECT_NAME}-static PROPERTIES OUTPUT_NAME ${PROJECT_NAME} PREFIX "")

add_library(${PROJECT_NAME}-shared SHARED ${NANOKV_SRC})
set_target_properties(${PROJECT_NAME}-shared PROPERTIES OUTPUT_NAME ${PROJECT_NAME} PREFIX "")
#target_link_libraries(${PROJECT_NAME}-shared ${HIREDIS_LIBRARIES})

add_subdirectory(tests)
