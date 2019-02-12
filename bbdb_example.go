package main

import (
	"fmt"
	"github.com/kiran/bbdb/engine"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// construct a 3-column table schema
	types := []engine.Type{new(engine.IntType), new(engine.IntType), new(engine.IntType)}
	names := []string{"field0", "field1", "field2"}
	descriptor := engine.NewTupleDesc(types, names)

	fmt.Println(descriptor)

	// create the table, associate it with some_data_file.dat
	f, err := os.Create("some_data_file.dat")
	check(err)
	defer f.Close()
	table1 := engine.NewHeapFile(f, descriptor)
	// Database.Catalog().AddTable(table1, "test_table")

	fmt.Println(table1)

	engine.ConvertToHeapFileTest()

	// Construct the query: we use a simple SeqScan, which spoonfeeds tuples
	// via its iterator
	// tid := engine.NewTransactionID()
	// seqScan := engine.NewSeqScan(tid, table1.getID())

	// and run it
	// for (seqScan.Open(); seqScan.hasNext(); tup := seqScan.Next()) {
	// 	fmt.Println(tup)
	// }
	// seqScan.Close()
	// Database.BufferPool().TransactionComplete(tid)
	// catch errors somewhere
}
