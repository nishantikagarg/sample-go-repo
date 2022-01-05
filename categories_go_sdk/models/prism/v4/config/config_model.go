/*
 * Generated file models/prism/v4/config/config_model.go.
 *
 * Product version: 16.2.0-SNAPSHOT
 *
 * Part of the Category API definitions
 *
 * (c) 2022 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module prism.v4.config of Category API definitions
*/
package config
import (
  "bytes"
  import3 "github.com/nishantikagarg/sample-go-repo/categories_go_sdk/v16/models/common/v1/config"
  import2 "github.com/nishantikagarg/sample-go-repo/categories_go_sdk/v16/models/common/v1/response"
  "encoding/json"
  "errors"
  "fmt"
  import1 "github.com/nishantikagarg/sample-go-repo/categories_go_sdk/v16/models/prism/v4/error"
)

/**
This attribute contains the list of entities which have been assigned the given category.<br>
These entities are grouped by entity types (like VM or HOST) or policy types (like PROTECTION_POLICY or
NGT_POLICY).<br>
Each associations object contains the total entities belonging to the given entity type, category extId, and
references (for example for VM it'd be VM uuid).
*/
type AssociationDetail struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  External identifier for the given category.
  */
  CategoryId *string `json:"categoryId,omitempty"`
  /**
  There are three category types:
- SYSTEM: Defined by an administrator. Cannot be modified by the user.
- INTERNAL: Defined by any internal process within Nutanix ecosystem. Cannot be viewed directly through a list call.
            Can be modified by the user.
- USER: Defined by a user through the v4 api
Currently v4 apis allow the creation of only User Category Type.
  */
  Count *int `json:"count,omitempty"`
  
  ResourceGroup *ResourceGroup `json:"resourceGroup,omitempty"`
  
  ResourceReferences []Reference `json:"resourceReferences,omitempty"`
  
  ResourceType *ResourceType `json:"resourceType,omitempty"`
}

