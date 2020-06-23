package generator

type JoinedTable struct {
	Name            string   `xml:"name"`
	OnColumn        []string `xml:"on_column"`
	JoinType        string   `xml:"join_type,attr"`
	ScanLimit       string   `xml:"scan_limit,attr"`
	EachScanLimit   string   `xml:"each_scan_limit,attr"`
	ResultLimit     string   `xml:"result_limit,attr"`
	EachResultLimit string   `xml:"each_result_limit,attr"`
}
