# cleango
A short program to learn principles of Clean Code with golang.

## Problem Statement

One fine day, you had to build a `Go` module that generates a summary of a provided directory, however you were told that such a module already exists but the the problem is that it is not written in the best way possible. It works and provides the required functionality (or we hope so!) however its a complete `NO!` to get this to production.

Since you are running on a tight timeline, you can't really start re-writing the whole thing from scratch but you can go ahead and make fixes to this code to make it Production ready!

Also you were able dig out the requirement document for this code and it looks like the following:

```
1. Takes as input a command line argument `--dir` which is a absolute path to a directory in the host filesystem.
2. Traverses over all the files in the dir (excluding the hidden files).
3. Generates a summary of the directory, the summary contains the following:
    1. Name and size of all files in the directory.
    2. Name and count of duplicate files (if any).
    3. Count of files grouped by extension.
```


You are expected to refactor this code using the `Clean Code` principles by [Uncle Bob](https://en.wikipedia.org/wiki/Robert_C._Martin) and raise a Pull Request with your changes.

### Instructions
1. Please keep your branch name as your Github username.
2. Feel free to add files, test cases, edge cases, comments to the code as long as you are making it better.
3. Make sure that the modified code works as expected by the requirements document.


That's all! Go ahead, `git` started!

Pss! Incase you need a refresher on Clean Code principles, [this](https://gist.github.com/wojteklu/73c6914cc446146b8b533c0988cf8d29) might help!

