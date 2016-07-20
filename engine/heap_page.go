package engine

// HeapPage stores pages of HeapFiles and implements the Page interface
// that is used by BufferPool
type HeapPage struct {
	pid      *HeapPageID
	td       *TupleDesc
	header   []byte
	tuples   []Tuple
	numSlots int

	oldData []byte
}

// New creates a HeapPage from a set of bytes of data read fromdisk.
// The format of a HeapPage is a set of header bytes indicating the slots of
// the page that are in use, some number of tuple slots.
// where tuple size is the size of tuples in this database table, which can be
// determined via Catalog.GetTupleDesc.
// The number of 8-bit header words is equal to:
//		ceiling(# of tuple slots / 8)
func NewHeapPage(id *HeapPageID, data []byte, td *TupleDesc) (*HeapPage, error) {
	hp := HeapPage{pid: id}
	// todo: more code goes here
	// hp.td = Database.GetCatalog().GetTupleDesc(id.tableID);
	hp.td = td
	hp.numSlots = hp.getNumTuples()
	// data input stream?

	// allocate and read the header slots of this page (TODO)
	return &hp, nil
}

// getNumTuples retruns the number of tuples on this page.
// Specifically, the number of tuples is equal to:
//		floor((BufferPool.PAGE_SIZE*8) / (tuple size * 8 + 1))
func (hp *HeapPage) getNumTuples() int {
	// TODO
	return 0
}

// getHeaderSize computs the number of bytes in the header of a page in a
// HeapFile with each tuple occupying tupleSize bytes
func (hp *HeapPage) getHeaderSize() int {
	// TODO
	return 0
}

// GetID returns the PageID associated with this page
func (hp *HeapPage) GetID() *HeapPageID {
	return hp.pid
}

// todo: getBeforeImage

// HeapPageID is a unique identifier for HeapPage structs.
type HeapPageID struct {
	tableID int
	pageNo  int
}

func NewHeapPageId(tableID int, pageNo int) *HeapPageID {
	hpi := HeapPageID{tableID: tableID, pageNo: pageNo}
	return &hpi
}

// TODO: might want to figure out how to return a hash code for this page,
// represented by the concatenation of the table number and the page number,
// which is needed if a PageID is used as a key in a hash table in the
// BufferPool or whatever

// might also need a comparison operator?

// Serialize returns a representation of this object as an array of integers,
// for writing to disk. Size of the returned array must contain the number of
// integers that corresponds to the number of args to one of the constructors.
// TODO ????
func (hpi *HeapPageID) Serialize() []int {
	data := make([]int, 2)
	data[0] = hpi.tableID
	data[1] = hpi.pageNo
	return data
}
