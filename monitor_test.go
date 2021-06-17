package kafkamonitor_test

import (
	"log"
	"testing"

	"github.com/alecthomas/assert"
	kafkamonitor "github.com/ringsaturn/kafka-monitor"
)

func TestMonitor(t *testing.T) {
	monitor, err := kafkamonitor.NewMonitor("localhost:9092")
	if err != nil {
		panic(err)
	}
	log.Println(monitor)
	defer monitor.Close()

	topics, err := monitor.GetTopics()
	if err != nil {
		panic(err)
	}
	log.Println(topics)
	assert.Equal(t, 1, 2)
}
