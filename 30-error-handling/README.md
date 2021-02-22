# Error handling

## Error
- Errors indicate an abnormal condition occurring in the program. Let's say we are trying to open a file and the file does not exist in the file system. This is an abnormal condition and it's represented as an error.

- Errors in Go are plain old values. Errors are represented using the built-in ```error``` type.

## Error Handling
- The idiomatic way of handling errors in Go is to compare the returned error to ```nil```. A ```nil``` value indicates that no error has occurred and a non-```nil``` value indicates the presence of an error.

- Error type representation:
    - Error is an interface tpye:
        ```
        type error interface {  
            Error() string
        }
        ```

    - It contains a single method with signature Error() string. Any type which implements this interface can be used as an error. This method provides the description of the error.

- Extract more information from the error:
    - Asserting the underlying struct type and getting more information from the struct fields. (__type assertion__)
        ```
        type PathError struct {  
            Op   string
            Path string
            Err  error
        }

        func (e *PathError) Error() string { return e.Op + " " + e.Path + ": " + e.Err.Error() }  
        ```
    
    - Asserting the underlying struct type and getting more information using methods

- __Never ever ignore errors__


## Notes
- If a function or method returns an error, then by convention it has to be the last value returned from the function.