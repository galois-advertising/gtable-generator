/* 
This code is generated automatically by gtable-generator, do not modify it. 
solopointer1202@gmail.com
*/
#include <unordered_map>

namespace galois::gtable {

template<class traits>
class idatatable {
public:
    using row_t = typename traits::row_t;
    using primary_key_t = typename traits::primary_key_t;
public:
    virtual bool on_insert(const row_t& tuple) = 0;
    virtual bool on_update(const row_t& tuple) = 0;
    virtual bool on_remove(const primary_key_t& pk) = 0;
    virtual const row_t* find(const primary_key_t& pk) = 0;
    virtual ~idatatable(){};
};

}