func NewAssociationDetail() *AssociationDetail {
  p := new(AssociationDetail)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "prism.v4.config.AssociationDetail"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a1.config.AssociationDetail"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST Response for all response codes in api path /prism/v4.0.a1/config/categories/{extId}/associations Get operation
*/
type AssociationDetailApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfAssociationDetailApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewAssociationDetailApiResponse() *AssociationDetailApiResponse {
  p := new(AssociationDetailApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "prism.v4.config.AssociationDetailApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a1.config.AssociationDetailApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *AssociationDetailApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *AssociationDetailApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfAssociationDetailApiResponseData()
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




type AssociationDetailProjection struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  External identifier for the given category.
  */
  CategoryId *string `json:"categoryId,omitempty"`
  /**
  There are three category types:
- SYSTEM: Defined by an administrator. Cannot be modified by the user.
- INTERNAL: Defined by any internal process within Nutanix ecosystem. Cannot be viewed directly through a list call.
            Can be modified by the user.
- USER: Defined by a user through the v4 api
Currently v4 apis allow the creation of only User Category Type.
  */
  Count *int `json:"count,omitempty"`
  
  ResourceGroup *ResourceGroup `json:"resourceGroup,omitempty"`
  
  ResourceReferences []Reference `json:"resourceReferences,omitempty"`
  
  ResourceType *ResourceType `json:"resourceType,omitempty"`
}

func NewAssociationDetailProjection() *AssociationDetailProjection {
  p := new(AssociationDetailProjection)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "prism.v4.config.AssociationDetailProjection"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a1.config.AssociationDetailProjection"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
There are three category types:
- SYSTEM: Defined by an administrator. Cannot be modified by the user.
- INTERNAL: Defined by any internal process within Nutanix ecosystem. Cannot be viewed directly through a list call.
            Can be modified by the user.
- USER: Defined by a user through the v4 api
Currently v4 apis allow the creation of only User Category Type.
*/
type Category struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  Associations []CategoryAssociationSummary `json:"associations,omitempty"`
  
  ChildCategories []CategorySummary `json:"childCategories,omitempty"`
  /**
  A string consisting of the description of the category as defined by the user.

The server does not validate this value nor does it enforce the uniqueness or any other constraints.<br>
It is the responsibility of the user to ensure that any semantic or syntactic constraints are retained when mutating
this field.
  */
  Description *string `json:"description,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  Fully qualified name of the category<br>
A URL query parameter that filters a list of categories by performing a regex match on the `fqName` field.<br>
It accepts a Perl compatible regex string. For example: `fqName=sample-.*` would return all categories whose
`fqName` matches the regex `sample-.*`; like - `test-sample-1`, `sample-test` etc.
  */
  FqName *string `json:"fqName,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import2.ApiLink `json:"links,omitempty"`
  /**
  Opaque metadata which can be associated to a category.<br>
It is a list of key-value pairs.<br>
For example, for a category 'California/SanJose' we can associate a geographical coordinate based metadata
like: {'latitude': '37.3382? N' , 'longitude': '121.8863? W'}.

The server does not validate this value nor does it enforce the uniqueness or any other constraints.
It is the responsibility of the user to ensure that any semantic or syntactic constraints are retained when mutating
this field.
  */
  Metadata []import3.KVPair `json:"metadata,omitempty"`
  /**
  Name of the category<br>
A URL query parameter that filters a list of categories by performing a regex match on the `name` field.<br>
It accepts a Perl compatible regex string. For example: `name=sample-name-.*` would return all categories whose
`name` matches the regex `sample-name-.*`; like - `test-sample-name-1`, `sample-name-test` etc.
  */
  Name *string `json:"name,omitempty"`
  /**
  The parent category of this category (may be null if this category is not part of a hierarchy).<br>
Each category can have at most one parent.<br>
A parent cannot be deleted until all the children categories are first deleted.<br>
Must be specified inside the post/put request body for child categories (if not specified, the service assumes
the category to be a parent category).<br>
This field is immutable.
  */
  ParentExtId *string `json:"parentExtId,omitempty"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
  
  Type *CategoryType `json:"type,omitempty"`
  /**
  The user specified name is a string that the user can specify; with syntax and semantics controlled by the user.

The server does not validate this value nor does it enforce the uniqueness or any other constraints.<br>
It is the responsibility of the user to ensure that any semantic or syntactic constraints are retained when mutating
this field. Unlike the name of the categories, which is immutable, the user name can be freely changed by the user
to meet their needs.
  */
  UserSpecifiedName *string `json:"userSpecifiedName,omitempty"`
}

func NewCategory() *Category {
  p := new(Category)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "prism.v4.config.Category"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a1.config.Category"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST Response for all response codes in api path /prism/v4.0.a1/config/categories Post operation
*/
type CategoryApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfCategoryApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewCategoryApiResponse() *CategoryApiResponse {
  p := new(CategoryApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "prism.v4.config.CategoryApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a1.config.CategoryApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *CategoryApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *CategoryApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfCategoryApiResponseData()
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
This attribute contains the list of entities which have been assigned the given category.<br>
These entities are grouped by entity types (like VM or HOST).<br>
Each associations object contains the total entities belonging to the given entity type, and category extId.
*/
type CategoryAssociationSummary struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  External identifier for the given category.
  */
  CategoryId *string `json:"categoryId,omitempty"`
  /**
  There are three category types:
- SYSTEM: Defined by an administrator. Cannot be modified by the user.
- INTERNAL: Defined by any internal process within Nutanix ecosystem. Cannot be viewed directly through a list call.
            Can be modified by the user.
- USER: Defined by a user through the v4 api
Currently v4 apis allow the creation of only User Category Type.
  */
  Count *int `json:"count,omitempty"`
  
  ResourceGroup *ResourceGroup `json:"resourceGroup,omitempty"`
  
  ResourceType *ResourceType `json:"resourceType,omitempty"`
}

func NewCategoryAssociationSummary() *CategoryAssociationSummary {
  p := new(CategoryAssociationSummary)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "prism.v4.config.CategoryAssociationSummary"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a1.config.CategoryAssociationSummary"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}





type CategoryAssociationSummaryProjection struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  External identifier for the given category.
  */
  CategoryId *string `json:"categoryId,omitempty"`
  /**
  There are three category types:
- SYSTEM: Defined by an administrator. Cannot be modified by the user.
- INTERNAL: Defined by any internal process within Nutanix ecosystem. Cannot be viewed directly through a list call.
            Can be modified by the user.
- USER: Defined by a user through the v4 api
Currently v4 apis allow the creation of only User Category Type.
  */
  Count *int `json:"count,omitempty"`
  
  ResourceGroup *ResourceGroup `json:"resourceGroup,omitempty"`
  
  ResourceType *ResourceType `json:"resourceType,omitempty"`
}

func NewCategoryAssociationSummaryProjection() *CategoryAssociationSummaryProjection {
  p := new(CategoryAssociationSummaryProjection)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "prism.v4.config.CategoryAssociationSummaryProjection"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a1.config.CategoryAssociationSummaryProjection"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST Response for all response codes in api path /prism/v4.0.a1/config/categories/{extId} Delete operation
*/
type CategoryDeleteApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfCategoryDeleteApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewCategoryDeleteApiResponse() *CategoryDeleteApiResponse {
  p := new(CategoryDeleteApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "prism.v4.config.CategoryDeleteApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a1.config.CategoryDeleteApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *CategoryDeleteApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *CategoryDeleteApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfCategoryDeleteApiResponseData()
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
REST Response for all response codes in api path /prism/v4.0.a1/config/categories/{extId} Get operation
*/
type CategoryGetResponses struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfCategoryGetResponsesData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewCategoryGetResponses() *CategoryGetResponses {
  p := new(CategoryGetResponses)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "prism.v4.config.CategoryGetResponses"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a1.config.CategoryGetResponses"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *CategoryGetResponses) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *CategoryGetResponses) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfCategoryGetResponsesData()
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




type CategoryProjection struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  Associations []CategoryAssociationSummary `json:"associations,omitempty"`
  
  CategoryAssociationSummaryProjection *CategoryAssociationSummaryProjection `json:"categoryAssociationSummaryProjection,omitempty"`
  
  CategorySummaryProjection *CategorySummaryProjection `json:"categorySummaryProjection,omitempty"`
  
  ChildCategories []CategorySummary `json:"childCategories,omitempty"`
  /**
  A string consisting of the description of the category as defined by the user.

The server does not validate this value nor does it enforce the uniqueness or any other constraints.<br>
It is the responsibility of the user to ensure that any semantic or syntactic constraints are retained when mutating
this field.
  */
  Description *string `json:"description,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  Fully qualified name of the category<br>
A URL query parameter that filters a list of categories by performing a regex match on the `fqName` field.<br>
It accepts a Perl compatible regex string. For example: `fqName=sample-.*` would return all categories whose
`fqName` matches the regex `sample-.*`; like - `test-sample-1`, `sample-test` etc.
  */
  FqName *string `json:"fqName,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import2.ApiLink `json:"links,omitempty"`
  /**
  Opaque metadata which can be associated to a category.<br>
It is a list of key-value pairs.<br>
For example, for a category 'California/SanJose' we can associate a geographical coordinate based metadata
like: {'latitude': '37.3382? N' , 'longitude': '121.8863? W'}.

The server does not validate this value nor does it enforce the uniqueness or any other constraints.
It is the responsibility of the user to ensure that any semantic or syntactic constraints are retained when mutating
this field.
  */
  Metadata []import3.KVPair `json:"metadata,omitempty"`
  /**
  Name of the category<br>
A URL query parameter that filters a list of categories by performing a regex match on the `name` field.<br>
It accepts a Perl compatible regex string. For example: `name=sample-name-.*` would return all categories whose
`name` matches the regex `sample-name-.*`; like - `test-sample-name-1`, `sample-name-test` etc.
  */
  Name *string `json:"name,omitempty"`
  /**
  The parent category of this category (may be null if this category is not part of a hierarchy).<br>
Each category can have at most one parent.<br>
A parent cannot be deleted until all the children categories are first deleted.<br>
Must be specified inside the post/put request body for child categories (if not specified, the service assumes
the category to be a parent category).<br>
This field is immutable.
  */
  ParentExtId *string `json:"parentExtId,omitempty"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
  
  Type *CategoryType `json:"type,omitempty"`
  /**
  The user specified name is a string that the user can specify; with syntax and semantics controlled by the user.

The server does not validate this value nor does it enforce the uniqueness or any other constraints.<br>
It is the responsibility of the user to ensure that any semantic or syntactic constraints are retained when mutating
this field. Unlike the name of the categories, which is immutable, the user name can be freely changed by the user
to meet their needs.
  */
  UserSpecifiedName *string `json:"userSpecifiedName,omitempty"`
}

func NewCategoryProjection() *CategoryProjection {
  p := new(CategoryProjection)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "prism.v4.config.CategoryProjection"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a1.config.CategoryProjection"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST Response for all response codes in api path /prism/v4.0.a1/config/categories/{extId} Put operation
*/
type CategoryPutResponses struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfCategoryPutResponsesData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewCategoryPutResponses() *CategoryPutResponses {
  p := new(CategoryPutResponses)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "prism.v4.config.CategoryPutResponses"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a1.config.CategoryPutResponses"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *CategoryPutResponses) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *CategoryPutResponses) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfCategoryPutResponsesData()
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




type CategorySummary struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  Associations []CategoryAssociationSummary `json:"associations,omitempty"`
  
  ChildCategories []CategorySummary `json:"childCategories,omitempty"`
  /**
  A string consisting of the description of the category as defined by the user.

The server does not validate this value nor does it enforce the uniqueness or any other constraints.<br>
It is the responsibility of the user to ensure that any semantic or syntactic constraints are retained when mutating
this field.
  */
  Description *string `json:"description,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  Fully qualified name of the category<br>
A URL query parameter that filters a list of categories by performing a regex match on the `fqName` field.<br>
It accepts a Perl compatible regex string. For example: `fqName=sample-.*` would return all categories whose
`fqName` matches the regex `sample-.*`; like - `test-sample-1`, `sample-test` etc.
  */
  FqName *string `json:"fqName,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import2.ApiLink `json:"links,omitempty"`
  /**
  Name of the category<br>
A URL query parameter that filters a list of categories by performing a regex match on the `name` field.<br>
It accepts a Perl compatible regex string. For example: `name=sample-name-.*` would return all categories whose
`name` matches the regex `sample-name-.*`; like - `test-sample-name-1`, `sample-name-test` etc.
  */
  Name *string `json:"name,omitempty"`
  /**
  The parent category of this category (may be null if this category is not part of a hierarchy).<br>
Each category can have at most one parent.<br>
A parent cannot be deleted until all the children categories are first deleted.<br>
Must be specified inside the post/put request body for child categories (if not specified, the service assumes
the category to be a parent category).<br>
This field is immutable.
  */
  ParentExtId *string `json:"parentExtId,omitempty"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
  
  Type *CategoryType `json:"type,omitempty"`
  /**
  The user specified name is a string that the user can specify; with syntax and semantics controlled by the user.

The server does not validate this value nor does it enforce the uniqueness or any other constraints.<br>
It is the responsibility of the user to ensure that any semantic or syntactic constraints are retained when mutating
this field. Unlike the name of the categories, which is immutable, the user name can be freely changed by the user
to meet their needs.
  */
  UserSpecifiedName *string `json:"userSpecifiedName,omitempty"`
}

func NewCategorySummary() *CategorySummary {
  p := new(CategorySummary)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "prism.v4.config.CategorySummary"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a1.config.CategorySummary"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST Response for all response codes in api path /prism/v4.0.a1/config/categories Get operation
*/
type CategorySummaryListApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfCategorySummaryListApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewCategorySummaryListApiResponse() *CategorySummaryListApiResponse {
  p := new(CategorySummaryListApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "prism.v4.config.CategorySummaryListApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a1.config.CategorySummaryListApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *CategorySummaryListApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *CategorySummaryListApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfCategorySummaryListApiResponseData()
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




type CategorySummaryProjection struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  Associations []CategoryAssociationSummary `json:"associations,omitempty"`
  
  CategoryAssociationSummaryProjection *CategoryAssociationSummaryProjection `json:"categoryAssociationSummaryProjection,omitempty"`
  
  CategorySummaryProjection *CategorySummaryProjection `json:"categorySummaryProjection,omitempty"`
  
  ChildCategories []CategorySummary `json:"childCategories,omitempty"`
  /**
  A string consisting of the description of the category as defined by the user.

The server does not validate this value nor does it enforce the uniqueness or any other constraints.<br>
It is the responsibility of the user to ensure that any semantic or syntactic constraints are retained when mutating
this field.
  */
  Description *string `json:"description,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  Fully qualified name of the category<br>
A URL query parameter that filters a list of categories by performing a regex match on the `fqName` field.<br>
It accepts a Perl compatible regex string. For example: `fqName=sample-.*` would return all categories whose
`fqName` matches the regex `sample-.*`; like - `test-sample-1`, `sample-test` etc.
  */
  FqName *string `json:"fqName,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import2.ApiLink `json:"links,omitempty"`
  /**
  Name of the category<br>
A URL query parameter that filters a list of categories by performing a regex match on the `name` field.<br>
It accepts a Perl compatible regex string. For example: `name=sample-name-.*` would return all categories whose
`name` matches the regex `sample-name-.*`; like - `test-sample-name-1`, `sample-name-test` etc.
  */
  Name *string `json:"name,omitempty"`
  /**
  The parent category of this category (may be null if this category is not part of a hierarchy).<br>
Each category can have at most one parent.<br>
A parent cannot be deleted until all the children categories are first deleted.<br>
Must be specified inside the post/put request body for child categories (if not specified, the service assumes
the category to be a parent category).<br>
This field is immutable.
  */
  ParentExtId *string `json:"parentExtId,omitempty"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
  
  Type *CategoryType `json:"type,omitempty"`
  /**
  The user specified name is a string that the user can specify; with syntax and semantics controlled by the user.

The server does not validate this value nor does it enforce the uniqueness or any other constraints.<br>
It is the responsibility of the user to ensure that any semantic or syntactic constraints are retained when mutating
this field. Unlike the name of the categories, which is immutable, the user name can be freely changed by the user
to meet their needs.
  */
  UserSpecifiedName *string `json:"userSpecifiedName,omitempty"`
}

func NewCategorySummaryProjection() *CategorySummaryProjection {
  p := new(CategorySummaryProjection)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "prism.v4.config.CategorySummaryProjection"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a1.config.CategorySummaryProjection"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
There are three category types:
- SYSTEM: Defined by an administrator. Cannot be modified by the user.
- INTERNAL: Defined by any internal process within Nutanix ecosystem. Cannot be viewed directly through a list call.
            Can be modified by the user.
- USER: Defined by a user through the v4 api
Currently v4 apis allow the creation of only User Category Type.
*/
type CategoryType int

const(
  CATEGORYTYPE_UNKNOWN CategoryType = 0
  CATEGORYTYPE_REDACTED CategoryType = 1
  CATEGORYTYPE_USER CategoryType = 2
  CATEGORYTYPE_SYSTEM CategoryType = 3
  CATEGORYTYPE_INTERNAL CategoryType = 4
)

// returns the name of the enum given an ordinal number
func (e *CategoryType) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "USER",
    "SYSTEM",
    "INTERNAL",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *CategoryType) index(name string) CategoryType {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "USER",
    "SYSTEM",
    "INTERNAL",
  }
  for idx := range names {
    if names[idx] == name {
      return CategoryType(idx)
    }
  }
  return CATEGORYTYPE_UNKNOWN
}

func (e *CategoryType) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for CategoryType:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *CategoryType) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e CategoryType) Ref() *CategoryType {
  return &e
}


