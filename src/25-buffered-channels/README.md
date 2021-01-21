# Buffered Channels

## Notes
- The capacity of the buffered channel is 2 and hence the write Goroutine will be able to write value 0 and 1 to the ch channel immediately and then it blocks until at least one value is read from ch channel.
- Deadlock: if we try to write to a full channel but no goroutine read from the channel, there will be a deadlock and the program will panic at runtime.
