package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/hiroyky/go_graphql_server_sample/graph/generated"
	"github.com/hiroyky/go_graphql_server_sample/graph/model"
)

func (r *employeeResolver) Department(ctx context.Context, obj *model.Employee) (*model.Department, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *employeeResolver) Company(ctx context.Context, obj *model.Employee) (*model.Company, error) {
	return r.CompanyLoader.Load(obj.CompanyID)
}

// Employee returns generated.EmployeeResolver implementation.
func (r *Resolver) Employee() generated.EmployeeResolver { return &employeeResolver{r} }

type employeeResolver struct{ *Resolver }
