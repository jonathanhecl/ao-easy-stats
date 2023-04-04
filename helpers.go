package main

import "os"

// OpenFile opens a file and returns the file pointer.
func OpenFile(path string) (*os.File, error) {

	file, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {

		return nil, err
	}

	return file, nil
}

// PathExists returns true if the file or folder exists.
func PathExists(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}

	return true
}

// CreateFolder creates a folder.
func CreateFolder(path string) error {
	return os.Mkdir(path, 0755)
}

// ArrayStringContains returns true if the value is in the list.
func ArrayStringContains(list []string, value string) bool {
	found := false

	for _, s := range list {
		if s == value {
			found = true
			break
		}
	}

	return found
}

// ArrayStringAddOnce adds a value to the list if it is not already in the list.
func ArrayStringAddOnce(list []string, value string) []string {
	if !ArrayStringContains(list, value) {
		list = append(list, value)
	}

	return list
}

// ArrayStringRemove removes a value from the list.
func ArrayStringRemove(list []string, value string) []string {
	var ret []string

	for _, s := range list {
		if s != value {
			ret = append(ret, s)
		}
	}

	return ret
}
