package day6;

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)


func fromBytes(data []byte) []string {
    strData := string(data)
    out := []string{}
    for _, line := range strings.Split(strData, "\n") {
        out = append(out, strings.Trim(line, "\n"))
    }
    return out
}


func loadInput() ([]string, error) {
    _, file, _, _ := runtime.Caller(0)
    path := filepath.Dir(file)
    data, err := os.ReadFile(filepath.Join(path, "input.txt"))
    if err != nil {
        return []string{}, err
    }
    return fromBytes(data), nil
}


func TestPart1(t *testing.T) {
    input, err := loadInput()
    if err != nil {
        t.Fatal(err)
    }
    output := partOne(input)
    expected := 400410
    if output != expected {
        t.Fatalf("expected %d but got %d", expected, output)
    }
}
