cmake_minimum_required(VERSION 3.9)
project(${CMAKE_PROJECT_NAME}-test)

#set(THREADS_PREFER_PTHREAD_FLAG ON)

#find_package(Threads REQUIRED)

add_executable(libra test_libra.cpp)

include_directories(${CMAKE_CURRENT_SOURCE_DIR}/../common/include /home/runner/work/libra/libra/lib/linux)

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

#target_link_libraries(libra PRIVATE Threads::Threads common CRYPTO_LIB)
target_link_libraries(libra PRIVATE common)

#install(TARGETS libra DESTINATION bin)
