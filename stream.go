package criu

// stream.go — streaming support for S3 checkpoint/restore via criu-image-streamer.
//
// Both dump and restore use go-criu's normal swrk Dump()/Restore() with
// opts.Stream=true. CRIU's service code calls img_streamer_init() when
// stream=true, connecting to the streamer socket in images_dir.

import (
	"strings"

	"github.com/checkpoint-restore/go-criu/v8/rpc"
)

// BuildStreamingCRIUConf builds criu.conf content for ext-mount-map entries,
// which are not available as RPC protobuf fields and must be in the config file.
func BuildStreamingCRIUConf(opts *rpc.CriuOpts) string {
	var sb strings.Builder
	for _, em := range opts.GetExtMnt() {
		sb.WriteString("ext-mount-map " + em.GetKey() + ":" + em.GetVal() + "\n")
	}
	return sb.String()
}
