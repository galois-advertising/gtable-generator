/* 
This code is generated automatically by gtable-generator, do not modify it. 
solopointer1202@gmail.com
*/
#pragma once
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
    virtual int initialize() = 0;
    virtual int load_base() = 0;
    virtual int reload() = 0;
};

class gtable_project: public igtable {
public:
    using _datasources_t = std::map<std::string, std::shared_ptr<galois::gtable::idatasource>>;
    gtable_project();
    virtual ~gtable_project();
    int initialize();
    virtual bool setup_dataupdator() = 0;
    virtual _datasources_t& _mutable_datasources() = 0;
    int load_base();
    int reload();
    int flush();
};

}