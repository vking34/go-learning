# Worker Pools

## WaitGroup
- A WaitGroup is used to wait for a collection of Goroutines to finish executing. The control is blocked until all Goroutines finish executing.

### Notes
- To understand worker pools, we need to to first know about ```WaitGroup``` as it will be used in the implementation of Worker pool.

- The way WaitGroup works is by using a counter. When we call Add on the WaitGroup and pass it an int, the WaitGroup's counter is incremented by the value passed to Add. The way to decrement the counter is by calling Done() method on the WaitGroup. The Wait() method blocks the main Goroutine to wait until the counter becomes zero.

- __It is important to pass the address of wg in line no. 21. If the address is not passed, then each Goroutine will have its own copy of the WaitGroup and main will not be notified when they finish executing.__

## Worker Pool Implementation
- One of the important uses of buffered channel is the implementation of worker pool.
- In general, a worker pool is a collection of threads which are waiting for tasks to be assigned to them. Once they finish the task assigned, they make themselves available again for the next task.
