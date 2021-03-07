cmake_minimum_required(VERSION 3.9)
project(${CMAKE_PROJECT_NAME}-common)

set(LIBRA_SRC
        src/libra.cpp)

set(LIBRA_HDR
        include)
        
add_library(common STATIC
        ${LIBRA_SRC} ${CMAKE_THREAD_LIBS_INIT})

#add_library(crypto SHARED IMPORTED)
if ( NOT crypto )
        message( WARNING "crypto NOT FOUND in common")
else()
        message( STATUS "crypto library : ${crypto}")
endif( NOT crypto )

target_link_libraries(common crypto)

install(TARGETS common DESTINATION /home/runner/work/libra/libra/libs)
