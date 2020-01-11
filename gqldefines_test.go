package generator

import (
	"fmt"
	"testing"
)

func TestGqlDefines(t *testing.T) {
	gd := GqlDefines{dataviews: make(map[string]string)}
	fmt.Println(gd.load(*ddl))
}
