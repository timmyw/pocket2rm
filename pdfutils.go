package pocket2rm

import (
	"os"
	"fmt"
)

// GenerateOutputFilename returns a full output path
func GenerateOutputFilename(outputPath string, source string) (string, error) {
	err := os.MkdirAll(outputPath, 0775)

	if err != nil {
		return "", err
	}
	
	return fmt.Sprintf("%s/%s.pdf", outputPath, source), nil
}
