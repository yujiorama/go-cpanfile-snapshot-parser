package cpanfilesnapshotparser_test

import (
	"testing"

	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"
)

func TestUnitCpanfileSnapshotParser(t *testing.T) {
	suite := spec.New("cpanfilesnapshotparser", spec.Report(report.Terminal{}), spec.Parallel())
	suite("CpanfileSnapshotParser", testCpanfileSnapshotParser)
	suite.Run(t)
}
