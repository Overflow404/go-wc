Word count
==============

*An implementation of the wc cli tool written in go*

### How to run the tests

Move inside the root directory of the project and run `go test` from the terminal

---

### How to use the tool

Move inside the root directory of the project and run `go run go-wc [flag] [filename]` from the terminal

### Arguments

#### flag(optional):

* `-c` count bytes
* `-l` count lines
* `-w` count words
* `-m` count unicode characters
* if not specified will print the equivalent of `-c` `-l` `-w` by default

#### filename(mandatory):

* the name of the file to process, it can be a relative path or an absolute one

---

### Optimisations
* Do not open and close the file between each counting operation
