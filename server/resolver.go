//go:generate gorunpkg github.com/99designs/gqlgen

package main

import (
	context "context"
	prisma "prisma-go-todo/prisma-client"
)

type Resolver struct {
	Prisma *prisma.Client
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateTodo(ctx context.Context, text string, done *bool) (prisma.Todo, error) {
	return r.Prisma.CreateTodo(&prisma.TodoCreateInput{
		Text: &text,
		Done: done,
	}).Exec()
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Todo(ctx context.Context, id string) (*prisma.Todo, error) {
	todo, err := r.Prisma.Todo(&prisma.TodoWhereUniqueInput{
		ID: &id,
	}).Exec()
	return &todo, err
}

// func (r *queryResolver) LastTodo(ctx context.Context) (*prisma.Todo, error) {
// 	last := int32(1)
// 	return r.Prisma.Todoes(&prisma.TodoesParams{
// 		Last: &last,
// 	}).Exec()
// 	//return &todo, err
// }

func (r *queryResolver) Todos(ctx context.Context) ([]prisma.Todo, error) {
	done := false
	return r.Prisma.Todoes(&prisma.TodoesParams{
		Where: &prisma.TodoWhereInput{Done: &done},
	}).Exec()
}
