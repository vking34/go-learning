# Custom Error

## Creating custom errors using the New function
- The implementation of the New function in the errors package:
    ```
    // Package errors implements functions to manipulate errors.
    package errors

    // New returns an error that formats as the given text.
    func New(text string) error {
        return &errorString{text}
    }

    // errorString is a trivial implementation of error.
    type errorString struct {
        s string
    }

    func (e *errorString) Error() string {
        return e.s
    }
    ```

## Adding more information to the error using Errorf
- The ```Errorf``` function of the ```fmt``` package.

## Providing more information about the error using struct type and fields
- It is also possible to use struct types which implement the error interface as errors. This gives us more flexibility with error handling.
- Creating a struct type that implements the error interface and using its fields to provide more information about the error will not break our code.

- The naming convention for error types is that the name should __end with the text ```Error```__

## Providing more information about error using methods on struct types.