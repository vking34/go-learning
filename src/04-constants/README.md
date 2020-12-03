# Constants

## Notes
- __The value of a constant should be known at compile time__. Hence it cannot be assigned to a value returned by a function call since the function call takes place at run time.
    ```
    var a = math.Sqrt(4)   //allowed
    const b = math.Sqrt(4) //not allowed
    ```

- Any value enclosed between double quotes is a string constant in Go. They are __untyped__

- Untyped constants have __a default type__ associated with them and they supply it if and only if a line of code demands it.

- Go's strong typing policy disallows variables of one type to be assigned to another. Even assign one type to its alias.
    ```
    var defaultName = "Sam" //allowed
    type myString string
    var customName myString = "Sam" //allowed
    customName = defaultName //not allowed
    ```

