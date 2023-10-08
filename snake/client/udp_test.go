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

		gameserverClient = client.NewGameserverClient(nil, nil, in, out)
		Expect(gameserverClient).NotTo(BeNil())
		Expect(gameserverClient.IsConnected()).To(BeTrue())

	})

	AfterEach(func() {
		err := gameserverClient.Cleanup()
		Expect(err).NotTo(HaveOccurred())
		Expect(gameserverClient.IsConnected()).To(BeFalse())
	})

	It("can get the current player count", func() {
		defer gameserverClient.Cleanup()
		go gameserverClient.HandleOutgoingEvents()
		go gameserverClient.HandleIncomingEvents()

		count := gameserverClient.GetPlayerCount()

		Expect(count).To(Equal(0))
	})

	//It("can register and deregister 2 players", func() {
	//	playerOneName := "test"
	//	err := gameserverClient.Register(playerOneName)
	//	Expect(err).NotTo(HaveOccurred())
	//
	//	playerTwoName := "test2"
	//	err = gameserverClient.Register(playerTwoName)
	//	Expect(err).NotTo(HaveOccurred())
	//
	//	Expect(gameserverClient.GetPlayerCount()).To(HaveLen(2))
	//
	//	err = gameserverClient.Deregister(playerOneName)
	//	Expect(err).NotTo(HaveOccurred())
	//
	//	err = gameserverClient.Deregister(playerTwoName)
	//	Expect(err).NotTo(HaveOccurred())
	//
	//	Expect(gameserverClient.GetPlayerCount()).To(HaveLen(0))
	//})

})
