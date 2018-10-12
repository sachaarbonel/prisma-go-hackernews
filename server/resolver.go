//go:generate gorunpkg github.com/99designs/gqlgen

package main

import (
	context "context"

	"github.com/Sach97/prisma-go-hackernews/prisma-client"
)

type Resolver struct {
	Prisma *prisma.Client
}

func (r *Resolver) Link() LinkResolver {
	return &linkResolver{r}
}
func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) User() UserResolver {
	return &userResolver{r}
}
func (r *Resolver) Vote() VoteResolver {
	return &voteResolver{r}
}

type linkResolver struct{ *Resolver }

func (r *linkResolver) PostedBy(ctx context.Context, obj *prisma.Link) (*prisma.User, error) {
	PostedBy, err := r.Prisma.Link(prisma.LinkWhereUniqueInput{
		ID: &obj.ID,
	}).PostedBy().Exec(ctx)
	return PostedBy, err
}
func (r *linkResolver) AllVotes(ctx context.Context, obj *prisma.Link) ([]prisma.Vote, error) {
	AllVotes, err := r.Prisma.Link(prisma.LinkWhereUniqueInput{
		ID: &obj.ID,
	}).AllVotes(nil).Exec(ctx)
	return AllVotes, err
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateUser(ctx context.Context, name string, email string, password string) (prisma.User, error) {
	user, err := r.Prisma.CreateUser(prisma.UserCreateInput{
		Name:     name,
		Email:    email,
		Password: password,
	}).Exec(ctx)
	return *user, err
}

func (r *mutationResolver) CreateLink(ctx context.Context, url string, description string) (prisma.Link, error) {

	//TODO : middleware to get userid
	userid := "cjmvspvhf00020906pgn00izs" //hardcoded - change this value according to the result of the mutation createUser
	link, err := r.Prisma.CreateLink(prisma.LinkCreateInput{
		Description: description,
		Url:         url,
		PostedBy: &prisma.UserCreateOneWithoutLinksInput{
			Connect: &prisma.UserWhereUniqueInput{
				ID: &userid,
			},
		},
	}).Exec(ctx)
	return *link, err
}

func (r *mutationResolver) UpVote(ctx context.Context, linkId string) (prisma.Vote, error) {
	//TODO : middleware to get userid
	userid := "cjmvspvhf00020906pgn00izs" //hardcoded - change this value according to the result of the mutation createUser
	vote, err := r.Prisma.CreateVote(prisma.VoteCreateInput{
		Link: prisma.LinkCreateOneWithoutAllVotesInput{
			Connect: &prisma.LinkWhereUniqueInput{
				ID: &linkId,
			},
		},
		VotedBy: prisma.UserCreateOneWithoutVotesInput{
			Connect: &prisma.UserWhereUniqueInput{
				ID: &userid,
			},
		},
	}).Exec(ctx)

	return *vote, err
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Links(ctx context.Context) ([]prisma.Link, error) {
	return r.Prisma.Links(nil).Exec(ctx)
}
func (r *queryResolver) Users(ctx context.Context) ([]prisma.User, error) {
	return r.Prisma.Users(nil).Exec(ctx)
}

func (r *queryResolver) Me(ctx context.Context, id string) (prisma.User, error) {
	panic("not implemented")
}

type userResolver struct{ *Resolver }

func (r *userResolver) Links(ctx context.Context, obj *prisma.User) ([]prisma.Link, error) {
	links, err := r.Prisma.User(prisma.UserWhereUniqueInput{
		ID: &obj.ID,
	}).Links(nil).Exec(ctx)

	return links, err
}
func (r *userResolver) Votes(ctx context.Context, obj *prisma.User) ([]prisma.Vote, error) {
	votes, err := r.Prisma.User(prisma.UserWhereUniqueInput{
		ID: &obj.ID,
	}).Votes(nil).Exec(ctx)

	return votes, err
}

type voteResolver struct{ *Resolver }

func (r *voteResolver) Link(ctx context.Context, obj *prisma.Vote) (prisma.Link, error) {
	link, err := r.Prisma.Vote(prisma.VoteWhereUniqueInput{
		ID: &obj.ID,
	}).Link().Exec(ctx)
	return *link, err
}
func (r *voteResolver) VotedBy(ctx context.Context, obj *prisma.Vote) (prisma.User, error) {
	votedBy, err := r.Prisma.Vote(prisma.VoteWhereUniqueInput{
		ID: &obj.ID,
	}).VotedBy().Exec(ctx)
	return *votedBy, err
}
