//go:generate go run github.com/99designs/gqlgen
package graph

import "github.com/dandelion293949/todo-list/graph/domains"
// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
  TodoRepository domains.TodoRepository
}
