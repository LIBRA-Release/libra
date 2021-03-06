cmake_minimum_required(VERSION 3.9)
project(${CMAKE_PROJECT_NAME}-test)

set(THREADS_PREFER_PTHREAD_FLAG ON)

find_package(Threads REQUIRED)

add_executable(libra test_libra.cpp)

include_directories(${CMAKE_CURRENT_SOURCE_DIR}/../common/include)

link_directories(${CMAKE_CURRENT_SOURCE_DIR}/../common/lib/linux)

link_directories(/home/runner/work/libra/libra/common/lib/linux)

target_link_libraries(libra Threads::Threads common crypto)

install(TARGETS libra DESTINATION bin)
