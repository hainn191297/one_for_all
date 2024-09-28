package rabbitmq

import "fmt"

// RabbitConf holds the basic RabbitMQ configuration.
type RabbitConf struct {
	Username string  // Username for RabbitMQ
	Password string  // Password for RabbitMQ
	Host     string  // Host for RabbitMQ
	Port     int     // Port for RabbitMQ
	VHost    *string `json:",omitempty"` // Virtual host (optional)
}

// RabbitListenerConf holds configuration for a RabbitMQ listener.
type RabbitListenerConf struct {
	RabbitConf                    // Embedding RabbitConf for common settings
	ListenerQueues []ConsumerConf // List of consumer configurations
}

// ConsumerConf holds configuration for a RabbitMQ consumer.
type ConsumerConf struct {
	Name      string // Name of the consumer
	AutoAck   bool   // Automatically acknowledge messages
	Exclusive bool   // Exclusive access to this consumer
	NoLocal   bool   // If true, prevents delivery from producers in the same connection
	NoWait    bool   // If true, do not wait for the server response
}

// RabbitSenderConf holds configuration for sending messages to RabbitMQ.
type RabbitSenderConf struct {
	RabbitConf         // Embedding RabbitConf for common settings
	ContentType string `json:",omitempty"` // MIME content type (optional)
}

// QueueConf holds configuration for a RabbitMQ queue.
type QueueConf struct {
	Name       string // Name of the queue
	Durable    bool   // If true, the queue will survive server restarts
	AutoDelete bool   // If true, the queue will be deleted when no consumers are connected
	Exclusive  bool   // If true, the queue can only be used by the current connection
	NoWait     bool   // If true, do not wait for the server response
}

type ExchangeConf struct {
	ExchangeName string      // Name of the exchange
	Type         string      // Exchange type (valid options: direct, fanout, topic, headers)
	Durable      bool        // If true, the exchange will survive server restarts
	AutoDelete   bool        // If true, the exchange will be deleted when no queues are bound
	Internal     bool        // If true, the exchange is internal and cannot be directly accessed by clients
	NoWait       bool        // If true, do not wait for the server response
	Queues       []QueueConf // List of queues bound to this exchange
}

// Validate checks if the Type is one of the valid options.
func (e *ExchangeConf) Validate() error {
	validTypes := map[string]struct{}{
		"direct":  {},
		"fanout":  {},
		"topic":   {},
		"headers": {},
	}

	if _, valid := validTypes[e.Type]; !valid {
		return fmt.Errorf("invalid exchange type: %s", e.Type)
	}
	return nil
}

// getRabbitURL constructs the AMQP URL for connecting to RabbitMQ.
func getRabbitURL(rabbitConf RabbitConf) string {
	vhost := ""
	if rabbitConf.VHost != nil {
		vhost = *rabbitConf.VHost
	}
	return fmt.Sprintf("amqp://%s:%s@%s:%d/%s", rabbitConf.Username, rabbitConf.Password,
		rabbitConf.Host, rabbitConf.Port, vhost)
}
