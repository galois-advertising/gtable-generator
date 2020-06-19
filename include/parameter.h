//solopointer1202@gmail.com
#pragma once 

#include <vector>
#include "log.h"

namespace galois::gtable {

class parameter {
public:
    virtual ~parameter(){
    }

    virtual void set_value(const uint64_t value) = 0;
    virtual void append_value(const uint64_t value) = 0;
    virtual void set_empty_value() = 0;
    virtual uint64_t get_value() const = 0;
    virtual uint64_t get_value(const size_t idx) const = 0;
    virtual const uint32_t value_count() const = 0;
    virtual bool is_setted() const = 0;
    virtual void reset() = 0;
};

class int_parameter : public parameter {
public:
    int_parameter() : _value(0), _is_setted(false) {
    }

    virtual ~int_parameter() {
    }

    uint64_t get_value() const {
        return _value;
    }

    uint64_t get_value(const size_t idx) const {
        FATAL("int_parameter does not support get_value", "");
        return _value;
    }

    void set_value(const uint64_t value) {
        _value = value;
        _is_setted = true;
    }

    void append_value(const uint64_t value) {
        FATAL("int_parameter does not support append_value", "");
    }

    virtual void set_empty_value() {
        FATAL("int_parameter does not support set_empty_value", "");
    }

    const uint32_t value_count() const {
        return (_is_setted ? 1 : 0);
    }

    bool is_setted() const {
        return _is_setted;
    }

    void reset() {
        _value = 0;
        _is_setted = false;
    }

protected:
    uint64_t _value;
    bool _is_setted;
};

class array_parameter : public parameter {
public:
    array_parameter() : _is_setted(false) {
        _values.clear();
    }

    virtual ~array_parameter() {
        _values.clear();
    }

    uint64_t get_value() const {
        if (_values.size() == 0) {
            return 0;
        }

        return _values[0];
    }

    uint64_t get_value(const size_t idx) const {
        if (idx >= _values.size()) {
            return 0;
        }

        return (_values[idx]);
    }

    void set_value(const uint64_t value) {
        FATAL("array_parameter does not support set_value", "");
    }

    // NOTICE：不要用到tablestarter的数组参数中
    void set_empty_value() {
        _values.clear();
        _is_setted = true;
    }

    void append_value(const uint64_t value) {
        _values.push_back(value);
        _is_setted = true;
    }

    const uint32_t value_count() const {
        return _values.size();
    }

    bool is_setted() const { 
        return _is_setted;
    }

    void reset() {
        _is_setted = false;
        _values.clear();
    }

protected:
    std::vector<uint64_t> _values;
    bool _is_setted;
};

}
