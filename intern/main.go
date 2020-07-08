package main

import (
	"context"
	"encoding/json"
	"log"
	"strconv"
	"time"
	"io/ioutil"
	"net/http"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/graphql-go/graphql"
)


type response struct {
	OperationName string
	Variables  interface{};
	Query string
}




// Handler is King of Codecamp:N!
func Handler(context context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	res := response{}
	json.Unmarshal([]byte(request.Body), &res)

	query := res.Query
	log.Printf("query %s\n", query)

	// Schema
	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Description: "Hello Campers",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "Hello world", nil
			},
		},
		"waitLong": &graphql.Field{
			Type: graphql.String,
			Description: "Give me an Integer",
			Args: graphql.FieldConfigArgument{
				"wait": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				id, _ := params.Args["wait"].(int)
				time.Sleep(time.Duration(id)*time.Second)
				idAsString := strconv.Itoa(id)
				return ("Hello world, Wait: " + idAsString + " Seconds"), nil
			},
		},
		"request": &graphql.Field{
			Type: graphql.String,
			Description: "I will query for you!",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				resp, _ := http.Get("https://reqres.in/api/users")
				body, _ := ioutil.ReadAll(resp.Body)
				return string(body), nil
			},
		},

		"request2": &graphql.Field{
			Type: graphql.String,
			Description: "I will query for you!",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				resp, _ := http.Get("https://reqres.in/api/users")
				body, _ := ioutil.ReadAll(resp.Body)
				return string(body), nil
			},
		},
		"request3": &graphql.Field{
			Type: graphql.String,
			Description: "I will query for you!",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				resp, _ := http.Get("https://reqres.in/api/users")
				body, _ := ioutil.ReadAll(resp.Body)
				return string(body), nil
			},
		},
		"request4": &graphql.Field{
			Type: graphql.String,
			Description: "I will query for you!",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				resp, _ := http.Get("https://reqres.in/api/users")
				body, _ := ioutil.ReadAll(resp.Body)
				return string(body), nil
			},
		},
	}

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params)

	if len(r.Errors) > 0 {
		log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	}
	response, err := json.Marshal(r)

	if err != nil {
		log.Print("Could not decode body")
	}


	headers := map[string]string{
		"Access-Control-Allow-Origin": "*",
	}

	log.Printf("responseJSON %s\n", response)
	return events.APIGatewayProxyResponse{
		Body:       string(response),
		StatusCode: 200,
		Headers: headers,
	}, nil

}


func main() {
	lambda.Start(Handler)
}