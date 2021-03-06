cmake_minimum_required(VERSION 3.9)
project(${CMAKE_PROJECT_NAME}-crypto)

find_library(cryptoPath crypto linux)
add_library(crypto SHARED IMPORTED GLOBAL)
set_property(TARGET crypto PROPERTY IMPORTED_LOCATION ${cryptoPath})
set_property(TARGET crypto APPEND PROPERTY INTERFACE_INCLUDE_DIRECTORIES ${CMAKE_CURRENT_SOURCE_DIR}/linux)
#install(TARGETS crypto DESTINATION lib)
