# prodo-golag-api

install go 

set up env path
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:$PATH

then set up env variable


then run server by this command
go run main.go

and for air configuration for dynamic reloading and compiling the code and run
1-go install github.com/cosmtrek/air@latest
2- alias air='$(go env GOPATH)/bin/air'
3-air     
