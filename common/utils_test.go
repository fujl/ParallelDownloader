package common

import "testing"

func TestIsUrl(t *testing.T) {
	err := IsUrl("test")
	if err == nil {
		t.Error("test is url")
	}

	err = IsUrl("http://my.parallel.downloader.com/test")
	if err != nil {
		t.Error("http://my.parallel.downloader.com/test is't url")
	}
}

func TestIsExists(t *testing.T) {
	exist, err := IsExists("./")
	if err != nil || !exist {
		t.Error("./ is exist")
	}
}

func TestGetHeaderRange(t *testing.T) {
	startId, endId := 0, 3
	if GetHeaderRange(startId, endId) != "bytes=0-3" {
		t.Error("not as expect")
	}

	if GetDefaultHeaderRange() != "bytes=0-3" {
		t.Error("not as expect")
	}
}
