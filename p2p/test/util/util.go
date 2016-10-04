package testutil

import (
	"context"
	"testing"

	bhost "github.com/libp2p/go-libp2p/p2p/host/basic"
	metrics "github.com/libp2p/go-libp2p/p2p/metrics"
	inet "github.com/libp2p/go-libp2p/p2p/net"
	swarm "github.com/libp2p/go-libp2p/p2p/net/swarm"

	pstore "github.com/ipfs/go-libp2p-peerstore"
	ma "github.com/jbenet/go-multiaddr"
	tu "github.com/libp2p/go-testutil"
)

func GenSwarmNetwork(t *testing.T, ctx context.Context) *swarm.Network {
	p := tu.RandPeerNetParamsOrFatal(t)
	ps := pstore.NewPeerstore()
	ps.AddPubKey(p.ID, p.PubKey)
	ps.AddPrivKey(p.ID, p.PrivKey)
	n, err := swarm.NewNetwork(ctx, []ma.Multiaddr{p.Addr}, p.ID, ps, metrics.NewBandwidthCounter())
	if err != nil {
		t.Fatal(err)
	}
	ps.AddAddrs(p.ID, n.ListenAddresses(), pstore.PermanentAddrTTL)
	return n
}

func DivulgeAddresses(a, b inet.Network) {
	id := a.LocalPeer()
	addrs := a.Peerstore().Addrs(id)
	b.Peerstore().AddAddrs(id, addrs, pstore.PermanentAddrTTL)
}

func GenHostSwarm(t *testing.T, ctx context.Context) *bhost.BasicHost {
	n := GenSwarmNetwork(t, ctx)
	return bhost.New(n)
}

var RandPeerID = tu.RandPeerID
