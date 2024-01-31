# Learning GO with test

This repository is a playground for learning GO while following the book **Learn Go with test**.
The goal for this repo is to keep track of my learning process and documment aspects of the languages I consider key
and extra relevant for covering the GO learning curve.


## Commands handbook

### Execute a go program
```bash
go run hello.go
```

### execute tests
```
go test hello
```

### Install Godoc

```bash
go install golang.org/x/tools/cmd/godoc@latest
```
Then add the bin folder to the PATH variable

```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

Test the docs server

```bash
godoc -http=:6060
```

Navigate to pkg on the browser at http://localhost:6060/pkg/
