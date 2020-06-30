package generator

type JoinedTable struct {
	Name            string  `xml:"name"`
	JoinType        string  `xml:"join_type,attr"`
	LeftOnColumns   []Field `xml:"left_on_columns>field"`
	RightOnColumns  []Field `xml:"right_on_columns>field"`
	ScanLimit       Field   `xml:"scan_limit"`
	EachScanLimit   Field   `xml:"each_scan_limit"`
	ResultLimit     Field   `xml:"result_limit"`
	EachResultLimit Field   `xml:"each_result_limit"`
}
