# Channels

## Definition
- Channels can be thought as pipes using which Goroutines communicate. 
- Similar to how water flows from one end to another in a pipe, data can be sent from one end and received from the another end using channels.

## Declaration
- ```chan T``` is a channel of type ```T```.
- Each channel has a type associated with it. This type is the type of data that the channel is allowed to transport. No other type is allowed to be transported using the channel.
- The zero value of channel is ```nil```. ```nil``` channels are not of any use and hence the channel has to be defined using ```make``` similar to maps and slices.

## Sending and receiving from channel
```
data := <- a // read from channel a  
a <- data // write to channel a 
```

- The direction of the arrow with respect to the channel specifies whether the data is sent or received.

- Sends and receives are blocking by default. When a data is sent to a channel, the control is blocked in the send statement until some other Goroutine reads from that channel. Similarly when data is read from a channel, the read is blocked until some Goroutine writes data to that channel. This property of channels is what helps Goroutines communicate effectively without the use of explicit locks or conditional variables that are quite common in other programming languages.

## Deadlock
- One important factor to consider while using channels is deadlock. If a Goroutine is sending data on a channel, then it is expected that some other Goroutine should be receiving the data. If this does not happen, then __the program will panic at runtime with ```Deadlock```__. Similarly if a Goroutine is waiting to receive data from a channel, then some other Goroutine is expected to write data on that channel, else the program will panic.

## Unidirectional Channel
- All the above channels we discussed so far are bidirectional channels, that is data can be sent and received on them.
- __Unidirectional channels__ is channels that __only send and receive data__.
- It is possible to __convert a bidirectional channel to a send only or receive only channel__ but not the vice versa.
- The point of writing to a send only channel if it cannot be read from or the vice versa!

## Closing channels
- Senders have the ability to close the channel to notify receivers that no more data will be sent on the channel.
- Receivers can use an additional variable while receiving data from the channel to check whether the channel has been closed.
    ```
    v, isOpen := <- ch  
    ```
- The __value__ read from a __closed channel__ will be the __zero value of the channel's type__.