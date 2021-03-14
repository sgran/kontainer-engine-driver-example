package dummy

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/imdario/mergo"
	"github.com/rancher/kontainer-engine/drivers/options"
	"github.com/rancher/kontainer-engine/types"
	"github.com/sirupsen/logrus"
)

// Driver is a thing
type Driver struct {
	createFlags *types.DriverFlags
	updateFlags *types.DriverFlags
	options     *types.DriverOptions

	types.UnimplementedClusterSizeAccess
	types.UnimplementedVersionAccess

	driverCapabilities types.Capabilities
}

func NewDriver() types.Driver {
	driver := &Driver{
		driverCapabilities: types.Capabilities{
			Capabilities: make(map[int64]bool),
		},
	}
	driver.driverCapabilities.AddCapability(types.GetVersionCapability)
	logrus.Infof("dummy driver initialised")

	return driver
}

// GetDriverCreateOptions returns cli flags that are used in create
func (d *Driver) GetDriverCreateOptions(ctx context.Context) (driverFlag *types.DriverFlags, err error) {
	logrus.Infof("dummy driver GetDriverCreateOptions")
	driverFlag = &types.DriverFlags{
		Options: make(map[string]*types.Flag),
	}

	driverFlag.Options["display-name"] = &types.Flag{
		Type:  types.StringType,
		Usage: "The displayed name of the cluster in the Rancher UI",
	}
	driverFlag.Options["kubernetes-version"] = &types.Flag{
		Type:    types.StringType,
		Usage:   "The kubernetes master version",
		Default: &types.Default{DefaultString: "1.18"},
	}
	driverFlag.Options["datacentre"] = &types.Flag{
		Type:    types.StringType,
		Usage:   "The datacenter for the cluster",
		Default: &types.Default{DefaultString: "DC12"},
	}

	return
}

// GetDriverUpdateOptions returns cli flags that are used in update
func (d *Driver) GetDriverUpdateOptions(ctx context.Context) (driverFlag *types.DriverFlags, err error) {
	logrus.Infof("dummy driver GetDriverUpdateOptions")

	driverFlag = &types.DriverFlags{
		Options: make(map[string]*types.Flag),
	}

	return
}

func getStateFromOptions(driverOptions *types.DriverOptions) (state, error) {
	s := state{}

	s.ClusterName = options.GetValueFromDriverOptions(driverOptions, types.StringType, "name").(string)
	s.DisplayName = options.GetValueFromDriverOptions(driverOptions, types.StringType, "display-name", "displayName").(string)
	s.DataCentre = options.GetValueFromDriverOptions(driverOptions, types.StringType, "datacentre").(string)
	s.ClusterInfo.Version = options.GetValueFromDriverOptions(driverOptions, types.StringType, "kubernetes-version").(string)

	return s, s.validate()
}

func storeState(info *types.ClusterInfo, state state) error {
	data, err := json.Marshal(state)

	if err != nil {
		return err
	}

	if info.Metadata == nil {
		info.Metadata = map[string]string{}
	}

	info.Metadata["state"] = string(data)

	return nil
}

func getState(info *types.ClusterInfo) (state, error) {
	state := state{}

	err := json.Unmarshal([]byte(info.Metadata["state"]), &state)
	if err != nil {
		logrus.Errorf("Error encountered while marshalling state: %v", err)
	}

	return state, err
}

// Create creates the cluster. clusterInfo is only set when we are retrying a failed or interrupted create
func (d *Driver) Create(ctx context.Context, opts *types.DriverOptions, clusterInfo *types.ClusterInfo) (info *types.ClusterInfo, err error) {
	logrus.Infof("dummy driver Create")

	state, err := getStateFromOptions(opts)
	if err != nil {
		return nil, fmt.Errorf("error parsing state: %v", err)
	}

	info = &types.ClusterInfo{}
	if err := storeState(info, state); err != nil {
		return nil, fmt.Errorf("error storing state")
	}

	return
}

