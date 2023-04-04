package e2e

import (
	"time"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe("Framework E2E", func() {
	// framework.NewDefaultFramework("framework-e2e")

	ginkgo.It("should run a failed test", func() {
		ginkgo.By("Running a failed test")
		gomega.Expect(false).To(gomega.BeTrue())
	})

	ginkgo.It("should run serial test", ginkgo.Serial, func(ctx ginkgo.SpecContext) {
		ginkgo.By("Running a serial test")
		time.Sleep(2 * time.Second)
	})
})
