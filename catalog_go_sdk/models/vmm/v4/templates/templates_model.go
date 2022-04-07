/*
 * Generated file models/vmm/v4/templates/templates_model.go.
 *
 * Product version: 16.2.0-SNAPSHOT
 *
 * Part of the Catalog REST API
 *
 * (c) 2022 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module vmm.v4.templates of Catalog REST API
*/
package templates
import (
  import1 "github.com/nutanix-core/ntnx-api-go-sdk-internal/catalog_go_sdk/v16/models/common/v1/config"
  import2 "github.com/nutanix-core/ntnx-api-go-sdk-internal/catalog_go_sdk/v16/models/common/v1/response"
  "encoding/json"
  "errors"
  "fmt"
  "time"
  import4 "github.com/nutanix-core/ntnx-api-go-sdk-internal/catalog_go_sdk/v16/models/vmm/v4/config"
  import3 "github.com/nutanix-core/ntnx-api-go-sdk-internal/catalog_go_sdk/v16/models/vmm/v4/error"
)

/**
VM reference for Template guest changes.
*/
type GuestChangesStatus struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  VM reference UUID.
  */
  DeployedVmReference *string `json:"deployedVmReference,omitempty"`
}

func NewGuestChangesStatus() *GuestChangesStatus {
  p := new(GuestChangesStatus)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "vmm.v4.templates.GuestChangesStatus"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.templates.GuestChangesStatus"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}





type Template struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  The current gold version.
  */
  ActiveVersionNumber *int `json:"activeVersionNumber,omitempty"`
  /**
  Created Time.
  */
  CreatedAt *time.Time `json:"createdAt,omitempty"`
  /**
  The user who created the Template.
  */
  CreatedBy *string `json:"createdBy,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  An array of entity specific metadata
  */
  ExtraInfo []import1.KVPair `json:"extraInfo,omitempty"`
  
  GuestChangesStatus *GuestChangesStatus `json:"guestChangesStatus,omitempty"`
  /**
  Indicates whether the version has to be a Gold version or not.
  */
  IsActiveVersion *bool `json:"isActiveVersion,omitempty"`
  /**
  Last Update Time.
  */
  LastUpdatedAt *time.Time `json:"lastUpdatedAt,omitempty"`
  /**
  The user who updated the Version in Template.
  */
  LastUpdatedBy *string `json:"lastUpdatedBy,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import2.ApiLink `json:"links,omitempty"`
  /**
  A description for the Template.
  */
  TemplateDescription *string `json:"templateDescription,omitempty"`
  /**
  Template Name.
  */
  TemplateName *string `json:"templateName,omitempty"`
  
  TemplateVersionDetails *TemplateVersionDetails `json:"templateVersionDetails,omitempty"`
  
  TemplateVersionSpec *TemplateSourceReference `json:"templateVersionSpec,omitempty"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
}

func NewTemplate() *Template {
  p := new(Template)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "vmm.v4.templates.Template"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.templates.Template"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST Response for all response codes in api path /vmm/v4.0.a1/templates/{templateExtId} Get operation
*/
type TemplateApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfTemplateApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewTemplateApiResponse() *TemplateApiResponse {
  p := new(TemplateApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "vmm.v4.templates.TemplateApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.templates.TemplateApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *TemplateApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *TemplateApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfTemplateApiResponseData()
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
Template Guest Complete Changes.
*/
type TemplateCompleteGuestChanges struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Indicates whether the version has to be a Gold version or not.
  */
  IsActiveVersion *bool `json:"isActiveVersion,omitempty"`
  /**
  A description for template version.
  */
  VersionDescription *string `json:"versionDescription"`
  /**
  Template version name. This will be the default VM name prefix when a VM is deployed from this template.
  */
  VersionName *string `json:"versionName"`
}

func NewTemplateCompleteGuestChanges() *TemplateCompleteGuestChanges {
  p := new(TemplateCompleteGuestChanges)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "vmm.v4.templates.TemplateCompleteGuestChanges"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.templates.TemplateCompleteGuestChanges"}
  p.UnknownFields_ = map[string]interface{}{}

  p.IsActiveVersion = new(bool)
  *p.IsActiveVersion = false


  return p
}




/**
Deploy VM Config from the Template.
*/
type TemplateDeployment struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  The UUID of the Cluster where the VM has to be deployed. This is mandatory to be specified for creating the VM.
  */
  ClusterReference *string `json:"clusterReference"`
  /**
  Number / Count of VMs to be deployed.
  */
  NumberOfVms *int `json:"numberOfVms"`
  
  OverrideVmConfigMap map[string]string `json:"overrideVmConfigMap,omitempty"`
  /**
  Start Index of the VM.
  */
  StartIndex *int `json:"startIndex"`
  
  VersionNumber *int `json:"versionNumber,omitempty"`
  /**
  vmName accepts the Name for the VM in moustache template format. For example 'sql_server_{{startIndex}}'. The marker will be replaced with startIndex value. The startIndex will monotonically increase up to the numberOfVms. User can provide this marker in any location of the vmName value.
  */
  VmName *string `json:"vmName"`
}

func NewTemplateDeployment() *TemplateDeployment {
  p := new(TemplateDeployment)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "vmm.v4.templates.TemplateDeployment"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.templates.TemplateDeployment"}
  p.UnknownFields_ = map[string]interface{}{}

  p.NumberOfVms = new(int)
  *p.NumberOfVms = 1
  p.StartIndex = new(int)
  *p.StartIndex = 0


  return p
}




/**
Template Guest Changes Input.
*/
type TemplateGuestChanges struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Template Version number
  */
  VersionNumber *int `json:"versionNumber,omitempty"`
}

func NewTemplateGuestChanges() *TemplateGuestChanges {
  p := new(TemplateGuestChanges)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "vmm.v4.templates.TemplateGuestChanges"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.templates.TemplateGuestChanges"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST Response for all response codes in api path /vmm/v4.0.a1/templates Get operation
*/
type TemplateListApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfTemplateListApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewTemplateListApiResponse() *TemplateListApiResponse {
  p := new(TemplateListApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "vmm.v4.templates.TemplateListApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.templates.TemplateListApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *TemplateListApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *TemplateListApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfTemplateListApiResponseData()
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




type TemplateQueryable struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Default value is false. Set spec to true to include the VM spec in the template response.
  */
  VmSpec *bool `json:"vmSpec,omitempty"`
}

func NewTemplateQueryable() *TemplateQueryable {
  p := new(TemplateQueryable)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "vmm.v4.templates.TemplateQueryable"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.templates.TemplateQueryable"}
  p.UnknownFields_ = map[string]interface{}{}

  p.VmSpec = new(bool)
  *p.VmSpec = false


  return p
}




/**
Spec for creating new version of template using VM Reference or Template Version number and override config.
*/
type TemplateSourceReference struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  ConfigItemDiscriminator_ *string `json:"$configItemDiscriminator,omitempty"`
  
  Config *OneOfTemplateSourceReferenceConfig `json:"config,omitempty"`
  
  ConfigItemDiscriminator *string `json:"configItemDiscriminator,omitempty"`
  /**
  Allow or Disallow override of the Guest Customization during Template Deployment using this version.
  */
  GcOverride *bool `json:"gcOverride,omitempty"`
  /**
  A description for template version.
  */
  VersionDescription *string `json:"versionDescription,omitempty"`
  /**
  Template version name. This will be the default VM name prefix when a VM is deployed from this template.
  */
  VersionName *string `json:"versionName,omitempty"`
}

func NewTemplateSourceReference() *TemplateSourceReference {
  p := new(TemplateSourceReference)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "vmm.v4.templates.TemplateSourceReference"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.templates.TemplateSourceReference"}
  p.UnknownFields_ = map[string]interface{}{}

  p.GcOverride = new(bool)
  *p.GcOverride = false


  return p
}

func (p *TemplateSourceReference) GetConfig() interface{} {
  if nil == p.Config {
    return nil
  }
  return p.Config.GetValue()
}

func (p *TemplateSourceReference) SetConfig(v interface{}) error {
  if nil == p.Config {
    p.Config = NewOneOfTemplateSourceReferenceConfig()
  }
  e := p.Config.SetValue(v)
  if nil == e {
    if nil == p.ConfigItemDiscriminator_ {
      p.ConfigItemDiscriminator_ = new(string)
    }
    *p.ConfigItemDiscriminator_ = *p.Config.Discriminator
  }
  return e
}



/**
REST Response for all response codes in api path /vmm/v4.0.a1/templates/{templateExtId}/versions/{versionNumber} Get operation
*/
type TemplateVersionApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfTemplateVersionApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewTemplateVersionApiResponse() *TemplateVersionApiResponse {
  p := new(TemplateVersionApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "vmm.v4.templates.TemplateVersionApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.templates.TemplateVersionApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *TemplateVersionApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *TemplateVersionApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfTemplateVersionApiResponseData()
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
A representation of template version entity.
*/
type TemplateVersionDetails struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Created Time.
  */
  CreatedAt *time.Time `json:"createdAt,omitempty"`
  /**
  The user who created the Version in Template.
  */
  CreatedBy *string `json:"createdBy,omitempty"`
  /**
  Allow or Disallow override of the Guest Customization during Template Deployment using this version.
  */
  GcOverride *bool `json:"gcOverride,omitempty"`
  /**
  A description for template version.
  */
  VersionDescription *string `json:"versionDescription,omitempty"`
  /**
  Template version name. This will be the default VM name prefix when a VM is deployed from this template.
  */
  VersionName *string `json:"versionName"`
  
  VersionNumber *int `json:"versionNumber"`
  /**
  VM configuration spec
  */
  VmSpec *string `json:"vmSpec"`
}

func NewTemplateVersionDetails() *TemplateVersionDetails {
  p := new(TemplateVersionDetails)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "vmm.v4.templates.TemplateVersionDetails"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.templates.TemplateVersionDetails"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST Response for all response codes in api path /vmm/v4.0.a1/templates/{templateExtId}/versions Get operation
*/
type TemplateVersionListApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfTemplateVersionListApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewTemplateVersionListApiResponse() *TemplateVersionListApiResponse {
  p := new(TemplateVersionListApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "vmm.v4.templates.TemplateVersionListApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.templates.TemplateVersionListApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *TemplateVersionListApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *TemplateVersionListApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfTemplateVersionListApiResponseData()
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
Template new version with Version Number and Override Config.
*/
type TemplateVersionReference struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  A valid VM spec
  */
  OverrideVmConfig *string `json:"overrideVmConfig"`
  
  Version *int `json:"version"`
}

func NewTemplateVersionReference() *TemplateVersionReference {
  p := new(TemplateVersionReference)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "vmm.v4.templates.TemplateVersionReference"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.templates.TemplateVersionReference"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST Response for all response codes in api path /vmm/v4.0.a1/templates/{templateExtId}/versions/{versionNumber}/$actions/publish Post operation
*/
type TemplatesTaskApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfTemplatesTaskApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewTemplatesTaskApiResponse() *TemplatesTaskApiResponse {
  p := new(TemplatesTaskApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "vmm.v4.templates.TemplatesTaskApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.templates.TemplatesTaskApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *TemplatesTaskApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *TemplatesTaskApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfTemplatesTaskApiResponseData()
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
VM UUID
*/
type VmReference struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  VM UUID
  */
  ExtId *string `json:"extId"`
  /**
  VM guests may be customized at boot time using one of several different methods. Currently, cloud-init w/ ConfigDriveV2 (for Linux VMs) and Sysprep (for Windows VMs) are supported. Only ONE OF Sysprep or CloudInit should be provided. Note that guest customization can currently only be set during VM creation. Attempting to change it after creation will result in an error. Additional properties can be specified. For example - in the context of VM template creation if 'isOverridable' is set to 'True' then the deployer can upload their own custom script.
  */
  GuestCustomization *string `json:"guestCustomization,omitempty"`
}

func NewVmReference() *VmReference {
  p := new(VmReference)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "vmm.v4.templates.VmReference"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.templates.VmReference"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}



type OneOfTemplateSourceReferenceConfig struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *VmReference `json:"-"`
  oneOfType1 *TemplateVersionReference `json:"-"`
}

func NewOneOfTemplateSourceReferenceConfig() *OneOfTemplateSourceReferenceConfig {
  p := new(OneOfTemplateSourceReferenceConfig)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfTemplateSourceReferenceConfig) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfTemplateSourceReferenceConfig is nil"))
  }
  switch v.(type) {
    case VmReference:
      if nil == p.oneOfType0 {p.oneOfType0 = new(VmReference)}
      *p.oneOfType0 = v.(VmReference)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    case TemplateVersionReference:
      if nil == p.oneOfType1 {p.oneOfType1 = new(TemplateVersionReference)}
      *p.oneOfType1 = v.(TemplateVersionReference)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType1.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType1.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfTemplateSourceReferenceConfig) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
    return *p.oneOfType1
  }
  return nil
}

func (p *OneOfTemplateSourceReferenceConfig) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new(VmReference)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "vmm.v4.templates.VmReference" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(VmReference)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  vOneOfType1 := new(TemplateVersionReference)
  if err := json.Unmarshal(b, vOneOfType1); err == nil {
    if "vmm.v4.templates.TemplateVersionReference" == *vOneOfType1.ObjectType_ {
          if nil == p.oneOfType1 {p.oneOfType1 = new(TemplateVersionReference)}
      *p.oneOfType1 = *vOneOfType1
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType1.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType1.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfTemplateSourceReferenceConfig"))
}

func (p *OneOfTemplateSourceReferenceConfig) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType1)
  }
  return nil, errors.New("No value to marshal for OneOfTemplateSourceReferenceConfig")
}

type OneOfTemplateVersionApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *TemplateVersionDetails `json:"-"`
  oneOfType400 *import3.ErrorResponse `json:"-"`
}

func NewOneOfTemplateVersionApiResponseData() *OneOfTemplateVersionApiResponseData {
  p := new(OneOfTemplateVersionApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfTemplateVersionApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfTemplateVersionApiResponseData is nil"))
  }
  switch v.(type) {
    case TemplateVersionDetails:
      if nil == p.oneOfType0 {p.oneOfType0 = new(TemplateVersionDetails)}
      *p.oneOfType0 = v.(TemplateVersionDetails)
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

func (p *OneOfTemplateVersionApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfTemplateVersionApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new(TemplateVersionDetails)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "vmm.v4.templates.TemplateVersionDetails" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(TemplateVersionDetails)}
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
    if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfTemplateVersionApiResponseData"))
}

func (p *OneOfTemplateVersionApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfTemplateVersionApiResponseData")
}

type OneOfTemplatesTaskApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *import4.Task `json:"-"`
  oneOfType400 *import3.ErrorResponse `json:"-"`
}

func NewOneOfTemplatesTaskApiResponseData() *OneOfTemplatesTaskApiResponseData {
  p := new(OneOfTemplatesTaskApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfTemplatesTaskApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfTemplatesTaskApiResponseData is nil"))
  }
  switch v.(type) {
    case import4.Task:
      if nil == p.oneOfType0 {p.oneOfType0 = new(import4.Task)}
      *p.oneOfType0 = v.(import4.Task)
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

func (p *OneOfTemplatesTaskApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfTemplatesTaskApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new(import4.Task)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "vmm.v4.config.Task" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(import4.Task)}
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
    if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfTemplatesTaskApiResponseData"))
}

func (p *OneOfTemplatesTaskApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfTemplatesTaskApiResponseData")
}

type OneOfTemplateListApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 []Template `json:"-"`
  oneOfType400 *import3.ErrorResponse `json:"-"`
}

func NewOneOfTemplateListApiResponseData() *OneOfTemplateListApiResponseData {
  p := new(OneOfTemplateListApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfTemplateListApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfTemplateListApiResponseData is nil"))
  }
  switch v.(type) {
    case []Template:
      p.oneOfType0 = v.([]Template)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<vmm.v4.templates.Template>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<vmm.v4.templates.Template>"
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

func (p *OneOfTemplateListApiResponseData) GetValue() interface{} {
  if "List<vmm.v4.templates.Template>" == *p.Discriminator {
    return p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfTemplateListApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new([]Template)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "vmm.v4.templates.Template" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<vmm.v4.templates.Template>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<vmm.v4.templates.Template>"
      return nil

    }
  }
  vOneOfType400 := new(import3.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfTemplateListApiResponseData"))
}

func (p *OneOfTemplateListApiResponseData) MarshalJSON() ([]byte, error) {
  if "List<vmm.v4.templates.Template>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfTemplateListApiResponseData")
}

type OneOfTemplateVersionListApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 []TemplateVersionDetails `json:"-"`
  oneOfType400 *import3.ErrorResponse `json:"-"`
}

func NewOneOfTemplateVersionListApiResponseData() *OneOfTemplateVersionListApiResponseData {
  p := new(OneOfTemplateVersionListApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfTemplateVersionListApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfTemplateVersionListApiResponseData is nil"))
  }
  switch v.(type) {
    case []TemplateVersionDetails:
      p.oneOfType0 = v.([]TemplateVersionDetails)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<vmm.v4.templates.TemplateVersionDetails>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<vmm.v4.templates.TemplateVersionDetails>"
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

func (p *OneOfTemplateVersionListApiResponseData) GetValue() interface{} {
  if "List<vmm.v4.templates.TemplateVersionDetails>" == *p.Discriminator {
    return p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfTemplateVersionListApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new([]TemplateVersionDetails)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "vmm.v4.templates.TemplateVersionDetails" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<vmm.v4.templates.TemplateVersionDetails>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<vmm.v4.templates.TemplateVersionDetails>"
      return nil

    }
  }
  vOneOfType400 := new(import3.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfTemplateVersionListApiResponseData"))
}

func (p *OneOfTemplateVersionListApiResponseData) MarshalJSON() ([]byte, error) {
  if "List<vmm.v4.templates.TemplateVersionDetails>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfTemplateVersionListApiResponseData")
}

type OneOfTemplateApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *Template `json:"-"`
  oneOfType400 *import3.ErrorResponse `json:"-"`
}

func NewOneOfTemplateApiResponseData() *OneOfTemplateApiResponseData {
  p := new(OneOfTemplateApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfTemplateApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfTemplateApiResponseData is nil"))
  }
  switch v.(type) {
    case Template:
      if nil == p.oneOfType0 {p.oneOfType0 = new(Template)}
      *p.oneOfType0 = v.(Template)
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

func (p *OneOfTemplateApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfTemplateApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new(Template)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "vmm.v4.templates.Template" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(Template)}
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
    if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfTemplateApiResponseData"))
}

func (p *OneOfTemplateApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfTemplateApiResponseData")
}

