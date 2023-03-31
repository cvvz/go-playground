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

var _ = ginkgo.SynchronizedAfterSuite(func(ctx context.Context) {}, func(ctx context.Context) {
	ginkgo.By("hahahahahaha")
}, ginkgo.NodeTimeout(1*time.Second))
