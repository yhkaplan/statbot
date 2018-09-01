package coverage

import (
	"fmt"
	"os"
)

// Run xcrun xccov view --only-targets --json derivedDataPath
func getCoverageInfo(derivedDataPath string) (string, error) {
}

func parseJSON() {}

// Convert"lineCoverage": 0.05829249107521026
func calculateCoverageStatistics(coverage float32) float32 {
	return coverage * 100
}

// TODO: Check for DerivedData path envar
// If not set, assume default
func getDerivedDataPath(targetName string) string {
	envarPath := os.Getenv("DERIVED_DATA_PATH")
	if envarPath == "" {
		return fmt.Sprintf("~/Library/Developer/Xcode/DerivedData/%s-*/Build/Products/Logs/Test/*.xccovreport", targetName)
	}

	return envarPath
}
