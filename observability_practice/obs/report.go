package obs

import "context"

// Information need to report:
// Request Rate and return value type.
// Duration distribution of request.

type ReportItem struct {
}

func NewReport(c context.Context) *ReportItem {
	r := &ReportItem{}
	return r
}

func (r *ReportItem) Report(event string) {

}
