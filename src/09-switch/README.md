# Switch

- Syntax:
    ```
    switch finger {
    case 1:
        fmt.Println("Thumb")
    case 2:
        fmt.Println("Index")
    
    default:
        fmt.Println("incorrect finger number")
    ```

## Notes:
- Duplicate cases are not allowed
- Multiple expressions in a case
- Expressionless
    ```
    num := 75
    switch { // expression is omitted
    case num >= 0 && num <= 50:
        fmt.Printf("%d is greater than 0 and less than 50", num)
    case num >= 51 && num <= 100:
        fmt.Printf("%d is greater than 51 and less than 100", num)
    }
    ```

- In Go, the control comes out of the switch statement immediately after a case is executed. A ```fallthrough``` statement is used to transfer control to the first statement of the case that is present immediately after the case which has been executed.