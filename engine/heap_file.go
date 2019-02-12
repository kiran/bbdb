package engine

import (
	"crypto/md5"
	"encoding/hex"
	"os"
)

// HeapFile is an implementation of a DbFile that stores a collection of tuples
// in no particular order. Tuples are stored on pages, each of which is a fixed
// size, and the file is simply a collection of those pages. HeapFile works
// closely with HeapPage. The format of HeapPages is described in the HeapPage.
type HeapFile struct {
	//??
	file *os.File
	td   *TupleDesc
}

// NewHeapFile constructs a HeapFile backed by the specified file
func NewHeapFile(f *os.File, td *TupleDesc) *HeapFile {
	// some code goes here
	return new(HeapFile)
}

// GetFile returns the File backing this HeapFile on disk
func (hp *HeapFile) GetFile() {
}

// GetID returns an ID uniquely identifying this HeapFile. Implementation note:
// this table id is generated somewhere to ensure that each HeapFile has a
// unique id, and that you always return the same value for a particular heap
// file. We suggest hashing the absolute file name of the file underlying the
// heap file, ie, f.getAbsoluteFile().hashCode().
func (hp *HeapFile) GetID() string {
	// TODO: this is probably not the absolute file path
	filename := []byte(hp.file.Name())
	hash := md5.Sum(filename)
	return hex.EncodeToString(hash[:])
}

// GetTupleDesc returns the TupleDesc of the table stored in this DbFile
func (hp *HeapFile) GetTupleDesc() *TupleDesc {
	return hp.td
}

// ReadPage is TODO
func (hp *HeapFile) ReadPage(pid PageID) *Page {
	// todo
	return new(Page)
}

// WritePage is TODO
func (hp *HeapFile) WritePage(page Page) error {
	// todo
	return nil
}

// NumPages returns the number of pages in this HeapFile
func (hp *HeapFile) NumPages() int {
	// todo
	return 0
}

// todo: addTuple
// todo: deleteTuple
// todo: DbFileIterator
