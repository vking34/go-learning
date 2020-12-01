# Types:

## Basic types:
- bool
- Numberic types:
    - int8, int16, int32, int64, int
    - uint8, uint16, uint32, uint64, uint
    - float32, float64
    - complex64, complex128
    - byte
    - rune
- string:
    

## Notes:

- ```int``` represents 32 or 64 bit integer depending on the underlying platform. 32 bits in 32-bit systems and 64 bits in 64-bit systems.

- ```string```: Strings are a collection of bytes in Go
- Strings can be concatenated using the + operator

- Go is very strict about explicit typing. There is no automatic type promotion or conversion. When you try to assign a variable to another variable of other type without any type conversion, the compiler will throw an error.

- ```unsafe``` package should be used with care as the code using it might have portability issues.


