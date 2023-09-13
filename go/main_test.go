package main

import (
	//"fmt"
	"testing"
	"math/rand"
)

func almostEqual(a, b float64) bool {
	return b-a <= 1
}

func randomMoneyAmount() float64 {
	return rand.Float64() * 10000
}

/*
func NTestScanFolderTotalWithExistingDirectory(t *testing.T) {
	// Provide the path to the existing "Costexample" directory for testing
	existingDir := "../Costexample"

	// Calculate the total cost in the "Costexample" directory
	total := costTotal()

	// Replace this with the expected total cost based on the contents of the "Costexample" directory
	expectedTotal := 1955.36
	fmt.Println(total)
	if !almostEqual(total, expectedTotal) {
		t.Errorf("Expected total cost %.2f for directory %s, but got %.2f", expectedTotal, existingDir, total)
	}
}*/

func TestCalcCost(t *testing.T) {
	// Test a valid filename with a cost
	filename1 := "file1_2.50.txt"
	cost1 := calcCost(filename1, "_")
	expectedCost1 := 2.50
	if cost1 != expectedCost1 {
		t.Errorf("Expected cost %.2f for %s, but got %.2f", expectedCost1, filename1, cost1)
	}

	// Test a filename without a cost
	filename2 := "file2.txt"
	cost2 := calcCost(filename2, "_")
	expectedCost2 := 0.0
	if cost2 != expectedCost2 {
		t.Errorf("Expected cost %.2f for %s, but got %.2f", expectedCost2, filename2, cost2)
	}

	// Test an invalid filename
	filename3 := "file3.abc.txt"
	cost3 := calcCost(filename3, "_")
	expectedCost3 := 0.0
	if cost3 != expectedCost3 {
		t.Errorf("Expected cost %.2f for %s, but got %.2f", expectedCost3, filename3, cost3)
	}

	filename4 := "file4_232_0.0.txt"
	cost4 := calcCost(filename4, "_")
	expectedCost4 := 0.0
	if cost4 != expectedCost4 {
		t.Errorf("Expected cost %.2f for %s, but got %.2f", expectedCost4, filename4, cost4)
	}
}