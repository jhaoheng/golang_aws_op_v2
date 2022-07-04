package rekognition

import (
	"errors"
	"log"
	"testing"

	"github.com/aws/smithy-go"
	"github.com/stretchr/testify/assert"
)

func Test_CompareFaces_Error(t *testing.T) {

	source_path := "./images/image_1.jpg"
	target_path := "./images/blur_2.jpg"

	_, err := CompareFaces(get_image_byte(source_path), get_image_byte(target_path))
	assert.NoError(t, err)
}

func Test_CompareFaces_OperationError(t *testing.T) {
	source_path := "./images/image_1.jpg"
	// target_path := "./images/blur_2.jpg"

	_, err := CompareFaces(get_image_byte(source_path), nil)

	// Service Client Errors
	if err != nil {
		var oe *smithy.OperationError
		if errors.As(err, &oe) {
			log.Printf("failed to call service: %s, operation: %s, error: %v", oe.Service(), oe.Operation(), oe.Unwrap())
		}
		return
	}
}

func Test_CompareFaces_APIError(t *testing.T) {

	source_path := "./images/image_1.jpg"
	// target_path := "./images/blur_2.jpg"

	_, err := CompareFaces(get_image_byte(source_path), nil)

	//
	if err != nil {
		var ae smithy.APIError
		if errors.As(err, &ae) {
			log.Printf("code: %s, message: %s, fault: %s", ae.ErrorCode(), ae.ErrorMessage(), ae.ErrorFault().String())
		}
		return
	}
}
