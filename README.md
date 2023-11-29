Word count
==============

*An implementation of the wc cli tool written in go*

### How to run the tests

Move inside the root directory of the project and run `go test` from the terminal

---

### How to use the tool

Move inside the root directory of the project and run `go run go-wc [flag] [filename]` from the terminal

### Arguments

#### flag:

* `-c` count bytes
* `-l` count lines
* if not specified will run with the `-c` by default

#### filename:

* the name of the file to process, it can be a relative path or an absolute one

---

