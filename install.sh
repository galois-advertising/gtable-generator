go build -o gtable-generator *.go
cp gtable-generator /usr/local/bin
rm -rf /usr/local/bin/gtable-generator-include
cp -r include /usr/local/bin/gtable-generator-include
rm -rf /usr/local/bin/gtable-generator-src
cp -r src /usr/local/bin/gtable-generator-src
rm -rf /usr/local/bin/gtable-generator-templates
cp -r templates /usr/local/bin/gtable-generator-templates