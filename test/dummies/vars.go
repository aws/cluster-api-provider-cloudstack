package dummies

import (
	csapi "github.com/apache/cloudstack-go/v2/cloudstack"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/pointer"
	capcv1 "sigs.k8s.io/cluster-api-provider-cloudstack/api/v1beta1"
	"sigs.k8s.io/cluster-api-provider-cloudstack/pkg/cloud"
	capiv1 "sigs.k8s.io/cluster-api/api/v1beta1"
)

var ( // Declare exported dummy vars.
	AffinityGroup      *cloud.AffinityGroup
	CSAffinityGroup    *capcv1.CloudStackAffinityGroup
	CSCluster          *capcv1.CloudStackCluster
	CAPIMachine        *capiv1.Machine
	CSMachine1         *capcv1.CloudStackMachine
	CAPICluster        *capiv1.Cluster
	CSMachineTemplate1 *capcv1.CloudStackMachineTemplate
	Zone1              capcv1.Zone
	Zone2              capcv1.Zone
	CSZone1            *capcv1.CloudStackZone
	CSZone2            *capcv1.CloudStackZone
	Net1               capcv1.Network
	Net2               capcv1.Network
	ISONet1            capcv1.Network
	CSISONet1          *capcv1.CloudStackIsolatedNetwork
	Domain             string
	DomainID           string
	RootDomain         string
	RootDomainID       string
	Level2Domain       string
	Level2DomainID     string
	Account            string
	Tags               map[string]string
	Tag1               map[string]string
	Tag2               map[string]string
	Tag1Key            string
	Tag1Val            string
	Tag2Key            string
	Tag2Val            string
	CSApiVersion       string
	CSClusterKind      string
	CSClusterName      string
	CSlusterNamespace  string
	TestTags           map[string]string
	CSClusterTagKey    string
	CSClusterTagVal    string
	CSClusterTag       map[string]string
	CreatedByCapcKey   string
	CreatedByCapcVal   string
	LBRuleID           string
	PublicIPID         string
	EndPointHost       string
	EndPointPort       int32
	ListDomainsParams  *csapi.ListDomainsParams
	ListDomainsResp    *csapi.ListDomainsResponse
	ListAccountsParams *csapi.ListAccountsParams
	ListAccountsResp   *csapi.ListAccountsResponse
	DiskOffering       = capcv1.CloudStackResourceDiskOffering{
		CloudStackResourceIdentifier: capcv1.CloudStackResourceIdentifier{
			Name: "Small",
		},
		MountPath:  "/data",
		Device:     "/dev/vdb",
		Filesystem: "ext4",
		Label:      "data_disk",
	}
)

// SetDummyVars sets/resets all dummy vars.
func SetDummyVars() {
	// These need to be in order as they build upon eachother.
	SetDummyZoneVars()
	SetDummyCAPCClusterVars()
	SetDummyCAPIClusterVars()
	SetDummyCAPIMachineVars()
	SetDummyCSMachineTemplateVars()
	SetDummyCSMachineVars()
	SetDummyTagVars()
	LBRuleID = "FakeLBRuleID"
}

func CAPCNetToCSAPINet(net *capcv1.Network) *csapi.Network {
	return &csapi.Network{
		Name: net.Name,
		Id:   net.ID,
		Type: net.Type,
	}
}

func CAPCZoneToCSAPIZone(net *capcv1.Zone) *csapi.Zone {
	return &csapi.Zone{
		Name: net.Name,
		Id:   net.ID,
	}
}

// SetDummyVars sets/resets tag related dummy vars.
func SetDummyTagVars() {
	CSClusterTagKey = "CAPC_cluster_" + string(CSCluster.ObjectMeta.UID)
	CSClusterTagVal = "1"
	CSClusterTag = map[string]string{CSClusterTagVal: CSClusterTagVal}
	CreatedByCapcKey = "create_by_CAPC"
	CreatedByCapcVal = ""
	Tag1Key = "test_tag1"
	Tag1Val = "arbitrary_value1"
	Tag2Key = "test_tag2"
	Tag2Val = "arbitrary_value2"
	Tag1 = map[string]string{Tag2Key: Tag2Val}
	Tag2 = map[string]string{Tag2Key: Tag2Val}
	Tags = map[string]string{Tag1Key: Tag1Val, Tag2Key: Tag2Val}
}

