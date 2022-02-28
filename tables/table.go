package tables

type Table struct {
	Opt   *Option
	Title Cells
	Body  []Cells
}

func (t *Table) String() (out string) {
	if len(t.Title) == 0 {
		return "error table"
	}
	mws := MaxCellsWidths(t.Title)
	for _, bodyCells := range t.Body {
		for idx, c := range MaxCellsWidths(bodyCells) {
			if mws[idx] < c {
				mws[idx] = c
			}
		}
	}

	header := t.Title.Parse(mws, t.Opt)
	out += t.Opt.Contour.SlideHeader(mws...) + header + t.Opt.Contour.SlideCenter(mws...)
	for _, c := range t.Body {
		out += c.Parse(mws, t.Opt)
	}
	out += t.Opt.Contour.SlideFooter(mws...)
	return
}

func emptyTable(opt *Option) *Table {
	return &Table{
		Opt:   opt,
		Title: Cells{},
		Body:  []Cells{},
	}
}
