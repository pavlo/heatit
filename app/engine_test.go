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

func TestProcessYamlParams(t *testing.T) {
	assertParams(
		t,
		"../fixtures/parameters/yaml/heat.yaml",
		"../fixtures/parameters/yaml/result.yaml",
	)
}

func TestProcessTextParams2(t *testing.T) {
	assertParams(
		t,
		"../fixtures/parameters/text/source.txt",
		"../fixtures/parameters/text/result.txt",
	)
}

func assertParams(t *testing.T, sourceFile string, expectedResultFile string) {
	data := readFixture(t, sourceFile)
	p, err := NewParameters("../fixtures/parameters/params.yaml")

	if err != nil {
		t.Fatal("Failed to read params fixture!")
	}

	actual, err := processParams(data, p)
	if err != nil {
		t.Fatalf("Failed to process parameters")
	}

	expected := readFixture(t, expectedResultFile)
	assertStringsEqual(t, expected, actual)
}

func assertInserts(t *testing.T, sourceFile string, expectedResultFile string) {
	data := readFixture(t, sourceFile)

	actual, err := processInserts(data, 0)
	if err != nil {
		t.Fatal("Failed to process inserts!")
	}

	expected := readFixture(t, expectedResultFile)
	assertStringsEqual(t, expected, actual)
}

func readFixture(t *testing.T, path string) string {
	result, err := utils.ReadTextFile(path)
	if err != nil {
		t.Fatalf("Failed to read fixture file: %s", path)
	}

	return result
}

func assertStringsEqual(t *testing.T, expected, actual string) {
	if strings.TrimSpace(actual) != strings.TrimSpace(expected) {
		t.Log("Failed to compare strings!")
		t.Logf("Expected: \n[%v]\n", expected)
		t.Logf("Actual: \n[%v]\n", actual)
		t.Fail()
	}
}
