package resolver

import (
	"time"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/nomkhonwaan/myblog-server/pkg/post"
)

type postResolver struct {
	p *post.Post
}

func newPostResolver(p *post.Post) *postResolver {
	return &postResolver{p: p}
}

func (r *postResolver) ID() graphql.ID {
	return relay.MarshalID("post", r.p.ID.Hex())
}

func (r *postResolver) Title() string {
	return r.p.Title
}

func (r *postResolver) Slug() string {
	return r.p.Slug
}

func (r *postResolver) Link() string {
	return r.p.PublishedAt.Format("/2006/01/02/") + r.p.Slug + "-" + r.p.ID.Hex()
}

func (r *postResolver) Status() string {
	return string(r.p.Status)
}

func (r *postResolver) HTML() *string {
	return &r.p.HTML
}

func (r *postResolver) Markdown() *string {
	return &r.p.Markdown
}

func (r *postResolver) Tags() *[]*tagResolver {
	ts := make([]*tagResolver, len(r.p.Tags))
	for i, t := range r.p.Tags {
		ts[i] = newTagResolver(&t)
	}
	return &ts
}
func (r *postResolver) CreatedAt() string {
	return r.p.CreatedAt.Format(time.RFC3339)
}

func (r *postResolver) UpdatedAt() *string {
	updatedAt := r.p.UpdatedAt.Format(time.RFC3339)

	return &updatedAt
}

func (r *postResolver) PublishedAt() *string {
	publisedAt := r.p.PublishedAt.Format(time.RFC3339)

	return &publisedAt
}
