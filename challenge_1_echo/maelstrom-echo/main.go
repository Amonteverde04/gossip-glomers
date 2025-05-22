package main

import (
	"encoding/json"
	"log"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func main() {
	n := maelstrom.NewNode()

	// Unmarshal message body as a loosely typed map.
	n.Handle("echo", func(msg maelstrom.Message) error {
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		// Update return message.
		body["type"] = "echo_ok"

		// Echo original messge back with updated message type.
		return n.Reply(msg, body)
	})

	// Execute node's message loop. Will return until STDIN is closed.
	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}
