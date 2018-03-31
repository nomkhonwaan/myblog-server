package tag_test

import (
	"context"
	"errors"
	"reflect"

	dld "github.com/nicksrandall/dataloader"
	. "github.com/nomkhonwaan/myblog-server/pkg/dataloader/mock"
	. "github.com/nomkhonwaan/myblog-server/pkg/mongodb/mock"
	. "github.com/nomkhonwaan/myblog-server/pkg/tag"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gopkg.in/mgo.v2/bson"
)

type mockQuery struct {
	*MockQuery
	result interface{}
}

func (q *mockQuery) All(result interface{}) error {
	resultv := reflect.ValueOf(result)
	slicev := resultv.Elem()
	for _, t := range q.result.([]*Tag) {
		slicev = reflect.Append(slicev, reflect.ValueOf(t))
	}
	resultv.Elem().Set(slicev.Slice(0, len(q.result.([]*Tag))))
	return nil
}

type mockErrorQuery struct {
	*MockQuery
	err error
}

func (q *mockErrorQuery) All(result interface{}) error {
	return q.err
}

var _ = Describe("Tag", func() {
	Describe("Tag.Key", func() {
		It("should return a hexadecimal string of ObjectID correctly", func() {
			t := &Tag{ID: bson.ObjectIdHex("5ab1cb7e47774b19578e6894")}
			Expect(t.Key()).To(Equal("5ab1cb7e47774b19578e6894"))
		})
	})

	Describe("Tags.Keys", func() {
		It("should return a list of hexadecimal string of the list of ObjectIDs correctly", func() {
			ts := Tags{
				&Tag{
					ID: bson.ObjectIdHex("5794a237b7655ba4eef7ad68"),
				},
				&Tag{
					ID: bson.ObjectIdHex("5794a237b7655ba4eef7ad6e"),
				},
			}
			Expect(ts.Keys()).To(Equal([]string{"5794a237b7655ba4eef7ad68", "5794a237b7655ba4eef7ad6e"}))
		})
	})

	Describe("NewPlaceholder", func() {
		It("should return a new placeholder which is an empty Tag object correctly", func() {
			placeholder := NewPlaceholder()
			Expect(reflect.TypeOf(placeholder).String()).To(Equal("*tag.Tag"))
			Expect(string(placeholder.(*Tag).ID)).To(Equal(""))
			Expect(placeholder.(*Tag).Name).To(Equal(""))
			Expect(placeholder.(*Tag).Slug).To(Equal(""))
		})
	})

	Describe("Repository.FindByID", func() {
		Context("find with existing Tag's ID", func() {
			It("should return a Tag from its ID", func() {
				loader := NewMockInterface(ctrl)
				id := bson.NewObjectId()
				repo := Repository{Loader: loader}
				loader.
					EXPECT().
					Load(context.TODO(), dld.StringKey(id.Hex())).
					Return(dld.Thunk(func() (interface{}, error) {
						return &Tag{
							ID: id,
						}, nil
					}))

				t, err := repo.FindByID(id.Hex())
				Expect(err).To(BeNil())
				Expect(t.ID).To(Equal(id))
			})
		})

		Context("find with non-existing Tag's ID", func() {
			It("should return a nil Tag's pointer", func() {
				loader := NewMockInterface(ctrl)
				id := bson.NewObjectId()
				repo := Repository{Loader: loader}
				loader.
					EXPECT().
					Load(context.TODO(), dld.StringKey(id.Hex())).
					Return(dld.Thunk(func() (interface{}, error) {
						return nil, nil
					}))

				t, err := repo.FindByID(id.Hex())
				Expect(err).To(BeNil())
				Expect(t).To(BeNil())
			})
		})

		Context("find with any ID but an error has occurred", func() {
			It("should return an error", func() {
				loader := NewMockInterface(ctrl)
				id := bson.NewObjectId()
				repo := Repository{Loader: loader}
				loader.
					EXPECT().
					Load(context.TODO(), dld.StringKey(id.Hex())).
					Return(dld.Thunk(func() (interface{}, error) {
						return nil, errors.New("something went wrong")
					}))

				t, err := repo.FindByID(id.Hex())
				Expect(err).To(Not(BeNil()))
				Expect(t).To(BeNil())
			})
		})
	})

	Describe("repository.FindAll", func() {
		Context("find first 10 Tags", func() {
			It("should return a list of first 10 Tags", func() {
				db := NewMockDatabase(ctrl)
				loader := NewMockInterface(ctrl)
				mockTags := make([]*Tag, 10)
				for i := range mockTags {
					mockTags[i] = &Tag{
						ID: bson.NewObjectId(),
					}
				}
				repo := Repository{
					Database: db,
					Loader:   loader,
				}
				q := &mockQuery{NewMockQuery(ctrl), mockTags}
				q.EXPECT().Select(bson.M{"_id": 1}).Return(q)
				q.EXPECT().Skip(0).Return(q)
				q.EXPECT().Limit(10).Return(q)
				c := NewMockCollection(ctrl)
				c.EXPECT().Find(nil).Return(q)
				db.EXPECT().C("tags").Return(c)
				loader.
					EXPECT().
					LoadMany(context.TODO(), dld.NewKeysFromStrings(Tags(mockTags).Keys())).
					Return(dld.ThunkMany(func() ([]interface{}, []error) {
						result := make([]interface{}, len(mockTags))
						for i, t := range mockTags {
							result[i] = t
						}
						return result, nil
					}))

				ts, err := repo.FindAll(0, 10, struct {
					Field     string
					Direction string
				}{"slug", ""})
				Expect(err).To(BeNil())
				Expect(len(ts)).To(Equal(10))
				for i, t := range ts {
					Expect(t.ID).To(Equal(mockTags[i].ID))
				}
			})
		})

		Context("find next 10 Tags", func() {
			It("should return a list of next 10 Tags", func() {
				db := NewMockDatabase(ctrl)
				loader := NewMockInterface(ctrl)
				mockTags := make([]*Tag, 10)
				for i := range mockTags {
					mockTags[i] = &Tag{
						ID: bson.NewObjectId(),
					}
				}
				repo := Repository{
					Database: db,
					Loader:   loader,
				}
				q := &mockQuery{NewMockQuery(ctrl), mockTags}
				q.EXPECT().Select(bson.M{"_id": 1}).Return(q)
				q.EXPECT().Skip(10).Return(q)
				q.EXPECT().Limit(10).Return(q)
				c := NewMockCollection(ctrl)
				c.EXPECT().Find(nil).Return(q)
				db.EXPECT().C("tags").Return(c)
				loader.
					EXPECT().
					LoadMany(context.TODO(), dld.NewKeysFromStrings(Tags(mockTags).Keys())).
					Return(dld.ThunkMany(func() ([]interface{}, []error) {
						result := make([]interface{}, len(mockTags))
						for i, t := range mockTags {
							result[i] = t
						}
						return result, nil
					}))

				ts, err := repo.FindAll(10, 10, struct {
					Field     string
					Direction string
				}{"slug", ""})
				Expect(err).To(BeNil())
				Expect(len(ts)).To(Equal(10))
				Expect(err).To(BeNil())
				for i, t := range ts {
					Expect(t.ID).To(Equal(mockTags[i].ID))
				}
			})
		})

		Context("find more 10 Tags but an error has occurred", func() {
			It("should return an error", func() {
				db := NewMockDatabase(ctrl)
				loader := NewMockInterface(ctrl)
				mockTags := make([]*Tag, 10)
				for i := range mockTags {
					mockTags[i] = &Tag{
						ID: bson.NewObjectId(),
					}
				}
				repo := Repository{
					Database: db,
					Loader:   loader,
				}
				q := &mockErrorQuery{NewMockQuery(ctrl), errors.New("something went wrong")}
				q.EXPECT().Select(bson.M{"_id": 1}).Return(q)
				q.EXPECT().Skip(10).Return(q)
				q.EXPECT().Limit(10).Return(q)
				c := NewMockCollection(ctrl)
				c.EXPECT().Find(nil).Return(q)
				db.EXPECT().C("tags").Return(c)
				loader.
					EXPECT().
					LoadMany(context.TODO(), dld.NewKeysFromStrings(Tags(mockTags).Keys())).
					Return(dld.ThunkMany(func() ([]interface{}, []error) {
						return nil, nil
					}))

				ts, err := repo.FindAll(10, 10, struct {
					Field     string
					Direction string
				}{"slug", ""})
				Expect(err).To(Not(BeNil()))
				Expect(ts).To(BeNil())
			})
		})
	})
})
