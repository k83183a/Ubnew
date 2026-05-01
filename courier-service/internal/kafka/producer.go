package kafka

type Producer struct{}

func NewProducer(brokers string) (*Producer, error) {
    return &Producer{}, nil
}
