---
description: "Fundamentals of the Go programming language."
title: "Struct embedding"
keywords:
  - Software engineering
  - Go
---

## Struct embedding

Let's say we want to create a workshop based on our `Course` struct.

We could create a new struct `Workshop` and copy all the fields from `Course` into it. But that would be a lot of duplication.

Let's start by creating a new struct `Workshop`:

```go
type Workshop struct {
    Course Course
    Date   time.Time
}
```

</br>

We can instantiate and print it:

```go
func main() {
    var w = Workshop{
        Course: Course{Name: "Go", Instructor: Instructor{FirstName: "Lasse", LastName: "Jensen"}},
        Date:   time.Now(),
    }
    fmt.Printf("%v", w)
}
```

</br>

Looks good? Well, it works, but we can do better.

</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>

To avoid duplication, we can embed the `Course` struct into the `Workshop` struct:

```go
type Workshop struct {
    Course
    Date time.Time
}
```

We simply remove the identifier `Course` and keep the type to embed it.

</br>

> Is it extending the `Course` struct? No, `Workshop` gets a field named `Course`, and the fields and methods of `Course` are *promoted*, so we can use them directly on `Workshop`, e.g. `w.Name` instead of `w.Course.Name`.

</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>

A downside of using embedding is that fields cannot be set directly when creating an instance of the `Workshop` struct.

```go
func main() {
    var w = Workshop{
        Name: "Go", // Error: unknown field 'Name' in struct literal
        Instructor: Instructor{FirstName: "Lasse", LastName: "Jensen"}, // Error: unknown field 'Instructor' in struct literal
        Date: time.Now(),
    }
    fmt.Printf("%v", w)
}
```

</br>

To set the fields, we need to use the `Course` type:

```go
func main() {
    var w = Workshop{
        Course: Course{Name: "Go", Instructor: Instructor{FirstName: "Lasse", LastName: "Jensen"}},
        Date:   time.Now(),
    }
    fmt.Printf("%v", w)
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

One way to make this look nicer is to use the **factory pattern**:

```go
func NewWorkshop(courseName string, instructor Instructor, date time.Time) Workshop {
    return Workshop{
        Course: Course{Name: courseName, Instructor: instructor},
        Date:   date,
    }
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

## Order of properties

What happens if we have the same field in both the `Workshop` and `Course` structs?

</br>

```go
type Workshop struct {
    Course
    Date time.Time
    Name string // Error: duplicate field 'Name' ???
}
```

</br>

> Will we get an error? No, Go will use the field from the `Workshop` struct as it is the "closest".

</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>

We can still set the name of the `Course.Name` field, but we need to be explicit:

</br>

```go
w.Name = "Go workshop"

w.Course.Name = "Go course"
```

</br>

> `w.Course.Name` is also called an embedding property.

</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>

## Sharing of methods

When you embed a struct, you also get access to its methods. Neat!

</br>
</br>
