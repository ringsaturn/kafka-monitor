package kafkamonitor

import (
	"github.com/segmentio/kafka-go"
)

type Monitor struct {
	conn *kafka.Conn
}

func NewMonitor(uri string) (*Monitor, error) {
	conn, err := kafka.Dial("tcp", uri) // "localhost:9092")
	if err != nil {
		return nil, err
	}
	m := &Monitor{
		conn: conn,
	}
	return m, nil
}

func (m *Monitor) Close() {
	if m.conn != nil {
		m.conn.Close()
	}
}

type AllTopics map[string]struct{}

func (m *Monitor) GetTopics() (AllTopics, error) {
	partitions, err := m.conn.ReadPartitions()
	if err != nil {
		return nil, err
	}

	topics := make(AllTopics)

	for _, p := range partitions {
		topics[p.Topic] = struct{}{}
	}
	return topics, nil
}
