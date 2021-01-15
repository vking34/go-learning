# Interfaces

## Definition
- An interface is a set of method signatures. When a type provides definition for all the methods in the interface, it is said to implement the interface.

## Notes
- This is quite different from other languages like Java where a class has to explicitly state that it implements an interface using the implements keyword. This is not needed in Go and Go interfaces are implemented implicitly if a type contains all the methods declared in the interface.

- An interface can be thought of as being represented internally by tuple (type, value).

- An interface that has zero methods is called an empty interface. It represented as ```interface{}```. Since the empty interface has zero methods, all types implement the empty interface.

- Type assertion:
    - Used to extract the underlying value of the interface
    - Syntax: ```i.(T)```
        - ```i```: interface
        - ```T```: type
    
    - If we assert the wrong type, the program will panic.
    - To solve this problem: ```v, ok := i.(T)```
        - if ```ok``` is false, ```v``` is the zero value of ```T``` 

- Type switch:
    - Used to compare the concrete type of an interface against multiple types specified in various case statements. It is similar to switch case. The only difference is the case specify types and not value as in normal switch