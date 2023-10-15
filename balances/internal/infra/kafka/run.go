package kafka

import (
	"encoding/json"
	"fmt"

	"github.com.br/Leodf/walletcore/balances/internal/app/factories"
	"github.com.br/Leodf/walletcore/balances/internal/domain/usecase"
	"github.com.br/Leodf/walletcore/balances/internal/infra/env"
	ckafka "github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type KafkaMsg struct {
	Name    string `json:"name"`
	Payload struct {
		AccountIdFrom        string  `json:"accountIdFrom"`
		AccountIdTo          string  `json:"accountIdTo"`
		BalanceAccountIdFrom float64 `json:"balanceAccountIdFrom"`
		BalanceAccountIdTo   float64 `json:"balanceAccountIdTo"`
	}
}

func RunKafka() {
	configMap := ckafka.ConfigMap{
		"bootstrap.servers": env.GetEnv("KAFKA_SERVER"),
		"group.id":          env.GetEnv("KAFKA_CONSUMER_GROUP_ID"),
	}
	kafkaMsgChan := make(chan *ckafka.Message)
	KafkaConsumer := NewConsumer(&configMap, []string{env.GetEnv("KAFKA_TOPIC")})
	go KafkaConsumer.Consume(kafkaMsgChan)
	go func() {
		for msg := range kafkaMsgChan {
			kafkaMsg := KafkaMsg{}
			err := json.Unmarshal(msg.Value, &kafkaMsg)
			if err != nil {
				panic(err)
			}
			uc := factories.MakeKafkaHandle()
			input := usecase.CreateBalancesInput{
				AccountID:   kafkaMsg.Payload.AccountIdFrom,
				AccountID_2: kafkaMsg.Payload.AccountIdTo,
				Amount:      kafkaMsg.Payload.BalanceAccountIdFrom,
				Amount_2:    kafkaMsg.Payload.BalanceAccountIdTo,
			}
			err = uc.Execute(input)
			if err != nil {
				fmt.Println("KAFKA CONSUMER", "msg", err.Error())
			} else {
				fmt.Println("KAFKA CONSUMER", kafkaMsg.Name, kafkaMsg.Payload)
			}
		}
	}()
}
