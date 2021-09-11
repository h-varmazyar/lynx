<p align="center"><img src="https://raw.githubusercontent.com/mrNobody95/lynx/master/lynx.png" width="360"></p>

# Lynx
Lynx is a Golang application for easily creating simple Go project

## Install
Just run below commands for install lynx in your system

```
go get -u github.com/mrNobody95/lynx
```

and then run

```sh
go install github.com/mrNobody95/lynx
```

## Usage

after installing lynx you can create project every where you want in your computer by executing lynx like below:

```sh
lynx new [project name]
```

customizing your project creation with lynx flags. all of flages listed as below:

```sh
lynx new [project name] [flages]
```

| Flag | short hand | description |
|------|------------|-------------|
|--skip-env|-e|Ignore adding env file in the project|
|--skip-go-module|-u|Ignore adding go module to the project|
|--skip-docker|-d|Ignore adding docker to the project|
|--skip-git|-g|Ignore adding git to the project|
|--force|-f|Forcibly create project. delete all previous files if available|

## Give a Star! :star:
If you like or are using this project, please give it a star. Thanks!
