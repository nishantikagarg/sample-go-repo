/*
 * Generated file models/microseg/v4/policies/policies_model.go.
 *
 * Product version: 16.2.0-SNAPSHOT
 *
 * Part of the Microseg REST APIs
 *
 * (c) 2022 Nutanix Inc.  All rights reserved
 *
 */

/*
  
*/
package policies
import (
  "bytes"
  import2 "github.com/nishantikagarg/sample-go-repo/flow_go_sdk/v16/models/common/v1/config"
  import1 "github.com/nishantikagarg/sample-go-repo/flow_go_sdk/v16/models/common/v1/response"
  "encoding/json"
  "errors"
  "fmt"
  import3 "github.com/nishantikagarg/sample-go-repo/flow_go_sdk/v16/models/microseg/v4/error"
  import4 "github.com/nishantikagarg/sample-go-repo/flow_go_sdk/v16/models/prism/v4/config"
  "time"
)

/**
A reference to an address group.
*/
type AddressGroupRefSpec struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  AddressGroupName *string `json:"addressGroupName,omitempty"`
  
  ExtId *string `json:"extId"`
}

func NewAddressGroupRefSpec() *AddressGroupRefSpec {
  p := new(AddressGroupRefSpec)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "microseg.v4.policies.AddressGroupRefSpec"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.policies.AddressGroupRefSpec"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}





type AllowType int

const(
  ALLOWTYPE_UNKNOWN AllowType = 0
  ALLOWTYPE_REDACTED AllowType = 1
  ALLOWTYPE_ALL AllowType = 2
)

// returns the name of the enum given an ordinal number
func (e *AllowType) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "ALL",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *AllowType) index(name string) AllowType {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "ALL",
  }
  for idx := range names {
    if names[idx] == name {
      return AllowType(idx)
    }
  }
  return ALLOWTYPE_UNKNOWN
}

func (e *AllowType) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for AllowType:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *AllowType) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e AllowType) Ref() *AllowType {
  return &e
}


/**
Specifies allow types - currently, all traffic is the only option.
*/
type AllowTypesSpec struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  AllowType *AllowType `json:"allowType"`
}

func NewAllowTypesSpec() *AllowTypesSpec {
  p := new(AllowTypesSpec)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "microseg.v4.policies.AllowTypesSpec"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.policies.AllowTypesSpec"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
A set of one or more categories that can be used for securing entities inside of a rule or allowing traffic as inbound or outbound traffic; membership in a secured entity group is defined as those entities which are a member of ALL categories in the secured entity group.
*/
type CategoryGroupSpec struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  Categories []CategoryValueRef `json:"categories"`
}

func NewCategoryGroupSpec() *CategoryGroupSpec {
  p := new(CategoryGroupSpec)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "microseg.v4.policies.CategoryGroupSpec"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.policies.CategoryGroupSpec"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
A reference to a category value
*/
type CategoryValueRef struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  CategoryPath *string `json:"categoryPath,omitempty"`
  
  ExtId *string `json:"extId"`
}

func NewCategoryValueRef() *CategoryValueRef {
  p := new(CategoryValueRef)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "microseg.v4.policies.CategoryValueRef"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.policies.CategoryValueRef"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Specifies if a rule is an inbound (IN), or outbound (OUT) rule.
*/
type Direction int

const(
  DIRECTION_UNKNOWN Direction = 0
  DIRECTION_REDACTED Direction = 1
  DIRECTION_IN Direction = 2
  DIRECTION_OUT Direction = 3
)

// returns the name of the enum given an ordinal number
func (e *Direction) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "IN",
    "OUT",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *Direction) index(name string) Direction {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "IN",
    "OUT",
  }
  for idx := range names {
    if names[idx] == name {
      return Direction(idx)
    }
  }
  return DIRECTION_UNKNOWN
}

func (e *Direction) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for Direction:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *Direction) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e Direction) Ref() *Direction {
  return &e
}


/**
A summary of counts for a specific field.
*/
type FieldCountSummary struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  Counts map[string]int `json:"counts,omitempty"`
  
  FieldName *string `json:"fieldName,omitempty"`
}

func NewFieldCountSummary() *FieldCountSummary {
  p := new(FieldCountSummary)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "microseg.v4.policies.FieldCountSummary"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.policies.FieldCountSummary"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}





type FileWrapper struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  File *string `json:"file,omitempty"`
}

func NewFileWrapper() *FileWrapper {
  p := new(FileWrapper)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "microseg.v4.policies.FileWrapper"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.policies.FileWrapper"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
A ICMP type and code value combination.
*/
type ICMPTypeCodeSpec struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  Code *int `json:"code,omitempty"`
  
  Type *int `json:"type,omitempty"`
}

func NewICMPTypeCodeSpec() *ICMPTypeCodeSpec {
  p := new(ICMPTypeCodeSpec)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "microseg.v4.policies.ICMPTypeCodeSpec"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.policies.ICMPTypeCodeSpec"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}





type IntraEntityGroupRuleAction int

const(
  INTRAENTITYGROUPRULEACTION_UNKNOWN IntraEntityGroupRuleAction = 0
  INTRAENTITYGROUPRULEACTION_REDACTED IntraEntityGroupRuleAction = 1
  INTRAENTITYGROUPRULEACTION_ALLOW IntraEntityGroupRuleAction = 2
  INTRAENTITYGROUPRULEACTION_DENY IntraEntityGroupRuleAction = 3
)

// returns the name of the enum given an ordinal number
func (e *IntraEntityGroupRuleAction) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "ALLOW",
    "DENY",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *IntraEntityGroupRuleAction) index(name string) IntraEntityGroupRuleAction {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "ALLOW",
    "DENY",
  }
  for idx := range names {
    if names[idx] == name {
      return IntraEntityGroupRuleAction(idx)
    }
  }
  return INTRAENTITYGROUPRULEACTION_UNKNOWN
}

func (e *IntraEntityGroupRuleAction) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for IntraEntityGroupRuleAction:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *IntraEntityGroupRuleAction) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e IntraEntityGroupRuleAction) Ref() *IntraEntityGroupRuleAction {
  return &e
}


/**
Contains a list of Flow 1.0 network security rules that need to be migrated.
*/
type MigrationDryRunInput struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  A list of Flow 1.0 network security rule names.
  */
  RuleNames []string `json:"ruleNames"`
}

func NewMigrationDryRunInput() *MigrationDryRunInput {
  p := new(MigrationDryRunInput)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "microseg.v4.policies.MigrationDryRunInput"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.policies.MigrationDryRunInput"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
A reference to a network function chain.
*/
type NFRef struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  ExtId *string `json:"extId"`
  
  NetworkFunctionChainName *string `json:"networkFunctionChainName,omitempty"`
}

func NewNFRef() *NFRef {
  p := new(NFRef)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "microseg.v4.policies.NFRef"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.policies.NFRef"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
A rule for specifying allowed traffic for an application.
*/
type NetworkPolicyApplicationRuleSpec struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  Direction *Direction `json:"direction"`
  
  NetworkFunctionChain *NFRef `json:"networkFunctionChain,omitempty"`
  
  SecuredGroup *CategoryGroupSpec `json:"securedGroup"`
  /**
  A list of services that are allowed for the listed allowed sources or destinations.
  */
  Services []Service `json:"services,omitempty"`
  
  SourceOrDest *SourceOrDest `json:"sourceOrDest,omitempty"`
}

