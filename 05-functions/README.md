# Functions

## Syntax
```
func functionName(param type) returnType {

}
```

- Parameters and return type are optional.
- If consecutive parameters are of the same type, we can avoid writing the type each time and it is enough to be written once at the end.
    ```
    func calculateBill(price, no int) int {
        
    }
    ```

- Multiple return values:
    ```
    func rectProps(length, width float64)(float64, float64) {  
        var area = length * width
        var perimeter = (length + width) * 2
        return area, perimeter
    }
    ```

- Named return values from a function. If a return value is named, it can be considered as being declared as a variable in the first line of the function.
    ```
    func getRectPros(length, width float64) (area, perimeter float64) {

    }
    ```

- Blank identifier: ```_``` can be used in place of any value of any type.
## Notes
- It is possible to return multiple values from a function.
- If a function returns multiple return values then they must be specified between ```( )```. 
- It is possible to return named values from a function.

