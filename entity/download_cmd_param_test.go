package entity

import "testing"

func TestVerify(t *testing.T) {
	dcp := DownloadCmdParam{
		Target:     "test",
		Out:        "./",
		Concurrent: 10,
		Block:      100,
		RetryTimes: 10,
		Timeout:    30,
		ProxyEnv:   "on",
	}
	err := dcp.Verify()
	if err == nil {
		t.Error("not as expect")
	}

	dcp.Target = "http://my.parallel.downloader.com/test"
	err = dcp.Verify()
	if err != nil {
		t.Error("not as expect")
	}

	dcp.Concurrent = 0
	err = dcp.Verify()
	if err == nil {
		t.Error("not as expect")
	}

	dcp.Concurrent = 1001
	err = dcp.Verify()
	if err == nil {
		t.Error("not as expect")
	}

	dcp.Concurrent = 10
	dcp.RetryTimes = 2
	err = dcp.Verify()
	if err == nil {
		t.Error("not as expect")
	}

	dcp.RetryTimes = 101
	err = dcp.Verify()
	if err == nil {
		t.Error("not as expect")
	}

	dcp.RetryTimes = 5
	dcp.Timeout = 0
	err = dcp.Verify()
	if err == nil {
		t.Error("not as expect")
	}

	dcp.Timeout = 1001
	err = dcp.Verify()
	if err == nil {
		t.Error("not as expect")
	}
}
