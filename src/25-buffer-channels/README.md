# Buffered Channels

- All the channels we discussed in the previous chapter are buffered channels.

- Sends to a buffered channel are __blocked only when the buffer is full__. Similarly receives from a buffered channel are __blocked only when the buffer is empty__.

- Create a buffered channel:
    ```
    ch := make(chan type, capacity) 
    ```

    - __capacity__ is __the times of sends (receives)__ to the channel. the value of capacity is 0 by default (for unbuffered channel) but should be greater than 0 (for buffered channel)
