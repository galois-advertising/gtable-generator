#pragma once

#define DEFINE_HANDLER(type, name) \
    public: \
        std::shared_ptr<type> name() const { \
            return _p_##name; \
        } \
    private: \
        std::shared_ptr<type> _p_##name{std::make_shared<type>(std::string(#name))};

#define BEGIN_DATASOUECE  _datasources_t _datasources = { 
#define DATASOURCE(type) { #type, std::make_shared<type>(std::string(#type)) },
#define END_DATASOURCE };\
   _datasources_t& _mutable_datasources() { return _datasources; };