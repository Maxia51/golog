# Golog

this is a small project to understang go

## How to use it

```shell
go get github.com/maxia51/golog
```

```go

package main

import "github.com/maxia51/golog"

func main() {

    // Set Log Level
    golog.SetLogLevel(golog.TRACE)

    // Set full color line or only error level color
    golog.SetColorLevel(true)

    // Print log - [ 08/11/2019 - 15:45:52 ]  [ TRACE ]  [ main.main ] could not connect to db
    golog.T("could not connect to db")
    // Print log - [ 08/11/2019 - 15:45:52 ]  [ DEBUG ]  [ main.main ] could not connect to db
    golog.D("could not connect to db")
    // Print log - [ 08/11/2019 - 15:45:52 ]  [ INGO ]  [ main.main ] could not connect to db
    golog.I("could not connect to db")
    // Print log - [ 08/11/2019 - 15:45:52 ]  [ WARNING ]  [ main.main ] could not connect to db
    golog.W("could not connect to db")
    // Print log - [ 08/11/2019 - 15:45:52 ]  [ ERROR ]  [ main.main ] could not connect to db
    golog.E("could not connect to db")
    // Print log - [ 08/11/2019 - 15:45:52 ]  [ FATAL ]  [ main.main ] could not connect to db
    golog.F("could not connect to db")
    // Print log - [ 08/11/2019 - 15:45:52 ]  [ PANIC ]  [ main.main ] could not connect to db
    golog.P("could not connect to db")

}

```
