package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/hiroyky/go_graphql_server_sample/graph/generated"
	"github.com/hiroyky/go_graphql_server_sample/graph/model"
)

func (r *queryResolver) Company(ctx context.Context, id string) (*model.Company, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Companies(ctx context.Context, limit int, offset *int) (*model.CompanyPagination, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Department(ctx context.Context, id string) (*model.Department, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Departments(ctx context.Context, limit int, offset *int) (*model.DepartmentPagination, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Employee(ctx context.Context, id string) (*model.Employee, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Employees(ctx context.Context, limit int, offset *int, email *string, gender *model.Gender, isManager *bool, hasDependent *bool) (*model.EmployeePagination, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
