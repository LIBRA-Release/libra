cmake_minimum_required(VERSION 3.9)
project(${CMAKE_PROJECT_NAME}-common)

add_library(cryptolib SHARED IMPORTED)
set_target_properties(
  cryptolib
  PROPERTIES
  IMPORTED_LOCATION "./../lib/linux/crypto.so"
)

set(LIBRA_SRC
        src/libra.cpp)

set(LIBRA_HDR
        include)

add_library(common STATIC
        ${LIBRA_SRC} ${CMAKE_THREAD_LIBS_INIT})
#message( ${CMAKE_THREAD_LIBS_INIT} )
        
#target_link_libraries(common ${CMAKE_CURRENT_SOURCE_DIR}/../lib/linux/crypto.so)

#message( ${CMAKE_CURRENT_SOURCE_DIR}/../lib/linux/crypto.so )

#target_include_directories(common PRIVATE ${LIBRA_HDR})

install(TARGETS common DESTINATION lib)
