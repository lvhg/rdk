// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package tflite

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type GatherNdOptions struct {
	_tab flatbuffers.Table
}

func GetRootAsGatherNdOptions(buf []byte, offset flatbuffers.UOffsetT) *GatherNdOptions {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &GatherNdOptions{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsGatherNdOptions(buf []byte, offset flatbuffers.UOffsetT) *GatherNdOptions {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &GatherNdOptions{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *GatherNdOptions) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *GatherNdOptions) Table() flatbuffers.Table {
	return rcv._tab
}

func GatherNdOptionsStart(builder *flatbuffers.Builder) {
	builder.StartObject(0)
}
func GatherNdOptionsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
