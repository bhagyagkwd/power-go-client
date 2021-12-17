package main

import (
	"context"
	"log"

	v "github.com/IBM-Cloud/power-go-client/clients/instance"
	ps "github.com/IBM-Cloud/power-go-client/ibmpisession"
	"github.com/IBM-Cloud/power-go-client/power/models"
)

func main() {

	//session Inputs
	token := " < IAM TOKEN > "
	region := " < REGION > "
	accountID := " < ACCOUNT ID > "
	zone := " < ZONE > "

	// volume inputs
	piID := "< POWER INSTANCE ID >"
	volumes := make([]string, 1)
	volumes[0] = "ef82d430-4cb0-46a5-be11-e61fb129fe18"

	//capture inputs
	captureDestination := " < Capture Destination > "
	captureName := " < Capture Name > "
	instance_name := " < Name of vm > "
	CloudStorageImagePath := " < Cloud Storage Image Path > "
	CloudStorageAccessKey := "< Cloud Storage Access key>"
	CloudStorageSecretKey := "< Cloud Storage Secret key>"
	CloudStorageRegion := "< Cloud Storage Region >"

	session, err := ps.New(token, region, true, accountID, zone)
	if err != nil {
		log.Fatal(err)
	}
	powerClient := v.NewIBMPIInstanceClient(context.Background(), session, piID)
	if err != nil {
		log.Fatal(err)
	}

	body := &models.PVMInstanceCapture{
		CaptureDestination:    &captureDestination,
		CaptureName:           &captureName,
		CaptureVolumeIds:      volumes,
		CloudStorageAccessKey: CloudStorageAccessKey,
		CloudStorageImagePath: CloudStorageImagePath,
		CloudStorageRegion:    CloudStorageRegion,
		CloudStorageSecretKey: CloudStorageSecretKey,
	}
	createRespOk, err := powerClient.CaptureInstanceToImageCatalogV2(instance_name, body)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("***************[1]****************** %+v\n", *createRespOk)

}
