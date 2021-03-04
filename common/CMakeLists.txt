cmake_minimum_required(VERSION 3.9)
project(${CMAKE_PROJECT_NAME}-common)

find_library(CRYPTO_LIB crypto)

set(LIBRA_SRC
        src/libra.cpp)

set(LIBRA_HDR
        include)

add_library(common STATIC
        ${LIBRA_SRC} ${CMAKE_THREAD_LIBS_INIT})
        
target_link_libraries(common PRIVATE ${CRYPTO_LIB})

install(TARGETS common DESTINATION lib)
