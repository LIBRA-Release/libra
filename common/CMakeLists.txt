cmake_minimum_required(VERSION 3.9)
project(${CMAKE_PROJECT_NAME}-common)

set(THREADS_PREFER_PTHREAD_FLAG ON)

find_package(Threads REQUIRED)

set(LIBRA_SRC
        src/libra.cpp)

set(LIBRA_HDR
        include)
        
add_library(common STATIC
        ${LIBRA_SRC} ${CMAKE_THREAD_LIBS_INIT})

#add_library(crypto SHARED IMPORTED)

target_link_libraries(common Threads::Thread crypto)

install(TARGETS common DESTINATION lib)
