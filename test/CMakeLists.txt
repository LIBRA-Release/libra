cmake_minimum_required(VERSION 3.9)
project(${CMAKE_PROJECT_NAME}-test)

set(THREADS_PREFER_PTHREAD_FLAG ON)

find_package(Threads REQUIRED)

add_executable(libra test_libra.cpp)

include_directories(${CMAKE_CURRENT_SOURCE_DIR}/../common/include)

link_directories(${CMAKE_CURRENT_SOURCE_DIR}/../common/lib/linux)

find_library(OFLOG_LIB NAMES oflog PATHS CMAKE_CURRENT_SOURCE_DIR}/../common/lib/linux))

target_link_libraries(common Threads::Threads crypto)
target_link_libraries(libra Threads::Threads crypto common)

install(TARGETS libra DESTINATION bin)
