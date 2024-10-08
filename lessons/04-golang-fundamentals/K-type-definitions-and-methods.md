---
description: "Fundamentals of the Go programming language."
title: "Type definitions and methods"
keywords:
  - Software engineering
  - Go
---

## Type definitions

Recap: packages can include variables, constants, and functions.

But they can also include types:

* **New types**
* **Aliases**

</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>

```go
// Type alias
type customString = string

// New type based on an existing type
type customString string

// A type has constructor and conversion functions
cs := customString("Hello")
```

> You can add methods to a new type, but not to an alias.

</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>

## Methods

Methods are extensions to types and are defined with a special syntax on functions.

```go
func (cs customString) PrintInChinese() string {
    // ...
}

cs.PrintInChinese()
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

### Exercise: Distance converter

> Create a new type 'distance' from an existing number type.

> Declare a variable of type 'distance'; lets call it 'miles'.

> Create a method for the 'distance' type that converts the distance to kilometers.

> Call the method on the 'miles' variable and print the result.

<!-- You can also think about the attribute argument as: func (this TheType) TheFunction() {} -->

</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>

## Complex types

We got two complex types in Go: **Structs** and **Interfaces**.

### Structures

* They are similar to the idea of a class
* Got a default constructor mechanism
* Can have methods
* A data type with strongly typed properties

### Interfaces

* A collection of method signatures
* Emulates polymorphism from OOP
* Interfaces can be embedded in other interfaces
* Implicit implementation
  
</br>
</br>
