package utils

import "io/ioutil"

func ReadFile(fileName string) ([]byte, error) {
	bytes, err := ioutil.ReadFile(fileName)
	CheckError(err)
	return bytes, nil
}
