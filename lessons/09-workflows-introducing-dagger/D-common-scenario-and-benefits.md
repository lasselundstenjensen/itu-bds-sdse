---
title: "The common scenario and Dagger benefits"
---

## The common scenario

</br>

### Disruptive experimentation

You have a pipeline, either new or existing, and when you make changes to it, you have to push the changes to your repository and wait for the pipeline to run (or run it manually).

This can be slow, but it can also be a problem if you are experimenting with changes in a pipeline that others depend on.

This is typically the case for pipelines on development branches.

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

## Dagger benefits

There are two main reasons for adopting Dagger:

1. **Running pipelines locally.** </br>
Meaning you can test most changes before pushing them to your repository. It allows for code to be reused.
2. **Writing pipelines as code.** </br>
Often we learn flavors of YAML or other schemas, but Dagger allows you to write pipelines in languages many of us are familiar with.

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

## Agentic workflow automation

Dagger supports encapsuling agents into Dagger workflows, which can be run locally or in the cloud, and made reproducible. This could be agents using tools to perform tasks, or agents using other agents to perform tasks.

We will not cover this in the course as it is a separate topic.

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

## Supported languages

Dagger currently offers SDKs for the following languages:

- [Go](https://pkg.go.dev/dagger.io/dagger)
- [Python](https://dagger-io.readthedocs.io/en/sdk-python-v0.19.6/)
- [TypeScript](https://docs.dagger.io/reference/typescript/modules)
- GraphQL
- PHP

For a list of supported SDKs and latest links, see the [Dagger documentation, API and SDKs](https://docs.dagger.io/reference#api-and-sdks).

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

## Supported operating systems

Dagger runs on the following platforms:

- MacOS
- Linux
- Windows

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

## Supported platforms

</br>

![Dagger platforms](../../images/lessons/dagger/supported-ci-cd-platforms.png)

</br>
</br>
