/*
 * Generated file models/networking/v4/config/config_model.go.
 *
 * Product version: 16.2.0-SNAPSHOT
 *
 * Part of the Networking API project
 *
 * (c) 2022 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module networking.v4.config of Networking API project
*/
package config
import (
  "bytes"
  import1 "github.com/nishantikagarg/sample-go-repo/networking_go_sdk/v16/models/common/v1/config"
  import2 "github.com/nishantikagarg/sample-go-repo/networking_go_sdk/v16/models/common/v1/response"
  "encoding/json"
  "errors"
  "fmt"
  import3 "github.com/nishantikagarg/sample-go-repo/networking_go_sdk/v16/models/networking/v4/error"
  import4 "github.com/nishantikagarg/sample-go-repo/networking_go_sdk/v16/models/prism/v4/config"
)

/**
Information pertaining to an assigned or reserved IP address on a subnet.
*/
type AddressAssignmentDto struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  AssignedDetails *AssignedAddressInfo `json:"assignedDetails,omitempty"`
  
  IpAddress *import1.IPAddress `json:"ipAddress,omitempty"`
  
  IsAssigned *bool `json:"isAssigned,omitempty"`
  
  IsLearned *bool `json:"isLearned,omitempty"`
  
  IsReserved *bool `json:"isReserved,omitempty"`
  
  LearnedDetails []LearnedAddressInfo `json:"learnedDetails,omitempty"`
  
  ReservedDetails *ReservedAddressInfo `json:"reservedDetails,omitempty"`
}

func NewAddressAssignmentDto() *AddressAssignmentDto {
  p := new(AddressAssignmentDto)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.AddressAssignmentDto"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.AddressAssignmentDto"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}





type AddressType int

const(
  ADDRESSTYPE_UNKNOWN AddressType = 0
  ADDRESSTYPE_REDACTED AddressType = 1
  ADDRESSTYPE_ANY AddressType = 2
  ADDRESSTYPE_EXTERNAL AddressType = 3
)

// returns the name of the enum given an ordinal number
func (e *AddressType) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "ANY",
    "EXTERNAL",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *AddressType) index(name string) AddressType {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "ANY",
    "EXTERNAL",
  }
  for idx := range names {
    if names[idx] == name {
      return AddressType(idx)
    }
  }
  return ADDRESSTYPE_UNKNOWN
}

func (e *AddressType) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for AddressType:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *AddressType) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e AddressType) Ref() *AddressType {
  return &e
}


/**
Address Type like "EXTERNAL" or "ANY".
*/
type AddressTypeObject struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  AddressType *AddressType `json:"addressType"`
}

func NewAddressTypeObject() *AddressTypeObject {
  p := new(AddressTypeObject)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.AddressTypeObject"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.AddressTypeObject"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}





type AncConfig struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Nameservers that resolve names of ANC services.
  */
  AncDomainNameServerList []import1.IPAddress `json:"ancDomainNameServerList,omitempty"`
  /**
  URL of Atlas Network Controller (ANC)
  */
  AncUrl *string `json:"ancUrl,omitempty"`
  /**
  True if atlas networking is enabled and false otherwise.
  */
  AtlasNetworkingEnabled *bool `json:"atlasNetworkingEnabled,omitempty"`
  /**
  Config version of the current ANC config. It is logical timestamp for now and will be updated by actual config version in future.
  */
  ConfigVersion *int64 `json:"configVersion,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import2.ApiLink `json:"links,omitempty"`
  
  Metadata *import1.Metadata `json:"metadata,omitempty"`
  /**
  OVN remote address of the Atlas Network Controller (ANC)
  */
  OvnRemoteAddress *string `json:"ovnRemoteAddress,omitempty"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
}

func NewAncConfig() *AncConfig {
  p := new(AncConfig)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.AncConfig"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.AncConfig"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST Response for  AncConfig type
*/
type AncConfigApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfAncConfigApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewAncConfigApiResponse() *AncConfigApiResponse {
  p := new(AncConfigApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.AncConfigApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.AncConfigApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *AncConfigApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *AncConfigApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfAncConfigApiResponseData()
  }
  e := p.Data.SetValue(v)
  if nil == e {
    if nil == p.DataItemDiscriminator_ {
      p.DataItemDiscriminator_ = new(string)
    }
    *p.DataItemDiscriminator_ = *p.Data.Discriminator
  }
  return e
}



/**
Information pertaining to an assigned IP address on a subnet.
*/
type AssignedAddressInfo struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  MacAddress *MacAddress `json:"macAddress,omitempty"`
  
  VmReference *import1.EntityReference `json:"vmReference,omitempty"`
}

func NewAssignedAddressInfo() *AssignedAddressInfo {
  p := new(AssignedAddressInfo)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.AssignedAddressInfo"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.AssignedAddressInfo"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Authentication algorithm
*/
type AuthenticationAlgorithm int

const(
  AUTHENTICATIONALGORITHM_UNKNOWN AuthenticationAlgorithm = 0
  AUTHENTICATIONALGORITHM_REDACTED AuthenticationAlgorithm = 1
  AUTHENTICATIONALGORITHM_MD5 AuthenticationAlgorithm = 2
  AUTHENTICATIONALGORITHM_SHA1 AuthenticationAlgorithm = 3
  AUTHENTICATIONALGORITHM_SHA256 AuthenticationAlgorithm = 4
  AUTHENTICATIONALGORITHM_SHA384 AuthenticationAlgorithm = 5
  AUTHENTICATIONALGORITHM_SHA512 AuthenticationAlgorithm = 6
)

// returns the name of the enum given an ordinal number
func (e *AuthenticationAlgorithm) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "MD5",
    "SHA1",
    "SHA256",
    "SHA384",
    "SHA512",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *AuthenticationAlgorithm) index(name string) AuthenticationAlgorithm {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "MD5",
    "SHA1",
    "SHA256",
    "SHA384",
    "SHA512",
  }
  for idx := range names {
    if names[idx] == name {
      return AuthenticationAlgorithm(idx)
    }
  }
  return AUTHENTICATIONALGORITHM_UNKNOWN
}

func (e *AuthenticationAlgorithm) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for AuthenticationAlgorithm:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *AuthenticationAlgorithm) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e AuthenticationAlgorithm) Ref() *AuthenticationAlgorithm {
  return &e
}


/**
Authentication type
*/
type AuthenticationType int

const(
  AUTHENTICATIONTYPE_UNKNOWN AuthenticationType = 0
  AUTHENTICATIONTYPE_REDACTED AuthenticationType = 1
  AUTHENTICATIONTYPE_PLAIN_TEXT AuthenticationType = 2
  AUTHENTICATIONTYPE_MD5 AuthenticationType = 3
)

// returns the name of the enum given an ordinal number
func (e *AuthenticationType) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "PLAIN_TEXT",
    "MD5",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *AuthenticationType) index(name string) AuthenticationType {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "PLAIN_TEXT",
    "MD5",
  }
  for idx := range names {
    if names[idx] == name {
      return AuthenticationType(idx)
    }
  }
  return AUTHENTICATIONTYPE_UNKNOWN
}

func (e *AuthenticationType) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for AuthenticationType:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *AuthenticationType) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e AuthenticationType) Ref() *AuthenticationType {
  return &e
}


/**
Azure config.
*/
type AzureConfig struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  External Subnet config list of Azure cloud.
  */
  ExternalSubnetConfigList []AzureExternalSubnetConfig `json:"externalSubnetConfigList"`
}

func NewAzureConfig() *AzureConfig {
  p := new(AzureConfig)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.AzureConfig"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.AzureConfig"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Binding of Atlas Flow Gateway external subnet with Azure external subnet.
*/
type AzureExternalSubnetBinding struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  AzureExternalNetworkPrefix *IPSubnet `json:"azureExternalNetworkPrefix"`
  
  PeerToPeerNetworkPrefix *IPSubnet `json:"peerToPeerNetworkPrefix,omitempty"`
  /**
  MAC Address of the Router port of Atlas Flow Gateway.
  */
  RouterPortMac *string `json:"routerPortMac"`
}

func NewAzureExternalSubnetBinding() *AzureExternalSubnetBinding {
  p := new(AzureExternalSubnetBinding)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.AzureExternalSubnetBinding"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.AzureExternalSubnetBinding"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
External Subnet config of Azure cloud.
*/
type AzureExternalSubnetConfig struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Gateway MAC address of the external subnet.
  */
  GatewayMacAddress *string `json:"gatewayMacAddress,omitempty"`
  
  IpConfig *IPConfig `json:"ipConfig"`
}

func NewAzureExternalSubnetConfig() *AzureExternalSubnetConfig {
  p := new(AzureExternalSubnetConfig)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.AzureExternalSubnetConfig"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.AzureExternalSubnetConfig"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}





type BgpAdvertisedPrefixes struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  IP prefixes learned from the remote gateway over BGP
  */
  ReceivedPrefixes []IPSubnet `json:"receivedPrefixes,omitempty"`
  /**
  IP prefixes advertised to the remote gateway over BGP
  */
  SentPrefixes []IPSubnet `json:"sentPrefixes,omitempty"`
}

func NewBgpAdvertisedPrefixes() *BgpAdvertisedPrefixes {
  p := new(BgpAdvertisedPrefixes)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.BgpAdvertisedPrefixes"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.BgpAdvertisedPrefixes"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
BGP configuration
*/
type BgpConfig struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Autonomous system number
  */
  Asn *int `json:"asn,omitempty"`
  /**
  BPG password
  */
  Password *string `json:"password,omitempty"`
  /**
  Redistribute routes over eBGP. Applicable only to network gateways deployed on VLAN subnets with eBGP over VPN.
  */
  RedistributeRoutes *bool `json:"redistributeRoutes,omitempty"`
}

func NewBgpConfig() *BgpConfig {
  p := new(BgpConfig)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.BgpConfig"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.BgpConfig"}
  p.UnknownFields_ = map[string]interface{}{}

  p.RedistributeRoutes = new(bool)
  *p.RedistributeRoutes = false


  return p
}




/**
BGP session.
*/
type BgpSession struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Advertised routes.
  */
  AdvertisedRoutes []Route `json:"advertisedRoutes,omitempty"`
  /**
  BGP session description.
  */
  Description *string `json:"description,omitempty"`
  /**
  The priority assigned to routes received over this BGP session.
  */
  DynamicRoutePriority *int `json:"dynamicRoutePriority,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  Received routes that are ignored either because the next hop is not L2-adjacent to the VPC or because the upper limit on received routes per session has been exceeded.
  */
  IgnoredRoutes []Route `json:"ignoredRoutes,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import2.ApiLink `json:"links,omitempty"`
  /**
  The local BGP gateway reference.
  */
  LocalGatewayReference *string `json:"localGatewayReference"`
  
  Metadata *import1.Metadata `json:"metadata,omitempty"`
  /**
  BGP session name.
  */
  Name *string `json:"name"`
  /**
  BGP password
  */
  Password *string `json:"password,omitempty"`
  /**
  Received routes.
  */
  ReceivedRoutes []Route `json:"receivedRoutes,omitempty"`
  /**
  The remote BGP gateway reference.
  */
  RemoteGatewayReference *string `json:"remoteGatewayReference"`
  /**
  Number of times the BGP session is dropped and reestablished per minute.
  */
  SessionDropsPerMinute *int `json:"sessionDropsPerMinute,omitempty"`
  
  Status *Status `json:"status,omitempty"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
}

func NewBgpSession() *BgpSession {
  p := new(BgpSession)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.BgpSession"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.BgpSession"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST Response for all response codes in api path /networking/v4.0.a1/config/bgp-sessions/{extId} Get operation
*/
type BgpSessionApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfBgpSessionApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewBgpSessionApiResponse() *BgpSessionApiResponse {
  p := new(BgpSessionApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.BgpSessionApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.BgpSessionApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *BgpSessionApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *BgpSessionApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfBgpSessionApiResponseData()
  }
  e := p.Data.SetValue(v)
  if nil == e {
    if nil == p.DataItemDiscriminator_ {
      p.DataItemDiscriminator_ = new(string)
    }
    *p.DataItemDiscriminator_ = *p.Data.Discriminator
  }
  return e
}



/**
REST Response for all response codes in api path /networking/v4.0.a1/config/bgp-sessions Get operation
*/
type BgpSessionListApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfBgpSessionListApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewBgpSessionListApiResponse() *BgpSessionListApiResponse {
  p := new(BgpSessionListApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.BgpSessionListApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.BgpSessionListApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *BgpSessionListApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *BgpSessionListApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfBgpSessionListApiResponseData()
  }
  e := p.Data.SetValue(v)
  if nil == e {
    if nil == p.DataItemDiscriminator_ {
      p.DataItemDiscriminator_ = new(string)
    }
    *p.DataItemDiscriminator_ = *p.Data.Discriminator
  }
  return e
}



/**
Cloud substrate of the network controller, for e.g. Azure.
*/
type CloudSubstrate int

const(
  CLOUDSUBSTRATE_UNKNOWN CloudSubstrate = 0
  CLOUDSUBSTRATE_REDACTED CloudSubstrate = 1
  CLOUDSUBSTRATE_AZURE CloudSubstrate = 2
)

// returns the name of the enum given an ordinal number
func (e *CloudSubstrate) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "AZURE",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *CloudSubstrate) index(name string) CloudSubstrate {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "AZURE",
  }
  for idx := range names {
    if names[idx] == name {
      return CloudSubstrate(idx)
    }
  }
  return CLOUDSUBSTRATE_UNKNOWN
}

func (e *CloudSubstrate) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for CloudSubstrate:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *CloudSubstrate) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e CloudSubstrate) Ref() *CloudSubstrate {
  return &e
}


/**
Get the Flow Networking usage of each registered Prism Element cluster.
*/
type ClusterFlowStatus struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Flow Networking usage status for every cluster.
  */
  ClusterStatusList []ClusterStatus `json:"clusterStatusList,omitempty"`
}

func NewClusterFlowStatus() *ClusterFlowStatus {
  p := new(ClusterFlowStatus)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.ClusterFlowStatus"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.ClusterFlowStatus"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST Response for all response codes in api path /networking/v4.0.a1/config/cluster-flow-status Get operation
*/
type ClusterFlowStatusApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfClusterFlowStatusApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewClusterFlowStatusApiResponse() *ClusterFlowStatusApiResponse {
  p := new(ClusterFlowStatusApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.ClusterFlowStatusApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.ClusterFlowStatusApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *ClusterFlowStatusApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *ClusterFlowStatusApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfClusterFlowStatusApiResponseData()
  }
  e := p.Data.SetValue(v)
  if nil == e {
    if nil == p.DataItemDiscriminator_ {
      p.DataItemDiscriminator_ = new(string)
    }
    *p.DataItemDiscriminator_ = *p.Data.Discriminator
  }
  return e
}



/**
Flow Networking usage status for a Prism Element cluster.
*/
type ClusterStatus struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  UUID of respective cluster.
  */
  ClusterReference *string `json:"clusterReference,omitempty"`
  /**
  True if cluster has at least one vNIC that is part of an Atlas subnet.
  */
  FlowStatus *bool `json:"flowStatus,omitempty"`
}

func NewClusterStatus() *ClusterStatus {
  p := new(ClusterStatus)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.ClusterStatus"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.ClusterStatus"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Current status of the network controller.
*/
type ControllerStatus int

const(
  CONTROLLERSTATUS_UNKNOWN ControllerStatus = 0
  CONTROLLERSTATUS_REDACTED ControllerStatus = 1
  CONTROLLERSTATUS_UP ControllerStatus = 2
  CONTROLLERSTATUS_DEGRADED ControllerStatus = 3
  CONTROLLERSTATUS_DOWN ControllerStatus = 4
)

// returns the name of the enum given an ordinal number
func (e *ControllerStatus) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "UP",
    "DEGRADED",
    "DOWN",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *ControllerStatus) index(name string) ControllerStatus {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "UP",
    "DEGRADED",
    "DOWN",
  }
  for idx := range names {
    if names[idx] == name {
      return ControllerStatus(idx)
    }
  }
  return CONTROLLERSTATUS_UNKNOWN
}

func (e *ControllerStatus) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for ControllerStatus:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *ControllerStatus) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e ControllerStatus) Ref() *ControllerStatus {
  return &e
}


/**
List of DHCP options to be configured.
*/
type DhcpOptions struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Bootfile name (option 67).
  */
  BootFileName *string `json:"bootFileName,omitempty"`
  /**
  The DNS domain name of the client (option 15).
  */
  DomainName *string `json:"domainName,omitempty"`
  /**
  List of Domain Name Server addresses (option 6).
  */
  DomainNameServers []import1.IPAddress `json:"domainNameServers,omitempty"`
  /**
  List of NTP Server addresses (option 42).
  */
  NtpServers []import1.IPAddress `json:"ntpServers,omitempty"`
  /**
  The DNS domain search list (option 119).
  */
  SearchDomains []string `json:"searchDomains,omitempty"`
  /**
  TFTP server name (option 66).
  */
  TftpServerName *string `json:"tftpServerName,omitempty"`
}

func NewDhcpOptions() *DhcpOptions {
  p := new(DhcpOptions)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.DhcpOptions"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.DhcpOptions"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Direct connect provides VPCs dedicated network connections
*/
type DirectConnect struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Direct connect bandwidth
  */
  BandwidthMbps *int `json:"bandwidthMbps"`
  /**
  Direct connect description field
  */
  Description *string `json:"description,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import2.ApiLink `json:"links,omitempty"`
  
  Metadata *import1.Metadata `json:"metadata,omitempty"`
  /**
  Direct connect name
  */
  Name *string `json:"name"`
  
  ProvisioningStatus *DirectConnectProvisioningStatus `json:"provisioningStatus,omitempty"`
  /**
  Direct connect service provider
  */
  ServiceProvider *string `json:"serviceProvider"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
  /**
  VPC
  */
  VpcReference *string `json:"vpcReference,omitempty"`
}

func NewDirectConnect() *DirectConnect {
  p := new(DirectConnect)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.DirectConnect"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.DirectConnect"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST Response for all response codes in api path /networking/v4.0.a1/config/direct-connects/{extId} Get operation
*/
type DirectConnectApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfDirectConnectApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDirectConnectApiResponse() *DirectConnectApiResponse {
  p := new(DirectConnectApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.DirectConnectApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.DirectConnectApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *DirectConnectApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *DirectConnectApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfDirectConnectApiResponseData()
  }
  e := p.Data.SetValue(v)
  if nil == e {
    if nil == p.DataItemDiscriminator_ {
      p.DataItemDiscriminator_ = new(string)
    }
    *p.DataItemDiscriminator_ = *p.Data.Discriminator
  }
  return e
}




type DirectConnectBgpConfig struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Nutanix local autonomous system number
  */
  LocalAsn *int `json:"localAsn,omitempty"`
  
  LocalBgpIp *import1.IPAddress `json:"localBgpIp,omitempty"`
  /**
  Password used in obtaining the md5 hash of TCP segments of the BGP peering session
  */
  Password *string `json:"password,omitempty"`
  /**
  Customer peering autonomous system number
  */
  PeerAsn *int `json:"peerAsn"`
  
  PeerBgpIp *import1.IPAddress `json:"peerBgpIp,omitempty"`
}

func NewDirectConnectBgpConfig() *DirectConnectBgpConfig {
  p := new(DirectConnectBgpConfig)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.DirectConnectBgpConfig"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.DirectConnectBgpConfig"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST Response for all response codes in api path /networking/v4.0.a1/config/direct-connects Get operation
*/
type DirectConnectListApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfDirectConnectListApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDirectConnectListApiResponse() *DirectConnectListApiResponse {
  p := new(DirectConnectListApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.DirectConnectListApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.DirectConnectListApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *DirectConnectListApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *DirectConnectListApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfDirectConnectListApiResponseData()
  }
  e := p.Data.SetValue(v)
  if nil == e {
    if nil == p.DataItemDiscriminator_ {
      p.DataItemDiscriminator_ = new(string)
    }
    *p.DataItemDiscriminator_ = *p.Data.Discriminator
  }
  return e
}



/**
Direct connect provisioning status
*/
type DirectConnectProvisioningStatus int

const(
  DIRECTCONNECTPROVISIONINGSTATUS_UNKNOWN DirectConnectProvisioningStatus = 0
  DIRECTCONNECTPROVISIONINGSTATUS_REDACTED DirectConnectProvisioningStatus = 1
  DIRECTCONNECTPROVISIONINGSTATUS_PENDING_CREATE DirectConnectProvisioningStatus = 2
  DIRECTCONNECTPROVISIONINGSTATUS_PENDING_UPDATE DirectConnectProvisioningStatus = 3
  DIRECTCONNECTPROVISIONINGSTATUS_PENDING_DELETE DirectConnectProvisioningStatus = 4
  DIRECTCONNECTPROVISIONINGSTATUS_PROVISIONED DirectConnectProvisioningStatus = 5
)

// returns the name of the enum given an ordinal number
func (e *DirectConnectProvisioningStatus) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "PENDING_CREATE",
    "PENDING_UPDATE",
    "PENDING_DELETE",
    "PROVISIONED",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *DirectConnectProvisioningStatus) index(name string) DirectConnectProvisioningStatus {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "PENDING_CREATE",
    "PENDING_UPDATE",
    "PENDING_DELETE",
    "PROVISIONED",
  }
  for idx := range names {
    if names[idx] == name {
      return DirectConnectProvisioningStatus(idx)
    }
  }
  return DIRECTCONNECTPROVISIONINGSTATUS_UNKNOWN
}

func (e *DirectConnectProvisioningStatus) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for DirectConnectProvisioningStatus:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *DirectConnectProvisioningStatus) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e DirectConnectProvisioningStatus) Ref() *DirectConnectProvisioningStatus {
  return &e
}


/**
Direct connect service provider
*/
type DirectConnectServiceProvider struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Offered connection bandwidths
  */
  ConnectionBandwidths []int `json:"connectionBandwidths,omitempty"`
  /**
  Service provider name
  */
  Name *string `json:"name,omitempty"`
}

func NewDirectConnectServiceProvider() *DirectConnectServiceProvider {
  p := new(DirectConnectServiceProvider)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.DirectConnectServiceProvider"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.DirectConnectServiceProvider"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST Response for all response codes in api path /networking/v4.0.a1/config/direct-connect-service-providers Get operation
*/
type DirectConnectServiceProviderListApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfDirectConnectServiceProviderListApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDirectConnectServiceProviderListApiResponse() *DirectConnectServiceProviderListApiResponse {
  p := new(DirectConnectServiceProviderListApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.DirectConnectServiceProviderListApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.DirectConnectServiceProviderListApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *DirectConnectServiceProviderListApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *DirectConnectServiceProviderListApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfDirectConnectServiceProviderListApiResponseData()
  }
  e := p.Data.SetValue(v)
  if nil == e {
    if nil == p.DataItemDiscriminator_ {
      p.DataItemDiscriminator_ = new(string)
    }
    *p.DataItemDiscriminator_ = *p.Data.Discriminator
  }
  return e
}



/**
Direct connect virtual interface defines a network interface over a direct connection
*/
type DirectConnectVirtualInterface struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  BgpAdvertisedPrefixes *BgpAdvertisedPrefixes `json:"bgpAdvertisedPrefixes,omitempty"`
  
  BgpConfig *DirectConnectBgpConfig `json:"bgpConfig"`
  /**
  Direct connect virtual interface description field
  */
  Description *string `json:"description,omitempty"`
  /**
  Direct connect to which this virtual interface belongs
  */
  DirectConnectReference *string `json:"directConnectReference"`
  /**
  Priority assigned to routes received on this virtual interface over BGP. A higher priority value indicates that the routes are more preferred
  */
  DynamicRoutePriority *int `json:"dynamicRoutePriority,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import2.ApiLink `json:"links,omitempty"`
  
  Metadata *import1.Metadata `json:"metadata,omitempty"`
  /**
  Direct connect virtual interface name
  */
  Name *string `json:"name"`
  
  State *DirectConnectVirtualInterfaceState `json:"state,omitempty"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
  /**
  VPC
  */
  VpcReference *string `json:"vpcReference"`
}

func NewDirectConnectVirtualInterface() *DirectConnectVirtualInterface {
  p := new(DirectConnectVirtualInterface)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.DirectConnectVirtualInterface"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.DirectConnectVirtualInterface"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST Response for all response codes in api path /networking/v4.0.a1/config/direct-connect-virtual-interfaces/{extId} Get operation
*/
type DirectConnectVirtualInterfaceApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfDirectConnectVirtualInterfaceApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDirectConnectVirtualInterfaceApiResponse() *DirectConnectVirtualInterfaceApiResponse {
  p := new(DirectConnectVirtualInterfaceApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.DirectConnectVirtualInterfaceApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.DirectConnectVirtualInterfaceApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *DirectConnectVirtualInterfaceApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *DirectConnectVirtualInterfaceApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfDirectConnectVirtualInterfaceApiResponseData()
  }
  e := p.Data.SetValue(v)
  if nil == e {
    if nil == p.DataItemDiscriminator_ {
      p.DataItemDiscriminator_ = new(string)
    }
    *p.DataItemDiscriminator_ = *p.Data.Discriminator
  }
  return e
}



/**
REST Response for all response codes in api path /networking/v4.0.a1/config/direct-connect-virtual-interfaces Get operation
*/
type DirectConnectVirtualInterfaceListApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfDirectConnectVirtualInterfaceListApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDirectConnectVirtualInterfaceListApiResponse() *DirectConnectVirtualInterfaceListApiResponse {
  p := new(DirectConnectVirtualInterfaceListApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.DirectConnectVirtualInterfaceListApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.DirectConnectVirtualInterfaceListApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *DirectConnectVirtualInterfaceListApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *DirectConnectVirtualInterfaceListApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfDirectConnectVirtualInterfaceListApiResponseData()
  }
  e := p.Data.SetValue(v)
  if nil == e {
    if nil == p.DataItemDiscriminator_ {
      p.DataItemDiscriminator_ = new(string)
    }
    *p.DataItemDiscriminator_ = *p.Data.Discriminator
  }
  return e
}



/**
Direct connect virtual interface reference
*/
type DirectConnectVirtualInterfaceReference struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Direct connect virtual interface identifier
  */
  ExtId *string `json:"extId"`
  
  ObjectType *string `json:"objectType,omitempty"`
}

func NewDirectConnectVirtualInterfaceReference() *DirectConnectVirtualInterfaceReference {
  p := new(DirectConnectVirtualInterfaceReference)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.DirectConnectVirtualInterfaceReference"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.DirectConnectVirtualInterfaceReference"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
The state of the direct connect virtual interface
*/
type DirectConnectVirtualInterfaceState int

const(
  DIRECTCONNECTVIRTUALINTERFACESTATE_UNKNOWN DirectConnectVirtualInterfaceState = 0
  DIRECTCONNECTVIRTUALINTERFACESTATE_REDACTED DirectConnectVirtualInterfaceState = 1
  DIRECTCONNECTVIRTUALINTERFACESTATE_UP DirectConnectVirtualInterfaceState = 2
  DIRECTCONNECTVIRTUALINTERFACESTATE_DOWN DirectConnectVirtualInterfaceState = 3
)

// returns the name of the enum given an ordinal number
func (e *DirectConnectVirtualInterfaceState) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "UP",
    "DOWN",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *DirectConnectVirtualInterfaceState) index(name string) DirectConnectVirtualInterfaceState {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "UP",
    "DOWN",
  }
  for idx := range names {
    if names[idx] == name {
      return DirectConnectVirtualInterfaceState(idx)
    }
  }
  return DIRECTCONNECTVIRTUALINTERFACESTATE_UNKNOWN
}

func (e *DirectConnectVirtualInterfaceState) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for DirectConnectVirtualInterfaceState:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *DirectConnectVirtualInterfaceState) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e DirectConnectVirtualInterfaceState) Ref() *DirectConnectVirtualInterfaceState {
  return &e
}


/**
Dead Peer Detection configuration for the VPN connection
*/
type DpdConfig struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  The amount of time the peer waits for traffic before sending a DPD request
  */
  IntervalSecs *int64 `json:"intervalSecs,omitempty"`
  
  Operation *DpdOperation `json:"operation,omitempty"`
  /**
  The maximum amount of time to wait for a DPD response before marking the peer as dead
  */
  TimeoutSecs *int64 `json:"timeoutSecs,omitempty"`
}

func NewDpdConfig() *DpdConfig {
  p := new(DpdConfig)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.DpdConfig"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.DpdConfig"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Operation to be performed on detecting a dead peer
*/
type DpdOperation int

const(
  DPDOPERATION_UNKNOWN DpdOperation = 0
  DPDOPERATION_REDACTED DpdOperation = 1
  DPDOPERATION_RESTART DpdOperation = 2
  DPDOPERATION_CLEAR DpdOperation = 3
  DPDOPERATION_HOLD DpdOperation = 4
)

// returns the name of the enum given an ordinal number
func (e *DpdOperation) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "RESTART",
    "CLEAR",
    "HOLD",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *DpdOperation) index(name string) DpdOperation {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "RESTART",
    "CLEAR",
    "HOLD",
  }
  for idx := range names {
    if names[idx] == name {
      return DpdOperation(idx)
    }
  }
  return DPDOPERATION_UNKNOWN
}

func (e *DpdOperation) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for DpdOperation:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *DpdOperation) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e DpdOperation) Ref() *DpdOperation {
  return &e
}


/**
Encryption algorithm
*/
type EncryptionAlgorithm int

const(
  ENCRYPTIONALGORITHM_UNKNOWN EncryptionAlgorithm = 0
  ENCRYPTIONALGORITHM_REDACTED EncryptionAlgorithm = 1
  ENCRYPTIONALGORITHM_AES128 EncryptionAlgorithm = 2
  ENCRYPTIONALGORITHM_AES256 EncryptionAlgorithm = 3
  ENCRYPTIONALGORITHM_TRIPLE_DES EncryptionAlgorithm = 4
)

// returns the name of the enum given an ordinal number
func (e *EncryptionAlgorithm) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "AES128",
    "AES256",
    "TRIPLE_DES",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *EncryptionAlgorithm) index(name string) EncryptionAlgorithm {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "AES128",
    "AES256",
    "TRIPLE_DES",
  }
  for idx := range names {
    if names[idx] == name {
      return EncryptionAlgorithm(idx)
    }
  }
  return ENCRYPTIONALGORITHM_UNKNOWN
}

func (e *EncryptionAlgorithm) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for EncryptionAlgorithm:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *EncryptionAlgorithm) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e EncryptionAlgorithm) Ref() *EncryptionAlgorithm {
  return &e
}


/**
Contains the uuid and scope type information for a particular export scope.
*/
type ExportScope struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  ScopeType *ScopeType `json:"scopeType,omitempty"`
  
  Uuid *string `json:"uuid,omitempty"`
}

func NewExportScope() *ExportScope {
  p := new(ExportScope)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.ExportScope"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.ExportScope"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Enum for protocol, can be one of TCP or UDP.
*/
type ExporterProtocol int

const(
  EXPORTERPROTOCOL_UNKNOWN ExporterProtocol = 0
  EXPORTERPROTOCOL_REDACTED ExporterProtocol = 1
  EXPORTERPROTOCOL_TCP ExporterProtocol = 2
  EXPORTERPROTOCOL_UDP ExporterProtocol = 3
)

// returns the name of the enum given an ordinal number
func (e *ExporterProtocol) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "TCP",
    "UDP",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *ExporterProtocol) index(name string) ExporterProtocol {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "TCP",
    "UDP",
  }
  for idx := range names {
    if names[idx] == name {
      return ExporterProtocol(idx)
    }
  }
  return EXPORTERPROTOCOL_UNKNOWN
}

func (e *ExporterProtocol) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for ExporterProtocol:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *ExporterProtocol) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e ExporterProtocol) Ref() *ExporterProtocol {
  return &e
}


/**
Information about the external subnet, SNAT IPs and the gateway nodes.
*/
type ExternalSubnet struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  ActiveGatewayNode *GatewayNodeReference `json:"activeGatewayNode,omitempty"`
  /**
  List of IP Addresses used for SNAT, if NAT is enabled on the external subnet. If NAT is not enabled, this specifies the IP address of the VPC port connected to the external gateway.
  */
  ExternalIps []import1.IPAddress `json:"externalIps,omitempty"`
  /**
  List of gateway nodes that can be used for external connectivity.
  */
  GatewayNodes []string `json:"gatewayNodes,omitempty"`
  /**
  External subnet reference.
  */
  SubnetReference *string `json:"subnetReference"`
}

func NewExternalSubnet() *ExternalSubnet {
  p := new(ExternalSubnet)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.ExternalSubnet"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.ExternalSubnet"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Schema to configure a floating IP.
*/
type FloatingIp struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  AssociationItemDiscriminator_ *string `json:"$associationItemDiscriminator,omitempty"`
  /**
  Association of the Floating IP with either NIC or Private IP
  */
  Association *OneOfFloatingIpAssociation `json:"association,omitempty"`
  /**
  Description for the Floating IP.
  */
  Description *string `json:"description,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  External Subnet Reference for the floating IP to be allocated in on-prem only.
  */
  ExternalSubnetReference *string `json:"externalSubnetReference,omitempty"`
  
  FloatingIp *import1.IPAddress `json:"floatingIp,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import2.ApiLink `json:"links,omitempty"`
  
  Metadata *import1.Metadata `json:"metadata,omitempty"`
  /**
  Name of the floating IP.
  */
  Name *string `json:"name,omitempty"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
}

