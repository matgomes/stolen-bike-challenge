package repository

import (
	"database/sql"
	"fmt"
	"testing"
)

func TestHandleID(t *testing.T) {

	id := 0
	expected := sql.NullInt64{
		Int64: 0,
		Valid: false,
	}

	result := getNullableID(id)

	msg := fmt.Sprintf("expected %v, got %v", expected, result)

	if expected == result {
		t.Logf("Success: %s", msg)
	} else {
		t.Errorf("Failed: %s", msg)
	}

}
