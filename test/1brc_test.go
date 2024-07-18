package main

import (
	"fmt"
	"log"
	oneBRC "oneBRC/src"
	"os"
	"path/filepath"
	"testing"
)

func runTest(t *testing.T, inputPath string, expectedAnswer string) {
	t.Helper()
	answer := oneBRC.BrcNaive(inputPath)
	if answer != expectedAnswer {
		t.Errorf("Got: %s, expected answer: %s", answer, expectedAnswer)
	}
}

func runTestBenchmark(b *testing.B, inputPath string, expectedAnswer string) {
	b.Helper()
	answer := oneBRC.BrcNaive(inputPath)
	if answer != expectedAnswer {
		b.Errorf("Got: %s, expected answer: %s", answer, expectedAnswer)
	}
}

func Test1Brc(t *testing.T) {
	inputFiles, err := filepath.Glob(filepath.Join("./resources/samples", "*.txt"))
	if err != nil {
		log.Fatalln("Error reading input files from test directory:", err)
	}

	for _, inputFile := range inputFiles {
		inputFileWithoutExt := inputFile[:len(inputFile)-4]
		outputFilePath := fmt.Sprintf("%s.out", inputFileWithoutExt)
		outputFileContent, err := os.ReadFile(outputFilePath)

		expectedAnswer := string(outputFileContent)

		if err != nil {
			t.Fatalf("Error in opening the output file: %s\n", err)
			return
		}
		t.Run(inputFile, func(t *testing.T) {
			runTest(t, fmt.Sprintf("../test/%s", inputFile), expectedAnswer)
		})
	}
}

func Benchmark1Brc(b *testing.B) {
	inputFiles, err := filepath.Glob(filepath.Join("./resources/samples", "*.txt"))
	if err != nil {
		log.Fatalln("Error reading input files from test directory:", err)
	}

	for _, inputFile := range inputFiles {
		inputFileWithoutExt := inputFile[:len(inputFile)-4]
		outputFilePath := fmt.Sprintf("%s.out", inputFileWithoutExt)
		outputFileContent, err := os.ReadFile(outputFilePath)

		expectedAnswer := string(outputFileContent)

		if err != nil {
			b.Fatalf("Error in opening the output file: %s\n", err)
			return
		}
		b.Run(inputFile, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				runTestBenchmark(b, fmt.Sprintf("../test/%s", inputFile), expectedAnswer)
			}
		})
	}
}