func NewFloatingIp() *FloatingIp {
  p := new(FloatingIp)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.FloatingIp"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.FloatingIp"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST Response for all response codes in api path /networking/v4.0.a1/config/floating-ips/{extId} Get operation
*/
type FloatingIpApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfFloatingIpApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewFloatingIpApiResponse() *FloatingIpApiResponse {
  p := new(FloatingIpApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.FloatingIpApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.FloatingIpApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *FloatingIpApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *FloatingIpApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfFloatingIpApiResponseData()
  }
  e := p.Data.SetValue(v)
  if nil == e {
    if nil == p.DataItemDiscriminator_ {
      p.DataItemDiscriminator_ = new(string)
    }
    *p.DataItemDiscriminator_ = *p.Data.Discriminator
  }
  return e
}



/**
REST Response for all response codes in api path /networking/v4.0.a1/config/floating-ips Get operation
*/
type FloatingIpListApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfFloatingIpListApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewFloatingIpListApiResponse() *FloatingIpListApiResponse {
  p := new(FloatingIpListApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.FloatingIpListApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.FloatingIpListApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *FloatingIpListApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *FloatingIpListApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfFloatingIpListApiResponseData()
  }
  e := p.Data.SetValue(v)
  if nil == e {
    if nil == p.DataItemDiscriminator_ {
      p.DataItemDiscriminator_ = new(string)
    }
    *p.DataItemDiscriminator_ = *p.Data.Discriminator
  }
  return e
}




type FlowGateway struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  List of bindings of Atlas Flow Gateway external subnet with Azure external subnet.
  */
  AzureExternalSubnetBindingList []AzureExternalSubnetBinding `json:"azureExternalSubnetBindingList"`
  /**
  Chassis UUID of the Atlas Flow Gateway.
  */
  ChassisUuid *string `json:"chassisUuid"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  
  GatewayStatus *FlowGatewayStatus `json:"gatewayStatus,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import2.ApiLink `json:"links,omitempty"`
  
  Metadata *import1.Metadata `json:"metadata,omitempty"`
  /**
  Version of the OVN controller
  */
  OvnControllerVersion *string `json:"ovnControllerVersion,omitempty"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
}

func NewFlowGateway() *FlowGateway {
  p := new(FlowGateway)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.FlowGateway"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.FlowGateway"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST Response for  FlowGateway type
*/
type FlowGatewayApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfFlowGatewayApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewFlowGatewayApiResponse() *FlowGatewayApiResponse {
  p := new(FlowGatewayApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.FlowGatewayApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.FlowGatewayApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *FlowGatewayApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *FlowGatewayApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfFlowGatewayApiResponseData()
  }
  e := p.Data.SetValue(v)
  if nil == e {
    if nil == p.DataItemDiscriminator_ {
      p.DataItemDiscriminator_ = new(string)
    }
    *p.DataItemDiscriminator_ = *p.Data.Discriminator
  }
  return e
}




type FlowGatewayKeepAliveRequest struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Chassis UUID of the Atlas Flow Gateway.
  */
  ChassisUuid *string `json:"chassisUuid"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  
  GatewayStatus *FlowGatewayStatus `json:"gatewayStatus,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import2.ApiLink `json:"links,omitempty"`
  /**
  Version of the OVN controller
  */
  OvnControllerVersion *string `json:"ovnControllerVersion,omitempty"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
}

func NewFlowGatewayKeepAliveRequest() *FlowGatewayKeepAliveRequest {
  p := new(FlowGatewayKeepAliveRequest)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.FlowGatewayKeepAliveRequest"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.FlowGatewayKeepAliveRequest"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}





type FlowGatewayKeepAliveResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  Tells if keep alive was successful or not.
  */
  KeepAliveResponse *string `json:"keepAliveResponse,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import2.ApiLink `json:"links,omitempty"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
}

func NewFlowGatewayKeepAliveResponse() *FlowGatewayKeepAliveResponse {
  p := new(FlowGatewayKeepAliveResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.FlowGatewayKeepAliveResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.FlowGatewayKeepAliveResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST Response for  FlowGatewayKeepAliveResponse type
*/
type FlowGatewayKeepAliveResponseApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfFlowGatewayKeepAliveResponseApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewFlowGatewayKeepAliveResponseApiResponse() *FlowGatewayKeepAliveResponseApiResponse {
  p := new(FlowGatewayKeepAliveResponseApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.FlowGatewayKeepAliveResponseApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.FlowGatewayKeepAliveResponseApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *FlowGatewayKeepAliveResponseApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *FlowGatewayKeepAliveResponseApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfFlowGatewayKeepAliveResponseApiResponseData()
  }
  e := p.Data.SetValue(v)
  if nil == e {
    if nil == p.DataItemDiscriminator_ {
      p.DataItemDiscriminator_ = new(string)
    }
    *p.DataItemDiscriminator_ = *p.Data.Discriminator
  }
  return e
}



/**
REST Response for  list of  FlowGateway types
*/
type FlowGatewayListApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfFlowGatewayListApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewFlowGatewayListApiResponse() *FlowGatewayListApiResponse {
  p := new(FlowGatewayListApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.FlowGatewayListApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.FlowGatewayListApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *FlowGatewayListApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *FlowGatewayListApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfFlowGatewayListApiResponseData()
  }
  e := p.Data.SetValue(v)
  if nil == e {
    if nil == p.DataItemDiscriminator_ {
      p.DataItemDiscriminator_ = new(string)
    }
    *p.DataItemDiscriminator_ = *p.Data.Discriminator
  }
  return e
}




type FlowGatewayState int

const(
  FLOWGATEWAYSTATE_UNKNOWN FlowGatewayState = 0
  FLOWGATEWAYSTATE_REDACTED FlowGatewayState = 1
  FLOWGATEWAYSTATE_HEALTHY FlowGatewayState = 2
  FLOWGATEWAYSTATE_DOWN FlowGatewayState = 3
  FLOWGATEWAYSTATE_PROVISIONING FlowGatewayState = 4
  FLOWGATEWAYSTATE_MAINTENANCE FlowGatewayState = 5
)

// returns the name of the enum given an ordinal number
func (e *FlowGatewayState) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "HEALTHY",
    "DOWN",
    "PROVISIONING",
    "MAINTENANCE",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *FlowGatewayState) index(name string) FlowGatewayState {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "HEALTHY",
    "DOWN",
    "PROVISIONING",
    "MAINTENANCE",
  }
  for idx := range names {
    if names[idx] == name {
      return FlowGatewayState(idx)
    }
  }
  return FLOWGATEWAYSTATE_UNKNOWN
}

func (e *FlowGatewayState) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for FlowGatewayState:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *FlowGatewayState) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e FlowGatewayState) Ref() *FlowGatewayState {
  return &e
}


/**
Status of the Atlas Flow Gateway.
*/
type FlowGatewayStatus struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  Detail *string `json:"detail,omitempty"`
  
  State *FlowGatewayState `json:"state,omitempty"`
}

func NewFlowGatewayStatus() *FlowGatewayStatus {
  p := new(FlowGatewayStatus)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.FlowGatewayStatus"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.FlowGatewayStatus"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}





type GatewayNodeReference struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  NodeId *string `json:"nodeId,omitempty"`
  
  NodeIpAddress *import1.IPAddress `json:"nodeIpAddress,omitempty"`
}

func NewGatewayNodeReference() *GatewayNodeReference {
  p := new(GatewayNodeReference)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.GatewayNodeReference"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.GatewayNodeReference"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Local gateway role (acceptor or initiator) in the connection
*/
type GatewayRole int

const(
  GATEWAYROLE_UNKNOWN GatewayRole = 0
  GATEWAYROLE_REDACTED GatewayRole = 1
  GATEWAYROLE_INITIATOR GatewayRole = 2
  GATEWAYROLE_ACCEPTOR GatewayRole = 3
)

// returns the name of the enum given an ordinal number
func (e *GatewayRole) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "INITIATOR",
    "ACCEPTOR",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *GatewayRole) index(name string) GatewayRole {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "INITIATOR",
    "ACCEPTOR",
  }
  for idx := range names {
    if names[idx] == name {
      return GatewayRole(idx)
    }
  }
  return GATEWAYROLE_UNKNOWN
}

func (e *GatewayRole) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for GatewayRole:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *GatewayRole) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e GatewayRole) Ref() *GatewayRole {
  return &e
}


/**
ICMP parameters to be matched in Routing Policy.
*/
type ICMP struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  IcmpCode *int `json:"icmpCode,omitempty"`
  
  IcmpType *int `json:"icmpType,omitempty"`
}

func NewICMP() *ICMP {
  p := new(ICMP)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.ICMP"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.ICMP"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
ICMP parameters to be matched in Routing Policy.
*/
type ICMPObject struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  Icmp *ICMP `json:"icmp"`
}

func NewICMPObject() *ICMPObject {
  p := new(ICMPObject)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.ICMPObject"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.ICMPObject"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
IP configuration.
*/
type IPConfig struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  Ipv4 *IPv4Config `json:"ipv4,omitempty"`
  
  Ipv6 *IPv6Config `json:"ipv6,omitempty"`
}

func NewIPConfig() *IPConfig {
  p := new(IPConfig)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.IPConfig"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.IPConfig"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}


func (i *IPConfig) HasIpv4() bool {
  return i.Ipv4 != nil
}
func (i *IPConfig) HasIpv6() bool {
  return i.Ipv6 != nil
}

func (i *IPConfig) IsValid() bool {
  return i.HasIpv4() || i.HasIpv6()
}



type IPFIXExporter struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  The IP address of the IPFIX collector.
  */
  CollectorIp *string `json:"collectorIp"`
  /**
  The port number of the IPFIX collector.
  */
  CollectorPort *int64 `json:"collectorPort"`
  /**
  IPFIX Exporter description.
  */
  Description *string `json:"description,omitempty"`
  /**
  The maximum export rate in bits per second(bps) at which the exporter should try to export data.
  */
  ExportRateLimitPerNode *int64 `json:"exportRateLimitPerNode,omitempty"`
  /**
  List of IPFIX Exporter scopes.
  */
  ExportScopes []ExportScope `json:"exportScopes"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import2.ApiLink `json:"links,omitempty"`
  
  Metadata *import1.Metadata `json:"metadata,omitempty"`
  /**
  Name of the IPFIX Exporter.
  */
  Name *string `json:"name"`
  
  Protocol *ExporterProtocol `json:"protocol"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
}

func NewIPFIXExporter() *IPFIXExporter {
  p := new(IPFIXExporter)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.IPFIXExporter"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.IPFIXExporter"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST Response for all response codes in api path /networking/v4.0.a1/config/ipfix-exporters/{extId} Get operation
*/
type IPFIXExporterApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfIPFIXExporterApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewIPFIXExporterApiResponse() *IPFIXExporterApiResponse {
  p := new(IPFIXExporterApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.IPFIXExporterApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.IPFIXExporterApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *IPFIXExporterApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *IPFIXExporterApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfIPFIXExporterApiResponseData()
  }
  e := p.Data.SetValue(v)
  if nil == e {
    if nil == p.DataItemDiscriminator_ {
      p.DataItemDiscriminator_ = new(string)
    }
    *p.DataItemDiscriminator_ = *p.Data.Discriminator
  }
  return e
}



/**
REST Response for all response codes in api path /networking/v4.0.a1/config/ipfix-exporters Get operation
*/
type IPFIXExporterListApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfIPFIXExporterListApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewIPFIXExporterListApiResponse() *IPFIXExporterListApiResponse {
  p := new(IPFIXExporterListApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.IPFIXExporterListApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.IPFIXExporterListApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *IPFIXExporterListApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *IPFIXExporterListApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfIPFIXExporterListApiResponseData()
  }
  e := p.Data.SetValue(v)
  if nil == e {
    if nil == p.DataItemDiscriminator_ {
      p.DataItemDiscriminator_ = new(string)
    }
    *p.DataItemDiscriminator_ = *p.Data.Discriminator
  }
  return e
}




type IPSubnet struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  Ipv4 *IPv4Subnet `json:"ipv4,omitempty"`
  
  Ipv6 *IPv6Subnet `json:"ipv6,omitempty"`
}

func NewIPSubnet() *IPSubnet {
  p := new(IPSubnet)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.IPSubnet"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.IPSubnet"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}


func (i *IPSubnet) HasIpv4() bool {
  return i.Ipv4 != nil
}
func (i *IPSubnet) HasIpv6() bool {
  return i.Ipv6 != nil
}