func NewNetworkPolicyApplicationRuleSpec() *NetworkPolicyApplicationRuleSpec {
  p := new(NetworkPolicyApplicationRuleSpec)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "microseg.v4.policies.NetworkPolicyApplicationRuleSpec"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.policies.NetworkPolicyApplicationRuleSpec"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
A rule for specifying allowed traffic inside of a secured entity group.
*/
type NetworkPolicyIntraEntityGroupRuleSpec struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  SecuredGroup *CategoryGroupSpec `json:"securedGroup"`
  
  SecuredGroupAction *IntraEntityGroupRuleAction `json:"securedGroupAction,omitempty"`
}

func NewNetworkPolicyIntraEntityGroupRuleSpec() *NetworkPolicyIntraEntityGroupRuleSpec {
  p := new(NetworkPolicyIntraEntityGroupRuleSpec)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "microseg.v4.policies.NetworkPolicyIntraEntityGroupRuleSpec"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.policies.NetworkPolicyIntraEntityGroupRuleSpec"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Policy-wide options for a security policy.
*/
type NetworkPolicyOptions struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  AllowIPv6Traffic *bool `json:"allowIPv6Traffic,omitempty"`
  
  EnableHitlog *bool `json:"enableHitlog,omitempty"`
}

func NewNetworkPolicyOptions() *NetworkPolicyOptions {
  p := new(NetworkPolicyOptions)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "microseg.v4.policies.NetworkPolicyOptions"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.policies.NetworkPolicyOptions"}
  p.UnknownFields_ = map[string]interface{}{}

  p.AllowIPv6Traffic = new(bool)
  *p.AllowIPv6Traffic = false
  p.EnableHitlog = new(bool)
  *p.EnableHitlog = false


  return p
}




/**
A rule for specifying that two environments should be isolated from each other. Specify both 'firstIsolationGroup' and 'secondIsolationGroup'.
*/
type NetworkPolicyTwoEnvIsolationRuleSpec struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  FirstIsolationGroup *CategoryGroupSpec `json:"firstIsolationGroup"`
  
  SecondIsolationGroup *CategoryGroupSpec `json:"secondIsolationGroup"`
}

func NewNetworkPolicyTwoEnvIsolationRuleSpec() *NetworkPolicyTwoEnvIsolationRuleSpec {
  p := new(NetworkPolicyTwoEnvIsolationRuleSpec)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "microseg.v4.policies.NetworkPolicyTwoEnvIsolationRuleSpec"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.policies.NetworkPolicyTwoEnvIsolationRuleSpec"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}





type NetworkSecurityPolicy struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  A user defined annotation for a policy.
  */
  Description *string `json:"description,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import1.ApiLink `json:"links,omitempty"`
  /**
  A short identifier for a policy.
  */
  Name *string `json:"name"`
  
  Options *NetworkPolicyOptions `json:"options,omitempty"`
  /**
  A list of rules that form a policy. For isolation policies, use isolation rules; for application or quarantine policies, use application rules.
  */
  Rules []Rule `json:"rules,omitempty"`
  
  State *PolicyState `json:"state,omitempty"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
  
  Type *PolicyType `json:"type"`
  
  Vpc *VPCRef `json:"vpc,omitempty"`
}

