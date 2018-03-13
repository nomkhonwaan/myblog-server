package resolver

import (
	"context"
	"time"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/nomkhonwaan/myblog-server/pkg/post"
)

type PostResolver struct {
	*post.Post
}

func NewPostResolver(p *post.Post) *PostResolver {
	return &PostResolver{Post: p}
}

func (r *PostResolver) ID() graphql.ID {
	return relay.MarshalID("post", r.Post.ID.Hex())
}

func (r *PostResolver) Title() string {
	return r.Post.Title
}

func (r *PostResolver) Slug() string {
	return r.Post.Slug
}

func (r *PostResolver) Link() string {
	return r.Post.PublishedAt.Format("/2006/01/02/") + r.Post.Slug + "-" + r.Post.ID.Hex()
}

func (r *PostResolver) Status() string {
	return string(r.Post.Status)
}

func (r *PostResolver) HTML() *string {
	return &r.Post.HTML
}

func (r *PostResolver) Markdown() *string {
	return &r.Post.Markdown
}

func (r *PostResolver) CreatedAt() string {
	return r.Post.CreatedAt.Format(time.RFC3339)
}

func (r *PostResolver) UpdatedAt() *string {
	updatedAt := r.Post.UpdatedAt.Format(time.RFC3339)
	return &updatedAt
}

func (r *PostResolver) PublishedAt() *string {
	publisedAt := r.Post.PublishedAt.Format(time.RFC3339)
	return &publisedAt
}

func (r *Resolver) PublishedPost(ctx context.Context, args struct{ ID graphql.ID }) (*PostResolver, error) {
	repo := (ctx.Value("repositories").(map[string]interface{})["post"]).(*post.Repository)

	var id string

	err := relay.UnmarshalSpec(args.ID, &id)
	if err != nil {
		return nil, err
	}

	p, err := repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if p.Status != post.Published && p.PublishedAt.IsZero() {
		return nil, nil
	}

	return NewPostResolver(p), nil
}