// Update updates the cluster
func (d *Driver) Update(ctx context.Context, clusterInfo *types.ClusterInfo, opts *types.DriverOptions) (info *types.ClusterInfo, err error) {
	logrus.Infof("dummy driver Update")

	state, err := getState(clusterInfo)
	if err != nil {
		return nil, fmt.Errorf("error parsing state: %v", err)
	}
	new, err := getStateFromOptions(opts)
	if err != nil {
		return nil, fmt.Errorf("error parsing update: %v", err)
	}

	mergo.Merge(&new, &state)

	info = &types.ClusterInfo{}
	if err := storeState(info, new); err != nil {
		return nil, fmt.Errorf("error storing state")
	}

	return
}

// PostCheck does post action after provisioning
func (d *Driver) PostCheck(ctx context.Context, clusterInfo *types.ClusterInfo) (*types.ClusterInfo, error) {
	logrus.Infof("dummy driver PostCheck")

	return clusterInfo, nil
}

// Remove removes the cluster
func (d *Driver) Remove(ctx context.Context, clusterInfo *types.ClusterInfo) (err error) {
	logrus.Infof("dummy driver Remove")
	return
}

func (d *Driver) GetVersion(ctx context.Context, clusterInfo *types.ClusterInfo) (version *types.KubernetesVersion, err error) {
	logrus.Infof("dummy driver GetVersion")

	s, err := getState(clusterInfo)
	if err != nil {
		return nil, fmt.Errorf("error parsing state: %v", err)
	}

	version = &types.KubernetesVersion{
		Version: s.ClusterInfo.Version,
	}
	return
}

/*
func (d *Driver) SetVersion(ctx context.Context, clusterInfo *types.ClusterInfo, version *types.KubernetesVersion) (err error) {
	return
}

func (d *Driver) GetClusterSize(ctx context.Context, clusterInfo *types.ClusterInfo) (count *types.NodeCount, err error) {
	count = &types.NodeCount{
		Count: 3,
	}
	return
}

func (d *Driver) SetClusterSize(ctx context.Context, clusterInfo *types.ClusterInfo, count *types.NodeCount) (err error) {
	return fmt.Errorf("Setting Node Count not implemented")
}
*/
// Get driver capabilities
func (d *Driver) GetCapabilities(ctx context.Context) (caps *types.Capabilities, err error) {
	logrus.Infof("dummy driver GetCapabilities")

	return &d.driverCapabilities, nil
}

// Remove legacy service account token
func (d *Driver) RemoveLegacyServiceAccount(ctx context.Context, clusterInfo *types.ClusterInfo) (err error) {
	logrus.Infof("dummy driver RemoveLegacyServiceAccount")

	return
}

func (d *Driver) ETCDSave(ctx context.Context, clusterInfo *types.ClusterInfo, opts *types.DriverOptions, snapshotName string) (err error) {
	return fmt.Errorf("ETCD backup operations are not implemented")
}

func (d *Driver) ETCDRestore(ctx context.Context, clusterInfo *types.ClusterInfo, opts *types.DriverOptions, snapshotName string) (info *types.ClusterInfo, err error) {
	return nil, fmt.Errorf("ETCD backup operations are not implemented")
}

func (d *Driver) ETCDRemoveSnapshot(ctx context.Context, clusterInfo *types.ClusterInfo, opts *types.DriverOptions, snapshotName string) (err error) {
	return fmt.Errorf("ETCD backup operations are not implemented")
}

func (d *Driver) GetK8SCapabilities(ctx context.Context, opts *types.DriverOptions) (caps *types.K8SCapabilities, err error) {
	logrus.Infof("dummy driver GetK8SCapabilities")

	caps = &types.K8SCapabilities{
		NodePoolScalingSupported: false,
		L4LoadBalancer: &types.LoadBalancerCapabilities{
			Provider: "dummy",
			Enabled:  true,
			ProtocolsSupported: []string{
				"http",
				"https",
			},
			HealthCheckSupported: false,
		},
		IngressControllers: []*types.IngressCapabilities{},
	}

	return
}
