package rekognition

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/rekognition"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
)

const region = "ap-southeast-1"

var reko_agnet *rekognition.Client

func CompareFaces(source, target []byte) (*rekognition.CompareFacesOutput, error) {
	//
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		return nil, err
	}

	reko_agnet = rekognition.NewFromConfig(cfg)
	return reko_agnet.CompareFaces(context.TODO(), &rekognition.CompareFacesInput{
		SourceImage: &types.Image{
			Bytes: source,
		},
		TargetImage: &types.Image{
			Bytes: target,
		},
	})
}
