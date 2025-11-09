---
title: "Running Dagger in a GitHub Action workflow"
---

<!-- TODO: Workflows and Actions -->

## Embedding is easy

Embedding Dagger in a GitHub Action workflow is straightforward as the Dagger team provides a GitHub Action that you can use to run Dagger pipelines.

Alternatively, the proces is the same as what we just did, if you do not want to use the provided GitHub Action.

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

### Dagger, GitHub and a runner

To run a Dagger pipeline in a GitHub Action workflow, you need to create a new workflow file in your repository and add the necessary steps to run the pipeline.

This eventually means running the Dagger workflow inside the GitHub workflow.

</br>

The general outline of how Dagger, GitHub workflows and GitHub runners work together looks like this:

</br>

<div style="background-color: white;">
<img src="../../images/lessons/dagger/dagger-on-github-logical.svg" />
</div>

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

### In more detail

</br>

<div style="background-color: white;">
<img src="../../images/lessons/dagger/dagger-on-github-logical-2.svg" />
</div>

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

## GitHub Action Workflow

Create a new file in the `.github/workflows` directory called `dagger.yml` and add the following content:

</br>

```yaml
name: Run Dagger pipeline

on:
  push:
    branches:
      - main
  workflow_dispatch:

jobs:
  dagger:
    name: Run Dagger pipeline
    runs-on: ubuntu-latest

    steps:
      - name: Checkout code
        uses: actions/checkout@v4

      - name: Set up Go
        uses: actions/setup-go@v6
        with:
          go-version: 1.25.4
          cache: false

      - name: Run Dagger pipeline
        uses: dagger/dagger-for-github@v8.2.0
        with:
          workdir: go
          verb: run
          args: go run pipeline.go
          version: "0.19.6"
```

</br>

This [workflow](https://github.com/lasselundstenjensen/itu-bds-sdse-dagger-and-github/blob/main/.github/workflows/dagger.yaml) runs on every push to the `main` branch and when [manually triggered](https://github.com/lasselundstenjensen/itu-bds-sdse-dagger-and-github/actions).

It uses the `dagger/dagger-for-github` action to run the pipeline.

Push the changes to your repository and navigate to the Actions tab in GitHub to see the workflow running.

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

![Finish smile](https://media1.tenor.com/m/rPOqhxQBOz0AAAAC/all-might-my-hero-academia.gif)

</br>
</br>
