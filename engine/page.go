package engine

// Page is the interface used to represent files that are resident in the
// BufferPool. Typically, DbFiles will read and write pages from disk.
// Pages may be "dirty", indicating that they have been modified since they
// were last written out to disk.
type Page interface {
	// GetId() PageId
	// IsDirty() bool
	// MarkDirty(dirty bool, tid TransactionId)

	GetPageData() []byte
}

// PageID is an interface to a specific page of a specific table.
type PageID interface {
	// return a representation of this page id object as a collection of ints
	// (used for logging)
	Serialize() []int
	// returns the unique table id
	GetTableId() int
	Pageno() int
}
