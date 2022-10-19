package util

func FailOnError(err error) {
	if err != nil {
		panic("Fatal Error")
	}
}
