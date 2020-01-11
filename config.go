package generator

import (
    "string"
)
Size_map := map[string]uint32{
    "bool": 1,
    "char": 1,
    "int": 4,
    "long": 8,
    "float": 4,
    "unsigned char": 1,
    "unsigned int": 4,
    "unsigned long": 8,
    "unsigned float": 4,
    "uint8_t": 1,
    "uint16_t": 2,
    "uint32_t": 4,
    "uint64_t": 8,
    "string": 0,
    "vector": 0,
    "binary": 0,
    "unsigned binary": 0,
    "uint32key": 4,
    "sign32": 4,
    "uint64key": 8,
    "sign64": 8,
    "uint96key": 4,
    "sign96": 4,
}

Key_map := map[string]string{
    "uint32key" : "uint32_t",
    "uint64key" : "uint64_t",
}