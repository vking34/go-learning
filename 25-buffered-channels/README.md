# Buffered Channels

- All the channels we discussed in the previous chapter are buffered channels.

- Sends to a buffered channel are __blocked only when the buffer is full__. Similarly receives from a buffered channel are __blocked only when the buffer is empty__.

- Create a buffered channel:
    ```
    ch := make(chan type, capacity) 
    ```

    - __capacity__ is __the times of sends (receives)__ to the channel. the value of capacity is 0 by default (for unbuffered channel) but should be greater than 0 (for buffered channel)


## Notes
- The capacity of the buffered channel is 2 and hence the write Goroutine will be able to write value 0 and 1 to the ch channel immediately and then it blocks until at least one value is read from ch channel.
- Deadlock: if we try to write to a full channel but no goroutine read from the channel, there will be a deadlock and the program will panic at runtime.
- Length vs Capacity:
    - Capacity: the number of values that channel can hold
    - Length: the number of elements currently queued in it
