cmake_minimum_required(VERSION 3.9)
project(${CMAKE_PROJECT_NAME}-test)

set(THREADS_PREFER_PTHREAD_FLAG ON)

find_library(CRYPTO_LIB crypto)

find_package(Threads REQUIRED)

add_executable(libra test_libra.cpp)

include_directories(${CMAKE_CURRENT_SOURCE_DIR}/../common/include)

target_link_libraries(libra PRIVATE ${CRYPTO_LIB} common Threads::Threads )

install(TARGETS libra DESTINATION bin)