/**
Contains references for entities given in EntityAssociation.<br>
This contains the entity ID and a list of links to fetch the associated entities.
*/
type Reference struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  Links []import2.ApiLink `json:"links,omitempty"`
  /**
  The external identifier of the resource which uniquely identifies it.
  */
  ResourceId *string `json:"resourceId,omitempty"`
}

func NewReference() *Reference {
  p := new(Reference)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "prism.v4.config.Reference"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a1.config.Reference"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
An enum denoting the resource group. Resources can be organised into either entity or policy. Hence it supports
two possible values:
- ENTITY
- POLICY
*/
type ResourceGroup int

const(
  RESOURCEGROUP_UNKNOWN ResourceGroup = 0
  RESOURCEGROUP_REDACTED ResourceGroup = 1
  RESOURCEGROUP_ENTITY ResourceGroup = 2
  RESOURCEGROUP_POLICY ResourceGroup = 3
)

// returns the name of the enum given an ordinal number
func (e *ResourceGroup) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "ENTITY",
    "POLICY",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *ResourceGroup) index(name string) ResourceGroup {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "ENTITY",
    "POLICY",
  }
  for idx := range names {
    if names[idx] == name {
      return ResourceGroup(idx)
    }
  }
  return RESOURCEGROUP_UNKNOWN
}

