package client_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/theboarderline/ebiten-utilities/snake/client"
	"github.com/theboarderline/ebiten-utilities/snake/events"
)

var _ = Describe("Udp Integration", func() {

	var (
		in               chan events.Event
		out              chan events.Event
		gameserverClient *client.GameserverClient
	)

	BeforeEach(func() {
		in = make(chan events.Event, 1)
		out = make(chan events.Event, 1)

		gameserverClient = client.NewGameserverClient(0, nil, &client.MockConn{}, in, out)
		Expect(gameserverClient).NotTo(BeNil())
		Expect(gameserverClient.IsConnected()).To(BeTrue())

		defer gameserverClient.Cleanup()
		go gameserverClient.HandleOutgoingEvents()
		go gameserverClient.HandleIncomingEvents()
	})

	AfterEach(func() {
		err := gameserverClient.Cleanup()
		Expect(err).NotTo(HaveOccurred())
		Expect(gameserverClient.IsConnected()).To(BeFalse())
	})

	It("can get the current player count", func() {
		count := gameserverClient.GetPlayerCount()
		Expect(count).To(Equal(0))
	})

	It("can register a player", func() {
		playerOneName := "test"
		gameserverClient.Register(playerOneName)

		count := gameserverClient.GetPlayerCount()
		Expect(count).To(Equal(1))

		gameserverClient.Deregister(playerOneName)

		count = gameserverClient.GetPlayerCount()
		Expect(count).To(Equal(0))
	})

})
