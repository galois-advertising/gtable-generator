#include "../{{.Name -}}.h"

{{.Namespace}} {

const char* {{.Name}}_schema::PRIMARY_KEY_FIELDS = "{{.Primary_key.Name}}";

void {{.Name}}_schema::make_object_tuple(
        const {{.Name}}ObjectTupleRef& ref,
        {{.Name}}ObjectTuple* tuple) {
${make_object_tuple_from_ref}
}

void {{.Name}}_schema::make_basic_tuple(
        const {{.Name}}ObjectTuple &tuple,
        {{.Name}}BasicTuple* basic) {
${copy_basic_field}
}

bool {{.Name}}_schema::is_valid_object_tuple(
        const {{.Name}}ObjectTuple &tuple) {
${check_is_valid_fields}
    return true;
}

${def_var_member}

int {{.Name}}::insert_var_pools(
        TBasicTuple &basic, 
        const TObjectTuple &tuple) {
    int succ_idx[TSchema::VAR_POOL_NUM];
    int succ_idx_num = 0;
    do {
${write_var_pool}

        return 0;
    } while (0);

    for (int i = 0; i < succ_idx_num; ++i) {
        int idx = succ_idx[i];
        _p_var_pools[idx]->free(basic._var_ptrs[idx]);
    }
    return -1;
}

int {{.Name}}::do_create_basic_pool(TBasicPool* p_basic_pool)
{
${create_basic_pool}
}
}
