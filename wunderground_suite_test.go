package wunderground_test

import (
	"testing"

	"github.com/jarcoal/httpmock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = BeforeSuite(func() {
	httpmock.Activate()
})

var _ = AfterEach(func() {
	httpmock.Reset()
})

var _ = AfterSuite(func() {
	httpmock.DeactivateAndReset()
})

func TestWunderground(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Wunderground Suite")
}
