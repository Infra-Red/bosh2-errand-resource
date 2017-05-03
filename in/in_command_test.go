package in_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"errors"
	"io/ioutil"
	"path/filepath"

	"github.com/starkandwayne/bosh2-errand-resource/bosh/boshfakes"
	"github.com/starkandwayne/bosh2-errand-resource/concourse"
	"github.com/starkandwayne/bosh2-errand-resource/in"
)

var _ = Describe("InCommand", func() {
	var (
		inCommand in.InCommand
		director  *boshfakes.FakeDirector
	)

	BeforeEach(func() {
		director = new(boshfakes.FakeDirector)
		inCommand = in.NewInCommand(director)
	})

	Describe("Run", func() {
		var inRequest concourse.InRequest
		var targetDir string
		sillyBytes := []byte{0xFE, 0xED, 0xDE, 0xAD, 0xBE, 0xEF}
		sillyBytesSha1 := "33bf00cb7a45258748f833a47230124fcc8fa3a4"
		wrongBytes := []byte{0x0F, 0xFE, 0xEF, 0xBE, 0xCF, 0xF0}

		BeforeEach(func() {
			inRequest = concourse.InRequest{
				Source: concourse.Source{
					Target: "director.example.com",
				},
				Version: concourse.Version{
					ManifestSha1: sillyBytesSha1,
					Target:       "director.example.com",
				},
			}

			var err error
			targetDir, err = ioutil.TempDir("", "")
			Expect(err).ToNot(HaveOccurred())

			director.DownloadManifestReturns(sillyBytes, nil)
		})

		It("writes the manifest and target to disk and returns the version as a response", func() {
			inResponse, err := inCommand.Run(inRequest, targetDir)
			Expect(err).ToNot(HaveOccurred())

			manifestBytes, err := ioutil.ReadFile(filepath.Join(targetDir, "manifest.yml"))
			Expect(err).ToNot(HaveOccurred())
			Expect(manifestBytes).To(Equal(sillyBytes))
			Expect(director.DownloadManifestCallCount()).To(Equal(1))

			targetBytes, err := ioutil.ReadFile(filepath.Join(targetDir, "target"))
			Expect(err).ToNot(HaveOccurred())
			Expect(string(targetBytes)).To(Equal("director.example.com"))

			Expect(inResponse).To(Equal(in.InResponse{
				Version: concourse.Version{
					ManifestSha1: sillyBytesSha1,
					Target:       "director.example.com",
				},
			}))
		})

		Context("when the manifest download fails", func() {
			BeforeEach(func() {
				director.DownloadManifestReturns(nil, errors.New("could not download manifest"))
			})

			It("returns an error", func() {
				_, err := inCommand.Run(inRequest, targetDir)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("could not download manifest"))
			})
		})

		Context("when downloaded manifest does not match the requested version", func() {
			BeforeEach(func() {
				director.DownloadManifestReturns(wrongBytes, nil)
			})

			It("returns an error", func() {
				_, err := inCommand.Run(inRequest, targetDir)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("Requested deployment version is not available"))
			})
		})

		Context("when director target does not match the requested version", func() {
			BeforeEach(func() {
				inRequest.Source.Target = "weird.example.com"
			})

			It("returns an error", func() {
				_, err := inCommand.Run(inRequest, targetDir)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("Requested deployment director is different than configured source"))
			})
		})
	})
})