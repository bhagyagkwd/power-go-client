package instance

import (
	"context"
	"fmt"

	"github.com/IBM-Cloud/power-go-client/helpers"
	"github.com/IBM-Cloud/power-go-client/ibmpisession"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_s_a_p"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/go-openapi/runtime"
)

// IBMPISAPInstanceClient ...
type IBMPISAPInstanceClient struct {
	session         *ibmpisession.IBMPISession
	cloudInstanceID string
	authInfo        runtime.ClientAuthInfoWriter
	ctx             context.Context
}

// NewIBMPISAPInstanceClient ...
func NewIBMPISAPInstanceClient(ctx context.Context, sess *ibmpisession.IBMPISession, cloudInstanceID string) *IBMPISAPInstanceClient {
	authInfo := ibmpisession.NewAuth(sess, cloudInstanceID)
	return &IBMPISAPInstanceClient{
		session:         sess,
		cloudInstanceID: cloudInstanceID,
		authInfo:        authInfo,
		ctx:             ctx,
	}
}

// Create a SAP instance
func (f *IBMPISAPInstanceClient) Create(body *models.SAPCreate) (*models.PVMInstanceList, error) {
	params := p_cloud_s_a_p.NewPcloudSapPostParams().
		WithContext(f.ctx).WithTimeout(helpers.PICreateTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithBody(body)
	postok, postcreated, postAccepted, err := f.session.Power.PCloudSAP.PcloudSapPost(params, f.authInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to Create SAP Instance :%v", err)
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

// Get SAP Profile
func (f *IBMPISAPInstanceClient) GetSAPProfile(id string) (*models.SAPProfile, error) {
	params := p_cloud_s_a_p.NewPcloudSapGetParams().
		WithContext(f.ctx).WithTimeout(helpers.PIGetTimeOut).
		WithCloudInstanceID(f.cloudInstanceID).WithSapProfileID(id)
	resp, err := f.session.Power.PCloudSAP.PcloudSapGet(params, f.authInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to get sap profile %s : %v", id, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to get sap profile %s", id)
	}
	return resp.Payload, nil
}

// GetAll SAP Profiles
func (f *IBMPISAPInstanceClient) GetAllSAPProfiles(cloudInstanceID string) (*models.SAPProfiles, error) {
	params := p_cloud_s_a_p.NewPcloudSapGetallParams().
		WithContext(f.ctx).WithTimeout(helpers.PIGetTimeOut).
		WithCloudInstanceID(f.cloudInstanceID)
	resp, err := f.session.Power.PCloudSAP.PcloudSapGetall(params, f.authInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to get all sap profiles for power instance %s : %s", cloudInstanceID, err)
	}
	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("failed to get all sap profiles for power instance %s", cloudInstanceID)
	}
	return resp.Payload, nil
}
