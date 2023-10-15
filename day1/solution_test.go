package day1

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
    "testing"
)


func loadInput() string {
    _, file, _, _ := runtime.Caller(0)
    path := filepath.Dir(file)
    content, err := os.ReadFile(filepath.Join(path, "input.txt"))
    if err != nil {
        log.Fatal(err)
    }
    return string(content)
}


func TestPart1(t *testing.T) {
    input := loadInput()
    output := part1(input)
    if output != 232 {
        t.Errorf("Expected 232, got %d", output)
    }
}


func TestPart2(t *testing.T) {
    input := loadInput()
    output, ok := part2(input)
    if !ok {
        t.Error("santa never entered the basement")
    }
    if output != 1783 {
        t.Errorf("Expected 1783, got %d", output)
    }
}
