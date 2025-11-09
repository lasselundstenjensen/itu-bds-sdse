---
title: "Typical Dagger pipeline structure"
---

## Before we dive in

We will be using Go for the example pipeline definition.

There are currently two ways to define a Dagger pipeline:

1. Dagger functions, e.g. a function for each stage (build, test, deploy etc.).
2. A single Dagger pipeline function that contains all the stages.

Dagger functions are relatively new and while they would in practice be more ideal to encapsulate and moduralize the pipeline, we will use the single Dagger pipeline function for the exercises. It is compact and simpler to maintain when workflows are small.

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

## Typical Dagger pipeline structure

1. A pipeline in Go in its simplest form is a single file, e.g. `pipeline.go`.
2. The pipeline starts through the entry point function `main()`.
3. A Dagger client and background context is created. This is what orchestrates Docker under the hood and is responsible for building the execution graph.
4. Pipeline stages are defined. These consist of the code you want to run, e.g. shell commands.

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

![Show me more](https://media1.tenor.com/m/4EYaYAMus0cAAAAd/michael-scott-the-office.gif)

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

## How it works

![Dagger how it works with OCI](../../images/lessons/dagger/dagger-oci.png)

</br>

1. Your Go program imports the Dagger Go library.
2. Using the Go library, your program opens a new session to a Dagger Engine: either by connecting to an existing engine, or by provisioning one on-the-fly.
3. Using the Go library, your program prepares API requests describing pipelines to run, then sends them to the engine. Docker containers are created on the fly as per the code specifications, in this case in Go.
4. When the engine receives an API request, it computes a Directed Acyclic Graph (DAG) of low-level operations required to compute the result, and starts processing operations concurrently.
5. When all operations in the pipeline have been resolved, the engine sends the pipeline result back to your program.
6. Your program may use the pipeline's result as input to new pipelines.

</br>

*Source: https://archive.docs.dagger.io/0.9/sdk/go/*

</br>
</br>

Ok.. Let's continue and start building a pipeline! <i class="fa-solid fa-rocket"></i>

</br>
</br>
