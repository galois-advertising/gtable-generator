#pragma once
#include <vector>
#include <string>

namespace galois::gtable {
// T : Handle type
// I : Info type
typedef unsigned int TAG_ID_T;
enum class TEST_STATUS_T {failed = 0, succeed = 1};
template <class T, class I>
class conditioner {
public:
    int _seek_round_idx;
    bool _last_ret;
private:
    int _tag_id;
public:
    conditioner() : _seek_round_idx(0), _last_ret(false), _tag_id(0) {}
    virtual ~conditioner() {}
    virtual void reset() {
        _seek_round_idx = 0;
        _last_ret = false;
    }
    virtual const char* name() const = 0;
    virtual TEST_STATUS_T test(const T&, I&, bool & result) = 0;
    int true_count() const;
    void set_tag_id(TAG_ID_T id) {
        _tag_id = id;
    }
    TAG_ID_T tag_id() {
        return _tag_id;
    }
};

template <class T, class I>
class tag_conditioner : public conditioner<T, I> {
    typedef conditioner<T, I> conditioner_t;
private:
    conditioner_t * _sub_conditioner;
    std::string _name;
public:
    tag_conditioner() : _sub_conditioner(nullptr), _name() {}
    ~tag_conditioner() {}
    std::string name() const {
        return _name;
    }
    void set_name(const std::string & name) {
        _name = name;
    }
    void set_sub_conditioner(conditioner_t * conditioner) {
        _sub_conditioner = conditioner;
    }
    TEST_STATUS_T test(const T& handle, I& debug, bool & result)
    {
        if (_sub_conditioner == nullptr) {
            return TEST_STATUS_T::failed;
        }
        bool temp;
        if(_sub_conditioner->test(handle, debug, temp) == TEST_STATUS_T::succeed) {
           debug.record_tag(this->get_tag_id(), this->name());
           msgr->record_debug(ret ? ",true)" : ",false)");
           result = temp;
           return TEST_STATUS_T::succeed;
        } else {
           return TEST_STATUS_T::failed;
        }
    }
    
};

template <class T, class I>
class not_conditioner : public conditioner<T, I> {
    typedef conditioner<T, I> conditioner_t;
private:
    conditioner_t * _sub_conditioner;
public:
    not_conditioner() {}
    ~not_conditioner() {}
    const char* name() const {
        return "not";
    }
    void set_sub_conditioner(conditioner_t* conditioner) {
        _sub_conditioner = conditioner;
    }
    TEST_STATUS_T test(const T& handle, I& debug, bool & result) {
        if (_sub_conditioner == nullptr) {
            return TEST_STATUS_T::failed;
        }
        bool temp;
        if(_sub_conditioner->test(handle, debug, temp) == TEST_STATUS_T::succeed) {
           debug.record_tag(this->get_tag_id(), this->name());
           msgr->record_debug(ret ? ",true)" : ",false)");
           result = !temp;
           return TEST_STATUS_T::succeed;
        } else {
           return TEST_STATUS_T::failed;
        }
    }
};

template <class T, class I>
class or_conditioner : public conditioner<T, I> {
    typedef conditioner<T, I> conditioner_t;
private:
    std::vector<conditioner_t*> _conditioners;
public:
    or_conditioner() {}
    ~or_conditioner() {}
    const char* name() const {
        return "or";
    }
    void append_sub_conditioner(conditioner_t * conditioner) {
        _conditioners.push_back(conditioner);
    }
    TEST_STATUS_T test(const T& handle, I& debug, bool & result) {
        for (const auto & c : _conditioners) {
            debug.record_tag(this->get_tag_id());
            bool temp = false;
            if (c->test(handle, debug, temp) == TEST_STATUS_T::succeed) {
                if (temp) {
                    result = temp;
                    return TEST_STATUS_T::succeed;
                }
            } else {
                return TEST_STATUS_T::failed;
            }
        }
        result = false;
        return TEST_STATUS_T::succeed;
    }
};

template <class T, class I>
class and_conditioner : public conditioner<T, I> {
    typedef conditioner<T, I> conditioner_t;
private:
    std::vector<conditioner_t*> _conditioners;
public:
    and_conditioner() {}
    ~and_conditioner() {}
    const char* name() const {
        return "and";
    }
    void append_sub_conditioner(conditioner_t* conditioner) {
        _conditioners.push_back(conditioner);
    }
    TEST_STATUS_T test(const T& handle, I& debug, bool & result) {
        for (const auto & c : _conditioners) {
            debug.record_tag(this->get_tag_id());
            bool temp = false;
            if (c->test(handle, debug, temp) == TEST_STATUS_T::succeed) {
                if (temp == false) {
                    result = false;
                    return TEST_STATUS_T::succeed;
                }
            } else {
                return TEST_STATUS_T::failed;
            }
        }
        result = true;
        return TEST_STATUS_T::succeed;
    }

};

template <class T, class I>
class unary_conditioner : public conditioner<T, I> {
    typedef conditioner<T, I> conditioner_t;
public:
    unary_conditioner() {}
    virtual ~unaryconditioner() {}
    void set_ph_pos(int pos) {
        _ph_pos = pos;
    }
protected:
    int _ph_pos;
};

template <class T, class I>
class field_conditioner : public conditioner<T, I> {
    typedef ValueGetter<I> TValueGetter;
protected:
    std::vector<int> _ph_poses;
private:
    TValueGetter* _left_value_getter;
    TValueGetter* _right_value_getter;

public:
    field_conditioner(): _left_value_getter(nullptr), _right_value_getter(nullptr)  {}
    virtual ~field_conditioner() {}
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
};
} 