// SetDummyCSMachineTemplateVars resets the values in each of the exported CloudStackMachinesTemplate dummy variables.
func SetDummyCSMachineTemplateVars() {
	CSMachineTemplate1 = &capcv1.CloudStackMachineTemplate{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "infrastructure.cluster.x-k8s.io/v1beta1",
			Kind:       "CloudStackMachineTemplate",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-machinetemplate-1",
			Namespace: "default",
		},
		Spec: capcv1.CloudStackMachineTemplateSpec{
			Spec: capcv1.CloudStackMachineTemplateResource{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-machinetemplateresource",
					Namespace: "default",
				},
				Spec: capcv1.CloudStackMachineSpec{
					IdentityRef: &capcv1.CloudStackIdentityReference{
						Kind: "Secret",
						Name: "IdentitySecret",
					},
					Template: capcv1.CloudStackResourceIdentifier{
						Name: "Template",
					},
					Offering: capcv1.CloudStackResourceIdentifier{
						Name: "Offering",
					},
					DiskOffering: capcv1.CloudStackResourceDiskOffering{
						CloudStackResourceIdentifier: capcv1.CloudStackResourceIdentifier{
							Name: "DiskOffering",
						},
						MountPath:  "/data",
						Device:     "/dev/vdb",
						Filesystem: "ext4",
						Label:      "data_disk",
					},
					Details: map[string]string{
						"memoryOvercommitRatio": "1.2",
					},
				},
			},
		},
	}
}

// SetDummyCSMachineVars resets the values in each of the exported CloudStackMachine dummy variables.
func SetDummyCSMachineVars() {
	CSMachine1 = &capcv1.CloudStackMachine{
		TypeMeta: metav1.TypeMeta{
			APIVersion: CSApiVersion,
			Kind:       "CloudStackMachine",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-machine-2",
			Namespace: "default",
		},
		Spec: capcv1.CloudStackMachineSpec{
			IdentityRef: &capcv1.CloudStackIdentityReference{
				Kind: "Secret",
				Name: "IdentitySecret",
			},
			InstanceID: pointer.String("Instance1"),
			Template: capcv1.CloudStackResourceIdentifier{
				Name: "Template",
			},
			Offering: capcv1.CloudStackResourceIdentifier{
				Name: "Offering",
			},
			DiskOffering: capcv1.CloudStackResourceDiskOffering{
				CloudStackResourceIdentifier: capcv1.CloudStackResourceIdentifier{
					Name: "DiskOffering",
				},
				MountPath:  "/data",
				Device:     "/dev/vdb",
				Filesystem: "ext4",
				Label:      "data_disk",
			},
			AffinityGroupIDs: []string{"41eeb6e4-946f-4a18-b543-b2184815f1e4"},
			Details: map[string]string{
				"memoryOvercommitRatio": "1.2",
			},
		},
	}
	CSMachine1.ObjectMeta.SetName("test-vm")
}

func SetDummyZoneVars() {
	Zone1 = capcv1.Zone{Network: Net1}
	Zone1.Name = "Zone1"
	Zone1.ID = "FakeZone1ID"
	Zone2 = capcv1.Zone{Network: Net2}
	Zone2.Name = "Zone2"
	Zone2.ID = "FakeZone2ID"
	CSZone1 = &capcv1.CloudStackZone{Spec: capcv1.CloudStackZoneSpec(Zone1)}
	CSZone2 = &capcv1.CloudStackZone{Spec: capcv1.CloudStackZoneSpec(Zone2)}
}

