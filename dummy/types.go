package dummy

import (
	"fmt"

	"github.com/rancher/kontainer-engine/types"
)

type NodePool struct {
	MinimumASGSize int64
	MaximumASGSize int64
	DesiredASGSize int64
	NodeVolumeSize *int64
	EBSEncryption  bool
	UserData       string
	InstanceType   string
	Image          string
	NodeCIDR       string
	PodCIDR        string
}

type state struct {
	ClusterName string
	DisplayName string
	DataCentre  string
	// ServiceCIDR string

	ClusterInfo types.ClusterInfo
}

func (s state) validate() error {
	nill := func(key string, val *string) (ok bool, err error) {
		if (val == nil) || (*val == "") {
			return false, fmt.Errorf("%s can not be nil", key)
		}
		return true, nil
	}

	if ok, err := nill("ClusterName", &s.ClusterName); !ok {
		return err
	}
	if ok, err := nill("DisplayName", &s.DisplayName); !ok {
		return err
	}
	if ok, err := nill("DataCentre", &s.DataCentre); !ok {
		return err
	}

	return nil
}
