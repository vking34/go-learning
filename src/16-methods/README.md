## Methods

## Definition
A method is just a function with a special receiver type between the ```func``` keyword and the method name. The receiver type can either be a struct type or non-struct type.

```
func (t Type) methodName(parameter list) {  
}
```

## Methods vs Funtions
- Go is not a pure object-oriented programming language and it does not support classes. Hence methods on types are a way to achieve behavior similar to classes.

- Methods with the same name can be defined on different types whereas functions with the same names are not allowed.

- Methods are used to implement interfaces

- The difference between value receiver and pointer receiver is, changes made inside a method with a pointer receiver is visible to the caller whereas this is not the case in value receiver.
    - A value receiver in a method will need the entire struct to be copied which will be expensive if the struct has many fields.

- A method has a value receiver, it will accept both pointer and value receivers.
    - Even pointer call that method, the argument passed to that method is still value receiver.

- Similarly, a method has a pointer receiver, it will accept both pointer and value receiver.

- To define a method on a type, the definition of the receiver type and the definition of the method should be present in the same package.

- Can not define methods to built-in types. To do that, we create a alias type.