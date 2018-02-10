package engine

// RecordID is a reference to a specific tuple on a specific page of a specific
// table.
type RecordID struct {
	pid     *PageID
	tupleno int
}

func NewRecordID(pid *PageID, tupleno int) *RecordID {
	return &RecordID{pid: pid, tupleno: tupleno}
}

func (rid *RecordID) TupleNo() int {
	return rid.tupleno
}

func (rid *RecordID) PageID() *PageID {
	return rid.pid
}

// consider implementing equals