// SetDummyCAPCClusterVars resets the values in each of the exported CloudStackCluster related dummy variables.
// It is intended to be called in BeforeEach() functions.
func SetDummyCAPCClusterVars() {
	Domain = "FakeDomainName"
	DomainID = "FakeDomainID"
	Level2Domain = "foo/FakeDomainName"
	Level2DomainID = "FakeLevel2DomainID"
	RootDomain = "ROOT"
	RootDomainID = "FakeRootDomainID"
	Account = "FakeAccountName"
	CSApiVersion = "infrastructure.cluster.x-k8s.io/v1beta1"
	CSClusterKind = "CloudStackCluster"
	CSClusterName = "test-cluster"
	EndPointHost = "EndpointHost"
	EndPointPort = int32(5309)
	PublicIPID = "FakePublicIPID"

	CSlusterNamespace = "default"
	AffinityGroup = &cloud.AffinityGroup{
		Name: "FakeAffinityGroup",
		Type: cloud.AffinityGroupType,
		ID:   "FakeAffinityGroupID"}
	CSAffinityGroup = &capcv1.CloudStackAffinityGroup{
		Spec: capcv1.CloudStackAffinityGroupSpec{Name: AffinityGroup.Name, Type: AffinityGroup.Type, ID: AffinityGroup.ID}}
	Net1 = capcv1.Network{Name: "SharedGuestNet1", Type: cloud.NetworkTypeShared, ID: "FakeSharedNetID1"}
	Net2 = capcv1.Network{Name: "SharedGuestNet2", Type: cloud.NetworkTypeShared, ID: "FakeSharedNetID2"}
	ISONet1 = capcv1.Network{Name: "IsoGuestNet1", Type: cloud.NetworkTypeIsolated, ID: "FakeIsolatedNetID1"}
	CSCluster = &capcv1.CloudStackCluster{
		TypeMeta: metav1.TypeMeta{
			APIVersion: CSApiVersion,
			Kind:       CSClusterKind,
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      CSClusterName,
			Namespace: "default",
			UID:       "0",
		},
		Spec: capcv1.CloudStackClusterSpec{
			IdentityRef: &capcv1.CloudStackIdentityReference{
				Kind: "Secret",
				Name: "IdentitySecret",
			},
			ControlPlaneEndpoint: capiv1.APIEndpoint{Host: EndPointHost, Port: EndPointPort},
			Zones:                []capcv1.Zone{Zone1, Zone2},
		},
		Status: capcv1.CloudStackClusterStatus{Zones: map[string]capcv1.Zone{}},
	}
	CSISONet1 = &capcv1.CloudStackIsolatedNetwork{
		Spec: capcv1.CloudStackIsolatedNetworkSpec{
			ControlPlaneEndpoint: CSCluster.Spec.ControlPlaneEndpoint}}
	CSISONet1.Spec.Name = ISONet1.Name
	CSISONet1.Spec.ID = ISONet1.ID
}

// SetDummyDomainAndAccount sets domain and account in the CSCluster Spec. This is not the default.
func SetDummyDomainAndAccount() {
	CSCluster.Spec.Account = Account
	CSCluster.Spec.Domain = Domain
}

// SetDummyDomainAndAccount sets domainID in the CSCluster Status. This is not the default.
func SetDummyDomainID() {
	CSCluster.Status.DomainID = "FakeDomainID"
}

// SetDummyCapiCluster resets the values in each of the exported CAPICluster related dummy variables.
func SetDummyCAPIClusterVars() {
	CAPICluster = &capiv1.Cluster{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: "capi-cluster-test-",
			Namespace:    "default",
		},
		Spec: capiv1.ClusterSpec{
			InfrastructureRef: &corev1.ObjectReference{
				APIVersion: capcv1.GroupVersion.String(),
				Kind:       "CloudStackCluster",
				Name:       "somename",
			},
		},
	}
}

func SetDummyIsoNetToNameOnly() {
	ISONet1.ID = ""
	ISONet1.Type = ""
	Zone1.Network = ISONet1
}

// Fills in cluster status vars.
func SetDummyClusterStatus() {
	CSCluster.Status.Zones = capcv1.ZoneStatusMap{Zone1.ID: Zone1, Zone2.ID: Zone2}
	CSCluster.Status.LBRuleID = LBRuleID
}

// Sets cluster spec to specified network.
func SetClusterSpecToNet(net *capcv1.Network) {
	Zone1.Network = *net
	CSCluster.Spec.Zones = []capcv1.Zone{Zone1}
}

func SetDummyCAPIMachineVars() {
	CAPIMachine = &capiv1.Machine{
		Spec: capiv1.MachineSpec{FailureDomain: pointer.String(Zone1.ID)},
	}
}

func SetDummyCSMachineStatuses() {
	CSMachine1.Status = capcv1.CloudStackMachineStatus{ZoneID: Zone1.ID}
}

func SetDummyCSApiResponse() {
	ListDomainsParams = &csapi.ListDomainsParams{}
	ListDomainsResp = &csapi.ListDomainsResponse{}
	ListDomainsResp.Count = 1
	ListDomainsResp.Domains = []*csapi.Domain{{Id: DomainID, Path: "ROOT/" + Domain}, {Id: RootDomainID, Path: "ROOT"}, {Id: Level2DomainID, Path: "ROOT/" + Level2Domain}}

	ListAccountsParams = &csapi.ListAccountsParams{}
	ListAccountsResp = &csapi.ListAccountsResponse{}
	ListAccountsResp.Count = 1
	ListAccountsResp.Accounts = []*csapi.Account{{Name: Account}}
}
