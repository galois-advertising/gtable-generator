#include "gtable.h"

namespace galois::gtable {

gtable_project::gtable_project() {

}

gtable_project::~gtable_project() {

}

int gtable_project::initialize() {
    for (auto &ds : _mutable_datasources()) {
        if (!ds.second->create(this)) {
            FATAL("Failed to create datasource[%s]", ds.first.c_str());
            return -1;
        }
    }
    if (!setup_dataupdator()) {
        FATAL("Failed to setup_dataupdator", "");
        return -1;
    }
    return 0;
}


int gtable_project::load_base() {
    for (auto &ds : _mutable_datasources()) {
        TRACE("Datasource [%s] begin load base.", ds.first.c_str());
        if (!ds.second->load_base()) {
            FATAL("Failed to load_base [%s]", ds.first.c_str());
            return -1;
        } else {
            TRACE("Datasource [%s] load_base succeed.", ds.first.c_str());
        }
    }
    return 0;
}

int gtable_project::reload() {
    return 0;
}

int gtable_project::flush() {
    return 0;
}

}