func NewNetworkSecurityPolicy() *NetworkSecurityPolicy {
  p := new(NetworkSecurityPolicy)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "microseg.v4.policies.NetworkSecurityPolicy"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.policies.NetworkSecurityPolicy"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST Response for all response codes in api path /microseg/v4.0.a1/policies/export Get operation
*/
type NetworkSecurityPolicyExportResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfNetworkSecurityPolicyExportResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewNetworkSecurityPolicyExportResponse() *NetworkSecurityPolicyExportResponse {
  p := new(NetworkSecurityPolicyExportResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "microseg.v4.policies.NetworkSecurityPolicyExportResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.policies.NetworkSecurityPolicyExportResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *NetworkSecurityPolicyExportResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *NetworkSecurityPolicyExportResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfNetworkSecurityPolicyExportResponseData()
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
REST Response for all response codes in api path /microseg/v4.0.a1/policies/counts Get operation
*/
type NetworkSecurityPolicyGetCountsResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfNetworkSecurityPolicyGetCountsResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewNetworkSecurityPolicyGetCountsResponse() *NetworkSecurityPolicyGetCountsResponse {
  p := new(NetworkSecurityPolicyGetCountsResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "microseg.v4.policies.NetworkSecurityPolicyGetCountsResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.policies.NetworkSecurityPolicyGetCountsResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *NetworkSecurityPolicyGetCountsResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *NetworkSecurityPolicyGetCountsResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfNetworkSecurityPolicyGetCountsResponseData()
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
REST Response for all response codes in api path /microseg/v4.0.a1/policies/$actions/migrate-config/dry-run Post operation
*/
type NetworkSecurityPolicyGetListResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfNetworkSecurityPolicyGetListResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewNetworkSecurityPolicyGetListResponse() *NetworkSecurityPolicyGetListResponse {
  p := new(NetworkSecurityPolicyGetListResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "microseg.v4.policies.NetworkSecurityPolicyGetListResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.policies.NetworkSecurityPolicyGetListResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *NetworkSecurityPolicyGetListResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *NetworkSecurityPolicyGetListResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfNetworkSecurityPolicyGetListResponseData()
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
REST Response for all response codes in api path /microseg/v4.0.a1/policies/{extId} Get operation
*/
type NetworkSecurityPolicyGetResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfNetworkSecurityPolicyGetResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewNetworkSecurityPolicyGetResponse() *NetworkSecurityPolicyGetResponse {
  p := new(NetworkSecurityPolicyGetResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "microseg.v4.policies.NetworkSecurityPolicyGetResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.policies.NetworkSecurityPolicyGetResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *NetworkSecurityPolicyGetResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *NetworkSecurityPolicyGetResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfNetworkSecurityPolicyGetResponseData()
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
Network security rule import response data.
*/
type NetworkSecurityPolicyImportEntity struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Name of the entity.
  */
  EntityName *string `json:"entityName"`
  
  EntityType *NetworkSecurityPolicyImportEntityType `json:"entityType"`
  
  EntityUpdateType *NetworkSecurityPolicyImportEntityUpdateType `json:"entityUpdateType"`
}

func NewNetworkSecurityPolicyImportEntity() *NetworkSecurityPolicyImportEntity {
  p := new(NetworkSecurityPolicyImportEntity)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "microseg.v4.policies.NetworkSecurityPolicyImportEntity"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.policies.NetworkSecurityPolicyImportEntity"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Type of the entity.
*/
type NetworkSecurityPolicyImportEntityType int

const(
  NETWORKSECURITYPOLICYIMPORTENTITYTYPE_UNKNOWN NetworkSecurityPolicyImportEntityType = 0
  NETWORKSECURITYPOLICYIMPORTENTITYTYPE_REDACTED NetworkSecurityPolicyImportEntityType = 1
  NETWORKSECURITYPOLICYIMPORTENTITYTYPE_CATEGORY NetworkSecurityPolicyImportEntityType = 2
  NETWORKSECURITYPOLICYIMPORTENTITYTYPE_NETWORK_FUNCTION_CHAIN NetworkSecurityPolicyImportEntityType = 3
  NETWORKSECURITYPOLICYIMPORTENTITYTYPE_POLICY NetworkSecurityPolicyImportEntityType = 4
  NETWORKSECURITYPOLICYIMPORTENTITYTYPE_ADDRESS_GROUP NetworkSecurityPolicyImportEntityType = 5
  NETWORKSECURITYPOLICYIMPORTENTITYTYPE_SERVICE_GROUP NetworkSecurityPolicyImportEntityType = 6
)

// returns the name of the enum given an ordinal number
func (e *NetworkSecurityPolicyImportEntityType) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "CATEGORY",
    "NETWORK_FUNCTION_CHAIN",
    "POLICY",
    "ADDRESS_GROUP",
    "SERVICE_GROUP",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *NetworkSecurityPolicyImportEntityType) index(name string) NetworkSecurityPolicyImportEntityType {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "CATEGORY",
    "NETWORK_FUNCTION_CHAIN",
    "POLICY",
    "ADDRESS_GROUP",
    "SERVICE_GROUP",
  }
  for idx := range names {
    if names[idx] == name {
      return NetworkSecurityPolicyImportEntityType(idx)
    }
  }
  return NETWORKSECURITYPOLICYIMPORTENTITYTYPE_UNKNOWN
}

func (e *NetworkSecurityPolicyImportEntityType) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for NetworkSecurityPolicyImportEntityType:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *NetworkSecurityPolicyImportEntityType) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e NetworkSecurityPolicyImportEntityType) Ref() *NetworkSecurityPolicyImportEntityType {
  return &e
}


/**
Type of update of the entity.
*/
type NetworkSecurityPolicyImportEntityUpdateType int

const(
  NETWORKSECURITYPOLICYIMPORTENTITYUPDATETYPE_UNKNOWN NetworkSecurityPolicyImportEntityUpdateType = 0
  NETWORKSECURITYPOLICYIMPORTENTITYUPDATETYPE_REDACTED NetworkSecurityPolicyImportEntityUpdateType = 1
  NETWORKSECURITYPOLICYIMPORTENTITYUPDATETYPE_ADDED NetworkSecurityPolicyImportEntityUpdateType = 2
  NETWORKSECURITYPOLICYIMPORTENTITYUPDATETYPE_DELETED NetworkSecurityPolicyImportEntityUpdateType = 3
  NETWORKSECURITYPOLICYIMPORTENTITYUPDATETYPE_MODIFIED NetworkSecurityPolicyImportEntityUpdateType = 4
)

// returns the name of the enum given an ordinal number
func (e *NetworkSecurityPolicyImportEntityUpdateType) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "ADDED",
    "DELETED",
    "MODIFIED",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *NetworkSecurityPolicyImportEntityUpdateType) index(name string) NetworkSecurityPolicyImportEntityUpdateType {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "ADDED",
    "DELETED",
    "MODIFIED",
  }
  for idx := range names {
    if names[idx] == name {
      return NetworkSecurityPolicyImportEntityUpdateType(idx)
    }
  }
  return NETWORKSECURITYPOLICYIMPORTENTITYUPDATETYPE_UNKNOWN
}

func (e *NetworkSecurityPolicyImportEntityUpdateType) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for NetworkSecurityPolicyImportEntityUpdateType:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *NetworkSecurityPolicyImportEntityUpdateType) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e NetworkSecurityPolicyImportEntityUpdateType) Ref() *NetworkSecurityPolicyImportEntityUpdateType {
  return &e
}


/**
REST Response for all response codes in api path /microseg/v4.0.a1/policies/$actions/import/dry_run Post operation
*/
type NetworkSecurityPolicyImportResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfNetworkSecurityPolicyImportResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewNetworkSecurityPolicyImportResponse() *NetworkSecurityPolicyImportResponse {
  p := new(NetworkSecurityPolicyImportResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "microseg.v4.policies.NetworkSecurityPolicyImportResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.policies.NetworkSecurityPolicyImportResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *NetworkSecurityPolicyImportResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *NetworkSecurityPolicyImportResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfNetworkSecurityPolicyImportResponseData()
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
Network security Policy Import response.
*/
type NetworkSecurityPolicyImportSummary struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Entity list is returned for a successful import dryrun.
  */
  EntitySummaries []NetworkSecurityPolicyImportEntity `json:"entitySummaries,omitempty"`
}

func NewNetworkSecurityPolicyImportSummary() *NetworkSecurityPolicyImportSummary {
  p := new(NetworkSecurityPolicyImportSummary)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "microseg.v4.policies.NetworkSecurityPolicyImportSummary"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.policies.NetworkSecurityPolicyImportSummary"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST Response for all response codes in api path /microseg/v4.0.a1/policies/$actions/set-allow-ipv6-traffic Post operation
*/
type NetworkSecurityPolicyTaskResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfNetworkSecurityPolicyTaskResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewNetworkSecurityPolicyTaskResponse() *NetworkSecurityPolicyTaskResponse {
  p := new(NetworkSecurityPolicyTaskResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "microseg.v4.policies.NetworkSecurityPolicyTaskResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.policies.NetworkSecurityPolicyTaskResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *NetworkSecurityPolicyTaskResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *NetworkSecurityPolicyTaskResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfNetworkSecurityPolicyTaskResponseData()
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




type NetworkSecurityPolicyWithMetadata struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  CreationTime *time.Time `json:"creationTime,omitempty"`
  /**
  A user defined annotation for a policy.
  */
  Description *string `json:"description,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  
  LastUpdateTime *time.Time `json:"lastUpdateTime,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import1.ApiLink `json:"links,omitempty"`
  /**
  A short identifier for a policy.
  */
  Name *string `json:"name"`
  
  Options *NetworkPolicyOptions `json:"options,omitempty"`
  /**
  A list of rules that form a policy. For isolation policies, use isolation rules; for application or quarantine policies, use application rules.
  */
  Rules []Rule `json:"rules,omitempty"`
  
  State *PolicyState `json:"state,omitempty"`
  
  SystemDefined *bool `json:"systemDefined,omitempty"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
  
  Type *PolicyType `json:"type"`
  
  Vpc *VPCRef `json:"vpc,omitempty"`
}

func NewNetworkSecurityPolicyWithMetadata() *NetworkSecurityPolicyWithMetadata {
  p := new(NetworkSecurityPolicyWithMetadata)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "microseg.v4.policies.NetworkSecurityPolicyWithMetadata"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.policies.NetworkSecurityPolicyWithMetadata"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
A list of external ids for a set of network security policies.
*/
type PolicyExtIdList struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  A list of external ids for a set of network security policies.
  */
  PolicyExtIds []string `json:"policyExtIds"`
}

func NewPolicyExtIdList() *PolicyExtIdList {
  p := new(PolicyExtIdList)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "microseg.v4.policies.PolicyExtIdList"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.policies.PolicyExtIdList"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}





type PolicyExtIdListWithBooleanValue struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  A list of external ids for a set of network security policies.
  */
  PolicyExtIds []string `json:"policyExtIds"`
  
  Value *bool `json:"value"`
}

func NewPolicyExtIdListWithBooleanValue() *PolicyExtIdListWithBooleanValue {
  p := new(PolicyExtIdListWithBooleanValue)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "microseg.v4.policies.PolicyExtIdListWithBooleanValue"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.policies.PolicyExtIdListWithBooleanValue"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}





type PolicyExtIdListWithStateValue struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  A list of external ids for a set of network security policies.
  */
  PolicyExtIds []string `json:"policyExtIds"`
  
  State *PolicyState `json:"state"`
}

func NewPolicyExtIdListWithStateValue() *PolicyExtIdListWithStateValue {
  p := new(PolicyExtIdListWithStateValue)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "microseg.v4.policies.PolicyExtIdListWithStateValue"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.policies.PolicyExtIdListWithStateValue"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Whether the policy is applied or monitored; can be omitted or set null to save the policy without applying or monitoring it.
*/
type PolicyState int

const(
  POLICYSTATE_UNKNOWN PolicyState = 0
  POLICYSTATE_REDACTED PolicyState = 1
  POLICYSTATE_SAVE PolicyState = 2
  POLICYSTATE_MONITOR PolicyState = 3
  POLICYSTATE_ENFORCE PolicyState = 4
)

// returns the name of the enum given an ordinal number
func (e *PolicyState) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "SAVE",
    "MONITOR",
    "ENFORCE",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *PolicyState) index(name string) PolicyState {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "SAVE",
    "MONITOR",
    "ENFORCE",
  }
  for idx := range names {
    if names[idx] == name {
      return PolicyState(idx)
    }
  }
  return POLICYSTATE_UNKNOWN
}

func (e *PolicyState) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for PolicyState:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *PolicyState) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e PolicyState) Ref() *PolicyState {
  return &e
}


/**
Defines the type of rules that can be used in a policy.
*/
type PolicyType int

const(
  POLICYTYPE_UNKNOWN PolicyType = 0
  POLICYTYPE_REDACTED PolicyType = 1
  POLICYTYPE_QUARANTINE PolicyType = 2
  POLICYTYPE_ISOLATION PolicyType = 3
  POLICYTYPE_APPLICATION PolicyType = 4
)

// returns the name of the enum given an ordinal number
func (e *PolicyType) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "QUARANTINE",
    "ISOLATION",
    "APPLICATION",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *PolicyType) index(name string) PolicyType {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "QUARANTINE",
    "ISOLATION",
    "APPLICATION",
  }
  for idx := range names {
    if names[idx] == name {
      return PolicyType(idx)
    }
  }
  return POLICYTYPE_UNKNOWN
}

func (e *PolicyType) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for PolicyType:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *PolicyType) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e PolicyType) Ref() *PolicyType {
  return &e
}


/**
A port range.
*/
type PortRangeSpec struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  EndPort *int `json:"endPort"`
  
  StartPort *int `json:"startPort"`
}

func NewPortRangeSpec() *PortRangeSpec {
  p := new(PortRangeSpec)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "microseg.v4.policies.PortRangeSpec"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.policies.PortRangeSpec"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}





type Rule struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  A user defined annotation for a rule.
  */
  Description *string `json:"description,omitempty"`
  /**
  Rule UUID; should be used to identify an individual rule uniquely.
  */
  Id *string `json:"id,omitempty"`
  
  SpecItemDiscriminator_ *string `json:"$specItemDiscriminator,omitempty"`
  
  Spec *OneOfRuleSpec `json:"spec"`
  
  Type *RuleType `json:"type"`
}

func NewRule() *Rule {
  p := new(Rule)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "microseg.v4.policies.Rule"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.policies.Rule"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *Rule) GetSpec() interface{} {
  if nil == p.Spec {
    return nil
  }
  return p.Spec.GetValue()
}

func (p *Rule) SetSpec(v interface{}) error {
  if nil == p.Spec {
    p.Spec = NewOneOfRuleSpec()
  }
  e := p.Spec.SetValue(v)
  if nil == e {
    if nil == p.SpecItemDiscriminator_ {
      p.SpecItemDiscriminator_ = new(string)
    }
    *p.SpecItemDiscriminator_ = *p.Spec.Discriminator
  }
  return e
}



/**
The type for a rule - the value chosen here restricts which specification can be chosen.
*/
type RuleType int

const(
  RULETYPE_UNKNOWN RuleType = 0
  RULETYPE_REDACTED RuleType = 1
  RULETYPE_QUARANTINE RuleType = 2
  RULETYPE_TWO_ENV_ISOLATION RuleType = 3
  RULETYPE_APPLICATION RuleType = 4
  RULETYPE_INTRA_GROUP RuleType = 5
)

// returns the name of the enum given an ordinal number
func (e *RuleType) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "QUARANTINE",
    "TWO_ENV_ISOLATION",
    "APPLICATION",
    "INTRA_GROUP",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *RuleType) index(name string) RuleType {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "QUARANTINE",
    "TWO_ENV_ISOLATION",
    "APPLICATION",
    "INTRA_GROUP",
  }
  for idx := range names {
    if names[idx] == name {
      return RuleType(idx)
    }
  }
  return RULETYPE_UNKNOWN
}

func (e *RuleType) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for RuleType:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *RuleType) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e RuleType) Ref() *RuleType {
  return &e
}


/**
Service wrapper for a rule, which specifies what type of traffic is allowed. Pick one of the options provided for 'spec' - a service group reference or a service specification.
*/
type Service struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  SpecItemDiscriminator_ *string `json:"$specItemDiscriminator,omitempty"`
  
  Spec *OneOfServiceSpec `json:"spec"`
}

func NewService() *Service {
  p := new(Service)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "microseg.v4.policies.Service"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.policies.Service"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *Service) GetSpec() interface{} {
  if nil == p.Spec {
    return nil
  }
  return p.Spec.GetValue()
}

func (p *Service) SetSpec(v interface{}) error {
  if nil == p.Spec {
    p.Spec = NewOneOfServiceSpec()
  }
  e := p.Spec.SetValue(v)
  if nil == e {
    if nil == p.SpecItemDiscriminator_ {
      p.SpecItemDiscriminator_ = new(string)
    }
    *p.SpecItemDiscriminator_ = *p.Spec.Discriminator
  }
  return e
}



/**
A reference to a service group.
*/
type ServiceGroupRefSpec struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  ExtId *string `json:"extId"`
  
  ServiceGroupName *string `json:"serviceGroupName,omitempty"`
}

func NewServiceGroupRefSpec() *ServiceGroupRefSpec {
  p := new(ServiceGroupRefSpec)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "microseg.v4.policies.ServiceGroupRefSpec"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.policies.ServiceGroupRefSpec"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
The protocol to allow (or all).
*/
type ServiceProtocol int

const(
  SERVICEPROTOCOL_UNKNOWN ServiceProtocol = 0
  SERVICEPROTOCOL_REDACTED ServiceProtocol = 1
  SERVICEPROTOCOL_ALL ServiceProtocol = 2
  SERVICEPROTOCOL_TCP ServiceProtocol = 3
  SERVICEPROTOCOL_UDP ServiceProtocol = 4
  SERVICEPROTOCOL_ICMP ServiceProtocol = 5
)

// returns the name of the enum given an ordinal number
func (e *ServiceProtocol) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "ALL",
    "TCP",
    "UDP",
    "ICMP",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *ServiceProtocol) index(name string) ServiceProtocol {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "ALL",
    "TCP",
    "UDP",
    "ICMP",
  }
  for idx := range names {
    if names[idx] == name {
      return ServiceProtocol(idx)
    }
  }
  return SERVICEPROTOCOL_UNKNOWN
}

func (e *ServiceProtocol) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for ServiceProtocol:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *ServiceProtocol) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e ServiceProtocol) Ref() *ServiceProtocol {
  return &e
}



type ServiceProtocolDetails struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  SpecItemDiscriminator_ *string `json:"$specItemDiscriminator,omitempty"`
  
  Spec *OneOfServiceProtocolDetailsSpec `json:"spec,omitempty"`
}

func NewServiceProtocolDetails() *ServiceProtocolDetails {
  p := new(ServiceProtocolDetails)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "microseg.v4.policies.ServiceProtocolDetails"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.policies.ServiceProtocolDetails"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *ServiceProtocolDetails) GetSpec() interface{} {
  if nil == p.Spec {
    return nil
  }
  return p.Spec.GetValue()
}

func (p *ServiceProtocolDetails) SetSpec(v interface{}) error {
  if nil == p.Spec {
    p.Spec = NewOneOfServiceProtocolDetailsSpec()
  }
  e := p.Spec.SetValue(v)
  if nil == e {
    if nil == p.SpecItemDiscriminator_ {
      p.SpecItemDiscriminator_ = new(string)
    }
    *p.SpecItemDiscriminator_ = *p.Spec.Discriminator
  }
  return e
}



/**
A specification for a service.
*/
type ServiceSpec struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  Protocol *ServiceProtocol `json:"protocol"`
  /**
  For tcp or udp, a list of port ranges, for icmp a type-code combination list.
  */
  ProtocolDetails []ServiceProtocolDetails `json:"protocolDetails,omitempty"`
}

func NewServiceSpec() *ServiceSpec {
  p := new(ServiceSpec)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "microseg.v4.policies.ServiceSpec"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.policies.ServiceSpec"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Source or destination specification for a rule. Pick one of the options provided for 'spec' - a blanket allow or deny, or a traffic whitelist via categories, subnets, or address groups.
*/
type SourceOrDest struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  SpecItemDiscriminator_ *string `json:"$specItemDiscriminator,omitempty"`
  
  Spec *OneOfSourceOrDestSpec `json:"spec"`
}

func NewSourceOrDest() *SourceOrDest {
  p := new(SourceOrDest)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "microseg.v4.policies.SourceOrDest"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.policies.SourceOrDest"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *SourceOrDest) GetSpec() interface{} {
  if nil == p.Spec {
    return nil
  }
  return p.Spec.GetValue()
}

func (p *SourceOrDest) SetSpec(v interface{}) error {
  if nil == p.Spec {
    p.Spec = NewOneOfSourceOrDestSpec()
  }
  e := p.Spec.SetValue(v)
  if nil == e {
    if nil == p.SpecItemDiscriminator_ {
      p.SpecItemDiscriminator_ = new(string)
    }
    *p.SpecItemDiscriminator_ = *p.Spec.Discriminator
  }
  return e
}



/**
A subnet in the form x.x.x.x/xx; /32s are permitted.
*/
type SubnetSpec struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  Subnet *string `json:"subnet,omitempty"`
}

func NewSubnetSpec() *SubnetSpec {
  p := new(SubnetSpec)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "microseg.v4.policies.SubnetSpec"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.policies.SubnetSpec"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
An optional reference to a VPC.
*/
type VPCRef struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  VPCName *string `json:"VPCName,omitempty"`
  
  ExtId *string `json:"extId"`
}

func NewVPCRef() *VPCRef {
  p := new(VPCRef)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "microseg.v4.policies.VPCRef"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.policies.VPCRef"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}



type OneOfNetworkSecurityPolicyImportResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType1 []import2.Message `json:"-"`
  oneOfType400 *import3.ErrorResponse `json:"-"`
  oneOfType0 *NetworkSecurityPolicyImportSummary `json:"-"`
}

func NewOneOfNetworkSecurityPolicyImportResponseData() *OneOfNetworkSecurityPolicyImportResponseData {
  p := new(OneOfNetworkSecurityPolicyImportResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfNetworkSecurityPolicyImportResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfNetworkSecurityPolicyImportResponseData is nil"))
  }
  switch v.(type) {
    case []import2.Message:
      p.oneOfType1 = v.([]import2.Message)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<common.v1.config.Message>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<common.v1.config.Message>"
    case import3.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = v.(import3.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case NetworkSecurityPolicyImportSummary:
      if nil == p.oneOfType0 {p.oneOfType0 = new(NetworkSecurityPolicyImportSummary)}
      *p.oneOfType0 = v.(NetworkSecurityPolicyImportSummary)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfNetworkSecurityPolicyImportResponseData) GetValue() interface{} {
  if "List<common.v1.config.Message>" == *p.Discriminator {
    return p.oneOfType1
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  return nil
}

func (p *OneOfNetworkSecurityPolicyImportResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType1 := new([]import2.Message)
  if err := json.Unmarshal(b, vOneOfType1); err == nil {
    if len(*vOneOfType1) == 0 || "common.v1.config.Message" == *((*vOneOfType1)[0].ObjectType_) {
      p.oneOfType1 = *vOneOfType1
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<common.v1.config.Message>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<common.v1.config.Message>"
      return nil

    }
  }
  vOneOfType400 := new(import3.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "microseg.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new(NetworkSecurityPolicyImportSummary)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "microseg.v4.policies.NetworkSecurityPolicyImportSummary" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(NetworkSecurityPolicyImportSummary)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfNetworkSecurityPolicyImportResponseData"))
}

func (p *OneOfNetworkSecurityPolicyImportResponseData) MarshalJSON() ([]byte, error) {
  if "List<common.v1.config.Message>" == *p.Discriminator {
    return json.Marshal(p.oneOfType1)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfNetworkSecurityPolicyImportResponseData")
}

type OneOfNetworkSecurityPolicyGetCountsResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 []FieldCountSummary `json:"-"`
  oneOfType1 []import2.Message `json:"-"`
  oneOfType400 *import3.ErrorResponse `json:"-"`
}

func NewOneOfNetworkSecurityPolicyGetCountsResponseData() *OneOfNetworkSecurityPolicyGetCountsResponseData {
  p := new(OneOfNetworkSecurityPolicyGetCountsResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfNetworkSecurityPolicyGetCountsResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfNetworkSecurityPolicyGetCountsResponseData is nil"))
  }
  switch v.(type) {
    case []FieldCountSummary:
      p.oneOfType0 = v.([]FieldCountSummary)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<microseg.v4.policies.FieldCountSummary>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<microseg.v4.policies.FieldCountSummary>"
    case []import2.Message:
      p.oneOfType1 = v.([]import2.Message)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<common.v1.config.Message>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<common.v1.config.Message>"
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

func (p *OneOfNetworkSecurityPolicyGetCountsResponseData) GetValue() interface{} {
  if "List<microseg.v4.policies.FieldCountSummary>" == *p.Discriminator {
    return p.oneOfType0
  }
  if "List<common.v1.config.Message>" == *p.Discriminator {
    return p.oneOfType1
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfNetworkSecurityPolicyGetCountsResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new([]FieldCountSummary)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "microseg.v4.policies.FieldCountSummary" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<microseg.v4.policies.FieldCountSummary>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<microseg.v4.policies.FieldCountSummary>"
      return nil

    }
  }
  vOneOfType1 := new([]import2.Message)
  if err := json.Unmarshal(b, vOneOfType1); err == nil {
    if len(*vOneOfType1) == 0 || "common.v1.config.Message" == *((*vOneOfType1)[0].ObjectType_) {
      p.oneOfType1 = *vOneOfType1
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<common.v1.config.Message>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<common.v1.config.Message>"
      return nil

    }
  }
  vOneOfType400 := new(import3.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "microseg.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfNetworkSecurityPolicyGetCountsResponseData"))
}

func (p *OneOfNetworkSecurityPolicyGetCountsResponseData) MarshalJSON() ([]byte, error) {
  if "List<microseg.v4.policies.FieldCountSummary>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if "List<common.v1.config.Message>" == *p.Discriminator {
    return json.Marshal(p.oneOfType1)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfNetworkSecurityPolicyGetCountsResponseData")
}

type OneOfNetworkSecurityPolicyGetListResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 []NetworkSecurityPolicyWithMetadata `json:"-"`
  oneOfType1 []import2.Message `json:"-"`
  oneOfType400 *import3.ErrorResponse `json:"-"`
}

func NewOneOfNetworkSecurityPolicyGetListResponseData() *OneOfNetworkSecurityPolicyGetListResponseData {
  p := new(OneOfNetworkSecurityPolicyGetListResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfNetworkSecurityPolicyGetListResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfNetworkSecurityPolicyGetListResponseData is nil"))
  }
  switch v.(type) {
    case []NetworkSecurityPolicyWithMetadata:
      p.oneOfType0 = v.([]NetworkSecurityPolicyWithMetadata)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<microseg.v4.policies.NetworkSecurityPolicyWithMetadata>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<microseg.v4.policies.NetworkSecurityPolicyWithMetadata>"
    case []import2.Message:
      p.oneOfType1 = v.([]import2.Message)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<common.v1.config.Message>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<common.v1.config.Message>"
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

func (p *OneOfNetworkSecurityPolicyGetListResponseData) GetValue() interface{} {
  if "List<microseg.v4.policies.NetworkSecurityPolicyWithMetadata>" == *p.Discriminator {
    return p.oneOfType0
  }
  if "List<common.v1.config.Message>" == *p.Discriminator {
    return p.oneOfType1
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfNetworkSecurityPolicyGetListResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new([]NetworkSecurityPolicyWithMetadata)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "microseg.v4.policies.NetworkSecurityPolicyWithMetadata" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<microseg.v4.policies.NetworkSecurityPolicyWithMetadata>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<microseg.v4.policies.NetworkSecurityPolicyWithMetadata>"
      return nil

    }
  }
  vOneOfType1 := new([]import2.Message)
  if err := json.Unmarshal(b, vOneOfType1); err == nil {
    if len(*vOneOfType1) == 0 || "common.v1.config.Message" == *((*vOneOfType1)[0].ObjectType_) {
      p.oneOfType1 = *vOneOfType1
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<common.v1.config.Message>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<common.v1.config.Message>"
      return nil

    }
  }
  vOneOfType400 := new(import3.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "microseg.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfNetworkSecurityPolicyGetListResponseData"))
}

func (p *OneOfNetworkSecurityPolicyGetListResponseData) MarshalJSON() ([]byte, error) {
  if "List<microseg.v4.policies.NetworkSecurityPolicyWithMetadata>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if "List<common.v1.config.Message>" == *p.Discriminator {
    return json.Marshal(p.oneOfType1)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfNetworkSecurityPolicyGetListResponseData")
}

type OneOfSourceOrDestSpec struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType3 *AddressGroupRefSpec `json:"-"`
  oneOfType2 *SubnetSpec `json:"-"`
  oneOfType0 *AllowTypesSpec `json:"-"`
  oneOfType1 *CategoryGroupSpec `json:"-"`
}

func NewOneOfSourceOrDestSpec() *OneOfSourceOrDestSpec {
  p := new(OneOfSourceOrDestSpec)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfSourceOrDestSpec) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfSourceOrDestSpec is nil"))
  }
  switch v.(type) {
    case AddressGroupRefSpec:
      if nil == p.oneOfType3 {p.oneOfType3 = new(AddressGroupRefSpec)}
      *p.oneOfType3 = v.(AddressGroupRefSpec)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType3.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType3.ObjectType_
    case SubnetSpec:
      if nil == p.oneOfType2 {p.oneOfType2 = new(SubnetSpec)}
      *p.oneOfType2 = v.(SubnetSpec)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType2.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType2.ObjectType_
    case AllowTypesSpec:
      if nil == p.oneOfType0 {p.oneOfType0 = new(AllowTypesSpec)}
      *p.oneOfType0 = v.(AllowTypesSpec)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    case CategoryGroupSpec:
      if nil == p.oneOfType1 {p.oneOfType1 = new(CategoryGroupSpec)}
      *p.oneOfType1 = v.(CategoryGroupSpec)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType1.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType1.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfSourceOrDestSpec) GetValue() interface{} {
  if p.oneOfType3 != nil && *p.oneOfType3.ObjectType_ == *p.Discriminator {
    return *p.oneOfType3
  }
  if p.oneOfType2 != nil && *p.oneOfType2.ObjectType_ == *p.Discriminator {
    return *p.oneOfType2
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
    return *p.oneOfType1
  }
  return nil
}

func (p *OneOfSourceOrDestSpec) UnmarshalJSON(b []byte) error {
  vOneOfType3 := new(AddressGroupRefSpec)
  if err := json.Unmarshal(b, vOneOfType3); err == nil {
    if "microseg.v4.policies.AddressGroupRefSpec" == *vOneOfType3.ObjectType_ {
          if nil == p.oneOfType3 {p.oneOfType3 = new(AddressGroupRefSpec)}
      *p.oneOfType3 = *vOneOfType3
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType3.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType3.ObjectType_
      return nil
    }
    }
  vOneOfType2 := new(SubnetSpec)
  if err := json.Unmarshal(b, vOneOfType2); err == nil {
    if "microseg.v4.policies.SubnetSpec" == *vOneOfType2.ObjectType_ {
          if nil == p.oneOfType2 {p.oneOfType2 = new(SubnetSpec)}
      *p.oneOfType2 = *vOneOfType2
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType2.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType2.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new(AllowTypesSpec)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "microseg.v4.policies.AllowTypesSpec" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(AllowTypesSpec)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  vOneOfType1 := new(CategoryGroupSpec)
  if err := json.Unmarshal(b, vOneOfType1); err == nil {
    if "microseg.v4.policies.CategoryGroupSpec" == *vOneOfType1.ObjectType_ {
          if nil == p.oneOfType1 {p.oneOfType1 = new(CategoryGroupSpec)}
      *p.oneOfType1 = *vOneOfType1
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType1.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType1.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfSourceOrDestSpec"))
}

func (p *OneOfSourceOrDestSpec) MarshalJSON() ([]byte, error) {
  if p.oneOfType3 != nil && *p.oneOfType3.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType3)
  }
  if p.oneOfType2 != nil && *p.oneOfType2.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType2)
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType1)
  }
  return nil, errors.New("No value to marshal for OneOfSourceOrDestSpec")
}

type OneOfNetworkSecurityPolicyTaskResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType1 []import2.Message `json:"-"`
  oneOfType0 *import4.TaskReference `json:"-"`
  oneOfType400 *import3.ErrorResponse `json:"-"`
}

func NewOneOfNetworkSecurityPolicyTaskResponseData() *OneOfNetworkSecurityPolicyTaskResponseData {
  p := new(OneOfNetworkSecurityPolicyTaskResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfNetworkSecurityPolicyTaskResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfNetworkSecurityPolicyTaskResponseData is nil"))
  }
  switch v.(type) {
    case []import2.Message:
      p.oneOfType1 = v.([]import2.Message)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<common.v1.config.Message>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<common.v1.config.Message>"
    case import4.TaskReference:
      if nil == p.oneOfType0 {p.oneOfType0 = new(import4.TaskReference)}
      *p.oneOfType0 = v.(import4.TaskReference)
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

func (p *OneOfNetworkSecurityPolicyTaskResponseData) GetValue() interface{} {
  if "List<common.v1.config.Message>" == *p.Discriminator {
    return p.oneOfType1
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfNetworkSecurityPolicyTaskResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType1 := new([]import2.Message)
  if err := json.Unmarshal(b, vOneOfType1); err == nil {
    if len(*vOneOfType1) == 0 || "common.v1.config.Message" == *((*vOneOfType1)[0].ObjectType_) {
      p.oneOfType1 = *vOneOfType1
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<common.v1.config.Message>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<common.v1.config.Message>"
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
  vOneOfType400 := new(import3.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "microseg.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfNetworkSecurityPolicyTaskResponseData"))
}

func (p *OneOfNetworkSecurityPolicyTaskResponseData) MarshalJSON() ([]byte, error) {
  if "List<common.v1.config.Message>" == *p.Discriminator {
    return json.Marshal(p.oneOfType1)
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfNetworkSecurityPolicyTaskResponseData")
}

type OneOfNetworkSecurityPolicyGetResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType1 []import2.Message `json:"-"`
  oneOfType400 *import3.ErrorResponse `json:"-"`
  oneOfType0 *NetworkSecurityPolicyWithMetadata `json:"-"`
}

func NewOneOfNetworkSecurityPolicyGetResponseData() *OneOfNetworkSecurityPolicyGetResponseData {
  p := new(OneOfNetworkSecurityPolicyGetResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfNetworkSecurityPolicyGetResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfNetworkSecurityPolicyGetResponseData is nil"))
  }
  switch v.(type) {
    case []import2.Message:
      p.oneOfType1 = v.([]import2.Message)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<common.v1.config.Message>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<common.v1.config.Message>"
    case import3.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = v.(import3.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case NetworkSecurityPolicyWithMetadata:
      if nil == p.oneOfType0 {p.oneOfType0 = new(NetworkSecurityPolicyWithMetadata)}
      *p.oneOfType0 = v.(NetworkSecurityPolicyWithMetadata)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfNetworkSecurityPolicyGetResponseData) GetValue() interface{} {
  if "List<common.v1.config.Message>" == *p.Discriminator {
    return p.oneOfType1
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  return nil
}

func (p *OneOfNetworkSecurityPolicyGetResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType1 := new([]import2.Message)
  if err := json.Unmarshal(b, vOneOfType1); err == nil {
    if len(*vOneOfType1) == 0 || "common.v1.config.Message" == *((*vOneOfType1)[0].ObjectType_) {
      p.oneOfType1 = *vOneOfType1
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<common.v1.config.Message>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<common.v1.config.Message>"
      return nil

    }
  }
  vOneOfType400 := new(import3.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "microseg.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new(NetworkSecurityPolicyWithMetadata)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "microseg.v4.policies.NetworkSecurityPolicyWithMetadata" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(NetworkSecurityPolicyWithMetadata)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfNetworkSecurityPolicyGetResponseData"))
}

func (p *OneOfNetworkSecurityPolicyGetResponseData) MarshalJSON() ([]byte, error) {
  if "List<common.v1.config.Message>" == *p.Discriminator {
    return json.Marshal(p.oneOfType1)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfNetworkSecurityPolicyGetResponseData")
}

type OneOfNetworkSecurityPolicyExportResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType1 []import2.Message `json:"-"`
  oneOfType400 *import3.ErrorResponse `json:"-"`
  oneOfType0 *FileWrapper `json:"-"`
}

func NewOneOfNetworkSecurityPolicyExportResponseData() *OneOfNetworkSecurityPolicyExportResponseData {
  p := new(OneOfNetworkSecurityPolicyExportResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfNetworkSecurityPolicyExportResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfNetworkSecurityPolicyExportResponseData is nil"))
  }
  switch v.(type) {
    case []import2.Message:
      p.oneOfType1 = v.([]import2.Message)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<common.v1.config.Message>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<common.v1.config.Message>"
    case import3.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = v.(import3.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case FileWrapper:
      if nil == p.oneOfType0 {p.oneOfType0 = new(FileWrapper)}
      *p.oneOfType0 = v.(FileWrapper)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfNetworkSecurityPolicyExportResponseData) GetValue() interface{} {
  if "List<common.v1.config.Message>" == *p.Discriminator {
    return p.oneOfType1
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  return nil
}

func (p *OneOfNetworkSecurityPolicyExportResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType1 := new([]import2.Message)
  if err := json.Unmarshal(b, vOneOfType1); err == nil {
    if len(*vOneOfType1) == 0 || "common.v1.config.Message" == *((*vOneOfType1)[0].ObjectType_) {
      p.oneOfType1 = *vOneOfType1
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<common.v1.config.Message>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<common.v1.config.Message>"
      return nil

    }
  }
  vOneOfType400 := new(import3.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "microseg.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new(FileWrapper)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "microseg.v4.policies.FileWrapper" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(FileWrapper)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfNetworkSecurityPolicyExportResponseData"))
}

func (p *OneOfNetworkSecurityPolicyExportResponseData) MarshalJSON() ([]byte, error) {
  if "List<common.v1.config.Message>" == *p.Discriminator {
    return json.Marshal(p.oneOfType1)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfNetworkSecurityPolicyExportResponseData")
}

type OneOfServiceProtocolDetailsSpec struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType1 *ICMPTypeCodeSpec `json:"-"`
  oneOfType0 *PortRangeSpec `json:"-"`
}

func NewOneOfServiceProtocolDetailsSpec() *OneOfServiceProtocolDetailsSpec {
  p := new(OneOfServiceProtocolDetailsSpec)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfServiceProtocolDetailsSpec) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfServiceProtocolDetailsSpec is nil"))
  }
  switch v.(type) {
    case ICMPTypeCodeSpec:
      if nil == p.oneOfType1 {p.oneOfType1 = new(ICMPTypeCodeSpec)}
      *p.oneOfType1 = v.(ICMPTypeCodeSpec)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType1.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType1.ObjectType_
    case PortRangeSpec:
      if nil == p.oneOfType0 {p.oneOfType0 = new(PortRangeSpec)}
      *p.oneOfType0 = v.(PortRangeSpec)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfServiceProtocolDetailsSpec) GetValue() interface{} {
  if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
    return *p.oneOfType1
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  return nil
}

func (p *OneOfServiceProtocolDetailsSpec) UnmarshalJSON(b []byte) error {
  vOneOfType1 := new(ICMPTypeCodeSpec)
  if err := json.Unmarshal(b, vOneOfType1); err == nil {
    if "microseg.v4.policies.ICMPTypeCodeSpec" == *vOneOfType1.ObjectType_ {
          if nil == p.oneOfType1 {p.oneOfType1 = new(ICMPTypeCodeSpec)}
      *p.oneOfType1 = *vOneOfType1
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType1.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType1.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new(PortRangeSpec)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "microseg.v4.policies.PortRangeSpec" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(PortRangeSpec)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfServiceProtocolDetailsSpec"))
}

func (p *OneOfServiceProtocolDetailsSpec) MarshalJSON() ([]byte, error) {
  if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType1)
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfServiceProtocolDetailsSpec")
}

type OneOfRuleSpec struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *NetworkPolicyTwoEnvIsolationRuleSpec `json:"-"`
  oneOfType2 *NetworkPolicyIntraEntityGroupRuleSpec `json:"-"`
  oneOfType1 *NetworkPolicyApplicationRuleSpec `json:"-"`
}

func NewOneOfRuleSpec() *OneOfRuleSpec {
  p := new(OneOfRuleSpec)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfRuleSpec) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfRuleSpec is nil"))
  }
  switch v.(type) {
    case NetworkPolicyTwoEnvIsolationRuleSpec:
      if nil == p.oneOfType0 {p.oneOfType0 = new(NetworkPolicyTwoEnvIsolationRuleSpec)}
      *p.oneOfType0 = v.(NetworkPolicyTwoEnvIsolationRuleSpec)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    case NetworkPolicyIntraEntityGroupRuleSpec:
      if nil == p.oneOfType2 {p.oneOfType2 = new(NetworkPolicyIntraEntityGroupRuleSpec)}
      *p.oneOfType2 = v.(NetworkPolicyIntraEntityGroupRuleSpec)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType2.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType2.ObjectType_
    case NetworkPolicyApplicationRuleSpec:
      if nil == p.oneOfType1 {p.oneOfType1 = new(NetworkPolicyApplicationRuleSpec)}
      *p.oneOfType1 = v.(NetworkPolicyApplicationRuleSpec)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType1.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType1.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfRuleSpec) GetValue() interface{} {
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

func (p *OneOfRuleSpec) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new(NetworkPolicyTwoEnvIsolationRuleSpec)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "microseg.v4.policies.NetworkPolicyTwoEnvIsolationRuleSpec" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(NetworkPolicyTwoEnvIsolationRuleSpec)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  vOneOfType2 := new(NetworkPolicyIntraEntityGroupRuleSpec)
  if err := json.Unmarshal(b, vOneOfType2); err == nil {
    if "microseg.v4.policies.NetworkPolicyIntraEntityGroupRuleSpec" == *vOneOfType2.ObjectType_ {
          if nil == p.oneOfType2 {p.oneOfType2 = new(NetworkPolicyIntraEntityGroupRuleSpec)}
      *p.oneOfType2 = *vOneOfType2
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType2.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType2.ObjectType_
      return nil
    }
    }
  vOneOfType1 := new(NetworkPolicyApplicationRuleSpec)
  if err := json.Unmarshal(b, vOneOfType1); err == nil {
    if "microseg.v4.policies.NetworkPolicyApplicationRuleSpec" == *vOneOfType1.ObjectType_ {
          if nil == p.oneOfType1 {p.oneOfType1 = new(NetworkPolicyApplicationRuleSpec)}
      *p.oneOfType1 = *vOneOfType1
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType1.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType1.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfRuleSpec"))
}

func (p *OneOfRuleSpec) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType2 != nil && *p.oneOfType2.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType2)
  }
  if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType1)
  }
  return nil, errors.New("No value to marshal for OneOfRuleSpec")
}

type OneOfServiceSpec struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *ServiceGroupRefSpec `json:"-"`
  oneOfType1 *ServiceSpec `json:"-"`
}

func NewOneOfServiceSpec() *OneOfServiceSpec {
  p := new(OneOfServiceSpec)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfServiceSpec) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfServiceSpec is nil"))
  }
  switch v.(type) {
    case ServiceGroupRefSpec:
      if nil == p.oneOfType0 {p.oneOfType0 = new(ServiceGroupRefSpec)}
      *p.oneOfType0 = v.(ServiceGroupRefSpec)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    case ServiceSpec:
      if nil == p.oneOfType1 {p.oneOfType1 = new(ServiceSpec)}
      *p.oneOfType1 = v.(ServiceSpec)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType1.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType1.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfServiceSpec) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
    return *p.oneOfType1
  }
  return nil
}

func (p *OneOfServiceSpec) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new(ServiceGroupRefSpec)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "microseg.v4.policies.ServiceGroupRefSpec" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(ServiceGroupRefSpec)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  vOneOfType1 := new(ServiceSpec)
  if err := json.Unmarshal(b, vOneOfType1); err == nil {
    if "microseg.v4.policies.ServiceSpec" == *vOneOfType1.ObjectType_ {
          if nil == p.oneOfType1 {p.oneOfType1 = new(ServiceSpec)}
      *p.oneOfType1 = *vOneOfType1
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType1.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType1.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfServiceSpec"))
}

func (p *OneOfServiceSpec) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType1)
  }
  return nil, errors.New("No value to marshal for OneOfServiceSpec")
}

