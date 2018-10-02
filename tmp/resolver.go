//go:generate gorunpkg github.com/99designs/gqlgen

package tmp

import (
	context "context"
	prisma "prisma-go-hackernews/prisma-client"
	main "prisma-go-hackernews/server"
)

type Resolver struct{}

func (r *Resolver) Link() main.LinkResolver {
	return &linkResolver{r}
}
func (r *Resolver) Mutation() main.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() main.QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) User() main.UserResolver {
	return &userResolver{r}
}
func (r *Resolver) Vote() main.VoteResolver {
	return &voteResolver{r}
}

type linkResolver struct{ *Resolver }

func (r *linkResolver) PostedBy(ctx context.Context, obj *prisma.Link) (*prisma.User, error) {
	panic("not implemented")
}
func (r *linkResolver) AllVotes(ctx context.Context, obj *prisma.Link) ([]prisma.Vote, error) {
	panic("not implemented")
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateUser(ctx context.Context, name string, email string, password string) (prisma.User, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateLink(ctx context.Context, url string, description string) (prisma.Link, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Links(ctx context.Context) ([]prisma.Link, error) {
	panic("not implemented")
}
func (r *queryResolver) Users(ctx context.Context) ([]prisma.User, error) {
	panic("not implemented")
}
func (r *queryResolver) Me(ctx context.Context, id string) (prisma.User, error) {
	panic("not implemented")
}

type userResolver struct{ *Resolver }

func (r *userResolver) Links(ctx context.Context, obj *prisma.User) ([]prisma.Link, error) {
	panic("not implemented")
}
func (r *userResolver) Votes(ctx context.Context, obj *prisma.User) ([]prisma.Vote, error) {
	panic("not implemented")
}

type voteResolver struct{ *Resolver }

func (r *voteResolver) Link(ctx context.Context, obj *prisma.Vote) (prisma.Link, error) {
	panic("not implemented")
}
func (r *voteResolver) VotedBy(ctx context.Context, obj *prisma.Vote) (prisma.User, error) {
	panic("not implemented")
}
