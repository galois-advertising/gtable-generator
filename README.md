# what is gtable

gtable(GaloisTable) is a low-latency in-menory database engine implement by c++ for relational data storage especially for advertisement data. Support SQL like pre-compiled query language GQL. You can embed gtable into any c++ project.

# what is gtable-generator

gtable-generator is a tool for translating *.ddl.xml file inito *.h/*.cpp file.
![gtable-generator](./gtable-generator.png)

# install

Installation dependency:
* go1.12.9+

```
$ go build -o gtable-generator *.go
$ cp gtable-generator /usr/local/bin
$ cp -r include /usr/local/bin/gtable-generator-include
$ cp -r include /usr/local/bin/gtable-generator-src
$ cp -r templates /usr/local/bin/gtable-generator-templates
```

# usage

```
$ gtable-generator -i /path/to/*.ddl.xml -o /path/to/output
```