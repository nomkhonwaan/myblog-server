package resolver

import (
	"time"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/nomkhonwaan/myblog-server/pkg/post"
	"github.com/nomkhonwaan/myblog-server/pkg/tag"
)

type PostResolver struct {
	*post.Post
	TagRepository tag.Repositorier
}

func NewPostResolver(p *post.Post, tagRepository tag.Repositorier) *PostResolver {
	if p == nil {
		return nil
	}
	return &PostResolver{
		Post:          p,
		TagRepository: tagRepository,
	}
}

func (r *PostResolver) ID() graphql.ID {
	return graphql.ID(r.Post.ID.Hex())
}

func (r *PostResolver) Title() string {
	return r.Post.Title
}

func (r *PostResolver) Slug() string {
	return r.Post.Slug
}

func (r *PostResolver) Link() *string {
	var link string

	if !r.Post.PublishedAt.IsZero() {
		link = r.Post.PublishedAt.Format("/2006/01/02") + r.Post.Slug + "-" + r.Post.Key()
	}

	return &link
}

func (r *PostResolver) Status() string {
	return r.Post.Status
}

func (r *PostResolver) HTML() *string {
	return &r.Post.HTML
}

func (r *PostResolver) Markdown() *string {
	return &r.Post.Markdown
}

func (r *PostResolver) Tags() (*[]*TagResolver, error) {
	var err error

	for i, t := range r.Post.Tags {
		r.Post.Tags[i], err = r.TagRepository.FindByID(t.ID.Hex())
		if err != nil {
			return nil, err
		}
	}
	return NewTagsResolver(r.Post.Tags), nil
}

func (r *PostResolver) CreatedAt() string {
	return r.Post.CreatedAt.Format(time.RFC3339)
}

func (r *PostResolver) UpdatedAt() *string {
	var updatedAt string

	if !r.Post.UpdatedAt.IsZero() {
		updatedAt = r.Post.UpdatedAt.Format(time.RFC3339)
	}

	return &updatedAt
}

func (r *PostResolver) PublishedAt() *string {
	var publishedAt string

	if !r.Post.PublishedAt.IsZero() {
		publishedAt = r.Post.PublishedAt.Format(time.RFC3339)
	}

	return &publishedAt
}
