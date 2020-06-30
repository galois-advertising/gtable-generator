
/* 
This code is generated automatically by gtable-generator, do not modify it. 
solopointer1202@gmail.com
*/
#pragma once
#include <memory>
#include <map>
#include "defines.h"
#include "gtable.h"


namespace {{ .Namespace }} {

{{ range $key, $value := .FieldsMap }}
{{- if $value.IsPlaceHolder }}
using {{ $value.PlaceHolderName -}}_valuegetter = placeholder<>;
{{- else }}
using {{ $value.TableName -}}_{{- $value.FieldName}}_valuegetter = valuegetter<decltype({{ $value.TableName -}}_row_tuple.{{$value.FieldName}})>;
{{- end }}
{{- end }}


}