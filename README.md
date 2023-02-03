# prodo-golag-api

install go 

set up env path
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:$PATH

then set up env variable


then run server by this command
go run main.go


//for server not end issue search with this command
lsof -i     
and kill the pid process with name main
kill -9  {{pid}}