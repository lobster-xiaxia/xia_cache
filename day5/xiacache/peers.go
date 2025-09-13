package xiacache

import pb "xiacache/xiacachepb/github.com/lobster-xiaxia/xia_cache/day5/xiacache/xiacachepb"

type PeerPicker interface {
	PickPeer(key string) (peer PeerGetter, ok bool)
}

type PeerGetter interface {
	Get(in *pb.Request, out *pb.Response) error
}
