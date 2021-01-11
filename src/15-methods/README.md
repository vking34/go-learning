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
- The difference between value and pointer receiver is, changes made inside a method with a pointer receiver is visible to the caller whereas this is not the case in value receiver. 