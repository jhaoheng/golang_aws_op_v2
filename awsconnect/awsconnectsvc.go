package awsconnect

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/connect"
)

type IAWSConnectSvc interface {
	StartOutboundVoiceContact(instanceID, contactFlowID, destinationPhoneNumber string) (output *connect.StartOutboundVoiceContactOutput, err error)
}

type AWSConnectSvc struct {
	Region string
}

func NewAWSConnectSvc(region string) IAWSConnectSvc {
	return &AWSConnectSvc{
		Region: region,
	}
}

/*
- instanceID: Amazon Connect Instance 的 ID（從 AWS 控制台中獲取）
- contactFlowID: 聯絡流程 ID
- destinationPhoneNumber: 客戶的電話號碼, 國際格式, ex: +11234567890
*/
func (svc *AWSConnectSvc) StartOutboundVoiceContact(instanceID, contactFlowID, destinationPhoneNumber string) (output *connect.StartOutboundVoiceContactOutput, err error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(svc.Region))
	if err != nil {
		return
	}

	// 建立 Connect 客戶端
	client := connect.NewFromConfig(cfg)

	// 撥打電話
	output, err = client.StartOutboundVoiceContact(
		context.TODO(),
		&connect.StartOutboundVoiceContactInput{
			InstanceId:             &instanceID,
			ContactFlowId:          &contactFlowID,
			DestinationPhoneNumber: &destinationPhoneNumber,
			SourcePhoneNumber:      nil, // 可選，指定呼出顯示的電話號碼
			Attributes: map[string]string{
				"CustomAttributeKey": "CustomAttributeValue", // 可傳遞自定義參數到聯絡流程
			},
		},
	)
	if err != nil {
		err = fmt.Errorf("[aws connect]撥打失敗, err: %v", err)
		return
	}

	// fmt.Printf("成功撥打電話，聯絡 ID: %s\n", *output.ContactId)
	return
}
