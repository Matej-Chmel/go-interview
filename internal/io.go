package internal

import (
	"io"
	"os"
	"path/filepath"
	"strings"

	nio "github.com/Matej-Chmel/go-number-io"
)

// Opens a file on the relative path
func OpenFile(relPath string) (*os.File, error) {
	filePath, err := ProgramPathJoin(relPath)

	if err != nil {
		return nil, err
	}

	return os.Open(filePath)
}

// For user code, returns absolute path to the main Go file.
// For tests, returns empty path.
func ProgramPath() (string, error) {
	if len(os.Args) < 2 {
		return "", nil
	}

	if strings.Contains(os.Args[1], "-test") {
		return "", nil
	}

	return filepath.Abs(os.Args[1])
}

// Joins relPath to the path of the main Go file.
func ProgramPathJoin(relPath string) (string, error) {
	progPath, err := ProgramPath()

	if err != nil {
		return "", err
	}

	return filepath.Join(progPath, relPath), nil
}

// Reads all text from a file on relative path relPath.
func ReadAllText(relPath string) (string, error) {
	file, err := OpenFile(relPath)

	if err != nil {
		return "", err
	}

	defer file.Close()
	content, err := io.ReadAll(file)

	if err != nil {
		return "", err
	}

	return strings.ReplaceAll(string(content), "\r\n", "\n"), nil
}

// Reads 1D, 2D or 3D slice from a file on relative path relPath
func ReadData[T any](relPath string) (T, error) {
	file, err := OpenFile(relPath)

	if err != nil {
		var res T
		return res, err
	}

	defer file.Close()
	return nio.Read[T](file)
}
