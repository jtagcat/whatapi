package whatapi_test

import (
	"testing"

	"github.com/charles-haynes/whatapi"
)

func TestFiles(t *testing.T) {
	to := whatapi.TorrentStruct{}
	f, err := to.Files()
	if f == nil || len(f) != 0 || err != nil {
		t.Errorf("expected Files of null TorrentStruct to return empty list, got %v, %s", f, err)
	}
	to = whatapi.TorrentStruct{FileList: ""}
	f, err = to.Files()
	if f == nil || len(f) != 0 || err != nil {
		t.Errorf("expected Files  of empty file list to return empty list, got %v, %s", f, err)
	}
	to = whatapi.TorrentStruct{FileList: "bad"}
	f, err = to.Files()
	if err == nil {
		t.Errorf("expected Files of bad file list to return error, got %v, %s", f, err)
	}
	to = whatapi.TorrentStruct{FileList: "|||a{{{1}}}"}
	f, err = to.Files()
	if err == nil {
		t.Errorf("expected Files of bad file list to return error, got %v, %s", f, err)
	}
	to = whatapi.TorrentStruct{FileList: "{{{}}}"}
	f, err = to.Files()
	if err == nil {
		t.Errorf("expected Files of bad file list to return error, got %v, %s", f, err)
	}
	exp := []struct {
		Name string
		Size int64
	}{
		{Name: "aaa", Size: 123},
		{Name: "bbb", Size: 456},
		{Name: "ccc", Size: 789},
	}
	to = whatapi.TorrentStruct{
		FileList: "aaa{{{123}}}|||bbb{{{456}}}|||ccc{{{789}}}",
	}
	f, err = to.Files()
	if err != nil {
		t.Errorf("Files returned an error: %s", err)
	}
	if len(exp) != len(f) {
		t.Errorf("Expected to get %d results but got %d", len(exp), len(f))
	}
	for i, v := range exp {
		if i >= len(f) {
			break
		}
		if v.Name != f[i].Name() {
			t.Errorf(`Expected f[%d].Name = "%s" but got "%s"`,
				i, v.Name, f[i].Name())
		}
		if v.Size != f[i].Size {
			t.Errorf(`Expected f[%d].Size = %d but got %d`,
				i, v.Size, f[i].Size)
		}
	}
}
