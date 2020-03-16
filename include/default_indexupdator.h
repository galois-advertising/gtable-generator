/* 
This code is generated automatically by gtable-generator, do not modify it. 
solopointer1202@gmail.com
*/
#pragma once

namespace galois::gtable {

template <typename datatable_traits, template<class> class indexupdator_traits>
class default_indexupdator : public iindexupdator<datatable_traits> {
public:
    using row_t = datatable_traits::row_t;
private: 
    iindextable_t * _indextable = nullptr;
public:
    void set_indextable(iindextable_t* it) {
       _datatable = it; 
    }

    bool notify_after_insert(const row_t& tuple) {
        if (_datatable) {
           _datatable->
        }
        return true;
    };

    bool notify_before_delete(const row_t& tuple) {
        if (_datatable) {
           _datatable->
        }
        return true;
    };
};

}
