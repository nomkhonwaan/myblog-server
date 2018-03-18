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

	p, err := repo.FindPublishedPostByID(id)
	if err != nil {
		return nil, err
	}

	return newPostResolver(p), nil
}

func (r *Resolver) PublishedPosts(
	ctx context.Context, args struct {
		Offset  *int32
		Limit   *int32
		OrderBy *struct {
			Direction *string
			Field     *string
		}
	},
) (*[]*postResolver, error) {
	repo := (ctx.Value("repositories").(map[string]interface{})["post"]).(*post.Repository)

	ps, err := repo.FindAllPublishedPosts(args.Offset, args.Limit, args.OrderBy)
	if err != nil {
		return nil, err
	}

	postResolvers := make([]*postResolver, len(ps))
	for i, p := range ps {
		postResolvers[i] = newPostResolver(p)
	}

	return &postResolvers, nil
}
