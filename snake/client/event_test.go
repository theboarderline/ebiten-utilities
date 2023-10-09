package client_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/theboarderline/ebiten-utilities/snake/client"
	"github.com/theboarderline/ebiten-utilities/snake/events"
)

var _ = Describe("Events", func() {

	var (
		in               chan events.Event
		out              chan events.Event
		gameserverClient *client.GameserverClient
	)

	BeforeEach(func() {
		in = make(chan events.Event, 10)
		out = make(chan events.Event, 10)

		gameserverClient = client.NewGameserverClient(0, nil, &client.MockConn{}, in, out)
		Expect(gameserverClient).NotTo(BeNil())
	})

	AfterEach(func() {
		err := gameserverClient.Cleanup()
		Expect(err).NotTo(HaveOccurred())
	})

	It("can receive an incoming event", func() {
		event := events.Event{
			PlayerName:     events.ENEMY,
			Type:           events.PLAYER_INPUT,
			InputDirection: events.NewRandomDirection(),
		}
		in <- event

		receivedEvent := gameserverClient.GetMessage()

		Expect(receivedEvent).NotTo(BeNil())
		assertEventsEqual(*receivedEvent, event)
	})

	It("can send an outgoing event", func() {
		event := events.Event{
			PlayerName:     events.ENEMY,
			Type:           events.PLAYER_INPUT,
			InputDirection: events.NewRandomDirection(),
		}

		gameserverClient.SendMessage(&event)

		retrievedEvent := <-out
		assertEventsEqual(retrievedEvent, event)
	})

})

func assertEventsEqual(receivedEvent, expectedEvent events.Event) {
	Expect(receivedEvent).NotTo(BeNil())
	Expect(receivedEvent.PlayerName).To(Equal(expectedEvent.PlayerName))
	Expect(receivedEvent.Type).To(Equal(expectedEvent.Type))
	Expect(receivedEvent.InputDirection).To(Equal(expectedEvent.InputDirection))
}
