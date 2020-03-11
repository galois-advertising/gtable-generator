#pragma once
// solopointer1202@gmail.com
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
    using dataupdator_t = std::shared_ptr<idataupdator<traits>>;


public:
    explicit dataview(const std::string& name) : _name(name) {}

    bool append_dataupdator(dataupdator_t p_notifier) {
        if (p_notifier == nullptr) {
            return false;
        }
        dataupdators.push_back(p_notifier);
        return true;
    }

    bool notify_insert(const insert_raw_t& raw, const insert_derivative_t& derivative) {
        for (auto du : dataupdators) {
            if (!du->notify_insert(raw, derivative)) {
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

    const std::string& name() const {
        return _name;
    }

private:
    std::string _name;
    std::list<dataupdator_t> dataupdators;
};

} 