---
title: "Multi-platform and common use cases"
---

# Multi-platform

- Can generate executable binary files for different platforms and operating systems
- Compiles to WebAssembly (WASM)
- Transpiles to JavaScript (GopherJS)

[Compatibility](https://gist.github.com/asukakenji/f15ba7e588ac42795f421b48b8aede63) with different platforms (unofficially).

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

# Common use cases

- Command Line Interface (CLI) tools
- Microservices
- Web servers
- Backend code
- DevOps/MLOps
- Glue code

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

# Can Go be used for Machine Learning?

Yes, but it depends for what. Python is the most popular language for Machine Learning.

Several reasons for this, but the main one is the number of libraries and frameworks available for Python and lack of [basic support in Go for e.g. CUDA and math/statistics libraries](https://blog.jetbrains.com/go/2023/06/22/does-machine-learning-in-go-have-a-future/).

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

Go is powerful when it comes to preprocessing, data cleaning, and manipulating large amounts of data.

It is in large due to its speed and concurrency model:

- High concurrency
- High performance
- Suitable for real-time AI applications
- Suitable for data processing, big data, and data streaming
- Great cloud support through libraries
- Efficient use of resources, fast cold start

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

## ML Go specific libraries

- Natural Language Processing (NLP) `go-nlp`
- Computer Vision `gocv`, with CUDA support

![gocv](../../images/lessons/golang-introduction/gocv-logo.png)

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

## Libraries and frameworks

Some [inspiration](https://github.com/avelino/awesome-go?tab=readme-ov-file).

</br>
</br>