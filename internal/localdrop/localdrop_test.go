package localdrop_test

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	. "github.com/ooesili/dead-drop/internal/localdrop"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
)

type errorReader struct {
	message string
}

func (e errorReader) Read(p []byte) (int, error) {
	return 0, errors.New(e.message)
}

var _ = Describe("Localdrop", func() {
	var (
		tempDir       string
		localFilename string
		file          io.Reader
		filename      = "file.txt"
		out           *gbytes.Buffer
		dropper       LocalDrop
	)

	BeforeEach(func() {
		var err error
		tempDir, err = ioutil.TempDir("", "dead-drop-")
		Expect(err).ToNot(HaveOccurred())

		out = gbytes.NewBuffer()
		dropper = LocalDrop{
			Out:     out,
			BaseDir: tempDir,
		}

		localFilename = filepath.Join(tempDir, filename)
		file = bytes.NewBufferString("file contents")
	})

	JustBeforeEach(func() {
		dropper.Drop(filename, file)
	})

	AfterEach(func() {
		os.RemoveAll(tempDir)
	})

	Context("when the file can successfully be downloaded", func() {
		It("creates the local file in the specified base dir", func() {
			Expect(localFilename).To(BeARegularFile())
		})

		It("writes the contents of the file to the local file", func() {
			Expect(ioutil.ReadFile(localFilename)).To(Equal([]byte("file contents")))
		})

		It("tells the user about the download progress", func() {
			Expect(out).To(gbytes.Say("downloading: file.txt\n"))
			Expect(out).To(gbytes.Say("download complete: file.txt\n"))
		})
	})

	Context("when the file cannot be created", func() {
		BeforeEach(func() {
			os.RemoveAll(tempDir)
		})

		It("tells the user about the error", func() {
			Expect(out).To(gbytes.Say("download failed: file.txt: "))
		})
	})

	Context("when there is an io error downloading the file", func() {
		BeforeEach(func() {
			file = errorReader{message: "shoot"}
		})

		It("lets the user know the download failed", func() {
			Expect(out).To(gbytes.Say("downloading: file.txt\n"))
			Expect(out).To(gbytes.Say("download failed: file.txt: "))
		})
	})

	Context("when the file already exists", func() {
		BeforeEach(func() {
			f, err := os.Create(localFilename)
			Expect(err).ToNot(HaveOccurred())
			fmt.Fprint(f, "existing file")
		})

		It("tells the user the file already exists", func() {
			Expect(out).To(gbytes.Say("download failed: file.txt: file already exists\n"))
		})

		It("does not overwrite the file", func() {
			Expect(ioutil.ReadFile(localFilename)).To(Equal([]byte("existing file")))
		})
	})
})
