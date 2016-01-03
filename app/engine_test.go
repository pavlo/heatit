package app

import (
	"github.com/pavlo/heatit/utils"
	"strings"
	"testing"
)

func TestEngineSimpleInsertion(t *testing.T) {

	e := engine(
		"../fixtures/engine/simple_insertion/a.yaml",
		"",
		"",
	)

	content, err := e.processInserts()
	if err != nil {
		t.Fatal("Failed to process inserts!")
	}

	expectedContent, err := utils.ReadTextFile("../fixtures/engine/simple_insertion/result.yaml")
	if err != nil {
		t.Error("Failed to load results YAML file!")
	}

	if content != expectedContent {
		t.Log("Failed to compare results!")
		t.Logf("Expected: \n[%v]\n", expectedContent)
		t.Logf("Actual: \n[%v]\n", content)
		t.Fail()
	}

}

func TestEngineRecursiveInsertion(t *testing.T) {

	e := engine(
		"../fixtures/engine/recursive_insertion/a.yaml",
		"",
		"",
	)

	content, err := e.processInserts()
	if err != nil {
		t.Fatal("Failed to process inserts!")
	}

	expectedContent, err := utils.ReadTextFile("../fixtures/engine/recursive_insertion/result.yaml")
	if err != nil {
		t.Error("Failed to load results YAML file!")
	}

	if strings.TrimSpace(content) != strings.TrimSpace(expectedContent) {
		t.Log("Failed to compare results!")
		t.Logf("Expected: \n[%v]\n", expectedContent)
		t.Logf("Actual: \n[%v]\n", content)
		t.Fail()
	}

}

func TestEngineComplete(t *testing.T) {

	e := engine(
		"../fixtures/engine/complete/heat.yaml",
		"",
		"",
	)

	content, err := e.processInserts()
	if err != nil {
		t.Fatal("Failed to process inserts!")
	}

	expectedContent, err := utils.ReadTextFile("../fixtures/engine/complete/result.yaml")
	if err != nil {
		t.Error("Failed to load results YAML file!")
	}

	if strings.TrimSpace(content) != strings.TrimSpace(expectedContent) {
		t.Log("Failed to compare results!")
		t.Logf("Expected: \n[%v]\n", expectedContent)
		t.Logf("Actual: \n[%v]\n", content)
		t.Fail()
	}

}
