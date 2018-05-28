package resolver

import (
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/nomkhonwaan/myblog-server/pkg/tag"
)

type TagResolver struct {
	*tag.Tag
}

func NewTagResolver(t *tag.Tag) *TagResolver {
	if t == nil {
		return nil
	}
	return &TagResolver{Tag: t}
}

func NewTagsResolver(ts []*tag.Tag) *[]*TagResolver {
	if ts == nil {
		return nil
	}
	tagsResolver := make([]*TagResolver, len(ts))
	for i, t := range ts {
		tagsResolver[i] = NewTagResolver(t)
	}
	return &tagsResolver
}

func (r *TagResolver) ID() graphql.ID {
	return graphql.ID(r.Tag.ID.Hex())
}

func (r *TagResolver) Name() string {
	return r.Tag.Name
}

func (r *TagResolver) Slug() string {
	return r.Tag.Slug
}

func (r *TagResolver) Link() string {
	return "/tags/" + r.Tag.Slug + "-" + r.Tag.ID.Hex()
}

func (r *TagResolver) PublishedPosts() (*[]*PostResolver, error) {
	return nil, nil
}
