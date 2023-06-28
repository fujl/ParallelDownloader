package entity

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/urfave/cli"
	"parallel.downloader.com/common"
)

type DownloadCmdParam struct {
	Target     string
	Out        string
	Concurrent int
	Block      int
	RetryTimes int
	Timeout    int
	ProxyEnv   string
	Headers    map[string]string
}

const DEFAULT_OUT = "./"
const RETRY_TIMES = 10
const TIMEOUT = 30

func (dcp *DownloadCmdParam) ExecCmd() error {
	app := cli.NewApp()
	app.Name = "ParallelDownloader"
	app.Version = "v1.0.0"

	dcp.Headers = make(map[string]string)

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "url, u",
			Usage:       "download target url",
			Destination: &dcp.Target,
			Required:    true,
		},
		cli.StringFlag{
			Name:        "out, o",
			Usage:       "download out location, default is ./",
			Value:       "./",
			Destination: &dcp.Out,
		},
		cli.IntFlag{
			Name:        "concurrent, c",
			Usage:       "download max concurrent number, default is 20, [1, 1000]",
			Value:       20,
			Destination: &dcp.Concurrent,
		},
		cli.IntFlag{
			Name:        "block, b",
			Usage:       "download block size, default is 1MB",
			Value:       1024 * 1024,
			Destination: &dcp.Block,
		},
		cli.IntFlag{
			Name:        "retrytimes, r",
			Usage:       "download retry times, default is 10",
			Value:       RETRY_TIMES,
			Destination: &dcp.RetryTimes,
		},
		cli.IntFlag{
			Name:        "timeout, t",
			Usage:       "download timeout, default is 30",
			Value:       TIMEOUT,
			Destination: &dcp.Timeout,
		},
		cli.StringFlag{
			Name:        "proxy, p",
			Usage:       "download proxy from env, default is false [on/off]",
			Value:       "on",
			Destination: &dcp.ProxyEnv,
		},
		cli.StringSliceFlag{
			Name:  "header",
			Usage: "download header, example --header Content-Length:1024",
		},
	}
	app.Action = func(ctx *cli.Context) error {
		header := ctx.String("header")
		headers := strings.Split(header, ",")
		for _, v := range headers {
			kv := strings.Split(v, ":")
			if len(kv) != 2 {
				log.Fatalf("%s's format invalid", v)
				return errors.New("header is invalid")
			}
			dcp.Headers[kv[0]] = kv[1]
		}
		log.Printf("%v", dcp.Headers)
		return nil
	}

	cli.HelpFlag = cli.BoolFlag{
		Name:  "help, h",
		Usage: "to offer help",
	}

	cli.VersionFlag = cli.BoolFlag{
		Name:  "version, v",
		Usage: "print version",
	}

	err := app.Run(os.Args)
	if err != nil {
		return err
	}

	return dcp.Verify()
}

func (dcp *DownloadCmdParam) Verify() error {
	err := common.IsUrl(dcp.Target)
	if err != nil {
		return errors.New("target is invalid")
	}

	absPath, err := filepath.Abs(dcp.Out)
	if err != nil {
		return err
	}

	dir := filepath.Dir(absPath)
	exist, err := common.IsExists(dir)
	if err != nil {
		return err
	}

	if !exist {
		return errors.New("out dir isn't exist")
	}

	if dcp.Concurrent < 1 || dcp.Concurrent > 1000 {
		return errors.New("concurrent is invalid, [1, 1000]")
	}

	if dcp.RetryTimes < 3 || dcp.RetryTimes > 100 {
		return errors.New("try times is invalid, [3, 100]")
	}

	if dcp.Timeout < 30 || dcp.Timeout > 1000 {
		return errors.New("timeout is invalid, [30, 1000]")
	}

	if dcp.ProxyEnv != "on" && dcp.ProxyEnv != "off" && dcp.ProxyEnv != "ON" && dcp.ProxyEnv != "OFF" {
		return errors.New("proxyenv is invalid,[on/off],[ON/OFF]")
	}

	return nil
}
