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


nice articles
https://tutorialedge.net/golang/go-graphql-beginners-tutorial/
https://tutorialedge.net/golang/go-graphql-beginners-tutorial-part-2/

https://dzone.com/articles/go-microservices-part-14-graphql

https://www.freecodecamp.org/news/deep-dive-into-graphql-with-golang-d3e02a429ac3/


https://graphqlmastery.com/blog/input-object-type-as-an-argument-for-graphql-mutations-and-queries

https://graphqlmastery.com/blog/graphql-quick-tip-how-to-pass-variables-in-graphiql


to create news via graphql send mutation
mutation CreateNewsMutation {
  CreateNewsMutation(title: "gr_news2", teaser: "gr_teaser1", body: "gr_body") {
    id
    title
    teaser
    body
  }
}

to add new comment to news via graphql send mutation
mutation AddCommentMutation {
  AddCommentMutation(newsID: "5db851d75ec36eb337561b8b", username: "ilyitch", body: "this is firat comment") {
    username
    body
  }
}


to request news by id submit query
{
  News(id: "5db307734a2ab9fd2bd7e695") {
    title
    teaser
    body
    comments {
      
      body
    }
  }
}

to get all news submut query
{
  AllNews {
    id
    title
    teaser
	comments {
	  body
      username
	}
  }
}
