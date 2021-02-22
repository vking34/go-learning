# Object Oriented Programming

- Go is not a pure object oriented programming. There is no hierarchy.

- Structs instead of Classes

- ```New()``` function instead of constructor:
    - Go does not support constructors. If the zero value of a type is usable, it is the job of the programmer to unexport the type to prevent access from other packages and also to provide a function named NewT(parameters) which initializes the type T with the required values.
    - It is a convention in Go to name function that creates a value of type T to ```NewT(params)```. It acts as constructor. 
    - If the package defines only one type, then it's a convention in Go to name this function just ```New(parameters)``` instead of ```NewT(parameters)```.