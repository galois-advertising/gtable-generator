#pragma once
// solopointer1202@gmail.com
#include <ctime>
#include <ratio>
#include <chrono>
#include <map>
#include <memory>
#include "log.h"
#include "datasource.h"

namespace galois::gtable {

class igtable {
public:
    virtual int create() = 0;
    virtual int initialize() = 0;
    virtual int load_base() = 0;
    virtual int reload() = 0;
};

class gtable_project: public igtable {
public:
    using _datasources_t = std::map<std::string, std::shared_ptr<galois::gtable::datasource>>;
    gtable_project();
    virtual ~gtable_project();
    int create();
    int initialize();
    virtual bool dataupdator_linkto_dataview() = 0;
    virtual _datasources_t& _mutable_datasources() = 0;
    int load_base();
    int reload();
    int flush();
};

}