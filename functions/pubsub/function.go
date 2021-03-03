// Package function provides a set of Cloud Functions samples.
package function

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/kelseyhightower/envconfig"
)

type Env struct {
	ProjectID              string `required:"true" split_words:"true"`
	DBUser                 string `required:"true" split_words:"true"`
	DBPass                 string `required:"true" split_words:"true"`
	DBSchema               string `required:"true" split_words:"true"`
	DBFormat               string `required:"true" split_words:"true"`
	InstanceConnectionName string `required:"true" split_words:"true"`
}

var env Env

// PubSubMessage is the payload of a Pub/Sub event.
type PubSubMessage struct {
	Data []byte `json:"data"`
}

func init() {
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.JSONFormatter{
		FieldMap: log.FieldMap{
			log.FieldKeyTime:  "time",
			log.FieldKeyLevel: "severity",
			log.FieldKeyMsg:   "message",
		},
		TimestampFormat: time.RFC3339Nano,
	})
	if err := envconfig.Process("", &env); err != nil {
		log.Fatal("fail to load config wiht env :", err)
	}
}

// HelloPubSub consumes a Pub/Sub message.
func HelloPubSub(ctx context.Context, m PubSubMessage) error {
	name := string(m.Data)
	if name == "" {
		name = "World"
	}
	log.Printf("Hello, %s!", name)
	return nil
}
