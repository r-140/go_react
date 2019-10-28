# go_react

curl http://localhost:6768/news/10000

curl http://localhost:6768/news

to build backend run

> export GOOS=linux
> go build -o newsservice-linux-amd64
> export GOOS=darwin



docker rm $(docker ps -a -q)
docker rmi $(docker images -q)

all env variables must be defined in .env files
Every variable must starts with the prefix REACT_APP_, otherwise, this variable wil be ignored by React

to install mongo


to run backend ervice in terminal
1.export GOPATH=/Users/Illia_Ushakov/go_workspace/react_go_graphql/bknd/newsservice
1.  export PORT=6767
2.  export MONGODB_URL=mongodb://mongodb:27017
3.  go run main.go


graphql
curl http://localhost:6767/graphql -k -d {News(id: "5db307734a2ab9fd2bd7e695")}


curl -d '{News(id: "5db307734a2ab9fd2bd7e695") {title, teaser, body comments{id, body}}}' -X POST -H "Content-Type: application/graphql" http://localhost:6767/graphql

