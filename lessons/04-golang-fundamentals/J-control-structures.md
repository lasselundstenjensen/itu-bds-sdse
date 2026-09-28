---
description: "Fundamentals of the Go programming language."
title: "Control structures"
keywords:
  - Software engineering
  - Go
---

## Control structures

</br>

* **if** - **else**
* **switch**
* **for**

</br>

* No **while** or **do**-**while** loops.

</br>

* No parenthesis around boolean conditions or values.
* Only equality operator is `==`.
* Other operators: `!=`, `<`, `>`, `<=`, `>=`.

</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>

## if - else

* Can have multiple conditions.

```go
if message != nil {
} else {
}

// Variable available in special scope, condition is last
if text := "hello"; message != nil {
} else {
}
```

* There is no ternary operator in Go—e.g. in Python you can do the following, but not in Go:

### **Python**

```python
# python ternary operator
min = "a is minimum" if a < b else "b is minimum"
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

## switch

* No `break` statement needed.
* `fallthrough` keyword is used to go to the next case.

```go
switch day {
    case "Wednesday":
        fmt.Println("Time for class! 📚")
    case "Friday":
        fmt.Println("Woooh! Parteeey! 🎉")
    case "Saturday":
        fallthrough
    case "Sunday":
        fmt.Println("Weekend 🌞")
    default:
        fmt.Println("It's another day 😴")
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

## switch (with conditions)

* Great for replacing large if-else chains.

```go
switch {
    case message == nil:
    // ...
    case message != nil && message.Text == "":
    // ...
    case message.Sent == true:
    // ...
    default:
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

## for

* Multi-purpose loop.

```go
// Classic loop
for i := 0; i < len(myCollection); i++ {
}

// Loop over a collection
for index := range myCollection {
}

// Loop over a map
for key, value := range myMap {
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

## for (continued)

* We can emulate a `while` loop with boolean expressions.

```go
// Emulating a while loop
quit := false
for !quit {
    // process ... then set 'quit' to true to exit
}

count := 0
for count < 10 {
    count++
}
```

</br>

```go
// Infinite loop
for {
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

## TRY IT OUT

![Rob Pike](../../images/rob-pike.png)

</br>

> Build a simple calculator command line application that takes two numbers and an operator as input and outputs the result.
>
> **HINT:** Use the `fmt.Scanf()` function to read input from the user.
> 
> **HINT:** Use a `switch` statement to determine the operation to perform.

</br>
</br>
