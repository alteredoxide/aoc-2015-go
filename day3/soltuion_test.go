package day3;

import (
    "os"
    "path/filepath"
    "runtime"
    "testing"
)


func loadInput() (string, error) {
    _, file, _, _ := runtime.Caller(0)
    path := filepath.Dir(file)
    data, err := os.ReadFile(filepath.Join(path, "input.txt"))
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
    output := part1(input)
    var expected int = 2565
    if output != expected {
        t.Fatalf("exptected %d but found %d", expected, output)
    }
}


func TestPart2(t *testing.T) {
    input, err := loadInput()
    if err != nil {
        t.Fatal(err)
    }
    output := part2(input)
    var expected int = 2639
    if output != expected {
        t.Fatalf("expected %d but found %d", expected, output)
    }
}
