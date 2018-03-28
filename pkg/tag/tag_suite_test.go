package tag_test

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

var (
	ctrl *gomock.Controller
)

func TestTag(t *testing.T) {
	RegisterFailHandler(Fail)
	ctrl = gomock.NewController(t)
	RunSpecs(t, "Tag Suite")
}