func (e *ResourceGroup) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for ResourceGroup:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *ResourceGroup) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e ResourceGroup) Ref() *ResourceGroup {
  return &e
}


/**
An enum denoting the associated resource types. Resource types are further grouped into ResourceGroups or a ResourceGroup.<br>
Following are the supported entity types:
- VM
- IMAGE
- SUBNET
- CLUSTER
- HOST
- REPORT
- MARKETPLACE_ITEM
- BLUEPRINT
- APP
Following are the supported policy types:
- IMAGE_PLACEMENT_POLICY
- NETWORK_SECURITY_POLICY
- NETWORK_SECURITY_RULE
- AFFINITY_RULE
- QOS_POLICY
- NGT_POLICY
- PROTECTION_RULE
- STORAGE_POLICY
- IMAGE_RATE_LIMIT
- RECOVERY_PLAN
*/
type ResourceType int

const(
  RESOURCETYPE_UNKNOWN ResourceType = 0
  RESOURCETYPE_REDACTED ResourceType = 1
  RESOURCETYPE_VM ResourceType = 2
  RESOURCETYPE_IMAGE ResourceType = 3
  RESOURCETYPE_SUBNET ResourceType = 4
  RESOURCETYPE_CLUSTER ResourceType = 5
  RESOURCETYPE_HOST ResourceType = 6
  RESOURCETYPE_REPORT ResourceType = 7
  RESOURCETYPE_MARKETPLACE_ITEM ResourceType = 8
  RESOURCETYPE_BLUEPRINT ResourceType = 9
  RESOURCETYPE_APP ResourceType = 10
  RESOURCETYPE_VOLUMEGROUP ResourceType = 11
  RESOURCETYPE_IMAGE_PLACEMENT_POLICY ResourceType = 12
  RESOURCETYPE_NETWORK_SECURITY_POLICY ResourceType = 13
  RESOURCETYPE_NETWORK_SECURITY_RULE ResourceType = 14
  RESOURCETYPE_VM_HOST_AFFINITY_POLICY ResourceType = 15
  RESOURCETYPE_QOS_POLICY ResourceType = 16
  RESOURCETYPE_NGT_POLICY ResourceType = 17
  RESOURCETYPE_PROTECTION_RULE ResourceType = 18
  RESOURCETYPE_ACCESS_CONTROL_POLICY ResourceType = 19
  RESOURCETYPE_STORAGE_POLICY ResourceType = 20
  RESOURCETYPE_IMAGE_RATE_LIMIT ResourceType = 21
  RESOURCETYPE_RECOVERY_PLAN ResourceType = 22
)

