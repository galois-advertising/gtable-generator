package generator

import "github.com/go-acme/lego/v3/log"

type GqlDefines struct {
	Queries         []Query `xml:"query"`
	Handler         string  `xml:"handler"`
	Namespace       string  `xml:"namespace"`
	ParserBuildTime string  `xml:"parser_build_time"`
}

func (gd *GqlDefines) Setup() (ok bool) {
	for i := 0; i < len(gd.Queries); i++ {
		if gd.Queries[i].setup() != true {
			log.Fatalf("[gtable] query [%s] setup failed", gd.Queries[i].Name)
			return false
		}
	}
	return true
}
