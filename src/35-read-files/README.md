# Read Files

## Reading an entire file into memory

- There are 3 ways to determine the file path:
    - Using absolute file path: not flexible
    - Passing the file path as a command line flag: flexible
    - Bundling the text file along with the binary: better but complex

## Reading a file in small chunks (stream)

- A more optimal way is to read the file in small chunks. This can be done with the help of the ```bufio``` package.


## Run app
```
go run main -f=./test.txt -fs=./test1.txt
```