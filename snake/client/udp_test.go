package client_test

import (
	"github.com/go-faker/faker/v4"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/theboarderline/ebiten-utilities/snake/client"
	"github.com/theboarderline/ebiten-utilities/snake/events"
	"github.com/theboarderline/ebiten-utilities/snake/param"
	"net"
)

var _ = Describe("Udp Integration", func() {

	var (
		in               chan events.Event
		out              chan events.Event
		gameserverClient *client.GameserverClient
		playerOneName    string
		playerTwoName    string
	)

	BeforeEach(func() {
		in = make(chan events.Event, 10)
		out = make(chan events.Event, 10)

		addr := net.UDPAddr{
			IP:   net.ParseIP(param.Localhost),
			Port: param.GameserverPort,
		}
		conn := client.NewGameserverConn(addr)
		gameserverClient = client.NewGameserverClient(0, &addr, conn, in, out)
		Expect(gameserverClient).NotTo(BeNil())
		Expect(gameserverClient.IsConnected()).To(BeTrue())

		go gameserverClient.HandleOutgoingEvents()
		go gameserverClient.HandleIncomingEvents()

		playerOneName = faker.FirstName()
		gameserverClient.Register(playerOneName)

		playerTwoName = faker.FirstName()
		gameserverClient.Register(playerTwoName)

	})

	AfterEach(func() {
		gameserverClient.Deregister(playerOneName)
		gameserverClient.Deregister(playerTwoName)

		err := gameserverClient.Cleanup()

		Expect(err).NotTo(HaveOccurred())
		Expect(gameserverClient.IsConnected()).To(BeFalse())
	})

	It("can get the list of players", func() {
		gameserverClient.GetPlayers(playerOneName)
	})

	It("can get the list of players", func() {
		gameserverClient.GetPlayerCount()
	})

})
