# go_react

curl http://localhost:6768/news/10000

curl http://localhost:6768/news

to build backend run

> export GOOS=linux
> go build -o newsservice-linux-amd64
> export GOOS=darwin


docker rm $(docker ps -a -q)
docker rmi $(docker images -q)
