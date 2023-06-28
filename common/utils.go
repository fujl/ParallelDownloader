package common

import (
	"errors"
	"fmt"
	"os"
	"regexp"
)

const (
	HEADER_RANGE                  string = "Range"
	DEFAULT_HEADER_RANGE_TEMPLATE string = "bytes=%d-%d"
	DEFAULT_HEADER_RANGE          string = "bytes %d-%d"
	DEFAULT_HEADER_RANGE_STARTID  int    = 0
	DEFAULT_HEADER_RANGE_ENDID    int    = 3
)

func IsUrl(url string) error {
	re := regexp.MustCompile("(http|https):\\/\\/[\\w\\-_]+(\\.[\\w\\-_]+)+([\\w\\-\\.,@?^=%&:/~\\+#]*[\\w\\-\\@?^=%&/~\\+#])?")
	if re == nil {
		return errors.New("MustCompile err")
	}

	result := re.FindAllStringSubmatch(url, -1)
	if result == nil {
		return errors.New("target is invalid")
	}

	return nil
}

func IsExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}

func GetHeaderRange(startId, endId int) string {
	return fmt.Sprintf(DEFAULT_HEADER_RANGE_TEMPLATE, startId, endId)
}

func GetDefaultHeaderRange() string {
	return GetHeaderRange(DEFAULT_HEADER_RANGE_STARTID, DEFAULT_HEADER_RANGE_ENDID)
}
