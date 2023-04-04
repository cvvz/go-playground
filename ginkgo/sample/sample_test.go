package sample

import (
	"testing"
	"time"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = ginkgo.SynchronizedAfterSuite(func(ctx ginkgo.SpecContext) {
	ginkgo.By("************** allProcessBody begin")
	time.Sleep(2 * time.Second)
	ginkgo.By("************** process1Body end")
}, func(ctx ginkgo.SpecContext) {
	ginkgo.By("************** process1Body begin")
	time.Sleep(2 * time.Second)
	ginkgo.By("************** process1Body end")
})

var _ = ginkgo.Describe("Framework E2E", func() {
	ginkgo.It("should run a failed test", func() {
		ginkgo.By("Running a failed test")
		gomega.Expect(false).To(gomega.BeTrue())
	})

	ginkgo.It("should run serial test", ginkgo.Serial, func(ctx ginkgo.SpecContext) {
		ginkgo.By("Running a serial test")
		time.Sleep(2 * time.Second)
	})
})

func TestE2E(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "Sample Suite")
}
