package gogomicromenu

// PagerDivType is a marker for pager slots.
const PagerDivType = PagerDiv

// PagerState tracks the current page for each paged Div.
type PagerState map[string]int

// PagerInfo holds paging metadata for a Div.
type PagerInfo struct {
	TotalPages int
	Page       int
	PageSize   int
}

// NextPage, PrevPage, and SetPage helpers for PagerState.
func (ps PagerState) NextPage(divName string, totalPages int) {
	if ps[divName] < totalPages-1 {
		ps[divName]++
	}
}
func (ps PagerState) PrevPage(divName string) {
	if ps[divName] > 0 {
		ps[divName]--
	}
}
func (ps PagerState) SetPage(divName string, page, totalPages int) {
	if page < 0 {
		ps[divName] = 0
	} else if page >= totalPages {
		ps[divName] = totalPages - 1
	} else {
		ps[divName] = page
	}
}
