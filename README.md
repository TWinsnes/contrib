# Contrib


![GitHub Release](https://img.shields.io/github/v/release/twinsnes/contrib)
![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/twinsnes/contrib/ci.yml)
![GitHub License](https://img.shields.io/github/license/twinsnes/contrib)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/twinsnes/contrib)
![Github Downloads](https://img.shields.io/github/downloads/twinsnes/contrib/total)

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

Check out the [docs](https://twinsnes.github.io/contrib/) for more details information.

<!---Contrib Block Start-->
## Contributors

Without these amazing people, this project wouldn't exist.

- Thomas Winsnes
<!---Contrib Block End-->

---
This app was scaffolded using [cligen](https://github.com/twinsnes/cligen) ❤️
