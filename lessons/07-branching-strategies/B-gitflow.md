---
description: "Branching strategies."
title: "Gitflow"
keywords:
  - Software engineering
  - Git
---

## Gitflow

Gitflow is a branching model that helps teams manage their codebase by defining a strict branching strategy. It was created by Vincent Driessen and is widely used in the software industry.

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

> **Only** use Gitflow if you are working on a **large project** with **numerous developers** and/or **multiple development teams**. For small projects, Gitflow **IS** overkill and will cause slowdown due to complexity. Large here means projects such as the Youtube mobile app, or projects with 100+ developers across multiple teams where releases are not happening on a daily basis (typically).

> **Advice:** Be pragmatic. And be cautious if you see Gitflow being used in a small project.

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

### Gitflow defines two primary branches:

- `main`: The main branch that contains the production-ready code.
- `develop`: The main branch that contains the latest development code.

</br>

### Additionally, it defines three types of supporting branches:

- `feature`: Branches used to develop new features.
- `release`: Branches used to prepare a new release.
- `hotfix`: Branches used to fix bugs in the production code.

</br>

<div style="background-color: white;">
<img src="../../images/lessons/git/gitflow-primary-branches.svg" />
</div>

</br>

*Source: https://www.atlassian.com/git/tutorials/comparing-workflows/gitflow-workflow*

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

## Feature branches

Feature branches are used to develop new features or bug fixes. They are created from the `develop` branch and merged back into the `develop` branch when the feature is complete.

</br>

<div style="background-color: white;">
<img src="../../images/lessons/git/feature-branches.svg" />
</div>

<!-- ![Feature branches](../../images/lessons/git/feature-branches.svg) -->

</br>

*Source: https://www.atlassian.com/git/tutorials/comparing-workflows/gitflow-workflow*

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

## Release branches

Release branches are used to prepare a new release. They are created from the `develop` branch and merged back into the `develop` and `main` branches when the release is complete.

</br>

<div style="background-color: white;">
<img src="../../images/lessons/git/release-branches.svg" />
</div>

</br>

*Source: https://www.atlassian.com/git/tutorials/comparing-workflows/gitflow-workflow*

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

## Hotfix branches

Hotfix branches are used to quickly implement bug fixes and patch production. They are similar to release and feature branches, but they are created from the `main` branch and merged back into the `main` branch (and `develop` or current `release` branch) when the hotfix is complete.

</br>

<div style="background-color: white;">
<img src="../../images/lessons/git/hotfix-branches.svg" />
</div>

</br>

*Source: https://www.atlassian.com/git/tutorials/comparing-workflows/gitflow-workflow*

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
</br>
</br>

## So.. what is challenging in this from a developer perspective?

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
</br>
</br>

- **Complexity**: Gitflow is complex and can be difficult to understand for new developers.
- **Overhead**: Gitflow introduces overhead in terms of managing multiple branches.
- **Slowdown**: Gitflow can slow down the development process due to the complexity of managing multiple branches.
- **Merge conflicts**: Gitflow can lead to more merge conflicts due to the number of branches that need to be merged.
- **Overkill**: Gitflow is overkill for small projects with a small team of developers.
- **Release management**: Gitflow can be cumbersome for managing releases.

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

However, despite the complexity, Gitflow is often used when products require large development teams (hundreds of developers, tens of teams) and coordination of releases is required. In these cases you also often have release managers coordinating releases and hotfixes.

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

There are simpler alternatives—One is called **trunk-based development**.

</br>
</br>
