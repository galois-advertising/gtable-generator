// solopointer1202@gmail.com
#pragma once
#include <vector>
#include <string>
#include "log.h"
#include "ValueGetter.h"
#include "../aql-util/AqlContainUtils.h"

namespace galois::gtable {

template <class THandle, class TMessenger>
class conditioner {
public:
    int _seek_round_idx;
    bool _last_ret;
private:
    int _tag_id;

public:
    conditioner() : _seek_round_idx(0), _last_ret(false), _tag_id(0){
    }

    virtual ~conditioner() {
    }

    virtual void reset() {
        _seek_round_idx = 0;
        _last_ret = false;
    }

    virtual const char* name() const = 0;
    virtual bool do_condition(const THandle& handle, TMessenger* msgr) = 0;
    virtual void destroy() = 0;
    int true_condition_count() const;
    void log_monitor_info(bsl::var::Dict& monitor_dict) const;

    void set_tag_id(int tag_id) {
        _tag_id = tag_id;
    }

    int get_tag_id() {
        return _tag_id;
    }

};

template <class THandle, class TMessenger>
class tag_conditioner : public conditioner<THandle, TMessenger> {
public:
    typedef conditioner<THandle, TMessenger> conditioner_t;
private:
    conditioner_t* _sub_conditioner;
    std::string _name;

public:
    tag_conditioner() {
        _sub_conditioner = nullptr;
        _name.clear();
    }

    ~tag_conditioner() {
        _sub_conditioner = nullptr;
        _name.clear();
    }

    const char* name() const {
        return _name.c_str();
    }

    void set_name(const char* name) {
        if (name != nullptr) {
            _name.setf("%s", name);
        }
    }

    void set_sub_conditioner(conditioner_t* conditioner) {
        _sub_conditioner = conditioner;
    }

    bool do_condition(const THandle& handle, TMessenger* msgr) {
        msgr->record_tag(this->get_tag_id(), this->name());
        //NOTE: TAG节点直接透传返回值
        bool ret = _sub_conditioner->do_condition(handle, msgr);
        const char* ret_str = ret ? ",true)" : ",false)";
        msgr->record_debug(ret_str);
        return ret;
    }
    
    void destroy() {
        if (_sub_conditioner != nullptr) {
            _sub_conditioner->destroy();
        }
        _name.clear();
    }

};

template <class THandle, class TMessenger>
class not_conditioner : public conditioner<THandle, TMessenger> {
private:
    conditioner_t* _sub_conditioner;
public:
    typedef conditioner<THandle, TMessenger> conditioner_t;
public:
    not_conditioner() {
    }

    ~not_conditioner() {
    }

    const char* name() const {
        return "not_conditioner";
    }

    void set_sub_conditioner(conditioner_t* conditioner) {
        _sub_conditioner = conditioner;
    }

    bool do_condition(const THandle& handle, TMessenger* msgr) {
        msgr->record_tag(this->get_tag_id());
        //NOTE: 子节点在出错情况下也会返回false，NOT之后会返回true
        //TODO: 需要将错误与逻辑判断从Conditioner的返回中独立出来
        return !(_sub_conditioner->do_condition(handle, msgr));
    }
    
    void destroy() {
        if (_sub_conditioner != nullptr) {
            _sub_conditioner->destroy();
        }
    }
};

template <class THandle, class TMessenger>
class or_conditioner : public conditioner<THandle, TMessenger> {
private:
    std::vector<conditioner_t*> _conditioners;
public:
    typedef conditioner<THandle, TMessenger> conditioner_t;

public:
    or_conditioner() {
    }

    ~or_conditioner() {
    }

    const char* name() const {
        return "or_conditioner";
    }

    void append_sub_conditioner(conditioner_t* conditioner) {
        _conditioners.push_back(conditioner);
    }

    bool do_condition(const THandle& handle, TMessenger* msgr) {
        for (auto iter : _conditioners) {
            // tag属性上移到逻辑节点
            msgr->record_tag(this->get_tag_id());
            if ((*iter)->do_condition(handle, msgr) == true) {
                return true;
            }
        }
        return false;
    }

    void destroy() {
        for (auto iter : _conditioners) {
            (*iter)->destroy();
        }
        _conditioners.clear();
    }

};

template <class THandle, class TMessenger>
class and_conditioner : public conditioner<THandle, TMessenger> {

private:
    std::vector<conditioner_t*> _conditioners;
public:
    typedef conditioner<THandle, TMessenger> conditioner_t;
public:
    and_conditioner() {
    }

    ~and_conditioner() {
    }

    const char* name() const {
        return "and_conditioner";
    }

    void append_sub_conditioner(conditioner_t* conditioner) {
        _conditioners.push_back(conditioner);
    }

    bool do_condition(const THandle& handle, TMessenger* msgr) {
        for (auto iter : _conditioners) {
            // tag属性上移到逻辑节点
            msgr->record_tag(this->get_tag_id());
            if ((*iter)->do_condition(handle, msgr) == false) {
                return false;
            }
        }
        return true;
    }

    void destroy() {
        for (auto iter : _conditioners) {
            (*iter)->destroy();
        }
        _conditioners.clear();
    }
};

template <class THandle, class TMessenger>
class unary_conditioner : public conditioner< THandle, TMessenger> {

protected:
    int _ph_pos;
public:
    unary_conditioner() {
    }

    ~unary_conditioner() {
    }

    void set_ph_pos(int pos) {
        _ph_pos = pos;
    }

    void destroy() {
    }
};

template <class THandle, class TMessenger>
class field_conditioner : public conditioner<THandle, TMessenger> {

protected:
    std::vector<int> _ph_poses;
private:
    TValueGetter* _left_value_getter;
    TValueGetter* _right_value_getter;
public:
    typedef ValueGetter<TMessenger> TValueGetter;

public:
    field_conditioner(): _left_value_getter(nullptr), _right_value_getter(nullptr) {
    }

    virtual ~field_conditioner() {
    }

    inline void set_left_value_getter(TValueGetter* value_getter) {
        _left_value_getter = value_getter;
    }

    inline void set_right_value_getter(TValueGetter* value_getter) {
        _right_value_getter = value_getter;
    }

    TValueGetter* left_value_getter() const {
        return _left_value_getter;
    }

    TValueGetter* right_value_getter() const {
        return _right_value_getter;
    }

    void append_ph_pos(int pos) {
        _ph_poses.push_back(pos);
    }

    void destroy() {
    }
};

} 