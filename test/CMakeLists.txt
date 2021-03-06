cmake_minimum_required(VERSION 3.9)
project(${CMAKE_PROJECT_NAME}-test)

set(THREADS_PREFER_PTHREAD_FLAG ON)

#find_package(Threads REQUIRED)

SET(CMAKE_CXX_FLAGS -pthread)

add_executable(libra test_libra.cpp)

include_directories(${CMAKE_CURRENT_SOURCE_DIR}/../common/include)
#target_link_libraries(libra ${CMAKE_THREAD_LIBS_INIT} common crypto)
target_link_libraries(libra ${CMAKE_THREAD_LIBS_INIT} common)

set(CMAKE_ARCHIVE_OUTPUT_DIRECTORY ${CMAKE_BINARY_DIR}/lib)
set(CMAKE_LIBRARY_OUTPUT_DIRECTORY ${CMAKE_BINARY_DIR}/lib)
set(CMAKE_RUNTIME_OUTPUT_DIRECTORY ${CMAKE_BINARY_DIR}/bin)

install(TARGETS libra DESTINATION bin)
