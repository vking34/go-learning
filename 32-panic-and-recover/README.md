# Panic and Recover

## Panic

### What is panic?
- The idiomatic way of handling abnormal conditions in a Go program is using errors. Errors are sufficient for most of the abnormal conditions arising in the program.

- But there are some situations where the program cannot continue execution after an abnormal condition. In this case, we use ```panic``` to prematurely terminate the program.

- When a function encounters a panic, its execution is stopped, __any deferred functions are executed and then the control returns to its caller__. This process continues until all the functions of the current goroutine have returned at which point the program prints the panic message, followed by the stack trace and then terminates.

- It is possible to regain control of a panicking program using ```recover```.

### When should panic be used?

- One important factor is that you should __avoid panic and recover__ and __use errors where ever possible__. Only in cases where the program just cannot continue execution should panic and recover mechanism be used.

- There are two valid use cases for panic:
    - __An unrecoverable error where the program cannot simply continue its execution__. One example is a web server that fails to bind to the required port. In this case, it's reasonable to panic as there is nothing else to do if the port binding itself fails.

    - __A programmer error__. Let's say we have a method that accepts a pointer as a parameter and someone calls this method using a nil argument. In this case, we can panic as it's a programmer error to call a method with nil argument which was expecting a valid pointer.

## Recover
- recover is a builtin function that is __used to regain control of a panicking program.__

- Recover is useful only when __called inside deferred functions__. Executing a __call to recover inside a deferred function stops the panicking sequence by restoring normal execution and retrieves the error message passed to the panic function call__. If __recover is called outside the deferred function__, it will __not stop a panicking sequence__.

-  Getting Stack trace after recover:
    - If we recover from a panic, we lose the stack trace about the panic. Even in the program above after recovery, we lost the stack trace.
    - There is a way to print the stack trace using the ```PrintStack``` function of the ```debug``` package

- Recover __works only__ when it is called from the __same goroutine__ which is panicking. It's __not possible to recover from a panic that has happened in a different goroutine__. 