// returns the name of the enum given an ordinal number
func (e *ResourceType) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "VM",
    "IMAGE",
    "SUBNET",
    "CLUSTER",
    "HOST",
    "REPORT",
    "MARKETPLACE_ITEM",
    "BLUEPRINT",
    "APP",
    "VOLUMEGROUP",
    "IMAGE_PLACEMENT_POLICY",
    "NETWORK_SECURITY_POLICY",
    "NETWORK_SECURITY_RULE",
    "VM_HOST_AFFINITY_POLICY",
    "QOS_POLICY",
    "NGT_POLICY",
    "PROTECTION_RULE",
    "ACCESS_CONTROL_POLICY",
    "STORAGE_POLICY",
    "IMAGE_RATE_LIMIT",
    "RECOVERY_PLAN",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *ResourceType) index(name string) ResourceType {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "VM",
    "IMAGE",
    "SUBNET",
    "CLUSTER",
    "HOST",
    "REPORT",
    "MARKETPLACE_ITEM",
    "BLUEPRINT",
    "APP",
    "VOLUMEGROUP",
    "IMAGE_PLACEMENT_POLICY",
    "NETWORK_SECURITY_POLICY",
    "NETWORK_SECURITY_RULE",
    "VM_HOST_AFFINITY_POLICY",
    "QOS_POLICY",
    "NGT_POLICY",
    "PROTECTION_RULE",
    "ACCESS_CONTROL_POLICY",
    "STORAGE_POLICY",
    "IMAGE_RATE_LIMIT",
    "RECOVERY_PLAN",
  }
  for idx := range names {
    if names[idx] == name {
      return ResourceType(idx)
    }
  }
  return RESOURCETYPE_UNKNOWN
}

