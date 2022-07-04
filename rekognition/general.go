package rekognition

import "io/ioutil"

func get_image_byte(path string) []byte {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return b
}