func (i *IPSubnet) IsValid() bool {
  return i.HasIpv4() || i.HasIpv6()
}


/**
IP address and prefix length of the subnet.
*/
type IPSubnetObject struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  IpSubnet *IPSubnet `json:"ipSubnet"`
}

func NewIPSubnetObject() *IPSubnetObject {
  p := new(IPSubnetObject)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.IPSubnetObject"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.IPSubnetObject"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
IP V4 configuration.
*/
type IPv4Config struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  DefaultGatewayIp *import1.IPv4Address `json:"defaultGatewayIp,omitempty"`
  
  DhcpServerAddress *import1.IPv4Address `json:"dhcpServerAddress,omitempty"`
  
  IpSubnet *IPv4Subnet `json:"ipSubnet"`
  /**
  Pool of IP addresses from where IPs are allocated.
  */
  PoolList []IPv4Pool `json:"poolList,omitempty"`
}

func NewIPv4Config() *IPv4Config {
  p := new(IPv4Config)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.IPv4Config"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.IPv4Config"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Start/end IP address range.
*/
type IPv4Pool struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  EndIP *import1.IPv4Address `json:"endIP"`
  
  StartIP *import1.IPv4Address `json:"startIP"`
}

func NewIPv4Pool() *IPv4Pool {
  p := new(IPv4Pool)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.IPv4Pool"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.IPv4Pool"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}





type IPv4Subnet struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  Ip *import1.IPv4Address `json:"ip"`
  
  PrefixLength *int `json:"prefixLength"`
}

func NewIPv4Subnet() *IPv4Subnet {
  p := new(IPv4Subnet)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.IPv4Subnet"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.IPv4Subnet"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
IP V6 configuration.
*/
type IPv6Config struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  DefaultGatewayIp *import1.IPv6Address `json:"defaultGatewayIp,omitempty"`
  
  DhcpServerAddress *import1.IPv6Address `json:"dhcpServerAddress,omitempty"`
  
  IpSubnet *IPv6Subnet `json:"ipSubnet"`
  /**
  Pool of IP addresses from where IPs are allocated.
  */
  PoolList []IPv6Pool `json:"poolList,omitempty"`
}

func NewIPv6Config() *IPv6Config {
  p := new(IPv6Config)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.IPv6Config"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.IPv6Config"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Start/end IP address range.
*/
type IPv6Pool struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  EndIP *import1.IPv6Address `json:"endIP"`
  
  StartIP *import1.IPv6Address `json:"startIP"`
}

func NewIPv6Pool() *IPv6Pool {
  p := new(IPv6Pool)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.IPv6Pool"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.IPv6Pool"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}





type IPv6Subnet struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  Ip *import1.IPv6Address `json:"ip"`
  
  PrefixLength *int `json:"prefixLength"`
}

func NewIPv6Subnet() *IPv6Subnet {
  p := new(IPv6Subnet)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.IPv6Subnet"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.IPv6Subnet"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Internal BGP configuration
*/
type IbgpConfig struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Autonomous system number
  */
  Asn *int `json:"asn,omitempty"`
  /**
  BPG password
  */
  Password *string `json:"password,omitempty"`
  
  PeerIp *import1.IPAddress `json:"peerIp,omitempty"`
  /**
  Redistribute routes over eBGP. Applicable only to network gateways deployed on VLAN subnets with eBGP over VPN.
  */
  RedistributeRoutes *bool `json:"redistributeRoutes,omitempty"`
}

func NewIbgpConfig() *IbgpConfig {
  p := new(IbgpConfig)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.IbgpConfig"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.IbgpConfig"}
  p.UnknownFields_ = map[string]interface{}{}

  p.RedistributeRoutes = new(bool)
  *p.RedistributeRoutes = false


  return p
}




/**
Describes the routing protocol configuration spec needed by this gateway to peer and learn routes from internal routers using either iBGP or OSFP
*/
type InternalRoutingConfig struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  iBGP configuration
  */
  IbgpConfigList []IbgpConfig `json:"ibgpConfigList,omitempty"`
  /**
  List of local prefixes to be advertised over eBGP
  */
  LocalPrefixList []IPSubnet `json:"localPrefixList,omitempty"`
  
  OspfConfig *OspfConfig `json:"ospfConfig,omitempty"`
}

func NewInternalRoutingConfig() *InternalRoutingConfig {
  p := new(InternalRoutingConfig)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.InternalRoutingConfig"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.InternalRoutingConfig"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Input to specify a count number of IP addresses to reserve on a subnet.
*/
type IpReserveCountInput struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  Count *int64 `json:"count,omitempty"`
}

func NewIpReserveCountInput() *IpReserveCountInput {
  p := new(IpReserveCountInput)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.IpReserveCountInput"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.IpReserveCountInput"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Input required to reserve IP addresses on a subnet.
*/
type IpReserveInput struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Optional context a client wishes to associate with a reservation of IP addresses.
  */
  ClientContext *string `json:"clientContext,omitempty"`
  /**
  
  */
  InputItemDiscriminator_ *string `json:"$inputItemDiscriminator,omitempty"`
  
  Input *OneOfIpReserveInputInput `json:"input,omitempty"`
}

func NewIpReserveInput() *IpReserveInput {
  p := new(IpReserveInput)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.IpReserveInput"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.IpReserveInput"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *IpReserveInput) GetInput() interface{} {
  if nil == p.Input {
    return nil
  }
  return p.Input.GetValue()
}

func (p *IpReserveInput) SetInput(v interface{}) error {
  if nil == p.Input {
    p.Input = NewOneOfIpReserveInputInput()
  }
  e := p.Input.SetValue(v)
  if nil == e {
    if nil == p.InputItemDiscriminator_ {
      p.InputItemDiscriminator_ = new(string)
    }
    *p.InputItemDiscriminator_ = *p.Input.Discriminator
  }
  return e
}



/**
Input to specify a list of IP addresses to reserve on a subnet.
*/
type IpReserveIpsInput struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  IpAddresses []import1.IPAddress `json:"ipAddresses,omitempty"`
}

func NewIpReserveIpsInput() *IpReserveIpsInput {
  p := new(IpReserveIpsInput)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.IpReserveIpsInput"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.IpReserveIpsInput"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Output that client can consume to see what IP addresses were reserved after the reservation is completed.
*/
type IpReserveOutput struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  IpAddresses []import1.IPAddress `json:"ipAddresses,omitempty"`
}

func NewIpReserveOutput() *IpReserveOutput {
  p := new(IpReserveOutput)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.IpReserveOutput"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.IpReserveOutput"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Input to specify a range of IP addresses to reserve on a subnet given a start IP and count range.
*/
type IpReserveRangeInput struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  Count *int64 `json:"count,omitempty"`
  
  StartIpAddress *import1.IPAddress `json:"startIpAddress,omitempty"`
}

func NewIpReserveRangeInput() *IpReserveRangeInput {
  p := new(IpReserveRangeInput)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.IpReserveRangeInput"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.IpReserveRangeInput"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Input where client can specify a context and all IP addresses associated with that context will be unreserved.
*/
type IpUnreserveClientContextInput struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  ClientContext *string `json:"clientContext,omitempty"`
}

func NewIpUnreserveClientContextInput() *IpUnreserveClientContextInput {
  p := new(IpUnreserveClientContextInput)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.IpUnreserveClientContextInput"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.IpUnreserveClientContextInput"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Input required to unreserve IP addresses on a subnet.
*/
type IpUnreserveInput struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  InputItemDiscriminator_ *string `json:"$inputItemDiscriminator,omitempty"`
  
  Input *OneOfIpUnreserveInputInput `json:"input,omitempty"`
}

func NewIpUnreserveInput() *IpUnreserveInput {
  p := new(IpUnreserveInput)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.IpUnreserveInput"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.IpUnreserveInput"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *IpUnreserveInput) GetInput() interface{} {
  if nil == p.Input {
    return nil
  }
  return p.Input.GetValue()
}

func (p *IpUnreserveInput) SetInput(v interface{}) error {
  if nil == p.Input {
    p.Input = NewOneOfIpUnreserveInputInput()
  }
  e := p.Input.SetValue(v)
  if nil == e {
    if nil == p.InputItemDiscriminator_ {
      p.InputItemDiscriminator_ = new(string)
    }
    *p.InputItemDiscriminator_ = *p.Input.Discriminator
  }
  return e
}



/**
Input to specify a list of IP addresses to unreserve on a subnet.
*/
type IpUnreserveIpsInput struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  IpAddresses []import1.IPAddress `json:"ipAddresses,omitempty"`
}

func NewIpUnreserveIpsInput() *IpUnreserveIpsInput {
  p := new(IpUnreserveIpsInput)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.IpUnreserveIpsInput"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.IpUnreserveIpsInput"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Input to specify a range of IP addresses to unreserve on a subnet given a start IP and count range.
*/
type IpUnreserveRangeInput struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  Count *int64 `json:"count,omitempty"`
  
  StartIpAddress *import1.IPAddress `json:"startIpAddress,omitempty"`
}

func NewIpUnreserveRangeInput() *IpUnreserveRangeInput {
  p := new(IpUnreserveRangeInput)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.IpUnreserveRangeInput"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.IpUnreserveRangeInput"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
IPSec configuration
*/
type IpsecConfig struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Diffie-Hellman group value of 14, 19 or 20 to be used for Perfect Forward Secrecy (PFS)
  */
  EspPfsDhGroupNumber *int `json:"espPfsDhGroupNumber,omitempty"`
  
  IkeAuthenticationAlgorithm *AuthenticationAlgorithm `json:"ikeAuthenticationAlgorithm,omitempty"`
  
  IkeEncryptionAlgorithm *EncryptionAlgorithm `json:"ikeEncryptionAlgorithm,omitempty"`
  /**
  IKE lifetime (seconds)
  */
  IkeLifetimeSecs *int64 `json:"ikeLifetimeSecs,omitempty"`
  
  IpsecAuthenticationAlgorithm *AuthenticationAlgorithm `json:"ipsecAuthenticationAlgorithm,omitempty"`
  
  IpsecEncryptionAlgorithm *EncryptionAlgorithm `json:"ipsecEncryptionAlgorithm,omitempty"`
  /**
  IPSec lifetime (seconds)
  */
  IpsecLifetimeSecs *int64 `json:"ipsecLifetimeSecs,omitempty"`
  /**
  Local IKE authentication ID used for this connection
  */
  LocalAuthenticationId *string `json:"localAuthenticationId,omitempty"`
  
  LocalVtiIp *import1.IPAddress `json:"localVtiIp,omitempty"`
  /**
  Shared secret for authentication between gateway peers
  */
  PreSharedKey *string `json:"preSharedKey"`
  /**
  IKE Authentication ID of the remote peer
  */
  RemoteAuthenticationId *string `json:"remoteAuthenticationId,omitempty"`
  
  RemoteVtiIp *import1.IPAddress `json:"remoteVtiIp,omitempty"`
}

func NewIpsecConfig() *IpsecConfig {
  p := new(IpsecConfig)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.IpsecConfig"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.IpsecConfig"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}





type Layer2Stretch struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  ConnectionType *StretchConnectionType `json:"connectionType,omitempty"`
  /**
  Layer2 stretch configuration details between subnets on two sites.
  */
  Description *string `json:"description,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import2.ApiLink `json:"links,omitempty"`
  
  LocalSiteParams *SiteParams `json:"localSiteParams,omitempty"`
  
  Metadata *import1.Metadata `json:"metadata,omitempty"`
  /**
  The MTU size setting for the VXLAN session.
  */
  Mtu *int `json:"mtu,omitempty"`
  /**
  Layer2 stretch configuration name.
  */
  Name *string `json:"name"`
  
  RemoteSiteParams *SiteParams `json:"remoteSiteParams,omitempty"`
  
  RemoteStretchStatus []RemoteVtepStretchStatus `json:"remoteStretchStatus,omitempty"`
  
  StretchStatus *StretchStatus `json:"stretchStatus,omitempty"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
  /**
  The VXLAN network identifier used to uniquely identify the VXLAN tunnel.
  */
  Vni *int `json:"vni,omitempty"`
}

func NewLayer2Stretch() *Layer2Stretch {
  p := new(Layer2Stretch)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.Layer2Stretch"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.Layer2Stretch"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST Response for all response codes in api path /networking/v4.0.a1/config/layer2-stretches/{extId} Get operation
*/
type Layer2StretchApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfLayer2StretchApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewLayer2StretchApiResponse() *Layer2StretchApiResponse {
  p := new(Layer2StretchApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.Layer2StretchApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.Layer2StretchApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *Layer2StretchApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *Layer2StretchApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfLayer2StretchApiResponseData()
  }
  e := p.Data.SetValue(v)
  if nil == e {
    if nil == p.DataItemDiscriminator_ {
      p.DataItemDiscriminator_ = new(string)
    }
    *p.DataItemDiscriminator_ = *p.Data.Discriminator
  }
  return e
}



/**
REST Response for all response codes in api path /networking/v4.0.a1/config/layer2-stretches Get operation
*/
type Layer2StretchListApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfLayer2StretchListApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewLayer2StretchListApiResponse() *Layer2StretchListApiResponse {
  p := new(Layer2StretchListApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.Layer2StretchListApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.Layer2StretchListApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *Layer2StretchListApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *Layer2StretchListApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfLayer2StretchListApiResponseData()
  }
  e := p.Data.SetValue(v)
  if nil == e {
    if nil == p.DataItemDiscriminator_ {
      p.DataItemDiscriminator_ = new(string)
    }
    *p.DataItemDiscriminator_ = *p.Data.Discriminator
  }
  return e
}



/**
Layer2 stretch-related entities retrieved from the specified PC cluster.
*/
type Layer2StretchRelatedEntities struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  Subnets []Layer2StretchSubnetInfo `json:"subnets,omitempty"`
  
  VpnConnections []Layer2StretchVpnConnectionInfo `json:"vpnConnections,omitempty"`
  
  VtepGateways []Layer2StretchVtepGatewayInfo `json:"vtepGateways,omitempty"`
}

func NewLayer2StretchRelatedEntities() *Layer2StretchRelatedEntities {
  p := new(Layer2StretchRelatedEntities)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.Layer2StretchRelatedEntities"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.Layer2StretchRelatedEntities"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST Response for all response codes in api path /networking/v4.0.a1/config/layer2-stretches/related-entities/{pcClusterExtId} Get operation
*/
type Layer2StretchRelatedEntitiesApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfLayer2StretchRelatedEntitiesApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewLayer2StretchRelatedEntitiesApiResponse() *Layer2StretchRelatedEntitiesApiResponse {
  p := new(Layer2StretchRelatedEntitiesApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.Layer2StretchRelatedEntitiesApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.Layer2StretchRelatedEntitiesApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *Layer2StretchRelatedEntitiesApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *Layer2StretchRelatedEntitiesApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfLayer2StretchRelatedEntitiesApiResponseData()
  }
  e := p.Data.SetValue(v)
  if nil == e {
    if nil == p.DataItemDiscriminator_ {
      p.DataItemDiscriminator_ = new(string)
    }
    *p.DataItemDiscriminator_ = *p.Data.Discriminator
  }
  return e
}



/**
Information about a subnet from the specified PC cluster.
*/
type Layer2StretchSubnetInfo struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  ClusterReference *import1.EntityReference `json:"clusterReference,omitempty"`
  
  DefaultGatewayIp *import1.IPAddress `json:"defaultGatewayIp,omitempty"`
  
  IpSubnet *IPSubnet `json:"ipSubnet,omitempty"`
  
  SubnetReference *import1.EntityReference `json:"subnetReference,omitempty"`
  /**
  VLAN ID (if this subnet is vlan-backed).
  */
  VlanId *int `json:"vlanId,omitempty"`
  
  VpcReference *import1.EntityReference `json:"vpcReference,omitempty"`
}

func NewLayer2StretchSubnetInfo() *Layer2StretchSubnetInfo {
  p := new(Layer2StretchSubnetInfo)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.Layer2StretchSubnetInfo"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.Layer2StretchSubnetInfo"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Information about a VPN connection from the specified PC cluster.
*/
type Layer2StretchVpnConnectionInfo struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  ClusterReference *import1.EntityReference `json:"clusterReference,omitempty"`
  
  ConnectionReference *import1.EntityReference `json:"connectionReference,omitempty"`
  
  LocalVtiIPAddress *import1.IPAddress `json:"localVtiIPAddress,omitempty"`
  
  PeerConnectionReference *import1.EntityReference `json:"peerConnectionReference,omitempty"`
  
  VpcReference *import1.EntityReference `json:"vpcReference,omitempty"`
}

func NewLayer2StretchVpnConnectionInfo() *Layer2StretchVpnConnectionInfo {
  p := new(Layer2StretchVpnConnectionInfo)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.Layer2StretchVpnConnectionInfo"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.Layer2StretchVpnConnectionInfo"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Information about a VTEP gateway.
*/
type Layer2StretchVtepGatewayInfo struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  ClusterReference *import1.EntityReference `json:"clusterReference,omitempty"`
  /**
  If true, VTEP gateway is local. If false, VTEP gateway is remote.
  */
  IsLocal *bool `json:"isLocal,omitempty"`
  /**
  VTEP gateway name.
  */
  NetworkGatewayName *string `json:"networkGatewayName,omitempty"`
  /**
  VTEP gateway ID.
  */
  NetworkGatewayReference *string `json:"networkGatewayReference,omitempty"`
  
  VpcReference *import1.EntityReference `json:"vpcReference,omitempty"`
  
  Vteps []Vtep `json:"vteps,omitempty"`
  /**
  VXLAN port
  */
  VxlanPort *int `json:"vxlanPort,omitempty"`
}

func NewLayer2StretchVtepGatewayInfo() *Layer2StretchVtepGatewayInfo {
  p := new(Layer2StretchVtepGatewayInfo)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.Layer2StretchVtepGatewayInfo"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.Layer2StretchVtepGatewayInfo"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Information pertaining to a learned IP address on a subnet.
*/
type LearnedAddressInfo struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  LastSeen *string `json:"lastSeen,omitempty"`
  
  MacAddress *MacAddress `json:"macAddress,omitempty"`
  
  VmReference *import1.EntityReference `json:"vmReference,omitempty"`
}

func NewLearnedAddressInfo() *LearnedAddressInfo {
  p := new(LearnedAddressInfo)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.LearnedAddressInfo"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.LearnedAddressInfo"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
BGP service hosted on this local gateway.
*/
type LocalBgpService struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Autonomous system number
  */
  Asn *int `json:"asn"`
  /**
  If true, routes are exchanged between BGP and VPN services.
  */
  ExchangeRoutesWithVpn *bool `json:"exchangeRoutesWithVpn,omitempty"`
  /**
  BGP graceful restart timeout seconds.
  */
  GracefulRestartTimeoutSeconds *int `json:"gracefulRestartTimeoutSeconds,omitempty"`
  /**
  Reference to the VPC that this network gateway serves as its BGP speaker.
  */
  VpcReference *string `json:"vpcReference,omitempty"`
}

func NewLocalBgpService() *LocalBgpService {
  p := new(LocalBgpService)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.LocalBgpService"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.LocalBgpService"}
  p.UnknownFields_ = map[string]interface{}{}

  p.ExchangeRoutesWithVpn = new(bool)
  *p.ExchangeRoutesWithVpn = false


  return p
}




/**
Services of this local gateway
*/
type LocalNetworkServices struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  LocalBgpService *LocalBgpService `json:"localBgpService,omitempty"`
  
  LocalVpnService *LocalVpnService `json:"localVpnService,omitempty"`
  
  LocalVtepService *LocalVtepService `json:"localVtepService,omitempty"`
  
  PublicIpAddress *import1.IPAddress `json:"publicIpAddress,omitempty"`
}

func NewLocalNetworkServices() *LocalNetworkServices {
  p := new(LocalNetworkServices)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.LocalNetworkServices"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.LocalNetworkServices"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
VPN service hosted on this local gateway
*/
type LocalVpnService struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  EbgpConfig *BgpConfig `json:"ebgpConfig,omitempty"`
  
  PeerIgpConfig *InternalRoutingConfig `json:"peerIgpConfig,omitempty"`
}

func NewLocalVpnService() *LocalVpnService {
  p := new(LocalVpnService)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.LocalVpnService"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.LocalVpnService"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
VTEP service hosted on this local gateway
*/
type LocalVtepService struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  VXLAN port
  */
  VxlanPort *int `json:"vxlanPort"`
}

func NewLocalVtepService() *LocalVtepService {
  p := new(LocalVtepService)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.LocalVtepService"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.LocalVtepService"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
MAC Address
*/
type MacAddress struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  Address *string `json:"address"`
}

func NewMacAddress() *MacAddress {
  p := new(MacAddress)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.MacAddress"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.MacAddress"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}





type NetworkCloudConfig struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  AzureConfig *AzureConfig `json:"azureConfig"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import2.ApiLink `json:"links,omitempty"`
  
  Metadata *import1.Metadata `json:"metadata,omitempty"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
}

func NewNetworkCloudConfig() *NetworkCloudConfig {
  p := new(NetworkCloudConfig)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.NetworkCloudConfig"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.NetworkCloudConfig"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST Response for  NetworkCloudConfig type
*/
type NetworkCloudConfigApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfNetworkCloudConfigApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewNetworkCloudConfigApiResponse() *NetworkCloudConfigApiResponse {
  p := new(NetworkCloudConfigApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.NetworkCloudConfigApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.NetworkCloudConfigApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *NetworkCloudConfigApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *NetworkCloudConfigApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfNetworkCloudConfigApiResponseData()
  }
  e := p.Data.SetValue(v)
  if nil == e {
    if nil == p.DataItemDiscriminator_ {
      p.DataItemDiscriminator_ = new(string)
    }
    *p.DataItemDiscriminator_ = *p.Data.Discriminator
  }
  return e
}



/**
REST Response for  list of  NetworkCloudConfig types
*/
type NetworkCloudConfigListApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfNetworkCloudConfigListApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewNetworkCloudConfigListApiResponse() *NetworkCloudConfigListApiResponse {
  p := new(NetworkCloudConfigListApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.NetworkCloudConfigListApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.NetworkCloudConfigListApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *NetworkCloudConfigListApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *NetworkCloudConfigListApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfNetworkCloudConfigListApiResponseData()
  }
  e := p.Data.SetValue(v)
  if nil == e {
    if nil == p.DataItemDiscriminator_ {
      p.DataItemDiscriminator_ = new(string)
    }
    *p.DataItemDiscriminator_ = *p.Data.Discriminator
  }
  return e
}




type NetworkController struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  CloudSubstrate *CloudSubstrate `json:"cloudSubstrate,omitempty"`
  
  ControllerStatus *ControllerStatus `json:"controllerStatus,omitempty"`
  /**
  Network controller software version.
  */
  ControllerVersion *string `json:"controllerVersion,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import2.ApiLink `json:"links,omitempty"`
  
  Metadata *import1.Metadata `json:"metadata,omitempty"`
  /**
  Minimum AHV version that this network controller supports.
  */
  MinimumAHVVersion *string `json:"minimumAHVVersion,omitempty"`
  /**
  Minimum NOS version that this network controller supports.
  */
  MinimumNOSVersion *string `json:"minimumNOSVersion,omitempty"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
}

