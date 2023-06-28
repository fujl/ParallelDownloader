package entity

import "testing"

func TestGetSaveInfo(t *testing.T) {
	var dcp Downloader
	dcp.Out = "./download_cmd_param_test.go"
	err := dcp.getSaveInfo()
	if err != nil {
		t.Error("not as expect")
	}

	if dcp.Option.SaveName != "download_cmd_param_test.go" {
		t.Error("not as expect")
	}
}
