// geecache
// @Time : 2020/3/16 11:35
// @Author : Jeffery Sun
// @File : peer
// @Software: GoLand

package geecache

import (
	pb "github.com/gofromzero/geecache/geecachepb"
)

// PeerPicker is the interface that must be implemented to locate
// the peer that owns a specific key.

type PeerPicker interface {
	PickPeer(key string) (peer PeerGetter, ok bool)
}

// PeerGetter is the interface that must be implemented by a peer.
type PeerGetter interface {
	Get(in *pb.Request, out *pb.Response) error
}
