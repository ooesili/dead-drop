package acceptance_test

import (
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
	"github.com/sclevine/agouti"

	"testing"
)

var (
	agoutiDriver *agouti.WebDriver
	ddropPath    string
)

func TestAcceptance(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Acceptance Suite")
}

var _ = BeforeSuite(func() {
	goGenerate := exec.Command("sh", "-c", "cd ../../ && echo go generate $(glide novendor)")
	Expect(goGenerate.Run()).To(Succeed())

	agoutiDriver = agouti.PhantomJS()
	Expect(agoutiDriver.Start()).To(Succeed())

	var err error
	ddropPath, err = gexec.Build("github.com/ooesili/dead-drop/cmd/ddrop")
	Expect(err).ToNot(HaveOccurred())
})

var _ = AfterSuite(func() {
	gexec.CleanupBuildArtifacts()
	Expect(agoutiDriver.Stop()).To(Succeed())
})
