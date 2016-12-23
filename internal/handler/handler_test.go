package handler_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"

	. "github.com/ooesili/dead-drop/internal/handler"
	"github.com/ooesili/dead-drop/test/mocks"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func fileUpload(field, filename, contents string) (*bytes.Buffer, string) {
	body := &bytes.Buffer{}

	multipartWriter := multipart.NewWriter(body)

	fileWriter, err := multipartWriter.CreateFormFile("file", filename)
	Expect(err).ToNot(HaveOccurred())

	fmt.Fprint(fileWriter, "file contents")
	multipartWriter.Close()

	return body, multipartWriter.FormDataContentType()
}

var _ = Describe("Handler", func() {
	var (
		body     *bytes.Buffer
		req      *http.Request
		resp     *httptest.ResponseRecorder
		renderer *mocks.Renderer
		dropper  *mocks.Dropper
		handler  http.Handler
	)

	BeforeEach(func() {
		body = nil
		req = nil
		resp = httptest.NewRecorder()
		dropper = &mocks.Dropper{}
		renderer = &mocks.Renderer{}
		handler = NewRouter(dropper, renderer)
	})

	Describe("GET /", func() {
		BeforeEach(func() {
			var err error
			req, err = http.NewRequest("GET", "/", nil)
			Expect(err).ToNot(HaveOccurred())

			renderer.RenderCall.Renders = "the drop form"

			handler.ServeHTTP(resp, req)
		})

		It("responds with a 200 code", func() {
			Expect(resp.Code).To(Equal(http.StatusOK))
		})

		It("renders the index template", func() {
			recieved := renderer.RenderCall.Recieved
			Expect(recieved.Template).To(Equal("index"))
			Expect(recieved.Request).ToNot(BeNil())
		})

		It("responds with the template", func() {
			Expect(resp.Body.String()).To(Equal("the drop form"))
		})
	})

	Describe("POST /drop", func() {
		Context("when the request is valid multipart form data", func() {
			BeforeEach(func() {
				var contentType string
				body, contentType = fileUpload("file", "data.txt", "file contents")

				var err error
				req, err = http.NewRequest("POST", "/drop", body)
				Expect(err).ToNot(HaveOccurred())
				req.Header.Add("Content-Type", contentType)

				handler.ServeHTTP(resp, req)
			})

			It("redirects with a 303 status code", func() {
				Expect(resp.Code).To(Equal(http.StatusSeeOther))
			})

			It("redirects to /complete", func() {
				Expect(resp.HeaderMap.Get("Location")).To(Equal("/complete"))
			})

			It("passes the correct filename to the dropper", func() {
				Expect(dropper.Recieved.Filename).To(Equal("data.txt"))
			})

			It("passes the uploaded file to the droper", func() {
				contents, err := ioutil.ReadAll(dropper.Recieved.File)
				Expect(err).ToNot(HaveOccurred())
				Expect(contents).To(Equal([]byte("file contents")))
			})
		})

		Context("when the filename has slashes in it", func() {
			BeforeEach(func() {
				var contentType string
				body, contentType = fileUpload("file", "path/with/slashes/data.txt", "file contents")

				var err error
				req, err = http.NewRequest("POST", "/drop", body)
				Expect(err).ToNot(HaveOccurred())
				req.Header.Add("Content-Type", contentType)

				handler.ServeHTTP(resp, req)
			})

			It("redirects to /complete", func() {
				Expect(resp.HeaderMap.Get("Location")).To(Equal("/complete"))
			})

			It("takes the basename of the file before calling dropper", func() {
				Expect(dropper.Recieved.Filename).To(Equal("data.txt"))
			})
		})

		Context("when the request body is invalid", func() {
			BeforeEach(func() {
				var contentType string
				_, contentType = fileUpload("file", "data.txt", "file contents")
				body = bytes.NewBufferString("totally not multipart/form-data")

				var err error
				req, err = http.NewRequest("POST", "/drop", body)
				Expect(err).ToNot(HaveOccurred())
				req.Header.Add("Content-Type", contentType)

				handler.ServeHTTP(resp, req)
			})

			It("responds with a 400 status code", func() {
				Expect(resp.Code).To(Equal(http.StatusBadRequest))
			})

			It("responds with no body", func() {
				Expect(resp.Body.String()).To(Equal(""))
			})
		})
	})

	Describe("GET /complete", func() {
		BeforeEach(func() {
			var err error
			req, err = http.NewRequest("GET", "/complete", nil)
			Expect(err).ToNot(HaveOccurred())

			renderer.RenderCall.Renders = "the complete page"

			handler.ServeHTTP(resp, req)
		})

		It("responds with a 200 code", func() {
			Expect(resp.Code).To(Equal(http.StatusOK))
		})

		It("renders the index template", func() {
			recieved := renderer.RenderCall.Recieved
			Expect(recieved.Template).To(Equal("complete"))
			Expect(recieved.Request).ToNot(BeNil())
		})

		It("responds with the template", func() {
			Expect(resp.Body.String()).To(Equal("the complete page"))
		})
	})
})
