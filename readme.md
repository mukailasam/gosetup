<h1 style="font-family:monospace;">Gosetup</h1>
<p style="text-align:center">
<img style ="border-radius:2%" src="gosetup.png">
</p>

## About

Gosetup is a Go project setup tool for Go backend developer. It speeds up the process of creating a Go project and including its dependency packages.

**_Add it to your toolbox_**

## Prerequisite

- Go

## Install

```
go install github.com/mukailasam/gosetup@latest
```

### Usage

## Setting up a new project

default command to setup Go project

```
gosetup
```

## Create a Go workspace prior to creating a new Go project

Make use of this command in a situation whereby you would like to have your Go projects in a Go workspace, in which you will be having multiple Go project, not just one.

```
gosetup -ws=true
```

## Go straight to installing depedency packages

Make use of the command in a situation whereby you already have a go project and all you wanted to do is include some dependency packages. make sure you are in the project root directory

```
gosetup -dp=true
```

Note: if you try creating a Go project or workspace with a name that already exist, you will giving options whether you would like to continue. if you are to pick "yes" make sure the directory is a real Go project or workspace directory
