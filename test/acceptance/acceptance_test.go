package acceptance_test

import (
	"os"
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
	"github.com/sclevine/agouti"
	. "github.com/sclevine/agouti/matchers"
)

var _ = Describe("Acceptance", func() {
	var (
		page    *agouti.Page
		session *gexec.Session
	)

	BeforeEach(func() {
		cmd := exec.Command(ddropPath)

		var err error
		session, err = gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
		Expect(err).ToNot(HaveOccurred())

		Eventually(session).Should(gbytes.Say("listening on 0.0.0.0:3000 for dead drop"))

		page, err = agoutiDriver.NewPage()
		Expect(err).NotTo(HaveOccurred())

		Expect(page.Navigate("http://localhost:3000")).To(Succeed())
	})

	AfterEach(func() {
		Expect(page.Destroy()).To(Succeed())
		session.Terminate().Wait()
	})

	It("allows the user to upload a file", func() {
		By("selecting a file to upload", func() {
			fileInput := page.FindByLabel("pick a file")
			Eventually(fileInput).Should(BeFound())
			Expect(fileInput.UploadFile("fixtures/file.txt")).To(Succeed())
		})

		By("clicking the upload button", func() {
			uploadInput := page.FindByButton("upload")
			Expect(uploadInput.Click()).To(Succeed())
		})

		By("looking at the server's filesystem", func() {
			Expect("./file.txt").To(BeARegularFile())
			os.Remove("./file.txt")
		})
	})
})
