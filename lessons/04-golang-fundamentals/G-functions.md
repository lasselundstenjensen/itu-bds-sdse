---
description: "Fundamentals of the Go programming language."
title: "Functions"
keywords:
  - Software engineering
  - Go
---

## Functions

Go functions are similar to functions in other programming languages.

- A function can take zero or more arguments.
- Arguments cannot have default values.
- The last argument can be variadic (variable lenght).

</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>

### **Some things to be aware of**

- Functions can return more than one value.
- Functions can return labeled variables.
- Functions always receive arguments by value.

</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>

**Syntax and use:**

```go
func load() {}
func load(file string) {}

func add(a int, b int) int {
    return a + b
}

func subtract(a, b int) int {
    return a - b
}

func addAndSubtract(a int, b int) (int, int) {
    return a + b, a - b
}
```

Labelled return values, but try to limit the use of this as it obscures where return values are set as functions grow tall.

```go
func multiply(a, b int) (result int) {
    result = a * b
    return
}
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

## Using function return values

Escape the return values you don't need with `_`.

```go
func main() {
    sum, _ := addAndSubtract(1, 2)
    fmt.Println(sum)
}
```

</br>
</br>

> Why would you want to escape a return value?

</br>
</br>
