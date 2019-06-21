package services

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/service/sns"
)

var typeJSON = "json"

type PushService struct {
	publisher *sns.SNS
	topicArn  string
}

func NewPushService(topicArn string, publisher *sns.SNS) *PushService {
	return &PushService{
		publisher: publisher,
		topicArn:  topicArn,
	}
}

func (p *PushService) PushMessage(message interface{}) (string, error) {
	m, err := json.Marshal(message)
	if err != nil {
		return "", fmt.Errorf("error mashalling message: %v", err)
	}
	mString := string(m)
	pushInput := &sns.PublishInput{
		Message:          &mString,
		MessageStructure: &typeJSON,
		TopicArn:         &p.topicArn,
	}

	pushOutput, err := p.publisher.Publish(pushInput)
	if err != nil {
		return "", fmt.Errorf("error publishing SNS message: %v", err)
	}

	return *pushOutput.MessageId, nil
}
