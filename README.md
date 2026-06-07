<!-- [![progress-banner](https://backend.codecrafters.io/progress/shell/bc674609-221b-493c-80eb-b8753540b0e0)](https://app.codecrafters.io/users/Black-tag?r=2qF)

This is a starting point for Go solutions to the
["Build Your Own Shell" Challenge](https://app.codecrafters.io/courses/shell/overview).

In this challenge, you'll build your own POSIX compliant shell that's capable of
interpreting shell commands, running external programs and builtin commands like
cd, pwd, echo and more. Along the way, you'll learn about shell command parsing,
REPLs, builtin commands, and more.

**Note**: If you're viewing this repo on GitHub, head over to
[codecrafters.io](https://codecrafters.io) to try the challenge.

# Passing the first stage

The entry point for your `shell` implementation is in `app/main.go`. Study and
uncomment the relevant code, then run the command below to execute the tests on
our servers:

```sh
codecrafters submit
```

Time to move on to the next stage!

# Stage 2 & beyond

Note: This section is for stages 2 and beyond.

1. Ensure you have `go (1.26)` installed locally
1. Run `./your_program.sh` to run your program, which is implemented in
   `app/main.go`.
1. Run `codecrafters submit` to submit your solution to CodeCrafters. Test
   output will be streamed to your terminal. -->
# 🐚 GoShell

A lightweight Unix-style shell built from scratch in Go.

GoShell is a command-line interpreter that supports built-in commands, command parsing, quoting, path resolution, process execution, and filesystem navigation. The project was created to deepen understanding of operating system fundamentals, process management, shell parsing, and command execution.

---

## 🚀 Features

### Built-in Commands

* `echo`
* `pwd`
* `cd`
* `type`
* `exit`

### Command Execution

* Execute programs from the system `PATH`
* Pass command-line arguments
* Support relative and absolute executable paths

### Command Parsing

* Tokenization of user input
* Single quote handling
* Double quote handling
* Escaped character support
* Whitespace normalization

### Filesystem Navigation

* Absolute paths
* Relative paths
* Home directory expansion (`~`)
* Current working directory management

---

## 🏗️ Architecture

```text
User Input
    │
    ▼
 Parser
    │
    ▼
Command{Name, Args}
    │
    ▼
Shell Executor
    │
 ┌──┴───────────┐
 ▼              ▼
Builtins    External Programs
```

The shell is organized into separate modules for parsing, execution, and built-in commands, making it easy to extend and maintain.

---

## 📂 Project Structure

```text
.
├── app
│   ├── internal
│   │   ├── builtins
│   │   │   ├── cd.go
│   │   │   ├── echo.go
│   │   │   ├── pwd.go
│   │   │   └── type.go
│   │   │
│   │   ├── parser
│   │   │   ├── parser.go
│   │   │   └── types.go
│   │   │
│   │   └── shell
│   │       ├── execute.go
│   │       └── shell.go
│   │
│   └── main.go
│
├── go.mod
├── go.sum
└── README.md
```

### Module Overview

| Module     | Responsibility                                   |
| ---------- | ------------------------------------------------ |
| `parser`   | Converts raw user input into structured commands |
| `shell`    | Dispatches commands and coordinates execution    |
| `builtins` | Implements shell-native commands                 |
| `main`     | Interactive REPL loop                            |

---

## ⚙️ Technologies

<p align="left">
  <img src="https://skillicons.dev/icons?i=go,bash,linux,git" />
</p>

* Go
* Bash/Shell Concepts
* Linux Process Model
* POSIX-style Command Execution

---

## 🧠 Concepts Explored

* Process creation and execution
* Command parsing
* Shell built-ins
* PATH resolution
* Working directory management
* Quoting rules
* Escape sequences
* REPL design
* Modular Go application structure

---

## ▶️ Running Locally

```bash
git clone https://github.com/Black-tag/shell-go.git
cd shell-go

go run ./app
```

Example usage:

```bash
$ pwd
/home/user/projects

$ echo Hello World
Hello World

$ cd ~/Documents

$ type cat
cat is /usr/bin/cat

$ ls
file1.txt
file2.txt
```

---

## 🎯 Future Improvements

* Environment variables
* Pipes (`|`)
* Input/output redirection (`>`, `<`, `>>`)
* Command history
* Auto-completion
* Job control
* Signal handling

---

## 📜 License

This project is intended for learning systems programming and shell internals using Go.
