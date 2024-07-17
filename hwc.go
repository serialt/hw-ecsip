package main

import (
	"fmt"
	"log/slog"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	ecs "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2/model"
	ecsregion "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2/region"
	eip "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/eip/v2"
	eipmodel "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/eip/v2/model"
	eipregion "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/eip/v2/region"
)

func NewECSClient(huawei Huaweicloud, region string) (client *ecs.EcsClient) {
	auth := basic.NewCredentialsBuilder().
		WithAk(huawei.AccessKeyId).
		WithSk(huawei.AccessKeySecret).
		Build()

	client = ecs.NewEcsClient(
		ecs.EcsClientBuilder().
			WithRegion(ecsregion.ValueOf(region)).
			WithCredential(auth).
			Build())

	return

}

func HWListServers(ecsClient *ecs.EcsClient) *[]model.ServerDetail {
	var listSerevrLimit int32 = 100
	// var listServerOffset int32 = 0
	request := &model.ListServersDetailsRequest{
		Limit: &listSerevrLimit,
		// Offset: &listServerOffset,
	}
	response, err := ecsClient.ListServersDetails(request)
	if err != nil {
		slog.Error("list serevr failed", "error", err)
		fmt.Println(err)
	}

	return response.Servers

}

func NewEIPClient(huawei Huaweicloud, region string) (client *eip.EipClient) {
	auth := basic.NewCredentialsBuilder().
		WithAk(huawei.AccessKeyId).
		WithSk(huawei.AccessKeySecret).
		Build()

	client = eip.NewEipClient(
		eip.EipClientBuilder().WithRegion(eipregion.ValueOf(region)).WithCredential(auth).
			Build())

	return

}

// []*eipmodel.PublicipShowResp
func HWEIPListIP(epiClient *eip.EipClient) (pubip *[]eipmodel.PublicipShowResp) {
	request := &eipmodel.ListPublicipsRequest{}
	response, err := epiClient.ListPublicips(request)
	if err != nil {
		slog.Info("get eip failed", "error", err)
		return
	}

	return response.Publicips

}
