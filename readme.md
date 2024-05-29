<h1 style="font-family:monospace;">Gosetup</h1>
<p style="text-align:center">
<img style ="border-radius:20%" src="gosetup.jpg" width="350px" height="300px">
</p>

## About

Gosetup is a Go project setup tool for Go backend developer. It automate the "process of creating a Go project and including its dependencies packages".

**_Add it to your toolbox_**

## Prerequisite

- Go

## Install

```
go install https://github.com/mukailasam/gosetup
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

Make use of the command in a situation whereby you already have a go project and all you wanted to do is include some dependency packages

```
gosetup -dp=true
```

Note: if you try creating a Go project or workspace with a name that already exist, you will giving options whether you would like to continue. if you are to pick "yes" make sure the directory is a real Go project or workspace directory
