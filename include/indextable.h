/* 
This code is generated automatically by gtable-generator, do not modify it. 
solopointer1202@gmail.com
*/
#pragma once
#include <unordered_map>
#include <set>
#include "iindextable.h"
#include "log.h"

namespace galois::gtable {

template<typename traits>
class indextable : public iindextable<traits> {
public:
    using row_t = typename traits::row_t;
    using index_key = typename traits::index_key;
    using index_ref = const row_t*;
    using primary_key_t = typename traits::primary_key_t;
    using index_t = std::unordered_map<index_key, std::unordered_map<primary_key_t, index_ref>>;

private:
    index_t indextable;
public:
    bool after_insert(const row_t& tuple) {
        try {
            auto k = traits::make_index_key(tuple);
            auto pk = tuple.primary_key();
            if (auto pos = indextable.find(k); pos == indextable.end()) {
                indextable[k] = {};
            }
            indextable[k][pk] = dynamic_cast<index_ref>(&tuple);
            return true;
        } catch (std::bad_alloc& ) {
            FATAL("Out of memory when after_inseert new index", "");
            return false;
        } catch (...) {
            FATAL("Unknown error when after_inseert new index", "");
            return false;
        }
        return true;
    }

    bool before_delete(const row_t& tuple) {
        return true;
    }

};

};