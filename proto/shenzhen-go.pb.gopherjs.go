//+build js
// Code generated by protoc-gen-gopherjs. DO NOT EDIT.
// source: shenzhen-go.proto

/*
	Package proto is a generated protocol buffer package.

	It is generated from these files:
		shenzhen-go.proto

	It has these top-level messages:
		Empty
		NodeConfig
		CreateChannelRequest
		CreateNodeRequest
		ConnectPinRequest
		DeleteChannelRequest
		DeleteNodeRequest
		DisconnectPinRequest
		SaveRequest
		SetGraphPropertiesRequest
		SetNodePropertiesRequest
		SetPositionRequest
*/
package proto

import jspb "github.com/johanbrandhorst/protobuf/jspb"

import (
	context "context"

	grpcweb "github.com/johanbrandhorst/protobuf/grpcweb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the jspb package it is being compiled against.
const _ = jspb.JspbPackageIsVersion2

type Empty struct {
}

// MarshalToWriter marshals Empty to the provided writer.
func (m *Empty) MarshalToWriter(writer jspb.Writer) {
	if m == nil {
		return
	}

	return
}

// Marshal marshals Empty to a slice of bytes.
func (m *Empty) Marshal() []byte {
	writer := jspb.NewWriter()
	m.MarshalToWriter(writer)
	return writer.GetResult()
}

// UnmarshalFromReader unmarshals a Empty from the provided reader.
func (m *Empty) UnmarshalFromReader(reader jspb.Reader) *Empty {
	for reader.Next() {
		if m == nil {
			m = &Empty{}
		}

		switch reader.GetFieldNumber() {
		default:
			reader.SkipField()
		}
	}

	return m
}

// Unmarshal unmarshals a Empty from a slice of bytes.
func (m *Empty) Unmarshal(rawBytes []byte) (*Empty, error) {
	reader := jspb.NewReader(rawBytes)

	m = m.UnmarshalFromReader(reader)

	if err := reader.Err(); err != nil {
		return nil, err
	}

	return m, nil
}

type NodeConfig struct {
	Name         string
	Enabled      bool
	Multiplicity uint32
	Wait         bool
	PartCfg      []byte
	PartType     string
	X            float64
	Y            float64
}

// GetName gets the Name of the NodeConfig.
func (m *NodeConfig) GetName() (x string) {
	if m == nil {
		return x
	}
	return m.Name
}

// GetEnabled gets the Enabled of the NodeConfig.
func (m *NodeConfig) GetEnabled() (x bool) {
	if m == nil {
		return x
	}
	return m.Enabled
}

// GetMultiplicity gets the Multiplicity of the NodeConfig.
func (m *NodeConfig) GetMultiplicity() (x uint32) {
	if m == nil {
		return x
	}
	return m.Multiplicity
}

// GetWait gets the Wait of the NodeConfig.
func (m *NodeConfig) GetWait() (x bool) {
	if m == nil {
		return x
	}
	return m.Wait
}

// GetPartCfg gets the PartCfg of the NodeConfig.
func (m *NodeConfig) GetPartCfg() (x []byte) {
	if m == nil {
		return x
	}
	return m.PartCfg
}

// GetPartType gets the PartType of the NodeConfig.
func (m *NodeConfig) GetPartType() (x string) {
	if m == nil {
		return x
	}
	return m.PartType
}

// GetX gets the X of the NodeConfig.
func (m *NodeConfig) GetX() (x float64) {
	if m == nil {
		return x
	}
	return m.X
}

// GetY gets the Y of the NodeConfig.
func (m *NodeConfig) GetY() (x float64) {
	if m == nil {
		return x
	}
	return m.Y
}

// MarshalToWriter marshals NodeConfig to the provided writer.
func (m *NodeConfig) MarshalToWriter(writer jspb.Writer) {
	if m == nil {
		return
	}

	if len(m.Name) > 0 {
		writer.WriteString(1, m.Name)
	}

	if m.Enabled {
		writer.WriteBool(2, m.Enabled)
	}

	if m.Multiplicity != 0 {
		writer.WriteUint32(3, m.Multiplicity)
	}

	if m.Wait {
		writer.WriteBool(4, m.Wait)
	}

	if len(m.PartCfg) > 0 {
		writer.WriteBytes(5, m.PartCfg)
	}

	if len(m.PartType) > 0 {
		writer.WriteString(6, m.PartType)
	}

	if m.X != 0 {
		writer.WriteFloat64(7, m.X)
	}

	if m.Y != 0 {
		writer.WriteFloat64(8, m.Y)
	}

	return
}

// Marshal marshals NodeConfig to a slice of bytes.
func (m *NodeConfig) Marshal() []byte {
	writer := jspb.NewWriter()
	m.MarshalToWriter(writer)
	return writer.GetResult()
}

// UnmarshalFromReader unmarshals a NodeConfig from the provided reader.
func (m *NodeConfig) UnmarshalFromReader(reader jspb.Reader) *NodeConfig {
	for reader.Next() {
		if m == nil {
			m = &NodeConfig{}
		}

		switch reader.GetFieldNumber() {
		case 1:
			m.Name = reader.ReadString()
		case 2:
			m.Enabled = reader.ReadBool()
		case 3:
			m.Multiplicity = reader.ReadUint32()
		case 4:
			m.Wait = reader.ReadBool()
		case 5:
			m.PartCfg = reader.ReadBytes()
		case 6:
			m.PartType = reader.ReadString()
		case 7:
			m.X = reader.ReadFloat64()
		case 8:
			m.Y = reader.ReadFloat64()
		default:
			reader.SkipField()
		}
	}

	return m
}

// Unmarshal unmarshals a NodeConfig from a slice of bytes.
func (m *NodeConfig) Unmarshal(rawBytes []byte) (*NodeConfig, error) {
	reader := jspb.NewReader(rawBytes)

	m = m.UnmarshalFromReader(reader)

	if err := reader.Err(); err != nil {
		return nil, err
	}

	return m, nil
}

type CreateChannelRequest struct {
	Graph string
	Name  string
	Type  string
	Anon  bool
	Cap   uint64
	// Also connect two pins together, to reduce chatter (1 Create RPC instead of [Create, Connect, Connect])
	Node1 string
	Pin1  string
	Node2 string
	Pin2  string
}

// GetGraph gets the Graph of the CreateChannelRequest.
func (m *CreateChannelRequest) GetGraph() (x string) {
	if m == nil {
		return x
	}
	return m.Graph
}

// GetName gets the Name of the CreateChannelRequest.
func (m *CreateChannelRequest) GetName() (x string) {
	if m == nil {
		return x
	}
	return m.Name
}

// GetType gets the Type of the CreateChannelRequest.
func (m *CreateChannelRequest) GetType() (x string) {
	if m == nil {
		return x
	}
	return m.Type
}

// GetAnon gets the Anon of the CreateChannelRequest.
func (m *CreateChannelRequest) GetAnon() (x bool) {
	if m == nil {
		return x
	}
	return m.Anon
}

// GetCap gets the Cap of the CreateChannelRequest.
func (m *CreateChannelRequest) GetCap() (x uint64) {
	if m == nil {
		return x
	}
	return m.Cap
}

// GetNode1 gets the Node1 of the CreateChannelRequest.
func (m *CreateChannelRequest) GetNode1() (x string) {
	if m == nil {
		return x
	}
	return m.Node1
}

// GetPin1 gets the Pin1 of the CreateChannelRequest.
func (m *CreateChannelRequest) GetPin1() (x string) {
	if m == nil {
		return x
	}
	return m.Pin1
}

// GetNode2 gets the Node2 of the CreateChannelRequest.
func (m *CreateChannelRequest) GetNode2() (x string) {
	if m == nil {
		return x
	}
	return m.Node2
}

// GetPin2 gets the Pin2 of the CreateChannelRequest.
func (m *CreateChannelRequest) GetPin2() (x string) {
	if m == nil {
		return x
	}
	return m.Pin2
}

// MarshalToWriter marshals CreateChannelRequest to the provided writer.
func (m *CreateChannelRequest) MarshalToWriter(writer jspb.Writer) {
	if m == nil {
		return
	}

	if len(m.Graph) > 0 {
		writer.WriteString(1, m.Graph)
	}

	if len(m.Name) > 0 {
		writer.WriteString(2, m.Name)
	}

	if len(m.Type) > 0 {
		writer.WriteString(3, m.Type)
	}

	if m.Anon {
		writer.WriteBool(4, m.Anon)
	}

	if m.Cap != 0 {
		writer.WriteUint64(5, m.Cap)
	}

	if len(m.Node1) > 0 {
		writer.WriteString(6, m.Node1)
	}

	if len(m.Pin1) > 0 {
		writer.WriteString(7, m.Pin1)
	}

	if len(m.Node2) > 0 {
		writer.WriteString(8, m.Node2)
	}

	if len(m.Pin2) > 0 {
		writer.WriteString(9, m.Pin2)
	}

	return
}

// Marshal marshals CreateChannelRequest to a slice of bytes.
func (m *CreateChannelRequest) Marshal() []byte {
	writer := jspb.NewWriter()
	m.MarshalToWriter(writer)
	return writer.GetResult()
}

// UnmarshalFromReader unmarshals a CreateChannelRequest from the provided reader.
func (m *CreateChannelRequest) UnmarshalFromReader(reader jspb.Reader) *CreateChannelRequest {
	for reader.Next() {
		if m == nil {
			m = &CreateChannelRequest{}
		}

		switch reader.GetFieldNumber() {
		case 1:
			m.Graph = reader.ReadString()
		case 2:
			m.Name = reader.ReadString()
		case 3:
			m.Type = reader.ReadString()
		case 4:
			m.Anon = reader.ReadBool()
		case 5:
			m.Cap = reader.ReadUint64()
		case 6:
			m.Node1 = reader.ReadString()
		case 7:
			m.Pin1 = reader.ReadString()
		case 8:
			m.Node2 = reader.ReadString()
		case 9:
			m.Pin2 = reader.ReadString()
		default:
			reader.SkipField()
		}
	}

	return m
}

// Unmarshal unmarshals a CreateChannelRequest from a slice of bytes.
func (m *CreateChannelRequest) Unmarshal(rawBytes []byte) (*CreateChannelRequest, error) {
	reader := jspb.NewReader(rawBytes)

	m = m.UnmarshalFromReader(reader)

	if err := reader.Err(); err != nil {
		return nil, err
	}

	return m, nil
}

type CreateNodeRequest struct {
	Graph string
	Props *NodeConfig
}

// GetGraph gets the Graph of the CreateNodeRequest.
func (m *CreateNodeRequest) GetGraph() (x string) {
	if m == nil {
		return x
	}
	return m.Graph
}

// GetProps gets the Props of the CreateNodeRequest.
func (m *CreateNodeRequest) GetProps() (x *NodeConfig) {
	if m == nil {
		return x
	}
	return m.Props
}

// MarshalToWriter marshals CreateNodeRequest to the provided writer.
func (m *CreateNodeRequest) MarshalToWriter(writer jspb.Writer) {
	if m == nil {
		return
	}

	if len(m.Graph) > 0 {
		writer.WriteString(1, m.Graph)
	}

	if m.Props != nil {
		writer.WriteMessage(2, func() {
			m.Props.MarshalToWriter(writer)
		})
	}

	return
}

// Marshal marshals CreateNodeRequest to a slice of bytes.
func (m *CreateNodeRequest) Marshal() []byte {
	writer := jspb.NewWriter()
	m.MarshalToWriter(writer)
	return writer.GetResult()
}

// UnmarshalFromReader unmarshals a CreateNodeRequest from the provided reader.
func (m *CreateNodeRequest) UnmarshalFromReader(reader jspb.Reader) *CreateNodeRequest {
	for reader.Next() {
		if m == nil {
			m = &CreateNodeRequest{}
		}

		switch reader.GetFieldNumber() {
		case 1:
			m.Graph = reader.ReadString()
		case 2:
			reader.ReadMessage(func() {
				m.Props = m.Props.UnmarshalFromReader(reader)
			})
		default:
			reader.SkipField()
		}
	}

	return m
}

// Unmarshal unmarshals a CreateNodeRequest from a slice of bytes.
func (m *CreateNodeRequest) Unmarshal(rawBytes []byte) (*CreateNodeRequest, error) {
	reader := jspb.NewReader(rawBytes)

	m = m.UnmarshalFromReader(reader)

	if err := reader.Err(); err != nil {
		return nil, err
	}

	return m, nil
}

type ConnectPinRequest struct {
	Graph   string
	Node    string
	Pin     string
	Channel string
}

// GetGraph gets the Graph of the ConnectPinRequest.
func (m *ConnectPinRequest) GetGraph() (x string) {
	if m == nil {
		return x
	}
	return m.Graph
}

// GetNode gets the Node of the ConnectPinRequest.
func (m *ConnectPinRequest) GetNode() (x string) {
	if m == nil {
		return x
	}
	return m.Node
}

// GetPin gets the Pin of the ConnectPinRequest.
func (m *ConnectPinRequest) GetPin() (x string) {
	if m == nil {
		return x
	}
	return m.Pin
}

// GetChannel gets the Channel of the ConnectPinRequest.
func (m *ConnectPinRequest) GetChannel() (x string) {
	if m == nil {
		return x
	}
	return m.Channel
}

// MarshalToWriter marshals ConnectPinRequest to the provided writer.
func (m *ConnectPinRequest) MarshalToWriter(writer jspb.Writer) {
	if m == nil {
		return
	}

	if len(m.Graph) > 0 {
		writer.WriteString(1, m.Graph)
	}

	if len(m.Node) > 0 {
		writer.WriteString(2, m.Node)
	}

	if len(m.Pin) > 0 {
		writer.WriteString(3, m.Pin)
	}

	if len(m.Channel) > 0 {
		writer.WriteString(4, m.Channel)
	}

	return
}

// Marshal marshals ConnectPinRequest to a slice of bytes.
func (m *ConnectPinRequest) Marshal() []byte {
	writer := jspb.NewWriter()
	m.MarshalToWriter(writer)
	return writer.GetResult()
}

// UnmarshalFromReader unmarshals a ConnectPinRequest from the provided reader.
func (m *ConnectPinRequest) UnmarshalFromReader(reader jspb.Reader) *ConnectPinRequest {
	for reader.Next() {
		if m == nil {
			m = &ConnectPinRequest{}
		}

		switch reader.GetFieldNumber() {
		case 1:
			m.Graph = reader.ReadString()
		case 2:
			m.Node = reader.ReadString()
		case 3:
			m.Pin = reader.ReadString()
		case 4:
			m.Channel = reader.ReadString()
		default:
			reader.SkipField()
		}
	}

	return m
}

// Unmarshal unmarshals a ConnectPinRequest from a slice of bytes.
func (m *ConnectPinRequest) Unmarshal(rawBytes []byte) (*ConnectPinRequest, error) {
	reader := jspb.NewReader(rawBytes)

	m = m.UnmarshalFromReader(reader)

	if err := reader.Err(); err != nil {
		return nil, err
	}

	return m, nil
}

type DeleteChannelRequest struct {
	Graph   string
	Channel string
}

// GetGraph gets the Graph of the DeleteChannelRequest.
func (m *DeleteChannelRequest) GetGraph() (x string) {
	if m == nil {
		return x
	}
	return m.Graph
}

// GetChannel gets the Channel of the DeleteChannelRequest.
func (m *DeleteChannelRequest) GetChannel() (x string) {
	if m == nil {
		return x
	}
	return m.Channel
}

// MarshalToWriter marshals DeleteChannelRequest to the provided writer.
func (m *DeleteChannelRequest) MarshalToWriter(writer jspb.Writer) {
	if m == nil {
		return
	}

	if len(m.Graph) > 0 {
		writer.WriteString(1, m.Graph)
	}

	if len(m.Channel) > 0 {
		writer.WriteString(2, m.Channel)
	}

	return
}

// Marshal marshals DeleteChannelRequest to a slice of bytes.
func (m *DeleteChannelRequest) Marshal() []byte {
	writer := jspb.NewWriter()
	m.MarshalToWriter(writer)
	return writer.GetResult()
}

// UnmarshalFromReader unmarshals a DeleteChannelRequest from the provided reader.
func (m *DeleteChannelRequest) UnmarshalFromReader(reader jspb.Reader) *DeleteChannelRequest {
	for reader.Next() {
		if m == nil {
			m = &DeleteChannelRequest{}
		}

		switch reader.GetFieldNumber() {
		case 1:
			m.Graph = reader.ReadString()
		case 2:
			m.Channel = reader.ReadString()
		default:
			reader.SkipField()
		}
	}

	return m
}

// Unmarshal unmarshals a DeleteChannelRequest from a slice of bytes.
func (m *DeleteChannelRequest) Unmarshal(rawBytes []byte) (*DeleteChannelRequest, error) {
	reader := jspb.NewReader(rawBytes)

	m = m.UnmarshalFromReader(reader)

	if err := reader.Err(); err != nil {
		return nil, err
	}

	return m, nil
}

type DeleteNodeRequest struct {
	Graph string
	Node  string
}

// GetGraph gets the Graph of the DeleteNodeRequest.
func (m *DeleteNodeRequest) GetGraph() (x string) {
	if m == nil {
		return x
	}
	return m.Graph
}

// GetNode gets the Node of the DeleteNodeRequest.
func (m *DeleteNodeRequest) GetNode() (x string) {
	if m == nil {
		return x
	}
	return m.Node
}

// MarshalToWriter marshals DeleteNodeRequest to the provided writer.
func (m *DeleteNodeRequest) MarshalToWriter(writer jspb.Writer) {
	if m == nil {
		return
	}

	if len(m.Graph) > 0 {
		writer.WriteString(1, m.Graph)
	}

	if len(m.Node) > 0 {
		writer.WriteString(2, m.Node)
	}

	return
}

// Marshal marshals DeleteNodeRequest to a slice of bytes.
func (m *DeleteNodeRequest) Marshal() []byte {
	writer := jspb.NewWriter()
	m.MarshalToWriter(writer)
	return writer.GetResult()
}

// UnmarshalFromReader unmarshals a DeleteNodeRequest from the provided reader.
func (m *DeleteNodeRequest) UnmarshalFromReader(reader jspb.Reader) *DeleteNodeRequest {
	for reader.Next() {
		if m == nil {
			m = &DeleteNodeRequest{}
		}

		switch reader.GetFieldNumber() {
		case 1:
			m.Graph = reader.ReadString()
		case 2:
			m.Node = reader.ReadString()
		default:
			reader.SkipField()
		}
	}

	return m
}

// Unmarshal unmarshals a DeleteNodeRequest from a slice of bytes.
func (m *DeleteNodeRequest) Unmarshal(rawBytes []byte) (*DeleteNodeRequest, error) {
	reader := jspb.NewReader(rawBytes)

	m = m.UnmarshalFromReader(reader)

	if err := reader.Err(); err != nil {
		return nil, err
	}

	return m, nil
}

type DisconnectPinRequest struct {
	Graph string
	Node  string
	Pin   string
}

// GetGraph gets the Graph of the DisconnectPinRequest.
func (m *DisconnectPinRequest) GetGraph() (x string) {
	if m == nil {
		return x
	}
	return m.Graph
}

// GetNode gets the Node of the DisconnectPinRequest.
func (m *DisconnectPinRequest) GetNode() (x string) {
	if m == nil {
		return x
	}
	return m.Node
}

// GetPin gets the Pin of the DisconnectPinRequest.
func (m *DisconnectPinRequest) GetPin() (x string) {
	if m == nil {
		return x
	}
	return m.Pin
}

// MarshalToWriter marshals DisconnectPinRequest to the provided writer.
func (m *DisconnectPinRequest) MarshalToWriter(writer jspb.Writer) {
	if m == nil {
		return
	}

	if len(m.Graph) > 0 {
		writer.WriteString(1, m.Graph)
	}

	if len(m.Node) > 0 {
		writer.WriteString(2, m.Node)
	}

	if len(m.Pin) > 0 {
		writer.WriteString(3, m.Pin)
	}

	return
}

// Marshal marshals DisconnectPinRequest to a slice of bytes.
func (m *DisconnectPinRequest) Marshal() []byte {
	writer := jspb.NewWriter()
	m.MarshalToWriter(writer)
	return writer.GetResult()
}

// UnmarshalFromReader unmarshals a DisconnectPinRequest from the provided reader.
func (m *DisconnectPinRequest) UnmarshalFromReader(reader jspb.Reader) *DisconnectPinRequest {
	for reader.Next() {
		if m == nil {
			m = &DisconnectPinRequest{}
		}

		switch reader.GetFieldNumber() {
		case 1:
			m.Graph = reader.ReadString()
		case 2:
			m.Node = reader.ReadString()
		case 3:
			m.Pin = reader.ReadString()
		default:
			reader.SkipField()
		}
	}

	return m
}

// Unmarshal unmarshals a DisconnectPinRequest from a slice of bytes.
func (m *DisconnectPinRequest) Unmarshal(rawBytes []byte) (*DisconnectPinRequest, error) {
	reader := jspb.NewReader(rawBytes)

	m = m.UnmarshalFromReader(reader)

	if err := reader.Err(); err != nil {
		return nil, err
	}

	return m, nil
}

type SaveRequest struct {
	Graph string
}

// GetGraph gets the Graph of the SaveRequest.
func (m *SaveRequest) GetGraph() (x string) {
	if m == nil {
		return x
	}
	return m.Graph
}

// MarshalToWriter marshals SaveRequest to the provided writer.
func (m *SaveRequest) MarshalToWriter(writer jspb.Writer) {
	if m == nil {
		return
	}

	if len(m.Graph) > 0 {
		writer.WriteString(1, m.Graph)
	}

	return
}

// Marshal marshals SaveRequest to a slice of bytes.
func (m *SaveRequest) Marshal() []byte {
	writer := jspb.NewWriter()
	m.MarshalToWriter(writer)
	return writer.GetResult()
}

// UnmarshalFromReader unmarshals a SaveRequest from the provided reader.
func (m *SaveRequest) UnmarshalFromReader(reader jspb.Reader) *SaveRequest {
	for reader.Next() {
		if m == nil {
			m = &SaveRequest{}
		}

		switch reader.GetFieldNumber() {
		case 1:
			m.Graph = reader.ReadString()
		default:
			reader.SkipField()
		}
	}

	return m
}

// Unmarshal unmarshals a SaveRequest from a slice of bytes.
func (m *SaveRequest) Unmarshal(rawBytes []byte) (*SaveRequest, error) {
	reader := jspb.NewReader(rawBytes)

	m = m.UnmarshalFromReader(reader)

	if err := reader.Err(); err != nil {
		return nil, err
	}

	return m, nil
}

type SetGraphPropertiesRequest struct {
	Graph       string
	Name        string
	PackagePath string
	IsCommand   bool
}

// GetGraph gets the Graph of the SetGraphPropertiesRequest.
func (m *SetGraphPropertiesRequest) GetGraph() (x string) {
	if m == nil {
		return x
	}
	return m.Graph
}

// GetName gets the Name of the SetGraphPropertiesRequest.
func (m *SetGraphPropertiesRequest) GetName() (x string) {
	if m == nil {
		return x
	}
	return m.Name
}

// GetPackagePath gets the PackagePath of the SetGraphPropertiesRequest.
func (m *SetGraphPropertiesRequest) GetPackagePath() (x string) {
	if m == nil {
		return x
	}
	return m.PackagePath
}

// GetIsCommand gets the IsCommand of the SetGraphPropertiesRequest.
func (m *SetGraphPropertiesRequest) GetIsCommand() (x bool) {
	if m == nil {
		return x
	}
	return m.IsCommand
}

// MarshalToWriter marshals SetGraphPropertiesRequest to the provided writer.
func (m *SetGraphPropertiesRequest) MarshalToWriter(writer jspb.Writer) {
	if m == nil {
		return
	}

	if len(m.Graph) > 0 {
		writer.WriteString(1, m.Graph)
	}

	if len(m.Name) > 0 {
		writer.WriteString(2, m.Name)
	}

	if len(m.PackagePath) > 0 {
		writer.WriteString(3, m.PackagePath)
	}

	if m.IsCommand {
		writer.WriteBool(4, m.IsCommand)
	}

	return
}

// Marshal marshals SetGraphPropertiesRequest to a slice of bytes.
func (m *SetGraphPropertiesRequest) Marshal() []byte {
	writer := jspb.NewWriter()
	m.MarshalToWriter(writer)
	return writer.GetResult()
}

// UnmarshalFromReader unmarshals a SetGraphPropertiesRequest from the provided reader.
func (m *SetGraphPropertiesRequest) UnmarshalFromReader(reader jspb.Reader) *SetGraphPropertiesRequest {
	for reader.Next() {
		if m == nil {
			m = &SetGraphPropertiesRequest{}
		}

		switch reader.GetFieldNumber() {
		case 1:
			m.Graph = reader.ReadString()
		case 2:
			m.Name = reader.ReadString()
		case 3:
			m.PackagePath = reader.ReadString()
		case 4:
			m.IsCommand = reader.ReadBool()
		default:
			reader.SkipField()
		}
	}

	return m
}

// Unmarshal unmarshals a SetGraphPropertiesRequest from a slice of bytes.
func (m *SetGraphPropertiesRequest) Unmarshal(rawBytes []byte) (*SetGraphPropertiesRequest, error) {
	reader := jspb.NewReader(rawBytes)

	m = m.UnmarshalFromReader(reader)

	if err := reader.Err(); err != nil {
		return nil, err
	}

	return m, nil
}

type SetNodePropertiesRequest struct {
	Graph string
	Node  string
	Props *NodeConfig
}

// GetGraph gets the Graph of the SetNodePropertiesRequest.
func (m *SetNodePropertiesRequest) GetGraph() (x string) {
	if m == nil {
		return x
	}
	return m.Graph
}

// GetNode gets the Node of the SetNodePropertiesRequest.
func (m *SetNodePropertiesRequest) GetNode() (x string) {
	if m == nil {
		return x
	}
	return m.Node
}

// GetProps gets the Props of the SetNodePropertiesRequest.
func (m *SetNodePropertiesRequest) GetProps() (x *NodeConfig) {
	if m == nil {
		return x
	}
	return m.Props
}

// MarshalToWriter marshals SetNodePropertiesRequest to the provided writer.
func (m *SetNodePropertiesRequest) MarshalToWriter(writer jspb.Writer) {
	if m == nil {
		return
	}

	if len(m.Graph) > 0 {
		writer.WriteString(1, m.Graph)
	}

	if len(m.Node) > 0 {
		writer.WriteString(2, m.Node)
	}

	if m.Props != nil {
		writer.WriteMessage(3, func() {
			m.Props.MarshalToWriter(writer)
		})
	}

	return
}

// Marshal marshals SetNodePropertiesRequest to a slice of bytes.
func (m *SetNodePropertiesRequest) Marshal() []byte {
	writer := jspb.NewWriter()
	m.MarshalToWriter(writer)
	return writer.GetResult()
}

// UnmarshalFromReader unmarshals a SetNodePropertiesRequest from the provided reader.
func (m *SetNodePropertiesRequest) UnmarshalFromReader(reader jspb.Reader) *SetNodePropertiesRequest {
	for reader.Next() {
		if m == nil {
			m = &SetNodePropertiesRequest{}
		}

		switch reader.GetFieldNumber() {
		case 1:
			m.Graph = reader.ReadString()
		case 2:
			m.Node = reader.ReadString()
		case 3:
			reader.ReadMessage(func() {
				m.Props = m.Props.UnmarshalFromReader(reader)
			})
		default:
			reader.SkipField()
		}
	}

	return m
}

// Unmarshal unmarshals a SetNodePropertiesRequest from a slice of bytes.
func (m *SetNodePropertiesRequest) Unmarshal(rawBytes []byte) (*SetNodePropertiesRequest, error) {
	reader := jspb.NewReader(rawBytes)

	m = m.UnmarshalFromReader(reader)

	if err := reader.Err(); err != nil {
		return nil, err
	}

	return m, nil
}

type SetPositionRequest struct {
	Graph string
	Node  string
	X     float64
	Y     float64
}

// GetGraph gets the Graph of the SetPositionRequest.
func (m *SetPositionRequest) GetGraph() (x string) {
	if m == nil {
		return x
	}
	return m.Graph
}

// GetNode gets the Node of the SetPositionRequest.
func (m *SetPositionRequest) GetNode() (x string) {
	if m == nil {
		return x
	}
	return m.Node
}

// GetX gets the X of the SetPositionRequest.
func (m *SetPositionRequest) GetX() (x float64) {
	if m == nil {
		return x
	}
	return m.X
}

// GetY gets the Y of the SetPositionRequest.
func (m *SetPositionRequest) GetY() (x float64) {
	if m == nil {
		return x
	}
	return m.Y
}

// MarshalToWriter marshals SetPositionRequest to the provided writer.
func (m *SetPositionRequest) MarshalToWriter(writer jspb.Writer) {
	if m == nil {
		return
	}

	if len(m.Graph) > 0 {
		writer.WriteString(1, m.Graph)
	}

	if len(m.Node) > 0 {
		writer.WriteString(2, m.Node)
	}

	if m.X != 0 {
		writer.WriteFloat64(3, m.X)
	}

	if m.Y != 0 {
		writer.WriteFloat64(4, m.Y)
	}

	return
}

// Marshal marshals SetPositionRequest to a slice of bytes.
func (m *SetPositionRequest) Marshal() []byte {
	writer := jspb.NewWriter()
	m.MarshalToWriter(writer)
	return writer.GetResult()
}

// UnmarshalFromReader unmarshals a SetPositionRequest from the provided reader.
func (m *SetPositionRequest) UnmarshalFromReader(reader jspb.Reader) *SetPositionRequest {
	for reader.Next() {
		if m == nil {
			m = &SetPositionRequest{}
		}

		switch reader.GetFieldNumber() {
		case 1:
			m.Graph = reader.ReadString()
		case 2:
			m.Node = reader.ReadString()
		case 3:
			m.X = reader.ReadFloat64()
		case 4:
			m.Y = reader.ReadFloat64()
		default:
			reader.SkipField()
		}
	}

	return m
}

// Unmarshal unmarshals a SetPositionRequest from a slice of bytes.
func (m *SetPositionRequest) Unmarshal(rawBytes []byte) (*SetPositionRequest, error) {
	reader := jspb.NewReader(rawBytes)

	m = m.UnmarshalFromReader(reader)

	if err := reader.Err(); err != nil {
		return nil, err
	}

	return m, nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpcweb.Client

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpcweb package it is being compiled against.
const _ = grpcweb.GrpcWebPackageIsVersion2

// Client API for ShenzhenGo service

type ShenzhenGoClient interface {
	// CreateChannel makes a new channel.
	CreateChannel(ctx context.Context, in *CreateChannelRequest, opts ...grpcweb.CallOption) (*Empty, error)
	// CreateNode makes a new node.
	CreateNode(ctx context.Context, in *CreateNodeRequest, opts ...grpcweb.CallOption) (*Empty, error)
	// ConnectPin connects a pin to a channel.
	ConnectPin(ctx context.Context, in *ConnectPinRequest, opts ...grpcweb.CallOption) (*Empty, error)
	// DeleteChannel deletes a channel (and all connections).
	DeleteChannel(ctx context.Context, in *DeleteChannelRequest, opts ...grpcweb.CallOption) (*Empty, error)
	// DeleteNode deletes a node (and all connections).
	DeleteNode(ctx context.Context, in *DeleteNodeRequest, opts ...grpcweb.CallOption) (*Empty, error)
	// DisconnectPin deletes the connection from a pin to a channel.
	DisconnectPin(ctx context.Context, in *DisconnectPinRequest, opts ...grpcweb.CallOption) (*Empty, error)
	// Save saves the graph to disk.
	Save(ctx context.Context, in *SaveRequest, opts ...grpcweb.CallOption) (*Empty, error)
	// SetGraphProperties changes metdata such as name and package path.
	SetGraphProperties(ctx context.Context, in *SetGraphPropertiesRequest, opts ...grpcweb.CallOption) (*Empty, error)
	// SetNodeProperties changes node metadata such as name and multiplicity.
	SetNodeProperties(ctx context.Context, in *SetNodePropertiesRequest, opts ...grpcweb.CallOption) (*Empty, error)
	// SetPosition changes the node position in the diagram.
	SetPosition(ctx context.Context, in *SetPositionRequest, opts ...grpcweb.CallOption) (*Empty, error)
}

type shenzhenGoClient struct {
	client *grpcweb.Client
}

// NewShenzhenGoClient creates a new gRPC-Web client.
func NewShenzhenGoClient(hostname string, opts ...grpcweb.DialOption) ShenzhenGoClient {
	return &shenzhenGoClient{
		client: grpcweb.NewClient(hostname, "proto.ShenzhenGo", opts...),
	}
}

func (c *shenzhenGoClient) CreateChannel(ctx context.Context, in *CreateChannelRequest, opts ...grpcweb.CallOption) (*Empty, error) {
	resp, err := c.client.RPCCall(ctx, "CreateChannel", in.Marshal(), opts...)
	if err != nil {
		return nil, err
	}

	return new(Empty).Unmarshal(resp)
}

func (c *shenzhenGoClient) CreateNode(ctx context.Context, in *CreateNodeRequest, opts ...grpcweb.CallOption) (*Empty, error) {
	resp, err := c.client.RPCCall(ctx, "CreateNode", in.Marshal(), opts...)
	if err != nil {
		return nil, err
	}

	return new(Empty).Unmarshal(resp)
}

func (c *shenzhenGoClient) ConnectPin(ctx context.Context, in *ConnectPinRequest, opts ...grpcweb.CallOption) (*Empty, error) {
	resp, err := c.client.RPCCall(ctx, "ConnectPin", in.Marshal(), opts...)
	if err != nil {
		return nil, err
	}

	return new(Empty).Unmarshal(resp)
}

func (c *shenzhenGoClient) DeleteChannel(ctx context.Context, in *DeleteChannelRequest, opts ...grpcweb.CallOption) (*Empty, error) {
	resp, err := c.client.RPCCall(ctx, "DeleteChannel", in.Marshal(), opts...)
	if err != nil {
		return nil, err
	}

	return new(Empty).Unmarshal(resp)
}

func (c *shenzhenGoClient) DeleteNode(ctx context.Context, in *DeleteNodeRequest, opts ...grpcweb.CallOption) (*Empty, error) {
	resp, err := c.client.RPCCall(ctx, "DeleteNode", in.Marshal(), opts...)
	if err != nil {
		return nil, err
	}

	return new(Empty).Unmarshal(resp)
}

func (c *shenzhenGoClient) DisconnectPin(ctx context.Context, in *DisconnectPinRequest, opts ...grpcweb.CallOption) (*Empty, error) {
	resp, err := c.client.RPCCall(ctx, "DisconnectPin", in.Marshal(), opts...)
	if err != nil {
		return nil, err
	}

	return new(Empty).Unmarshal(resp)
}

func (c *shenzhenGoClient) Save(ctx context.Context, in *SaveRequest, opts ...grpcweb.CallOption) (*Empty, error) {
	resp, err := c.client.RPCCall(ctx, "Save", in.Marshal(), opts...)
	if err != nil {
		return nil, err
	}

	return new(Empty).Unmarshal(resp)
}

func (c *shenzhenGoClient) SetGraphProperties(ctx context.Context, in *SetGraphPropertiesRequest, opts ...grpcweb.CallOption) (*Empty, error) {
	resp, err := c.client.RPCCall(ctx, "SetGraphProperties", in.Marshal(), opts...)
	if err != nil {
		return nil, err
	}

	return new(Empty).Unmarshal(resp)
}

func (c *shenzhenGoClient) SetNodeProperties(ctx context.Context, in *SetNodePropertiesRequest, opts ...grpcweb.CallOption) (*Empty, error) {
	resp, err := c.client.RPCCall(ctx, "SetNodeProperties", in.Marshal(), opts...)
	if err != nil {
		return nil, err
	}

	return new(Empty).Unmarshal(resp)
}

func (c *shenzhenGoClient) SetPosition(ctx context.Context, in *SetPositionRequest, opts ...grpcweb.CallOption) (*Empty, error) {
	resp, err := c.client.RPCCall(ctx, "SetPosition", in.Marshal(), opts...)
	if err != nil {
		return nil, err
	}

	return new(Empty).Unmarshal(resp)
}
