cmake_minimum_required(VERSION 3.9)
project(${CMAKE_PROJECT_NAME}-common)

set(LIBRA_SRC
        src/libra.cpp)

set(LIBRA_HDR
        include)
        
add_library(common STATIC
        ${LIBRA_SRC} ${CMAKE_THREAD_LIBS_INIT})

find_library(cryptolib
         NAMES 
            crypto
         PATHS 
            /home/runner/work/libra/libra/lib/linux
)

#add_library(crypto SHARED IMPORTED)
if ( NOT cryptolib )
        message( WARNING "cryptolib NOT FOUND in common")
else()
        message( STATUS "cryptolib : ${cryptolib}")
endif( NOT cryptolib )

target_link_libraries(common cryptolib)

install(TARGETS common DESTINATION /home/runner/work/libra/libra/libs)
