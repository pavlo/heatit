package app

import (
	"github.com/pavlo/heatit/utils"
	"strings"
	"testing"
)

func TestEngineProcessInsertsSimpleInsertion(t *testing.T) {
	assertInserts(
		t,
		"../fixtures/engine/simple_insertion/a.yaml",
		"../fixtures/engine/simple_insertion/result.yaml",
	)
}

func TestEngineProcessInsertsRecursiveInsertion(t *testing.T) {
	assertInserts(
		t,
		"../fixtures/engine/recursive_insertion/a.yaml",
		"../fixtures/engine/recursive_insertion/result.yaml",
	)
}

func TestEngineProcessInsertsComplete(t *testing.T) {
	assertInserts(
		t,
		"../fixtures/engine/complete/heat.yaml",
		"../fixtures/engine/complete/result.yaml",
	)
}

func TestProcessParams(t *testing.T) {
	data := readFixture(t, "../fixtures/parameters/heat.yaml")
	p, err := NewParameters("../fixtures/parameters/params.yaml")

	if err != nil {
		t.Fatal("Failed to read params fixture!")
	}

	actual, err := processParams(data, p)
	if err != nil {
		t.Fatalf("Failed to process parameters")
	}

	expected := readFixture(t, "../fixtures/parameters/result.yaml")

	if strings.TrimSpace(actual) != strings.TrimSpace(expected) {
		t.Log("Failed to compare results!")
		t.Logf("Expected: \n[%v]\n", expected)
		t.Logf("Actual: \n[%v]\n", actual)
		t.Fail()
	}

}

func assertInserts(t *testing.T, sourceFile string, expectedResultFile string) {
	data := readFixture(t, sourceFile)

	content, err := processInserts(data, 0)
	if err != nil {
		t.Fatal("Failed to process inserts!")
	}

	expectedContent := readFixture(t, expectedResultFile)
	if strings.TrimSpace(content) != strings.TrimSpace(expectedContent) {
		t.Log("Failed to compare results!")
		t.Logf("Expected: \n[%v]\n", expectedContent)
		t.Logf("Actual: \n[%v]\n", content)
		t.Fail()
	}
}

func readFixture(t *testing.T, path string) string {
	result, err := utils.ReadTextFile(path)
	if err != nil {
		t.Fatalf("Failed to read fixture file: %s", path)
	}

	return result
}
