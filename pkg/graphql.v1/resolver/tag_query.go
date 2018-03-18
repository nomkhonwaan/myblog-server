package resolver

import (
	"context"

	"github.com/graph-gophers/graphql-go/relay"
	"github.com/nomkhonwaan/myblog-server/pkg/tag"

	graphql "github.com/graph-gophers/graphql-go"
)

func (r *Resolver) Tag(ctx context.Context, args struct{ ID graphql.ID }) (*tagResolver, error) {
	var id string
	err := relay.UnmarshalSpec(args.ID, &id)
	if err != nil {
		return nil, err
	}

	repo := (ctx.Value("repositories").(map[string]interface{})["tag"]).(*tag.Repository)

	t, err := repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return newTagResolver(t), nil
}

func (r *Resolver) Tags(
	ctx context.Context,
	args struct {
		Offset  *int32
		Limit   *int32
		OrderBy *struct {
			Direction *string
			Field     *string
		}
	},
) (*[]*tagResolver, error) {
	repo := (ctx.Value("repositories").(map[string]interface{})["tag"]).(*tag.Repository)

	ts, err := repo.FindAll(args.Offset, args.Limit, args.OrderBy)
	if err != nil {
		return nil, err
	}

	tagResolvers := make([]*tagResolver, len(ts))
	for i, t := range ts {
		tagResolvers[i] = newTagResolver(t)
	}

	return &tagResolvers, nil
}
