package ff

import (
	"github.com/qwenode/gogo/hashes"
	"io/ioutil"
	"os"
)

// CreateTempFileWithString create the temporary file with given string
func CreateTempFileWithString(str string) (*os.File, error) {
	tempFile, err := ioutil.TempFile(os.TempDir(), hashes.Sha1(str))
	if err != nil {
		return nil, err
	}
	_, err = tempFile.WriteString(str)
	if err != nil {
		return nil, err
	}
	return tempFile, nil
}
