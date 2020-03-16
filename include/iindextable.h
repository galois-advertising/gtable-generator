/* 
This code is generated automatically by gtable-generator, do not modify it. 
solopointer1202@gmail.com
*/
#pragma once

namespace galois::gtable {

template<class traits>
class iindextable {
public:
    using row_t = typename traits::row_t;

    virtual bool after_insert(const row_t& tuple) = 0;
    virtual bool before_delete(const row_t& tuple) = 0;
    virtual ~iindextable(){};
};

};