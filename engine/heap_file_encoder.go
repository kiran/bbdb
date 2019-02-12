package engine

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// HeapFileEncoder reads a comma delimited text file or accepts an array of
// tuples and converts it to pages of binary data in the appropriate format for
// simpledb heap pages.
// Pages are padded out to a specified length, and written consecutive in a
// data file.

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ConvertToTempFile(tuples [][]int, numFields int) error {
	f, err := os.Create("./tempTable.txt")
	check(err)
	w := bufio.NewWriter(f)
	defer f.Close()
	defer w.Flush()

	for _, tup := range tuples {
		writtenFields := 0
		for _, field := range tup {
			writtenFields++
			if writtenFields > numFields {
				// TODO add which field
				return errors.New("Tuple has more than " + string(numFields) + " fields.")
			}
			w.WriteString(strconv.Itoa(field))
			if writtenFields < numFields {
				w.WriteRune(',')
			}
		}
		w.WriteRune('\n')
	}
	return nil
}

func calculateRecdBits(numFields, npagebytes int) (nrecbytes, nrecords, nheaderbytes int) {
	// assumes 32-bit integers
	nrecbytes = 4 * numFields
	// each tuple requires one header byte + the number of bytes needed to repr it
	nrecords = (npagebytes * 8) / (nrecbytes*8 + 1)
	// per record, we need one bit; there are nrecords per page, so we need
	// nrecords bits, ie, ((nrecords/32)+1) integers
	nheaderbytes = nrecords / 8
	if nheaderbytes*8 < nrecords {
		nheaderbytes++
	}
	return

}

// ConvertToHeapFole converts the specified input text file into a binary page
// file.
// Assume format of the input file is (note that only integer fields are supported)
// int, ... , int<br>
// int, ... , int<br>
// where each row represents a tuple.
// The format of the output file will be as specified in HeapPage and HeapFile.
// npagebytes is the number of bytes per page in the output file.
func ConvertToHeapFile(outFilePath string, npagebytes int, numFields int) error {
	nrecbytes, nrecords, nheaderbytes := calculateRecdBits(numFields, npagebytes)

	// set up the readers
	in_f, err := os.Open("./tempTable.txt")
	check(err)
	scanner := bufio.NewScanner(in_f)

	outf, err := os.Create(outFilePath)
	check(err)
	defer outf.Close()

	// parse out lines.
	// 2 writers:
	// header stream
	// page stream

	headerStream := bytes.NewBuffer(make([]byte, 0, nheaderbytes))
	pageStream := bytes.NewBuffer(make([]byte, 0, npagebytes))
	hasMore := true
	recordCount := 0
	npages := 0

	// portion out tuples into pages, buffer up the data
	// and write to stream once filled.
	// repeat until we've gotten through all the tuples
	for hasMore {
		// maybe refactor this into a read
		hasMore = scanner.Scan()
		line := scanner.Text()
		check(scanner.Err())
		if hasMore {
			// split the line on commas
			fields := strings.Split(line, ",")
			fmt.Println(fields)

			// parse out the integer fields
			if len(fields) > numFields {
				fmt.Print("BAD LINE -- too many fields")
			}
			// write out the fields to the buffer
			for _, field := range fields {
				intfield, atoi_err := strconv.ParseInt(field, 10, 32)
				check(atoi_err)
				err := binary.Write(pageStream, binary.LittleEndian, intfield)
				check(err)
			}
			recordCount += numFields
		}

		// if we wrote a full page of records, or if we're done altogether,
		// write out the header of the page.
		if recordCount >= nrecords || !hasMore && recordCount > 0 || !hasMore && npages == 0 {
			fmt.Println("flushing to disk")
			// write out the bitmap for the header
			// in the header, write a 1 for bits that correspond to records we've
			// written and a 0 for empty slots.
			var headerbyte byte

			for i := 0; i < nheaderbytes*8; i++ {
				if i < recordCount {
					// shift in a headerbyte
					headerbyte |= (1 << uint(i%8))
				}
				if ((i + 1) % 8) == 0 {
					headerStream.WriteByte(headerbyte)
					headerbyte = 0
				}
			}
			// flush out the rest of the header
			headerStream.WriteByte(headerbyte)

			// pad out the rest of the page with zeros
			for i := 0; i < npagebytes-(recordCount*nrecbytes+nheaderbytes); i++ {
				pageStream.WriteByte(0)
			}

			// write header and body to file
			// flush, write to files
			_, err := headerStream.WriteTo(outf)
			check(err)
			_, err = pageStream.WriteTo(outf)
			check(err)

			// reset header and body for next page
			// reset record count, inc numpages
			headerStream = bytes.NewBuffer(make([]byte, 0, nheaderbytes))
			pageStream = bytes.NewBuffer(make([]byte, 0, npagebytes))
			recordCount = 0
			npages++
		}
	}

	return nil
}

func ConvertToHeapFileTest() {
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := 4
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	ConvertToTempFile(twoD, 4)
	ConvertToHeapFile("some_data_file.dat", 4096, 4)
}
