package directives

import "testing"

func TestNewInsertDirective(t *testing.T) {
	assertDirective("    @insert: \"file:assets/flavors.yaml\"", t)
	assertDirective("    @insert : \"file:assets/flavors.yaml\"", t)
}

func assertDirective(s string, t *testing.T) {
	d, err := NewInsertDirective(s)
	if err != nil {
		t.Fatalf("Failed to create InsertDirective from %s", s)
	}

	if d.SourceType != INSERT_DIRECTIVE_TYPE_FILE {
		t.Fatalf("Wrong source type for %s", s)
	}

	if d.SourceValue != "assets/flavors.yaml" {
		t.Fatalf("Wrong source value for %s", s)
	}

	if d.Indent != 4 {
		t.Fatalf("Wrong ident value for %s", s)
	}
}
