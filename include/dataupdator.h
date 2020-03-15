#pragma once
// solopointer1202@gmail.com
#include <string>
namespace galois::gtable {

template <typename traits>
class idataupdator {
public:
    using insert_raw_t = typename traits::insert_raw_t;
    using insert_derivative_t = typename traits::insert_derivative_t;
    using update_raw_t = typename traits::update_raw_t;
    using update_derivative_t = typename traits::update_derivative_t;
    using delete_key_t = typename traits::delete_key_t;

    virtual bool notify_insert(const insert_raw_t&, const insert_derivative_t&) = 0;
    virtual bool notify_update(const update_raw_t&, const update_derivative_t&) = 0;
    virtual bool notify_remove(const delete_key_t&) = 0;
    virtual ~idataupdator() {}
};

}  