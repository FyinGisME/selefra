package apply

import (
	"context"
	"github.com/selefra/selefra/global"
	"testing"
)

func TestApplyOnLine(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
		return
	}
	global.SERVER = "dev-api.selefra.io"
	global.LOGINTOKEN = "4fe8ed36488c479d0ba7292fe09a4132"
	*global.WORKSPACE = "../../tests/workspace/online"
	err := Apply(context.Background())
	if err != nil {
		t.Error(err)
	}
}
