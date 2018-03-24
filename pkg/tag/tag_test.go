package tag_test

import (
	"reflect"

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

	// Describe("Repository.FindByID", func() {
	// 	Context("with existing Post'd ID", func() {
	// 		It("should return a Post from its ID", func() {
	// 			repo := Repository{db}
	// 		})
	// 	})
	// })
})
