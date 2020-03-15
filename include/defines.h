#pragma once

#define DEFINE_HANDLER(type, name) \
    public: \
        type& name() { \
            return _handler_##name; \
        } \
    private: \
        type _handler_##name;

#define BEGIN_DATASOUECE  _datasources_t _datasources = { 
#define DATASOURCE(type) { #type, std::shared_ptr<galois::gtable::idatasource>(new type(#type)) },
#define END_DATASOURCE };\
   _datasources_t& _mutable_datasources() { return _datasources; };

#define NORMAL_COLUMN(type, name) type _##name;\
    bool _is_set_##name = false;\
    const type & name() const {return _##name;}\
    void set_##name(const type& _v) {_##name = _v; _is_set_##name = true;}\
    bool is_set_##name() const {return _is_set_##name;}

#define ARRAY_COLUMN(type, name) std::vector<type> _##name;\
    bool _is_set_##name = false;\
    const std::vector<type> & name() const {return _##name;}\
    bool append_##name(const type& item) { _##name.push_back(item); _is_set_##name = true; return true;}\
    bool clear_##name() { _##name.clear(); _is_set_##name = true; return true;}\
    bool is_set_##name() const {return _is_set_##name;}


typedef uint64_t uint64key;
typedef uint32_t uint32key;
