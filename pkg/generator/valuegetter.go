package generator

type ValueGetter struct {
	Name          string
	Namespace     string
	ParamType     string
	IsPlaceholder bool
	TableType     string
}

func (vg *ValueGetter) Setup() (ok bool) {
	return true
}
