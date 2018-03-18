package resolver

import (
	"context"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/nomkhonwaan/myblog-server/pkg/post"
)

// Resolver is a GraphQL's Root Resolver which provides root queries and mutations.
type Resolver struct {
	PostRepository post.Repositorier `inject:"pkg/post.Repositorier"`
}

// PublishedPost is an implemented of GraphQL's query
func (r *Resolver) PublishedPost(_ context.Context, args struct{ ID graphql.ID }) (*PostResolver, error) {
	p, err := r.PostRepository.FindPublishedByID(string(args.ID))
	if err != nil {
		return nil, err
	}
	if p == nil {
		return nil, nil
	}

	return &PostResolver{Post: p}, nil
}
