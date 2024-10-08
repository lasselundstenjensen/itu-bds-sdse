---
description: "Fundamentals of the Go programming language."
title: "Singly Linked List"
keywords:
  - Software engineering
  - Go
---

## Exercise: Singly Linked List

### Instructions

One basic data structure used in many cases in the linked list.  One special case of this data structure is the "singly linked list" data structure.  In this exercise, we will build a simplified singly linked list data structure in Go, with the following functionalities:

* `InsertLast` - inserts a new node at the end of the list.
* `DeleteLast` - deletes the last node from the list.
* `Display` - prints the list from beginning to end.

The linked list should accept a data of type `string`, e.g. `InsertLast(data string)`.

</br>

Using it should look something like this:
  
```go
list := SinglyLinkedList{}

list.InsertLast("Hello")
list.InsertLast("My")
list.InsertLast("World")
list.InsertLast("!")

list.Display()
```

</br>

When displaying the list, print the data of each node—e.g. output might look like this:

```plaintext
"Hello" -> "My" -> "World" -> "!"
```

</br>

### Extra (optional) — Implement one or more of the following methods

* `InsertFirst` - inserts a new node at the beginning of the list.
* `DeleteFirst` - deletes the first node from the list.
* `InsertAt` - inserts a new node at a given position.
* `DeleteAt` - deletes a node at a given position.
* `Reverse` - reverses the list.

</br>

> **More information:** 
> - https://en.wikipedia.org/wiki/Linked_list
> - https://www.wscubetech.com/resources/dsa/singly-linked-list-data-structure

</br>
