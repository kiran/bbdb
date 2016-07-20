# bbdb

A baby proof-of-concept relational database.

This is based off of the interfaces in [MIT's databases course](https://github.com/MIT-DB-Class/course-info).

bbdb is a disk-oriented relational database, with very minimal functionality. Tuples are stored on disk in heap files, and only fixed-length string and integer fields are allowed.