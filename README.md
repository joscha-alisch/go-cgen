# Go CGen

[![Go Version](https://img.shields.io/github/go-mod/go-version/joscha-alisch/go-cgen.svg)](https://github.com/joscha-alisch/go-cgen)
[![GoDoc reference](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/joscha-alisch/go-cgen)
[![GoReportCard](https://goreportcard.com/badge/github.com/joscha-alisch/go-cgen)](https://goreportcard.com/report/github.com/joscha-alisch/go-cgen)
[![Maintainability](https://api.codeclimate.com/v1/badges/75d5757fe6a001f6ea1b/maintainability)](https://codeclimate.com/github/joscha-alisch/go-cgen/maintainability)
[![Test Coverage](https://codecov.io/gh/joscha-alisch/go-cgen/branch/main/graph/badge.svg?token=898J1INMMX)](https://codecov.io/gh/joscha-alisch/go-cgen)
[![Sourcegraph](https://sourcegraph.com/github.com/joscha-alisch/go-cgen/-/badge.svg)](https://sourcegraph.com/joscha-alisch/go-cgen?badge)

# WIP

This package is currently being implemented and nowhere near what this readme outlines. The usage is only described here
as a guide towards our envisioned API interface.

## About

CGen is a go generator for automatic creation of idiomatic go bindings for any C headers.

## Install

```bash
go install github.com/joscha-alisch/go-cgen/cmd/cgen@latest
```

## Quickstart

Put the C header files of the library into your project. For example:

`include/header.h`

```c
typedef enum MyEnum {
    Option = 1
} MyEnum; 
```

Add the following to a go file in the root of your package, just before the package name.

`example.go`

```go
//go:generate cgen -h include/*.h
package mypackage
```

Run `go generate`.

CGen will now generate all bindings at the end of your go file.

`example.go`

```go
//go:generate cgen -h include/header.h
package mypackage

// The following code is autogenerated by CGen. DO NOT EDIT.

type MyEnum int

const (
	Option MyEnum = 1
)
```

## Customisation

To build proper idiomatic go bindings for your C library you will probably want to modify the behaviour of CGen. There
are two parts to look at here.

1. Configure the way your Go code is generated via the [CGen config file](#configuration-file). This includes how
   identifiers are named, how to deal with C pointers, arrays, etc.
2. Configure the way the generated go code is distributed across files and packages. This is done
   via [cgen commands in package files](#code-distribution).

### Configuration File

TODO

### Code Distribution

You can add cgen commands in any go file in the root package or any subpackage.

The commands have the form

```go 
// +cgen:<COMMAND> <CONFIG>
```

For a list of currently supported commands, see [Commands](#commands)

* If they appear at the top of your file (right before the package name), the generated code is appended at the end of
  the file.

  So the following:
    ```go 
    // +cgen:something
    package example
  
    func SomeCustomHandwrittenFunc() {...}
    ```
  Results in:
    ```go 
    // +cgen:something
    package example
  
    func SomeCustomHandwrittenFunc() {...}
  
    // The following code until the end of file is autogenerated by CGen. EDITS WILL BE LOST ON NEXT GENERATION.
    func SomeGeneratedFunc() {...}  
    ```

* If they appear somewhere else in the file, the generated code is inserted right after the tag.

  So the following:
  ```go 
  package example
  
  // +cgen:something
  
  func SomeCustomHandwrittenFunc() {...}
  ```

  results in:
  ```go 
  package example
  
  // +cgen:something
  // The following code is autogenerated by CGen. EDITS WILL BE LOST ON NEXT GENERATION.
  func SomeGeneratedFunc() {...}  
  // End of autogenerated code.
   
  func SomeCustomHandwrittenFunc() {...}
  ```
* When a tag appears right before a handwritten line of code, it applies to this piece of code. Only a subset of cgen
  commands are available there.

  So the following:
  ```go 
  // +cgen:somestructcommand
  type MyStruct struct {
    Field string
  }
  ```
  might result in:
  ```go 
  // +cgen:somestructcommand
  type MyStruct struct {
    Field string
  
    // The following code until the end of the struct is autogenerated by CGen. EDITS WILL BE LOST ON NEXT GENERATION.
    SomeGeneratedField string
  }
  
  // The following code is autogenerated by CGen. EDITS WILL BE LOST ON NEXT GENERATION.
  func (m *MyStruct) SomeGeneratedMethod() {...}
  // End of autogenerated code.
  ```


* When multiple cgen commands are specified on consecutive lines, they are grouped together. With at least one empty
  line between them, they are considered separate and will result in multiple code blocks being generated.

  i.e.
  ```go
  // +cgen:something
  // +cgen:somethingelse
  ```
  is not the same as
  ```go
  // +cgen:something
  
  // +cgen:somethingelse
  ```

### Commands

* [Tags](#tags)

  `// +cgen:tags KEY=VALUE ...`

  Code that matches the given tags is put here.

#### Tags

```go
// +cgen:tags KEY=VALUE KEY=VALUE ...
```

Results in all code that matches the tags being generated here. If some code matches multiple of these lines, the most
specific - i.e. the one with the highest number of tags - is chosen. Given the grouping of cgen commands, the above can
also be written as

```go
// +cgen:tags KEY=VALUE
// +cgen:tags KEY=VALUE
```

## Advanced

If the built-in configurations are not enough for your use-case. You can also build your own generator on top of CGen,
giving you the power to configure almost every part.

(TODO)