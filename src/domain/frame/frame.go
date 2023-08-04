package frame

import (
	"testing"

	"github.com/jhamiltonjunior/saas/src/domain/frame/validation"
)

func Exec(t *testing.T) {
	title := validation.Title{}

	// isValid, err := title.Create("jose")

	if _, err := title.Create("jose"); err != nil {
		t.Fatalf(`%q`, err)
	}
}