package entity

import (
	"errors"
	"net/http"
	"os"
	"path/filepath"
	"regexp"

	"parallel.downloader.com/common"
)

type Downloader struct {
	DownloadCmdParam
	Size   int
	Handle *os.File
	Option DownloadOption
	Queue  []*DownloadQueue
}

type DownloadOption struct {
	SaveName  string
	SavePath  string
	ProxyHost string
}

type DownloadQueue struct {
	Start int
	End   int
	Type  QueueType
}

type QueueType int

const (
	QUEUE_TYPE_NONE        QueueType = 0
	QUEUE_TYPE_DOWNLOADING QueueType = 1
	QUEUE_TYPE_FINISH      QueueType = 2
)

func (d *Downloader) Preload() error {
	if d.Out != DEFAULT_OUT {
		d.getSaveInfo()
	}
	return nil
}

func (d *Downloader) getSaveInfo() error {
	absPath, err := filepath.Abs(d.Out)
	if err != nil {
		return err
	}

	dir := filepath.Dir(absPath)
	exist, err := common.IsExists(dir)
	if err != nil {
		return err
	}

	if !exist {
		return errors.New("dir is't exist")
	}

	d.Option.SavePath = dir
	d.Option.SaveName = filepath.Base(absPath)

	return nil
}

func (d *Downloader) matchFilename(disposition string) error {
	re := regexp.MustCompile(`(?m)filename="(.*)"`)
	if re == nil {
		return errors.New("failed to match filename")
	}

	list := re.FindAllStringSubmatch(disposition, -1)
	if len(list) <= 0 {
		return errors.New("failed to check list")
	}

	if len(list[0]) < 1 {
		return errors.New("failed to check list 0")
	}

	d.Option.SaveName = list[0][1]

	return nil

}

func (d *Downloader) SetDownloadFileInfo(header http.Header) {
	disposition := header.Get("Content-Disposition")
	if disposition != "" {

	}
}
