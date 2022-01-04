/*
 * Generated file models/storage/v4/config/config_model.go.
 *
 * Product version: 16.2.0-SNAPSHOT
 *
 * Part of the Storage APIs
 *
 * (c) 2022 Nutanix Inc.  All rights reserved
 *
 */

/*
  config
*/
package config
import (
  "bytes"
  import4 "github.com/nishantikagarg/sample-go-repo/storage_go_sdk/v16/models/common/v1/config"
  import2 "github.com/nishantikagarg/sample-go-repo/storage_go_sdk/v16/models/common/v1/response"
  "encoding/json"
  "errors"
  "fmt"
  import3 "github.com/nishantikagarg/sample-go-repo/storage_go_sdk/v16/models/prism/v4/config"
  import1 "github.com/nishantikagarg/sample-go-repo/storage_go_sdk/v16/models/storage/v4/error"
)

/**
REST Response for all response codes in api path /storage/v4.0.a3/config/volume-groups/{volumeGroupExtId}/$actions/associate-category Post operation
*/
type AssociateCategoryApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfAssociateCategoryApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewAssociateCategoryApiResponse() *AssociateCategoryApiResponse {
  p := new(AssociateCategoryApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "storage.v4.config.AssociateCategoryApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "storage.v4.r0.a3.config.AssociateCategoryApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *AssociateCategoryApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *AssociateCategoryApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfAssociateCategoryApiResponseData()
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
REST Response for all response codes in api path /storage/v4.0.a3/config/volume-groups/{volumeGroupExtId}/$actions/attach-iscsi-client Post operation
*/
type AttachIscsiClientApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfAttachIscsiClientApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewAttachIscsiClientApiResponse() *AttachIscsiClientApiResponse {
  p := new(AttachIscsiClientApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "storage.v4.config.AttachIscsiClientApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "storage.v4.r0.a3.config.AttachIscsiClientApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *AttachIscsiClientApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *AttachIscsiClientApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfAttachIscsiClientApiResponseData()
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
REST Response for all response codes in api path /storage/v4.0.a3/config/volume-groups/{volumeGroupExtId}/$actions/attach-vm Post operation
*/
type AttachVmApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfAttachVmApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewAttachVmApiResponse() *AttachVmApiResponse {
  p := new(AttachVmApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "storage.v4.config.AttachVmApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "storage.v4.r0.a3.config.AttachVmApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *AttachVmApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *AttachVmApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfAttachVmApiResponseData()
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
The type of authentication enabled for the Volume Group. If this is set to CHAP, the target/client secret must be provided.
*/
type AuthenticationType int

const(
  AUTHENTICATIONTYPE_UNKNOWN AuthenticationType = 0
  AUTHENTICATIONTYPE_REDACTED AuthenticationType = 1
  AUTHENTICATIONTYPE_CHAP AuthenticationType = 2
  AUTHENTICATIONTYPE_NONE AuthenticationType = 3
)

// returns the name of the enum given an ordinal number
func (e *AuthenticationType) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "CHAP",
    "NONE",
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
    "CHAP",
    "NONE",
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
An existing category details associated with the Volume Group.
*/
type CategoryDetails struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  EntityType *import4.EntityType `json:"entityType,omitempty"`
  
  ExtId *string `json:"extId,omitempty"`
  
  Name *string `json:"name,omitempty"`
  
  Uris []string `json:"uris,omitempty"`
}

func NewCategoryDetails() *CategoryDetails {
  p := new(CategoryDetails)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "storage.v4.config.CategoryDetails"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "storage.v4.r0.a3.config.CategoryDetails"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
The list of categories to be associated/disassociated with the Volume Group.
*/
type CategoryEntityReferences struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  Categories []import4.EntityReference `json:"categories,omitempty"`
}

func NewCategoryEntityReferences() *CategoryEntityReferences {
  p := new(CategoryEntityReferences)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "storage.v4.config.CategoryEntityReferences"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "storage.v4.r0.a3.config.CategoryEntityReferences"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST Response for all response codes in api path /storage/v4.0.a3/config/volume-groups/{volumeGroupExtId}/disks Post operation
*/
type CreateVolumeDiskApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfCreateVolumeDiskApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewCreateVolumeDiskApiResponse() *CreateVolumeDiskApiResponse {
  p := new(CreateVolumeDiskApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "storage.v4.config.CreateVolumeDiskApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "storage.v4.r0.a3.config.CreateVolumeDiskApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *CreateVolumeDiskApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *CreateVolumeDiskApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfCreateVolumeDiskApiResponseData()
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
REST Response for all response codes in api path /storage/v4.0.a3/config/volume-groups Post operation
*/
type CreateVolumeGroupApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfCreateVolumeGroupApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewCreateVolumeGroupApiResponse() *CreateVolumeGroupApiResponse {
  p := new(CreateVolumeGroupApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "storage.v4.config.CreateVolumeGroupApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "storage.v4.r0.a3.config.CreateVolumeGroupApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *CreateVolumeGroupApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *CreateVolumeGroupApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfCreateVolumeGroupApiResponseData()
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
REST Response for all response codes in api path /storage/v4.0.a3/config/volume-groups/{volumeGroupExtId}/disks/{diskExtId} Delete operation
*/
type DeleteVolumeDiskApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfDeleteVolumeDiskApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDeleteVolumeDiskApiResponse() *DeleteVolumeDiskApiResponse {
  p := new(DeleteVolumeDiskApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "storage.v4.config.DeleteVolumeDiskApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "storage.v4.r0.a3.config.DeleteVolumeDiskApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *DeleteVolumeDiskApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *DeleteVolumeDiskApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfDeleteVolumeDiskApiResponseData()
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
REST Response for all response codes in api path /storage/v4.0.a3/config/volume-groups/{extId} Delete operation
*/
type DeleteVolumeGroupApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfDeleteVolumeGroupApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDeleteVolumeGroupApiResponse() *DeleteVolumeGroupApiResponse {
  p := new(DeleteVolumeGroupApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "storage.v4.config.DeleteVolumeGroupApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "storage.v4.r0.a3.config.DeleteVolumeGroupApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *DeleteVolumeGroupApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *DeleteVolumeGroupApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfDeleteVolumeGroupApiResponseData()
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
REST Response for all response codes in api path /storage/v4.0.a3/config/volume-groups/{volumeGroupExtId}/$actions/detach-iscsi-client/{clientId} Post operation
*/
type DetachIscsiClientApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfDetachIscsiClientApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDetachIscsiClientApiResponse() *DetachIscsiClientApiResponse {
  p := new(DetachIscsiClientApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "storage.v4.config.DetachIscsiClientApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "storage.v4.r0.a3.config.DetachIscsiClientApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *DetachIscsiClientApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *DetachIscsiClientApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfDetachIscsiClientApiResponseData()
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
REST Response for all response codes in api path /storage/v4.0.a3/config/volume-groups/{volumeGroupExtId}/$actions/detach-vm/{vmExtId} Post operation
*/
type DetachVmApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfDetachVmApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDetachVmApiResponse() *DetachVmApiResponse {
  p := new(DetachVmApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "storage.v4.config.DetachVmApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "storage.v4.r0.a3.config.DetachVmApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *DetachVmApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *DetachVmApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfDetachVmApiResponseData()
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
REST Response for all response codes in api path /storage/v4.0.a3/config/volume-groups/{volumeGroupExtId}/$actions/disassociate-category Post operation
*/
type DisassociateCategoryApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfDisassociateCategoryApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDisassociateCategoryApiResponse() *DisassociateCategoryApiResponse {
  p := new(DisassociateCategoryApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "storage.v4.config.DisassociateCategoryApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "storage.v4.r0.a3.config.DisassociateCategoryApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *DisassociateCategoryApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *DisassociateCategoryApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfDisassociateCategoryApiResponseData()
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
Storage features for the Volume disks.
*/
type DiskStorageFeatures struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  FlashMode *FlashMode `json:"flashMode,omitempty"`
}

func NewDiskStorageFeatures() *DiskStorageFeatures {
  p := new(DiskStorageFeatures)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "storage.v4.config.DiskStorageFeatures"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "storage.v4.r0.a3.config.DiskStorageFeatures"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Performance setting to avoid down migration of data from the hot tier. When specified, all the virtual disks of the volume group are pinned to the higher tier, unless overrides are specified for some of the virtual disks.
*/
type FlashMode struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  IsEnabled *bool `json:"isEnabled,omitempty"`
}

func NewFlashMode() *FlashMode {
  p := new(FlashMode)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "storage.v4.config.FlashMode"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "storage.v4.r0.a3.config.FlashMode"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST Response for all response codes in api path /storage/v4.0.a3/config/volume-groups/{volumeGroupExtId}/category-associations Get operation
*/
type GetCategoryAssociationsApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfGetCategoryAssociationsApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetCategoryAssociationsApiResponse() *GetCategoryAssociationsApiResponse {
  p := new(GetCategoryAssociationsApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "storage.v4.config.GetCategoryAssociationsApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "storage.v4.r0.a3.config.GetCategoryAssociationsApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *GetCategoryAssociationsApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *GetCategoryAssociationsApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfGetCategoryAssociationsApiResponseData()
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
REST Response for all response codes in api path /storage/v4.0.a3/config/volume-groups/{volumeGroupExtId}/iscsi-client-attachments Get operation
*/
type GetExternalAttachmentsApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfGetExternalAttachmentsApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetExternalAttachmentsApiResponse() *GetExternalAttachmentsApiResponse {
  p := new(GetExternalAttachmentsApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "storage.v4.config.GetExternalAttachmentsApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "storage.v4.r0.a3.config.GetExternalAttachmentsApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *GetExternalAttachmentsApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *GetExternalAttachmentsApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfGetExternalAttachmentsApiResponseData()
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
REST Response for all response codes in api path /storage/v4.0.a3/config/iscsi-clients/{clientId} Get operation
*/
type GetIscsiClientApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfGetIscsiClientApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetIscsiClientApiResponse() *GetIscsiClientApiResponse {
  p := new(GetIscsiClientApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "storage.v4.config.GetIscsiClientApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "storage.v4.r0.a3.config.GetIscsiClientApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *GetIscsiClientApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *GetIscsiClientApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfGetIscsiClientApiResponseData()
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
REST Response for all response codes in api path /storage/v4.0.a3/config/iscsi-clients Get operation
*/
type GetIscsiClientsApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfGetIscsiClientsApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetIscsiClientsApiResponse() *GetIscsiClientsApiResponse {
  p := new(GetIscsiClientsApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "storage.v4.config.GetIscsiClientsApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "storage.v4.r0.a3.config.GetIscsiClientsApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *GetIscsiClientsApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *GetIscsiClientsApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfGetIscsiClientsApiResponseData()
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
REST Response for all response codes in api path /storage/v4.0.a3/config/volume-groups/{volumeGroupExtId}/vm-attachments Get operation
*/
type GetVmAttachmentsApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfGetVmAttachmentsApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetVmAttachmentsApiResponse() *GetVmAttachmentsApiResponse {
  p := new(GetVmAttachmentsApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "storage.v4.config.GetVmAttachmentsApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "storage.v4.r0.a3.config.GetVmAttachmentsApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *GetVmAttachmentsApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *GetVmAttachmentsApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfGetVmAttachmentsApiResponseData()
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
REST Response for all response codes in api path /storage/v4.0.a3/config/volume-groups/{volumeGroupExtId}/disks/{diskExtId} Get operation
*/
type GetVolumeDiskApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfGetVolumeDiskApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetVolumeDiskApiResponse() *GetVolumeDiskApiResponse {
  p := new(GetVolumeDiskApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "storage.v4.config.GetVolumeDiskApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "storage.v4.r0.a3.config.GetVolumeDiskApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *GetVolumeDiskApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *GetVolumeDiskApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfGetVolumeDiskApiResponseData()
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
REST Response for all response codes in api path /storage/v4.0.a3/config/volume-groups/{volumeGroupExtId}/disks Get operation
*/
type GetVolumeDisksApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfGetVolumeDisksApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetVolumeDisksApiResponse() *GetVolumeDisksApiResponse {
  p := new(GetVolumeDisksApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "storage.v4.config.GetVolumeDisksApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "storage.v4.r0.a3.config.GetVolumeDisksApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *GetVolumeDisksApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *GetVolumeDisksApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfGetVolumeDisksApiResponseData()
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
REST Response for all response codes in api path /storage/v4.0.a3/config/volume-groups/{extId} Get operation
*/
type GetVolumeGroupApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfGetVolumeGroupApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetVolumeGroupApiResponse() *GetVolumeGroupApiResponse {
  p := new(GetVolumeGroupApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "storage.v4.config.GetVolumeGroupApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "storage.v4.r0.a3.config.GetVolumeGroupApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *GetVolumeGroupApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *GetVolumeGroupApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfGetVolumeGroupApiResponseData()
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
REST Response for all response codes in api path /storage/v4.0.a3/config/volume-groups Get operation
*/
type GetVolumeGroupsApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfGetVolumeGroupsApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetVolumeGroupsApiResponse() *GetVolumeGroupsApiResponse {
  p := new(GetVolumeGroupsApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "storage.v4.config.GetVolumeGroupsApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "storage.v4.r0.a3.config.GetVolumeGroupsApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *GetVolumeGroupsApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *GetVolumeGroupsApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfGetVolumeGroupsApiResponseData()
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
A model that represents iSCSI Client that can be associated with a volume group as an external attachment.
*/
type IscsiClient struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  iSCSI initiator Client Secret in case of CHAP authentication. This field should not be provided in case the authentication type is not set to CHAP.
  */
  ClientSecret *string `json:"clientSecret,omitempty"`
  /**
  The UUID of cluster that will host the iSCSI Client.
  */
  ClusterReference *string `json:"clusterReference,omitempty"`
  
  EnabledAuthentications *AuthenticationType `json:"enabledAuthentications,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  iSCSI Initiator Name.
  */
  IscsiInitiatorName *string `json:"iscsiInitiatorName,omitempty"`
  
  IscsiInitiatorNetworkId *import4.IPAddressOrFQDN `json:"iscsiInitiatorNetworkId,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import2.ApiLink `json:"links,omitempty"`
  
  TargetParams []TargetParam `json:"targetParams,omitempty"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
}

func NewIscsiClient() *IscsiClient {
  p := new(IscsiClient)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "storage.v4.config.IscsiClient"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "storage.v4.r0.a3.config.IscsiClient"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}





type IscsiClientProjection struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  iSCSI initiator Client Secret in case of CHAP authentication. This field should not be provided in case the authentication type is not set to CHAP.
  */
  ClientSecret *string `json:"clientSecret,omitempty"`
  /**
  The UUID of cluster that will host the iSCSI Client.
  */
  ClusterReference *string `json:"clusterReference,omitempty"`
  
  EnabledAuthentications *AuthenticationType `json:"enabledAuthentications,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  iSCSI Initiator Name.
  */
  IscsiInitiatorName *string `json:"iscsiInitiatorName,omitempty"`
  
  IscsiInitiatorNetworkId *import4.IPAddressOrFQDN `json:"iscsiInitiatorNetworkId,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import2.ApiLink `json:"links,omitempty"`
  
  TargetParams []TargetParam `json:"targetParams,omitempty"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
}

func NewIscsiClientProjection() *IscsiClientProjection {
  p := new(IscsiClientProjection)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "storage.v4.config.IscsiClientProjection"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "storage.v4.r0.a3.config.IscsiClientProjection"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST Response for all response codes in api path /storage/v4.0.a3/config/volume-groups/{volumeGroupExtId}/$actions/migrate Post operation
*/
type MigrateVolumeGroupApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfMigrateVolumeGroupApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewMigrateVolumeGroupApiResponse() *MigrateVolumeGroupApiResponse {
  p := new(MigrateVolumeGroupApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "storage.v4.config.MigrateVolumeGroupApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "storage.v4.r0.a3.config.MigrateVolumeGroupApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *MigrateVolumeGroupApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *MigrateVolumeGroupApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfMigrateVolumeGroupApiResponseData()
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
REST Response for all response codes in api path /storage/v4.0.a3/config/volume-groups/{volumeGroupExtId}/$actions/pause-synchronous-replication Post operation
*/
type PauseVolumeGroupSynchronousReplicationApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfPauseVolumeGroupSynchronousReplicationApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewPauseVolumeGroupSynchronousReplicationApiResponse() *PauseVolumeGroupSynchronousReplicationApiResponse {
  p := new(PauseVolumeGroupSynchronousReplicationApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "storage.v4.config.PauseVolumeGroupSynchronousReplicationApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "storage.v4.r0.a3.config.PauseVolumeGroupSynchronousReplicationApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *PauseVolumeGroupSynchronousReplicationApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *PauseVolumeGroupSynchronousReplicationApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfPauseVolumeGroupSynchronousReplicationApiResponseData()
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
REST Response for all response codes in api path /storage/v4.0.a3/config/volume-groups/{volumeGroupExtId}/$actions/resume-synchronous-replication Post operation
*/
type ResumeVolumeGroupSynchronousReplicationApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfResumeVolumeGroupSynchronousReplicationApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewResumeVolumeGroupSynchronousReplicationApiResponse() *ResumeVolumeGroupSynchronousReplicationApiResponse {
  p := new(ResumeVolumeGroupSynchronousReplicationApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "storage.v4.config.ResumeVolumeGroupSynchronousReplicationApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "storage.v4.r0.a3.config.ResumeVolumeGroupSynchronousReplicationApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *ResumeVolumeGroupSynchronousReplicationApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *ResumeVolumeGroupSynchronousReplicationApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfResumeVolumeGroupSynchronousReplicationApiResponseData()
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
REST Response for all response codes in api path /storage/v4.0.a3/config/volume-groups/{volumeGroupExtId}/$actions/revert Post operation
*/
type RevertVolumeGroupApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfRevertVolumeGroupApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewRevertVolumeGroupApiResponse() *RevertVolumeGroupApiResponse {
  p := new(RevertVolumeGroupApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "storage.v4.config.RevertVolumeGroupApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "storage.v4.r0.a3.config.RevertVolumeGroupApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *RevertVolumeGroupApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *RevertVolumeGroupApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfRevertVolumeGroupApiResponseData()
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
Whether the Volume Group can be shared across multiple iSCSI initiators. The mode cannot be changed from SHARED to NOT_SHARED on a volume group with multiple attachments. Similarly, a volume group cannot be associated with more than one attachment as long as it is in exclusive mode.
*/
type SharingStatus int

const(
  SHARINGSTATUS_UNKNOWN SharingStatus = 0
  SHARINGSTATUS_REDACTED SharingStatus = 1
  SHARINGSTATUS_SHARED SharingStatus = 2
  SHARINGSTATUS_NOT_SHARED SharingStatus = 3
)

// returns the name of the enum given an ordinal number
func (e *SharingStatus) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "SHARED",
    "NOT_SHARED",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *SharingStatus) index(name string) SharingStatus {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "SHARED",
    "NOT_SHARED",
  }
  for idx := range names {
    if names[idx] == name {
      return SharingStatus(idx)
    }
  }
  return SHARINGSTATUS_UNKNOWN
}

func (e *SharingStatus) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for SharingStatus:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *SharingStatus) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e SharingStatus) Ref() *SharingStatus {
  return &e
}


/**
Storage features for the Volume Groups.
*/
type StorageFeatures struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  FlashMode *FlashMode `json:"flashMode,omitempty"`
}

func NewStorageFeatures() *StorageFeatures {
  p := new(StorageFeatures)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "storage.v4.config.StorageFeatures"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "storage.v4.r0.a3.config.StorageFeatures"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
List of iSCSI targets' parameters that will be visible and accessible to the iSCSI client.
*/
type TargetParam struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Number of virtual targets to generate for the iSCSI target.
  */
  NumVirtualTargets *int `json:"numVirtualTargets,omitempty"`
}

func NewTargetParam() *TargetParam {
  p := new(TargetParam)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "storage.v4.config.TargetParam"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "storage.v4.r0.a3.config.TargetParam"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Object encapsulating Task ID return value.
*/
type Task struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  The external identifier of the task.
  */
  ExtId *string `json:"extId,omitempty"`
}

func NewTask() *Task {
  p := new(Task)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "storage.v4.config.Task"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "storage.v4.r0.a3.config.Task"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST Response for all response codes in api path /storage/v4.0.a3/config/iscsi-clients/{clientId} Patch operation
*/
type UpdateIscsiClientApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfUpdateIscsiClientApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewUpdateIscsiClientApiResponse() *UpdateIscsiClientApiResponse {
  p := new(UpdateIscsiClientApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "storage.v4.config.UpdateIscsiClientApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "storage.v4.r0.a3.config.UpdateIscsiClientApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *UpdateIscsiClientApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *UpdateIscsiClientApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfUpdateIscsiClientApiResponseData()
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
REST Response for all response codes in api path /storage/v4.0.a3/config/volume-groups/{volumeGroupExtId}/disks/{diskExtId} Patch operation
*/
type UpdateVolumeDiskApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfUpdateVolumeDiskApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewUpdateVolumeDiskApiResponse() *UpdateVolumeDiskApiResponse {
  p := new(UpdateVolumeDiskApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "storage.v4.config.UpdateVolumeDiskApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "storage.v4.r0.a3.config.UpdateVolumeDiskApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *UpdateVolumeDiskApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *UpdateVolumeDiskApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfUpdateVolumeDiskApiResponseData()
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
REST Response for all response codes in api path /storage/v4.0.a3/config/volume-groups/{extId} Patch operation
*/
type UpdateVolumeGroupApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfUpdateVolumeGroupApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewUpdateVolumeGroupApiResponse() *UpdateVolumeGroupApiResponse {
  p := new(UpdateVolumeGroupApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "storage.v4.config.UpdateVolumeGroupApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "storage.v4.r0.a3.config.UpdateVolumeGroupApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *UpdateVolumeGroupApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *UpdateVolumeGroupApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfUpdateVolumeGroupApiResponseData()
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
Expected usage type for the volume group.
*/
type UsageType int

const(
  USAGETYPE_UNKNOWN UsageType = 0
  USAGETYPE_REDACTED UsageType = 1
  USAGETYPE_USER UsageType = 2
  USAGETYPE_INTERNAL UsageType = 3
  USAGETYPE_TEMPORARY UsageType = 4
  USAGETYPE_BACKUP_TARGET UsageType = 5
)

// returns the name of the enum given an ordinal number
func (e *UsageType) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "USER",
    "INTERNAL",
    "TEMPORARY",
    "BACKUP_TARGET",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *UsageType) index(name string) UsageType {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "USER",
    "INTERNAL",
    "TEMPORARY",
    "BACKUP_TARGET",
  }
  for idx := range names {
    if names[idx] == name {
      return UsageType(idx)
    }
  }
  return USAGETYPE_UNKNOWN
}

func (e *UsageType) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for UsageType:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *UsageType) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e UsageType) Ref() *UsageType {
  return &e
}


/**
A model that represents a VM reference that can be associated with a volume group as a hypervisor attachment.
*/
type VmAttachment struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  The external identifier of the VM.
  */
  ExtId *string `json:"extId"`
  /**
  The index on the scsi bus to attach the VM to the Volume Group.
  */
  Index *int `json:"index,omitempty"`
}

func NewVmAttachment() *VmAttachment {
  p := new(VmAttachment)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "storage.v4.config.VmAttachment"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "storage.v4.r0.a3.config.VmAttachment"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
A model that represents volume disk which is associated with a volume group, and is supported by a backing file on DSF.
*/
type VolumeDisk struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Description of the Volume disk.
  */
  Description *string `json:"description,omitempty"`
  
  DiskDataSourceReference *import4.EntityReference `json:"diskDataSourceReference,omitempty"`
  /**
  Size of the disk in bytes.
  */
  DiskSizeBytes *int64 `json:"diskSizeBytes,omitempty"`
  
  DiskStorageFeatures *DiskStorageFeatures `json:"diskStorageFeatures,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  Index of the disk in a Volume Group. This field is immutable.
  */
  Index *int `json:"index,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import2.ApiLink `json:"links,omitempty"`
  /**
  Storage container on which the disk must be created.
  */
  StorageContainerId *string `json:"storageContainerId,omitempty"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
}

func NewVolumeDisk() *VolumeDisk {
  p := new(VolumeDisk)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "storage.v4.config.VolumeDisk"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "storage.v4.r0.a3.config.VolumeDisk"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
A model that represents Volume Group resources.
*/
type VolumeGroup struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  The UUID of cluster that will host the Volume Group. This is mandatory to be specified for creating a volume group on PC.
  */
  ClusterReference *string `json:"clusterReference,omitempty"`
  /**
  Service/user who created this Volume Group.
  */
  CreatedBy *string `json:"createdBy,omitempty"`
  /**
  The description of the Volume Group.
  */
  Description *string `json:"description,omitempty"`
  
  EnabledAuthentications *AuthenticationType `json:"enabledAuthentications,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  Whether the VG is meant to be hidden or not.
  */
  IsHidden *bool `json:"isHidden,omitempty"`
  /**
  iSCSI target full name. The spec should not comprise both iscsi target prefix and iscsi target name.
  */
  IscsiTargetName *string `json:"iscsiTargetName,omitempty"`
  /**
  iSCSI target prefix-name. The spec should not comprise both iscsi target prefix and iscsi target name.
  */
  IscsiTargetPrefix *string `json:"iscsiTargetPrefix,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import2.ApiLink `json:"links,omitempty"`
  /**
  Whether to enable volume group load balancing for VM attachments. This cannot be enabled if there are iSCSI client attachments already associated with the Volume Group, and vice-versa.
  */
  LoadBalanceVmAttachments *bool `json:"loadBalanceVmAttachments,omitempty"`
  /**
  The name of the Volume Group.
  */
  Name *string `json:"name,omitempty"`
  
  SharingStatus *SharingStatus `json:"sharingStatus,omitempty"`
  
  StorageFeatures *StorageFeatures `json:"storageFeatures,omitempty"`
  /**
  Target Secret in case of CHAP authentication. This field should not be provided in case the authentication type is not set to CHAP.
  */
  TargetSecret *string `json:"targetSecret,omitempty"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
  
  UsageType *UsageType `json:"usageType,omitempty"`
}

func NewVolumeGroup() *VolumeGroup {
  p := new(VolumeGroup)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "storage.v4.config.VolumeGroup"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "storage.v4.r0.a3.config.VolumeGroup"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Specification for the migrate action on the Volume Group.
*/
type VolumeGroupMigrationSpec struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Reference to the target Availability Zone where the Volume Group must be migrated.
  */
  TargetAvailabilityZoneId *string `json:"targetAvailabilityZoneId"`
  /**
  Reference to the cluster in the target Availability Zone where the Volume Group must be migrated.
  */
  TargetClusterId *string `json:"targetClusterId,omitempty"`
}

func NewVolumeGroupMigrationSpec() *VolumeGroupMigrationSpec {
  p := new(VolumeGroupMigrationSpec)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "storage.v4.config.VolumeGroupMigrationSpec"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "storage.v4.r0.a3.config.VolumeGroupMigrationSpec"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}





type VolumeGroupProjection struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  The UUID of cluster that will host the Volume Group. This is mandatory to be specified for creating a volume group on PC.
  */
  ClusterReference *string `json:"clusterReference,omitempty"`
  /**
  Service/user who created this Volume Group.
  */
  CreatedBy *string `json:"createdBy,omitempty"`
  /**
  The description of the Volume Group.
  */
  Description *string `json:"description,omitempty"`
  
  EnabledAuthentications *AuthenticationType `json:"enabledAuthentications,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  Whether the VG is meant to be hidden or not.
  */
  IsHidden *bool `json:"isHidden,omitempty"`
  /**
  iSCSI target full name. The spec should not comprise both iscsi target prefix and iscsi target name.
  */
  IscsiTargetName *string `json:"iscsiTargetName,omitempty"`
  /**
  iSCSI target prefix-name. The spec should not comprise both iscsi target prefix and iscsi target name.
  */
  IscsiTargetPrefix *string `json:"iscsiTargetPrefix,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import2.ApiLink `json:"links,omitempty"`
  /**
  Whether to enable volume group load balancing for VM attachments. This cannot be enabled if there are iSCSI client attachments already associated with the Volume Group, and vice-versa.
  */
  LoadBalanceVmAttachments *bool `json:"loadBalanceVmAttachments,omitempty"`
  /**
  The name of the Volume Group.
  */
  Name *string `json:"name,omitempty"`
  
  SharingStatus *SharingStatus `json:"sharingStatus,omitempty"`
  
  StorageFeatures *StorageFeatures `json:"storageFeatures,omitempty"`
  /**
  Target Secret in case of CHAP authentication. This field should not be provided in case the authentication type is not set to CHAP.
  */
  TargetSecret *string `json:"targetSecret,omitempty"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
  
  UsageType *UsageType `json:"usageType,omitempty"`
}

func NewVolumeGroupProjection() *VolumeGroupProjection {
  p := new(VolumeGroupProjection)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "storage.v4.config.VolumeGroupProjection"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "storage.v4.r0.a3.config.VolumeGroupProjection"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Specify the Volume Group Recovery Point ID to which the Volume Group would be reverted.
*/
type VolumeGroupRevertSpec struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  The external identifier of the Volume Group Recovery point.
  */
  VolumeGroupRecoveryPointExtId *string `json:"volumeGroupRecoveryPointExtId"`
}

func NewVolumeGroupRevertSpec() *VolumeGroupRevertSpec {
  p := new(VolumeGroupRevertSpec)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "storage.v4.config.VolumeGroupRevertSpec"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "storage.v4.r0.a3.config.VolumeGroupRevertSpec"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}



type OneOfGetVolumeDisksApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 []VolumeDisk `json:"-"`
  oneOfType400 *import1.ErrorResponse `json:"-"`
}

func NewOneOfGetVolumeDisksApiResponseData() *OneOfGetVolumeDisksApiResponseData {
  p := new(OneOfGetVolumeDisksApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfGetVolumeDisksApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfGetVolumeDisksApiResponseData is nil"))
  }
  switch v.(type) {
    case []VolumeDisk:
      p.oneOfType0 = v.([]VolumeDisk)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<storage.v4.config.VolumeDisk>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<storage.v4.config.VolumeDisk>"
    case import1.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = v.(import1.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfGetVolumeDisksApiResponseData) GetValue() interface{} {
  if "List<storage.v4.config.VolumeDisk>" == *p.Discriminator {
    return p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfGetVolumeDisksApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new([]VolumeDisk)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "storage.v4.config.VolumeDisk" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<storage.v4.config.VolumeDisk>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<storage.v4.config.VolumeDisk>"
      return nil

    }
  }
  vOneOfType400 := new(import1.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "storage.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetVolumeDisksApiResponseData"))
}

func (p *OneOfGetVolumeDisksApiResponseData) MarshalJSON() ([]byte, error) {
  if "List<storage.v4.config.VolumeDisk>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfGetVolumeDisksApiResponseData")
}

type OneOfResumeVolumeGroupSynchronousReplicationApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType1 *interface{} `json:"-"`
  oneOfType400 *import1.ErrorResponse `json:"-"`
}

func NewOneOfResumeVolumeGroupSynchronousReplicationApiResponseData() *OneOfResumeVolumeGroupSynchronousReplicationApiResponseData {
  p := new(OneOfResumeVolumeGroupSynchronousReplicationApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfResumeVolumeGroupSynchronousReplicationApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfResumeVolumeGroupSynchronousReplicationApiResponseData is nil"))
  }
  if nil == v {
    if nil == p.oneOfType1 {p.oneOfType1 = new(interface {})}
    *p.oneOfType1 = nil
    if nil == p.Discriminator {p.Discriminator = new(string)}
    *p.Discriminator = "EMPTY"
    if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
    *p.ObjectType_ = "EMPTY"
    return nil
  }
  switch v.(type) {
    case import1.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = v.(import1.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfResumeVolumeGroupSynchronousReplicationApiResponseData) GetValue() interface{} {
  if "EMPTY" == *p.Discriminator {
    return *p.oneOfType1
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfResumeVolumeGroupSynchronousReplicationApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType1 := new(interface {})
  if err := json.Unmarshal(b, vOneOfType1); err == nil {
    if nil == *vOneOfType1 {
      if nil == p.oneOfType1 {p.oneOfType1 = new(interface {})}
      *p.oneOfType1 = nil
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "EMPTY"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "EMPTY"
      return nil
    }
  }
  vOneOfType400 := new(import1.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "storage.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfResumeVolumeGroupSynchronousReplicationApiResponseData"))
}

func (p *OneOfResumeVolumeGroupSynchronousReplicationApiResponseData) MarshalJSON() ([]byte, error) {
  if "EMPTY" == *p.Discriminator {
    return json.Marshal(p.oneOfType1)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfResumeVolumeGroupSynchronousReplicationApiResponseData")
}

type OneOfAssociateCategoryApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType1 *interface{} `json:"-"`
  oneOfType400 *import1.ErrorResponse `json:"-"`
}

func NewOneOfAssociateCategoryApiResponseData() *OneOfAssociateCategoryApiResponseData {
  p := new(OneOfAssociateCategoryApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfAssociateCategoryApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfAssociateCategoryApiResponseData is nil"))
  }
  if nil == v {
    if nil == p.oneOfType1 {p.oneOfType1 = new(interface {})}
    *p.oneOfType1 = nil
    if nil == p.Discriminator {p.Discriminator = new(string)}
    *p.Discriminator = "EMPTY"
    if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
    *p.ObjectType_ = "EMPTY"
    return nil
  }
  switch v.(type) {
    case import1.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = v.(import1.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfAssociateCategoryApiResponseData) GetValue() interface{} {
  if "EMPTY" == *p.Discriminator {
    return *p.oneOfType1
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfAssociateCategoryApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType1 := new(interface {})
  if err := json.Unmarshal(b, vOneOfType1); err == nil {
    if nil == *vOneOfType1 {
      if nil == p.oneOfType1 {p.oneOfType1 = new(interface {})}
      *p.oneOfType1 = nil
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "EMPTY"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "EMPTY"
      return nil
    }
  }
  vOneOfType400 := new(import1.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "storage.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfAssociateCategoryApiResponseData"))
}

func (p *OneOfAssociateCategoryApiResponseData) MarshalJSON() ([]byte, error) {
  if "EMPTY" == *p.Discriminator {
    return json.Marshal(p.oneOfType1)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfAssociateCategoryApiResponseData")
}

type OneOfAttachVmApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *Task `json:"-"`
  oneOfType1 *import3.TaskReference `json:"-"`
  oneOfType400 *import1.ErrorResponse `json:"-"`
}

func NewOneOfAttachVmApiResponseData() *OneOfAttachVmApiResponseData {
  p := new(OneOfAttachVmApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfAttachVmApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfAttachVmApiResponseData is nil"))
  }
  switch v.(type) {
    case Task:
      if nil == p.oneOfType0 {p.oneOfType0 = new(Task)}
      *p.oneOfType0 = v.(Task)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    case import3.TaskReference:
      if nil == p.oneOfType1 {p.oneOfType1 = new(import3.TaskReference)}
      *p.oneOfType1 = v.(import3.TaskReference)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType1.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType1.ObjectType_
    case import1.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = v.(import1.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfAttachVmApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
    return *p.oneOfType1
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfAttachVmApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new(Task)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "storage.v4.config.Task" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(Task)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  vOneOfType1 := new(import3.TaskReference)
  if err := json.Unmarshal(b, vOneOfType1); err == nil {
    if "prism.v4.config.TaskReference" == *vOneOfType1.ObjectType_ {
          if nil == p.oneOfType1 {p.oneOfType1 = new(import3.TaskReference)}
      *p.oneOfType1 = *vOneOfType1
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType1.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType1.ObjectType_
      return nil
    }
    }
  vOneOfType400 := new(import1.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "storage.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfAttachVmApiResponseData"))
}

func (p *OneOfAttachVmApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType1)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfAttachVmApiResponseData")
}

type OneOfRevertVolumeGroupApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *import3.TaskReference `json:"-"`
  oneOfType400 *import1.ErrorResponse `json:"-"`
}

func NewOneOfRevertVolumeGroupApiResponseData() *OneOfRevertVolumeGroupApiResponseData {
  p := new(OneOfRevertVolumeGroupApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfRevertVolumeGroupApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfRevertVolumeGroupApiResponseData is nil"))
  }
  switch v.(type) {
    case import3.TaskReference:
      if nil == p.oneOfType0 {p.oneOfType0 = new(import3.TaskReference)}
      *p.oneOfType0 = v.(import3.TaskReference)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    case import1.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = v.(import1.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfRevertVolumeGroupApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfRevertVolumeGroupApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new(import3.TaskReference)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(import3.TaskReference)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  vOneOfType400 := new(import1.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "storage.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfRevertVolumeGroupApiResponseData"))
}

func (p *OneOfRevertVolumeGroupApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfRevertVolumeGroupApiResponseData")
}

type OneOfCreateVolumeDiskApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *Task `json:"-"`
  oneOfType1 *import3.TaskReference `json:"-"`
  oneOfType400 *import1.ErrorResponse `json:"-"`
}

func NewOneOfCreateVolumeDiskApiResponseData() *OneOfCreateVolumeDiskApiResponseData {
  p := new(OneOfCreateVolumeDiskApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfCreateVolumeDiskApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfCreateVolumeDiskApiResponseData is nil"))
  }
  switch v.(type) {
    case Task:
      if nil == p.oneOfType0 {p.oneOfType0 = new(Task)}
      *p.oneOfType0 = v.(Task)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    case import3.TaskReference:
      if nil == p.oneOfType1 {p.oneOfType1 = new(import3.TaskReference)}
      *p.oneOfType1 = v.(import3.TaskReference)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType1.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType1.ObjectType_
    case import1.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = v.(import1.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfCreateVolumeDiskApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
    return *p.oneOfType1
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfCreateVolumeDiskApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new(Task)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "storage.v4.config.Task" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(Task)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  vOneOfType1 := new(import3.TaskReference)
  if err := json.Unmarshal(b, vOneOfType1); err == nil {
    if "prism.v4.config.TaskReference" == *vOneOfType1.ObjectType_ {
          if nil == p.oneOfType1 {p.oneOfType1 = new(import3.TaskReference)}
      *p.oneOfType1 = *vOneOfType1
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType1.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType1.ObjectType_
      return nil
    }
    }
  vOneOfType400 := new(import1.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "storage.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateVolumeDiskApiResponseData"))
}

func (p *OneOfCreateVolumeDiskApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType1)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfCreateVolumeDiskApiResponseData")
}

type OneOfCreateVolumeGroupApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *Task `json:"-"`
  oneOfType1 *import3.TaskReference `json:"-"`
  oneOfType400 *import1.ErrorResponse `json:"-"`
}

func NewOneOfCreateVolumeGroupApiResponseData() *OneOfCreateVolumeGroupApiResponseData {
  p := new(OneOfCreateVolumeGroupApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfCreateVolumeGroupApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfCreateVolumeGroupApiResponseData is nil"))
  }
  switch v.(type) {
    case Task:
      if nil == p.oneOfType0 {p.oneOfType0 = new(Task)}
      *p.oneOfType0 = v.(Task)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    case import3.TaskReference:
      if nil == p.oneOfType1 {p.oneOfType1 = new(import3.TaskReference)}
      *p.oneOfType1 = v.(import3.TaskReference)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType1.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType1.ObjectType_
    case import1.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = v.(import1.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfCreateVolumeGroupApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
    return *p.oneOfType1
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfCreateVolumeGroupApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new(Task)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "storage.v4.config.Task" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(Task)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  vOneOfType1 := new(import3.TaskReference)
  if err := json.Unmarshal(b, vOneOfType1); err == nil {
    if "prism.v4.config.TaskReference" == *vOneOfType1.ObjectType_ {
          if nil == p.oneOfType1 {p.oneOfType1 = new(import3.TaskReference)}
      *p.oneOfType1 = *vOneOfType1
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType1.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType1.ObjectType_
      return nil
    }
    }
  vOneOfType400 := new(import1.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "storage.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateVolumeGroupApiResponseData"))
}

func (p *OneOfCreateVolumeGroupApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType1)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfCreateVolumeGroupApiResponseData")
}

type OneOfGetVolumeGroupApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *VolumeGroup `json:"-"`
  oneOfType400 *import1.ErrorResponse `json:"-"`
}

func NewOneOfGetVolumeGroupApiResponseData() *OneOfGetVolumeGroupApiResponseData {
  p := new(OneOfGetVolumeGroupApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfGetVolumeGroupApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfGetVolumeGroupApiResponseData is nil"))
  }
  switch v.(type) {
    case VolumeGroup:
      if nil == p.oneOfType0 {p.oneOfType0 = new(VolumeGroup)}
      *p.oneOfType0 = v.(VolumeGroup)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    case import1.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = v.(import1.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfGetVolumeGroupApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfGetVolumeGroupApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new(VolumeGroup)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "storage.v4.config.VolumeGroup" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(VolumeGroup)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  vOneOfType400 := new(import1.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "storage.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetVolumeGroupApiResponseData"))
}

func (p *OneOfGetVolumeGroupApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfGetVolumeGroupApiResponseData")
}

type OneOfAttachIscsiClientApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *Task `json:"-"`
  oneOfType1 *import3.TaskReference `json:"-"`
  oneOfType400 *import1.ErrorResponse `json:"-"`
}

func NewOneOfAttachIscsiClientApiResponseData() *OneOfAttachIscsiClientApiResponseData {
  p := new(OneOfAttachIscsiClientApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfAttachIscsiClientApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfAttachIscsiClientApiResponseData is nil"))
  }
  switch v.(type) {
    case Task:
      if nil == p.oneOfType0 {p.oneOfType0 = new(Task)}
      *p.oneOfType0 = v.(Task)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    case import3.TaskReference:
      if nil == p.oneOfType1 {p.oneOfType1 = new(import3.TaskReference)}
      *p.oneOfType1 = v.(import3.TaskReference)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType1.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType1.ObjectType_
    case import1.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = v.(import1.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfAttachIscsiClientApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
    return *p.oneOfType1
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfAttachIscsiClientApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new(Task)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "storage.v4.config.Task" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(Task)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  vOneOfType1 := new(import3.TaskReference)
  if err := json.Unmarshal(b, vOneOfType1); err == nil {
    if "prism.v4.config.TaskReference" == *vOneOfType1.ObjectType_ {
          if nil == p.oneOfType1 {p.oneOfType1 = new(import3.TaskReference)}
      *p.oneOfType1 = *vOneOfType1
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType1.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType1.ObjectType_
      return nil
    }
    }
  vOneOfType400 := new(import1.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "storage.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfAttachIscsiClientApiResponseData"))
}

func (p *OneOfAttachIscsiClientApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType1)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfAttachIscsiClientApiResponseData")
}

type OneOfPauseVolumeGroupSynchronousReplicationApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *import3.TaskReference `json:"-"`
  oneOfType400 *import1.ErrorResponse `json:"-"`
}

func NewOneOfPauseVolumeGroupSynchronousReplicationApiResponseData() *OneOfPauseVolumeGroupSynchronousReplicationApiResponseData {
  p := new(OneOfPauseVolumeGroupSynchronousReplicationApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfPauseVolumeGroupSynchronousReplicationApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfPauseVolumeGroupSynchronousReplicationApiResponseData is nil"))
  }
  switch v.(type) {
    case import3.TaskReference:
      if nil == p.oneOfType0 {p.oneOfType0 = new(import3.TaskReference)}
      *p.oneOfType0 = v.(import3.TaskReference)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    case import1.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = v.(import1.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfPauseVolumeGroupSynchronousReplicationApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfPauseVolumeGroupSynchronousReplicationApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new(import3.TaskReference)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(import3.TaskReference)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  vOneOfType400 := new(import1.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "storage.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfPauseVolumeGroupSynchronousReplicationApiResponseData"))
}

func (p *OneOfPauseVolumeGroupSynchronousReplicationApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfPauseVolumeGroupSynchronousReplicationApiResponseData")
}

type OneOfDetachIscsiClientApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *Task `json:"-"`
  oneOfType1 *import3.TaskReference `json:"-"`
  oneOfType400 *import1.ErrorResponse `json:"-"`
}

func NewOneOfDetachIscsiClientApiResponseData() *OneOfDetachIscsiClientApiResponseData {
  p := new(OneOfDetachIscsiClientApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfDetachIscsiClientApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfDetachIscsiClientApiResponseData is nil"))
  }
  switch v.(type) {
    case Task:
      if nil == p.oneOfType0 {p.oneOfType0 = new(Task)}
      *p.oneOfType0 = v.(Task)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    case import3.TaskReference:
      if nil == p.oneOfType1 {p.oneOfType1 = new(import3.TaskReference)}
      *p.oneOfType1 = v.(import3.TaskReference)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType1.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType1.ObjectType_
    case import1.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = v.(import1.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfDetachIscsiClientApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
    return *p.oneOfType1
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfDetachIscsiClientApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new(Task)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "storage.v4.config.Task" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(Task)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  vOneOfType1 := new(import3.TaskReference)
  if err := json.Unmarshal(b, vOneOfType1); err == nil {
    if "prism.v4.config.TaskReference" == *vOneOfType1.ObjectType_ {
          if nil == p.oneOfType1 {p.oneOfType1 = new(import3.TaskReference)}
      *p.oneOfType1 = *vOneOfType1
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType1.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType1.ObjectType_
      return nil
    }
    }
  vOneOfType400 := new(import1.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "storage.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDetachIscsiClientApiResponseData"))
}

func (p *OneOfDetachIscsiClientApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType1)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfDetachIscsiClientApiResponseData")
}

type OneOfUpdateVolumeDiskApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *Task `json:"-"`
  oneOfType1 *import3.TaskReference `json:"-"`
  oneOfType400 *import1.ErrorResponse `json:"-"`
}

func NewOneOfUpdateVolumeDiskApiResponseData() *OneOfUpdateVolumeDiskApiResponseData {
  p := new(OneOfUpdateVolumeDiskApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfUpdateVolumeDiskApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfUpdateVolumeDiskApiResponseData is nil"))
  }
  switch v.(type) {
    case Task:
      if nil == p.oneOfType0 {p.oneOfType0 = new(Task)}
      *p.oneOfType0 = v.(Task)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    case import3.TaskReference:
      if nil == p.oneOfType1 {p.oneOfType1 = new(import3.TaskReference)}
      *p.oneOfType1 = v.(import3.TaskReference)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType1.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType1.ObjectType_
    case import1.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = v.(import1.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfUpdateVolumeDiskApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
    return *p.oneOfType1
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfUpdateVolumeDiskApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new(Task)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "storage.v4.config.Task" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(Task)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  vOneOfType1 := new(import3.TaskReference)
  if err := json.Unmarshal(b, vOneOfType1); err == nil {
    if "prism.v4.config.TaskReference" == *vOneOfType1.ObjectType_ {
          if nil == p.oneOfType1 {p.oneOfType1 = new(import3.TaskReference)}
      *p.oneOfType1 = *vOneOfType1
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType1.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType1.ObjectType_
      return nil
    }
    }
  vOneOfType400 := new(import1.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "storage.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateVolumeDiskApiResponseData"))
}

func (p *OneOfUpdateVolumeDiskApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType1)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfUpdateVolumeDiskApiResponseData")
}

type OneOfGetVolumeDiskApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *VolumeDisk `json:"-"`
  oneOfType400 *import1.ErrorResponse `json:"-"`
}

func NewOneOfGetVolumeDiskApiResponseData() *OneOfGetVolumeDiskApiResponseData {
  p := new(OneOfGetVolumeDiskApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfGetVolumeDiskApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfGetVolumeDiskApiResponseData is nil"))
  }
  switch v.(type) {
    case VolumeDisk:
      if nil == p.oneOfType0 {p.oneOfType0 = new(VolumeDisk)}
      *p.oneOfType0 = v.(VolumeDisk)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    case import1.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = v.(import1.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfGetVolumeDiskApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfGetVolumeDiskApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new(VolumeDisk)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "storage.v4.config.VolumeDisk" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(VolumeDisk)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  vOneOfType400 := new(import1.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "storage.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetVolumeDiskApiResponseData"))
}

func (p *OneOfGetVolumeDiskApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfGetVolumeDiskApiResponseData")
}

type OneOfUpdateVolumeGroupApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *Task `json:"-"`
  oneOfType1 *import3.TaskReference `json:"-"`
  oneOfType400 *import1.ErrorResponse `json:"-"`
}

func NewOneOfUpdateVolumeGroupApiResponseData() *OneOfUpdateVolumeGroupApiResponseData {
  p := new(OneOfUpdateVolumeGroupApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfUpdateVolumeGroupApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfUpdateVolumeGroupApiResponseData is nil"))
  }
  switch v.(type) {
    case Task:
      if nil == p.oneOfType0 {p.oneOfType0 = new(Task)}
      *p.oneOfType0 = v.(Task)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    case import3.TaskReference:
      if nil == p.oneOfType1 {p.oneOfType1 = new(import3.TaskReference)}
      *p.oneOfType1 = v.(import3.TaskReference)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType1.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType1.ObjectType_
    case import1.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = v.(import1.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfUpdateVolumeGroupApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
    return *p.oneOfType1
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfUpdateVolumeGroupApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new(Task)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "storage.v4.config.Task" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(Task)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  vOneOfType1 := new(import3.TaskReference)
  if err := json.Unmarshal(b, vOneOfType1); err == nil {
    if "prism.v4.config.TaskReference" == *vOneOfType1.ObjectType_ {
          if nil == p.oneOfType1 {p.oneOfType1 = new(import3.TaskReference)}
      *p.oneOfType1 = *vOneOfType1
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType1.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType1.ObjectType_
      return nil
    }
    }
  vOneOfType400 := new(import1.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "storage.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateVolumeGroupApiResponseData"))
}

func (p *OneOfUpdateVolumeGroupApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType1)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfUpdateVolumeGroupApiResponseData")
}

type OneOfDeleteVolumeGroupApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *Task `json:"-"`
  oneOfType1 *import3.TaskReference `json:"-"`
  oneOfType400 *import1.ErrorResponse `json:"-"`
}

func NewOneOfDeleteVolumeGroupApiResponseData() *OneOfDeleteVolumeGroupApiResponseData {
  p := new(OneOfDeleteVolumeGroupApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfDeleteVolumeGroupApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfDeleteVolumeGroupApiResponseData is nil"))
  }
  switch v.(type) {
    case Task:
      if nil == p.oneOfType0 {p.oneOfType0 = new(Task)}
      *p.oneOfType0 = v.(Task)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    case import3.TaskReference:
      if nil == p.oneOfType1 {p.oneOfType1 = new(import3.TaskReference)}
      *p.oneOfType1 = v.(import3.TaskReference)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType1.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType1.ObjectType_
    case import1.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = v.(import1.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfDeleteVolumeGroupApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
    return *p.oneOfType1
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfDeleteVolumeGroupApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new(Task)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "storage.v4.config.Task" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(Task)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  vOneOfType1 := new(import3.TaskReference)
  if err := json.Unmarshal(b, vOneOfType1); err == nil {
    if "prism.v4.config.TaskReference" == *vOneOfType1.ObjectType_ {
          if nil == p.oneOfType1 {p.oneOfType1 = new(import3.TaskReference)}
      *p.oneOfType1 = *vOneOfType1
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType1.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType1.ObjectType_
      return nil
    }
    }
  vOneOfType400 := new(import1.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "storage.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteVolumeGroupApiResponseData"))
}

func (p *OneOfDeleteVolumeGroupApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType1)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfDeleteVolumeGroupApiResponseData")
}

type OneOfGetVolumeGroupsApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 []VolumeGroup `json:"-"`
  oneOfType401 []VolumeGroupProjection `json:"-"`
  oneOfType400 *import1.ErrorResponse `json:"-"`
}

func NewOneOfGetVolumeGroupsApiResponseData() *OneOfGetVolumeGroupsApiResponseData {
  p := new(OneOfGetVolumeGroupsApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfGetVolumeGroupsApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfGetVolumeGroupsApiResponseData is nil"))
  }
  switch v.(type) {
    case []VolumeGroup:
      p.oneOfType0 = v.([]VolumeGroup)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<storage.v4.config.VolumeGroup>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<storage.v4.config.VolumeGroup>"
    case []VolumeGroupProjection:
      p.oneOfType401 = v.([]VolumeGroupProjection)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<storage.v4.config.VolumeGroupProjection>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<storage.v4.config.VolumeGroupProjection>"
    case import1.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = v.(import1.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfGetVolumeGroupsApiResponseData) GetValue() interface{} {
  if "List<storage.v4.config.VolumeGroup>" == *p.Discriminator {
    return p.oneOfType0
  }
  if "List<storage.v4.config.VolumeGroupProjection>" == *p.Discriminator {
    return p.oneOfType401
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfGetVolumeGroupsApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new([]VolumeGroup)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "storage.v4.config.VolumeGroup" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<storage.v4.config.VolumeGroup>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<storage.v4.config.VolumeGroup>"
      return nil

    }
  }
  vOneOfType401 := new([]VolumeGroupProjection)
  if err := json.Unmarshal(b, vOneOfType401); err == nil {
    if len(*vOneOfType401) == 0 || "storage.v4.config.VolumeGroupProjection" == *((*vOneOfType401)[0].ObjectType_) {
      p.oneOfType401 = *vOneOfType401
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<storage.v4.config.VolumeGroupProjection>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<storage.v4.config.VolumeGroupProjection>"
      return nil

    }
  }
  vOneOfType400 := new(import1.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "storage.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetVolumeGroupsApiResponseData"))
}

func (p *OneOfGetVolumeGroupsApiResponseData) MarshalJSON() ([]byte, error) {
  if "List<storage.v4.config.VolumeGroup>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if "List<storage.v4.config.VolumeGroupProjection>" == *p.Discriminator {
    return json.Marshal(p.oneOfType401)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfGetVolumeGroupsApiResponseData")
}

type OneOfGetIscsiClientsApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType401 []IscsiClientProjection `json:"-"`
  oneOfType0 []IscsiClient `json:"-"`
  oneOfType400 *import1.ErrorResponse `json:"-"`
}

func NewOneOfGetIscsiClientsApiResponseData() *OneOfGetIscsiClientsApiResponseData {
  p := new(OneOfGetIscsiClientsApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfGetIscsiClientsApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfGetIscsiClientsApiResponseData is nil"))
  }
  switch v.(type) {
    case []IscsiClientProjection:
      p.oneOfType401 = v.([]IscsiClientProjection)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<storage.v4.config.IscsiClientProjection>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<storage.v4.config.IscsiClientProjection>"
    case []IscsiClient:
      p.oneOfType0 = v.([]IscsiClient)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<storage.v4.config.IscsiClient>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<storage.v4.config.IscsiClient>"
    case import1.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = v.(import1.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfGetIscsiClientsApiResponseData) GetValue() interface{} {
  if "List<storage.v4.config.IscsiClientProjection>" == *p.Discriminator {
    return p.oneOfType401
  }
  if "List<storage.v4.config.IscsiClient>" == *p.Discriminator {
    return p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfGetIscsiClientsApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType401 := new([]IscsiClientProjection)
  if err := json.Unmarshal(b, vOneOfType401); err == nil {
    if len(*vOneOfType401) == 0 || "storage.v4.config.IscsiClientProjection" == *((*vOneOfType401)[0].ObjectType_) {
      p.oneOfType401 = *vOneOfType401
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<storage.v4.config.IscsiClientProjection>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<storage.v4.config.IscsiClientProjection>"
      return nil

    }
  }
  vOneOfType0 := new([]IscsiClient)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "storage.v4.config.IscsiClient" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<storage.v4.config.IscsiClient>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<storage.v4.config.IscsiClient>"
      return nil

    }
  }
  vOneOfType400 := new(import1.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "storage.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetIscsiClientsApiResponseData"))
}

func (p *OneOfGetIscsiClientsApiResponseData) MarshalJSON() ([]byte, error) {
  if "List<storage.v4.config.IscsiClientProjection>" == *p.Discriminator {
    return json.Marshal(p.oneOfType401)
  }
  if "List<storage.v4.config.IscsiClient>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfGetIscsiClientsApiResponseData")
}

type OneOfDeleteVolumeDiskApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *Task `json:"-"`
  oneOfType1 *import3.TaskReference `json:"-"`
  oneOfType400 *import1.ErrorResponse `json:"-"`
}

func NewOneOfDeleteVolumeDiskApiResponseData() *OneOfDeleteVolumeDiskApiResponseData {
  p := new(OneOfDeleteVolumeDiskApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfDeleteVolumeDiskApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfDeleteVolumeDiskApiResponseData is nil"))
  }
  switch v.(type) {
    case Task:
      if nil == p.oneOfType0 {p.oneOfType0 = new(Task)}
      *p.oneOfType0 = v.(Task)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    case import3.TaskReference:
      if nil == p.oneOfType1 {p.oneOfType1 = new(import3.TaskReference)}
      *p.oneOfType1 = v.(import3.TaskReference)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType1.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType1.ObjectType_
    case import1.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = v.(import1.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfDeleteVolumeDiskApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
    return *p.oneOfType1
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfDeleteVolumeDiskApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new(Task)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "storage.v4.config.Task" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(Task)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  vOneOfType1 := new(import3.TaskReference)
  if err := json.Unmarshal(b, vOneOfType1); err == nil {
    if "prism.v4.config.TaskReference" == *vOneOfType1.ObjectType_ {
          if nil == p.oneOfType1 {p.oneOfType1 = new(import3.TaskReference)}
      *p.oneOfType1 = *vOneOfType1
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType1.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType1.ObjectType_
      return nil
    }
    }
  vOneOfType400 := new(import1.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "storage.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteVolumeDiskApiResponseData"))
}

func (p *OneOfDeleteVolumeDiskApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType1)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfDeleteVolumeDiskApiResponseData")
}

type OneOfGetIscsiClientApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *IscsiClient `json:"-"`
  oneOfType400 *import1.ErrorResponse `json:"-"`
}

func NewOneOfGetIscsiClientApiResponseData() *OneOfGetIscsiClientApiResponseData {
  p := new(OneOfGetIscsiClientApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfGetIscsiClientApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfGetIscsiClientApiResponseData is nil"))
  }
  switch v.(type) {
    case IscsiClient:
      if nil == p.oneOfType0 {p.oneOfType0 = new(IscsiClient)}
      *p.oneOfType0 = v.(IscsiClient)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    case import1.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = v.(import1.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfGetIscsiClientApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfGetIscsiClientApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new(IscsiClient)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "storage.v4.config.IscsiClient" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(IscsiClient)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  vOneOfType400 := new(import1.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "storage.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetIscsiClientApiResponseData"))
}

func (p *OneOfGetIscsiClientApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfGetIscsiClientApiResponseData")
}

type OneOfMigrateVolumeGroupApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *Task `json:"-"`
  oneOfType1 *import3.TaskReference `json:"-"`
  oneOfType400 *import1.ErrorResponse `json:"-"`
}

func NewOneOfMigrateVolumeGroupApiResponseData() *OneOfMigrateVolumeGroupApiResponseData {
  p := new(OneOfMigrateVolumeGroupApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfMigrateVolumeGroupApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfMigrateVolumeGroupApiResponseData is nil"))
  }
  switch v.(type) {
    case Task:
      if nil == p.oneOfType0 {p.oneOfType0 = new(Task)}
      *p.oneOfType0 = v.(Task)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    case import3.TaskReference:
      if nil == p.oneOfType1 {p.oneOfType1 = new(import3.TaskReference)}
      *p.oneOfType1 = v.(import3.TaskReference)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType1.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType1.ObjectType_
    case import1.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = v.(import1.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfMigrateVolumeGroupApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
    return *p.oneOfType1
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfMigrateVolumeGroupApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new(Task)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "storage.v4.config.Task" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(Task)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  vOneOfType1 := new(import3.TaskReference)
  if err := json.Unmarshal(b, vOneOfType1); err == nil {
    if "prism.v4.config.TaskReference" == *vOneOfType1.ObjectType_ {
          if nil == p.oneOfType1 {p.oneOfType1 = new(import3.TaskReference)}
      *p.oneOfType1 = *vOneOfType1
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType1.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType1.ObjectType_
      return nil
    }
    }
  vOneOfType400 := new(import1.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "storage.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfMigrateVolumeGroupApiResponseData"))
}

func (p *OneOfMigrateVolumeGroupApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType1)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfMigrateVolumeGroupApiResponseData")
}

type OneOfDetachVmApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *Task `json:"-"`
  oneOfType1 *import3.TaskReference `json:"-"`
  oneOfType400 *import1.ErrorResponse `json:"-"`
}

func NewOneOfDetachVmApiResponseData() *OneOfDetachVmApiResponseData {
  p := new(OneOfDetachVmApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfDetachVmApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfDetachVmApiResponseData is nil"))
  }
  switch v.(type) {
    case Task:
      if nil == p.oneOfType0 {p.oneOfType0 = new(Task)}
      *p.oneOfType0 = v.(Task)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    case import3.TaskReference:
      if nil == p.oneOfType1 {p.oneOfType1 = new(import3.TaskReference)}
      *p.oneOfType1 = v.(import3.TaskReference)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType1.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType1.ObjectType_
    case import1.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = v.(import1.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfDetachVmApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
    return *p.oneOfType1
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfDetachVmApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new(Task)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "storage.v4.config.Task" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(Task)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  vOneOfType1 := new(import3.TaskReference)
  if err := json.Unmarshal(b, vOneOfType1); err == nil {
    if "prism.v4.config.TaskReference" == *vOneOfType1.ObjectType_ {
          if nil == p.oneOfType1 {p.oneOfType1 = new(import3.TaskReference)}
      *p.oneOfType1 = *vOneOfType1
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType1.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType1.ObjectType_
      return nil
    }
    }
  vOneOfType400 := new(import1.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "storage.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDetachVmApiResponseData"))
}

func (p *OneOfDetachVmApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType1)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfDetachVmApiResponseData")
}

type OneOfGetCategoryAssociationsApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 []CategoryDetails `json:"-"`
  oneOfType400 *import1.ErrorResponse `json:"-"`
}

func NewOneOfGetCategoryAssociationsApiResponseData() *OneOfGetCategoryAssociationsApiResponseData {
  p := new(OneOfGetCategoryAssociationsApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfGetCategoryAssociationsApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfGetCategoryAssociationsApiResponseData is nil"))
  }
  switch v.(type) {
    case []CategoryDetails:
      p.oneOfType0 = v.([]CategoryDetails)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<storage.v4.config.CategoryDetails>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<storage.v4.config.CategoryDetails>"
    case import1.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = v.(import1.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfGetCategoryAssociationsApiResponseData) GetValue() interface{} {
  if "List<storage.v4.config.CategoryDetails>" == *p.Discriminator {
    return p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfGetCategoryAssociationsApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new([]CategoryDetails)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "storage.v4.config.CategoryDetails" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<storage.v4.config.CategoryDetails>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<storage.v4.config.CategoryDetails>"
      return nil

    }
  }
  vOneOfType400 := new(import1.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "storage.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetCategoryAssociationsApiResponseData"))
}

func (p *OneOfGetCategoryAssociationsApiResponseData) MarshalJSON() ([]byte, error) {
  if "List<storage.v4.config.CategoryDetails>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfGetCategoryAssociationsApiResponseData")
}

type OneOfUpdateIscsiClientApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *Task `json:"-"`
  oneOfType1 *import3.TaskReference `json:"-"`
  oneOfType400 *import1.ErrorResponse `json:"-"`
}

func NewOneOfUpdateIscsiClientApiResponseData() *OneOfUpdateIscsiClientApiResponseData {
  p := new(OneOfUpdateIscsiClientApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfUpdateIscsiClientApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfUpdateIscsiClientApiResponseData is nil"))
  }
  switch v.(type) {
    case Task:
      if nil == p.oneOfType0 {p.oneOfType0 = new(Task)}
      *p.oneOfType0 = v.(Task)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    case import3.TaskReference:
      if nil == p.oneOfType1 {p.oneOfType1 = new(import3.TaskReference)}
      *p.oneOfType1 = v.(import3.TaskReference)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType1.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType1.ObjectType_
    case import1.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = v.(import1.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfUpdateIscsiClientApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
    return *p.oneOfType1
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfUpdateIscsiClientApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new(Task)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "storage.v4.config.Task" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(Task)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  vOneOfType1 := new(import3.TaskReference)
  if err := json.Unmarshal(b, vOneOfType1); err == nil {
    if "prism.v4.config.TaskReference" == *vOneOfType1.ObjectType_ {
          if nil == p.oneOfType1 {p.oneOfType1 = new(import3.TaskReference)}
      *p.oneOfType1 = *vOneOfType1
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType1.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType1.ObjectType_
      return nil
    }
    }
  vOneOfType400 := new(import1.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "storage.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateIscsiClientApiResponseData"))
}

func (p *OneOfUpdateIscsiClientApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType1)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfUpdateIscsiClientApiResponseData")
}

type OneOfDisassociateCategoryApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType1 *interface{} `json:"-"`
  oneOfType400 *import1.ErrorResponse `json:"-"`
}

func NewOneOfDisassociateCategoryApiResponseData() *OneOfDisassociateCategoryApiResponseData {
  p := new(OneOfDisassociateCategoryApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfDisassociateCategoryApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfDisassociateCategoryApiResponseData is nil"))
  }
  if nil == v {
    if nil == p.oneOfType1 {p.oneOfType1 = new(interface {})}
    *p.oneOfType1 = nil
    if nil == p.Discriminator {p.Discriminator = new(string)}
    *p.Discriminator = "EMPTY"
    if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
    *p.ObjectType_ = "EMPTY"
    return nil
  }
  switch v.(type) {
    case import1.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = v.(import1.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfDisassociateCategoryApiResponseData) GetValue() interface{} {
  if "EMPTY" == *p.Discriminator {
    return *p.oneOfType1
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfDisassociateCategoryApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType1 := new(interface {})
  if err := json.Unmarshal(b, vOneOfType1); err == nil {
    if nil == *vOneOfType1 {
      if nil == p.oneOfType1 {p.oneOfType1 = new(interface {})}
      *p.oneOfType1 = nil
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "EMPTY"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "EMPTY"
      return nil
    }
  }
  vOneOfType400 := new(import1.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "storage.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDisassociateCategoryApiResponseData"))
}

func (p *OneOfDisassociateCategoryApiResponseData) MarshalJSON() ([]byte, error) {
  if "EMPTY" == *p.Discriminator {
    return json.Marshal(p.oneOfType1)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfDisassociateCategoryApiResponseData")
}

type OneOfGetExternalAttachmentsApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType401 []IscsiClientProjection `json:"-"`
  oneOfType0 []IscsiClient `json:"-"`
  oneOfType400 *import1.ErrorResponse `json:"-"`
}

func NewOneOfGetExternalAttachmentsApiResponseData() *OneOfGetExternalAttachmentsApiResponseData {
  p := new(OneOfGetExternalAttachmentsApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfGetExternalAttachmentsApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfGetExternalAttachmentsApiResponseData is nil"))
  }
  switch v.(type) {
    case []IscsiClientProjection:
      p.oneOfType401 = v.([]IscsiClientProjection)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<storage.v4.config.IscsiClientProjection>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<storage.v4.config.IscsiClientProjection>"
    case []IscsiClient:
      p.oneOfType0 = v.([]IscsiClient)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<storage.v4.config.IscsiClient>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<storage.v4.config.IscsiClient>"
    case import1.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = v.(import1.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfGetExternalAttachmentsApiResponseData) GetValue() interface{} {
  if "List<storage.v4.config.IscsiClientProjection>" == *p.Discriminator {
    return p.oneOfType401
  }
  if "List<storage.v4.config.IscsiClient>" == *p.Discriminator {
    return p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfGetExternalAttachmentsApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType401 := new([]IscsiClientProjection)
  if err := json.Unmarshal(b, vOneOfType401); err == nil {
    if len(*vOneOfType401) == 0 || "storage.v4.config.IscsiClientProjection" == *((*vOneOfType401)[0].ObjectType_) {
      p.oneOfType401 = *vOneOfType401
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<storage.v4.config.IscsiClientProjection>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<storage.v4.config.IscsiClientProjection>"
      return nil

    }
  }
  vOneOfType0 := new([]IscsiClient)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "storage.v4.config.IscsiClient" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<storage.v4.config.IscsiClient>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<storage.v4.config.IscsiClient>"
      return nil

    }
  }
  vOneOfType400 := new(import1.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "storage.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetExternalAttachmentsApiResponseData"))
}

func (p *OneOfGetExternalAttachmentsApiResponseData) MarshalJSON() ([]byte, error) {
  if "List<storage.v4.config.IscsiClientProjection>" == *p.Discriminator {
    return json.Marshal(p.oneOfType401)
  }
  if "List<storage.v4.config.IscsiClient>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfGetExternalAttachmentsApiResponseData")
}

type OneOfGetVmAttachmentsApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 []VmAttachment `json:"-"`
  oneOfType400 *import1.ErrorResponse `json:"-"`
}

func NewOneOfGetVmAttachmentsApiResponseData() *OneOfGetVmAttachmentsApiResponseData {
  p := new(OneOfGetVmAttachmentsApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfGetVmAttachmentsApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfGetVmAttachmentsApiResponseData is nil"))
  }
  switch v.(type) {
    case []VmAttachment:
      p.oneOfType0 = v.([]VmAttachment)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<storage.v4.config.VmAttachment>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<storage.v4.config.VmAttachment>"
    case import1.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = v.(import1.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfGetVmAttachmentsApiResponseData) GetValue() interface{} {
  if "List<storage.v4.config.VmAttachment>" == *p.Discriminator {
    return p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfGetVmAttachmentsApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new([]VmAttachment)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "storage.v4.config.VmAttachment" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<storage.v4.config.VmAttachment>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<storage.v4.config.VmAttachment>"
      return nil

    }
  }
  vOneOfType400 := new(import1.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "storage.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetVmAttachmentsApiResponseData"))
}

func (p *OneOfGetVmAttachmentsApiResponseData) MarshalJSON() ([]byte, error) {
  if "List<storage.v4.config.VmAttachment>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfGetVmAttachmentsApiResponseData")
}

