package cpanfilesnapshotparser

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type CpanfileSnapshotParser struct{}

func NewCpanfileSnapshotParser() CpanfileSnapshotParser {
	return CpanfileSnapshotParser{}
}

func (p CpanfileSnapshotParser) Parse(path string, distributionNamePart string) (bool, error) {
	file, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, err
		}

		return false, fmt.Errorf("failed to parse cpanfile.snapshot: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.TrimSpace(scanner.Text()) == "DISTRIBUTIONS" {
			for scanner.Scan() {
				distributionLinePattern := regexp.MustCompile("^  " + distributionNamePart)
				line := scanner.Bytes()
				if distributionLinePattern.Match(line) {
					return true, nil
				}
			}
		}
	}
	return false, nil
}
