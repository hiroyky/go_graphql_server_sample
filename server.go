package main

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/hiroyky/go_graphql_server_sample/dataloader"
	"github.com/hiroyky/go_graphql_server_sample/graph/model"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/hiroyky/go_graphql_server_sample/graph"
	"github.com/hiroyky/go_graphql_server_sample/graph/generated"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	companyLoader := dataloader.NewCompanyLoader(dataloader.CompanyLoaderConfig{
		MaxBatch: 100,
		Wait:     2 * time.Millisecond,
		Fetch: func(keys []string) ([]*model.Company, []error) {
			companies := make([]*model.Company, len(keys))
			errors := make([]error, len(keys))

			// 取得処理を書く

			// 引数のkeysに対応する順番の配列で返す。
			return companies, errors
		},
	})

	c := generated.Config{Resolvers: &graph.Resolver{CompanyLoader: companyLoader}}
	c.Complexity.Mutation.CreateCompany = func(childComplexity int, input model.CreateCompanyInput) int {
		return 5
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(c))
	// エラー処理を書く
	srv.SetErrorPresenter(func(ctx context.Context, e error) *gqlerror.Error {
		err := graphql.DefaultErrorPresenter(ctx, e)
		err.Message = e.Error()
		err.Extensions = map[string]interface{}{
			"key1": "value1",
			"key2": "value2",
		}
		return err
	})

	srv.Use(extension.FixedComplexityLimit(10)) // 重さが10を超えたらエラーにする

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
