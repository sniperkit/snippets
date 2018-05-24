if (NOT BUILD_TESTING)
        message(STATUS "tests disabled")
        return()
endif()
message(STATUS "tests enabled")

set(TESTS
        backend_dummy
#        backend_redis
)

foreach(test ${TESTS})
        add_executable(${test} ${test}.c)
        add_dependencies(${test} ${PROJECT_NAME}-static)
        target_link_libraries(${test} ${PROJECT_NAME}-static ${HIREDIS_LIBRARIES})
        add_test(NAME "${test}" COMMAND "${CMAKE_CURRENT_BINARY_DIR}/${test}")
endforeach()
