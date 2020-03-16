
/* 
This code is generated automatically by gtable-generator, do not modify it. 
solopointer1202@gmail.com
*/
#pragma once

namespace galois::gtable {

template <typename traits>
class iindexupdator {
public:
    using row_t = traits::row_t;
    bool notify_after_insert(const row_t& tuple) = 0;
    bool notify_before_delete(const row_t& tuple) = 0;
    virtual ~iindexupdator() {};
}

};