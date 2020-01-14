package transport

import (
	flowmessage "github.com/cloudflare/goflow/pb"
	"github.com/cloudflare/goflow/utils"
)

type PromState struct {
	ch chan *flowmessage.FlowMessage
}

func RegisterFlags() {
	// KafkaTLS = flag.Bool("kafka.tls", false, "Use TLS to connect to Kafka")
}

func StartPromProducerFromArgs(flowChan chan *flowmessage.FlowMessage, log utils.Logger) (*PromState, error) {
	state := &PromState{
		ch: flowChan,
	}
	return state, nil
}

func (s PromState) Publish(msgs []*flowmessage.FlowMessage) {
	for _, msg := range msgs {
		s.ch <- msg
	}
}
