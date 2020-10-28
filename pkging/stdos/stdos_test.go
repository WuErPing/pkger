package stdos

import (
	"testing"

	"github.com/WuErPing/pkger/pkging"
	"github.com/WuErPing/pkger/pkging/pkgtest"
)

func Test_Pkger(t *testing.T) {
	pkgtest.All(t, func(ref *pkgtest.Ref) (pkging.Pkger, error) {
		return New(ref.Info)
	})
}
