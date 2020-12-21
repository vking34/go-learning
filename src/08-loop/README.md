# Loop

- __Go does not have ```while``` or ```do while```__

- for loop syntax
    ```
    for initialisation; condition; post {  
    }
    ```

- label:
    ```
    func main() {  
    outer:  
        for i := 0; i < 3; i++ {
            for j := 1; j < 4; j++ {
                fmt.Printf("i = %d , j = %d\n", i, j)
                if i == j {
                    break outer
                }
            }

        }
    }
    ```

- All the three components of the for loop namely initialisation, condition and post are optional. The semicolons in the for loop of the above program can also be omitted.

- More examples:
    ```
    for no, i := 10, 1; i <= 10 && no <= 19; i, no = i+1, no+1 { //multiple initialisation and increment
        fmt.Printf("%d * %d = %d\n", no, i, no*i)
    }
    ```

    ```
    // infinite loop
    for {  
    }
    ```