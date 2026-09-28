---
description: "Fundamentals of the Go programming language."
title: "Numbers and collections"
keywords:
  - Software engineering
  - Go
---

## Numbers

Conversion between numbers can be done using a global built-in function with the type name.

```go
myId := 10

myTicketPrice := 55.7



myIdAsFloat := float64(myId)
myTicketPriceAsInt := int(myTicketPrice)
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
</br>
</br>

> Not every conversion is possible, e.g. a `string` into a `bool`.

</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>

## Strings

Multiline strings can be created using backticks.

```go
strSingle := "Just a single line of text."

strMulti := `We are
on multiple
lines here!`
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
</br>
</br>

## Collections

What do we have?

</br>

### **Arrays**

Fixed length

```go
[5]int
```

### **Slices**

Similar to dynamic length arrays, but are in practice chunks of arrays

```go
[]int
```

### **Maps**

Key-value dictionaries

```go
map[keyType]valueType

// Example
map[string]int
```

### **Generics**

Since Go version 1.18

</br>
</br>
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

- Declare a new array with a length of `20` and type `string`.
- Declare a new slice with the type `int`.
- Declare a new map with a key of type `int` and a value of type `string`.
- Discuss the special function `init`. This is not the same as **main()** ! Not a constructor.
- The `len()` function returns the length of a collection.

</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>

> Collections are not objects. We use global functions to work with collections, such as `len()` and `cap()`.

</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>

## Arrays and slices

Remember, arrays are fixed length and slices are dynamic length.

</br>

### Arrays

```go
var countries [180]string
countries[0] = "Denmark"

// Initialize with values, shorthand notation
accounts := [5]int{1, 2, 3, 4, 5}
```

</br>

### Slices

**Declaring a string slice.**

We use append to add values and note that the new slice is assigned back to the variable.

```go
var countries []string
countries = append(countries, "Denmark")
```

**Initialize with values.**

Overrides any capacity defined in the declaration.

```go
var countries = []string{"Denmark", "Sweden", "Norway"}
println(len(countries))
```

**Initialize with a length of 5.**

The slice already holds 5 empty strings, so an `append` adds a 6th value.

```go
var countries = make([]string, 5)
```

**Initialize with a length of 0 and a capacity of 5.**

Capacity is the space reserved up front. Go only allocates more memory when an `append` goes beyond it.

```go
var countries = make([]string, 0, 5)
println(len(countries), cap(countries)) // 0 5
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
</br>
</br>

## Maps

Pay attention to the type definition.

- For maps the default value is `nil`—initialization is required.

</br>

```go
var platformUsers = map[string]int{"Facebook": 1000000, "Twitter": 500000}
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
</br>
</br>

The `make()` function can also be used to initialize a map and this is most commonly what is done as maps are rarely small in size.

```go
var platformUsers = make(map[string]int)

platformUsers["Facebook"] = 1000000
platformUsers["Twitter"] = 500000
```

</br>
</br>
