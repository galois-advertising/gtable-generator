#pragma once
#include <memory>
#include <new>
#include "datasource.h"
#include "loader.h"
#include "log.h"

namespace galois::gtable {
template <typename databus_traits>
class databus_datasource : public datasource {
    std::string my_name;
public:
    class my_loader : public galois::gdatabus::loader<databus_traits> {
    public:
        gdatabus::file_path_t databus_root_path() const {return "../test/shared/data/GALOIS";};
        gdatabus::file_path_t stream_path() const {return "./stream/";};
        std::string stream_prefix() const {return "stream_GALOIS_DATA_";};
        gdatabus::file_path_t snapshot_path() const {return "./snapshot/";};
        std::string snapshot_prefix() const {return "snapshot_GALOIS_DATA_";};
    };
    typedef my_loader databus_loader_t;

    std::string name() const {return my_name;};

    explicit databus_datasource(const std::string& _name) : my_name(_name), loader(NULL) {
    }

    virtual ~databus_datasource() {}

    bool create(void* param) {
        loader = std::make_shared<databus_loader_t>();
        if (loader == nullptr) {
            FATAL("Failed to create databus loader handler", "");
            return false;
        }

        //if (loader->create(_loader_env) < 0) {
        //    FATAL("failed to create", "");
        //    return false;
        //}
        //loader->set_update_env(p_handle);
        //loader->set_parse_env(p_handle);
        return false;
    }

    bool init_load(void* param) {
        if (!loader->load_base()) {
            FATAL("Failed to load base", "");
            return false;
        }
        return true;
    }

    bool reload(uint32_t token_num) {
        if (loader != nullptr) {
            //return loader->load_inc();
        } else {
            FATAL("The datasource [%s] is uninitialized",
                this->name().c_str());
        }
        return false;
    }

    int add_record(const galois::gformat::pack_header_t& header,
            const char *data, size_t data_len) {
        //return loader->load_inc(header, data, (uint32_t)data_len);
        return false;
    }
private:
    std::shared_ptr<databus_loader_t> loader;
};

}  
