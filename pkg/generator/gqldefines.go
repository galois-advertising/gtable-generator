package generator

type GqlDefines struct {
	Queries         []Query `xml:"query"`
	Handler         string  `xml:"handler"`
	Namespace       string  `xml:"namespace"`
	ParserBuildTime string  `xml:"parser_build_time"`
	ValueGetters    []ValueGetter
}

func (gd *GqlDefines) Setup() (ok bool) {
	for i := 0; i < len(gd.Queries); i++ {
		if gd.Queries[i].setup() != true {
			return false
		}
	}
	return true
}
