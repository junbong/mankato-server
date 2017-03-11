# mankato-server


## Introduction
Makato is key-value type in-memory database project.
It`s under development.


## Key Features
#### Key-Value NO-SQL Database
Mankato is key-value type in-memory database.
Key is character sequence, and value can be *anything* you want!

#### Blazing Fast
Mankato made by Go lang which is amazingly fast 
and all of it`s data structures use native structures.

#### TTL  / Expiration Callback
Mankato provides expiration callback.


## Project structure
```
mankato
└── src
    └── github.com  
        ├── Junbong
        │   ├── mankato-go-client
        │   └── mankato-server
        └── other_libraries_in_here  
```


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

Build sources. Before build, you need to get dependencies in section below.
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

### mux
Use MUX HTTP router.

https://github.com/gorilla/mux
```sh
$ go get github.com/gorilla/mux
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

### amqp
Use for AMQP expiration callback.

https://github.com/streadway/amqp
```sh
go get github.com/streadway/amqp
```
