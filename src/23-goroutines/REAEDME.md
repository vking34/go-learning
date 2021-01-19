# Goroutines

## Definition
- Goroutines are functions or methods that run concurrently with other functions or methods.
- Gorountines can be thought of as light weight threads. The cost of creating a goroutine is tiny when compared to a thread. Hence it's common for Go applications to have thousends of goroutines running concureently.

## Advantages of goroutines over threads
- Goroutines are __extremely cheap__ when compared to threads. They are only a few kb in stack size and the stack can grow and shrink according to needs of the application whereas in the case of threads the stack size has to be specified and is fixed.

- The Goroutines are __multiplexed__ to fewer number of OS threads. There might be only __one thread in a program with thousands of Goroutines__. If any Goroutine in that __thread blocks__ say waiting for user input, then __another OS thread is created__ and the __remaining Goroutines are moved to the new OS thread__. All these are taken care by the runtime and we as programmers are abstracted from these intricate details and are given a clean API to work with concurrency.

- Goroutines __communicate using channels__. Channels by design prevent race conditions from happening when accessing shared memory using Goroutines. Channels can be thought of as a pipe using which Goroutines communicate.