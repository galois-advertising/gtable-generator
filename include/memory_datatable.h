/* 
This code is generated automatically by gtable-generator, do not modify it. 
solopointer1202@gmail.com
*/
#pragma once
#include <unordered_map>
#include <sstream>
#include <list>
#include "log.h"
#include "idatatable.h"
#include "indexupdator.h"

namespace galois::gtable {

template <typename traits>
class memory_datatable : public idatatable<traits> {
public:
    using row_t = typename traits::row_t;
    using primary_key_t = typename traits::primary_key_t;
    using idatatable_t = idatatable<traits>;
    using indexupdator_t = iindexupdator<traits>*;

    template <typename t>
    static primary_key_t primary_key(const t& row_or_key) {
        return traits::primary_key(row_or_key);
    }

    static bool merge_row(const row_t& new_tuple, row_t& old_tuple) {
        return traits::merge_row(new_tuple, old_tuple);
    }

    bool append_indexupdator(indexupdator_t p_notifier) {
        if (p_notifier == nullptr) {
            FATAL("p_notifier is nullptr", "");
            return false;
        }
        indexupdators.push_back(p_notifier);
        return true;
    }

    bool on_insert(const row_t& tuple) {
        try {
            database[traits::primary_key(tuple)] = tuple;
#ifdef _DEBUG
        std::stringstream ss;
        ss << tuple;
        DEBUG("Insert datatable: %s", ss.str().c_str());
#endif
            for (auto &du : indexupdators) {
                if (!du->notify_after_insert(tuple)) {
                    FATAL("notify_after_insert failed.", "");
                }
            }
        } catch (std::bad_alloc & ) {
            FATAL("Out of memory when inserting new row", "");
        } catch (...) {
            FATAL("Unknown error when inserting new row", "");
        }
        return true;
    }

    bool on_update(const row_t& tuple) {
        auto pk = traits::primary_key(tuple);
        if (auto pos = database.find(pk); pos != database.end()) {
            merge_row(tuple, pos->second);
            on_remove(pk);
            on_insert(tuple);
        } 
        return false;
    }

    bool on_remove(const primary_key_t&pk) {
        if (auto pos = database.find(pk); pos != database.end()) {
            for (auto &du : indexupdators) {
                if (!du->notify_before_delete(pos->second)) {
                    FATAL("notify_before_delete failed.", "");
                }
            }
            database.erase(pos);
        }
        return true;
    }

    const row_t* find(const primary_key_t& pk) {
        if (auto pos = database.find(pk); pos != database.end()) {
            return &pos->second;
        }
        return nullptr;
    }
private:
    std::unordered_map<primary_key_t, row_t> database;
private:
    std::list<indexupdator_t> indexupdators;
};
}