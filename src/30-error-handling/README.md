# Error Handling

## Error
- Errors indicate an abnormal condition occurring in the program. Let's say we are trying to open a file and the file does not exist in the file system. This is an abnormal condition and it's represented as an error.

- Errors in Go are plain old values. Errors are represented using the built-in ```error``` type.



## Notes:
- If a function or method returns an error, then by convention it has to be the last value returned from the function.