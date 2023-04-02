package books_test

import (
	"context"
	"testing"
	"time"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

func TestBooks(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "Books Suite")
}

var _ = ginkgo.SynchronizedBeforeSuite(func(ctx context.Context) []byte {
	ginkgo.By("=================")
	return []byte{}
}, func(ctx context.Context, data []byte) {
	ginkgo.By("****************")
})

var _ = ginkgo.SynchronizedAfterSuite(func(ctx context.Context) {
	ginkgo.By("****************")
}, func(ctx context.Context) {
	ginkgo.By("=================")
}, ginkgo.NodeTimeout(1*time.Second))
