package files

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func ReadAllLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func SaveLines(path string, lines []string) error {
	file, err := create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		_, _ = fmt.Fprintln(w, line)
	}
	return w.Flush()
}

func create(p string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(p), os.ModePerm); err != nil {
		return nil, err
	}
	return os.Create(p)
}
