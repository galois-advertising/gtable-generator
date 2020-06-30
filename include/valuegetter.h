//solopointer1202@gmail.com
#pragma once
#include "log.h"

namespace galois::gtable {

template <class T>
class value_getter{};

template <class T>
class value_getter<std::vector<T>> {

};


template <class TMessenger>
class value_getter {
public:
    virtual ~value_getter() = 0;
    virtual const char* name() const = 0;
    virtual void* get_value(TMessenger* msgr) const = 0;
    virtual const bool is_enumerate() const = 0;
    virtual void next() const = 0;
    virtual const bool is_end() const = 0;
    virtual void reset() = 0;
};

template <class TMessenger>
class field_getter : public value_getter<TMessenger> {
public:
    typedef typename TMessenger::TTableIters TTableIters;

public:
    virtual ~field_getter() {
    }

    void* get_value(TMessenger* msgr) const {
        if (msgr == nullptr) {
            return nullptr;
        }
        return get_field_value(&msgr->table_iters);
    }

    const bool is_enumerate() const {
        return false;
    }

    void next() const {
    }

    const bool is_end() const {
        return true;
    }

    void reset() {
    }
protected:
    virtual void* get_field_value(TTableIters* table_iters) const = 0;
};

template <class TMessenger>
class placeholder_getter : public value_getter<TMessenger> {
protected:
    mutable int _param_pos;
public:
    typedef typename TMessenger::TQueryData TQueryData;
    placeholder_getter(): _param_pos(0) {
    }

    virtual ~placeholder_getter() {
    }

    void set_param_pos(int pos) {
        _param_pos = pos;
    }

    void* get_value(TMessenger* msgr) const {
        if (msgr) {
            return get_param_value(&msgr->query_data);
        }
        return nullptr;
    }

    const bool is_enumerate() const {
        return false;
    }

    void next() const {
    }

    const bool is_end() const {
        return true;
    }

    void reset() {
    }
protected:
    virtual void* get_param_value(TQueryData* query_data) const = 0;
};

template <class TMessenger>
class arrayplaceholder_getter : public value_getter<TMessenger> {
public:
    typedef typename TMessenger::TQueryData TQueryData;

public:
    arrayplaceholder_getter()
        : _param_pos(0),
        _inner_idx(0),
        _inner_size(0)
    {
        // do nothing
    }

    virtual ~arrayplaceholder_getter()
    {
        // do nothing
    }

    void set_param_pos(int pos)
    {
        _param_pos = pos;
    }

    void* get_value(TMessenger* msgr) const
    {
        if (msgr == nullptr) {
            return nullptr;
        }

        return get_param_value(&msgr->query_data);
    }

    const bool is_enumerate() const 
    {
        return true;
    }

    void next() const
    {
        _inner_idx = (_inner_idx < _inner_size) ? (_inner_idx + 1) : _inner_idx;
    }

    const bool is_end() const
    {
        AQL_WRITE_LOG_TRACE("inner_idx[%u] equal inner_size[%u]", _inner_idx, _inner_size);
        return (_inner_idx == _inner_size);
    }

    void reset() 
    {
        _inner_idx = 0;
        _inner_size = 0;
    }

protected:
    virtual void* get_param_value(TQueryData* query_data) const = 0;

protected:
    mutable int _param_pos;
    mutable int _inner_idx;
    mutable int _inner_size;
};

template <class TMessenger>
class multivalue_getter {
public:
    typedef std::vector<value_getter<TMessenger>*> TVectorType;

public:
    multivalue_getter() {
    }

    ~multivalue_getter() {
    }

    void append_value_getter(value_getter<TMessenger>* value_getter) {
        if (value_getter == nullptr) {
            return;
        }
        _multi_value_getters.push_back(value_getter);
    }

    const TVectorType& get_value_getters() const {
        return _multi_value_getters;
    }

    void destroy() {
        _multi_value_getters.clear();
    }

    // only if one value_getter is_enumerate
    const bool is_enumerate() const {
        if (_multi_value_getters.size() == 0) {
            return false;
        }
        return _multi_value_getters[0]->is_enumerate();
    }

    void next() const {
        typename TVectorType::const_iterator iter = _multi_value_getters.begin();
        for(; iter != _multi_value_getters.end(); ++iter) {
            (*iter)->next();
        }
    }

    const bool is_end() const {
        typename TVectorType::const_iterator iter = _multi_value_getters.begin();
        for(; iter != _multi_value_getters.end(); ++iter) {
            if (!(*iter)->is_end()) {
                return false;
            }
        }
        return true;
    }

private:
    TVectorType _multi_value_getters;
};

} 