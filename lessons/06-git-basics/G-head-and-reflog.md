---
description: "Basics of Git."
title: "HEAD and reflog"
keywords:
  - Software engineering
  - Git
---

## HEAD

By now you might have noticed that it says `HEAD` several places in the logs.

Can you guess what it means...?

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

### Exercise: Check out `HEAD` in the `.git` directory

(not git checkout—just check it out as in: look at it)

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

### Solution: Check out `HEAD` in the `.git` directory

`HEAD` is a reference to the current commit you have checked out. It is a pointer to the branch you are currently on.

We should see something like this.

```bash
cat .git/HEAD
```

**Output:**

```plaintext
ref: refs/heads/main
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
</br>
</br>

And additionally, we can see the contents of the file `refs/heads/main`.

```bash
cat .git/refs/heads/main
```

**Output:**

```plaintext
562f95864c377acd831d1027eeba98fa86875d6d
```

So, literally just a reference to the commit hash you are on.

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

## Reflog

This is a bit of a cool command.

It shows a log of changes to `HEAD`. So—where `HEAD` has been pointing to.

```bash
git reflog
```

</br>

(try it out now)

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

**Output:**

```plaintext
562f958 (HEAD -> main) HEAD@{0}: checkout: moving from feature to main
312c2e9 (feature) HEAD@{1}: checkout: moving from feature-rebase-main to feature
109ad4e (feature-rebase-main) HEAD@{2}: checkout: moving from main to feature-rebase-main
562f958 (HEAD -> main) HEAD@{3}: checkout: moving from feature-rebase-main to main
109ad4e (feature-rebase-main) HEAD@{4}: rebase (finish): returning to refs/heads/feature-rebase-main
109ad4e (feature-rebase-main) HEAD@{5}: rebase (pick): C
b12a712 HEAD@{6}: rebase (pick): B
562f958 (HEAD -> main) HEAD@{7}: rebase (start): checkout main
312c2e9 (feature) HEAD@{8}: checkout: moving from feature to feature-rebase-main
312c2e9 (feature) HEAD@{9}: checkout: moving from main to feature
562f958 (HEAD -> main) HEAD@{10}: merge feature-two: Fast-forward
1fef215 HEAD@{11}: checkout: moving from feature-two to main
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
</br>
</br>

> **Question:** What do you think is stored in the `.git/logs` directory?

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

You can limit search like this.

```bash
git reflog -3  # shows the last 3 entries
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
</br>
</br>

## Recovering lost commits

If you accidentally delete a branch, you can recover it by checking out the commit hash from the reflog.

```bash
git reflog 
# find the commit hash
git merge <commit-hash>
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
</br>
</br>

### Exercise: Recover a lost commit

> * Create a new branch `feature-delete-test` based on `main`.
> * Checkout `feature-delete-test`.
> * Add a new file with some content (important).
> * Commit the changes.
> * Switch back to `main`.
> * Delete the branch `feature-delete-test`.
> * Try to recover the branch using the reflog and merge.

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

### Solution: Recover a lost commit

Create the branch and the change and commit it.

```bash
git checkout main
git checkout -b feature-delete-test
echo "some text" >> tester.md
git add .
git commit -m "some commit message"
```

**Output:**

```plaintext
[feature-delete-test 058f2b5] some commit message
 1 file changed, 1 insertion(+)
 create mode 100644 tester.md
```

</br>

Delete the branch.

```bash
git checkout main
git branch -D feature-delete-test
```

</br>

Recover the lost commit.

```bash
git reflog -3
```

**Output:**

```plaintext
562f958 (HEAD -> main) HEAD@{0}: checkout: moving from feature-delete-test to main
058f2b5 HEAD@{1}: commit: some commit message
562f958 (HEAD -> main) HEAD@{2}: checkout: moving from main to feature-delete-test
```

</br>

Copy the hash of the commit, in this case `058f2b5`.

Create a new branch and recover the commit. Creating a new branch is purely for keeping the changes in a branch.

```bash
git checkout -b feature-delete-test
git merge 058f2b5
```

**Output:**

```plaintext
Updating 562f958..058f2b5
Fast-forward
 tester.md | 1 +
 1 file changed, 1 insertion(+)
 create mode 100644 tester.md
```

</br>

Check the log to see that the commit is back.

```bash
git log
```

**Output:**

```plaintext
commit 058f2b5e13773f4e6aca1fa4ebe412ddab64ea5a (HEAD -> feature-delete-test)
Author: Lasse Lund Sten Jensen <lajl@itu.dk>
Date:   Tue Sep 24 00:04:13 2024 +0200

    some commit message

commit 562f95864c377acd831d1027eeba98fa86875d6d (main)
Author: Lasse Lund Sten Jensen <lajl@itu.dk>
Date:   Mon Sep 23 18:54:56 2024 +0200

    Y

commit 37fb15232e7b8cc90d9f7e7470d95331936956e5
Author: Lasse Lund Sten Jensen <lajl@itu.dk>
Date:   Mon Sep 23 18:54:37 2024 +0200

    X

commit 1fef21501625d1cb1ab99318ad0fa8487d6ef5cb
Author: Lasse Lund Sten Jensen <lajl@itu.dk>
Date:   Fri Sep 20 22:27:19 2024 +0200

    E

commit 8ae69da5d2938a5fee06207f2c3320a4c5b90a9b
Author: Lasse Lund Sten Jensen <lajl@itu.dk>
Date:   Fri Sep 20 22:27:06 2024 +0200

    D

commit 16b8f7f933c4c92b83ebb8602109c84c2d799359
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
</br>
</br>

This may land us a lot of diffs, which may not be desirable if there are a lot of changes between where we are in the branch and the commit hash we are trying to bring back.

There is an even better way—cherry-pick the commit hash.

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

```bash
git cherry-pick <commit-hash>
```

Know that the command exists, but we will not go into it further in this course.

</br>
</br>
