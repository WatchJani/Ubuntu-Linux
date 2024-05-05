package reader

import "os"

func Read(path string) {
	os.ReadFile(path)
}
