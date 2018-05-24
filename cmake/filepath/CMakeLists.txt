cmake_minimum_required(VERSION 2.6)
project(filepath C)

string(LENGTH "${CMAKE_SOURCE_DIR}/" CMAKE_SOURCE_DIR_LENGTH)
add_definitions(-DCMAKE_SOURCE_DIR_LENGTH=${CMAKE_SOURCE_DIR_LENGTH})

add_executable(${PROJECT_NAME} ${PROJECT_NAME}.c)
