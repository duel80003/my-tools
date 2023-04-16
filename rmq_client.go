package tools

import (
	"context"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"github.com/apache/rocketmq-client-go/v2/rlog"
)

var (
	p rocketmq.Producer
)

type RmqConfig struct {
	NameServers []string
}

type Rmq struct {
	Producer *rocketmq.Producer
}

type RmqMsg struct {
	Topic string
	Tag   string
	Keys  []string
	Body  []byte
}

func SetRLogLevelToError() {
	rlog.SetLogLevel("error")
}

func InitializePublisher(config *RmqConfig) *Rmq {
	var err error
	p, err = rocketmq.NewProducer(
		producer.WithNameServer(config.NameServers),
		producer.WithRetry(2),
	)
	if err != nil {
		Logger.Fatal("NewProducer error ", err)
	}

	err = p.Start()
	if err != nil {
		Logger.Fatal("start produce error ", err)
	}

	return &Rmq{
		Producer: &p,
	}
}

func (rmq *Rmq) Send(msg *RmqMsg) error {
	message := primitive.NewMessage(msg.Topic, msg.Body)
	message.WithTag(msg.Tag)
	message.WithKeys(msg.Keys)
	_, err := p.SendSync(context.TODO(), message)
	if err != nil {
		Logger.Errorf("RocketMQ send message error: %s", err)
	}
	return err
}

func (rmq *Rmq) SendAsync(msg *RmqMsg) error {
	message := primitive.NewMessage(msg.Topic, msg.Body)
	message.WithTag(msg.Tag)
	message.WithKeys(msg.Keys)
	return p.SendAsync(context.TODO(), func(ctx context.Context, result *primitive.SendResult, err error) {
		if err != nil {
			Logger.Errorf("rmq send async error: %s, message: %+v", err, message)
		}
	}, message)
}
