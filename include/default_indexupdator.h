/* 
This code is generated automatically by gtable-generator, do not modify it. 
solopointer1202@gmail.com
*/
#pragma once
#include "indexupdator.h"

namespace galois::gtable {

template <typename datatable_traits, template<class> class indexupdator_traits>
class default_indexupdator : public iindexupdator<datatable_traits> {
public:
    // from datatable_traits
    using row_t = typename datatable_traits::row_t;

    // from indexupdator_traits
    using index_table = typename indexupdator_traits<datatable_traits>::index_table;
    using iindextable_t = typename index_table::iindextable_t;
private: 
    iindextable_t * _indextable = nullptr;
public:
    void set_indextable(iindextable_t* it) {
       _indextable = it; 
    }

    bool notify_after_insert(const row_t& row) {
        if (_indextable) {
            if (!_indextable->after_insert(row)) {
                FATAL("notify_after_insert fail", "");
                return false;
            }
        }
        return true;
    };

    bool notify_before_delete(const row_t& row) {
        if (_indextable) {
           //_indextable->
        }
        return true;
    };
};

}
