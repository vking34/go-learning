# If Else statement

```
if condition {  
}
```

- Unlike in other languages like C, the braces ```{  }``` are mandatory even if there is only one line of code between the braces ```{  }```

- If with assignment:
    ```
    if assignment-statement; condition {  
    }
    ```

    ```
    if num := 10; num % 2 == 0 { //checks if number is even
        fmt.Println(num,"is even") 
    } else {
        fmt.Println(num,"is odd")
    }
    ```

- ```else``` statement should start in the same line after the closing curly brace ```}```. If not the compiler will complain because of the way Go inserts semicolons automatically.
    ```
    } else {
    ```

- In Go's philosophy, it is better to avoid unnecessary branches and indentation of code. It is also considered better to return as early as possible.

    ```
    num := 10;
    if num%2 == 0 { //checks if number is even
        fmt.Println(num, "is even")
        return
    }
    ```