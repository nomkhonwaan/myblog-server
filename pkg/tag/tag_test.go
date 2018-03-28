package tag_test

import (
	"context"
	"reflect"

	dld "github.com/nicksrandall/dataloader"
	"github.com/nomkhonwaan/myblog-server/pkg/dataloader/mock"
	. "github.com/nomkhonwaan/myblog-server/pkg/tag"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gopkg.in/mgo.v2/bson"
)

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
				loader := dataloader_mock.NewMockInterface(ctrl)
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
				loader := dataloader_mock.NewMockInterface(ctrl)
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
	})

	Describe("repository.FindAll", func() {
		Context("find first 10 Tags", func() {
			It("should return a list of first 10 Tags order by alphabet", func() {
				// db := mongodb_mock.NewMockDatabase(ctrl)
				// loader := dataloader_mock.NewMockInterface(ctrl)
				// repo := Repository{
				// 	Database: db,
				// 	Loader:   loader,
				// }
				// q := mongodb_mock.NewMockQuery(ctrl)
				// q.EXPECT().Select(bson.M{"_id": 1}).Return(q)
				// c := mongodb_mock.NewMockCollection(ctrl)
				// c.EXPECT().Find(nil).Return(q)
				// db.EXPECT().C("tags").Return(c)
			})
		})
	})
})
