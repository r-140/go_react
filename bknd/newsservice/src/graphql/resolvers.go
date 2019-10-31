package graphql

import (
	"dbclient"
	"log"
	"model"

	"github.com/graphql-go/graphql"
)

// GraphQLResolvers provides resolvers methods for  graphql schema
type GraphQLResolvers interface {
	NewsResolverFunc(p graphql.ResolveParams) (interface{}, error)
	AllNewsResolverFunc(p graphql.ResolveParams) (interface{}, error)
	CreateNewsResolverFunc(p graphql.ResolveParams) (interface{}, error)
	AddCommentToNewsResolverFunc(p graphql.ResolveParams) (interface{}, error)
}

// LiveGraphQLResolvers implementations
type LiveGraphQLResolvers struct {
}

// NewsResolverFunc grapgql resolver for getNews query
func (gqlres *LiveGraphQLResolvers) NewsResolverFunc(p graphql.ResolveParams) (interface{}, error) {

	log.Println("opening NewsResolverFunc() ")
	newsID, _ := p.Args["id"].(string)

	log.Println(" NewsResolverFunc() ID from argument ", newsID)
	// if newsID == nil {
	// 	panic("erroor parsing news id as an argument")
	// }

	news, err := fetchNews(newsID)
	if err != nil {
		return nil, err
	}

	log.Println(news.ID.Hex())
	return news, nil
}

// AllNewsResolverFunc graphql resolver for AllNews query
func (gqlres *LiveGraphQLResolvers) AllNewsResolverFunc(p graphql.ResolveParams) (interface{}, error) {

	log.Println("opening AllNewsResolverFunc() ")

	news, err := dbclient.DBClient.QueryAllNews()

	log.Println("leaving  AllNewsResolverFunc() found ", news)
	return news, err
}

// CreateNewsResolverFunc graphql resolver for Create news mutation query
func (gqlres *LiveGraphQLResolvers) CreateNewsResolverFunc(p graphql.ResolveParams) (interface{}, error) {

	log.Println("opening CreateNewsResolverFunc() ")

	news := model.News{
		Title:  p.Args["title"].(string),
		Teaser: p.Args["teaser"].(string),
		Body:   p.Args["body"].(string),
	}

	log.Println("CreateNewsResolverFunc: news from body ", news)

	result, err := dbclient.DBClient.CreateNews(news)

	log.Println("leaving  CreateNewsResolverFunc() found ", result)
	return result, err
}

// AddCommentToNewsResolverFunc add comment mutation resolver
func (gqlres *LiveGraphQLResolvers) AddCommentToNewsResolverFunc(p graphql.ResolveParams) (interface{}, error) {

	log.Println("opening AddCommentToNewsResolverFunc() ", p.Args)
	newsID, _ := p.Args["newsID"].(string)

	if len(newsID) == 0 {
		panic("AddCommentToNewsResolverFunc(): newsId is empty")
	}

	log.Println(" AddCommentToNewsResolverFunc() newsID ", newsID)
	username, _ := p.Args["username"].(string)

	body, _ := p.Args["body"].(string)

	comment := model.Comment{
		Username: username,
		Body:     body,
	}

	result, err := dbclient.DBClient.CreateComment(newsID, comment)

	log.Println("leaving AddCommentToNewsResolverFunc() result ", result)
	return result, err
}

func fetchNews(newsID string) (model.News, error) {

	log.Println("fetchNews newsID ", newsID)

	news, err := dbclient.DBClient.QueryNews(newsID)

	log.Println("fetchNews result ", news)

	return news, err
}

// TestGraphQLResolvers test implementations
// type TestGraphQLResolvers struct {
// }

// func (gqlres *TestGraphQLResolvers) NewsResolverFunc(p graphql.ResolveParams) (interface{}, error) {
// 	logrus.Infof("ENTER - resolve function for Account with params %v", p.Args)
// 	id, _ := p.Args["id"].(string)
// 	for _, account := range accounts {
// 		if account.ID == id {
// 			return account, nil
// 		}
// 	}

// 	return nil, fmt.Errorf("No account found matching ID %v", id)
// }

// func (gqlres *TestGraphQLResolvers) AllNewsResolverFunc(p graphql.ResolveParams) (interface{}, error) {
// 	logrus.Infof("ENTER - resolve function for AllAccounts with params %v", p.Args)
// 	return accounts, nil
// }
