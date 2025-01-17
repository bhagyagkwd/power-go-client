package instance

import (
	"context"
	"fmt"

	"github.com/IBM-Cloud/power-go-client/helpers"
	"github.com/IBM-Cloud/power-go-client/ibmpisession"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_p_vm_instances"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/go-openapi/runtime"
)

/*  ChangeLog

2020-June-05 : Added the timeout variable to the clients since a lot of the SB / Powervc calls are timing out.

*/

// IBMPIInstanceClient ...
type IBMPIInstanceClient struct {
	session         *ibmpisession.IBMPISession
	cloudInstanceID string
	authInfo        runtime.ClientAuthInfoWriter
	ctx             context.Context
}

// NewIBMPIInstanceClient ...
func NewIBMPIInstanceClient(ctx context.Context, sess *ibmpisession.IBMPISession, cloudInstanceID string) *IBMPIInstanceClient {
	authInfo := ibmpisession.NewAuth(sess, cloudInstanceID)
	return &IBMPIInstanceClient{
		session:         sess,
		cloudInstanceID: cloudInstanceID,
		authInfo:        authInfo,
		ctx:             ctx,
	}
}

//Get information about a single pvm only
func (f *IBMPIInstanceClient) Get(id string) (*models.PVMInstance, error) {
	params := p_cloud_p_vm_instances.NewPcloudPvminstancesGetParams().
		WithContext(f.ctx).WithTimeout(helpers.PIGetTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithPvmInstanceID(id)
	resp, err := f.session.Power.PCloudPVMInstances.PcloudPvminstancesGet(params, f.authInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to Get PVM Instance %s :%v", id, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get PVM Instance %s", id)
	}
	return resp.Payload, nil
}

// GetAll Information about all the PVM Instances for a Client
func (f *IBMPIInstanceClient) GetAll() (*models.PVMInstances, error) {
	params := p_cloud_p_vm_instances.NewPcloudPvminstancesGetallParams().
		WithContext(f.ctx).WithTimeout(helpers.PIGetTimeOut).
		WithCloudInstanceID(f.cloudInstanceID)
	resp, err := f.session.Power.PCloudPVMInstances.PcloudPvminstancesGetall(params, f.authInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to Get all PVM Instances of Power Instance %s :%v", f.cloudInstanceID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get all PVM Instances of Power Instance %s", f.cloudInstanceID)
	}
	return resp.Payload, nil
}

