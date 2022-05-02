package cpanfilesnapshotparser_test

import (
	"io"
	"os"
	"runtime"
	"strings"
	"testing"

	. "github.com/onsi/gomega"
	"github.com/sclevine/spec"
	cpanfilesnapshotparser "github.com/yujiorama/go-cpanfile-snapshot-parser"
)

func testCpanfileSnapshotParser(t *testing.T, context spec.G, it spec.S) {
	var (
		Expect = NewWithT(t).Expect

		path   string
		parser cpanfilesnapshotparser.CpanfileSnapshotParser
	)

	it.Before(func() {
		source, err := os.Open("./testdata/cpanfile.snapshot")
		Expect(err).NotTo(HaveOccurred())
		defer source.Close()
		file, err := os.CreateTemp("", "cpanfile.snapshot")
		Expect(err).NotTo(HaveOccurred())
		defer file.Close()

		_, err = io.Copy(file, source)
		Expect(err).NotTo(HaveOccurred())

		path = file.Name()

		parser = cpanfilesnapshotparser.NewCpanfileSnapshotParser()
	})

	it.After(func() {
		Expect(os.RemoveAll(path)).To(Succeed())
	})

	context("Parse", func() {
		it("parses the cpanfile.snapshot file to check for plack module", func() {
			foundPlack, err := parser.Parse(path, "Plack-")
			Expect(err).NotTo(HaveOccurred())
			Expect(foundPlack).To(Equal(true))
		})

		context("when the cpanfile.snapshot file does not exist", func() {
			it.Before(func() {
				Expect(os.Remove(path)).To(Succeed())
			})

			it("returns an ErrNotExist error", func() {
				_, err := parser.Parse(path, "Plack-")
				Expect(os.IsNotExist(err)).To(Equal(true))
			})
		})

		context("failure cases", func() {
			context("when the cpanfile.snapshot cannot be opened", func() {
				it.Before(func() {
					Expect(os.Chmod(path, 0000)).To(Succeed())
				})

				it("returns an error", func() {
					if strings.Compare("windows", runtime.GOOS) == 0 {
						t.SkipNow()
					}
					_, err := parser.Parse(path, "Plack-")
					Expect(err).To(MatchError(ContainSubstring("failed to parse cpanfile.snapshot:")))
					Expect(err).To(MatchError(ContainSubstring("permission denied")))
				})
			})
		})
	})
}
