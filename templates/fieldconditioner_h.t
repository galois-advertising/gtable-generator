// solopointer1202@gmail.com 
#pragma once
#include "conditioner.h"
#include "Messenger.h"

namespace {{.Namespace}} {

class {{ .Name -}}_conditioner : public galois::gtable::field_conditioner<{{- .Handler -}}, messenger_t> {
public:
    {{- .SomeDefinitions -}}

    const char* name() const {
        return "{{- .Name -}}";
    }

    bool do_condition(const {{- .Handler -}}& handle, messenger_t* msgr) {
        if (_seek_round_idx == msgr->p_monitor_info->_visited_num) {
            return _last_ret;
        }
        {{- .DoCondition }}
        _seek_round_idx = msgr->p_monitor_info->_visited_num;
        _last_ret = ret;
        return ret;
    }

private:
    {{- .GetValueStmts }}
};

}