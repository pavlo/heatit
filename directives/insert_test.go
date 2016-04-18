package directives

import "testing"

type InsertDirectiveTestData struct {
	Source string
	ExpectedValue string
	ExpectedScheme string
	ExpectedIndent int
}

func TestNewInsertDirective(t *testing.T) {

	assertDirective( &InsertDirectiveTestData{
		Source: "@insert: file://assets/flavors.yaml",
		ExpectedValue: "assets/flavors.yaml",
		ExpectedScheme: "file",
		ExpectedIndent: 0,
	}, t)

	assertDirective( &InsertDirectiveTestData{
		Source: "    @insert: \"file://assets/flavors.yaml\"",
		ExpectedValue: "assets/flavors.yaml",
		ExpectedScheme: "file",
		ExpectedIndent: 4,
	}, t)

	// no explicit scheme defaults to file
	assertDirective( &InsertDirectiveTestData{
		Source: "  @insert: \"assets/flavors.yaml\"",
		ExpectedValue: "assets/flavors.yaml",
		ExpectedScheme: "file",
		ExpectedIndent: 2,
	}, t)

	assertDirective( &InsertDirectiveTestData{
		Source: "@insert: \"http://google.com/q=foobar\"",
		ExpectedValue: "http://google.com/q=foobar",
		ExpectedScheme: "http",
		ExpectedIndent: 0,
	}, t)

	assertDirective( &InsertDirectiveTestData{
		Source: "@insert: \"https://discovery.etcd.io/new?size=3\"",
		ExpectedValue: "https://discovery.etcd.io/new?size=3",
		ExpectedScheme: "https",
		ExpectedIndent: 0,
	}, t)
}

func assertDirective(data *InsertDirectiveTestData, t *testing.T) {
	d, err := NewInsertDirective(data.Source)

	if err != nil {
		t.Fatalf("Failed to create InsertDirective from %s", data.Source)
	}

	if d.Scheme != data.ExpectedScheme {
		t.Fatalf("Wrong scheme for %s", data.Source)
	}

	if d.SourceValue != data.ExpectedValue {
		t.Fatalf("Wrong source value for %s, received: '%s', expected: '%s'", data.Source, d.SourceValue, data.ExpectedValue)
	}

	if d.Indent != data.ExpectedIndent {
		t.Fatalf("Wrong indent value for %s", data.Source)
	}
}

