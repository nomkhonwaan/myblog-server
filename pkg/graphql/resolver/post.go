package resolver

import (
	"time"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/nomkhonwaan/myblog-server/pkg/post"
	"github.com/nomkhonwaan/myblog-server/pkg/tag"
)

// PostResolver is a Post's resolver which resolves all Post's fields
type PostResolver struct {
	*post.Post
	TagRepository tag.Repositorier
}

// NewPostResolver creates and returns a new PostResolver
func NewPostResolver(p *post.Post, tagRepository tag.Repositorier) *PostResolver {
	if p == nil {
		return nil
	}
	return &PostResolver{
		Post:          p,
		TagRepository: tagRepository,
	}
}

// ID returns a Post's ID
func (r *PostResolver) ID() graphql.ID {
	return graphql.ID(r.Post.ID.Hex())
}

// Title returns a Post's title
func (r *PostResolver) Title() string {
	return r.Post.Title
}

// Slug returns a Post's slug
func (r *PostResolver) Slug() string {
	return r.Post.Slug
}

// Link returns a Post's link
func (r *PostResolver) Link() *string {
	var link string

	if !r.Post.PublishedAt.IsZero() {
		link = r.Post.PublishedAt.Format("/2006/01/02") + r.Post.Slug + "-" + r.Post.Key()
	}

	return &link
}

// Status returns a Post's status
func (r *PostResolver) Status() string {
	return r.Post.Status
}

// HTML returns a Post's content in HTML format
func (r *PostResolver) HTML() *string {
	return &r.Post.HTML
}

// Markdown returns a Post's content in Markdown format
func (r *PostResolver) Markdown() *string {
	return &r.Post.Markdown
}

// Tags returns a list of Post's tag
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

// CreatedAt returns a date that this Post was created
func (r *PostResolver) CreatedAt() string {
	return r.Post.CreatedAt.Format(time.RFC3339)
}

// UpdatedAt returns a date that this Post was updated
func (r *PostResolver) UpdatedAt() *string {
	var updatedAt string

	if !r.Post.UpdatedAt.IsZero() {
		updatedAt = r.Post.UpdatedAt.Format(time.RFC3339)
	}

	return &updatedAt
}

// PublishedAt returns a date that this Post was published
func (r *PostResolver) PublishedAt() *string {
	var publishedAt string

	if !r.Post.PublishedAt.IsZero() {
		publishedAt = r.Post.PublishedAt.Format(time.RFC3339)
	}

	return &publishedAt
}
