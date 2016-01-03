package app

import (
	"github.com/pavlo/heatit/utils"
	"strings"
	"testing"
)

func TestEngineProcessInsertsSimpleInsertion(t *testing.T) {
	data := readFixture(t, "../fixtures/engine/simple_insertion/a.yaml")

	content, err := processInserts(data, 0)
	if err != nil {
		t.Fatal("Failed to process inserts!")
	}

	expectedContent := readFixture(t, "../fixtures/engine/simple_insertion/result.yaml")
	if content != expectedContent {
		t.Log("Failed to compare results!")
		t.Logf("Expected: \n[%v]\n", expectedContent)
		t.Logf("Actual: \n[%v]\n", content)
		t.Fail()
	}

}

func TestEngineProcessInsertsRecursiveInsertion(t *testing.T) {
	data := readFixture(t, "../fixtures/engine/recursive_insertion/a.yaml")

	content, err := processInserts(data, 0)
	if err != nil {
		t.Fatal("Failed to process inserts!")
	}

	expectedContent := readFixture(t, "../fixtures/engine/recursive_insertion/result.yaml")
	if strings.TrimSpace(content) != strings.TrimSpace(expectedContent) {
		t.Log("Failed to compare results!")
		t.Logf("Expected: \n[%v]\n", expectedContent)
		t.Logf("Actual: \n[%v]\n", content)
		t.Fail()
	}

}

func TestEngineProcessInsertsComplete(t *testing.T) {
	data := readFixture(t, "../fixtures/engine/complete/heat.yaml")

	content, err := processInserts(data, 0)
	if err != nil {
		t.Fatal("Failed to process inserts!")
	}

	expectedContent := readFixture(t, "../fixtures/engine/complete/result.yaml")
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
