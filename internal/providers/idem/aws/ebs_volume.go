package aws

import (
	"github.com/infracost/infracost/internal/resources/aws"
	"github.com/infracost/infracost/internal/schema"
	"github.com/tidwall/gjson"
)

func GetEBSVolumeRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:  "states.aws.ec2.volume.present",
		RFunc: NewEBSVolume,
	}
}

func NewEBSVolume(d *schema.ResourceData, u *schema.UsageData) *schema.Resource {
	var size *int64
	if d.Get("size").Type != gjson.Null {
		size = intPtr(d.Get("size").Int())
	}

	a := &aws.EBSVolume{
		Address:    d.Address,
		Region:     d.Get("region").String(),
		Type:       d.Get("volume_type").String(),
		IOPS:       d.Get("iops").Int(),
		Throughput: d.Get("throughput").Int(),
		Size:       size,
	}

	a.PopulateUsage(u)

	return a.BuildResource()
}
