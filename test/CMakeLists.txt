cmake_minimum_required(VERSION 3.9)
project(${CMAKE_PROJECT_NAME}-test)

set(THREADS_PREFER_PTHREAD_FLAG ON)

find_package(Threads REQUIRED)

add_executable(libra test_libra.cpp)

add_library(crytolib SHARED IMPORTED)
set_target_properties(
  crytolib
  PROPERTIES
  IMPORTED_LOCATION "./../lib/linux/crypto.so"
)

target_link_libraries(libra PRIVATE common Threads::Threads)

install(TARGETS libra DESTINATION bin)
