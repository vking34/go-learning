# Mutex

## Critical Section
- When a program runs concurrently, the parts of code which modify shared resources should not be accessed by multiple Goroutines at the same time. This section of code which modifies shared resources is called critical section. 

- Depending on how content switching happens, undesirable situation where the output of the program depends on the sequence of execution of goroutines is called race condition.

## Mutex
- A mutex is used provide a locking mechanism to ensure that only one goroutine is running the critical section of code at any point of time to prevent race condition from happening.

- Mutex is available in the sync package. There are two methods defined on Mutex namely ```Lock``` and ```Unlock```. Any code that is present between a call to ```Lock``` and ```Unlock``` will be executed by only one Goroutine, thus avoiding race condition.

- If one Goroutine already holds the lock and if a new Goroutine is trying to acquire a lock, the new Goroutine will be blocked until the mutex is unlocked.

- Mutex is a __struct type__ and we create a zero valued variable m of type Mutex.

## Solving the race condition using channel
- Create a buffered channel with one capacity. This buffered channel is used to ensure that only one Goroutine access the critical section of code which increments x. 
- Write data to the channel right before critical section. All other Goroutines trying to write to this channel are blocked until the value is read from this channel.

## Mutex vs Channels
- How do we decide what to use when. __The answer lies in the problem__ you are trying to solve. If the problem you are trying to solve is a better fit for mutexes then go ahead and use mutex. Do not hesitate to use mutex if needed. If the problem seems to be a better fit for channels, then use it.

- Most Go newbies try to solve every concurrency problem using a channel as it is a cool feature of the language. This is wrong. The language gives us the option to either use Mutex or Channel and there is no wrong in choosing either.

- In general __use channels when Goroutines need to communicate__ with each other and __mutexes when only one Goroutine should access the critical section of code.__

- My advice would be to __choose the tool for the problem and do not try to fit the problem for the tool__ :)
