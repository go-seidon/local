package uploading_test

import (
	"testing"
	"time"

	. "github.com/go-seidon/local/internal/uploading"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestHealthCheck(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Uploading Package")
}

func NewPredefinedFile() *File {
	file, _ := NewFile(
		"uniqueId",
		"name",
		"clientExtension",
		"extension",
		"mimetype",
		"/dir/path",
		1000,
		time.Now(),
		time.Now(),
	)
	return file
}

var _ = Describe("File entity", func() {
	file, _ := NewFile(
		"uniqueId",
		"name",
		"clientExtension",
		"extension",
		"mimetype",
		"/dir/path",
		1000,
		time.Now(),
		time.Now(),
	)

	It("Clone() behave correctly", func() {
		Expect(file.ToDTO()).To(Equal(file.Clone().ToDTO()))
	})

	It("setter and getter methods behave correctly", func() {
		createdAt := time.Now()
		updatedAt := time.Now().AddDate(1, 0, 0)

		f := file.Clone()
		f.SetUniqueId("uniqueId")
		f.SetName("name")
		f.SetExtension("jpg")
		f.SetClientExtension("png")
		f.SetMimetype("image/jpeg")
		f.SetDirpath("/dir/path")
		f.SetSize(1000)
		f.SetCreatedAt(createdAt)
		f.SetUpdatedAt(updatedAt)

		Expect(f.GetFullpath()).To(Equal("/dir/path/uniqueId"))
		Expect(f.GetSize()).To(Equal(uint32(1000)))
		Expect(f.GetDirectoryPath()).To(Equal("/dir/path"))
		Expect(f.GetUniqueId()).To(Equal("uniqueId"))
		Expect(f.GetName()).To(Equal("name"))
		Expect(f.GetExtension()).To(Equal("jpg"))
		Expect(f.GetClientExtension()).To(Equal("png"))
		Expect(f.GetMimetype()).To(Equal("image/jpeg"))
		Expect(f.GetCreatedAt()).To(Equal(createdAt))
		Expect(f.GetUpdatedAt()).To(Equal(updatedAt))
	})

	expectErrorWithMsg := func(message string, err error) {
		Expect(err).NotTo(BeNil())
		Expect(err.Error()).To(Equal(message))
	}

	Context("NewFile() factory", func() {
		When("Given invalid param", func() {
			It("returns error", func() {
				const badUniqueId = ""
				_, err := NewFile(
					badUniqueId,
					"name",
					"clientExtension",
					"extension",
					"mimetype",
					"dirpath",
					1000,
					time.Now(),
					time.Now(),
				)
				expectErrorWithMsg("uniqueId is mandatory", err)
			})
		})

		When("Given valid param", func() {
			It("returns no error", func() {
				_, err := NewFile(
					"uniqueId",
					"name",
					"clientExtension",
					"extension",
					"mimetype",
					"dirpath",
					1000,
					time.Now(),
					time.Now(),
				)
				Expect(err).To(BeNil())
			})
		})
	})

	Context("Validate()", func() {
		When("Required data is empty", func() {
			It("returns error", func() {
				var err error

				err = file.Clone().SetUniqueId("").Validate()
				expectErrorWithMsg("uniqueId is mandatory", err)

				err = file.Clone().SetName("").Validate()
				expectErrorWithMsg("name is mandatory", err)

				err = file.Clone().SetExtension("").Validate()
				expectErrorWithMsg("extension is mandatory", err)

				err = file.Clone().SetMimetype("").Validate()
				expectErrorWithMsg("mimetype is mandatory", err)

				err = file.Clone().SetDirpath("").Validate()
				expectErrorWithMsg("dirpath is mandatory", err)
			})
		})
	})
})
