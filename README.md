# bbdb

A baby proof-of-concept relational database.

This is based off of the interfaces in [MIT's databases course](https://github.com/MIT-DB-Class/course-info).

bbdb is a disk-oriented relational database, with very minimal functionality. Tuples are stored on disk in heap files, and only fixed-length string and integer fields are allowed.

## Structure

- access methods (just heap files for now) store relations on disk, and provide a way to iterate through tuples of these relations
- a buffer pool caches active pages in memory and handles concurrency/transactions (to come!)
- tuples are described by tuple schemas, and consist of fields (only fixed-length string and integer fields)
- tuples are processed via a collection of operators (join, select, filter, insert, etc.)

Similar to the MIT SimpleDB, bbdb does not have many things that are typical of a DBMS, such as:

- Indexes
- a query optimizer
- views
- a SQL parser (for now). Like SimpleDB, bbdb's queries are built up by hand, by chaining together sets of tuple-at-a-time operators.