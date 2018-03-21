package resolver

import (
	"context"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/nomkhonwaan/myblog-server/pkg/post"
	"github.com/nomkhonwaan/myblog-server/pkg/tag"
)

// Resolver is a GraphQL's Root Resolver which provides root queries and mutations.
type Resolver struct {
	PostRepository post.Repositorier `inject:"pkg/post.Repositorier"`
	TagRepository  tag.Repositorier  `inject:"pkg/tag.Repositorier"`
}

// PublishedPost is an implemented function of GraphQL's queries which returns a published Post from its ID
func (r *Resolver) PublishedPost(_ context.Context, args struct{ ID graphql.ID }) (*PostResolver, error) {
	p, err := r.PostRepository.FindPublishedByID(string(args.ID))
	if err != nil {
		return nil, err
	}
	return NewPostResolver(p), nil
}

// Tag is an implemented function of GraphQL's queries which returns a Tag from its ID
func (r *Resolver) Tag(_ context.Context, args struct{ ID graphql.ID }) (*TagResolver, error) {
	t, err := r.TagRepository.FindByID(string(args.ID))
	if err != nil {
		return nil, err
	}
	return NewTagResolver(t), nil
}
