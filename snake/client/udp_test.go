package client_test

import (
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

		playerName := "test"
		gameserverClient.Register(playerName)

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

	//It("can register a player", func() {
	//	playerOneName := "test"
	//	gameserverClient.Register(playerOneName)
	//
	//	count := gameserverClient.GetPlayerCount()
	//	Expect(count).To(Equal(1))
	//
	//	gameserverClient.Deregister(playerOneName)
	//
	//	count = gameserverClient.GetPlayerCount()
	//	Expect(count).To(Equal(0))
	//})

})
