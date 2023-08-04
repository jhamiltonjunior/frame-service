package validation_test

import (
	"fmt"
	"testing"

	"github.com/jhamiltonjunior/saas/src/domain/frame/validation"
)

func TestCreateWithFewCharacters(t *testing.T) {
	title := validation.Title{}

	// isValid, err := title.Create("jose")

	if _, err := title.Create("jose"); err == nil {
		t.Errorf(`%q`, err)
	}
}

func TestCreateWithManyCharacters(t *testing.T) {
	title := validation.Title{}

	increment := ""

	for i := 0; i < 254; i++ {
		increment += "c"
	}

	fmt.Println("%a", increment)

	if _, err := title.Create(increment); err == nil {
		t.Errorf(`%q`, err)
	}
}
