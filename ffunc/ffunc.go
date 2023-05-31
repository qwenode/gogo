package ffunc

// Retry auto retry when fun got an error
func Retry(callable func() error, retries int) error {
	var err error
	if retries <= 0 {
		retries = 1
	}
	for i := 0; i < retries; i++ {
		err = callable()
		if err != nil {
			continue
		}
		break
	}
	return nil
}
