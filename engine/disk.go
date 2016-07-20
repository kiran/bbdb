package engine

// import (
// 	"bytes"
// 	"fmt"
// 	"strings"
// 	"time"
// )

// interface DbFile
// each table is represented by a single DbFile
// DbFiles can fetch pages and iterate through tuples
// each file has a unique ID used to store metadata about the table in the catalog
// DbFiles are accessed through the buffer pool page cache, rather than directly

// the db file implements a few methods:
// writePage(page p): write page to disk
// readPage(page id): reads from disk
// getTupleDesc
// getId
// deleteTuple
// insert tuple(transaction id, tuple)

// the iterator is a really common pattern for implementing query executors
// all operators in a query plan are implemented as subclasses of the iterator
// interface
