/* 
This code is generated automatically by gtable-generator, do not modify it. 
solopointer1202@gmail.com
*/
#pragma once
#include <list>
#include <string>
#include "dataupdator.h"

namespace galois::gtable {

template <typename traits>
class dataview {
public:
    using insert_raw_t = typename traits::insert_raw_t;
    using insert_derivative_t = typename traits::insert_derivative_t;
    using update_raw_t = typename traits::update_raw_t;
    using update_derivative_t = typename traits::update_derivative_t;
    using delete_key_t = typename traits::delete_key_t;
    using dataupdator_t = idataupdator<traits>*;


public:
    const char* name() {
        return traits::name();
    }

    bool append_dataupdator(dataupdator_t p_notifier) {
        if (p_notifier == nullptr) {
            FATAL("p_notifier is nullptr", "");
            return false;
        }
        dataupdators.push_back(p_notifier);
        return true;
    }

    bool notify_insert(const insert_raw_t& raw, const insert_derivative_t& derivative) {
        DEBUG("begin notify_insert.", "");
        for (auto du : dataupdators) {
            if (!du->notify_insert(raw, derivative)) {
                FATAL("notify_insert fail.", "");
                return false;
            }
        }
        return true;
    }

    bool notify_update(const update_raw_t& raw, const update_derivative_t& derivative) {
        for (auto du : dataupdators) {
            if (!du->notify_update(raw, derivative)) {
                return false;
            }
        }
        return true;
    }

    bool notify_delete(const delete_key_t& key) {
        for (auto du : dataupdators) {
            if (!du->notify_delete(key)) {
                return false;
            }
        }
        return true;
    }

private:
    std::list<dataupdator_t> dataupdators;
};

} 