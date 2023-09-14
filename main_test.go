package main

import (
	"fmt"
	"math"
	"math/rand"
	"slices"
	"testing"
)

func randomMoneyAmount() float64 {
	return math.Round(rand.Float64()*1000*100) / 100
}

func expectSliceEq(a, b []string) bool {
	return slices.Equal[[]string](a, b)
}

func TestScanFolderWithExistingDirectory(t *testing.T) {
	// Provide the path to the existing "Costexample" directory for testing
	existingDir := "./Costexample"

	// Calculate the total cost in the "Costexample" directory
	files, _ := scanFolder(existingDir)

	expected := []string{"Costexample\\Abcam Order_081523_698.pdf", "Costexample\\Amazon_Gift Card_072123_250.pdf", "Costexample\\Amazon_Gift Card_081423_250.pdf", "Costexample\\Amazon_OBOC_072423_25.70.pdf", "Costexample\\Amazon_Office suplly_080323_314.14.pdf", "Costexample\\Amazon_Office suplly_080323_338.10.pdf", "Costexample\\Amazon_Office suplly_080323_9.44.pdf", "Costexample\\Amazon_Office suplly_081523_34.99.pdf", "Costexample\\Amazon_Office suplly_081623_34.99.pdf"}
	if !expectSliceEq(files, expected) {
		t.Errorf("Expected %v, but got %v", expected, files)
	}
}

func TestCalcCost(t *testing.T) {
	// Test a valid filename with a cost
	for i := 0; i < 100; i++ {
		amt := randomMoneyAmount()
		filename1 := fmt.Sprintf("file1_%v.pdf", amt)
		cost1, _ := calcCost(filename1, "_")
		if cost1 != amt {
			t.Errorf("Expected cost %.2f for %s, but got %.2f", amt, filename1, cost1)
		}
	}

	filename4 := "file4_232_0.0.txt"
	cost4, _ := calcCost(filename4, "_")
	expectedCost4 := 0.0
	if cost4 != expectedCost4 {
		t.Errorf("Expected cost %.2f for %s, but got %.2f", expectedCost4, filename4, cost4)
	}
}
