/* 
This code is generated automatically by gtable-generator, do not modify it. 
solopointer1202@gmail.com
*/
#pragma once

namespace galois::gtable {

template <typename dataview_traits, template<class> class dataupdator_traits>
class default_dataupdator : public idataupdator<dataview_traits> {
public:
    using insert_raw_t = typename dataview_traits::insert_raw_t;
    using insert_derivative_t = typename dataview_traits::insert_derivative_t;
    using update_raw_t = typename dataview_traits::update_raw_t;
    using update_derivative_t = typename dataview_traits::update_derivative_t;
    using delete_key_t = typename dataview_traits::delete_key_t;
    using data_table = typename dataupdator_traits<dataview_traits>::data_table;
    using idatatable_t = typename data_table::idatatable_t;

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
                return this->_datatable->insert(row);
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
    bool notify_update(const update_raw_t&, const update_derivative_t&) {
        //typename TDataTable::TPrimaryKey pk;
        //if (0 != traits::make_primary_key(portal_tuple, &pk)) {
        //    FATAL("Make primary key failed.", "");
        //    this->_old_iter.reset();
        //    reset_old_data_tag();
        //    return false;
        //}
        //this->_old_iter = this->_p_data_table->seek(pk);
        //set_old_data_tag(portal_tuple);
        //return true;
        return false;

    }
    bool notify_remove(const delete_key_t&) {
        //typename TDataTable::TPrimaryKey pk;
        //if (0 != traits::make_primary_key(portal_tuple, &pk)) {
        //    FATAL("Make primary key failed.", "");
        //    return false;
        //}
        //return this->_p_data_table->remove(pk);
        return false;

    }
};

}