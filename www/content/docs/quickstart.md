+++
weight = 100
date = "2025-10-11T17:00:00+10:00"
draft = false
author = "Thomas Winsnes"
title = "Quickstart"
icon = "rocket_launch"
toc = true
description = "A quickstart guide to install and use contrib"
publishdate = "2025-10-11T17:00:00+10:00"
tags = ["Beginners"]
+++
## Installation

Install with homebrew
```shell
brew install twinsnes/tap/contrib
```

More options are available on the [installation](docs/install) page.

## Usage

To use contrib, call the binary from the root of your git project. It wil scan the git history from the current HEAD and write a list of contributors sorted by number of commits to a file.

```shell
contrib
```

To see more usage options use the `--help | -h` flag

```shell
cotrib -h
```

### Controlling where in the file the list is written
The contributor list will be appended at the end of the file by default. You can also control where in the file the list is written by adding the stand and ends tags where you would like the list to be written.

```markdown
<!---Contrib Block Start-->
This will be replaced by the list of contributors.
<!---Contrib Block End-->
```


### Output file
By default, it will write these to the README.md file in the root of your project.

To change the output file, use the `--path | -p` flag to set the path

```shell
contrib --path CONTRIBUTORS.md
```

or

```shell
contrib -p CONTRIBUTORS.md
```
