// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package login

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type WidgetFB struct {
	_tab flatbuffers.Table
}

func GetRootAsWidgetFB(buf []byte, offset flatbuffers.UOffsetT) *WidgetFB {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &WidgetFB{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *WidgetFB) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *WidgetFB) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *WidgetFB) Value1() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *WidgetFB) MutateValue1(n float64) bool {
	return rcv._tab.MutateFloat64Slot(4, n)
}

func (rcv *WidgetFB) Value2() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func WidgetFBStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func WidgetFBAddValue1(builder *flatbuffers.Builder, value1 float64) {
	builder.PrependFloat64Slot(0, value1, 0.0)
}
func WidgetFBAddValue2(builder *flatbuffers.Builder, value2 flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(value2), 0)
}
func WidgetFBEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
