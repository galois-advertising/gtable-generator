// solopointer1202@gmail.com

#pragma once

#include "adtable/StringOstream.h"

#include "atsc/aql-engine/Macros.h"
#include "atsc/aql-engine/TableSeeker.h"

#include "../atsc/{{- .Handler -}}.h"
#include "Messenger.h"

namespace {{.Namespace}} {

class {{- .Name -}}Seeker : public galois::gtable::TableSeeker<{{- .Handler -}}, messenger_t> {
public:
    typedef galois::gtable::TableSeeker<{{- .Handler -}}, messenger_t> TSelf;

public:
    {{- .Name -}}Seeker() {
    }

    ~{{- .Name -}}Seeker() {
    }

    galois::gtable::seeker_type_t type() const {
        return galois::gtable::DATATABLE_SEEKER;
    }

    const char* name() const {
        return "{{- .Name -}}Seeker";
    }

protected:
    galois::gtable::seek_status_t do_seek(const {{- .Handler -}}& handle, messenger_t* msgr) {
        {{- .Namespace -}}::{{- .Name -}}::TPrimaryKey pk(0);
        int ret = this->get_primary_key(msgr, pk);
        if (ret < 0) {
            msgr->table_iters.${name_lower}_iter.reset();
            AQL_WRITE_LOG_TRACE("get primary key from [{{- .Name -}}] failed!");
            if (msgr->is_debug_query()) {
                msgr->debug_msgr.${name_lower}_info = "_MEM:NULL\t";
            }
            if (is_left_joined()) {
                return galois::gtable::SEEK_LEFT_JOINED_FAIL;
            }
            return galois::gtable::SEEK_SELF_FAIL;
        }

        {{- .Namespace -}}::{{- .Name -}}::Iterator ${name_lower}_iter = handle.${name_lower}()->seek(pk);
        msgr->table_iters.${name_lower}_iter = ${name_lower}_iter;
        if (${name_lower}_iter.is_null()) {
            AQL_WRITE_LOG_TRACE("seek primary key from [{{- .Name -}}] failed!");
            if (msgr->is_debug_query()) {
                msgr->debug_msgr.${name_lower}_info = "_MEM:NULL\t";
            }
            if (is_left_joined()) {
                return galois::gtable::SEEK_LEFT_JOINED_FAIL;
            }
            return galois::gtable::SEEK_SELF_FAIL;
        }

        if (msgr->is_debug_query()) {
            galois::gtable::StringOstream os;
            os << *(${name_lower}_iter);
            msgr->debug_msgr.${name_lower}_info = os.get_string();
        }

        ${debug_record_pk}
        return galois::gtable::SEEK_SUCC;
    }
    
    uint64_t get_primary_key_id(const {{- .Handler -}}& handle, messenger_t* msgr) {
        {{- .Namespace -}}::{{- .Name -}}::TPrimaryKey pk(0);
        int ret = this->get_primary_key(msgr, pk);
        if (ret < 0) {
            AQL_WRITE_LOG_TRACE("get primary key from [{{- .Name -}}] failed!");
            return 0;
        }

        {{- .Namespace -}}::{{- .Name -}}::Iterator ${name_lower}_iter = handle.${name_lower}()->seek(pk);
        if (${name_lower}_iter.is_null()) {
            AQL_WRITE_LOG_TRACE("seek primary key from [{{- .Name -}}] failed!");
            return 0;
        }

        // example: return winfo_table_iter->winfo_id();
        ${return_pk_id}
    }

    bool next_item(const {{- .Handler -}}& handle, messenger_t* msgr) {
        return false;
    }

    bool next_primary_key(const {{- .Handler -}}& handle, messenger_t* msgr) const {
        if (!this->has_enumerate_primary_key()) {
            return false;
        }

        this->primary_key_getter()->next();

        if (this->primary_key_getter()->is_end()) {
            return false;
        }

        // zero item of every key is recorded
        ++msgr->p_monitor_info->_visited_num;

        if (msgr->is_debug_query()) {
            msgr->debug_msgr.reset();
        }
        msgr->_tag_succ = 0;

        return true;
    }

    int do_monitor(messenger_t* msgr) const {
        int monitor_idx = msgr->p_monitor_info->_monitor_name;
        if (monitor_idx != NOTAG && monitor_idx != MAXTAG) {
            if (FLAGS_adtable_aql_key_status_monitor) {
               msgr->p_monitor_info->key_status.push_back(monitor_idx);
            }
            msgr->p_monitor_info->_monitor_vec[monitor_idx]++;
            msgr->p_monitor_info->_monitor_name = NOTAG;
        }
        return 0;
    }

    int do_recount_hit_result_limit(messenger_t* msgr) const {
        if (FLAGS_adtable_aql_key_status_monitor) {
            msgr->p_monitor_info->key_status.push_back(galois::gtable::KEY_STATUS_HIT_KEY_LIMIT);
        }
        ++msgr->p_monitor_info->_hit_result_limit_num;
        return 0;
    }

    int do_recount_hit_each_result_limit(messenger_t* msgr) const {
        if (FLAGS_adtable_aql_key_status_monitor) {
            msgr->p_monitor_info->key_status.push_back(galois::gtable::KEY_STATUS_HIT_EACH_KEY_LIMIT);
        }
        ++msgr->p_monitor_info->_hit_each_result_limit_num;
        return 0;
    }

private:
    int get_primary_key(messenger_t* msgr, {{- .Namespace -}}::{{- .Name -}}::TPrimaryKey& pk) {
        const TSelf::TMultiValueGetter* 
            primary_key_getter = this->primary_key_getter();
        if (primary_key_getter == NULL) {
            AQL_WRITE_LOG_FATAL("%s has no primary key getter!", name());
            return -1;
        }

        const TSelf::TMultiValueGetter::TVectorType& multi_value_getters = 
            primary_key_getter->get_value_getters();
        if (multi_value_getters.size() == 0) {
            AQL_WRITE_LOG_FATAL("%s primary key has no primary getter!", name());
            return -1;
        }
${calc_primary_key_str}
        return 0;
    }

};

}

