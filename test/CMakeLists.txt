cmake_minimum_required(VERSION 3.9)
project(${CMAKE_PROJECT_NAME}-test)

set(THREADS_PREFER_PTHREAD_FLAG ON)

find_package(Threads REQUIRED)

add_library(cryptolib SHARED IMPORTED)
set_property(
        TARGET cryptolib
        PROPERTY
        IMPORTED_LOCATION "/home/runner/work/libra/libra/lib/linux"
        INTERFACE_INCLUDE_DIRECTORIES "/home/runner/work/libra/libra/lib/linux"
)

add_executable(libra test_libra.cpp)

include_directories(${CMAKE_CURRENT_SOURCE_DIR}/../common/include)

target_link_libraries(libra Threads::Threads common cryptolib)

#install(TARGETS libra DESTINATION bin)
