package proxy

import (
	"log"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/opensds/nbp/csi/util"
	"golang.org/x/net/context"
)

// Controller define
type Controller struct {
	client csi.ControllerClient
}

////////////////////////////////////////////////////////////////////////////////
//                            Controller Client Init                          //
////////////////////////////////////////////////////////////////////////////////

// GetController return controller struct
func GetController() (client Controller, err error) {
	conn, err := util.GetCSIClientConn()
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	c := csi.NewControllerClient(conn)

	return Controller{client: c}, nil
}

////////////////////////////////////////////////////////////////////////////////
//                            Controller Client Proxy                         //
////////////////////////////////////////////////////////////////////////////////

// CreateVolume proxy
func (c *Controller) CreateVolume(
	ctx context.Context,
	version *csi.Version,
	name string,
	capacity *csi.CapacityRange, /*Optional*/
	capabilities []*csi.VolumeCapability,
	params map[string]string, /*Optional*/
	credentials *csi.Credentials /*Optional*/) (volume *csi.VolumeInfo, err error) {

	req := &csi.CreateVolumeRequest{
		Version:            version,
		Name:               name,
		CapacityRange:      capacity,
		VolumeCapabilities: capabilities,
		Parameters:         params,
		UserCredentials:    credentials,
	}

	rs, err := c.client.CreateVolume(ctx, req)
	if err != nil {
		return nil, err
	}

	return rs.GetResult().VolumeInfo, nil
}

// DeleteVolume proxy
func (c *Controller) DeleteVolume(
	ctx context.Context,
	version *csi.Version,
	handle *csi.VolumeHandle,
	credentials *csi.Credentials /*Optional*/) error {

	req := &csi.DeleteVolumeRequest{
		Version:         version,
		VolumeHandle:    handle,
		UserCredentials: credentials,
	}

	_, err := c.client.DeleteVolume(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

// ControllerPublishVolume proxy
func (c *Controller) ControllerPublishVolume(
	ctx context.Context,
	version *csi.Version,
	handle *csi.VolumeHandle,
	nodeID *csi.NodeID, /*Optional*/
	capabilities *csi.VolumeCapability,
	readonly bool,
	credentials *csi.Credentials /*Optional*/) (*csi.PublishVolumeInfo, error) {

	req := &csi.ControllerPublishVolumeRequest{
		Version:          version,
		VolumeHandle:     handle,
		NodeId:           nodeID,
		VolumeCapability: capabilities,
		Readonly:         readonly,
		UserCredentials:  credentials,
	}

	rs, err := c.client.ControllerPublishVolume(ctx, req)
	if err != nil {
		return nil, err
	}

	return rs.GetResult().PublishVolumeInfo, nil
}

// ControllerUnpublishVolume proxy
func (c *Controller) ControllerUnpublishVolume(
	ctx context.Context,
	version *csi.Version,
	handle *csi.VolumeHandle,
	nodeID *csi.NodeID, /*Optional*/
	credentials *csi.Credentials /*Optional*/) error {

	req := &csi.ControllerUnpublishVolumeRequest{
		Version:         version,
		VolumeHandle:    handle,
		NodeId:          nodeID,
		UserCredentials: credentials,
	}

	_, err := c.client.ControllerUnpublishVolume(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

// ValidateVolumeCapabilities proxy
func (c *Controller) ValidateVolumeCapabilities(
	ctx context.Context,
	version *csi.Version,
	info *csi.VolumeInfo,
	capabilities []*csi.VolumeCapability) (*csi.ValidateVolumeCapabilitiesResponse_Result, error) {

	req := &csi.ValidateVolumeCapabilitiesRequest{
		Version:            version,
		VolumeInfo:         info,
		VolumeCapabilities: capabilities,
	}

	rs, err := c.client.ValidateVolumeCapabilities(ctx, req)
	if err != nil {
		return nil, err
	}

	return rs.GetResult(), nil
}

// ListVolumes proxy
func (c *Controller) ListVolumes(
	ctx context.Context,
	version *csi.Version,
	maxentries uint32, /*Optional*/
	startingtoken string /*Optional*/) (entries []*csi.ListVolumesResponse_Result_Entry, nextToken string, err error) {

	req := &csi.ListVolumesRequest{
		Version:       version,
		MaxEntries:    maxentries,
		StartingToken: startingtoken,
	}

	rs, err := c.client.ListVolumes(ctx, req)
	if err != nil {
		return nil, "", err
	}

	result := rs.GetResult()
	if len(result.Entries) == 0 {
		return nil, nextToken, nil
	}

	return result.Entries, result.NextToken, nil
}

// GetCapacity proxy
func (c *Controller) GetCapacity(
	ctx context.Context,
	version *csi.Version,
	capabilities []*csi.VolumeCapability /*Optional*/) (uint64, error) {

	req := &csi.GetCapacityRequest{
		Version:            version,
		VolumeCapabilities: capabilities,
	}

	rs, err := c.client.GetCapacity(ctx, req)
	if err != nil {
		return 0, err
	}

	return rs.GetResult().AvailableCapacity, nil
}

// ControllerGetCapabilities proxy
func (c *Controller) ControllerGetCapabilities(
	ctx context.Context,
	version *csi.Version) (capabilties []*csi.ControllerServiceCapability, err error) {

	req := &csi.ControllerGetCapabilitiesRequest{
		Version: version,
	}

	rs, err := c.client.ControllerGetCapabilities(ctx, req)
	if err != nil {
		return nil, err
	}

	return rs.GetResult().Capabilities, nil
}