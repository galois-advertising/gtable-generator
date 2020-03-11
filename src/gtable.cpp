#include "gtable.h"

namespace galois::gtable {

gtable_project::gtable_project() {

}

gtable_project::~gtable_project() {

}

int gtable_project::create() {
    return 0;
}
int gtable_project::initialize() {
    for (auto ds : _mutable_datasources()) {
        if (!ds.second->create(this)) {
            FATAL("Failed to create datasource[%s]", ds.first.c_str());
            return -1;
        }
    }
    if (!dataupdator_linkto_dataview()) {
        FATAL("Failed to dataupdator_linkto_dataview", "");
        return -1;
    }
    return 0;
}


int gtable_project::load_base() {
    return 0;
}

int gtable_project::reload() {
    return 0;
}

int gtable_project::flush() {
    return 0;
}

}