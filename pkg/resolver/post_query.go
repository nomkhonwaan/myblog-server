package resolver

import (
	"context"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/nomkhonwaan/myblog-server/pkg/post"
)

func (r *Resolver) PublishedPost(ctx context.Context, args struct{ ID graphql.ID }) (*postResolver, error) {
	var id string
	err := relay.UnmarshalSpec(args.ID, &id)
	if err != nil {
		return nil, err
	}

	repo := (ctx.Value("repositories").(map[string]interface{})["post"]).(*post.Repository)
	p, err := repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if p.Status != post.Published && p.PublishedAt.IsZero() {
		return nil, nil
	}

	return newPostResolver(p), nil
}

// func (r *Resolver) PublishedPosts(ctx context.Context, args struct {
// 	Before  *string
// 	After   *string
// 	First   *int32
// 	Last    *int32
// 	OrderBy *struct {
// 		Direction *string
// 		Field     *string
// 	}
// }) (*postConnectionResolver, error) {
// 	repo := (ctx.Value("repositories").(map[string]interface{})["post"]).(*post.Repository)
// 	ps, err := repo.FindAll(args.OrderBy)
// 	if err != nil {
// 		return nil, err
// 	}

// 	publishedPosts := make([]*post.Post, len(ps))
// 	for i, p := range ps {
// 		if p.Status == post.Published || !p.PublishedAt.IsZero() {
// 			publishedPosts[i] = p
// 		}
// 	}

// 	return newPostConnectionResolver(
// 		publishedPosts,
// 		connectionArguments{args.Before, args.After, args.First, args.Last},
// 		arraySliceMetaInfo{0, len(publishedPosts)},
// 	), nil
// }
