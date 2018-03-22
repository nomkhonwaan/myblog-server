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

	Describe("NewPlaceholder", func() {
		It("should return a new placeholder which is an empty Tag object correctly", func() {
			placeholder := NewPlaceholder()

			Expect(reflect.TypeOf(placeholder).String()).To(Equal("*tag.Tag"))
			Expect(string(placeholder.(*Tag).ID)).To(Equal(""))
			Expect(placeholder.(*Tag).Name).To(Equal(""))
			Expect(placeholder.(*Tag).Slug).To(Equal(""))
		})
	})
})
