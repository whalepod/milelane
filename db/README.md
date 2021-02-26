# DB
## Question
### Why type column is int in devices table?
Currently, type column in devices table is int type.
Because we want to change the value to be saved depending on the type of device to be received.
In the future, we plan to implement an int type that is saved depending on the received character string like below.  
https://github.com/whalepod/milelane/blob/18030f9b32d2aef949bf4fa1709ce2a81bd43022/app/domain/device.go#L85
