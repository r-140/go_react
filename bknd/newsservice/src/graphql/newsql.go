package graphql

import (
	"log"
	internalModel "model"

	"github.com/graphql-go/graphql"
)

// Schema graphql schema definition for http endpoint
var Schema graphql.Schema
var schemaInitialized = false

var newsList []internalModel.News

func init() {
	newsList = make([]internalModel.News, 0)
}

// InitQL initializes graphql schema
func InitQL(resolvers GraphQLResolvers) {
	if schemaInitialized {
		return
	}
	// ----------- Start declare types ------------------

	// Comment
	var commentType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Comment",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"username": &graphql.Field{
				Type: graphql.String,
			},
			"body": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

	// News
	var newsType = graphql.NewObject(graphql.ObjectConfig{
		Name: "News",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"teaser": &graphql.Field{
				Type: graphql.String,
			},
			"body": &graphql.Field{
				Type: graphql.String,
			},
			"comments": &graphql.Field{
				Type: graphql.NewList(commentType),
				Args: graphql.FieldConfigArgument{

					"username": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					log.Println("ENTER - resolve function for comments with params %v", p.Args)
					news := p.Source.(internalModel.News)

					if len(p.Args) == 0 {
						return news.Comments, nil
					}

					response := make([]internalModel.Comment, 0)
					for _, item := range news.Comments {
						if item.Username == p.Args["username"] {
							response = append(response, item)
						}
					}
					return response, nil
				},
			},
		},
	})

	var newsMutationType = graphql.NewObject(graphql.ObjectConfig{
		Name: "CreateNewsMutation",
		Fields: graphql.Fields{
			"CreateNewsMutation": &graphql.Field{
				Type:        newsType,
				Description: "Add News",
				Args: graphql.FieldConfigArgument{
					"title": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"body": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"teaser": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: resolvers.CreateNewsResolverFunc,
			},
		},
	})

	// var commentMutationType = graphql.NewObject(graphql.ObjectConfig{
	// 	Name: "CommentMutation",
	// 	Fields: graphql.Fields{
	// 		"create": &graphql.Field{
	// 			Type:        commentType,
	// 			Description: "Add Comment to News",
	// 			Args: graphql.FieldConfigArgument{
	// 				"newsID": &graphql.ArgumentConfig{
	// 					Type: graphql.NewNonNull(graphql.String),
	// 				},
	// 				"body": &graphql.ArgumentConfig{
	// 					Type: graphql.NewNonNull(graphql.String),
	// 				},
	// 				"username": &graphql.ArgumentConfig{
	// 					Type: graphql.NewNonNull(graphql.String),
	// 				},
	// 			},
	// 			Resolve: resolvers.AddCommentToNewsResolverFunc,
	// 		},
	// 	},
	// })

	// var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	// 	Name: "RootMutation",
	// 	Fields: graphql.Fields{
	// 		"CreateNewsMutation":    newsMutationType,
	// 		"CreateCommentMutation": commentMutationType,
	// 	},
	// })

	// Schema
	fields := graphql.Fields{
		"News": &graphql.Field{
			Type: graphql.Type(newsType),
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"title": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: resolvers.NewsResolverFunc,
		},
		"AllNews": &graphql.Field{
			Type: graphql.NewList(newsType),

			Resolve: resolvers.AllNewsResolverFunc,
		},
	}

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{
		Query:    graphql.NewObject(rootQuery),
		Mutation: newsMutationType,
	}
	var err error
	Schema, err = graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}
	log.Println("Successfully initialized GraphQL")
	schemaInitialized = true
}
