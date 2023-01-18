package table

func SetHeader(tb Table, cells ...interface{}) Table { return tb.UpdateRow(Headers, 0, cells) }
func SetFoots(tb Table, cells ...interface{}) Table  { return tb.UpdateRow(Foots, 0, cells) }
