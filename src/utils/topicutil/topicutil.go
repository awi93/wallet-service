package topicutil

import (
	"log"

	"github.com/lovoo/goka"
)

// EnsureStreamExists is a convenience wrapper for TopicManager.EnsureStreamExists
func EnsureStreamExists(topic string, brokers []string) {
	tm := createTopicManager(brokers)
	defer tm.Close()
	err := tm.EnsureStreamExists(topic, 8)
	if err != nil {
		log.Printf("Error creating kafka topic %s: %v", topic, err)
	}
}

// EnsureTableExists is a convenience wrapper for TopicManager.EnsureTableExists
func EnsureTableExists(topic string, brokers []string) {
	tm := createTopicManager(brokers)
	defer tm.Close()
	err := tm.EnsureTableExists(topic, 8)
	if err != nil {
		log.Printf("Error creating kafka topic %s: %v", topic, err)
	}
}

func createTopicManager(brokers []string) goka.TopicManager {
	tmc := goka.NewTopicManagerConfig()
	tmc.Stream.Replication = 1
	tmc.Table.Replication = 1
	tm, err := goka.NewTopicManager(brokers, goka.DefaultConfig(), tmc)
	if err != nil {
		log.Fatalf("Error creating topic manager: %v", err)
	}
	return tm
}
