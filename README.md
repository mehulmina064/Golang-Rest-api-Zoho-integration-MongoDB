# golag-rest-api

install go 

set up env path
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:$PATH

then set up env variable


**for air auto Refresh and rerun setup**
go install github.com/cosmtrek/air@latest

//for the path set up
alias air='$(go env GOPATH)/bin/air'   

 air init     

//for run the server
 air 

then run server by this command for manual refresh
go run main.go


//for server not end issue search with this command
lsof -i     
and kill the pid process with name main
kill -9  {{pid}}
