---
description: "SSH and GitHub"
title: "Connecting to GitHub via SSH keys"
keywords:
  - Software engineering
  - Git
---

# Connecting to GitHub via SSH keys

Before getting into SSH keys, we need to understand what **git** and **GitHub** are and how they are different.

</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>

## Git vs. GitHub

**Git** is a piece of software that you can install and run locally. It is a version control system (VCS) for tracking changes in, primarily, code files that are within the same directory, which then constitutes a repository.

**GitHub** is an online website that functions as a storage and collaboration platform for **git** repositories. It essentially hosts your code on their servers, and provides a graphical user interface as well as features for easing code collaboration, e.g. pull requests.

</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>

## Why SSH keys for GitHub?

In short, it's a secure way to identify yourself (so people know it's really you who sent a pull request with 10,000+ changed lines in one commit), and it saves you from having to log in each time you want to store files on GitHub.

In GitHub's own words:
> Using the SSH protocol, you can connect and authenticate to remote servers and services. With SSH keys, you can connect to GitHub without supplying your username and personal access token at each visit. You can also use an SSH key to sign commits.

**But for it to be secure, you should *never* give out your *private* key to anyone**, though the public key is fine to give out.

See the private key as your ID card, and the public key as your full name. It's generally fine to give out your name to other people, and in this case, GitHub will ask for your ID to confirm you are whom you claim to be, that you are allowed to work in this repository, etc.

</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>

# How to generate and add SSH keys for GitHub

This is specifically for generating SSH keys using the terminal.

</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>

## Step 0 - Check for existing SSH keys
```bash
ls -al ~/.ssh
```

You might already have a key pair or two, or maybe even none. Either way is fine. Just know that the rest of the guide will use `~/.ssh/id_rsa_github` as the new key file's location.

If that file already exists on your computer, remember to replace the string for the rest of the guide, e.g. `~/.ssh/id_rsa_datascience`.

</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>

## Step 1 - Generate SSH key pair

**Option 1:** With your e-mail as a comment (Optional)

Replace `<your_email@example.com>` below with your actual e-mail address.  See it as metadata to help you (or a future organisation administrator in charge of your team) to differentiate between multiple key pairs.

``` bash
ssh-keygen -t rsa -b 4096 -C "<your_email@example.com>"
```

**Option 2:** Without your e-mail as a comment

``` bash
ssh-keygen -t rsa -b 4096
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

## Step 1.5 - Specify location

The agent will prompt you for which file to write the secret key to. Don't worry about the public key here.

If the secret key file doesn't exist, it will create one. The same is true for the public key file.

</br>

Replace `<user>` below with the username you're using on the machine.

```
/home/<user>/.ssh/id_rsa_github
```

This will write the generated secret key to the specified file, and the public key to a new file `~/.ssh/id_rsa_github.pub`.
</br>
</br>

Know that `~` is a shell expansion (do not worry about this term), which expands to `/home/<user>/`, but the SSH agent might complain about such a directory not existing due to it being interpreted as the literal character rather than being expanded. Writing it out explicitly should avoid that. A possible alternative could also be to use the environment expansion `$HOME` instead as this may be more widely supported than `~`.

</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>

## FYI
If you had just pressed 'ENTER' (agreed to default behaviour), it will write the generated keys to the following files:
- `~/.ssh/id_rsa` (secret key)
- `~/.ssh/id_rsa.pub` (public key)

So, if you save your SSH keys in that same folder (`~/.ssh/`), you can easily find all of your SSH key in there.

</br>

> **Question:** Say you did just press 'ENTER' here and go about your day. Later, you generate a new key pair, e.g. for your new GitHub account used specifically for work. You follow all the steps, except again, when prompted on where to write the generated keys, you merely press 'ENTER'. What would happen then?

</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>

## Step 2: Add a passphrase (Optional but HIGHLY recommended)

Imagine having a student ID without a pincode, or even a phone without a pincode...

Please take the recommendation to heart.

</br>

**[ No, thanks? ]** Press 'ENTER', and continue to the next step.

**[ Yes, please? ]** Ensure that you will remember it: Once lost, it cannot be recovered, and you'll just have to replace your SSH key pair with new ones.

When typing in your passphrase, you will not see any characters being typed --- that's just security by obscurity. It will still take your input.

Afterward, it will prompt you to type in the same passphrase again...

</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>

## Step 3: Start the SSH agent

```bash
eval "$(ssh-agent -s)"
```

**Curious about what's happening?**

Running `ssh-agent -s` prints shell commands meant to configure the shell environment. These commands would be run in a shell sub-process rather than the current shell process, so the commands are unable to do the necessary environment changes to the current shell process. You can see this as if a child is babbling on about how to redecorate the house. Without the parent listening and decorating, nothing would change.

So, `eval` takes the shell commands as arguments, and **eval**uates each command, essentially executing them within the current shell process. Thus, the current shell environment has been configured by the agent. Continuing with the house decoration allegory, the parent can now hear the child, evaluates their wishes, and redecorates accordingly.

</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>

## Step 4: Add SSH key to SSH agent

```bash
ssh-add ~/.ssh/id_rsa_github
```

It will prompt you for the key's passphrase if one was set.

</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>

## Step 5: Display the contents of your public key

```bash
cat ~/.ssh/id_rsa_github.pub
```

... which outputs a string with the content: `ssh-rsa <long gibberish>= <computer_name>`, or something similar to that.

</br>

**Fun fact:** The command `cat` is used to read a file, or multiple files, sequentially and print it to the standard output. Shorthand for con**cat**enate files.

</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>

## Step 6: Add SSH key to GitHub
1. Copy the contents of the SSH **public key** file (.pub)
2. Access your GitHub account settings.
3. Navigate to **SSH and GPG keys**.
4. Click on **New SSH key**.
5. Give it a title, and ensure that **Key type** is set to 'Authentication Key'.
6. Paste your SSH public key into the provided **Key** field.
7. Click **Add SSH key**.

</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>
</br>

## Step 7: Start cloning repositories via SSH, not HTTPS links

You might have tried cloning git repositories from GitHub using HTTPS links like so:

```bash
git clone https://github.com/<repo_owner>/<repo_name>
```

</br>

But to use the new SSH key pair to clone and skip manually logging into GitHub, we run:

```bash
git clone git@github.com:<repo_owner>/<repo_name>.git
```

</br>
</br>

**To find this string...**
1. Go to the repository you want to clone from GitHub.
2. Click on the green **<> Code** button.
3. Clicon on the **SSH** tab.
4. Copy the SSH URL to the clipboard.
5. Paste it as an argument to the `git clone` command.

</br>
</br>
