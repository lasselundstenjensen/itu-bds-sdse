---
description: "Branching strategies."
title: "Trunk-based development"
keywords:
  - Software engineering
  - Git
---

## Trunk-based development

![Trunk-based development summary](../../images/lessons/git/trunk-based-one-line-summary.png)

</br>

*Source: https://trunkbaseddevelopment.com*

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

Trunk-based development is a branching model where all developers work on a single branch. This branch is often called `main` or `trunk`. The main branch is the source of truth for the project. Developers commit their changes directly to the main branch.

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

### Shared branches off mainline/main/trunk are bad at any release cadence:

</br>

<div style="background-color: white;">
<img src="../../images/lessons/git/trunk1a.png" />
</div>

</br>

*Source: https://trunkbaseddevelopment.com*

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

### Trunk-Based Development For Smaller Teams:

</br>

<div style="background-color: white;">
<img src="../../images/lessons/git/trunk1b.png" />
</div>

</br>

*Source: https://trunkbaseddevelopment.com*

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

### Scaled Trunk-Based Development:

</br>

<div style="background-color: white;">
<img src="../../images/lessons/git/trunk1c.png" />
</div>

</br>

*Source: https://trunkbaseddevelopment.com*

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

## Why trunk-based development?

- **Simplicity**: Trunk-based development is simple. There is only one branch to worry about.
- **Speed**: Trunk-based development is fast. Developers can commit their changes directly to the main branch.
- **Feedback**: Trunk-based development provides quick feedback. Developers can see the impact of their changes immediately. Continuous integration and automated testing help catch issues early.
- **Collaboration**: Trunk-based development encourages collaboration. Developers can work together on the same codebase.
- **Quality**: Trunk-based development promotes quality. It is an great enabler strategy for Continuous Integration and Continuous Delivery (CI/CD).
- **Release**: Trunk-based development makes releases easier. Developers can release new features with less overhead.

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

## Why is this important?

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

<img src="https://media1.tenor.com/m/KsLzwY-GHCYAAAAd/jeremy-clarkson-speed.gif" alt="Speed"  width="100%"/>

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

### Speed

Empirical evidence shows that high performing teams commit to the main branch more frequently, release more frequently, and have shorter lead times. They also have lower change failure rates and faster mean time to recovery.

This practice leads to faster feedback loops and higher quality code.

</br>

> In short and maybe counter intuitively—**going faster** and releasing more frequently actually **leads to higher quality** products.

</br>

![Accelerate, Nicole Forsgren](../../images/lessons/git/accelerate-book.png)

</br>

*Source: https://itrevolution.com/product/accelerate/*

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

> **Advice:** Use this branching strategy as the default option whenever possible. It is simple, fast, and promotes collaboration and quality.

</br>
</br>
