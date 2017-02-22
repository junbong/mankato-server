# mankato-server


## Introduction
Makato is key-value type in-memory database project.
It`s under development.


## Build and run
First, checkout project to your local.
```sh
cd ~/workspace
mkdir -p mankato/src/github.com/Junbong
cd mankato/src/github.com/Junbong
git clone git@github.com:Junbong/mankato-server.git
```

Set your `$GOPATH` to checked out directory,
```sh
export GOPATH=~/workspace/mankato
```

Build sources. Before build, you need to get dependencies below.
```sh
cd $GOPATH/src/github.com/Junbong
go build
```

An executable file may created; `mankato-server` and just run this:
```sh
./mankato-server
```


## Dependencies
Before get dependencies, change directory to:
```sh
cd $GOPATH
```

### httprouter
Use HTTP router.

https://github.com/julienschmidt/httprouter
```sh
go get github.com/julienschmidt/httprouter
```

### gods
gods is go-data-structure.

https://github.com/emirpasic/gods
```sh
go get github.com/emirpasic/gods
```


### go-yaml
Use for yaml parser for next feature.

https://github.com/go-yaml/yaml
```sh
go get gopkg.in/yaml.v2
```
