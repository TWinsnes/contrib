# contrib

Command line tool to help document contributors to a git project. These will be written to a file in markdown format.

![Demo](demo.gif)

## Installation

Install with homebrew
```shell
brew install twinsnes/tap/contrib
```

Or install by downlaoding the binary from the [releases page](https://github.com/twinsnes/contrib/releases)

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



<!---Contrib Block Start-->
## Contributors

Without these amazing people, this project wouldn't exist.

- Thomas Winsnes
<!---Contrib Block End-->

---
This app was scaffolded using [cligen](https://github.com/twinsnes/cligen) ❤️
