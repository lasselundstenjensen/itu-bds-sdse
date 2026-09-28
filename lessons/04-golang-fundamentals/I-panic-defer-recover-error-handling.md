---
description: "Fundamentals of the Go programming language."
title: "panic, defer, recover, and error handling"
keywords:
  - Software engineering
  - Go
---

# Special functions

- Package **init** func
- **panic**
- **defer**
- **recover**

</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>

## init

Used to run code that runs before the main function. 

All init functions across all packages are executed.

</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>

## panic

This is an interruption. It will terminate the program, if not handled.

</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>

## defer

It defers the execution of "something" until the end of the function.

Last in, first out.

</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>

It works with `panic()` and `recover()`.

</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>

## recover

It is used to recover from a panic.

</br>

```go
func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from panic:", r)
        }
    }()

    functionThatCanPanic()
}

func functionThatCanPanic() {
    panic("A problem happened!")
}
```

**Output:**

```text
Recovered from panic: A problem happened!
```

</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>

## Error handling

We do not have exceptions in Go.

</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>

This is the typical design pattern when dealing with errors (~kind of pseudocode):

```go
func readFile(file string) (string, error) {
    // attempt to read the file and set the 'ok' boolean
    if ok {
        return data, nil
    } else {
        return "", errorDetails
    }
}

func main() {
    data, err := readFile("file.txt")
}
```

</br>
</br>
