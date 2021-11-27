package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/hiroyky/go_graphql_server_sample/graph/generated"
	"github.com/hiroyky/go_graphql_server_sample/graph/model"
)

func (r *companyResolver) Departments(ctx context.Context, obj *model.Company) (*model.DepartmentPagination, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *companyResolver) Employees(ctx context.Context, obj *model.Company) (*model.EmployeePagination, error) {
	panic(fmt.Errorf("not implemented"))
}

// Company returns generated.CompanyResolver implementation.
func (r *Resolver) Company() generated.CompanyResolver { return &companyResolver{r} }

type companyResolver struct{ *Resolver }