func (e *ResourceType) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for ResourceType:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *ResourceType) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e ResourceType) Ref() *ResourceType {
  return &e
}

type OneOfCategoryDeleteApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType1 *interface{} `json:"-"`
  oneOfType400 *import1.ErrorResponse `json:"-"`
}

func NewOneOfCategoryDeleteApiResponseData() *OneOfCategoryDeleteApiResponseData {
  p := new(OneOfCategoryDeleteApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfCategoryDeleteApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfCategoryDeleteApiResponseData is nil"))
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

func (p *OneOfCategoryDeleteApiResponseData) GetValue() interface{} {
  if "EMPTY" == *p.Discriminator {
    return *p.oneOfType1
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfCategoryDeleteApiResponseData) UnmarshalJSON(b []byte) error {
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
    if "prism.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCategoryDeleteApiResponseData"))
}

func (p *OneOfCategoryDeleteApiResponseData) MarshalJSON() ([]byte, error) {
  if "EMPTY" == *p.Discriminator {
    return json.Marshal(p.oneOfType1)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfCategoryDeleteApiResponseData")
}

type OneOfCategoryPutResponsesData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import1.ErrorResponse `json:"-"`
  oneOfType0 *Category `json:"-"`
}

func NewOneOfCategoryPutResponsesData() *OneOfCategoryPutResponsesData {
  p := new(OneOfCategoryPutResponsesData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfCategoryPutResponsesData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfCategoryPutResponsesData is nil"))
  }
  switch v.(type) {
    case import1.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = v.(import1.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case Category:
      if nil == p.oneOfType0 {p.oneOfType0 = new(Category)}
      *p.oneOfType0 = v.(Category)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfCategoryPutResponsesData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  return nil
}

func (p *OneOfCategoryPutResponsesData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import1.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "prism.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new(Category)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "prism.v4.config.Category" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(Category)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCategoryPutResponsesData"))
}

func (p *OneOfCategoryPutResponsesData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfCategoryPutResponsesData")
}

type OneOfAssociationDetailApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 []AssociationDetail `json:"-"`
  oneOfType400 *import1.ErrorResponse `json:"-"`
  oneOfType401 []AssociationDetailProjection `json:"-"`
}