func NewNetworkController() *NetworkController {
  p := new(NetworkController)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.NetworkController"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.NetworkController"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST Response for all response codes in api path /networking/v4.0.a1/config/controllers/{extId} Get operation
*/
type NetworkControllerApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfNetworkControllerApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewNetworkControllerApiResponse() *NetworkControllerApiResponse {
  p := new(NetworkControllerApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.NetworkControllerApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.NetworkControllerApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *NetworkControllerApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *NetworkControllerApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfNetworkControllerApiResponseData()
  }
  e := p.Data.SetValue(v)
  if nil == e {
    if nil == p.DataItemDiscriminator_ {
      p.DataItemDiscriminator_ = new(string)
    }
    *p.DataItemDiscriminator_ = *p.Data.Discriminator
  }
  return e
}



/**
REST Response for all response codes in api path /networking/v4.0.a1/config/controllers Get operation
*/
type NetworkControllerListApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfNetworkControllerListApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewNetworkControllerListApiResponse() *NetworkControllerListApiResponse {
  p := new(NetworkControllerListApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.NetworkControllerListApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.NetworkControllerListApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *NetworkControllerListApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *NetworkControllerListApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfNetworkControllerListApiResponseData()
  }
  e := p.Data.SetValue(v)
  if nil == e {
    if nil == p.DataItemDiscriminator_ {
      p.DataItemDiscriminator_ = new(string)
    }
    *p.DataItemDiscriminator_ = *p.Data.Discriminator
  }
  return e
}



/**
Network gateway
*/
type NetworkGateway struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  Deployment *NetworkGatewayDeployment `json:"deployment,omitempty"`
  /**
  Description of the gateway
  */
  Description *string `json:"description,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  Third-party gateway vendor
  */
  GatewayDeviceVendor *string `json:"gatewayDeviceVendor,omitempty"`
  /**
  Software version installed on the gateway appliance
  */
  InstalledSoftwareVersion *string `json:"installedSoftwareVersion,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import2.ApiLink `json:"links,omitempty"`
  
  Metadata *import1.Metadata `json:"metadata,omitempty"`
  /**
  Name of the gateway
  */
  Name *string `json:"name,omitempty"`
  /**
  
  */
  ServicesItemDiscriminator_ *string `json:"$servicesItemDiscriminator,omitempty"`
  
  Services *OneOfNetworkGatewayServices `json:"services,omitempty"`
  
  Status *Status `json:"status,omitempty"`
  /**
  Supported gateway appliance version
  */
  SupportedSoftwareVersion *string `json:"supportedSoftwareVersion,omitempty"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
  /**
  Reference to a dedicated VM on which a local gateway is deployed
  */
  VmReference *string `json:"vmReference,omitempty"`
  /**
  VPC
  */
  VpcReference *string `json:"vpcReference,omitempty"`
}

func NewNetworkGateway() *NetworkGateway {
  p := new(NetworkGateway)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.NetworkGateway"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.NetworkGateway"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST Response for all response codes in api path /networking/v4.0.a1/config/network-gateways/{extId} Get operation
*/
type NetworkGatewayApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfNetworkGatewayApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewNetworkGatewayApiResponse() *NetworkGatewayApiResponse {
  p := new(NetworkGatewayApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.NetworkGatewayApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.NetworkGatewayApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *NetworkGatewayApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *NetworkGatewayApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfNetworkGatewayApiResponseData()
  }
  e := p.Data.SetValue(v)
  if nil == e {
    if nil == p.DataItemDiscriminator_ {
      p.DataItemDiscriminator_ = new(string)
    }
    *p.DataItemDiscriminator_ = *p.Data.Discriminator
  }
  return e
}



/**
Network gateway deployment configuration
*/
type NetworkGatewayDeployment struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Cluster reference required to identify which on-prem cluster to deploy the gateway VM on
  */
  ClusterReference *string `json:"clusterReference,omitempty"`
  
  ManagementInterface *NetworkGatewayManagementInterface `json:"managementInterface,omitempty"`
  /**
  vCenter datastore to which the gateway disks and images will be uploaded during deployment
  */
  VcenterDatastoreName *string `json:"vcenterDatastoreName,omitempty"`
}

func NewNetworkGatewayDeployment() *NetworkGatewayDeployment {
  p := new(NetworkGatewayDeployment)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.NetworkGatewayDeployment"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.NetworkGatewayDeployment"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST Response for all response codes in api path /networking/v4.0.a1/config/network-gateways Get operation
*/
type NetworkGatewayListApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfNetworkGatewayListApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewNetworkGatewayListApiResponse() *NetworkGatewayListApiResponse {
  p := new(NetworkGatewayListApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.NetworkGatewayListApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.NetworkGatewayListApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *NetworkGatewayListApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *NetworkGatewayListApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfNetworkGatewayListApiResponseData()
  }
  e := p.Data.SetValue(v)
  if nil == e {
    if nil == p.DataItemDiscriminator_ {
      p.DataItemDiscriminator_ = new(string)
    }
    *p.DataItemDiscriminator_ = *p.Data.Discriminator
  }
  return e
}



/**
Network interface used to deliver network services and for managing the gateway. If a VPC reference is supplied then the gateway will be deployed on an dedicated subnet within the VPC. If a VPC reference is not supplied, then this interface defines the subnet on which the gateway will be deployed, and the address it will be assigned. When a VPC reference is not present, either a subnet reference or a VLAN id must be supplied, along with the address and default gateway of the subnet. A VLAN network without IPAM may be used.
*/
type NetworkGatewayManagementInterface struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  Address *import1.IPAddress `json:"address,omitempty"`
  
  DefaultGateway *import1.IPAddress `json:"defaultGateway,omitempty"`
  /**
  MTU of management interface
  */
  Mtu *int `json:"mtu,omitempty"`
  /**
  The on-prem vlan subnet to deploy the network gateway VM on
  */
  SubnetReference *string `json:"subnetReference,omitempty"`
  /**
  The on-prem VLAN to deploy the gateway on
  */
  VlanId *int `json:"vlanId,omitempty"`
}

func NewNetworkGatewayManagementInterface() *NetworkGatewayManagementInterface {
  p := new(NetworkGatewayManagementInterface)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.NetworkGatewayManagementInterface"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.NetworkGatewayManagementInterface"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Networking resource base model with metadata support
*/
type NetworkingBaseModel struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import2.ApiLink `json:"links,omitempty"`
  
  Metadata *import1.Metadata `json:"metadata,omitempty"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
}

func NewNetworkingBaseModel() *NetworkingBaseModel {
  p := new(NetworkingBaseModel)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.NetworkingBaseModel"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.NetworkingBaseModel"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Next hop type.
*/
type NexthopType int

const(
  NEXTHOPTYPE_UNKNOWN NexthopType = 0
  NEXTHOPTYPE_REDACTED NexthopType = 1
  NEXTHOPTYPE_IP_ADDRESS NexthopType = 2
  NEXTHOPTYPE_DIRECT_CONNECT_VIF NexthopType = 3
  NEXTHOPTYPE_INTERNAL_SUBNET NexthopType = 4
  NEXTHOPTYPE_EXTERNAL_SUBNET NexthopType = 5
  NEXTHOPTYPE_VPN_CONNECTION NexthopType = 6
)

// returns the name of the enum given an ordinal number
func (e *NexthopType) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "IP_ADDRESS",
    "DIRECT_CONNECT_VIF",
    "INTERNAL_SUBNET",
    "EXTERNAL_SUBNET",
    "VPN_CONNECTION",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *NexthopType) index(name string) NexthopType {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "IP_ADDRESS",
    "DIRECT_CONNECT_VIF",
    "INTERNAL_SUBNET",
    "EXTERNAL_SUBNET",
    "VPN_CONNECTION",
  }
  for idx := range names {
    if names[idx] == name {
      return NexthopType(idx)
    }
  }
  return NEXTHOPTYPE_UNKNOWN
}

func (e *NexthopType) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for NexthopType:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *NexthopType) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e NexthopType) Ref() *NexthopType {
  return &e
}


/**
OSPF configuration for route peering with internal routers
*/
type OspfConfig struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  OSPF area id of this gateway
  */
  AreaId *string `json:"areaId,omitempty"`
  
  AuthenticationType *AuthenticationType `json:"authenticationType,omitempty"`
  /**
  Password for authentication
  */
  Password *string `json:"password,omitempty"`
}

func NewOspfConfig() *OspfConfig {
  p := new(OspfConfig)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.OspfConfig"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.OspfConfig"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Range of TCP/UDP ports.
*/
type PortRange struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  EndPort *int `json:"endPort"`
  
  StartPort *int `json:"startPort"`
}

func NewPortRange() *PortRange {
  p := new(PortRange)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.PortRange"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.PortRange"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Private IP and VPC to which the floating IP is associated with.
*/
type PrivateIpAssociation struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  PrivateIp *import1.IPAddress `json:"privateIp"`
  /**
  VPC in which the Private IP exists.
  */
  VpcReference *string `json:"vpcReference"`
}

func NewPrivateIpAssociation() *PrivateIpAssociation {
  p := new(PrivateIpAssociation)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.PrivateIpAssociation"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.PrivateIpAssociation"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}





type ProtocolNumberObject struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  ProtocolNumber *int `json:"protocolNumber"`
}

func NewProtocolNumberObject() *ProtocolNumberObject {
  p := new(ProtocolNumberObject)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.ProtocolNumberObject"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.ProtocolNumberObject"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Routing Policy IP protocol type.
*/
type ProtocolType int

const(
  PROTOCOLTYPE_UNKNOWN ProtocolType = 0
  PROTOCOLTYPE_REDACTED ProtocolType = 1
  PROTOCOLTYPE_ANY ProtocolType = 2
  PROTOCOLTYPE_TCP ProtocolType = 3
  PROTOCOLTYPE_UDP ProtocolType = 4
  PROTOCOLTYPE_ICMP ProtocolType = 5
  PROTOCOLTYPE_PROTOCOL_NUMBER ProtocolType = 6
)

// returns the name of the enum given an ordinal number
func (e *ProtocolType) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "ANY",
    "TCP",
    "UDP",
    "ICMP",
    "PROTOCOL_NUMBER",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *ProtocolType) index(name string) ProtocolType {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "ANY",
    "TCP",
    "UDP",
    "ICMP",
    "PROTOCOL_NUMBER",
  }
  for idx := range names {
    if names[idx] == name {
      return ProtocolType(idx)
    }
  }
  return PROTOCOLTYPE_UNKNOWN
}

func (e *ProtocolType) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for ProtocolType:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *ProtocolType) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e ProtocolType) Ref() *ProtocolType {
  return &e
}


/**
Quality of Service configuration for the VPN IPSec tunnel
*/
type QosConfig struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Egress traffic limit (Mbps)
  */
  EgressLimitMbps *int64 `json:"egressLimitMbps,omitempty"`
  /**
  Ingress traffic limit (Mbps)
  */
  IngressLimitMbps *int64 `json:"ingressLimitMbps,omitempty"`
}

func NewQosConfig() *QosConfig {
  p := new(QosConfig)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.QosConfig"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.QosConfig"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
BGP service hosted on this remote gateway.
*/
type RemoteBgpService struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  Address *import1.IPAddress `json:"address"`
  /**
  Autonomous system number
  */
  Asn *int `json:"asn"`
}

func NewRemoteBgpService() *RemoteBgpService {
  p := new(RemoteBgpService)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.RemoteBgpService"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.RemoteBgpService"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Services of this remote gateway
*/
type RemoteNetworkServices struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  RemoteBgpService *RemoteBgpService `json:"remoteBgpService,omitempty"`
  
  RemoteVpnService *RemoteVpnService `json:"remoteVpnService,omitempty"`
  
  RemoteVtepService *RemoteVtepService `json:"remoteVtepService,omitempty"`
}

func NewRemoteNetworkServices() *RemoteNetworkServices {
  p := new(RemoteNetworkServices)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.RemoteNetworkServices"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.RemoteNetworkServices"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
VPN service hosted on this remote gateway
*/
type RemoteVpnService struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  EbgpConfig *BgpConfig `json:"ebgpConfig,omitempty"`
  /**
  Boolean flag indicating user opt-in for installing Xi LB route in on-prem PC and PE CVMs provided on-prem PC, PE and VPN VM are in the same subnet.
  */
  InstallXiRoute *bool `json:"installXiRoute,omitempty"`
  
  PeerIgpConfig *InternalRoutingConfig `json:"peerIgpConfig,omitempty"`
  
  PublicIpAddress *import1.IPAddress `json:"publicIpAddress,omitempty"`
}

func NewRemoteVpnService() *RemoteVpnService {
  p := new(RemoteVpnService)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.RemoteVpnService"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.RemoteVpnService"}
  p.UnknownFields_ = map[string]interface{}{}

  p.InstallXiRoute = new(bool)
  *p.InstallXiRoute = false


  return p
}




/**
VTEP service hosted on this remote gateway
*/
type RemoteVtepService struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Remote VXLAN Tunnel Endpoints configuration
  */
  Vteps []Vtep `json:"vteps,omitempty"`
  /**
  VXLAN port
  */
  VxlanPort *int `json:"vxlanPort,omitempty"`
}

func NewRemoteVtepService() *RemoteVtepService {
  p := new(RemoteVtepService)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.RemoteVtepService"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.RemoteVtepService"}
  p.UnknownFields_ = map[string]interface{}{}

  p.VxlanPort = new(int)
  *p.VxlanPort = 4789


  return p
}




/**
Status of each VTEP. Applicable only when connectionType is VXLAN.
*/
type RemoteVtepStretchStatus struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  Address *import1.IPAddress `json:"address,omitempty"`
  
  LearnedMacAddresses []MacAddress `json:"learnedMacAddresses,omitempty"`
  
  Status *StretchStatus `json:"status,omitempty"`
}

func NewRemoteVtepStretchStatus() *RemoteVtepStretchStatus {
  p := new(RemoteVtepStretchStatus)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.RemoteVtepStretchStatus"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.RemoteVtepStretchStatus"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}





type RerouteFallbackAction int

const(
  REROUTEFALLBACKACTION_UNKNOWN RerouteFallbackAction = 0
  REROUTEFALLBACKACTION_REDACTED RerouteFallbackAction = 1
  REROUTEFALLBACKACTION_PERMIT RerouteFallbackAction = 2
  REROUTEFALLBACKACTION_DENY RerouteFallbackAction = 3
  REROUTEFALLBACKACTION_PASSTHROUGH RerouteFallbackAction = 4
)

// returns the name of the enum given an ordinal number
func (e *RerouteFallbackAction) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "PERMIT",
    "DENY",
    "PASSTHROUGH",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *RerouteFallbackAction) index(name string) RerouteFallbackAction {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "PERMIT",
    "DENY",
    "PASSTHROUGH",
  }
  for idx := range names {
    if names[idx] == name {
      return RerouteFallbackAction(idx)
    }
  }
  return REROUTEFALLBACKACTION_UNKNOWN
}

func (e *RerouteFallbackAction) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for RerouteFallbackAction:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *RerouteFallbackAction) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e RerouteFallbackAction) Ref() *RerouteFallbackAction {
  return &e
}


/**
Parameters for the reroute action which includes the reroute service IP and the fallback action when the service IP is down.
*/
type RerouteParam struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  RerouteFallbackAction *RerouteFallbackAction `json:"rerouteFallbackAction,omitempty"`
  
  ServiceIp *import1.IPv4Address `json:"serviceIp"`
}

func NewRerouteParam() *RerouteParam {
  p := new(RerouteParam)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.RerouteParam"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.RerouteParam"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Information pertaining to a reserved IP address on a subnet.
*/
type ReservedAddressInfo struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  ClientContext *string `json:"clientContext,omitempty"`
}

func NewReservedAddressInfo() *ReservedAddressInfo {
  p := new(ReservedAddressInfo)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.ReservedAddressInfo"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.ReservedAddressInfo"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Route.
*/
type Route struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Route active or inactive.
  */
  Active *bool `json:"active,omitempty"`
  
  Destination *IPSubnet `json:"destination"`
  
  NexthopIpAddress *import1.IPAddress `json:"nexthopIpAddress,omitempty"`
  /**
  Reference to a link, such as a VPN connection or a subnet.
  */
  NexthopReference *string `json:"nexthopReference,omitempty"`
  
  NexthopType *NexthopType `json:"nexthopType"`
  /**
  Higher value implies greater preference is assigned to route.
  */
  Priority *int `json:"priority,omitempty"`
}

func NewRoute() *Route {
  p := new(Route)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.Route"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.Route"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Route table
*/
type RouteTable struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Dynamic routes
  */
  DynamicRoutes []Route `json:"dynamicRoutes,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import2.ApiLink `json:"links,omitempty"`
  /**
  Routes to local subnets
  */
  LocalRoutes []Route `json:"localRoutes,omitempty"`
  
  Metadata *import1.Metadata `json:"metadata,omitempty"`
  /**
  Static routes
  */
  StaticRoutes []Route `json:"staticRoutes,omitempty"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
  /**
  VPC
  */
  VpcReference *string `json:"vpcReference"`
}

func NewRouteTable() *RouteTable {
  p := new(RouteTable)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.RouteTable"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.RouteTable"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST Response for all response codes in api path /networking/v4.0.a1/config/route-tables/{extId} Get operation
*/
type RouteTableApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfRouteTableApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewRouteTableApiResponse() *RouteTableApiResponse {
  p := new(RouteTableApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.RouteTableApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.RouteTableApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *RouteTableApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *RouteTableApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfRouteTableApiResponseData()
  }
  e := p.Data.SetValue(v)
  if nil == e {
    if nil == p.DataItemDiscriminator_ {
      p.DataItemDiscriminator_ = new(string)
    }
    *p.DataItemDiscriminator_ = *p.Data.Discriminator
  }
  return e
}



/**
REST Response for all response codes in api path /networking/v4.0.a1/config/route-tables Get operation
*/
type RouteTableListApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfRouteTableListApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewRouteTableListApiResponse() *RouteTableListApiResponse {
  p := new(RouteTableListApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.RouteTableListApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.RouteTableListApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *RouteTableListApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *RouteTableListApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfRouteTableListApiResponseData()
  }
  e := p.Data.SetValue(v)
  if nil == e {
    if nil == p.DataItemDiscriminator_ {
      p.DataItemDiscriminator_ = new(string)
    }
    *p.DataItemDiscriminator_ = *p.Data.Discriminator
  }
  return e
}



/**
Schema to configure a Routing Policy.
*/
type RoutingPolicy struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  A description of the Routing Policy.
  */
  Description *string `json:"description,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import2.ApiLink `json:"links,omitempty"`
  
  Metadata *import1.Metadata `json:"metadata,omitempty"`
  /**
  Name of the Routing Policy.
  */
  Name *string `json:"name"`
  
  Policies []RoutingPolicyRule `json:"policies"`
  /**
  Priority of the Routing Policy.
  */
  Priority *int `json:"priority"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
  /**
  ExtId of the VPC to which the Routing Policy belongs to.
  */
  VpcExtId *string `json:"vpcExtId"`
}

func NewRoutingPolicy() *RoutingPolicy {
  p := new(RoutingPolicy)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.RoutingPolicy"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.RoutingPolicy"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
The action to be taken on the traffic matching the Routing Policy.
*/
type RoutingPolicyAction struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  ActionType *RoutingPolicyActionType `json:"actionType"`
  
  RerouteParams []RerouteParam `json:"rerouteParams,omitempty"`
}

func NewRoutingPolicyAction() *RoutingPolicyAction {
  p := new(RoutingPolicyAction)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.RoutingPolicyAction"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.RoutingPolicyAction"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}





type RoutingPolicyActionType int

const(
  ROUTINGPOLICYACTIONTYPE_UNKNOWN RoutingPolicyActionType = 0
  ROUTINGPOLICYACTIONTYPE_REDACTED RoutingPolicyActionType = 1
  ROUTINGPOLICYACTIONTYPE_PERMIT RoutingPolicyActionType = 2
  ROUTINGPOLICYACTIONTYPE_DENY RoutingPolicyActionType = 3
  ROUTINGPOLICYACTIONTYPE_REROUTE RoutingPolicyActionType = 4
)

// returns the name of the enum given an ordinal number
func (e *RoutingPolicyActionType) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "PERMIT",
    "DENY",
    "REROUTE",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *RoutingPolicyActionType) index(name string) RoutingPolicyActionType {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "PERMIT",
    "DENY",
    "REROUTE",
  }
  for idx := range names {
    if names[idx] == name {
      return RoutingPolicyActionType(idx)
    }
  }
  return ROUTINGPOLICYACTIONTYPE_UNKNOWN
}

func (e *RoutingPolicyActionType) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for RoutingPolicyActionType:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *RoutingPolicyActionType) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e RoutingPolicyActionType) Ref() *RoutingPolicyActionType {
  return &e
}


/**
REST Response for all response codes in api path /networking/v4.0.a1/config/routing-policies/{extId} Get operation
*/
type RoutingPolicyApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfRoutingPolicyApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewRoutingPolicyApiResponse() *RoutingPolicyApiResponse {
  p := new(RoutingPolicyApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.RoutingPolicyApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.RoutingPolicyApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *RoutingPolicyApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *RoutingPolicyApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfRoutingPolicyApiResponseData()
  }
  e := p.Data.SetValue(v)
  if nil == e {
    if nil == p.DataItemDiscriminator_ {
      p.DataItemDiscriminator_ = new(string)
    }
    *p.DataItemDiscriminator_ = *p.Data.Discriminator
  }
  return e
}



/**
REST Response for all response codes in api path /networking/v4.0.a1/config/routing-policies Get operation
*/
type RoutingPolicyListApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfRoutingPolicyListApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewRoutingPolicyListApiResponse() *RoutingPolicyListApiResponse {
  p := new(RoutingPolicyListApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.RoutingPolicyListApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.RoutingPolicyListApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *RoutingPolicyListApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *RoutingPolicyListApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfRoutingPolicyListApiResponseData()
  }
  e := p.Data.SetValue(v)
  if nil == e {
    if nil == p.DataItemDiscriminator_ {
      p.DataItemDiscriminator_ = new(string)
    }
    *p.DataItemDiscriminator_ = *p.Data.Discriminator
  }
  return e
}



/**
Match condition for the traffic that is entering the VPC.
*/
type RoutingPolicyMatchCondition struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DestinationItemDiscriminator_ *string `json:"$destinationItemDiscriminator,omitempty"`
  /**
  Destination of the traffic this is entering the VPC.
  */
  Destination *OneOfRoutingPolicyMatchConditionDestination `json:"destination"`
  /**
  
  */
  ProtocolParametersItemDiscriminator_ *string `json:"$protocolParametersItemDiscriminator,omitempty"`
  
  ProtocolParameters *OneOfRoutingPolicyMatchConditionProtocolParameters `json:"protocolParameters,omitempty"`
  
  ProtocolType *ProtocolType `json:"protocolType"`
  /**
  
  */
  SourceItemDiscriminator_ *string `json:"$sourceItemDiscriminator,omitempty"`
  /**
  Source of the traffic that is entering the VPC.
  */
  Source *OneOfRoutingPolicyMatchConditionSource `json:"source"`
}

