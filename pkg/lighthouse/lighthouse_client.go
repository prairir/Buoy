/*
	as lighthouse is centralized, the lighthouse client simply connects to the ip address given to it.
	it will then communicate with a simple protocol with the following queries:

	get(networkName string) -> gets a node that has information about a network name
	join(networkName string, passwordHash string, node node) -> attempts to join the network.
	leave(networkName string) -> this may be done via timeout, but the client can explicitly leave as well)
*/
package lighthouse

import "net"

// a boat is a client to a lighthouse
type Boat struct {
	conn *net.UDPConn
}
