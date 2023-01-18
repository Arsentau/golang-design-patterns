package servicesExample

import "fmt"

// Set main interface
type DnsProvider interface {
	SetCredentials()
	GetZones()
}

// Implement interface for Ns1
type Ns1 struct {
	provider *DnsProvider
}

func (n *Ns1) setCredentials() {
	fmt.Println("Setting NS1 credentials")
}

func (n *Ns1) GetZones() {
	n.setCredentials()
	fmt.Println("Getting zones information from NS1")
}

// Implement interface for Route83
type Route83 struct {
	provider DnsProvider
}

func (n *Route83) setCredentials() {
	fmt.Println("Setting AWS credentials")
}

func (n *Route83) GetZones() {
	n.setCredentials()
	fmt.Println("Getting zones information from Route83")
}

// Implement interface for Cloudflare
type Cloudflare struct {
	provider DnsProvider
}

func (n *Cloudflare) setCredentials() {
	fmt.Println("Setting Cloudflare credentials")
}

func (n *Cloudflare) GetZones() {
	n.setCredentials()
	fmt.Println("Getting zones information from Cloudflare")
}

// Application
func Bridge() {
	ns1 := &Ns1{}
	route83 := &Route83{}
	cloudflare := &Cloudflare{}

	ns1.GetZones()
	route83.GetZones()
	cloudflare.GetZones()
}
