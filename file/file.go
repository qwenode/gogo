package file

import "os"

func WriteFileAppend(filename string, c []byte) error {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	_, err = f.Write(c)
	if err1 := f.Close(); err == nil {
		err = err1
	}

	return nil
}
