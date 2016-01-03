package directives

import "testing"

func TestNewParamDirective(t *testing.T) {
	s := "@param:foobar"
	d, err := NewParamDirective(s)
	if err != nil {
		t.Fatalf("Failed to create ParamDirective from %s", s)
	}

	if d.Name != "foobar" {
		t.Fatalf("Value is wrong for %s", s)
	}
}

func TestNewInvalidSourceForParamDirective(t *testing.T) {
	s := "To be or not to be?"

	_, err := NewParamDirective(s)
	if err == nil {
		t.Fatalf("Expected to get an error!")
	}
}
