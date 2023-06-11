package main

import (
	"encoding/json"
	"go-nats-jetstream-example/config"
	"go-nats-jetstream-example/models"
	"log"

	"github.com/nats-io/nats.go"
)

func consumeReviews(js nats.JetStreamContext) {
	_, err := js.Subscribe(config.SubjectNameReviewCreated, func(m *nats.Msg) {
		err := m.Ack()

		if err != nil {
			log.Println("Unable to Ack", err)
			return
		}

		var employee models.Employee
		err = json.Unmarshal(m.Data, &employee)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Consumer  =>  Subject: %s  -  ID: %s  -  Author: %s  -  Rating: %d\n", m.Subject, employee.EmpId, employee.EmpFName, employee.EmpLName)

		// send answer via JetStream using another subject if you need
		// js.Publish(config.SubjectNameReviewAnswered, []byte(review.Id))
	})

	if err != nil {
		log.Println("Subscribe failed")
		return
	}
}
