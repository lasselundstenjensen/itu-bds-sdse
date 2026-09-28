---
description: "Fundamentals of the Go programming language."
title: "Channels"
keywords:
  - Software engineering
  - Go
---

## Channels

A common challenge when using threads is to coordinate the communication between them. Go provides a mechanism to solve this problem: **channels**.

</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>

* **goroutines** can communicate with each other using *channels*, which is a special type of variable.
* A **channel** can contain a *value* of *any kind*.
* A **goroutine** can send a value to a **channel** and another **goroutine** can receive it.
* Channels can be *buffered* or *unbuffered*.

</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>

## Using channels

### Unbuffered

```go
var ch chan string

ch := make(chan string)

ch <- "hello"

message := <-ch
```

</br>
</br>

### Example

```go
func main() {
    ch := make(chan int)
    go func() {
        ch <- 42
    }()
    fmt.Println(<-ch)
}
```

</br>

> The `func()` here is an example of an **anonymous function**, or a **lambda**.

</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>

### Buffered

We can also have buffered channels.

In this case we specify the number of values that can be stored in the channel.

```go
ch := make(chan string, 2)
ch <- "hello"
ch <- "world"

fmt.Println(<-ch)
fmt.Println(<-ch)
```

> If the buffer is empty, receiving waits until a message is sent. If the buffer is full, sending waits until a message is received.

</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>

## Things to be aware of

* **Deadlocks** occur when all goroutines are waiting, e.g. sending to an unbuffered channel that nothing receives from.
* Close a channel when no more values will be sent. Otherwise a receiver looping over it with `range` waits forever.

```go
close(channel)
```

</br>

> **Good practice:** Use the `defer` keyword to close the channel and to keep code in related blocks.

</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>

## Waitgroups

A **waitgroup** is a way to wait for a collection of goroutines to finish.

```go
var wg sync.WaitGroup

wg.Add(1)
go func() {
    defer wg.Done()
    fmt.Println("Hello")
}()
wg.Wait()
```

</br>
</br>
