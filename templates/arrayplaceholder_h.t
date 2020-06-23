// solopointer1202@gmail.com
#pragma once
#include "valuegetter.h"
#include "messenger.h"

namespace {{ .Namespace }} {


class {{- .Name -}}_getter: public galois::gtable::arrayplaceholder_getter<messenger_t> {
public:
    {{- .Name -}}_getter(): _value(0) {
    }
    const char* name() const {
        return "{{- .Name -}}_getter";
    }

protected:
    void* get_param_value(TQueryData* query_data) const 
    {
        if (query_data == nullptr) {
            FATAL("query data is nullptr!", "");
            return nullptr;
        }

        _value = 0;
        size_t param_size = query_data->params.size();
        if (_param_pos < 0 || (size_t)_param_pos >= param_size) {
            FATAL("pos[%d] out of range[0, %zu]", _param_pos, param_size);
        } else if (!query_data->params[_param_pos]->is_setted()) {
            FATAL("pos[%d] not setted", _param_pos, param_size);
        } else {
            auto* param = query_data->params[_param_pos];
            if (param != nullptr) {
                _inner_size = param->value_count();
                _value = static_cast<{{- .FieldType -}}>(param->get_value(_inner_idx));
            } else {
                FATAL("pos[%d] has nullptr param", _param_pos);
            }
        }
        return &_value;
    }

private:
    mutable {{ .FieldType }} _value;
};

}

