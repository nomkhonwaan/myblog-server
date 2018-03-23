package resolver

import (
	"context"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/nomkhonwaan/myblog-server/pkg/post"
	"github.com/nomkhonwaan/myblog-server/pkg/tag"
	"github.com/sirupsen/logrus"
)

// OrderByDirection is a represent type of string for sorting direction
type OrderByDirection string

const (
	// ASC is a default sorting option which is empty string in MongoDB
	ASC = OrderByDirection("")
	// DESC is a sorting which will start with dashed (-) on any field that need to sort by descending
	DESC = OrderByDirection("-")
)

// Resolver is a GraphQL's Root Resolver which provides root queries and mutations.
type Resolver struct {
	PostRepository post.Repositorier `inject:"pkg/post.Repositorier"`
	TagRepository  tag.Repositorier  `inject:"pkg/tag.Repositorier"`
}

// PublishedPost is an implemented function of GraphQL's queries which returns a published Post from its ID
func (r *Resolver) PublishedPost(_ context.Context, args struct{ ID graphql.ID }) (*PostResolver, error) {
	p, err := r.PostRepository.FindPublishedByID(string(args.ID))
	if err != nil {
		return nil, err
	}
	return NewPostResolver(p, r.TagRepository), nil
}

// Tag is an implemented function of GraphQL's queries which returns a Tag from its ID
func (r *Resolver) Tag(_ context.Context, args struct{ ID graphql.ID }) (*TagResolver, error) {
	t, err := r.TagRepository.FindByID(string(args.ID))
	if err != nil {
		return nil, err
	}
	return NewTagResolver(t), nil
}

// Tags is an implemented function of GraphQL's queries which returns a list of Tags
func (r *Resolver) Tags(
	_ context.Context,
	args struct {
		Offset, Limit *int32
		OrderBy       *struct {
			Field     *string
			Direction *string
		}
	},
) (*[]*TagResolver, error) {
	var orderBy struct {
		Field     string
		Direction string
	}
	if args.OrderBy != nil {
		orderBy.Field = withDefaultString(args.OrderBy.Field, "slug")
		orderBy.Direction = withDefaultString(args.OrderBy.Direction, string(ASC))
	}

	ts, err := r.TagRepository.FindAll(
		int(withDefaultInt32(args.Offset, 0)),
		int(withDefaultInt32(args.Limit, 5)),
		orderBy,
	)
	if err != nil {
		return nil, err
	}
	return NewTagsResolver(ts), nil
}

// withDefaultInt returns a default number of int if the given number is equal zero or nil
func withDefaultInt(num *int, dnum int) int {
	if num == nil {
		return dnum
	}
	if *num == 0 {
		return dnum
	}
	return *num
}

// withDefaultInt32 returns a default number of int32 if the given number is equal zero or nil
func withDefaultInt32(num *int32, dnum int32) int32 {
	if num == nil {
		return dnum
	}
	if *num == 0 {
		return dnum
	}
	return *num
}

// withDefaultString returns a default string if the given string is empty or nil
func withDefaultString(str *string, dstr string) string {
	logrus.Info(str)
	if str == nil {
		return dstr
	}
	if *str == "" {
		return dstr
	}
	return *str
}
