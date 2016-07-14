package stylist

import "testing"

func TestEquality(t *testing.T) {
	var Brian = Stylist{"1", "Brian", "I've been cutting hair for ages"}

	if Brian.Name != "Brian" {
		t.Error("expected Brian")
	}
}
