package e2e

import (
	"flag"
	"os"
	"testing"
	"time"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"k8s.io/kubernetes/test/e2e/framework"
	"k8s.io/kubernetes/test/e2e/framework/config"
	_ "k8s.io/kubernetes/test/e2e/framework/debug/init"
)

//	var _ = ginkgo.AfterSuite(func(ctx ginkgo.SpecContext) {
//		ginkgo.By("************** AfterSuite begin")
//		time.Sleep(3 * time.Second)
//		<-ctx.Done()
//		ginkgo.By("************** AfterSuite ctx done")
//	})
var a int

var _ = ginkgo.SynchronizedAfterSuite(func(ctx ginkgo.SpecContext) {
	// ginkgo.By("************** allProcessBody begin")
	// fmt.Println("!!!!!!!!!!!!!!! allProcessBody begin")
	// time.Sleep(3 * time.Second)
	// ginkgo.By("************** allProcessBody end")
	// // <-ctx.Done()
	// ginkgo.By(fmt.Sprintf("a = %d", a))
	// ginkgo.By("************** allProcessBody ctx done")
	// fmt.Println("!!!!!!!!!!!!!!! allProcessBody end")
}, func(ctx ginkgo.SpecContext) {
	ginkgo.By("************** process1Body begin")
	time.Sleep(2 * time.Second)
	ginkgo.By("************** process1Body end")
})

func TestMain(m *testing.M) {
	config.CopyFlags(config.Flags, flag.CommandLine)
	framework.RegisterCommonFlags(flag.CommandLine)
	framework.RegisterClusterFlags(flag.CommandLine)
	flag.Parse()
	framework.AfterReadingAllFlags(&framework.TestContext)

	os.Exit(m.Run())
}

func TestE2E(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "kubernetes e2e framework Suite")
}