//Create ...
func (f *IBMPIInstanceClient) Create(body *models.PVMInstanceCreate) (*models.PVMInstanceList, error) {
	params := p_cloud_p_vm_instances.NewPcloudPvminstancesPostParams().
		WithContext(f.ctx).WithTimeout(helpers.PICreateTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithBody(body)
	postok, postcreated, postAccepted, err := f.session.Power.PCloudPVMInstances.PcloudPvminstancesPost(params, ibmpisession.NewAuth(f.session, f.cloudInstanceID))
	if err != nil {
		return nil, fmt.Errorf("failed to Create PVM Instance :%v", err)
	}
	if postok != nil && len(postok.Payload) > 0 {
		return &postok.Payload, nil
	}
	if postcreated != nil && len(postcreated.Payload) > 0 {
		return &postcreated.Payload, nil
	}
	if postAccepted != nil && len(postAccepted.Payload) > 0 {
		return &postAccepted.Payload, nil
	}
	return nil, nil
}

// Delete PVM Instances
func (f *IBMPIInstanceClient) Delete(id string) error {
	params := p_cloud_p_vm_instances.NewPcloudPvminstancesDeleteParams().
		WithContext(f.ctx).WithTimeout(helpers.PIDeleteTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithPvmInstanceID(id)
	_, err := f.session.Power.PCloudPVMInstances.PcloudPvminstancesDelete(params, f.authInfo)
	if err != nil {
		return fmt.Errorf("failed to Delete PVM Instance %s :%s", id, err)
	}
	return nil
}

// Update PVM Instances
func (f *IBMPIInstanceClient) Update(id string, body *models.PVMInstanceUpdate) (*models.PVMInstanceUpdateResponse, error) {
	params := p_cloud_p_vm_instances.NewPcloudPvminstancesPutParams().
		WithContext(f.ctx).WithTimeout(helpers.PICreateTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithPvmInstanceID(id).WithBody(body)
	resp, err := f.session.Power.PCloudPVMInstances.PcloudPvminstancesPut(params, f.authInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to Update PVM Instance %s :%v", id, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Update PVM Instance %s", id)
	}
	return resp.Payload, nil
}

// Action PVM Instances Operations
func (f *IBMPIInstanceClient) Action(id string, body *models.PVMInstanceAction) error {
	params := p_cloud_p_vm_instances.NewPcloudPvminstancesActionPostParams().
		WithContext(f.ctx).WithTimeout(helpers.PICreateTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithPvmInstanceID(id).
		WithBody(body)
	_, err := f.session.Power.PCloudPVMInstances.PcloudPvminstancesActionPost(params, f.authInfo)
	if err != nil {
		return fmt.Errorf("failed to perform Action on PVM Instance %s :%v", id, err)
	}
	return nil

}

// PostConsoleURL Generate the Console URL
func (f *IBMPIInstanceClient) PostConsoleURL(id string) (*models.PVMInstanceConsole, error) {
	params := p_cloud_p_vm_instances.NewPcloudPvminstancesConsolePostParams().
		WithContext(f.ctx).WithTimeout(helpers.PICreateTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithPvmInstanceID(id)
	postok, err := f.session.Power.PCloudPVMInstances.PcloudPvminstancesConsolePost(params, f.authInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to Generate the Console URL PVM Instance %s :%v", id, err)
	}
	if postok == nil || postok.Payload == nil {
		return nil, fmt.Errorf("failed to Generate the Console URL PVM Instance %s", id)
	}
	return postok.Payload, nil
}

// List the available console languages for an instance
func (f *IBMPIInstanceClient) GetConsoleLanguages(id string) (*models.ConsoleLanguages, error) {
	params := p_cloud_p_vm_instances.NewPcloudPvminstancesConsoleGetParams().
		WithContext(f.ctx).WithTimeout(helpers.PIGetTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithPvmInstanceID(id)
	resp, err := f.session.Power.PCloudPVMInstances.PcloudPvminstancesConsoleGet(params, f.authInfo)
	if err != nil {
		return nil, err
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to get the Console Languages for PVM Instance %s", id)
	}
	return resp.Payload, nil
}

// List the available console languages for an instance
func (f *IBMPIInstanceClient) UpdateConsoleLanguage(body *models.ConsoleLanguage, id string) (*models.ConsoleLanguage, error) {
	params := p_cloud_p_vm_instances.NewPcloudPvminstancesConsolePutParams().
		WithContext(f.ctx).WithTimeout(helpers.PIUpdateTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithPvmInstanceID(id).
		WithBody(body)
	resp, err := f.session.Power.PCloudPVMInstances.PcloudPvminstancesConsolePut(params, f.authInfo)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// CaptureInstanceToImageCatalog Captures an instance
func (f *IBMPIInstanceClient) CaptureInstanceToImageCatalog(id string, body *models.PVMInstanceCapture) error {
	params := p_cloud_p_vm_instances.NewPcloudPvminstancesCapturePostParams().
		WithContext(f.ctx).WithTimeout(helpers.PIGetTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithPvmInstanceID(id).
		WithBody(body)
	_, _, err := f.session.Power.PCloudPVMInstances.PcloudPvminstancesCapturePost(params, f.authInfo)
	if err != nil {
		return fmt.Errorf("failed to Capture the PVM Instance %s: %v", id, err)
	}
	return nil

}

// CreatePvmSnapShot Create a snapshot of the instance
func (f *IBMPIInstanceClient) CreatePvmSnapShot(id string, body *models.SnapshotCreate) (*models.SnapshotCreateResponse, error) {
	params := p_cloud_p_vm_instances.NewPcloudPvminstancesSnapshotsPostParams().
		WithContext(f.ctx).WithTimeout(helpers.PICreateTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithPvmInstanceID(id).
		WithBody(body)
	snapshotpostaccepted, err := f.session.Power.PCloudPVMInstances.PcloudPvminstancesSnapshotsPost(params, f.authInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to Create the snapshot for the pvminstance %s: %s", id, err)
	}
	if snapshotpostaccepted == nil || snapshotpostaccepted.Payload == nil {
		return nil, fmt.Errorf("failed to Create the snapshot for the pvminstance %s", id)
	}
	return snapshotpostaccepted.Payload, nil
}

// CreateClone ...
func (f *IBMPIInstanceClient) CreateClone(id string, body *models.PVMInstanceClone) (*models.PVMInstance, error) {
	params := p_cloud_p_vm_instances.NewPcloudPvminstancesClonePostParams().
		WithContext(f.ctx).WithTimeout(helpers.PICreateTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithPvmInstanceID(id).
		WithBody(body)
	clonePost, err := f.session.Power.PCloudPVMInstances.PcloudPvminstancesClonePost(params, f.authInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to create the clone of the pvm instance %s: %v", id, err)
	}
	if clonePost == nil || clonePost.Payload == nil {
		return nil, fmt.Errorf("failed to create the clone of the pvm instance %s", id)
	}
	return clonePost.Payload, nil
}

// GetSnapShotVM Get information about the snapshots for a vm
func (f *IBMPIInstanceClient) GetSnapShotVM(id string) (*models.Snapshots, error) {
	params := p_cloud_p_vm_instances.NewPcloudPvminstancesSnapshotsGetallParams().
		WithContext(f.ctx).WithTimeout(helpers.PICreateTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithPvmInstanceID(id)
	resp, err := f.session.Power.PCloudPVMInstances.PcloudPvminstancesSnapshotsGetall(params, f.authInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to Get the snapshot for the pvminstance %s: %v", id, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to Get the snapshot for the pvminstance %s", id)
	}
	return resp.Payload, nil

}

// RestoreSnapShotVM Restore a snapshot
func (f *IBMPIInstanceClient) RestoreSnapShotVM(id, snapshotid, restoreAction string, body *models.SnapshotRestore) (*models.Snapshot, error) {
	params := p_cloud_p_vm_instances.NewPcloudPvminstancesSnapshotsRestorePostParams().
		WithContext(f.ctx).WithTimeout(helpers.PICreateTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithPvmInstanceID(id).
		WithSnapshotID(snapshotid).WithRestoreFailAction(&restoreAction).
		WithBody(body)
	resp, err := f.session.Power.PCloudPVMInstances.PcloudPvminstancesSnapshotsRestorePost(params, f.authInfo)
	if err != nil || resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to restrore the snapshot for the pvminstance %s: %v", id, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to restrore the snapshot for the pvminstance %s", id)
	}
	return resp.Payload, nil
}

// AddNetwork Add a network to the instance
func (f *IBMPIInstanceClient) AddNetwork(id string, body *models.PVMInstanceAddNetwork) (*models.PVMInstanceNetwork, error) {
	params := p_cloud_p_vm_instances.NewPcloudPvminstancesNetworksPostParams().
		WithContext(f.ctx).WithTimeout(helpers.PICreateTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithPvmInstanceID(id).
		WithBody(body)
	resp, err := f.session.Power.PCloudPVMInstances.PcloudPvminstancesNetworksPost(params, f.authInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to attach the network to the pvminstanceid %s: %v", id, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to attach the network to the pvminstanceid %s", id)
	}
	return resp.Payload, nil
}

// Delete a network from an instance
func (f *IBMPIInstanceClient) DeleteNetwork(id string, body *models.PVMInstanceRemoveNetwork) error {
	params := p_cloud_p_vm_instances.NewPcloudPvminstancesNetworksDeleteParams().
		WithContext(f.ctx).WithTimeout(helpers.PICreateTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithPvmInstanceID(id).
		WithBody(body)
	_, err := f.session.Power.PCloudPVMInstances.PcloudPvminstancesNetworksDelete(params, f.authInfo)
	if err != nil {
		return fmt.Errorf("failed to delete the network to the pvminstanceid %s: %v", id, err)
	}
	return nil
}
