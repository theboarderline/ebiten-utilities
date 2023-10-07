package client_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/theboarderline/ebiten-utilities/snake/client"
	"github.com/theboarderline/ebiten-utilities/snake/events"
)

var _ = Describe("Udp", func() {

	var (
		udpAddress       = "127.0.0.1"
		udpPort          = 7777
		gameserverClient *client.GameserverClient
	)

	BeforeEach(func() {
		gameserverClient = client.NewGameserverClient(udpAddress, udpPort)
		err := gameserverClient.Connect()
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		err := gameserverClient.Cleanup()
		Expect(err).NotTo(HaveOccurred())
	})

	It("can disconnect and reconnect", func() {
		err := gameserverClient.Cleanup()
		Expect(err).NotTo(HaveOccurred())

		Expect(gameserverClient.IsConnected()).To(BeFalse())

		err = gameserverClient.Connect()
		Expect(err).NotTo(HaveOccurred())

		Expect(gameserverClient.IsConnected()).To(BeTrue())
	})

	It("can register and deregister 2 players", func() {
		playerOneName := "test"
		err := gameserverClient.Register(playerOneName)
		Expect(err).NotTo(HaveOccurred())

		playerTwoName := "test2"
		err = gameserverClient.Register(playerTwoName)
		Expect(err).NotTo(HaveOccurred())

		players := gameserverClient.GetPlayers()
		Expect(players).To(HaveLen(2))

		err = gameserverClient.Deregister(playerOneName)
		Expect(err).NotTo(HaveOccurred())

		err = gameserverClient.Deregister(playerTwoName)
		Expect(err).NotTo(HaveOccurred())
	})

	It("send and receive messages", func() {
		event := events.Event{Type: "ACK"}
		Expect(gameserverClient.SendMessage(event)).To(Succeed())

		response, err := gameserverClient.GetMessage()
		Expect(err).NotTo(HaveOccurred())
		Expect(response.Message).To(ContainSubstring("Acknowledged"))
	})

})
