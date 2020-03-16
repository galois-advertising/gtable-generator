/* 
This code is generated automatically by gtable-generator, do not modify it. 
solopointer1202@gmail.com
*/
#pragma once
#include <unordered_map>
#include <sstream>
#include "log.h"
// solopointer1202@gmail.com

#include "idatatable.h"

namespace galois::gtable {

template <typename traits>
class memory_datatable : public idatatable<traits> {
public:
    using row_t = typename traits::row_t;
    using primary_key_t = typename traits::primary_key_t;
    using idatatable_t = idatatable<traits>;
    bool insert(const row_t& tuple) {
        try {
            database[tuple.primary_key()] = tuple;
        } catch (std::bad_alloc & ) {
            FATAL("Out of memory when inserting new row", "");
        } catch (...) {
            FATAL("Unknown error when inserting new row", "");
        }
#ifdef _DEBUG
        std::stringstream ss;
        ss << tuple;
        DEBUG("Insert->%s", ss.str().c_str());
#endif
        return true;
    }
    bool update(const row_t& tuple, row_t& old) {
        if (auto pos = database.find(tuple.primary_key()); pos != database.end()) {
            return traits::update(tuple, pos->second);
        } else {
            return false;
        }
    }
    bool del(const primary_key_t&pk) {
        database.erase(pk);
        return true;
    }
private:
    std::unordered_map<primary_key_t, row_t> database;
};
}