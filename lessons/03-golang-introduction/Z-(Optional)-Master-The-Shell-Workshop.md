---
title: "Mastering The Shell"
description: "Basics of how to get around and interact with the shell"
keywords:
  - Software engineering
  - Data science
  - Basic shell learning
  - Theis Per Holm 
---

# The shell as a tool

This is a small overview/exercise sheet on basic shell commands that will help you interact with files and programs in a graceful manner. The shell is a valuable tool and can in most instances be the fastest and most optimal way to work. People who are not used to working in the shell may see it purely as a tool to compile and run code but out of the box it comes with many powerful tools that let you skip clunky graphical user interfaces entirely.

</br>
</br>

## Which shell to use?
For this overview I will assume we are using either bash or zsh (all things covered in this guide should work the same in both out of the box)

## A great start

Most of this guide will be based on the ["missing-semester" - From MIT](https://missing.csail.mit.edu/2020/course-shell/)

- Introduction and basics [here](https://missing.csail.mit.edu/2020/course-shell/)
- Shell scripting [here](https://missing.csail.mit.edu/2020/shell-tools/)
- Command-line environment [here](https://missing.csail.mit.edu/2020/command-line/)

We will mostly cover basics and a small amount of shell scripting but feel free to read through these 3 ressources in depth.

# Lets stretch our legs

One of the most important skills to learn is how to get get around the shell and interact with files. Without the ability to move, read, move or delete files we can't really do a whole lot :)

Explaining the symbols:
user:<something>: user is a standin for your own username the : and then some word is our current directory name  
\$: Anytime you see a \$ it purely marks the start of a shell line this might look different in your shell/terminal if you want to run a command feel free to copy it from the code blocks but do not take the \$ with you it will not run if you do. 
\>: The carrot symbol (>) will be used to show the output of a given command

### cd (change directory)
One of the most used commands will be `cd` which will change your directory to whatever directory you pass it. If left blank it will move you to (~/$HOME) which is your home directory.
```bash
user:current$ cd 
> user:~$ 
```


```bash
user:current$ cd cool_directory
> user:cool_directory$
```

### ls (list) 
To be able to move around you must know what is around you the ls command is useful here as it lists files and directories.

```bash
user:current$ ls
> cool_directory cool_file.md
```

`ls` has flags that allow you to view things like hidden files, show file sizes and even file ownership. The most command flag combination is `ls -la` the `a` flag shows all and `l` is long list this shows a lot more information

```bash
user:current$ ls -la
> drwxr-xr-x - user 22 Sep 08:43 cool_directory
  .rw-r--r-- 0 user 22 Sep 08:43 cool_file.md
```

`ls` doesn't just list files locally it can even list files in other directories by passing a path into it
```bash
user:~$ ls current
> cool_directory cool_file.md
```
Notice how we get the exact same output as when we were in current and ran `ls`

### cat (con(cat)enate)
This is mostly used for reading files but as the name suggests it can also be used to concatenate files together

```bash
user:current$ cat cool_file.md
> # this is a super cool file
it even spans multiple lines
```

### echo (command tool used for displaying lines of text or string)
This will work somewhat like a print in python but its a little more powerful since it can be chained with other shell commands and options to achieve broader goals.

```bash
user:~$ echo hello
> hello
```

We could as an example capture the output of echo in a file (note this file doesn't even have to exist echo will create it for us)


```bash
user:~$ echo hello > file_name
> 
user:~$ cat file_name
> hello
user:~$ echo hello >> file_name
> hello
hello
```

Notice the carrot (>) behind `echo hello` this means "output to x" in this case we are telling it to output to the file named `file_name`. The double carrot (>>) means append so add to the end of.

### pwd (print working directory)
As the name suggests this command will print your current working directory aka. where you are
Note this command will print the full path and thus may look slightly different on Linux and MacOS as the home folder is stored in different locations. If you are Linux `/User` will be `/home`

```bash
user:current$ pwd
> /Users/user/current
```

### man (manual) 
This command is useful for figuring out what other commands do. If you ever want to know what a tool does then man is your friend

```bash
user:~$ man ls 
> 
```
The output here is displayed as empty as this will open a page to view in your terminal

### less (a terminal pager)
Reading longer files using cat is kind of annoying so instead we have another tool that we can use. If you want to know more then you can try `man less` 

```bash
user:current$ less cool_file.md
>
```
This will open a page allowing you to scroll through a file and even search for things. To exit use q

### tail 
This will give you the tail of a file aka the end. You can specify how much you want using flags and can even watch a file using this command. By default it will give you the whole file use `-n` with a number to specify how many lines e.g. `-n1` for just the last line

```bash
user:current$ tail -n1 cool_file.md
> it even spans multiple lines
```

You can watch the file using the -f flag. If somebody appended to the cool_file.md it would display under the cursor
```bash
user:current$ tail -f append_to_me
> 
```
WARNING: This will take control of your shell. If you want out your can always call a SIGKILL using ctrl+C

### find 
This command will let your find or search for files using a pattern. We will not cover what a pattern is here but a simple google search or even the man page may give you a good idea of how to create one that suits your case
#TODO: FINISH
```bash
user:~$ find ./current -name "cool_*"
> ./current/cool_directory
./current/cool_file.md
```
> \* means wildcard aka match anything. so this find command found all files and directories matching cool_<anything> in the `current` directory

### | (piping) 
This header is not a typo a useful tool is the pipe operator denoted by `|` this can be used to carry output through to a different command. This is by far the most useful operator to know on top of the file output and input operators. If we have a basic python program which reads input from stdin and prints it to the user we can pipe the input to it

```bash
user:current$ cat main.py
> print(input())
user:current$ echo hello | python3 main.py
> hello
```
Normally when running this python program the shell would be hijacked and we would have to write something and press return to get the program to do anything. Instead we have now given the program input by using the echo command seen before. 

### > < >> << (file modifiers)
As shown previously we can output to a file using `>` and append to a file using `>>` but this also works the other way we can actually give input to a program using `<`. If we take the same python program from above. `<<` is known as heredoc and is a multiline input passer.

```bash
user:current$ echo hello > input_file
> 
user:current$ cat input_file
> hello
user:current$ python3 main.py < input_file
> hello
```
As seen we get the same output as from the piping example above but in this case we actually read the data from a file and pass that along to the python program. 

> Note: This exact workflow can be extremely useful when developing programs as we can simply pass a file into the program instead of having to type out examples manually. This is extra useful when developing small programs for say programming contests like lille kat where inputs can be quite long and tedious to type out

### Text editors in your terminal
I will list 2 of the common editors that are built into most shells that being `nano` and `vim`. The order of these programs listed is in **my opinion** also how hard they are to learn to use. `nano` is a minimal text editor you can easily see how to write and exit to a file at the bottom using the shortcuts (these are done by holding control and pressing the key shown). `vim` is an extremely powerful text editor but also a mouthful to learn, it can be rewarding but could in of itself be a whole course (there's even a book written about how to learn it which is roughly 300 pages called practical vim). 


# Bash scripting
Everything we've learned so far has been in a single line and even with piping and file modifiers this can only get us so far. What if we could write a file with many bash expressions and commands and run that just like we would a python program?

## Bash scripts
Bash scripts are in essence just a file containing what we've already written above. This file can be run and thus we save ourselves some time by defining a file we can run multiple times. Bash is actually turing complete (don't worry if you don't know what that means) which means you could code entire projects in it (would not suggest this unless you truly hate yourself).

> We are now working in a mix of files and shell commands. If an example is in a file no starting notation will be shown like below. If we are running a shell command it will be notated as seen in the chapter above.

## Small example of a bash script
```bash

```

