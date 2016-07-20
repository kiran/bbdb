package engine

// DbFile is the interface for database files on disk. Each table is represented by
// a single DbFile. DbFiles can fetch pages and iterate through tuples. Each file
// has a unique id used to store metadata about the table in the Catalog.
// DbFiles are generally accessed through the buffer pool, rather than directly
// by operators.
type DbFile interface {
	// read the specified page from disk.
	// returns an error if the page does not exist in this file.
	ReadPage(pageid int) (Page, error)
	// writes the specified page to disk
	// returns an error if the write fails
	WritePage(page Page) error

	// returns a unique ID used to identify this DbFile in the Catalog. This id
	// can be used to look up the table.
	// TODO: generate this tableid somewhere, and ensure that each HeapFile has
	// a unique id, and that you always return the same value for a particular
	// heapfile. A simple implementation is to use the hash code of the
	// absolute path of the file underlying the HeapFile.
	GetId() int
}
