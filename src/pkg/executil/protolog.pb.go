// Code generated by protoc-gen-go.
// source: protolog.proto
// DO NOT EDIT!

/*
Package executil is a generated protocol buffer package.

It is generated from these files:
	protolog.proto

It has these top-level messages:
	RunningCommand
*/
package executil

import proto "github.com/golang/protobuf/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

type RunningCommand struct {
	Args string `protobuf:"bytes,1,opt,name=args" json:"args,omitempty"`
}

func (m *RunningCommand) Reset()         { *m = RunningCommand{} }
func (m *RunningCommand) String() string { return proto.CompactTextString(m) }
func (*RunningCommand) ProtoMessage()    {}
