package iptbutil

import (
	"github.com/ipfs/go-ipfs/repo/config"
)

// IpfsNode defines the interface iptb requires to work
// with an IPFS node
type IpfsNode interface {
	Init() error
	Kill() error
	Start(args []string) error
	APIAddr() (string, error)
	GetPeerID() string
	RunCmd(args ...string) (string, error)
	Shell() error
	String() string

	GetAttr(string) (string, error)
	SetAttr(string, string) error

	GetConfig() (*config.Config, error)
	WriteConfig(*config.Config) error
}
