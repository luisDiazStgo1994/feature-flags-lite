package broker

type Broker interface {
	PublishEvent() error
}
