# Write files

## Create the file
```
os.Create("filename")
```

## Write to file

- Write string:
    ```
    fd.WriteString()
    ```

- Write bytes:
    ```
    fd.Write()
    ```

- Write strings line by line:
    ```
    fmt.Fprintln(fd, "string line")
    ```

## Writing to file concurrently
- When multiple goroutines write to a file concurrently, we will end up with a race condition. Hence concurrent writes to a file should be co-ordinated using a channel.

- Problem: We will write a program that creates 100 goroutines. Each of this goroutine will generate a random number concurrently, thus generating hundred random numbers in total. These random numbers will be written to a file. We will solve this problem by using the following approach.

- Approach:
    - Create a channel which will be used to read and write the generated random numbers.

    - Create 100 producer goroutines. Each goroutine will generate a random number and will also write the random number to a channel.

    - Create a consumer goroutine which will read from the channel and write the generated random number to the file. Thus we have only one goroutine writing to a file concurrently thereby avoiding race condition
    