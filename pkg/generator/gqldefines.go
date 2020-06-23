package generator

type GqlDefines struct {
	Queries      []Query `xml:"gql>query"`
	ValueGetters []ValueGetter
}

func (gd *GqlDefines) Setup() (ok bool) {
	for i := 0; i < len(gd.Queries); i++ {
		if gd.Queries[i].setup() != true {
			return false
		}
	}
	return true
}
