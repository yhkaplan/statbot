package coverage

import (
	"fmt"
	"math"
)

// Run xcrun xccov view --only-targets --json derivedDataPath
// func getCoverageInfo(derivedDataPath string) (string, err) {
// }

func parseJSON() {}

// Convert"lineCoverage": 0.05829249107521026
func ConvertToPercentage(coverage float64) string {
	percentage := math.Round(coverage*10000.0) / 10000.0 * 100.0
	return fmt.Sprintf("%.2f", percentage)
}

// TODO: Check for DerivedData path envar
// If not set, assume default
// func getDerivedDataPath(targetName string) (string, err) {
// 	envarPath, _ := os.Getenv("DERIVED_DATA_PATH")
// 	if envarPath {
// 		return envarPath
// 	}

// 	return fmt.Sprintf("~/Library/Developer/Xcode/DerivedData/%s-*/Build/Products/Logs/Test/*.xccovreport", targetName)
// }