func NewOneOfAssociationDetailApiResponseData() *OneOfAssociationDetailApiResponseData {
  p := new(OneOfAssociationDetailApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfAssociationDetailApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfAssociationDetailApiResponseData is nil"))
  }
  switch v.(type) {
    case []AssociationDetail:
      p.oneOfType0 = v.([]AssociationDetail)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<prism.v4.config.AssociationDetail>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<prism.v4.config.AssociationDetail>"
    case import1.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = v.(import1.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case []AssociationDetailProjection:
      p.oneOfType401 = v.([]AssociationDetailProjection)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<prism.v4.config.AssociationDetailProjection>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<prism.v4.config.AssociationDetailProjection>"
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfAssociationDetailApiResponseData) GetValue() interface{} {
  if "List<prism.v4.config.AssociationDetail>" == *p.Discriminator {
    return p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if "List<prism.v4.config.AssociationDetailProjection>" == *p.Discriminator {
    return p.oneOfType401
  }
  return nil
}

func (p *OneOfAssociationDetailApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new([]AssociationDetail)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "prism.v4.config.AssociationDetail" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<prism.v4.config.AssociationDetail>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<prism.v4.config.AssociationDetail>"
      return nil

    }
  }
  vOneOfType400 := new(import1.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "prism.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType401 := new([]AssociationDetailProjection)
  if err := json.Unmarshal(b, vOneOfType401); err == nil {
    if len(*vOneOfType401) == 0 || "prism.v4.config.AssociationDetailProjection" == *((*vOneOfType401)[0].ObjectType_) {
      p.oneOfType401 = *vOneOfType401
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<prism.v4.config.AssociationDetailProjection>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<prism.v4.config.AssociationDetailProjection>"
      return nil

    }
  }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfAssociationDetailApiResponseData"))
}

func (p *OneOfAssociationDetailApiResponseData) MarshalJSON() ([]byte, error) {
  if "List<prism.v4.config.AssociationDetail>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if "List<prism.v4.config.AssociationDetailProjection>" == *p.Discriminator {
    return json.Marshal(p.oneOfType401)
  }
  return nil, errors.New("No value to marshal for OneOfAssociationDetailApiResponseData")
}

type OneOfCategoryApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import1.ErrorResponse `json:"-"`
  oneOfType0 *Category `json:"-"`
}

