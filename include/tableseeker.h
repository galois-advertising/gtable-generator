// solopointer1202@gmail.com
#pragma once
#include <vector>
#include <string>

#include "Gflags.h"
#include "Macros.h"
#include "Conditioner.h"
#include "ResultFiller.h"
#include "ValueGetter.h"
#include "../aql-util/AqlDebug.h"

namespace galois::gtable {

enum seeker_type_t {
    DATATABLE_SEEKER = 0,
    INDEXTABLE_SEEKER
};

enum seeker_joined_way_t {
    JOIN = 0,
    LEFT_JOIN
};

enum seek_status_t {
    SEEK_SUCC = 0,
    SEEK_ERROR,
    SEEK_SELF_FAIL,
    SEEK_CONDITION_FAIL,
    SEEK_JOIN_FAIL,
    SEEK_FILL_FAIL,
    SEEK_LEFT_JOINED_FAIL
};

enum key_status_t {
    KEY_STATUS_SUCCESS = 0,
    KEY_STATUS_SEEK_FAILED = -1,
    KEY_STATUS_HIT_KEY_LIMIT = -2,
    KEY_STATUS_HIT_EACH_KEY_LIMIT = -3,
    KEY_STATUS_ERROR = -4
};

template <class THandle, class TMessenger>
class TableSeeker {
public:
    typedef typename TMessenger::TQueryResponse TQueryResponse;
    typedef typename ::google::protobuf::RepeatedPtrField<TQueryResponse> TQueryResponseList;
    typedef TableSeeker<THandle, TMessenger> TTableSeeker;
    typedef Conditioner<THandle, TMessenger> TConditioner;
    typedef ResultFiller<TMessenger> TResultFiller;
    typedef ValueGetter<TMessenger> TValueGetter;
    typedef MultiValueGetter<TMessenger> TMultiValueGetter;

public:
    TableSeeker() : _conditioner(nullptr), 
                    _vaddr_primary_key_getter(nullptr),
                    _primary_key_getter(nullptr),
                    _scan_limit_getter(nullptr),
                    _result_limit_getter(nullptr),
                    _each_scan_limit_getter(nullptr),
                    _each_result_limit_getter(nullptr),
                    _left_joined(false),
                    _is_table_seeker_starter(false) {
    }

    virtual ~TableSeeker() {
    }

    seek_status_t seek(const THandle& handle, TMessenger* msgr, TQueryResponse* qry_rsp);

    inline void append_joined_table_seeker(TTableSeeker* table_seeker) {
        _joined_table_seekers.push_back(table_seeker);
    }

    inline void append_result_filler(TResultFiller* result_filler) {
        _result_fillers.push_back(result_filler);
    }

    inline void set_conditioner(TConditioner* conditioner) {
        _conditioner = conditioner;
    }

    inline void set_primary_key_getter(TMultiValueGetter* primary_key_getter) {
        _primary_key_getter = primary_key_getter;
    }

    inline void set_vaddr_primary_key_getter(TValueGetter* vaddr_primary_key_getter) {
        _vaddr_primary_key_getter = vaddr_primary_key_getter;
    }

    inline void set_scan_limit_getter(TValueGetter* scan_limit_getter) {
        _scan_limit_getter = scan_limit_getter;
    }

    inline void set_result_limit_getter(TValueGetter* result_limit_getter) {
        _result_limit_getter = result_limit_getter;
    }

    inline void set_each_scan_limit_getter(TValueGetter* each_scan_limit_getter) {
        _each_scan_limit_getter = each_scan_limit_getter;
    }

    inline void set_each_result_limit_getter(TValueGetter* each_result_limit_getter) {
        _each_result_limit_getter = each_result_limit_getter;
    }

    void set_left_joined(bool left_joined) {
        _left_joined = left_joined;
    }

    bool is_left_joined() {
        return _left_joined;
    }

    void set_is_table_seeker_starter() {
        _is_table_seeker_starter = true;
    }

