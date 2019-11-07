# go_react

this project demonstrates how to work with go, mongodb, graphql, and react+redux

One service wiritten in Go produce content, other two (written in GO and Python) acts as a consumers. They communicates to each oher over grpc

to run docker 

1. build frontend:

- npm install react-redux redux
- npm install prop-types --save
- npm install react-router-dom --save
- npm install redux react-redux redux-thunk --save
- npm install bootstrap 

 install graphql depenedensies
- npm install apollo-boost graphql --save
- npm install apollo-cache-inmemory apollo-link-http --save
- npm install graphql graphql-tag --save
- npm install react-apollo --save

- npm run build

2. build backend

install go 

 export GOPATH=/Users/Illia_Ushakov/go_workspace/react_go_graphql/bknd/newsservice
 export GOOS=linux
 go build -o newsservice-linux-amd64
 export GOOS=darwin

3. docker-compose up


to remove all docker containers and images
docker rm $(docker ps -a -q)
docker rmi $(docker images -q)

to run backend ervice in terminal
1.export GOPATH=/Users/Illia_Ushakov/go_workspace/react_go_graphql/bknd/newsservice
1.  export PORT=6767
2.  export MONGODB_URL=mongodb://mongodb:27017
3.  go run main.go


to debug graphql via curl:
curl http://localhost:6767/graphql -k -d {News(id: "5db307734a2ab9fd2bd7e695")}


curl -d '{News(id: "5db307734a2ab9fd2bd7e695") {title, teaser, body comments{id, body}}}' -X POST -H "Content-Type: application/graphql" http://localhost:6767/graphql


to debug graphql via GUI Graphoql use:
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
  News(id: "5db851d75ec36eb337561b8b") {
    title
    teaser
    body
    comments {
      username
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

nice articles
https://tutorialedge.net/golang/go-graphql-beginners-tutorial/
https://tutorialedge.net/golang/go-graphql-beginners-tutorial-part-2/

https://dzone.com/articles/go-microservices-part-14-graphql

https://www.freecodecamp.org/news/deep-dive-into-graphql-with-golang-d3e02a429ac3/


https://graphqlmastery.com/blog/input-object-type-as-an-argument-for-graphql-mutations-and-queries

https://graphqlmastery.com/blog/graphql-quick-tip-how-to-pass-variables-in-graphiql


https://medium.com/codingthesmartway-com-blog/getting-started-with-react-and-graphql-395311c1e8da



