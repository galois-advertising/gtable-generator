#pragma once
#include <string>
#include "pack_header.h" 

namespace galois::gtable {
class datasource {
public:
    virtual std::string name() const = 0;
    virtual bool create(void* p_handle) = 0;
    virtual bool init_load(void* p_handle) = 0;
    virtual bool reload(uint32_t token_num) = 0;
    virtual int add_record(const galois::gformat::pack_header_t& heade,
            const char *data, size_t data_len) = 0;
    virtual ~datasource(){}
};
} 