func NewRoutingPolicyMatchCondition() *RoutingPolicyMatchCondition {
  p := new(RoutingPolicyMatchCondition)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.RoutingPolicyMatchCondition"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.RoutingPolicyMatchCondition"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *RoutingPolicyMatchCondition) GetProtocolParameters() interface{} {
  if nil == p.ProtocolParameters {
    return nil
  }
  return p.ProtocolParameters.GetValue()
}

func (p *RoutingPolicyMatchCondition) SetProtocolParameters(v interface{}) error {
  if nil == p.ProtocolParameters {
    p.ProtocolParameters = NewOneOfRoutingPolicyMatchConditionProtocolParameters()
  }
  e := p.ProtocolParameters.SetValue(v)
  if nil == e {
    if nil == p.ProtocolParametersItemDiscriminator_ {
      p.ProtocolParametersItemDiscriminator_ = new(string)
    }
    *p.ProtocolParametersItemDiscriminator_ = *p.ProtocolParameters.Discriminator
  }
  return e
}



/**
Policy indicating the match rule and the action.
*/
type RoutingPolicyRule struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  If True, policies in reverse direction will be installed with the same action but source and destination swapped.
  */
  Bidirectional *bool `json:"bidirectional,omitempty"`
  
  PolicyAction *RoutingPolicyAction `json:"policyAction"`
  
  PolicyMatch *RoutingPolicyMatchCondition `json:"policyMatch"`
}

func NewRoutingPolicyRule() *RoutingPolicyRule {
  p := new(RoutingPolicyRule)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.RoutingPolicyRule"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.RoutingPolicyRule"}
  p.UnknownFields_ = map[string]interface{}{}

  p.Bidirectional = new(bool)
  *p.Bidirectional = false


  return p
}




/**
Enum for scope type, can be one of PC or PE.
*/
type ScopeType int

const(
  SCOPETYPE_UNKNOWN ScopeType = 0
  SCOPETYPE_REDACTED ScopeType = 1
  SCOPETYPE_PC ScopeType = 2
  SCOPETYPE_PE ScopeType = 3
)

// returns the name of the enum given an ordinal number
func (e *ScopeType) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "PC",
    "PE",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *ScopeType) index(name string) ScopeType {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "PC",
    "PE",
  }
  for idx := range names {
    if names[idx] == name {
      return ScopeType(idx)
    }
  }
  return SCOPETYPE_UNKNOWN
}

func (e *ScopeType) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for ScopeType:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *ScopeType) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e ScopeType) Ref() *ScopeType {
  return &e
}


/**
Site-specific stretch configuration parameters.
*/
type SiteParams struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  The VPN connection or network gateway with VTEP service used for this Layer2 stretch.
  */
  ConnectionReference *string `json:"connectionReference,omitempty"`
  
  DefaultGatewayIPAddress *import1.IPAddress `json:"defaultGatewayIPAddress,omitempty"`
  /**
  PC cluster reference. This parameter is deprecated, use pcClusterReference instead.
  */
  PcClusterId *string `json:"pcClusterId,omitempty"`
  /**
  PC cluster reference.
  */
  PcClusterReference *string `json:"pcClusterReference,omitempty"`
  
  StretchInterfaceIpAddress *import1.IPAddress `json:"stretchInterfaceIpAddress,omitempty"`
  /**
  Subnet reference. This parameter is deprecated, use stretchSubnetReference instead.
  */
  StretchSubnetId *string `json:"stretchSubnetId,omitempty"`
  /**
  Subnet reference.
  */
  StretchSubnetReference *string `json:"stretchSubnetReference,omitempty"`
  /**
  VPN Connection reference. This is a deprecated parameter which will not be available in the upcoming release, please use Layer2StretchConnection instead.
  */
  VpnConnectionId *string `json:"vpnConnectionId,omitempty"`
  
  VpnInterfaceIPAddress *import1.IPAddress `json:"vpnInterfaceIPAddress,omitempty"`
}

func NewSiteParams() *SiteParams {
  p := new(SiteParams)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.SiteParams"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.SiteParams"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}





type State int

const(
  STATE_UNKNOWN State = 0
  STATE_REDACTED State = 1
  STATE_UP State = 2
  STATE_DOWN State = 3
)

// returns the name of the enum given an ordinal number
func (e *State) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "UP",
    "DOWN",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *State) index(name string) State {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "UP",
    "DOWN",
  }
  for idx := range names {
    if names[idx] == name {
      return State(idx)
    }
  }
  return STATE_UNKNOWN
}

func (e *State) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for State:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *State) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e State) Ref() *State {
  return &e
}


/**
Up/Down status of component and message
*/
type Status struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  Message *string `json:"message,omitempty"`
  
  State *State `json:"state,omitempty"`
}

func NewStatus() *Status {
  p := new(Status)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.Status"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.Status"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Type of the connection used for stretching the subnet. The default is VPN.
*/
type StretchConnectionType int

const(
  STRETCHCONNECTIONTYPE_UNKNOWN StretchConnectionType = 0
  STRETCHCONNECTIONTYPE_REDACTED StretchConnectionType = 1
  STRETCHCONNECTIONTYPE_VPN StretchConnectionType = 2
  STRETCHCONNECTIONTYPE_VXLAN StretchConnectionType = 3
)

// returns the name of the enum given an ordinal number
func (e *StretchConnectionType) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "VPN",
    "VXLAN",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *StretchConnectionType) index(name string) StretchConnectionType {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "VPN",
    "VXLAN",
  }
  for idx := range names {
    if names[idx] == name {
      return StretchConnectionType(idx)
    }
  }
  return STRETCHCONNECTIONTYPE_UNKNOWN
}

func (e *StretchConnectionType) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for StretchConnectionType:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *StretchConnectionType) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e StretchConnectionType) Ref() *StretchConnectionType {
  return &e
}


/**
Enum describing the runtime status of this stretch configuation. This is a deprecated parameter which will not be available in the upcoming release, please use Layer2StretchTunnelState and Layer2StretchInterfaceState instead.
*/
type StretchState int

const(
  STRETCHSTATE_UNKNOWN StretchState = 0
  STRETCHSTATE_REDACTED StretchState = 1
  STRETCHSTATE_UP StretchState = 2
  STRETCHSTATE_DOWN StretchState = 3
)

// returns the name of the enum given an ordinal number
func (e *StretchState) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "UP",
    "DOWN",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *StretchState) index(name string) StretchState {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "UP",
    "DOWN",
  }
  for idx := range names {
    if names[idx] == name {
      return StretchState(idx)
    }
  }
  return STRETCHSTATE_UNKNOWN
}

func (e *StretchState) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for StretchState:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *StretchState) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e StretchState) Ref() *StretchState {
  return &e
}


/**
Current status of the layer2 extension between subnets.
*/
type StretchStatus struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Detailed text describing the runtime status of this stretch configuration.
  */
  Detail *string `json:"detail,omitempty"`
  
  InterfaceState *State `json:"interfaceState,omitempty"`
  /**
  The round-trip time, in milliseconds, between subnets in this stretch configuration.
  */
  RoundTripTimeMilliseconds *float32 `json:"roundTripTimeMilliseconds,omitempty"`
  
  State *StretchState `json:"state,omitempty"`
  
  TunnelState *State `json:"tunnelState,omitempty"`
}

func NewStretchStatus() *StretchStatus {
  p := new(StretchStatus)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.StretchStatus"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.StretchStatus"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}





type Subnet struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Name of the bridge on the host for the subnet.
  */
  BridgeName *string `json:"bridgeName,omitempty"`
  /**
  UUID of the cluster this subnet belongs to.
  */
  ClusterReference *string `json:"clusterReference,omitempty"`
  /**
  Description of the subnet.
  */
  Description *string `json:"description,omitempty"`
  
  DhcpOptions *DhcpOptions `json:"dhcpOptions,omitempty"`
  /**
  List of IPs, which are a subset from the Reserved IP address List, that must be advertised to the SDN Gateway.
  */
  DynamicIpAddresses []import1.IPAddress `json:"dynamicIpAddresses,omitempty"`
  /**
  Whether NAT must be enabled for the subnet (type Overlay only).
  */
  EnableNat *bool `json:"enableNat,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  IP configuration for the subnet.
  */
  IpConfig []IPConfig `json:"ipConfig,omitempty"`
  /**
  Whether the subnet is used for external connectivity.
  */
  IsExternal *bool `json:"isExternal,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import2.ApiLink `json:"links,omitempty"`
  
  Metadata *import1.Metadata `json:"metadata,omitempty"`
  /**
  Name of the subnet.
  */
  Name *string `json:"name"`
  /**
  UUID of the Network Function Chain entity that this subnet belongs to (type VLAN only).
  */
  NetworkFunctionChainReference *string `json:"networkFunctionChainReference,omitempty"`
  /**
  For subnet type VLAN, VLAN ID from 0..4095; for type OVERLAY, 24-bit VNI, readonly.
  */
  NetworkId *int `json:"networkId,omitempty"`
  /**
  List of IPs that are excluded while allocating IP addresses to VM ports.
  */
  ReservedIpAddresses []import1.IPAddress `json:"reservedIpAddresses,omitempty"`
  
  SubnetType *SubnetType `json:"subnetType"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
  /**
  UUID of the virtual switch this subnet belongs to (type VLAN only).
  */
  VirtualSwitchReference *string `json:"virtualSwitchReference,omitempty"`
  /**
  UUID of Virtual Private Cloud (Subnet) this subnet belongs to (type Overlay only).
  */
  VpcReference *string `json:"vpcReference,omitempty"`
}

func NewSubnet() *Subnet {
  p := new(Subnet)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.Subnet"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.Subnet"}
  p.UnknownFields_ = map[string]interface{}{}

  p.EnableNat = new(bool)
  *p.EnableNat = false
  p.IsExternal = new(bool)
  *p.IsExternal = false


  return p
}




/**
REST Response for all response codes in api path /networking/v4.0.a1/config/subnets/{subnetExtId}/addresses Get operation
*/
type SubnetAddressAssignmentListApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfSubnetAddressAssignmentListApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewSubnetAddressAssignmentListApiResponse() *SubnetAddressAssignmentListApiResponse {
  p := new(SubnetAddressAssignmentListApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.SubnetAddressAssignmentListApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.SubnetAddressAssignmentListApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *SubnetAddressAssignmentListApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *SubnetAddressAssignmentListApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfSubnetAddressAssignmentListApiResponseData()
  }
  e := p.Data.SetValue(v)
  if nil == e {
    if nil == p.DataItemDiscriminator_ {
      p.DataItemDiscriminator_ = new(string)
    }
    *p.DataItemDiscriminator_ = *p.Data.Discriminator
  }
  return e
}



/**
REST Response for all response codes in api path /networking/v4.0.a1/config/subnets/{extId} Get operation
*/
type SubnetApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfSubnetApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewSubnetApiResponse() *SubnetApiResponse {
  p := new(SubnetApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.SubnetApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.SubnetApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *SubnetApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *SubnetApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfSubnetApiResponseData()
  }
  e := p.Data.SetValue(v)
  if nil == e {
    if nil == p.DataItemDiscriminator_ {
      p.DataItemDiscriminator_ = new(string)
    }
    *p.DataItemDiscriminator_ = *p.Data.Discriminator
  }
  return e
}



/**
REST Response for all response codes in api path /networking/v4.0.a1/config/subnets Get operation
*/
type SubnetListApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfSubnetListApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewSubnetListApiResponse() *SubnetListApiResponse {
  p := new(SubnetListApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.SubnetListApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.SubnetListApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *SubnetListApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *SubnetListApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfSubnetListApiResponseData()
  }
  e := p.Data.SetValue(v)
  if nil == e {
    if nil == p.DataItemDiscriminator_ {
      p.DataItemDiscriminator_ = new(string)
    }
    *p.DataItemDiscriminator_ = *p.Data.Discriminator
  }
  return e
}



/**
Local subnet reference
*/
type SubnetReference struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Local subnet identifier
  */
  ExtId *string `json:"extId"`
  
  ObjectType *string `json:"objectType,omitempty"`
}

func NewSubnetReference() *SubnetReference {
  p := new(SubnetReference)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.SubnetReference"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.SubnetReference"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Type of subnet.
*/
type SubnetType int

const(
  SUBNETTYPE_UNKNOWN SubnetType = 0
  SUBNETTYPE_REDACTED SubnetType = 1
  SUBNETTYPE_OVERLAY SubnetType = 2
  SUBNETTYPE_VLAN SubnetType = 3
)

// returns the name of the enum given an ordinal number
func (e *SubnetType) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "OVERLAY",
    "VLAN",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *SubnetType) index(name string) SubnetType {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "OVERLAY",
    "VLAN",
  }
  for idx := range names {
    if names[idx] == name {
      return SubnetType(idx)
    }
  }
  return SUBNETTYPE_UNKNOWN
}

func (e *SubnetType) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for SubnetType:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *SubnetType) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e SubnetType) Ref() *SubnetType {
  return &e
}


/**
TCP parameters to be matched in Routing Policy.
*/
type TCP struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  DestinationPortRanges []PortRange `json:"destinationPortRanges,omitempty"`
  
  SourcePortRanges []PortRange `json:"sourcePortRanges,omitempty"`
}

func NewTCP() *TCP {
  p := new(TCP)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.TCP"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.TCP"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
TCP parameters to be matched in Routing Policy.
*/
type TCPObject struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  Tcp *TCP `json:"tcp"`
}

func NewTCPObject() *TCPObject {
  p := new(TCPObject)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.TCPObject"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.TCPObject"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Object encapsulating Task ID Return Value.
*/
type Task struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  The external identifier that can be used to retrieve the task using its URL.
  */
  ExtId *string `json:"extId,omitempty"`
}

func NewTask() *Task {
  p := new(Task)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.Task"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.Task"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST Response for  Task type
*/
type TaskApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfTaskApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewTaskApiResponse() *TaskApiResponse {
  p := new(TaskApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.TaskApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.TaskApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *TaskApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *TaskApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfTaskApiResponseData()
  }
  e := p.Data.SetValue(v)
  if nil == e {
    if nil == p.DataItemDiscriminator_ {
      p.DataItemDiscriminator_ = new(string)
    }
    *p.DataItemDiscriminator_ = *p.Data.Discriminator
  }
  return e
}



/**
REST Response for all response codes in api path /networking/v4.0.a1/config/vpn-connections/{extId} Delete operation
*/
type TaskReferenceApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfTaskReferenceApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewTaskReferenceApiResponse() *TaskReferenceApiResponse {
  p := new(TaskReferenceApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.TaskReferenceApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.TaskReferenceApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *TaskReferenceApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *TaskReferenceApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfTaskReferenceApiResponseData()
  }
  e := p.Data.SetValue(v)
  if nil == e {
    if nil == p.DataItemDiscriminator_ {
      p.DataItemDiscriminator_ = new(string)
    }
    *p.DataItemDiscriminator_ = *p.Data.Discriminator
  }
  return e
}



/**
UDP parameters to be matched in Routing Policy.
*/
type UDP struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  DestinationPortRanges []PortRange `json:"destinationPortRanges,omitempty"`
  
  SourcePortRanges []PortRange `json:"sourcePortRanges,omitempty"`
}

func NewUDP() *UDP {
  p := new(UDP)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.UDP"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.UDP"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
UDP parameters to be matched in Routing Policy.
*/
type UDPObject struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  Udp *UDP `json:"udp"`
}

func NewUDPObject() *UDPObject {
  p := new(UDPObject)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.UDPObject"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.UDPObject"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
VM NIC and VPC in which the VM NIC subnet belongs to.
*/
type VmNicAssociation struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  VM NIC reference.
  */
  VmNicReference *string `json:"vmNicReference"`
  /**
  VPC reference in which the VM NIC subnet belongs to.
  */
  VpcReference *string `json:"vpcReference,omitempty"`
}

func NewVmNicAssociation() *VmNicAssociation {
  p := new(VmNicAssociation)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.VmNicAssociation"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.VmNicAssociation"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}





type Vpc struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  CommonDhcpOptions *DhcpOptions `json:"commonDhcpOptions,omitempty"`
  /**
  Description of the VPC.
  */
  Description *string `json:"description,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  List of external subnets that the VPC is attached to.
  */
  ExternalSubnets []ExternalSubnet `json:"externalSubnets,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import2.ApiLink `json:"links,omitempty"`
  
  Metadata *import1.Metadata `json:"metadata,omitempty"`
  /**
  Name of the VPC.
  */
  Name *string `json:"name"`
  /**
  List of IP Addresses used for SNAT.
  */
  SnatIps []import1.IPAddress `json:"snatIps,omitempty"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
}

