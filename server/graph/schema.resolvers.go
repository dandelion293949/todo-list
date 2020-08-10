package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/dandelion293949/todo-list/graph/domains"
	generated1 "github.com/dandelion293949/todo-list/graph/generated"
	model1 "github.com/dandelion293949/todo-list/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model1.CreateTodoInput) (*domains.Todo, error) {
	return r.TodoRepository.Create(ctx, &domains.Todo{
		Text: input.Text,
	})
}

func (r *mutationResolver) UpdateTodo(ctx context.Context, input model1.UpdateTodoInput) (*domains.Todo, error) {
	todo, err := r.TodoRepository.Get(ctx, domains.TodoID(input.ID))
	if err != nil {
		return nil, err
	}

	if input.Text != nil {
		todo.Text = *input.Text
	}

	if input.Done != nil {
		todo.Done = *input.Done
	}

	return r.TodoRepository.Update(ctx, todo)
}

func (r *mutationResolver) DeleteTodo(ctx context.Context, input model1.DeleteTodoInput) (*domains.Todo, error) {
	todo, err := r.TodoRepository.Get(ctx, domains.TodoID(input.ID))
	if err != nil {
		return nil, err
	}

	return r.TodoRepository.Delete(ctx, todo)
}

func (r *queryResolver) Todos(ctx context.Context) ([]*domains.Todo, error) {
	return r.TodoRepository.GetAll(ctx)
}

func (r *todoResolver) ID(ctx context.Context, obj *domains.Todo) (string, error) {
	return string(obj.ID), nil
}

// Mutation returns generated1.MutationResolver implementation.
func (r *Resolver) Mutation() generated1.MutationResolver { return &mutationResolver{r} }

// Query returns generated1.QueryResolver implementation.
func (r *Resolver) Query() generated1.QueryResolver { return &queryResolver{r} }

// Todo returns generated1.TodoResolver implementation.
func (r *Resolver) Todo() generated1.TodoResolver { return &todoResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
