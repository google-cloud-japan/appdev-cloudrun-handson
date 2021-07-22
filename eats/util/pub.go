// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package util

import (
	"cloud.google.com/go/pubsub"
	"context"
	"eats.com/model"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

var projectId = os.Getenv("PROJECT_ID")
var topicId = os.Getenv("TOPIC_ID")

func Publish(eventName, purchaser string, orderId, itemId uint){
	ctx := context.Background()

	// make client
	client, err := pubsub.NewClient(ctx, projectId)
	if err != nil {
		fmt.Print("client error.err:",err)
		os.Exit(1)
	}
	// make json message
	event := &model.Event{
	 	EventName: eventName,
		Purchaser: purchaser,
		OrderID:   orderId,
		ItemID:    itemId,
	}
	msg, _ := json.Marshal(event)
	// publish message
	t := client.Topic(topicId)
	result := t.Publish(ctx, &pubsub.Message{
		Data: []byte(msg),
	})
	// error handling
	id, err := result.Get(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Published a message; msg ID: %v\n", id)
}