func NewVpc() *Vpc {
  p := new(Vpc)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.Vpc"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.Vpc"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST Response for all response codes in api path /networking/v4.0.a1/config/vpcs/{extId} Get operation
*/
type VpcApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfVpcApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewVpcApiResponse() *VpcApiResponse {
  p := new(VpcApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.VpcApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.VpcApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *VpcApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *VpcApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfVpcApiResponseData()
  }
  e := p.Data.SetValue(v)
  if nil == e {
    if nil == p.DataItemDiscriminator_ {
      p.DataItemDiscriminator_ = new(string)
    }
    *p.DataItemDiscriminator_ = *p.Data.Discriminator
  }
  return e
}



/**
REST Response for all response codes in api path /networking/v4.0.a1/config/vpcs Get operation
*/
type VpcListApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfVpcListApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewVpcListApiResponse() *VpcListApiResponse {
  p := new(VpcListApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.VpcListApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.VpcListApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *VpcListApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *VpcListApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfVpcListApiResponseData()
  }
  e := p.Data.SetValue(v)
  if nil == e {
    if nil == p.DataItemDiscriminator_ {
      p.DataItemDiscriminator_ = new(string)
    }
    *p.DataItemDiscriminator_ = *p.Data.Discriminator
  }
  return e
}



/**
VPN connection
*/
type VpnConnection struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  IP prefixes advertised to the remote gateway over BGP
  */
  AdvertisedPrefixes []IPSubnet `json:"advertisedPrefixes,omitempty"`
  /**
  VPN connection description
  */
  Description *string `json:"description,omitempty"`
  
  DpdConfig *DpdConfig `json:"dpdConfig,omitempty"`
  /**
  Priority assigned to routes received on this connection over eBGP. A higher priority value indicates that the routes are more preferred
  */
  DynamicRoutePriority *int `json:"dynamicRoutePriority,omitempty"`
  
  EbgpStatus *Status `json:"ebgpStatus,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  
  IpsecConfig *IpsecConfig `json:"ipsecConfig"`
  
  IpsecTunnelStatus *Status `json:"ipsecTunnelStatus,omitempty"`
  /**
  IP prefixes learned from the remote gateway over BGP
  */
  LearnedPrefixes []IPSubnet `json:"learnedPrefixes,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import2.ApiLink `json:"links,omitempty"`
  
  LocalGatewayRole *GatewayRole `json:"localGatewayRole"`
  /**
  The local VPN gateway reference
  */
  LocalNetworkGatewayReference *string `json:"localNetworkGatewayReference"`
  
  Metadata *import1.Metadata `json:"metadata,omitempty"`
  /**
  VPN connection name
  */
  Name *string `json:"name,omitempty"`
  
  QosConfig *QosConfig `json:"qosConfig,omitempty"`
  /**
  The remote VPN gateway reference
  */
  RemoteNetworkGatewayReference *string `json:"remoteNetworkGatewayReference"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
}

func NewVpnConnection() *VpnConnection {
  p := new(VpnConnection)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.VpnConnection"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.VpnConnection"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST Response for all response codes in api path /networking/v4.0.a1/config/vpn-connections/{extId} Get operation
*/
type VpnConnectionApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfVpnConnectionApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewVpnConnectionApiResponse() *VpnConnectionApiResponse {
  p := new(VpnConnectionApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.VpnConnectionApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.VpnConnectionApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *VpnConnectionApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *VpnConnectionApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfVpnConnectionApiResponseData()
  }
  e := p.Data.SetValue(v)
  if nil == e {
    if nil == p.DataItemDiscriminator_ {
      p.DataItemDiscriminator_ = new(string)
    }
    *p.DataItemDiscriminator_ = *p.Data.Discriminator
  }
  return e
}



/**
REST Response for all response codes in api path /networking/v4.0.a1/config/vpn-connections Get operation
*/
type VpnConnectionListApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfVpnConnectionListApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewVpnConnectionListApiResponse() *VpnConnectionListApiResponse {
  p := new(VpnConnectionListApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.VpnConnectionListApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.VpnConnectionListApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *VpnConnectionListApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *VpnConnectionListApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfVpnConnectionListApiResponseData()
  }
  e := p.Data.SetValue(v)
  if nil == e {
    if nil == p.DataItemDiscriminator_ {
      p.DataItemDiscriminator_ = new(string)
    }
    *p.DataItemDiscriminator_ = *p.Data.Discriminator
  }
  return e
}




type VpnConnectionReference struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  EntityType *import1.EntityType `json:"entityType,omitempty"`
  
  ExtId *string `json:"extId,omitempty"`
  
  Name *string `json:"name,omitempty"`
  
  Uris []string `json:"uris,omitempty"`
}

func NewVpnConnectionReference() *VpnConnectionReference {
  p := new(VpnConnectionReference)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.VpnConnectionReference"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.VpnConnectionReference"}
  p.UnknownFields_ = map[string]interface{}{}


  // set value for EntityType
  p.EntityType = new(import1.EntityType)
  *p.EntityType = import1.ENTITYTYPE_VPN_CONNECTION
  
  return p
}




/**
VPN device vendor name and device version.
*/
type VpnVendor struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  List of VPN device versions.
  */
  DeviceVersions []string `json:"deviceVersions,omitempty"`
  /**
  VPN vendor name.
  */
  Name *string `json:"name,omitempty"`
}

func NewVpnVendor() *VpnVendor {
  p := new(VpnVendor)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.VpnVendor"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.VpnVendor"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST Response for all response codes in api path /networking/v4.0.a1/config/vpn-vendor-configs/{vpnConnectionId}/available Get operation
*/
type VpnVendorListApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfVpnVendorListApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewVpnVendorListApiResponse() *VpnVendorListApiResponse {
  p := new(VpnVendorListApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.VpnVendorListApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.VpnVendorListApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *VpnVendorListApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *VpnVendorListApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfVpnVendorListApiResponseData()
  }
  e := p.Data.SetValue(v)
  if nil == e {
    if nil == p.DataItemDiscriminator_ {
      p.DataItemDiscriminator_ = new(string)
    }
    *p.DataItemDiscriminator_ = *p.Data.Discriminator
  }
  return e
}



/**
VXLAN Tunnel Endpoint
*/
type Vtep struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  Address *import1.IPAddress `json:"address,omitempty"`
}

func NewVtep() *Vtep {
  p := new(Vtep)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.Vtep"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.Vtep"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}





type VtepGatewayReference struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  EntityType *import1.EntityType `json:"entityType,omitempty"`
  
  ExtId *string `json:"extId,omitempty"`
  
  Name *string `json:"name,omitempty"`
  
  Uris []string `json:"uris,omitempty"`
}

func NewVtepGatewayReference() *VtepGatewayReference {
  p := new(VtepGatewayReference)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "networking.v4.config.VtepGatewayReference"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.config.VtepGatewayReference"}
  p.UnknownFields_ = map[string]interface{}{}


  // set value for EntityType
  p.EntityType = new(import1.EntityType)
  *p.EntityType = import1.ENTITYTYPE_VTEP_GATEWAY
  
  return p
}



type OneOfAncConfigApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType1003 *AncConfig `json:"-"`
  oneOfType1002 []import1.Message `json:"-"`
}

func NewOneOfAncConfigApiResponseData() *OneOfAncConfigApiResponseData {
  p := new(OneOfAncConfigApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfAncConfigApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfAncConfigApiResponseData is nil"))
  }
  switch v.(type) {
    case AncConfig:
      if nil == p.oneOfType1003 {p.oneOfType1003 = new(AncConfig)}
      *p.oneOfType1003 = v.(AncConfig)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType1003.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType1003.ObjectType_
    case []import1.Message:
      p.oneOfType1002 = v.([]import1.Message)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<common.v1.config.Message>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<common.v1.config.Message>"
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfAncConfigApiResponseData) GetValue() interface{} {
  if p.oneOfType1003 != nil && *p.oneOfType1003.ObjectType_ == *p.Discriminator {
    return *p.oneOfType1003
  }
  if "List<common.v1.config.Message>" == *p.Discriminator {
    return p.oneOfType1002
  }
  return nil
}

func (p *OneOfAncConfigApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType1003 := new(AncConfig)
  if err := json.Unmarshal(b, vOneOfType1003); err == nil {
    if "networking.v4.config.AncConfig" == *vOneOfType1003.ObjectType_ {
          if nil == p.oneOfType1003 {p.oneOfType1003 = new(AncConfig)}
      *p.oneOfType1003 = *vOneOfType1003
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType1003.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType1003.ObjectType_
      return nil
    }
    }
  vOneOfType1002 := new([]import1.Message)
  if err := json.Unmarshal(b, vOneOfType1002); err == nil {
    if len(*vOneOfType1002) == 0 || "common.v1.config.Message" == *((*vOneOfType1002)[0].ObjectType_) {
      p.oneOfType1002 = *vOneOfType1002
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<common.v1.config.Message>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<common.v1.config.Message>"
      return nil

    }
  }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfAncConfigApiResponseData"))
}

func (p *OneOfAncConfigApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType1003 != nil && *p.oneOfType1003.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType1003)
  }
  if "List<common.v1.config.Message>" == *p.Discriminator {
    return json.Marshal(p.oneOfType1002)
  }
  return nil, errors.New("No value to marshal for OneOfAncConfigApiResponseData")
}

type OneOfFloatingIpAssociation struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType1 *PrivateIpAssociation `json:"-"`
  oneOfType0 *VmNicAssociation `json:"-"`
}

func NewOneOfFloatingIpAssociation() *OneOfFloatingIpAssociation {
  p := new(OneOfFloatingIpAssociation)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfFloatingIpAssociation) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfFloatingIpAssociation is nil"))
  }
  switch v.(type) {
    case PrivateIpAssociation:
      if nil == p.oneOfType1 {p.oneOfType1 = new(PrivateIpAssociation)}
      *p.oneOfType1 = v.(PrivateIpAssociation)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType1.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType1.ObjectType_
    case VmNicAssociation:
      if nil == p.oneOfType0 {p.oneOfType0 = new(VmNicAssociation)}
      *p.oneOfType0 = v.(VmNicAssociation)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfFloatingIpAssociation) GetValue() interface{} {
  if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
    return *p.oneOfType1
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  return nil
}

func (p *OneOfFloatingIpAssociation) UnmarshalJSON(b []byte) error {
  vOneOfType1 := new(PrivateIpAssociation)
  if err := json.Unmarshal(b, vOneOfType1); err == nil {
    if "networking.v4.config.PrivateIpAssociation" == *vOneOfType1.ObjectType_ {
          if nil == p.oneOfType1 {p.oneOfType1 = new(PrivateIpAssociation)}
      *p.oneOfType1 = *vOneOfType1
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType1.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType1.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new(VmNicAssociation)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "networking.v4.config.VmNicAssociation" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(VmNicAssociation)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfFloatingIpAssociation"))
}

func (p *OneOfFloatingIpAssociation) MarshalJSON() ([]byte, error) {
  if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType1)
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfFloatingIpAssociation")
}

type OneOfTaskApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType1003 *Task `json:"-"`
  oneOfType1002 []import1.Message `json:"-"`
}

func NewOneOfTaskApiResponseData() *OneOfTaskApiResponseData {
  p := new(OneOfTaskApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfTaskApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfTaskApiResponseData is nil"))
  }
  switch v.(type) {
    case Task:
      if nil == p.oneOfType1003 {p.oneOfType1003 = new(Task)}
      *p.oneOfType1003 = v.(Task)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType1003.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType1003.ObjectType_
    case []import1.Message:
      p.oneOfType1002 = v.([]import1.Message)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<common.v1.config.Message>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<common.v1.config.Message>"
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfTaskApiResponseData) GetValue() interface{} {
  if p.oneOfType1003 != nil && *p.oneOfType1003.ObjectType_ == *p.Discriminator {
    return *p.oneOfType1003
  }
  if "List<common.v1.config.Message>" == *p.Discriminator {
    return p.oneOfType1002
  }
  return nil
}

func (p *OneOfTaskApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType1003 := new(Task)
  if err := json.Unmarshal(b, vOneOfType1003); err == nil {
    if "networking.v4.config.Task" == *vOneOfType1003.ObjectType_ {
          if nil == p.oneOfType1003 {p.oneOfType1003 = new(Task)}
      *p.oneOfType1003 = *vOneOfType1003
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType1003.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType1003.ObjectType_
      return nil
    }
    }
  vOneOfType1002 := new([]import1.Message)
  if err := json.Unmarshal(b, vOneOfType1002); err == nil {
    if len(*vOneOfType1002) == 0 || "common.v1.config.Message" == *((*vOneOfType1002)[0].ObjectType_) {
      p.oneOfType1002 = *vOneOfType1002
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<common.v1.config.Message>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<common.v1.config.Message>"
      return nil

    }
  }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfTaskApiResponseData"))
}

func (p *OneOfTaskApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType1003 != nil && *p.oneOfType1003.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType1003)
  }
  if "List<common.v1.config.Message>" == *p.Discriminator {
    return json.Marshal(p.oneOfType1002)
  }
  return nil, errors.New("No value to marshal for OneOfTaskApiResponseData")
}

type OneOfNetworkCloudConfigListApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType1003 []NetworkCloudConfig `json:"-"`
  oneOfType1002 []import1.Message `json:"-"`
}

func NewOneOfNetworkCloudConfigListApiResponseData() *OneOfNetworkCloudConfigListApiResponseData {
  p := new(OneOfNetworkCloudConfigListApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfNetworkCloudConfigListApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfNetworkCloudConfigListApiResponseData is nil"))
  }
  switch v.(type) {
    case []NetworkCloudConfig:
      p.oneOfType1003 = v.([]NetworkCloudConfig)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<networking.v4.config.NetworkCloudConfig>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<networking.v4.config.NetworkCloudConfig>"
    case []import1.Message:
      p.oneOfType1002 = v.([]import1.Message)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<common.v1.config.Message>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<common.v1.config.Message>"
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfNetworkCloudConfigListApiResponseData) GetValue() interface{} {
  if "List<networking.v4.config.NetworkCloudConfig>" == *p.Discriminator {
    return p.oneOfType1003
  }
  if "List<common.v1.config.Message>" == *p.Discriminator {
    return p.oneOfType1002
  }
  return nil
}

func (p *OneOfNetworkCloudConfigListApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType1003 := new([]NetworkCloudConfig)
  if err := json.Unmarshal(b, vOneOfType1003); err == nil {
    if len(*vOneOfType1003) == 0 || "networking.v4.config.NetworkCloudConfig" == *((*vOneOfType1003)[0].ObjectType_) {
      p.oneOfType1003 = *vOneOfType1003
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<networking.v4.config.NetworkCloudConfig>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<networking.v4.config.NetworkCloudConfig>"
      return nil

    }
  }
  vOneOfType1002 := new([]import1.Message)
  if err := json.Unmarshal(b, vOneOfType1002); err == nil {
    if len(*vOneOfType1002) == 0 || "common.v1.config.Message" == *((*vOneOfType1002)[0].ObjectType_) {
      p.oneOfType1002 = *vOneOfType1002
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<common.v1.config.Message>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<common.v1.config.Message>"
      return nil

    }
  }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfNetworkCloudConfigListApiResponseData"))
}

func (p *OneOfNetworkCloudConfigListApiResponseData) MarshalJSON() ([]byte, error) {
  if "List<networking.v4.config.NetworkCloudConfig>" == *p.Discriminator {
    return json.Marshal(p.oneOfType1003)
  }
  if "List<common.v1.config.Message>" == *p.Discriminator {
    return json.Marshal(p.oneOfType1002)
  }
  return nil, errors.New("No value to marshal for OneOfNetworkCloudConfigListApiResponseData")
}

type OneOfBgpSessionListApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import3.ErrorResponse `json:"-"`
  oneOfType0 []BgpSession `json:"-"`
}

func NewOneOfBgpSessionListApiResponseData() *OneOfBgpSessionListApiResponseData {
  p := new(OneOfBgpSessionListApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfBgpSessionListApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfBgpSessionListApiResponseData is nil"))
  }
  switch v.(type) {
    case import3.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = v.(import3.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case []BgpSession:
      p.oneOfType0 = v.([]BgpSession)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<networking.v4.config.BgpSession>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<networking.v4.config.BgpSession>"
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfBgpSessionListApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if "List<networking.v4.config.BgpSession>" == *p.Discriminator {
    return p.oneOfType0
  }
  return nil
}

func (p *OneOfBgpSessionListApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import3.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new([]BgpSession)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "networking.v4.config.BgpSession" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<networking.v4.config.BgpSession>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<networking.v4.config.BgpSession>"
      return nil

    }
  }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfBgpSessionListApiResponseData"))
}

func (p *OneOfBgpSessionListApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if "List<networking.v4.config.BgpSession>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfBgpSessionListApiResponseData")
}

type OneOfSubnetListApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 []Subnet `json:"-"`
  oneOfType400 *import3.ErrorResponse `json:"-"`
}

func NewOneOfSubnetListApiResponseData() *OneOfSubnetListApiResponseData {
  p := new(OneOfSubnetListApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfSubnetListApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfSubnetListApiResponseData is nil"))
  }
  switch v.(type) {
    case []Subnet:
      p.oneOfType0 = v.([]Subnet)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<networking.v4.config.Subnet>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<networking.v4.config.Subnet>"
    case import3.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = v.(import3.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfSubnetListApiResponseData) GetValue() interface{} {
  if "List<networking.v4.config.Subnet>" == *p.Discriminator {
    return p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfSubnetListApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new([]Subnet)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "networking.v4.config.Subnet" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<networking.v4.config.Subnet>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<networking.v4.config.Subnet>"
      return nil

    }
  }
  vOneOfType400 := new(import3.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfSubnetListApiResponseData"))
}

func (p *OneOfSubnetListApiResponseData) MarshalJSON() ([]byte, error) {
  if "List<networking.v4.config.Subnet>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfSubnetListApiResponseData")
}

type OneOfNetworkControllerListApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 []NetworkController `json:"-"`
  oneOfType400 *import3.ErrorResponse `json:"-"`
}

func NewOneOfNetworkControllerListApiResponseData() *OneOfNetworkControllerListApiResponseData {
  p := new(OneOfNetworkControllerListApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfNetworkControllerListApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfNetworkControllerListApiResponseData is nil"))
  }
  switch v.(type) {
    case []NetworkController:
      p.oneOfType0 = v.([]NetworkController)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<networking.v4.config.NetworkController>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<networking.v4.config.NetworkController>"
    case import3.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = v.(import3.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfNetworkControllerListApiResponseData) GetValue() interface{} {
  if "List<networking.v4.config.NetworkController>" == *p.Discriminator {
    return p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfNetworkControllerListApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new([]NetworkController)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "networking.v4.config.NetworkController" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<networking.v4.config.NetworkController>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<networking.v4.config.NetworkController>"
      return nil

    }
  }
  vOneOfType400 := new(import3.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfNetworkControllerListApiResponseData"))
}

func (p *OneOfNetworkControllerListApiResponseData) MarshalJSON() ([]byte, error) {
  if "List<networking.v4.config.NetworkController>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfNetworkControllerListApiResponseData")
}

type OneOfVpcApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import3.ErrorResponse `json:"-"`
  oneOfType0 *Vpc `json:"-"`
}

func NewOneOfVpcApiResponseData() *OneOfVpcApiResponseData {
  p := new(OneOfVpcApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfVpcApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfVpcApiResponseData is nil"))
  }
  switch v.(type) {
    case import3.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = v.(import3.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case Vpc:
      if nil == p.oneOfType0 {p.oneOfType0 = new(Vpc)}
      *p.oneOfType0 = v.(Vpc)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfVpcApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  return nil
}

func (p *OneOfVpcApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import3.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new(Vpc)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "networking.v4.config.Vpc" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(Vpc)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfVpcApiResponseData"))
}

func (p *OneOfVpcApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfVpcApiResponseData")
}

type OneOfFloatingIpApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *FloatingIp `json:"-"`
  oneOfType400 *import3.ErrorResponse `json:"-"`
}

func NewOneOfFloatingIpApiResponseData() *OneOfFloatingIpApiResponseData {
  p := new(OneOfFloatingIpApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfFloatingIpApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfFloatingIpApiResponseData is nil"))
  }
  switch v.(type) {
    case FloatingIp:
      if nil == p.oneOfType0 {p.oneOfType0 = new(FloatingIp)}
      *p.oneOfType0 = v.(FloatingIp)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    case import3.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = v.(import3.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfFloatingIpApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfFloatingIpApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new(FloatingIp)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "networking.v4.config.FloatingIp" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(FloatingIp)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  vOneOfType400 := new(import3.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfFloatingIpApiResponseData"))
}

func (p *OneOfFloatingIpApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfFloatingIpApiResponseData")
}

type OneOfIPFIXExporterListApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import3.ErrorResponse `json:"-"`
  oneOfType0 []IPFIXExporter `json:"-"`
}

func NewOneOfIPFIXExporterListApiResponseData() *OneOfIPFIXExporterListApiResponseData {
  p := new(OneOfIPFIXExporterListApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfIPFIXExporterListApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfIPFIXExporterListApiResponseData is nil"))
  }
  switch v.(type) {
    case import3.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = v.(import3.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case []IPFIXExporter:
      p.oneOfType0 = v.([]IPFIXExporter)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<networking.v4.config.IPFIXExporter>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<networking.v4.config.IPFIXExporter>"
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfIPFIXExporterListApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if "List<networking.v4.config.IPFIXExporter>" == *p.Discriminator {
    return p.oneOfType0
  }
  return nil
}

func (p *OneOfIPFIXExporterListApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import3.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new([]IPFIXExporter)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "networking.v4.config.IPFIXExporter" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<networking.v4.config.IPFIXExporter>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<networking.v4.config.IPFIXExporter>"
      return nil

    }
  }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfIPFIXExporterListApiResponseData"))
}

func (p *OneOfIPFIXExporterListApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if "List<networking.v4.config.IPFIXExporter>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfIPFIXExporterListApiResponseData")
}

type OneOfSubnetApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import3.ErrorResponse `json:"-"`
  oneOfType0 *Subnet `json:"-"`
}

func NewOneOfSubnetApiResponseData() *OneOfSubnetApiResponseData {
  p := new(OneOfSubnetApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfSubnetApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfSubnetApiResponseData is nil"))
  }
  switch v.(type) {
    case import3.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = v.(import3.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case Subnet:
      if nil == p.oneOfType0 {p.oneOfType0 = new(Subnet)}
      *p.oneOfType0 = v.(Subnet)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfSubnetApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  return nil
}

func (p *OneOfSubnetApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import3.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new(Subnet)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "networking.v4.config.Subnet" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(Subnet)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfSubnetApiResponseData"))
}

func (p *OneOfSubnetApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfSubnetApiResponseData")
}

type OneOfClusterFlowStatusApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *ClusterFlowStatus `json:"-"`
  oneOfType400 *import3.ErrorResponse `json:"-"`
}

func NewOneOfClusterFlowStatusApiResponseData() *OneOfClusterFlowStatusApiResponseData {
  p := new(OneOfClusterFlowStatusApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfClusterFlowStatusApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfClusterFlowStatusApiResponseData is nil"))
  }
  switch v.(type) {
    case ClusterFlowStatus:
      if nil == p.oneOfType0 {p.oneOfType0 = new(ClusterFlowStatus)}
      *p.oneOfType0 = v.(ClusterFlowStatus)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    case import3.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = v.(import3.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfClusterFlowStatusApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfClusterFlowStatusApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new(ClusterFlowStatus)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "networking.v4.config.ClusterFlowStatus" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(ClusterFlowStatus)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  vOneOfType400 := new(import3.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfClusterFlowStatusApiResponseData"))
}

func (p *OneOfClusterFlowStatusApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfClusterFlowStatusApiResponseData")
}

type OneOfRoutingPolicyMatchConditionDestination struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType1 *IPSubnetObject `json:"-"`
  oneOfType0 *AddressTypeObject `json:"-"`
}

func NewOneOfRoutingPolicyMatchConditionDestination() *OneOfRoutingPolicyMatchConditionDestination {
  p := new(OneOfRoutingPolicyMatchConditionDestination)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfRoutingPolicyMatchConditionDestination) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfRoutingPolicyMatchConditionDestination is nil"))
  }
  switch v.(type) {
    case IPSubnetObject:
      if nil == p.oneOfType1 {p.oneOfType1 = new(IPSubnetObject)}
      *p.oneOfType1 = v.(IPSubnetObject)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType1.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType1.ObjectType_
    case AddressTypeObject:
      if nil == p.oneOfType0 {p.oneOfType0 = new(AddressTypeObject)}
      *p.oneOfType0 = v.(AddressTypeObject)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfRoutingPolicyMatchConditionDestination) GetValue() interface{} {
  if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
    return *p.oneOfType1
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  return nil
}

func (p *OneOfRoutingPolicyMatchConditionDestination) UnmarshalJSON(b []byte) error {
  vOneOfType1 := new(IPSubnetObject)
  if err := json.Unmarshal(b, vOneOfType1); err == nil {
    if "networking.v4.config.IPSubnetObject" == *vOneOfType1.ObjectType_ {
          if nil == p.oneOfType1 {p.oneOfType1 = new(IPSubnetObject)}
      *p.oneOfType1 = *vOneOfType1
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType1.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType1.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new(AddressTypeObject)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "networking.v4.config.AddressTypeObject" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(AddressTypeObject)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfRoutingPolicyMatchConditionDestination"))
}

func (p *OneOfRoutingPolicyMatchConditionDestination) MarshalJSON() ([]byte, error) {
  if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType1)
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfRoutingPolicyMatchConditionDestination")
}

type OneOfFlowGatewayListApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType1003 []FlowGateway `json:"-"`
  oneOfType1002 []import1.Message `json:"-"`
}

func NewOneOfFlowGatewayListApiResponseData() *OneOfFlowGatewayListApiResponseData {
  p := new(OneOfFlowGatewayListApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfFlowGatewayListApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfFlowGatewayListApiResponseData is nil"))
  }
  switch v.(type) {
    case []FlowGateway:
      p.oneOfType1003 = v.([]FlowGateway)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<networking.v4.config.FlowGateway>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<networking.v4.config.FlowGateway>"
    case []import1.Message:
      p.oneOfType1002 = v.([]import1.Message)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<common.v1.config.Message>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<common.v1.config.Message>"
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfFlowGatewayListApiResponseData) GetValue() interface{} {
  if "List<networking.v4.config.FlowGateway>" == *p.Discriminator {
    return p.oneOfType1003
  }
  if "List<common.v1.config.Message>" == *p.Discriminator {
    return p.oneOfType1002
  }
  return nil
}

func (p *OneOfFlowGatewayListApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType1003 := new([]FlowGateway)
  if err := json.Unmarshal(b, vOneOfType1003); err == nil {
    if len(*vOneOfType1003) == 0 || "networking.v4.config.FlowGateway" == *((*vOneOfType1003)[0].ObjectType_) {
      p.oneOfType1003 = *vOneOfType1003
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<networking.v4.config.FlowGateway>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<networking.v4.config.FlowGateway>"
      return nil

    }
  }
  vOneOfType1002 := new([]import1.Message)
  if err := json.Unmarshal(b, vOneOfType1002); err == nil {
    if len(*vOneOfType1002) == 0 || "common.v1.config.Message" == *((*vOneOfType1002)[0].ObjectType_) {
      p.oneOfType1002 = *vOneOfType1002
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<common.v1.config.Message>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<common.v1.config.Message>"
      return nil

    }
  }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfFlowGatewayListApiResponseData"))
}

func (p *OneOfFlowGatewayListApiResponseData) MarshalJSON() ([]byte, error) {
  if "List<networking.v4.config.FlowGateway>" == *p.Discriminator {
    return json.Marshal(p.oneOfType1003)
  }
  if "List<common.v1.config.Message>" == *p.Discriminator {
    return json.Marshal(p.oneOfType1002)
  }
  return nil, errors.New("No value to marshal for OneOfFlowGatewayListApiResponseData")
}

type OneOfIpReserveInputInput struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType1 *IpReserveRangeInput `json:"-"`
  oneOfType0 *IpReserveCountInput `json:"-"`
  oneOfType2 *IpReserveIpsInput `json:"-"`
}

func NewOneOfIpReserveInputInput() *OneOfIpReserveInputInput {
  p := new(OneOfIpReserveInputInput)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfIpReserveInputInput) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfIpReserveInputInput is nil"))
  }
  switch v.(type) {
    case IpReserveRangeInput:
      if nil == p.oneOfType1 {p.oneOfType1 = new(IpReserveRangeInput)}
      *p.oneOfType1 = v.(IpReserveRangeInput)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType1.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType1.ObjectType_
    case IpReserveCountInput:
      if nil == p.oneOfType0 {p.oneOfType0 = new(IpReserveCountInput)}
      *p.oneOfType0 = v.(IpReserveCountInput)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    case IpReserveIpsInput:
      if nil == p.oneOfType2 {p.oneOfType2 = new(IpReserveIpsInput)}
      *p.oneOfType2 = v.(IpReserveIpsInput)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType2.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType2.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfIpReserveInputInput) GetValue() interface{} {
  if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
    return *p.oneOfType1
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType2 != nil && *p.oneOfType2.ObjectType_ == *p.Discriminator {
    return *p.oneOfType2
  }
  return nil
}

func (p *OneOfIpReserveInputInput) UnmarshalJSON(b []byte) error {
  vOneOfType1 := new(IpReserveRangeInput)
  if err := json.Unmarshal(b, vOneOfType1); err == nil {
    if "networking.v4.config.IpReserveRangeInput" == *vOneOfType1.ObjectType_ {
          if nil == p.oneOfType1 {p.oneOfType1 = new(IpReserveRangeInput)}
      *p.oneOfType1 = *vOneOfType1
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType1.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType1.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new(IpReserveCountInput)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "networking.v4.config.IpReserveCountInput" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(IpReserveCountInput)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  vOneOfType2 := new(IpReserveIpsInput)
  if err := json.Unmarshal(b, vOneOfType2); err == nil {
    if "networking.v4.config.IpReserveIpsInput" == *vOneOfType2.ObjectType_ {
          if nil == p.oneOfType2 {p.oneOfType2 = new(IpReserveIpsInput)}
      *p.oneOfType2 = *vOneOfType2
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType2.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType2.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfIpReserveInputInput"))
}

func (p *OneOfIpReserveInputInput) MarshalJSON() ([]byte, error) {
  if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType1)
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType2 != nil && *p.oneOfType2.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType2)
  }
  return nil, errors.New("No value to marshal for OneOfIpReserveInputInput")
}

type OneOfVpcListApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import3.ErrorResponse `json:"-"`
  oneOfType0 []Vpc `json:"-"`
}

func NewOneOfVpcListApiResponseData() *OneOfVpcListApiResponseData {
  p := new(OneOfVpcListApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfVpcListApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfVpcListApiResponseData is nil"))
  }
  switch v.(type) {
    case import3.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = v.(import3.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case []Vpc:
      p.oneOfType0 = v.([]Vpc)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<networking.v4.config.Vpc>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<networking.v4.config.Vpc>"
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfVpcListApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if "List<networking.v4.config.Vpc>" == *p.Discriminator {
    return p.oneOfType0
  }
  return nil
}

func (p *OneOfVpcListApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import3.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new([]Vpc)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "networking.v4.config.Vpc" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<networking.v4.config.Vpc>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<networking.v4.config.Vpc>"
      return nil

    }
  }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfVpcListApiResponseData"))
}

func (p *OneOfVpcListApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if "List<networking.v4.config.Vpc>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfVpcListApiResponseData")
}

type OneOfLayer2StretchApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import3.ErrorResponse `json:"-"`
  oneOfType0 *Layer2Stretch `json:"-"`
}

func NewOneOfLayer2StretchApiResponseData() *OneOfLayer2StretchApiResponseData {
  p := new(OneOfLayer2StretchApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfLayer2StretchApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfLayer2StretchApiResponseData is nil"))
  }
  switch v.(type) {
    case import3.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = v.(import3.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case Layer2Stretch:
      if nil == p.oneOfType0 {p.oneOfType0 = new(Layer2Stretch)}
      *p.oneOfType0 = v.(Layer2Stretch)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfLayer2StretchApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  return nil
}

func (p *OneOfLayer2StretchApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import3.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new(Layer2Stretch)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "networking.v4.config.Layer2Stretch" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(Layer2Stretch)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfLayer2StretchApiResponseData"))
}

func (p *OneOfLayer2StretchApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfLayer2StretchApiResponseData")
}

type OneOfNetworkGatewayListApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import3.ErrorResponse `json:"-"`
  oneOfType0 []NetworkGateway `json:"-"`
}

func NewOneOfNetworkGatewayListApiResponseData() *OneOfNetworkGatewayListApiResponseData {
  p := new(OneOfNetworkGatewayListApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfNetworkGatewayListApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfNetworkGatewayListApiResponseData is nil"))
  }
  switch v.(type) {
    case import3.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = v.(import3.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case []NetworkGateway:
      p.oneOfType0 = v.([]NetworkGateway)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<networking.v4.config.NetworkGateway>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<networking.v4.config.NetworkGateway>"
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfNetworkGatewayListApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if "List<networking.v4.config.NetworkGateway>" == *p.Discriminator {
    return p.oneOfType0
  }
  return nil
}

func (p *OneOfNetworkGatewayListApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import3.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new([]NetworkGateway)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "networking.v4.config.NetworkGateway" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<networking.v4.config.NetworkGateway>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<networking.v4.config.NetworkGateway>"
      return nil

    }
  }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfNetworkGatewayListApiResponseData"))
}

func (p *OneOfNetworkGatewayListApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if "List<networking.v4.config.NetworkGateway>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfNetworkGatewayListApiResponseData")
}

type OneOfRoutingPolicyListApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import3.ErrorResponse `json:"-"`
  oneOfType0 []RoutingPolicy `json:"-"`
}

func NewOneOfRoutingPolicyListApiResponseData() *OneOfRoutingPolicyListApiResponseData {
  p := new(OneOfRoutingPolicyListApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfRoutingPolicyListApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfRoutingPolicyListApiResponseData is nil"))
  }
  switch v.(type) {
    case import3.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = v.(import3.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case []RoutingPolicy:
      p.oneOfType0 = v.([]RoutingPolicy)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<networking.v4.config.RoutingPolicy>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<networking.v4.config.RoutingPolicy>"
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfRoutingPolicyListApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if "List<networking.v4.config.RoutingPolicy>" == *p.Discriminator {
    return p.oneOfType0
  }
  return nil
}

func (p *OneOfRoutingPolicyListApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import3.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new([]RoutingPolicy)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "networking.v4.config.RoutingPolicy" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<networking.v4.config.RoutingPolicy>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<networking.v4.config.RoutingPolicy>"
      return nil

    }
  }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfRoutingPolicyListApiResponseData"))
}

func (p *OneOfRoutingPolicyListApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if "List<networking.v4.config.RoutingPolicy>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfRoutingPolicyListApiResponseData")
}

type OneOfNetworkControllerApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import3.ErrorResponse `json:"-"`
  oneOfType0 *NetworkController `json:"-"`
}

func NewOneOfNetworkControllerApiResponseData() *OneOfNetworkControllerApiResponseData {
  p := new(OneOfNetworkControllerApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfNetworkControllerApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfNetworkControllerApiResponseData is nil"))
  }
  switch v.(type) {
    case import3.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = v.(import3.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case NetworkController:
      if nil == p.oneOfType0 {p.oneOfType0 = new(NetworkController)}
      *p.oneOfType0 = v.(NetworkController)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfNetworkControllerApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  return nil
}

func (p *OneOfNetworkControllerApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import3.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new(NetworkController)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "networking.v4.config.NetworkController" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(NetworkController)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfNetworkControllerApiResponseData"))
}

func (p *OneOfNetworkControllerApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfNetworkControllerApiResponseData")
}

type OneOfVpnVendorListApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import3.ErrorResponse `json:"-"`
  oneOfType0 []VpnVendor `json:"-"`
}

func NewOneOfVpnVendorListApiResponseData() *OneOfVpnVendorListApiResponseData {
  p := new(OneOfVpnVendorListApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfVpnVendorListApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfVpnVendorListApiResponseData is nil"))
  }
  switch v.(type) {
    case import3.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = v.(import3.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case []VpnVendor:
      p.oneOfType0 = v.([]VpnVendor)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<networking.v4.config.VpnVendor>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<networking.v4.config.VpnVendor>"
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfVpnVendorListApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if "List<networking.v4.config.VpnVendor>" == *p.Discriminator {
    return p.oneOfType0
  }
  return nil
}

func (p *OneOfVpnVendorListApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import3.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new([]VpnVendor)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "networking.v4.config.VpnVendor" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<networking.v4.config.VpnVendor>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<networking.v4.config.VpnVendor>"
      return nil

    }
  }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfVpnVendorListApiResponseData"))
}

func (p *OneOfVpnVendorListApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if "List<networking.v4.config.VpnVendor>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfVpnVendorListApiResponseData")
}

type OneOfDirectConnectVirtualInterfaceApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *DirectConnectVirtualInterface `json:"-"`
  oneOfType400 *import3.ErrorResponse `json:"-"`
}

func NewOneOfDirectConnectVirtualInterfaceApiResponseData() *OneOfDirectConnectVirtualInterfaceApiResponseData {
  p := new(OneOfDirectConnectVirtualInterfaceApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfDirectConnectVirtualInterfaceApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfDirectConnectVirtualInterfaceApiResponseData is nil"))
  }
  switch v.(type) {
    case DirectConnectVirtualInterface:
      if nil == p.oneOfType0 {p.oneOfType0 = new(DirectConnectVirtualInterface)}
      *p.oneOfType0 = v.(DirectConnectVirtualInterface)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    case import3.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = v.(import3.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfDirectConnectVirtualInterfaceApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfDirectConnectVirtualInterfaceApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new(DirectConnectVirtualInterface)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "networking.v4.config.DirectConnectVirtualInterface" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(DirectConnectVirtualInterface)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  vOneOfType400 := new(import3.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDirectConnectVirtualInterfaceApiResponseData"))
}

func (p *OneOfDirectConnectVirtualInterfaceApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfDirectConnectVirtualInterfaceApiResponseData")
}

type OneOfIPFIXExporterApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import3.ErrorResponse `json:"-"`
  oneOfType0 *IPFIXExporter `json:"-"`
}

func NewOneOfIPFIXExporterApiResponseData() *OneOfIPFIXExporterApiResponseData {
  p := new(OneOfIPFIXExporterApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfIPFIXExporterApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfIPFIXExporterApiResponseData is nil"))
  }
  switch v.(type) {
    case import3.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = v.(import3.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case IPFIXExporter:
      if nil == p.oneOfType0 {p.oneOfType0 = new(IPFIXExporter)}
      *p.oneOfType0 = v.(IPFIXExporter)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfIPFIXExporterApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  return nil
}

func (p *OneOfIPFIXExporterApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import3.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new(IPFIXExporter)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "networking.v4.config.IPFIXExporter" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(IPFIXExporter)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfIPFIXExporterApiResponseData"))
}

func (p *OneOfIPFIXExporterApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfIPFIXExporterApiResponseData")
}

type OneOfRoutingPolicyApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import3.ErrorResponse `json:"-"`
  oneOfType0 *RoutingPolicy `json:"-"`
}

func NewOneOfRoutingPolicyApiResponseData() *OneOfRoutingPolicyApiResponseData {
  p := new(OneOfRoutingPolicyApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfRoutingPolicyApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfRoutingPolicyApiResponseData is nil"))
  }
  switch v.(type) {
    case import3.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = v.(import3.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case RoutingPolicy:
      if nil == p.oneOfType0 {p.oneOfType0 = new(RoutingPolicy)}
      *p.oneOfType0 = v.(RoutingPolicy)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfRoutingPolicyApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  return nil
}

func (p *OneOfRoutingPolicyApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import3.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new(RoutingPolicy)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "networking.v4.config.RoutingPolicy" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(RoutingPolicy)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfRoutingPolicyApiResponseData"))
}

func (p *OneOfRoutingPolicyApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfRoutingPolicyApiResponseData")
}

type OneOfNetworkGatewayApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *NetworkGateway `json:"-"`
  oneOfType400 *import3.ErrorResponse `json:"-"`
}

func NewOneOfNetworkGatewayApiResponseData() *OneOfNetworkGatewayApiResponseData {
  p := new(OneOfNetworkGatewayApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfNetworkGatewayApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfNetworkGatewayApiResponseData is nil"))
  }
  switch v.(type) {
    case NetworkGateway:
      if nil == p.oneOfType0 {p.oneOfType0 = new(NetworkGateway)}
      *p.oneOfType0 = v.(NetworkGateway)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    case import3.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = v.(import3.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfNetworkGatewayApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfNetworkGatewayApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new(NetworkGateway)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "networking.v4.config.NetworkGateway" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(NetworkGateway)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  vOneOfType400 := new(import3.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfNetworkGatewayApiResponseData"))
}

func (p *OneOfNetworkGatewayApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfNetworkGatewayApiResponseData")
}

type OneOfDirectConnectApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import3.ErrorResponse `json:"-"`
  oneOfType0 *DirectConnect `json:"-"`
}

func NewOneOfDirectConnectApiResponseData() *OneOfDirectConnectApiResponseData {
  p := new(OneOfDirectConnectApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfDirectConnectApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfDirectConnectApiResponseData is nil"))
  }
  switch v.(type) {
    case import3.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = v.(import3.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case DirectConnect:
      if nil == p.oneOfType0 {p.oneOfType0 = new(DirectConnect)}
      *p.oneOfType0 = v.(DirectConnect)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfDirectConnectApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  return nil
}

func (p *OneOfDirectConnectApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import3.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new(DirectConnect)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "networking.v4.config.DirectConnect" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(DirectConnect)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDirectConnectApiResponseData"))
}

func (p *OneOfDirectConnectApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfDirectConnectApiResponseData")
}

type OneOfIpUnreserveInputInput struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *IpUnreserveIpsInput `json:"-"`
  oneOfType2 *IpUnreserveClientContextInput `json:"-"`
  oneOfType1 *IpUnreserveRangeInput `json:"-"`
}

func NewOneOfIpUnreserveInputInput() *OneOfIpUnreserveInputInput {
  p := new(OneOfIpUnreserveInputInput)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfIpUnreserveInputInput) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfIpUnreserveInputInput is nil"))
  }
  switch v.(type) {
    case IpUnreserveIpsInput:
      if nil == p.oneOfType0 {p.oneOfType0 = new(IpUnreserveIpsInput)}
      *p.oneOfType0 = v.(IpUnreserveIpsInput)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    case IpUnreserveClientContextInput:
      if nil == p.oneOfType2 {p.oneOfType2 = new(IpUnreserveClientContextInput)}
      *p.oneOfType2 = v.(IpUnreserveClientContextInput)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType2.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType2.ObjectType_
    case IpUnreserveRangeInput:
      if nil == p.oneOfType1 {p.oneOfType1 = new(IpUnreserveRangeInput)}
      *p.oneOfType1 = v.(IpUnreserveRangeInput)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType1.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType1.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfIpUnreserveInputInput) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType2 != nil && *p.oneOfType2.ObjectType_ == *p.Discriminator {
    return *p.oneOfType2
  }
  if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
    return *p.oneOfType1
  }
  return nil
}

func (p *OneOfIpUnreserveInputInput) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new(IpUnreserveIpsInput)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "networking.v4.config.IpUnreserveIpsInput" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(IpUnreserveIpsInput)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  vOneOfType2 := new(IpUnreserveClientContextInput)
  if err := json.Unmarshal(b, vOneOfType2); err == nil {
    if "networking.v4.config.IpUnreserveClientContextInput" == *vOneOfType2.ObjectType_ {
          if nil == p.oneOfType2 {p.oneOfType2 = new(IpUnreserveClientContextInput)}
      *p.oneOfType2 = *vOneOfType2
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType2.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType2.ObjectType_
      return nil
    }
    }
  vOneOfType1 := new(IpUnreserveRangeInput)
  if err := json.Unmarshal(b, vOneOfType1); err == nil {
    if "networking.v4.config.IpUnreserveRangeInput" == *vOneOfType1.ObjectType_ {
          if nil == p.oneOfType1 {p.oneOfType1 = new(IpUnreserveRangeInput)}
      *p.oneOfType1 = *vOneOfType1
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType1.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType1.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfIpUnreserveInputInput"))
}

func (p *OneOfIpUnreserveInputInput) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType2 != nil && *p.oneOfType2.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType2)
  }
  if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType1)
  }
  return nil, errors.New("No value to marshal for OneOfIpUnreserveInputInput")
}

type OneOfDirectConnectVirtualInterfaceListApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 []DirectConnectVirtualInterface `json:"-"`
  oneOfType400 *import3.ErrorResponse `json:"-"`
}

func NewOneOfDirectConnectVirtualInterfaceListApiResponseData() *OneOfDirectConnectVirtualInterfaceListApiResponseData {
  p := new(OneOfDirectConnectVirtualInterfaceListApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfDirectConnectVirtualInterfaceListApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfDirectConnectVirtualInterfaceListApiResponseData is nil"))
  }
  switch v.(type) {
    case []DirectConnectVirtualInterface:
      p.oneOfType0 = v.([]DirectConnectVirtualInterface)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<networking.v4.config.DirectConnectVirtualInterface>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<networking.v4.config.DirectConnectVirtualInterface>"
    case import3.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = v.(import3.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfDirectConnectVirtualInterfaceListApiResponseData) GetValue() interface{} {
  if "List<networking.v4.config.DirectConnectVirtualInterface>" == *p.Discriminator {
    return p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfDirectConnectVirtualInterfaceListApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new([]DirectConnectVirtualInterface)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "networking.v4.config.DirectConnectVirtualInterface" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<networking.v4.config.DirectConnectVirtualInterface>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<networking.v4.config.DirectConnectVirtualInterface>"
      return nil

    }
  }
  vOneOfType400 := new(import3.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDirectConnectVirtualInterfaceListApiResponseData"))
}

func (p *OneOfDirectConnectVirtualInterfaceListApiResponseData) MarshalJSON() ([]byte, error) {
  if "List<networking.v4.config.DirectConnectVirtualInterface>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfDirectConnectVirtualInterfaceListApiResponseData")
}

type OneOfFloatingIpListApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 []FloatingIp `json:"-"`
  oneOfType400 *import3.ErrorResponse `json:"-"`
}

func NewOneOfFloatingIpListApiResponseData() *OneOfFloatingIpListApiResponseData {
  p := new(OneOfFloatingIpListApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfFloatingIpListApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfFloatingIpListApiResponseData is nil"))
  }
  switch v.(type) {
    case []FloatingIp:
      p.oneOfType0 = v.([]FloatingIp)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<networking.v4.config.FloatingIp>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<networking.v4.config.FloatingIp>"
    case import3.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = v.(import3.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfFloatingIpListApiResponseData) GetValue() interface{} {
  if "List<networking.v4.config.FloatingIp>" == *p.Discriminator {
    return p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfFloatingIpListApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new([]FloatingIp)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "networking.v4.config.FloatingIp" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<networking.v4.config.FloatingIp>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<networking.v4.config.FloatingIp>"
      return nil

    }
  }
  vOneOfType400 := new(import3.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfFloatingIpListApiResponseData"))
}

func (p *OneOfFloatingIpListApiResponseData) MarshalJSON() ([]byte, error) {
  if "List<networking.v4.config.FloatingIp>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfFloatingIpListApiResponseData")
}

type OneOfLayer2StretchListApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import3.ErrorResponse `json:"-"`
  oneOfType0 []Layer2Stretch `json:"-"`
}

func NewOneOfLayer2StretchListApiResponseData() *OneOfLayer2StretchListApiResponseData {
  p := new(OneOfLayer2StretchListApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfLayer2StretchListApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfLayer2StretchListApiResponseData is nil"))
  }
  switch v.(type) {
    case import3.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = v.(import3.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case []Layer2Stretch:
      p.oneOfType0 = v.([]Layer2Stretch)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<networking.v4.config.Layer2Stretch>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<networking.v4.config.Layer2Stretch>"
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfLayer2StretchListApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if "List<networking.v4.config.Layer2Stretch>" == *p.Discriminator {
    return p.oneOfType0
  }
  return nil
}

func (p *OneOfLayer2StretchListApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import3.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new([]Layer2Stretch)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "networking.v4.config.Layer2Stretch" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<networking.v4.config.Layer2Stretch>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<networking.v4.config.Layer2Stretch>"
      return nil

    }
  }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfLayer2StretchListApiResponseData"))
}

func (p *OneOfLayer2StretchListApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if "List<networking.v4.config.Layer2Stretch>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfLayer2StretchListApiResponseData")
}

type OneOfRouteTableListApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import3.ErrorResponse `json:"-"`
  oneOfType0 []RouteTable `json:"-"`
}

func NewOneOfRouteTableListApiResponseData() *OneOfRouteTableListApiResponseData {
  p := new(OneOfRouteTableListApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfRouteTableListApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfRouteTableListApiResponseData is nil"))
  }
  switch v.(type) {
    case import3.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = v.(import3.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case []RouteTable:
      p.oneOfType0 = v.([]RouteTable)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<networking.v4.config.RouteTable>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<networking.v4.config.RouteTable>"
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfRouteTableListApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if "List<networking.v4.config.RouteTable>" == *p.Discriminator {
    return p.oneOfType0
  }
  return nil
}

func (p *OneOfRouteTableListApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import3.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new([]RouteTable)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "networking.v4.config.RouteTable" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<networking.v4.config.RouteTable>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<networking.v4.config.RouteTable>"
      return nil

    }
  }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfRouteTableListApiResponseData"))
}

func (p *OneOfRouteTableListApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if "List<networking.v4.config.RouteTable>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfRouteTableListApiResponseData")
}

type OneOfSubnetAddressAssignmentListApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import3.ErrorResponse `json:"-"`
  oneOfType0 []AddressAssignmentDto `json:"-"`
}

func NewOneOfSubnetAddressAssignmentListApiResponseData() *OneOfSubnetAddressAssignmentListApiResponseData {
  p := new(OneOfSubnetAddressAssignmentListApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfSubnetAddressAssignmentListApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfSubnetAddressAssignmentListApiResponseData is nil"))
  }
  switch v.(type) {
    case import3.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = v.(import3.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case []AddressAssignmentDto:
      p.oneOfType0 = v.([]AddressAssignmentDto)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<networking.v4.config.AddressAssignmentDto>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<networking.v4.config.AddressAssignmentDto>"
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfSubnetAddressAssignmentListApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if "List<networking.v4.config.AddressAssignmentDto>" == *p.Discriminator {
    return p.oneOfType0
  }
  return nil
}

func (p *OneOfSubnetAddressAssignmentListApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import3.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new([]AddressAssignmentDto)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "networking.v4.config.AddressAssignmentDto" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<networking.v4.config.AddressAssignmentDto>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<networking.v4.config.AddressAssignmentDto>"
      return nil

    }
  }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfSubnetAddressAssignmentListApiResponseData"))
}

func (p *OneOfSubnetAddressAssignmentListApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if "List<networking.v4.config.AddressAssignmentDto>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfSubnetAddressAssignmentListApiResponseData")
}

type OneOfNetworkCloudConfigApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType1003 *NetworkCloudConfig `json:"-"`
  oneOfType1002 []import1.Message `json:"-"`
}

func NewOneOfNetworkCloudConfigApiResponseData() *OneOfNetworkCloudConfigApiResponseData {
  p := new(OneOfNetworkCloudConfigApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfNetworkCloudConfigApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfNetworkCloudConfigApiResponseData is nil"))
  }
  switch v.(type) {
    case NetworkCloudConfig:
      if nil == p.oneOfType1003 {p.oneOfType1003 = new(NetworkCloudConfig)}
      *p.oneOfType1003 = v.(NetworkCloudConfig)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType1003.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType1003.ObjectType_
    case []import1.Message:
      p.oneOfType1002 = v.([]import1.Message)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<common.v1.config.Message>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<common.v1.config.Message>"
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfNetworkCloudConfigApiResponseData) GetValue() interface{} {
  if p.oneOfType1003 != nil && *p.oneOfType1003.ObjectType_ == *p.Discriminator {
    return *p.oneOfType1003
  }
  if "List<common.v1.config.Message>" == *p.Discriminator {
    return p.oneOfType1002
  }
  return nil
}

func (p *OneOfNetworkCloudConfigApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType1003 := new(NetworkCloudConfig)
  if err := json.Unmarshal(b, vOneOfType1003); err == nil {
    if "networking.v4.config.NetworkCloudConfig" == *vOneOfType1003.ObjectType_ {
          if nil == p.oneOfType1003 {p.oneOfType1003 = new(NetworkCloudConfig)}
      *p.oneOfType1003 = *vOneOfType1003
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType1003.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType1003.ObjectType_
      return nil
    }
    }
  vOneOfType1002 := new([]import1.Message)
  if err := json.Unmarshal(b, vOneOfType1002); err == nil {
    if len(*vOneOfType1002) == 0 || "common.v1.config.Message" == *((*vOneOfType1002)[0].ObjectType_) {
      p.oneOfType1002 = *vOneOfType1002
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<common.v1.config.Message>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<common.v1.config.Message>"
      return nil

    }
  }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfNetworkCloudConfigApiResponseData"))
}

func (p *OneOfNetworkCloudConfigApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType1003 != nil && *p.oneOfType1003.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType1003)
  }
  if "List<common.v1.config.Message>" == *p.Discriminator {
    return json.Marshal(p.oneOfType1002)
  }
  return nil, errors.New("No value to marshal for OneOfNetworkCloudConfigApiResponseData")
}

type OneOfRoutingPolicyMatchConditionProtocolParameters struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType3 *ProtocolNumberObject `json:"-"`
  oneOfType1 *UDPObject `json:"-"`
  oneOfType0 *TCPObject `json:"-"`
  oneOfType2 *ICMPObject `json:"-"`
}

func NewOneOfRoutingPolicyMatchConditionProtocolParameters() *OneOfRoutingPolicyMatchConditionProtocolParameters {
  p := new(OneOfRoutingPolicyMatchConditionProtocolParameters)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfRoutingPolicyMatchConditionProtocolParameters) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfRoutingPolicyMatchConditionProtocolParameters is nil"))
  }
  switch v.(type) {
    case ProtocolNumberObject:
      if nil == p.oneOfType3 {p.oneOfType3 = new(ProtocolNumberObject)}
      *p.oneOfType3 = v.(ProtocolNumberObject)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType3.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType3.ObjectType_
    case UDPObject:
      if nil == p.oneOfType1 {p.oneOfType1 = new(UDPObject)}
      *p.oneOfType1 = v.(UDPObject)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType1.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType1.ObjectType_
    case TCPObject:
      if nil == p.oneOfType0 {p.oneOfType0 = new(TCPObject)}
      *p.oneOfType0 = v.(TCPObject)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    case ICMPObject:
      if nil == p.oneOfType2 {p.oneOfType2 = new(ICMPObject)}
      *p.oneOfType2 = v.(ICMPObject)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType2.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType2.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfRoutingPolicyMatchConditionProtocolParameters) GetValue() interface{} {
  if p.oneOfType3 != nil && *p.oneOfType3.ObjectType_ == *p.Discriminator {
    return *p.oneOfType3
  }
  if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
    return *p.oneOfType1
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType2 != nil && *p.oneOfType2.ObjectType_ == *p.Discriminator {
    return *p.oneOfType2
  }
  return nil
}

func (p *OneOfRoutingPolicyMatchConditionProtocolParameters) UnmarshalJSON(b []byte) error {
  vOneOfType3 := new(ProtocolNumberObject)
  if err := json.Unmarshal(b, vOneOfType3); err == nil {
    if "networking.v4.config.ProtocolNumberObject" == *vOneOfType3.ObjectType_ {
          if nil == p.oneOfType3 {p.oneOfType3 = new(ProtocolNumberObject)}
      *p.oneOfType3 = *vOneOfType3
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType3.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType3.ObjectType_
      return nil
    }
    }
  vOneOfType1 := new(UDPObject)
  if err := json.Unmarshal(b, vOneOfType1); err == nil {
    if "networking.v4.config.UDPObject" == *vOneOfType1.ObjectType_ {
          if nil == p.oneOfType1 {p.oneOfType1 = new(UDPObject)}
      *p.oneOfType1 = *vOneOfType1
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType1.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType1.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new(TCPObject)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "networking.v4.config.TCPObject" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(TCPObject)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  vOneOfType2 := new(ICMPObject)
  if err := json.Unmarshal(b, vOneOfType2); err == nil {
    if "networking.v4.config.ICMPObject" == *vOneOfType2.ObjectType_ {
          if nil == p.oneOfType2 {p.oneOfType2 = new(ICMPObject)}
      *p.oneOfType2 = *vOneOfType2
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType2.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType2.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfRoutingPolicyMatchConditionProtocolParameters"))
}

func (p *OneOfRoutingPolicyMatchConditionProtocolParameters) MarshalJSON() ([]byte, error) {
  if p.oneOfType3 != nil && *p.oneOfType3.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType3)
  }
  if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType1)
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType2 != nil && *p.oneOfType2.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType2)
  }
  return nil, errors.New("No value to marshal for OneOfRoutingPolicyMatchConditionProtocolParameters")
}

type OneOfDirectConnectListApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import3.ErrorResponse `json:"-"`
  oneOfType0 []DirectConnect `json:"-"`
}

func NewOneOfDirectConnectListApiResponseData() *OneOfDirectConnectListApiResponseData {
  p := new(OneOfDirectConnectListApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfDirectConnectListApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfDirectConnectListApiResponseData is nil"))
  }
  switch v.(type) {
    case import3.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = v.(import3.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case []DirectConnect:
      p.oneOfType0 = v.([]DirectConnect)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<networking.v4.config.DirectConnect>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<networking.v4.config.DirectConnect>"
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfDirectConnectListApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if "List<networking.v4.config.DirectConnect>" == *p.Discriminator {
    return p.oneOfType0
  }
  return nil
}

func (p *OneOfDirectConnectListApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import3.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new([]DirectConnect)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "networking.v4.config.DirectConnect" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<networking.v4.config.DirectConnect>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<networking.v4.config.DirectConnect>"
      return nil

    }
  }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDirectConnectListApiResponseData"))
}

func (p *OneOfDirectConnectListApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if "List<networking.v4.config.DirectConnect>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfDirectConnectListApiResponseData")
}

type OneOfRouteTableApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import3.ErrorResponse `json:"-"`
  oneOfType0 *RouteTable `json:"-"`
}

func NewOneOfRouteTableApiResponseData() *OneOfRouteTableApiResponseData {
  p := new(OneOfRouteTableApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfRouteTableApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfRouteTableApiResponseData is nil"))
  }
  switch v.(type) {
    case import3.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = v.(import3.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case RouteTable:
      if nil == p.oneOfType0 {p.oneOfType0 = new(RouteTable)}
      *p.oneOfType0 = v.(RouteTable)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfRouteTableApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  return nil
}

func (p *OneOfRouteTableApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import3.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new(RouteTable)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "networking.v4.config.RouteTable" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(RouteTable)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfRouteTableApiResponseData"))
}

func (p *OneOfRouteTableApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfRouteTableApiResponseData")
}

type OneOfVpnConnectionApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import3.ErrorResponse `json:"-"`
  oneOfType0 *VpnConnection `json:"-"`
}

func NewOneOfVpnConnectionApiResponseData() *OneOfVpnConnectionApiResponseData {
  p := new(OneOfVpnConnectionApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfVpnConnectionApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfVpnConnectionApiResponseData is nil"))
  }
  switch v.(type) {
    case import3.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = v.(import3.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case VpnConnection:
      if nil == p.oneOfType0 {p.oneOfType0 = new(VpnConnection)}
      *p.oneOfType0 = v.(VpnConnection)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfVpnConnectionApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  return nil
}

func (p *OneOfVpnConnectionApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import3.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new(VpnConnection)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "networking.v4.config.VpnConnection" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(VpnConnection)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfVpnConnectionApiResponseData"))
}

func (p *OneOfVpnConnectionApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfVpnConnectionApiResponseData")
}

type OneOfVpnConnectionListApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import3.ErrorResponse `json:"-"`
  oneOfType0 []VpnConnection `json:"-"`
}

func NewOneOfVpnConnectionListApiResponseData() *OneOfVpnConnectionListApiResponseData {
  p := new(OneOfVpnConnectionListApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfVpnConnectionListApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfVpnConnectionListApiResponseData is nil"))
  }
  switch v.(type) {
    case import3.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = v.(import3.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case []VpnConnection:
      p.oneOfType0 = v.([]VpnConnection)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<networking.v4.config.VpnConnection>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<networking.v4.config.VpnConnection>"
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfVpnConnectionListApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if "List<networking.v4.config.VpnConnection>" == *p.Discriminator {
    return p.oneOfType0
  }
  return nil
}

func (p *OneOfVpnConnectionListApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import3.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new([]VpnConnection)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "networking.v4.config.VpnConnection" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<networking.v4.config.VpnConnection>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<networking.v4.config.VpnConnection>"
      return nil

    }
  }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfVpnConnectionListApiResponseData"))
}

func (p *OneOfVpnConnectionListApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if "List<networking.v4.config.VpnConnection>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfVpnConnectionListApiResponseData")
}

type OneOfDirectConnectServiceProviderListApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import3.ErrorResponse `json:"-"`
  oneOfType0 []DirectConnectServiceProvider `json:"-"`
}

func NewOneOfDirectConnectServiceProviderListApiResponseData() *OneOfDirectConnectServiceProviderListApiResponseData {
  p := new(OneOfDirectConnectServiceProviderListApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfDirectConnectServiceProviderListApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfDirectConnectServiceProviderListApiResponseData is nil"))
  }
  switch v.(type) {
    case import3.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = v.(import3.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case []DirectConnectServiceProvider:
      p.oneOfType0 = v.([]DirectConnectServiceProvider)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<networking.v4.config.DirectConnectServiceProvider>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<networking.v4.config.DirectConnectServiceProvider>"
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfDirectConnectServiceProviderListApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if "List<networking.v4.config.DirectConnectServiceProvider>" == *p.Discriminator {
    return p.oneOfType0
  }
  return nil
}

func (p *OneOfDirectConnectServiceProviderListApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import3.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new([]DirectConnectServiceProvider)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "networking.v4.config.DirectConnectServiceProvider" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<networking.v4.config.DirectConnectServiceProvider>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<networking.v4.config.DirectConnectServiceProvider>"
      return nil

    }
  }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDirectConnectServiceProviderListApiResponseData"))
}

func (p *OneOfDirectConnectServiceProviderListApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if "List<networking.v4.config.DirectConnectServiceProvider>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfDirectConnectServiceProviderListApiResponseData")
}

type OneOfLayer2StretchRelatedEntitiesApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *Layer2StretchRelatedEntities `json:"-"`
  oneOfType400 *import3.ErrorResponse `json:"-"`
}

func NewOneOfLayer2StretchRelatedEntitiesApiResponseData() *OneOfLayer2StretchRelatedEntitiesApiResponseData {
  p := new(OneOfLayer2StretchRelatedEntitiesApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfLayer2StretchRelatedEntitiesApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfLayer2StretchRelatedEntitiesApiResponseData is nil"))
  }
  switch v.(type) {
    case Layer2StretchRelatedEntities:
      if nil == p.oneOfType0 {p.oneOfType0 = new(Layer2StretchRelatedEntities)}
      *p.oneOfType0 = v.(Layer2StretchRelatedEntities)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    case import3.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = v.(import3.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfLayer2StretchRelatedEntitiesApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfLayer2StretchRelatedEntitiesApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new(Layer2StretchRelatedEntities)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "networking.v4.config.Layer2StretchRelatedEntities" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(Layer2StretchRelatedEntities)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  vOneOfType400 := new(import3.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfLayer2StretchRelatedEntitiesApiResponseData"))
}

func (p *OneOfLayer2StretchRelatedEntitiesApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfLayer2StretchRelatedEntitiesApiResponseData")
}

type OneOfRoutingPolicyMatchConditionSource struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType1 *IPSubnetObject `json:"-"`
  oneOfType0 *AddressTypeObject `json:"-"`
}

func NewOneOfRoutingPolicyMatchConditionSource() *OneOfRoutingPolicyMatchConditionSource {
  p := new(OneOfRoutingPolicyMatchConditionSource)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfRoutingPolicyMatchConditionSource) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfRoutingPolicyMatchConditionSource is nil"))
  }
  switch v.(type) {
    case IPSubnetObject:
      if nil == p.oneOfType1 {p.oneOfType1 = new(IPSubnetObject)}
      *p.oneOfType1 = v.(IPSubnetObject)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType1.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType1.ObjectType_
    case AddressTypeObject:
      if nil == p.oneOfType0 {p.oneOfType0 = new(AddressTypeObject)}
      *p.oneOfType0 = v.(AddressTypeObject)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfRoutingPolicyMatchConditionSource) GetValue() interface{} {
  if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
    return *p.oneOfType1
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  return nil
}

func (p *OneOfRoutingPolicyMatchConditionSource) UnmarshalJSON(b []byte) error {
  vOneOfType1 := new(IPSubnetObject)
  if err := json.Unmarshal(b, vOneOfType1); err == nil {
    if "networking.v4.config.IPSubnetObject" == *vOneOfType1.ObjectType_ {
          if nil == p.oneOfType1 {p.oneOfType1 = new(IPSubnetObject)}
      *p.oneOfType1 = *vOneOfType1
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType1.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType1.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new(AddressTypeObject)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "networking.v4.config.AddressTypeObject" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(AddressTypeObject)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfRoutingPolicyMatchConditionSource"))
}

func (p *OneOfRoutingPolicyMatchConditionSource) MarshalJSON() ([]byte, error) {
  if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType1)
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfRoutingPolicyMatchConditionSource")
}

type OneOfBgpSessionApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import3.ErrorResponse `json:"-"`
  oneOfType0 *BgpSession `json:"-"`
}

func NewOneOfBgpSessionApiResponseData() *OneOfBgpSessionApiResponseData {
  p := new(OneOfBgpSessionApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfBgpSessionApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfBgpSessionApiResponseData is nil"))
  }
  switch v.(type) {
    case import3.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = v.(import3.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case BgpSession:
      if nil == p.oneOfType0 {p.oneOfType0 = new(BgpSession)}
      *p.oneOfType0 = v.(BgpSession)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfBgpSessionApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  return nil
}

func (p *OneOfBgpSessionApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import3.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new(BgpSession)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "networking.v4.config.BgpSession" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(BgpSession)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfBgpSessionApiResponseData"))
}

func (p *OneOfBgpSessionApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfBgpSessionApiResponseData")
}

type OneOfFlowGatewayApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType1002 []import1.Message `json:"-"`
  oneOfType1003 *FlowGateway `json:"-"`
}

func NewOneOfFlowGatewayApiResponseData() *OneOfFlowGatewayApiResponseData {
  p := new(OneOfFlowGatewayApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfFlowGatewayApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfFlowGatewayApiResponseData is nil"))
  }
  switch v.(type) {
    case []import1.Message:
      p.oneOfType1002 = v.([]import1.Message)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<common.v1.config.Message>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<common.v1.config.Message>"
    case FlowGateway:
      if nil == p.oneOfType1003 {p.oneOfType1003 = new(FlowGateway)}
      *p.oneOfType1003 = v.(FlowGateway)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType1003.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType1003.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfFlowGatewayApiResponseData) GetValue() interface{} {
  if "List<common.v1.config.Message>" == *p.Discriminator {
    return p.oneOfType1002
  }
  if p.oneOfType1003 != nil && *p.oneOfType1003.ObjectType_ == *p.Discriminator {
    return *p.oneOfType1003
  }
  return nil
}

func (p *OneOfFlowGatewayApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType1002 := new([]import1.Message)
  if err := json.Unmarshal(b, vOneOfType1002); err == nil {
    if len(*vOneOfType1002) == 0 || "common.v1.config.Message" == *((*vOneOfType1002)[0].ObjectType_) {
      p.oneOfType1002 = *vOneOfType1002
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<common.v1.config.Message>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<common.v1.config.Message>"
      return nil

    }
  }
  vOneOfType1003 := new(FlowGateway)
  if err := json.Unmarshal(b, vOneOfType1003); err == nil {
    if "networking.v4.config.FlowGateway" == *vOneOfType1003.ObjectType_ {
          if nil == p.oneOfType1003 {p.oneOfType1003 = new(FlowGateway)}
      *p.oneOfType1003 = *vOneOfType1003
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType1003.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType1003.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfFlowGatewayApiResponseData"))
}

func (p *OneOfFlowGatewayApiResponseData) MarshalJSON() ([]byte, error) {
  if "List<common.v1.config.Message>" == *p.Discriminator {
    return json.Marshal(p.oneOfType1002)
  }
  if p.oneOfType1003 != nil && *p.oneOfType1003.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType1003)
  }
  return nil, errors.New("No value to marshal for OneOfFlowGatewayApiResponseData")
}

type OneOfNetworkGatewayServices struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType1 *RemoteNetworkServices `json:"-"`
  oneOfType0 *LocalNetworkServices `json:"-"`
}

func NewOneOfNetworkGatewayServices() *OneOfNetworkGatewayServices {
  p := new(OneOfNetworkGatewayServices)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfNetworkGatewayServices) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfNetworkGatewayServices is nil"))
  }
  switch v.(type) {
    case RemoteNetworkServices:
      if nil == p.oneOfType1 {p.oneOfType1 = new(RemoteNetworkServices)}
      *p.oneOfType1 = v.(RemoteNetworkServices)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType1.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType1.ObjectType_
    case LocalNetworkServices:
      if nil == p.oneOfType0 {p.oneOfType0 = new(LocalNetworkServices)}
      *p.oneOfType0 = v.(LocalNetworkServices)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfNetworkGatewayServices) GetValue() interface{} {
  if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
    return *p.oneOfType1
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  return nil
}

func (p *OneOfNetworkGatewayServices) UnmarshalJSON(b []byte) error {
  vOneOfType1 := new(RemoteNetworkServices)
  if err := json.Unmarshal(b, vOneOfType1); err == nil {
    if "networking.v4.config.RemoteNetworkServices" == *vOneOfType1.ObjectType_ {
          if nil == p.oneOfType1 {p.oneOfType1 = new(RemoteNetworkServices)}
      *p.oneOfType1 = *vOneOfType1
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType1.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType1.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new(LocalNetworkServices)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "networking.v4.config.LocalNetworkServices" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(LocalNetworkServices)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfNetworkGatewayServices"))
}

func (p *OneOfNetworkGatewayServices) MarshalJSON() ([]byte, error) {
  if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType1)
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfNetworkGatewayServices")
}

type OneOfFlowGatewayKeepAliveResponseApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType1002 []import1.Message `json:"-"`
  oneOfType1003 *FlowGatewayKeepAliveResponse `json:"-"`
}

func NewOneOfFlowGatewayKeepAliveResponseApiResponseData() *OneOfFlowGatewayKeepAliveResponseApiResponseData {
  p := new(OneOfFlowGatewayKeepAliveResponseApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfFlowGatewayKeepAliveResponseApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfFlowGatewayKeepAliveResponseApiResponseData is nil"))
  }
  switch v.(type) {
    case []import1.Message:
      p.oneOfType1002 = v.([]import1.Message)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<common.v1.config.Message>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<common.v1.config.Message>"
    case FlowGatewayKeepAliveResponse:
      if nil == p.oneOfType1003 {p.oneOfType1003 = new(FlowGatewayKeepAliveResponse)}
      *p.oneOfType1003 = v.(FlowGatewayKeepAliveResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType1003.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType1003.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfFlowGatewayKeepAliveResponseApiResponseData) GetValue() interface{} {
  if "List<common.v1.config.Message>" == *p.Discriminator {
    return p.oneOfType1002
  }
  if p.oneOfType1003 != nil && *p.oneOfType1003.ObjectType_ == *p.Discriminator {
    return *p.oneOfType1003
  }
  return nil
}

func (p *OneOfFlowGatewayKeepAliveResponseApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType1002 := new([]import1.Message)
  if err := json.Unmarshal(b, vOneOfType1002); err == nil {
    if len(*vOneOfType1002) == 0 || "common.v1.config.Message" == *((*vOneOfType1002)[0].ObjectType_) {
      p.oneOfType1002 = *vOneOfType1002
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<common.v1.config.Message>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<common.v1.config.Message>"
      return nil

    }
  }
  vOneOfType1003 := new(FlowGatewayKeepAliveResponse)
  if err := json.Unmarshal(b, vOneOfType1003); err == nil {
    if "networking.v4.config.FlowGatewayKeepAliveResponse" == *vOneOfType1003.ObjectType_ {
          if nil == p.oneOfType1003 {p.oneOfType1003 = new(FlowGatewayKeepAliveResponse)}
      *p.oneOfType1003 = *vOneOfType1003
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType1003.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType1003.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfFlowGatewayKeepAliveResponseApiResponseData"))
}

func (p *OneOfFlowGatewayKeepAliveResponseApiResponseData) MarshalJSON() ([]byte, error) {
  if "List<common.v1.config.Message>" == *p.Discriminator {
    return json.Marshal(p.oneOfType1002)
  }
  if p.oneOfType1003 != nil && *p.oneOfType1003.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType1003)
  }
  return nil, errors.New("No value to marshal for OneOfFlowGatewayKeepAliveResponseApiResponseData")
}

type OneOfTaskReferenceApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import3.ErrorResponse `json:"-"`
  oneOfType0 *import4.TaskReference `json:"-"`
}

func NewOneOfTaskReferenceApiResponseData() *OneOfTaskReferenceApiResponseData {
  p := new(OneOfTaskReferenceApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfTaskReferenceApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfTaskReferenceApiResponseData is nil"))
  }
  switch v.(type) {
    case import3.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = v.(import3.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case import4.TaskReference:
      if nil == p.oneOfType0 {p.oneOfType0 = new(import4.TaskReference)}
      *p.oneOfType0 = v.(import4.TaskReference)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfTaskReferenceApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  return nil
}

func (p *OneOfTaskReferenceApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import3.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new(import4.TaskReference)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(import4.TaskReference)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfTaskReferenceApiResponseData"))
}

func (p *OneOfTaskReferenceApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfTaskReferenceApiResponseData")
}