    const bool is_table_seeker_starter() const {
        return _is_table_seeker_starter;
    }

    virtual seeker_type_t type() const = 0;
    virtual const char* name() const = 0;

    virtual void reset() {
    }

    void destroy() {
        typename std::vector<TTableSeeker*>::iterator iter = _joined_table_seekers.begin();
        for (; iter != _joined_table_seekers.end(); ++iter) {
            (*iter)->destroy();
        }
        _joined_table_seekers.clear();
        if (_conditioner != nullptr) {
            _conditioner->destroy();
        }
        if (_primary_key_getter != nullptr) {
            _primary_key_getter->destroy();
        }
        _result_fillers.clear();
    }

    virtual uint32_t result_limit_num() const {
        return 0xFFFFFFFF;
    };

    virtual uint32_t each_result_limit_num() const {
        return 0xFFFFFFFF;
    };

protected:
    const TMultiValueGetter* primary_key_getter() const {
        return _primary_key_getter;
    }

    const TValueGetter* vaddr_primary_key_getter() const {
        return _vaddr_primary_key_getter;
    }

    const TValueGetter* scan_limit_getter() const {
        return _scan_limit_getter;
    }

    const TValueGetter* result_limit_getter() const {
        return _result_limit_getter;
    }

    const TValueGetter* each_scan_limit_getter() const {
        return _each_scan_limit_getter;
    }

    const TValueGetter* each_result_limit_getter() const {
        return _each_result_limit_getter;
    }

    const bool has_enumerate_primary_key() const {
        if (_primary_key_getter == nullptr) {
            return false;
        }

        return _primary_key_getter->is_enumerate();
    }

    virtual seek_status_t do_seek(const THandle& handle, TMessenger* msgr) = 0;
    virtual bool next_item(const THandle& handle, TMessenger* msgr) = 0;
    virtual bool next_primary_key(const THandle& handle, TMessenger* msgr) const = 0;
    virtual uint64_t get_primary_key_id(const THandle& handle, TMessenger* msgr) = 0;
    virtual int do_monitor(TMessenger* msgr) const = 0;
    virtual int do_recount_hit_result_limit(TMessenger* msgr) const = 0;
    virtual int do_recount_hit_each_result_limit(TMessenger* msgr) const = 0;
    void count_seek_failed(TMessenger* msgr) const {
        if (msgr->p_monitor_info != nullptr) {
            msgr->p_monitor_info->_seek_failed_num += 1;
            // 倒排seek fail不记录
            if (FLAGS_adtable_aql_key_status_monitor && this->type() != INDEXTABLE_SEEKER) {
                msgr->p_monitor_info->key_status.push_back(KEY_STATUS_SEEK_FAILED);
            }
        }
    }
    
private:
    bool call_conditioners(const THandle& handle, TMessenger* msgr);

    seek_status_t call_joined_table_seekers(const THandle& handle, TMessenger* msgr, TQueryResponse* qry_rsp);

    int call_result_fillers(TMessenger* msgr, TQueryResponse* qry_rsp) const;

    inline TQueryResponse* add_one_response_to(TQueryResponseList* qry_rsps) const {
        return qry_rsps->Add();
    }

    inline void remove_last_response_from(TQueryResponseList* qry_rsps) const {
        qry_rsps->RemoveLast();
    }

private:
    std::vector<TTableSeeker*> _joined_table_seekers;
    std::vector<TResultFiller*> _result_fillers;
    TConditioner* _conditioner;
    TValueGetter* _vaddr_primary_key_getter;
    TMultiValueGetter* _primary_key_getter;
    TValueGetter* _scan_limit_getter;
    TValueGetter* _result_limit_getter;
    TValueGetter* _each_scan_limit_getter;
    TValueGetter* _each_result_limit_getter;
    bool _left_joined;
    bool _is_table_seeker_starter;
};

} 

#include "tableseeker.hpp"