cmake_minimum_required(VERSION 3.9)
project(${CMAKE_PROJECT_NAME}-common)

#set(THREADS_PREFER_PTHREAD_FLAG ON)

#find_package(Threads REQUIRED)

set(LIBRA_SRC
        src/libra.cpp)

set(LIBRA_HDR
        include)
        
set(LIBRA_COMM_HDR
        /home/runner/work/libra/libra/lib/linux)
        
add_library(common STATIC
        ${LIBRA_SRC} ${CMAKE_THREAD_LIBS_INIT} ${LIBRA_HDR} ${LIBRA_COMM_HDR}/crypto.h)

find_library(CRYPTO_LIB
         NAMES 
            crypto.so
         PATHS 
            /home/runner/work/libra/libra/lib/linux
)

#add_library(crypto SHARED IMPORTED)
if ( NOT CRYPTO_LIB )
        message( WARNING "CRYPTO_LIB NOT FOUND in common")
else()
        message( STATUS "CRYPTO_LIB : ${CRYPTO_LIB}")
endif( NOT CRYPTO_LIB )

target_link_libraries(common CRYPTO_LIB)

install(TARGETS common DESTINATION /home/runner/work/libra/libra/libs)
