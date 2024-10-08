---
description: "Basics of Git."
title: "Creating a new repository"
keywords:
  - Software engineering
  - Git
---

## Creating a new repository

Before we can use git, we need to create a repo. To create a repo use the `git init` command.

But, before we do that, we also need to create a new or use an existing directory to store our files. There can be no existing version control in the directory.

When inside the directory, type `git init` to create a new repository. Notice it will create a `.git` directory.

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

### Exercise: Create a new repository

> Create a new git repo `my-git-repo` in a new directory. You can put this anywhere as long as it is not in an existing git repo.

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

### Solution: Create a new repository

```bash
mkdir my-git-repo
cd my-git-repo
git init
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

## Everything is in the repository

Every single git repo comes with a `.git` directory. This directory contains all the information about the repository.

> This means, if you delete the `.git` directory, you will lose all the history and changes. But it is also how you can remove git from a directory.

</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>

### Exercise: Verify the git repository

> Verify that you have created a repository by listing out the files and directories in the git state directory.

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

### Solution: Verify the git repository

From within `my-git-repo`, run `find .git`:

**Output:**

```bash
.git
.git/config
.git/objects
.git/objects/pack
.git/objects/info
.git/HEAD
.git/info
.git/info/exclude
.git/description
.git/hooks
.git/hooks/commit-msg.sample
.git/hooks/pre-rebase.sample
.git/hooks/sendemail-validate.sample
.git/hooks/pre-commit.sample
.git/hooks/applypatch-msg.sample
.git/hooks/fsmonitor-watchman.sample
.git/hooks/pre-receive.sample
.git/hooks/prepare-commit-msg.sample
.git/hooks/post-update.sample
.git/hooks/pre-merge-commit.sample
.git/hooks/pre-applypatch.sample
.git/hooks/pre-push.sample
.git/hooks/update.sample
.git/hooks/push-to-checkout.sample
.git/refs
.git/refs/heads
.git/refs/tags
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

This is basically what **peak empty** git repository looks like.

Essentially git is stateless. Every time we run a command, git has to figure out where we are and what has changed etc.

</br>
</br>
