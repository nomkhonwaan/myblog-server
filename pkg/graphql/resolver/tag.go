package resolver

import (
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/nomkhonwaan/myblog-server/pkg/tag"
)

// TagResolver is a Tag's resolver which resolves all Tag's fields
type TagResolver struct {
	*tag.Tag
}

// ID returns a Tag's ID
func (r *TagResolver) ID() graphql.ID {
	return graphql.ID(r.Tag.ID.Hex())
}

// Name returns a Tag's name
func (r *TagResolver) Name() string {
	return r.Tag.Name
}

// Slug returns a Tag's slug
func (r *TagResolver) Slug() string {
	return r.Tag.Slug
}

// Link returns a Tag's link
func (r *TagResolver) Link() string {
	return "/tags/" + r.Tag.Slug + "-" + r.Tag.ID.Hex()
}

// PublishedPosts returns a list of published Posts that belong to this Tag
func (r *TagResolver) PublishedPosts() (*[]*PostResolver, error) {
	return nil, nil
}
