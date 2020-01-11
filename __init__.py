# -*- coding: utf-8 -*-
"""
solopointer1202@gmail.com
20191222
"""
import os
import sys
from xml.dom.minidom import parseString as xmlparse
import pyddl2xml
import pygql2xml

class ddlparser_t(pyddl2xml.pyddl2xml): 
    def log(self, type, log): 
        sys.stderr.write("[ddl2xml][%s] %s\n" % (type, log))

class gqlparser_t(pygql2xml.pygql2xml): 
    def log(self, type, log): 
        sys.stderr.write("[gql2xml][%s] %s\n" % (type, log))

ddlparser = ddlparser_t()
gqlparser = gqlparser_t()

def process_xml(file_name): 
    with open(file_name, 'r') as f: 
        text = f.read().strip('\n')
    if file_name.endswith('.ddl'): 
        xml = ddlparser.parse(text)
    elif file_name.endswith('.gql'): 
        xml = gqlparser.parse(text)
    result = xmlparse(xml)
    if result.documentElement.nodeName == 'ddl': 
        process_ddl(result)
    elif result.documentElement.nodeName == 'gql': 
        process_gql(result)

def process_gql(xml):
    for node in xml.documentElement.childNodes: 
        print node.nodeName

def process_ddl(xml): 
    for node in xml.documentElement.childNodes: 
        print node.nodeName
    pass

if __name__ == '__main__': 
    file_name = sys.argv[1]
    process_xml(file_name)









