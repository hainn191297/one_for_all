package kq

import (
	"context"
	"strconv"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/protocol"
	"github.com/zeromicro/go-zero/core/executors"
	"github.com/zeromicro/go-zero/core/logx"
	"go.opentelemetry.io/otel/trace"
)

type (
	PushOption func(options *chunkOptions)

	Pusher struct {
		produer  *kafka.Writer
		topic    string
		executor *executors.ChunkExecutor
	}

	Message struct {
		context context.Context
		p       Pusher
		msq     kafka.Message
	}

	chunkOptions struct {
		chunkSize     int
		flushInterval time.Duration
	}
)

func NewPusher(addrs []string, topic string, opts ...PushOption) *Pusher {
	producer := &kafka.Writer{
		Addr:        kafka.TCP(addrs...),
		Topic:       topic,
		Balancer:    &kafka.LeastBytes{},
		Compression: kafka.Snappy,
	}
	pusher := &Pusher{
		produer: producer,
		topic:   topic,
	}
	pusher.executor = executors.NewChunkExecutor(func(tasks []interface{}) {
		chunk := make([]kafka.Message, len(tasks))
		for i := range tasks {
			chunk[i] = tasks[i].(kafka.Message)
		}
		if err := pusher.produer.WriteMessages(context.Background(), chunk...); err != nil {
			logx.Error(err)
		}
	}, newOptions(opts)...)

	return pusher
}

func (p *Pusher) Close() error {
	if p.executor != nil {
		p.executor.Flush()
	}
	return p.produer.Close()
}

func (p *Pusher) Name() string {
	return p.topic
}

func (p *Pusher) Message(v string) *Message {
	return p.defaultMessage(v)
}

func (p *Pusher) defaultMessage(v string) *Message {
	msq := kafka.Message{
		Key:   []byte(strconv.FormatInt(time.Now().UnixNano(), 10)),
		Value: []byte(v),
	}
	return &Message{
		p:   *p,
		msq: msq,
	}
}

func (m *Message) Push() error {
	if m.p.executor != nil {
		return m.p.executor.Add(m.msq, len(m.msq.Value))
	}

	return m.p.produer.WriteMessages(context.Background(), m.msq)
}

func (m *Message) WithContext(ctx context.Context) *Message {
	mq := *m
	if m.context == nil {
		mq.context = context.Background()
	}
	mq.context = ctx
	spanValue := ""
	traceValue := ""

	span := trace.SpanContextFromContext(mq.context)
	if span.HasSpanID() {
		spanValue = span.SpanID().String()
	}

	spanCtx := trace.SpanContextFromContext(mq.context)
	if spanCtx.HasTraceID() {
		traceValue = spanCtx.TraceID().String()
	}

	mq.msq.Headers = []protocol.Header{
		{
			Key:   "X-Span",
			Value: []byte(spanValue),
		},
		{
			Key:   "X-Trace",
			Value: []byte(traceValue),
		},
	}

	return &mq
}

func (m *Message) WithHeader(key string, value string) *Message {
	mq := *m

	header := protocol.Header{
		Key:   key,
		Value: []byte(value),
	}

	mq.msq.Headers = append(mq.msq.Headers, header)

	return &mq
}

func WithChunkSize(chunkSize int) PushOption {
	return func(options *chunkOptions) {
		options.chunkSize = chunkSize
	}
}

func WithFlushInterval(interval time.Duration) PushOption {
	return func(options *chunkOptions) {
		options.flushInterval = interval
	}
}

func newOptions(opts []PushOption) []executors.ChunkOption {
	var options chunkOptions
	for _, opt := range opts {
		opt(&options)
	}

	var chunkOpts []executors.ChunkOption
	if options.chunkSize > 0 {
		chunkOpts = append(chunkOpts, executors.WithChunkBytes(options.chunkSize))
	}
	if options.flushInterval > 0 {
		chunkOpts = append(chunkOpts, executors.WithFlushInterval(options.flushInterval))
	}
	return chunkOpts
}
