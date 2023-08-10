package ping

import (
	"fmt"
	"net"
	"time"

	"github.com/tatsushid/go-fastping"
)

func ExecutePing(host string) {
	p := fastping.NewPinger()

	ra, err := net.ResolveIPAddr("ip4:icmp", host)
	if err != nil {
		fmt.Println("Error resolving host:", err)
		return
	}

	p.AddIPAddr(ra)

	received := 0
	sent := 0

	p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
		received++
		fmt.Printf("Received reply from %s: RTT = %s\n", addr, rtt)
	}

	p.OnIdle = func() {
		fmt.Println("Ping statistics:")
		fmt.Printf("  Packets: Sent = %d, Received = %d, Lost = %d (%.2f%% loss)\n",
			sent, received, sent-received, (1-float64(received)/float64(sent))*100)
	}

	fmt.Println("Pinging", host)

	for i := 0; i < 4; i++ {
		p.AddIPAddr(ra)
		sent++
	}

	err = p.Run()
	if err != nil {
		fmt.Println("Error running ping:", err)
	}
}
