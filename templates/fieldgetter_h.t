// solopointer1202@gmail.com
#pragma once
#include "valuegetter.h"
#include "messenger.h"

namespace {{ .Namespace }} {

class {{- .Name -}}_getter : public galois::gtable::field_getter<messenger_t> {
public:
    const char* name() const {
        return "{{- .Name -}}_getter";
    }

protected:
    void* get_field_value(TTableIters* table_iters) const {
        if (table_iters == nullptr) {
            return nullptr;
        }

        if (table_iters->{{- .TableName -}}_iter.is_null()) {
            return nullptr;
        }

        return static_cast<void*>(&(table_iters->{{- .TableName -}}_iter->{{- .FieldName -}}()));
    }
};

}