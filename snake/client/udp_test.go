package client_test

import (
	"github.com/theboarderline/ebiten-utilities/snake/client"
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
		message := "test"
		Expect(gameserverClient.SendMessage([]byte(message))).To(Succeed())

		response, err := gameserverClient.GetMessage()
		Expect(err).NotTo(HaveOccurred())
		Expect(response).To(ContainSubstring("Acknowledged"))
	})

})
