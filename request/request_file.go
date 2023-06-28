package request

import (
	"context"
	"fmt"
	"net/http"

	"parallel.downloader.com/common"
	"parallel.downloader.com/entity"
)

func CheckFileSupportPartialAndFileName(d *entity.Downloader, ctx *context.Context) error {
	rangeStr := common.GetDefaultHeaderRange()
	resp, err := HttpRequest(ctx, d, HTTP_METHOD_GET, rangeStr)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusPartialContent {
		return fmt.Errorf("failed to get file size statuscode err : %d", resp.StatusCode)
	}

	return nil
}
