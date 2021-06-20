package cobertura

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/bvobart/gocover-cobertura/gocover"
)

func TestParseProfileDoesntExist(t *testing.T) {
	v := Coverage{}
	profile := gocover.Profile{FileName: "does-not-exist"}
	err := v.parseProfile(&profile)
	if err == nil || !strings.Contains(err.Error(), `can't find "does-not-exist"`) {
		t.Fatalf("Expected \"can't find\" error; got: %+v", err)
	}
}

func TestParseProfileNotReadable(t *testing.T) {
	v := Coverage{}
	profile := gocover.Profile{FileName: os.DevNull}
	err := v.parseProfile(&profile)
	if err == nil || !strings.Contains(err.Error(), `expected 'package', found 'EOF'`) {
		t.Fatalf("Expected \"expected 'package', found 'EOF'\" error; got: %+v", err)
	}
}

func TestParseProfilePermissionDenied(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "not-readable")
	defer os.Remove(tmpfile.Name())
	tmpfile.Chmod(000)
	v := Coverage{}
	profile := gocover.Profile{FileName: tmpfile.Name()}
	err = v.parseProfile(&profile)
	if err == nil || !strings.Contains(err.Error(), `permission denied`) {
		t.Fatalf("Expected \"permission denied\" error; got: %+v", err)
	}
}
