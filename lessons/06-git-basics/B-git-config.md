---
description: "Basics of Git."
title: "git config"
keywords:
  - Software engineering
  - Git
---

## First, inspect your Git state

Make sure you have Git installed and configured.  You can check your Git version by running:

```bash
git version
```

**Output:**
> git version 2.46.0

</br>

Ensure the default branch is set correctly:

```bash
git config --get --global init.defaultBranch
```

**Output:**
> main

</br>
</br>
</br>

(you should do that now)

</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>

## Configuring your own **git**

Before getting started, we need to configure git to contain _your_ information. Every commit comes with the author's name and email address.

To ensure we all get **credit** and **blame** for our code, we need to set our name and email.

</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>

### git config

In Git there is a global configuration file for all projects on your machine. You can think of the global setting as "good default". There is also a local project level configuratin.

Project settings can override global settings.

In some projects you may want to use your personal details, like a different email.

</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>

**Some facts about `git config`:**

- All gif config keys are in the shape of `<section>.<key>`. we refer to them as `<key>` for now.
- You can set a key with `--global` to set it globally, meaning it counts for all new repos as well, unless overridden.
- The keys used in creating a commit tied to you are `user.name` and `user.email`.
- Add a key value pair with `git config --add --global <key> "<value>"`.
- You can view any value of git config by executing `git config --get <key>`.
- You can view all git config by executing `git config --list`, and `git config --list --local` for local settings specifically.

</br>

> **Note**: You _change_ a value by setting it again.

> **Note**: If you have multiple key-value pairs, the last one set will get used. You can always remove all of them with `git config --unset-all <key>`.

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

### Exercise: Configure git

> Add your `user.name` and `user.email` to your global git configuration IF you do not already have values for them.

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

### Solution: Configure git

Check if you have a name and email set:

```bash
git config --get user.name
git config --get user.email
```

If you get nothing out, we need to add both `user.name` and `user.email`.

```bash
git config --add --global user.name "Your Name"
git config --add --global user.email "your@email.com"
```

</br>
</br>
