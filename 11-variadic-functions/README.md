# Variadic Functions

## Definition
- A variadic function is a function that accepts a variable number of arguments. If the last parameter of a function definition is prefixed by ellipsis ..., then the function can accept any number (including zero) of arguments for that parameter.
- __Only the last parameter of a function can be variadic.__

## Syntax
```
func hello(a int, b ...int) {  
}
```

## Notes:
- The way variadic functions work is by converting the variable number of arguments to a slice of the type of the variadic parameter.

- The advantages of using variadic arguments instead of slices:
    - No need to createe a slice during each function call.
    - In some case, we dont have any variadic arguments. So this is totally not needed to create a slice.
    - More readable

- We can pass a slice to a variadic function. We have to suffix the slice with ```...```. So the slice is directly passed to the function without a new slice being created.

- Just be sure you know what you are doing when you are modifying a slice inside a variadic function.
    - the length and capacity of a slice is not modified after function call 
