package resolver

import (
	"context"
	"fmt"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/nomkhonwaan/myblog-server/pkg/tag"
)

type tagResolver struct {
	t *tag.Tag
}

func newTagResolver(t *tag.Tag) *tagResolver {
	return &tagResolver{t: t}
}

func (r *tagResolver) ID() graphql.ID {
	return relay.MarshalID("post", r.t.ID.Hex())
}

func (r *tagResolver) Name() string {
	return r.t.Name
}

func (r *tagResolver) Slug() string {
	return r.t.Slug
}

func (r *tagResolver) Link() string {
	return fmt.Sprintf("/tags/%s-%s", r.t.Slug, r.t.ID.Hex())
}

func (r *tagResolver) PublishedPosts(
	ctx context.Context,
	args struct {
		Offset  *int32
		Limit   *int32
		OrderBy *struct {
			Direction *string
			Field     *string
		}
	},
) (*[]*postResolver, error) {
	return nil, nil
}
