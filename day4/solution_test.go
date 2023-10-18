package day4

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
)


func loadInput() (string, error) {
    _, file, _, _ := runtime.Caller(0)
    path := filepath.Join(filepath.Dir(file), "input.txt")
    data, err := os.ReadFile(path)
    if err != nil {
        return "", err
    }
    return string(data), nil
}


func TestPart1(t *testing.T) {
    input, err := loadInput()
    if err != nil {
        t.Fatal(err)
    }
    prefix := "00000"
    output := part1(input, prefix)
    expected := 254575
    if output != expected {
        t.Fatalf("expected %d but found %d", expected, output)
    }
}
