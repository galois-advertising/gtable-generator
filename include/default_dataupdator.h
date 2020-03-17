/* 
This code is generated automatically by gtable-generator, do not modify it. 
solopointer1202@gmail.com
*/
#pragma once
#include "dataupdator.h"

namespace galois::gtable {

template <typename dataview_traits, template<class> class dataupdator_traits>
class default_dataupdator : public idataupdator<dataview_traits> {
public:
    // from dataview_traits
    using insert_raw_t = typename dataview_traits::insert_raw_t;
    using insert_derivative_t = typename dataview_traits::insert_derivative_t;
    using update_raw_t = typename dataview_traits::update_raw_t;
    using update_derivative_t = typename dataview_traits::update_derivative_t;
    using delete_key_t = typename dataview_traits::delete_key_t;

    // from dataupdator_traits
    using data_table = typename dataupdator_traits<dataview_traits>::data_table;
    using idatatable_t = typename data_table::idatatable_t;
    using primary_key_t = typename data_table::primary_key_t;

private: 
    idatatable_t * _datatable = nullptr;
public:
    void set_datatable(idatatable_t* dt) {
       _datatable = dt; 
    }
protected:
    bool notify_insert(const insert_raw_t& original, const insert_derivative_t& derivative) {
        typename data_table::row_t row;
        if (dataupdator_traits<dataview_traits>::create_row_tuple(original, derivative, row)) {
            if (_datatable) {
                return this->_datatable->on_insert(row);
            } else {
                FATAL("_datatable is nill.", "");
                return false;
            }
        } else {
            FATAL("create_row_tuple fail.", "");
            return false;
        }
        return true;

    }

    bool notify_update(const update_raw_t& original, const update_derivative_t& derivative) {
        typename data_table::row_t update_info;
        if (dataupdator_traits<dataview_traits>::create_row_tuple(original, derivative, update_info)) {
            if (_datatable) {
                primary_key_t pk = dataupdator_traits<dataview_traits>::primary_key(update_info);
                auto old_row_p = _datatable->find(pk);
                if (old_row_p != nullptr) {
                    return this->_datatable->on_update(*old_row_p);
                } else {
                    FATAL("Try to update a none-exists row", "");
                    return false;
                }
            } else {
                FATAL("Should not be here, _datatable is nullptr", "");
            }
        } else {
            FATAL("create_row_tuple fail.", "");
        }
        return false;
    }

    bool notify_remove(const delete_key_t& delete_key) {
        primary_key_t pk = dataupdator_traits<dataview_traits>::primary_key(delete_key);
        return this->_datatable->on_remove(pk);
    }
};

}