func NewOneOfCategoryApiResponseData() *OneOfCategoryApiResponseData {
  p := new(OneOfCategoryApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfCategoryApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfCategoryApiResponseData is nil"))
  }
  switch v.(type) {
    case import1.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = v.(import1.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case Category:
      if nil == p.oneOfType0 {p.oneOfType0 = new(Category)}
      *p.oneOfType0 = v.(Category)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfCategoryApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  return nil
}

func (p *OneOfCategoryApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import1.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "prism.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new(Category)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "prism.v4.config.Category" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(Category)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCategoryApiResponseData"))
}

func (p *OneOfCategoryApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfCategoryApiResponseData")
}

type OneOfCategoryGetResponsesData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import1.ErrorResponse `json:"-"`
  oneOfType0 *Category `json:"-"`
}

func NewOneOfCategoryGetResponsesData() *OneOfCategoryGetResponsesData {
  p := new(OneOfCategoryGetResponsesData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfCategoryGetResponsesData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfCategoryGetResponsesData is nil"))
  }
  switch v.(type) {
    case import1.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = v.(import1.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case Category:
      if nil == p.oneOfType0 {p.oneOfType0 = new(Category)}
      *p.oneOfType0 = v.(Category)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfCategoryGetResponsesData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  return nil
}

func (p *OneOfCategoryGetResponsesData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import1.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "prism.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new(Category)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "prism.v4.config.Category" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(Category)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCategoryGetResponsesData"))
}

func (p *OneOfCategoryGetResponsesData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfCategoryGetResponsesData")
}

type OneOfCategorySummaryListApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 []CategorySummary `json:"-"`
  oneOfType400 *import1.ErrorResponse `json:"-"`
  oneOfType401 []CategorySummaryProjection `json:"-"`
}

func NewOneOfCategorySummaryListApiResponseData() *OneOfCategorySummaryListApiResponseData {
  p := new(OneOfCategorySummaryListApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfCategorySummaryListApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfCategorySummaryListApiResponseData is nil"))
  }
  switch v.(type) {
    case []CategorySummary:
      p.oneOfType0 = v.([]CategorySummary)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<prism.v4.config.CategorySummary>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<prism.v4.config.CategorySummary>"
    case import1.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = v.(import1.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case []CategorySummaryProjection:
      p.oneOfType401 = v.([]CategorySummaryProjection)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<prism.v4.config.CategorySummaryProjection>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<prism.v4.config.CategorySummaryProjection>"
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfCategorySummaryListApiResponseData) GetValue() interface{} {
  if "List<prism.v4.config.CategorySummary>" == *p.Discriminator {
    return p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if "List<prism.v4.config.CategorySummaryProjection>" == *p.Discriminator {
    return p.oneOfType401
  }
  return nil
}

func (p *OneOfCategorySummaryListApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new([]CategorySummary)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "prism.v4.config.CategorySummary" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<prism.v4.config.CategorySummary>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<prism.v4.config.CategorySummary>"
      return nil

    }
  }
  vOneOfType400 := new(import1.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "prism.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType401 := new([]CategorySummaryProjection)
  if err := json.Unmarshal(b, vOneOfType401); err == nil {
    if len(*vOneOfType401) == 0 || "prism.v4.config.CategorySummaryProjection" == *((*vOneOfType401)[0].ObjectType_) {
      p.oneOfType401 = *vOneOfType401
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<prism.v4.config.CategorySummaryProjection>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<prism.v4.config.CategorySummaryProjection>"
      return nil

    }
  }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCategorySummaryListApiResponseData"))
}

func (p *OneOfCategorySummaryListApiResponseData) MarshalJSON() ([]byte, error) {
  if "List<prism.v4.config.CategorySummary>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if "List<prism.v4.config.CategorySummaryProjection>" == *p.Discriminator {
    return json.Marshal(p.oneOfType401)
  }
  return nil, errors.New("No value to marshal for OneOfCategorySummaryListApiResponseData")
}

