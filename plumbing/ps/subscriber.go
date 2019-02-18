package ps

import "gx/ipfs/QmVRxA4J3UPQpw74dLrQ6NJkfysCA1H4GU28gVpXQt9zMU/go-libp2p-pubsub"

// Subscriber subscribes to pubsub topics
type Subscriber struct {
	pubsub *pubsub.PubSub
}

// NewSubscriber builds a new subscriber
func NewSubscriber(sub *pubsub.PubSub) *Subscriber {
	return &Subscriber{pubsub: sub}
}

// Subscribe subscribes to a pubsub topic
func (s *Subscriber) Subscribe(topic string, opts ...pubsub.SubOpt) (*pubsub.Subscription, error) {
	return s.pubsub.Subscribe(topic, opts...)
}
