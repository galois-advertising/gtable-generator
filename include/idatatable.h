#include <unordered_map>

namespace galois::gtable {

template<class traits>
class idatatable {
public:
    using row_t = typename traits::row_t;
    using primary_key_t = typename traits::primary_key_t;
public:
    virtual bool insert(const row_t& tuple) = 0;
    virtual bool update(const row_t& tuple, row_t& old) = 0;
    virtual bool del(const primary_key_t& pk) = 0;
};

}