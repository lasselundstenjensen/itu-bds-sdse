---
description: "Branching strategies."
title: "Working with branches"
keywords:
  - Software engineering
  - Git
---

## Working with branches

</br>

> Why do we need branches?


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

## Some reasons to use branches

- **Isolation**: Branches allow you to work on a feature or bug fix without affecting the main codebase.
- **Collaboration**: Branches allow multiple developers to work on different features simultaneously.
- **Testing**: Branches allow you to test new features or bug fixes before merging them into the main codebase.
- **Experimentation**: Branches allow you to experiment with new features or ideas without affecting the main codebase.
- **Backup**: Branches allow you to create a backup of your work in case something goes wrong.
- **Code review**: Branches allow you to create a pull request for code review before merging it into the main codebase.
- **Release management**: Branches allow you to manage releases by creating release branches for each release.

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

That said, there are several branching strategies available in the industry. Some are:

- **Feature branching**: Create a branch for each feature or bug fix.
- **Release branching**: Create a branch for each release.
- **Gitflow**: A branching model that defines a strict branching strategy.
- **Trunk-based development**: A branching model where all developers work on a single branch.

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

### Branching falls into two categories:

- **Long-lived branches**: Branches that exist for a long time, such as `main`.
- **Short-lived branches**: Branches that are created for a specific feature or bug fix and are deleted after merging.

</br>

As a contributor to a codebase, you use branches to work on new features, bug fixes, or experiments. You create a new branch for each task and work on it until it is *ready* to be merged into the main codebase.

**Ready** (usually, but varies with team contracts) means that the code is *tested*, *reviewed*, and *approved by the team*. Once the code is ready, you **merge** the branch into the main codebase and **delete the branch**.

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

When you become part of a cross-functional team working with software, it is important to master versioning your code and not leave the shared codebase—"the source of truth"—broken for other team members. You should always leave the shared codebase in a working state. Branches play a crucial role here and is a great place to experiment.

But do not fall in the trap of having long-lived branches. **You** will end up with the responsibility of merging and resolving conflicts. This is not a fun task.

</br>
</br>
