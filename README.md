# See - Search for Executable Examples

Welcome to See, a command-line tool that allows you to search for executable examples of command-line programs. 
With See, you no longer have to struggle with finding the correct syntax for a specific command.

## Why?

If you create a new CLI, it's natural that you use a framework such as Cobra or `urfave/cli`. Each of your defined commands also has an **Example** section. Let's take the `helm` CLI as an example:

```plaintext
Examples:

    # Search for stable release versions matching the keyword "nginx"
    $ helm search repo nginx

    # Search for release versions matching the keyword "nginx", including pre-release versions
    $ helm search repo nginx --devel

    # Search for the latest stable release for nginx-ingress with a major version of 1
    $ helm search repo nginx-ingress --version ^1.0.0
```

However, to display those examples, you need to run `helm search repo --help`. Isn't that odd? To get the help for what you are looking for, you already need to know at least that it is possible and know the CLI syntax.

What if you could do:

```bash
helm help
```

which opens an interactive window where you can type "search helm chart," and all related example commands are displayed together with information about the command syntax, e.g. `helm search repo [keyword] [flags]`.

## How?

We traverse all commands and index them in the same way as CLI documentation is generated for the Cobra library. 
Next, by using fuzzy search, we display all matching examples and the source for them, e.g., `helm search repo --help` if a given example comes from the `helm search repo` command help message.

## Like the idea? Give a GitHub star ⭐!

The library is work in progress. Please add a star to let me know that you would be interested in such a library!
