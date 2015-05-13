package main

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/dronemill/eventsocket-client-go"
)

// Client is the main ES client of the Maestro
type Client struct {
	// Portal is the portal client
	Portal *eventsocketclient.Client
}

// NewClient returns a new, connected Portal client
func NewClient() *Client {
	portal, err := eventsocketclient.NewClient(fmt.Sprintf("127.0.0.1:%d", config.Eventsocket.Port))
	if err != nil {
		log.WithField("error", err.Error()).
			Fatal("Failed creating Portal client")
	}

	if err := portal.DialWs(); err != nil {
		log.WithField("error", err.Error()).
			Fatal("Failed dialing ws")
	}
	log.Info("Successfully Dialed WS")

	portal.SetMaxMessageSize(5242880) // 5MB
	log.WithField("size", 5242880).Debug("Set ES server max message size")

	log.WithField("clientID", portal.Id).Info("Connected to portal")

	client := &Client{
		Portal: portal,
	}

	return client
}

func (c *Client) run() {
	log.WithField("clientID", c.Portal.Id).Info("Running portal client")

	go c.Portal.Recv()

	portalBootChan, err := c.Portal.Suscribe("batond_boot")
	if err != nil {
		log.WithField("error", err.Error()).
			WithField("event", "batond_boot").
			Fatal("Failed suscribing to event")
	}
	log.WithField("event", "batond_boot").
		Info("Suscribed to event")

	for {
		select {
		case r := <-portalBootChan:
			c.handleBatondBoot(r)
		}
	}
}

func (c *Client) handleBatondBoot(r *eventsocketclient.Received) {
	clientID := (*r.Message.Payload)["ClientID"].(string)
	machineID := (*r.Message.Payload)["MachineID"].(string)

	log.WithField("clientID", c.Portal.Id).
		WithField("clientID", clientID).
		WithField("machineID", machineID).
		Info("A Batond has booted!")
}
