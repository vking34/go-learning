# First Class Function

## Definition
- __First class functions allows functions to be assigned to variables, passed as arguments to other functions and returned from other functions. Go has support for first class functions.__

## Anonymous functions
- A function having no name is assigned directly to a variable.
    ```
    a := func() {
        fmt.Println("hello world first class function")
    }
    ```

- It is also possible to call a anonymous function without assigning it to a variable.

## User define function types
- Just like we define our own struct type, it is possible to define our own function types.
    ```
    type add func(a int, b int) int  
    ```

## Higher-order function

- Definition: A function which does at least one of the following:
    - takes one or more functions as parameters
    - return a function as its result

## Closures

- Closures are a special case of anonymous functions. Closures are anonymous functions which access the variables defined outside the body of the function.

- Every closure is __bound to its own surrounding variable, not sharing variable__.