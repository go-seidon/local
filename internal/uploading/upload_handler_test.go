package uploading_test

import (
	"errors"

	"github.com/go-seidon/local/cmd"
	"github.com/go-seidon/local/internal/clock"
	"github.com/go-seidon/local/internal/tests"
	. "github.com/go-seidon/local/internal/uploading"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

/**
 * ----------------------------------------------------------------------------
 * Spy FileRepository
 * ----------------------------------------------------------------------------
 */
type SpyFileRepository struct {
	CreateIsCalled  bool
	CreateLastParam *File
	CreateResult    error

	GetByUniqueIdIsCalled    bool
	GetByUniqueIdLastParam   string
	GetByUniqueIdFileResult  *File
	GetByUniqueIdErrorResult error
}

func (o *SpyFileRepository) Create(file *File) error {
	o.CreateIsCalled = true
	o.CreateLastParam = file
	return o.CreateResult
}
func (o *SpyFileRepository) GetByUniqueId(name string) (*File, error) {
	o.GetByUniqueIdIsCalled = true
	o.GetByUniqueIdLastParam = name
	return o.GetByUniqueIdFileResult, o.GetByUniqueIdErrorResult
}

/**
 * ----------------------------------------------------------------------------
 * Spy Filesytem
 * ----------------------------------------------------------------------------
 */
type SpyFilesystem struct {
	IsDirExistIsCalled  bool
	IsDirExistLastParam string
	IsDirExistResult    bool

	WriteBinaryFileToDiskIsCalled            bool
	WriteBinaryFileToDiskLastParamBinaryFile []byte
	WriteBinaryFileToDiskLastParamFullpath   string
	WriteBinaryFileToDiskResult              error
}

func (o *SpyFilesystem) IsDirExist(dirpath string) bool {
	o.IsDirExistIsCalled = true
	o.IsDirExistLastParam = dirpath
	return o.IsDirExistResult
}

func (o *SpyFilesystem) WriteBinaryFileToDisk(binaryFile []byte, fullpath string) error {
	o.WriteBinaryFileToDiskIsCalled = true
	o.WriteBinaryFileToDiskLastParamBinaryFile = binaryFile
	o.WriteBinaryFileToDiskLastParamFullpath = fullpath
	return o.WriteBinaryFileToDiskResult
}

/**
 * ----------------------------------------------------------------------------
 * TESTING
 * ----------------------------------------------------------------------------
 */

func expectErrorWithMsg(message string, err error) {
	Expect(err).NotTo(BeNil())
	Expect(err.Error()).To(Equal(message))
}

var _ = Describe("NewUploadHandler()", func() {
	fileRepo := &SpyFileRepository{}
	filesystem := &SpyFilesystem{}
	config := cmd.Config{FileDirectory: "/dir/path"}

	When("Given valid params", func() {
		It("returns no error", func() {
			_, err := NewUploadHandler(fileRepo, filesystem, config)
			Expect(err).To(BeNil())
		})
	})

	When("Given invalid params", func() {
		It("returns error if file repository object is not provided", func() {
			_, err := NewUploadHandler(nil, filesystem, config)
			expectErrorWithMsg("File repo is needed", err)
		})

		It("returns error if file directory path is not provided", func() {
			_, err := NewUploadHandler(fileRepo, filesystem, cmd.Config{FileDirectory: ""})
			expectErrorWithMsg("FileDirectory config is not set", err)
		})

		It("returns error if filesystem object is not provided", func() {
			_, err := NewUploadHandler(fileRepo, nil, config)
			expectErrorWithMsg("Filesystem is needed", err)
		})
	})
})

var _ = Describe("Upload handler", func() {
	/**
	 * Mocking for all test cases
	 */
	tests.MockClock()

	/**
	 * Param for NewUploadHandler that used by all test cases
	 */
	config := cmd.Config{FileDirectory: "/dir/path"}

	/**
	 * Param for UploadHandler that used by all test cases
	 */
	fileParam := UploadFileParam{
		FileData:            []byte{0, 1, 2},
		FileId:              "uniqueId",
		FileName:            "name",
		FileExtension:       "jpg",
		FileClientExtension: "jpg",
		FileMimetype:        "image/jpeg",
		FileSize:            1000,
	}

	/**
	 * File entity that used by all test cases
	 */
	file, _ := NewFile(
		fileParam.FileId,
		fileParam.FileName,
		fileParam.FileClientExtension,
		fileParam.FileExtension,
		fileParam.FileMimetype,
		config.FileDirectory,
		fileParam.FileSize,
		clock.Now(),
		clock.Now(),
	)

	/**
	 * Dependencies initialization that used by all test cases
	 * It is also doing reset before test run
	 */
	var spyFilesystem *SpyFilesystem
	var spyRepo *SpyFileRepository
	BeforeEach(func() {
		spyFilesystem = &SpyFilesystem{}
		spyRepo = &SpyFileRepository{}

		spyFilesystem.IsDirExistResult = true
		spyFilesystem.WriteBinaryFileToDiskResult = nil

		spyRepo.CreateResult = nil
		spyRepo.GetByUniqueIdFileResult = NewPredefinedFile()
		spyRepo.GetByUniqueIdErrorResult = nil
	})

	Describe("Integration with other objects", func() {
		It("UploadFile() calls filesystem.IsDirExist() correctly", func() {
			spyFilesystem.IsDirExistResult = false

			path := "/dir/path"
			uploadHandler, _ := NewUploadHandler(spyRepo, spyFilesystem, cmd.Config{FileDirectory: path})
			_, _ = uploadHandler.UploadFile(fileParam)

			Expect(spyFilesystem.IsDirExistIsCalled).To(BeTrue())
			Expect(spyFilesystem.IsDirExistLastParam).To(Equal(path))
		})

		It("UploadFile() calls fileRepo.Create() correctly", func() {
			uploadHandler, _ := NewUploadHandler(spyRepo, spyFilesystem, config)
			_, _ = uploadHandler.UploadFile(fileParam)

			expectedCreateParam := file

			Expect(spyRepo.CreateIsCalled).To(BeTrue())
			Expect(spyRepo.CreateLastParam.ToDTO()).To(Equal(expectedCreateParam.ToDTO()))
		})

		It("UploadFile() calls filesystem.WriteBinaryFileToDisk() correctly", func() {
			uploadHandler, _ := NewUploadHandler(spyRepo, spyFilesystem, config)
			_, _ = uploadHandler.UploadFile(fileParam)

			Expect(spyFilesystem.WriteBinaryFileToDiskIsCalled).To(BeTrue())
			Expect(spyFilesystem.WriteBinaryFileToDiskLastParamBinaryFile).To(Equal(fileParam.FileData))
			Expect(spyFilesystem.WriteBinaryFileToDiskLastParamFullpath).To(Equal(file.GetFullpath()))
		})
	})

	Context("UploadFile", func() {
		When("Given directory where file would be stored is not exists", func() {
			It("returns error", func() {
				spyFilesystem.IsDirExistResult = false
				uploadHandler, _ := NewUploadHandler(spyRepo, spyFilesystem, cmd.Config{FileDirectory: "/dir/path"})
				_, err := uploadHandler.UploadFile(fileParam)
				expectErrorWithMsg("directory is not exist", err)
			})
		})

		When("Fail to save file data to DB", func() {
			It("returns error ", func() {
				spyRepo := &SpyFileRepository{}
				spyRepo.CreateResult = errors.New("Failure on DB")

				uploadHandler, _ := NewUploadHandler(spyRepo, spyFilesystem, config)
				_, err := uploadHandler.UploadFile(fileParam)
				expectErrorWithMsg("Failure on DB", err)
			})
		})

		When("Fail to write binary file to filesystem storage", func() {
			It("returns error ", func() {
				spyFilesystem.WriteBinaryFileToDiskResult = errors.New("Failed to write to filesystem")

				uploadHandler, _ := NewUploadHandler(spyRepo, spyFilesystem, config)
				_, err := uploadHandler.UploadFile(fileParam)
				expectErrorWithMsg("Failed to write to filesystem", err)
			})
		})

		When("Success to store upload file", func() {
			It("returns UploadFileResult data and no error", func() {
				uploadHandler, _ := NewUploadHandler(spyRepo, spyFilesystem, config)
				result, err := uploadHandler.UploadFile(fileParam)

				Expect(err).To(BeNil())
				Expect(result).To(Equal(&UploadFileResult{
					FileId:     file.GetUniqueId(),
					FileName:   file.GetName(),
					UploadedAt: file.GetCreatedAt(),
				}))
			})
		})
	})
})
