/*
 * Generated file models/petstore/v1/defaultapi/defaultapi_model.go.
 *
 * Product version: 1.0.0-SNAPSHOT
 *
 * Part of the Petstore API project
 *
 * (c) 2021 Nutanix Inc.  All rights reserved
 *
 */

/*
  Default models for petstore
*/
package defaultapi
import (
  "bytes"
  import4 "github.com/nishantikagarg/sample-go-repo/petstore_go_sdk/models/common/v1/config"
  import1 "github.com/nishantikagarg/sample-go-repo/petstore_go_sdk/models/common/v1/response"
  "encoding/json"
  "errors"
  "fmt"
  import2 "github.com/nishantikagarg/sample-go-repo/petstore_go_sdk/models/petstore/v1/error"
  import3 "github.com/nishantikagarg/sample-go-repo/petstore_go_sdk/models/petstore/v1/utils"
)


type BirthStatus int

const(
  BIRTHSTATUS_WILD BirthStatus = 0
  BIRTHSTATUS_DOMESTIC BirthStatus = 1
  BIRTHSTATUS_UNKNOWN BirthStatus = 2
)

// returns the name of the enum given an ordinal number
func (e *BirthStatus) name(index int) string {
  names := [...]string {
    "WILD",
    "DOMESTIC",
    "$UNKNOWN",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *BirthStatus) index(name string) BirthStatus {
  names := [...]string {
    "WILD",
    "DOMESTIC",
    "$UNKNOWN",
  }
  for idx := range names {
    if names[idx] == name {
      return BirthStatus(idx)
    }
  }
  return BIRTHSTATUS_UNKNOWN
}

func (e *BirthStatus) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for BirthStatus:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *BirthStatus) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e BirthStatus) Ref() *BirthStatus {
  return &e
}



type BoolVal struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  BoolVal *bool `json:"boolVal,omitempty"`
}

func NewBoolVal() *BoolVal {
  p := new(BoolVal)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "petstore.v1.defaultapi.BoolVal"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "petstore.v1.r0.a1.defaultapi.BoolVal"}
  p.UnknownFields_ = map[string]interface{}{}

  p.BoolVal = new(bool)
  *p.BoolVal = true


  return p
}





type Cat struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  Category *Category `json:"category,omitempty"`
  
  Color *string `json:"color,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  
  FavoriteColor *string `json:"favorite_color,omitempty"`
  
  FavoriteFood *string `json:"favorite_food,omitempty"`
  
  Id *int64 `json:"id"`
  
  LastOwner *string `json:"lastOwner,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address at which the particular resource can be retrieved.
  */
  Links []import1.ApiLink `json:"links,omitempty"`
  
  Name *string `json:"name"`
  
  OriginDetails *Origin `json:"originDetails,omitempty"`
  
  PhotoFiles []string `json:"photoFiles,omitempty"`
  
  Point *Point `json:"point,omitempty"`
  
  Size *PetSize `json:"size,omitempty"`
  
  Status *PetStatus `json:"status,omitempty"`
  
  Tags []Tag `json:"tags,omitempty"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this id to change - for instance if a use case requires transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
  /**
  Whether the pet is vaccinated or not
  */
  Vaccinated *bool `json:"vaccinated,omitempty"`
}

func NewCat() *Cat {
  p := new(Cat)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "petstore.v1.defaultapi.Cat"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "petstore.v1.r0.a1.defaultapi.Cat"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}





type CatProjection struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  Category *Category `json:"category,omitempty"`
  
  Color *string `json:"color,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  
  FavoriteColor *string `json:"favorite_color,omitempty"`
  
  FavoriteFood *string `json:"favorite_food,omitempty"`
  
  Id *int64 `json:"id"`
  
  LastOwner *string `json:"lastOwner,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address at which the particular resource can be retrieved.
  */
  Links []import1.ApiLink `json:"links,omitempty"`
  
  Name *string `json:"name"`
  
  OriginDetails *Origin `json:"originDetails,omitempty"`
  
  PhotoFiles []string `json:"photoFiles,omitempty"`
  
  Point *Point `json:"point,omitempty"`
  
  Size *PetSize `json:"size,omitempty"`
  
  Status *PetStatus `json:"status,omitempty"`
  
  Tags []Tag `json:"tags,omitempty"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this id to change - for instance if a use case requires transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
  /**
  Whether the pet is vaccinated or not
  */
  Vaccinated *bool `json:"vaccinated,omitempty"`
}

func NewCatProjection() *CatProjection {
  p := new(CatProjection)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "petstore.v1.defaultapi.CatProjection"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "petstore.v1.r0.a1.defaultapi.CatProjection"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}





type Category struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  Id *int64 `json:"id,omitempty"`
  
  LastUpdated *int64 `json:"lastUpdated,omitempty"`
  
  Name *string `json:"name,omitempty"`
}

func NewCategory() *Category {
  p := new(Category)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "petstore.v1.defaultapi.Category"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "petstore.v1.r0.a1.defaultapi.Category"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST Response for all response codes in api path /petstore/v1.0.a1/defaultapi/pet-categories/{categoryId} Get operation
*/
type CategoryApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfCategoryApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewCategoryApiResponse() *CategoryApiResponse {
  p := new(CategoryApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "petstore.v1.defaultapi.CategoryApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "petstore.v1.r0.a1.defaultapi.CategoryApiResponse"}
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
REST Response for all response codes in api path /petstore/v1.0.a1/defaultapi/pet-categories Get operation
*/
type CategoryListApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfCategoryListApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewCategoryListApiResponse() *CategoryListApiResponse {
  p := new(CategoryListApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "petstore.v1.defaultapi.CategoryListApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "petstore.v1.r0.a1.defaultapi.CategoryListApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *CategoryListApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *CategoryListApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfCategoryListApiResponseData()
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
  
  Id *int64 `json:"id,omitempty"`
  
  LastUpdated *int64 `json:"lastUpdated,omitempty"`
  
  Name *string `json:"name,omitempty"`
}

func NewCategoryProjection() *CategoryProjection {
  p := new(CategoryProjection)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "petstore.v1.defaultapi.CategoryProjection"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "petstore.v1.r0.a1.defaultapi.CategoryProjection"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST Response for all response codes in api path /petstore/v1.0.a1/defaultapi/pets/{petId}/downloadImage Delete operation
*/
type DeletePetApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfDeletePetApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDeletePetApiResponse() *DeletePetApiResponse {
  p := new(DeletePetApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "petstore.v1.defaultapi.DeletePetApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "petstore.v1.r0.a1.defaultapi.DeletePetApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *DeletePetApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *DeletePetApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfDeletePetApiResponseData()
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




type DetailedPetSize int

const(
  DETAILEDPETSIZE_UNKNOWN DetailedPetSize = 0
  DETAILEDPETSIZE_REDACTED DetailedPetSize = 1
  DETAILEDPETSIZE_BIG DetailedPetSize = 2
  DETAILEDPETSIZE_MEDIUM DetailedPetSize = 3
  DETAILEDPETSIZE_SMALL DetailedPetSize = 4
  DETAILEDPETSIZE_NOT_APPLICABLE DetailedPetSize = 5
  DETAILEDPETSIZE_VERY_BIG DetailedPetSize = 6
  DETAILEDPETSIZE_VERY_SMALL DetailedPetSize = 7
)

// returns the name of the enum given an ordinal number
func (e *DetailedPetSize) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "BIG",
    "MEDIUM",
    "SMALL",
    "NOT_APPLICABLE",
    "VERY_BIG",
    "VERY_SMALL",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *DetailedPetSize) index(name string) DetailedPetSize {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "BIG",
    "MEDIUM",
    "SMALL",
    "NOT_APPLICABLE",
    "VERY_BIG",
    "VERY_SMALL",
  }
  for idx := range names {
    if names[idx] == name {
      return DetailedPetSize(idx)
    }
  }
  return DETAILEDPETSIZE_UNKNOWN
}

func (e *DetailedPetSize) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for DetailedPetSize:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *DetailedPetSize) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e DetailedPetSize) Ref() *DetailedPetSize {
  return &e
}



type Dog struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  Id *int64 `json:"id,omitempty"`
  
  Name *string `json:"name,omitempty"`
}

func NewDog() *Dog {
  p := new(Dog)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "petstore.v1.defaultapi.Dog"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "petstore.v1.r0.a1.defaultapi.Dog"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}





type DogProjection struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  CategoryProjection *CategoryProjection `json:"categoryProjection,omitempty"`
  
  Id *int64 `json:"id,omitempty"`
  
  Name *string `json:"name,omitempty"`
}

func NewDogProjection() *DogProjection {
  p := new(DogProjection)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "petstore.v1.defaultapi.DogProjection"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "petstore.v1.r0.a1.defaultapi.DogProjection"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}





type Echo struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  Alpha *import3.Alpha `json:"alpha,omitempty"`
  
  Bravo *import3.Bravo `json:"bravo,omitempty"`
  
  Charlie *import3.Charlie `json:"charlie,omitempty"`
  
  Name *string `json:"name,omitempty"`
  
  Severity *import4.MessageSeverity `json:"severity,omitempty"`
}

func NewEcho() *Echo {
  p := new(Echo)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "petstore.v1.defaultapi.Echo"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "petstore.v1.r0.a1.defaultapi.Echo"}
  p.UnknownFields_ = map[string]interface{}{}


  // set value for Alpha
  p.Alpha = new(import3.Alpha)
  *p.Alpha = import3.ALPHA_BAR
    // set value for Bravo

  p.Bravo = import3.NewBravo()
  p.Bravo.Id = new(int)
  *p.Bravo.Id = 2
  p.Bravo.Alpha = new(import3.Alpha)
  *p.Bravo.Alpha = import3.ALPHA_BAR
  // set value for Charlie

  p.Charlie = import3.NewCharlie()
  p.Charlie.Id = new(int)
  *p.Charlie.Id = 3
  p.Charlie.Tag = new(string)
  *p.Charlie.Tag = "tree"
  p.Charlie.Flag = new(bool)
  *p.Charlie.Flag = true
  p.Charlie.Alpha = new(import3.Alpha)
  *p.Charlie.Alpha = import3.ALPHA_BAR
  p.Charlie.Severity = new(import4.MessageSeverity)
  *p.Charlie.Severity = import4.MESSAGESEVERITY_INFO
  // set value for Severity
  p.Severity = new(import4.MessageSeverity)
  *p.Severity = import4.MESSAGESEVERITY_INFO
  
  return p
}





type IntVal struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  IntVal *int `json:"intVal,omitempty"`
}

func NewIntVal() *IntVal {
  p := new(IntVal)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "petstore.v1.defaultapi.IntVal"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "petstore.v1.r0.a1.defaultapi.IntVal"}
  p.UnknownFields_ = map[string]interface{}{}

  p.IntVal = new(int)
  *p.IntVal = 101


  return p
}





type Location struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  Country *string `json:"country,omitempty"`
  
  Region *string `json:"region,omitempty"`
}

func NewLocation() *Location {
  p := new(Location)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "petstore.v1.defaultapi.Location"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "petstore.v1.r0.a1.defaultapi.Location"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}





type OpenDays struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  WorkDay *WeekDays `json:"workDay,omitempty"`
}

func NewOpenDays() *OpenDays {
  p := new(OpenDays)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "petstore.v1.defaultapi.OpenDays"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "petstore.v1.r0.a1.defaultapi.OpenDays"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}





type Origin struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  Ancestors []string `json:"ancestors,omitempty"`
  
  BirthDate *string `json:"birthDate,omitempty"`
  
  BirthPlace *Location `json:"birthPlace,omitempty"`
  
  BirthStatus *BirthStatus `json:"birthStatus,omitempty"`
}

func NewOrigin() *Origin {
  p := new(Origin)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "petstore.v1.defaultapi.Origin"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "petstore.v1.r0.a1.defaultapi.Origin"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}





type Pet struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  Category *Category `json:"category,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  
  FavoriteColor *string `json:"favorite_color,omitempty"`
  
  FavoriteFood *string `json:"favorite_food,omitempty"`
  
  Id *int64 `json:"id"`
  
  LastOwner *string `json:"lastOwner,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address at which the particular resource can be retrieved.
  */
  Links []import1.ApiLink `json:"links,omitempty"`
  
  Name *string `json:"name"`
  
  OriginDetails *Origin `json:"originDetails,omitempty"`
  
  PhotoFiles []string `json:"photoFiles,omitempty"`
  
  Point *Point `json:"point,omitempty"`
  
  Size *PetSize `json:"size,omitempty"`
  
  Status *PetStatus `json:"status,omitempty"`
  
  Tags []Tag `json:"tags,omitempty"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this id to change - for instance if a use case requires transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
  /**
  Whether the pet is vaccinated or not
  */
  Vaccinated *bool `json:"vaccinated,omitempty"`
}

func NewPet() *Pet {
  p := new(Pet)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "petstore.v1.defaultapi.Pet"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "petstore.v1.r0.a1.defaultapi.Pet"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST Response for all response codes in api path /petstore/v1.0.a1/defaultapi/pets/error Get operation
*/
type PetApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfPetApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewPetApiResponse() *PetApiResponse {
  p := new(PetApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "petstore.v1.defaultapi.PetApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "petstore.v1.r0.a1.defaultapi.PetApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *PetApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *PetApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfPetApiResponseData()
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
REST Response for all response codes in api path /petstore/v1.0.a1/defaultapi/pets Get operation
*/
type PetListApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfPetListApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewPetListApiResponse() *PetListApiResponse {
  p := new(PetListApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "petstore.v1.defaultapi.PetListApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "petstore.v1.r0.a1.defaultapi.PetListApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *PetListApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *PetListApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfPetListApiResponseData()
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




type PetProjection struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  Category *Category `json:"category,omitempty"`
  
  CategoryProjection *CategoryProjection `json:"categoryProjection,omitempty"`
  
  DogProjection *DogProjection `json:"dogProjection,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  
  FavoriteColor *string `json:"favorite_color,omitempty"`
  
  FavoriteFood *string `json:"favorite_food,omitempty"`
  
  Id *int64 `json:"id"`
  
  LastOwner *string `json:"lastOwner,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address at which the particular resource can be retrieved.
  */
  Links []import1.ApiLink `json:"links,omitempty"`
  
  Name *string `json:"name"`
  
  OriginDetails *Origin `json:"originDetails,omitempty"`
  
  PhotoFiles []string `json:"photoFiles,omitempty"`
  
  Point *Point `json:"point,omitempty"`
  
  Size *PetSize `json:"size,omitempty"`
  
  Status *PetStatus `json:"status,omitempty"`
  
  Tags []Tag `json:"tags,omitempty"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this id to change - for instance if a use case requires transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
  /**
  Whether the pet is vaccinated or not
  */
  Vaccinated *bool `json:"vaccinated,omitempty"`
}

func NewPetProjection() *PetProjection {
  p := new(PetProjection)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "petstore.v1.defaultapi.PetProjection"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "petstore.v1.r0.a1.defaultapi.PetProjection"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}





type PetSize int

const(
  PETSIZE_UNKNOWN PetSize = 0
  PETSIZE_REDACTED PetSize = 1
  PETSIZE_BIG PetSize = 2
  PETSIZE_MEDIUM PetSize = 3
  PETSIZE_SMALL PetSize = 4
  PETSIZE_NOT_APPLICABLE PetSize = 5
)

// returns the name of the enum given an ordinal number
func (e *PetSize) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "BIG",
    "MEDIUM",
    "SMALL",
    "NOT_APPLICABLE",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *PetSize) index(name string) PetSize {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "BIG",
    "MEDIUM",
    "SMALL",
    "NOT_APPLICABLE",
  }
  for idx := range names {
    if names[idx] == name {
      return PetSize(idx)
    }
  }
  return PETSIZE_UNKNOWN
}

func (e *PetSize) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for PetSize:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *PetSize) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e PetSize) Ref() *PetSize {
  return &e
}



type PetStatus int

const(
  PETSTATUS_UNKNOWN PetStatus = 0
  PETSTATUS_REDACTED PetStatus = 1
  PETSTATUS_AVAILABLE PetStatus = 2
  PETSTATUS_PENDING PetStatus = 3
  PETSTATUS_SOLD PetStatus = 4
  PETSTATUS_NOT_APPLICABLE PetStatus = 5
)

// returns the name of the enum given an ordinal number
func (e *PetStatus) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "AVAILABLE",
    "PENDING",
    "SOLD",
    "NOT_APPLICABLE",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *PetStatus) index(name string) PetStatus {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "AVAILABLE",
    "PENDING",
    "SOLD",
    "NOT_APPLICABLE",
  }
  for idx := range names {
    if names[idx] == name {
      return PetStatus(idx)
    }
  }
  return PETSTATUS_UNKNOWN
}

func (e *PetStatus) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for PetStatus:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *PetStatus) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e PetStatus) Ref() *PetStatus {
  return &e
}



type Point struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  TimestampEpoch *int `json:"timestampEpoch,omitempty"`
  
  ValueItemDiscriminator_ *string `json:"$valueItemDiscriminator,omitempty"`
  
  Value *OneOfPointValue `json:"value,omitempty"`
  
  ValueItemDiscriminator *string `json:"valueItemDiscriminator,omitempty"`
}

func NewPoint() *Point {
  p := new(Point)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "petstore.v1.defaultapi.Point"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "petstore.v1.r0.a1.defaultapi.Point"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *Point) GetValue() interface{} {
  if nil == p.Value {
    return nil
  }
  return p.Value.GetValue()
}

func (p *Point) SetValue(v interface{}) error {
  if nil == p.Value {
    p.Value = NewOneOfPointValue()
  }
  e := p.Value.SetValue(v)
  if nil == e {
    if nil == p.ValueItemDiscriminator_ {
      p.ValueItemDiscriminator_ = new(string)
    }
    *p.ValueItemDiscriminator_ = *p.Value.Discriminator
  }
  return e
}




type ReducedPet struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  Category *Category `json:"category,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  
  Id *int64 `json:"id,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address at which the particular resource can be retrieved.
  */
  Links []import1.ApiLink `json:"links,omitempty"`
  
  Name *string `json:"name,omitempty"`
  
  PhotoFiles []string `json:"photoFiles,omitempty"`
  
  Point *Point `json:"point,omitempty"`
  
  Size *PetSize `json:"size,omitempty"`
  
  Status *PetStatus `json:"status,omitempty"`
  
  Tags []Tag `json:"tags,omitempty"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this id to change - for instance if a use case requires transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
}

func NewReducedPet() *ReducedPet {
  p := new(ReducedPet)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "petstore.v1.defaultapi.ReducedPet"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "petstore.v1.r0.a1.defaultapi.ReducedPet"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}





type StrVal struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  StrVal *string `json:"strVal,omitempty"`
}

func NewStrVal() *StrVal {
  p := new(StrVal)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "petstore.v1.defaultapi.StrVal"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "petstore.v1.r0.a1.defaultapi.StrVal"}
  p.UnknownFields_ = map[string]interface{}{}

  p.StrVal = new(string)
  *p.StrVal = "foo"


  return p
}





type Tag struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  Id *int64 `json:"id,omitempty"`
  
  Name *string `json:"name,omitempty"`
}

func NewTag() *Tag {
  p := new(Tag)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "petstore.v1.defaultapi.Tag"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "petstore.v1.r0.a1.defaultapi.Tag"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST Response for all response codes in api path /petstore/v1.0.a1/defaultapi/tags Get operation
*/
type TagListApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfTagListApiResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewTagListApiResponse() *TagListApiResponse {
  p := new(TagListApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "petstore.v1.defaultapi.TagListApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "petstore.v1.r0.a1.defaultapi.TagListApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *TagListApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *TagListApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfTagListApiResponseData()
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




type Url struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  Url *string `json:"url,omitempty"`
}

func NewUrl() *Url {
  p := new(Url)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "petstore.v1.defaultapi.Url"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "petstore.v1.r0.a1.defaultapi.Url"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST Response for all response codes in api path /petstore/v1.0.a1/defaultapi/pets/{petId}/uploadImage Post operation
*/
type UrlResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfUrlResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewUrlResponse() *UrlResponse {
  p := new(UrlResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "petstore.v1.defaultapi.UrlResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "petstore.v1.r0.a1.defaultapi.UrlResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *UrlResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *UrlResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfUrlResponseData()
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




type WeekDays int

const(
  WEEKDAYS_UNKNOWN WeekDays = 0
  WEEKDAYS_REDACTED WeekDays = 1
  WEEKDAYS_MONDAY WeekDays = 2
  WEEKDAYS_TUESDAY WeekDays = 3
  WEEKDAYS_WEDNESDAY WeekDays = 4
  WEEKDAYS_THURSDAY WeekDays = 5
  WEEKDAYS_FRIDAY WeekDays = 6
  WEEKDAYS_SATURDAY WeekDays = 7
  WEEKDAYS_SUNDAY WeekDays = 8
)

// returns the name of the enum given an ordinal number
func (e *WeekDays) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "MONDAY",
    "TUESDAY",
    "WEDNESDAY",
    "THURSDAY",
    "FRIDAY",
    "SATURDAY",
    "SUNDAY",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *WeekDays) index(name string) WeekDays {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "MONDAY",
    "TUESDAY",
    "WEDNESDAY",
    "THURSDAY",
    "FRIDAY",
    "SATURDAY",
    "SUNDAY",
  }
  for idx := range names {
    if names[idx] == name {
      return WeekDays(idx)
    }
  }
  return WEEKDAYS_UNKNOWN
}

func (e *WeekDays) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for WeekDays:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *WeekDays) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e WeekDays) Ref() *WeekDays {
  return &e
}

type OneOfCategoryListApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType2080 []Category `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
  oneOfType401 []CategoryProjection `json:"-"`
}

func NewOneOfCategoryListApiResponseData() *OneOfCategoryListApiResponseData {
  p := new(OneOfCategoryListApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfCategoryListApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfCategoryListApiResponseData is nil"))
  }
  switch v.(type) {
    case []Category:
      p.oneOfType2080 = v.([]Category)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<petstore.v1.defaultapi.Category>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<petstore.v1.defaultapi.Category>"
    case import2.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = v.(import2.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case []CategoryProjection:
      p.oneOfType401 = v.([]CategoryProjection)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<petstore.v1.defaultapi.CategoryProjection>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<petstore.v1.defaultapi.CategoryProjection>"
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfCategoryListApiResponseData) GetValue() interface{} {
  if "List<petstore.v1.defaultapi.Category>" == *p.Discriminator {
    return p.oneOfType2080
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if "List<petstore.v1.defaultapi.CategoryProjection>" == *p.Discriminator {
    return p.oneOfType401
  }
  return nil
}

func (p *OneOfCategoryListApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType2080 := new([]Category)
  if err := json.Unmarshal(b, vOneOfType2080); err == nil {
    if len(*vOneOfType2080) == 0 || "petstore.v1.defaultapi.Category" == *((*vOneOfType2080)[0].ObjectType_) {
      p.oneOfType2080 = *vOneOfType2080
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<petstore.v1.defaultapi.Category>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<petstore.v1.defaultapi.Category>"
      return nil

    }
  }
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "petstore.v1.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType401 := new([]CategoryProjection)
  if err := json.Unmarshal(b, vOneOfType401); err == nil {
    if len(*vOneOfType401) == 0 || "petstore.v1.defaultapi.CategoryProjection" == *((*vOneOfType401)[0].ObjectType_) {
      p.oneOfType401 = *vOneOfType401
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<petstore.v1.defaultapi.CategoryProjection>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<petstore.v1.defaultapi.CategoryProjection>"
      return nil

    }
  }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCategoryListApiResponseData"))
}

func (p *OneOfCategoryListApiResponseData) MarshalJSON() ([]byte, error) {
  if "List<petstore.v1.defaultapi.Category>" == *p.Discriminator {
    return json.Marshal(p.oneOfType2080)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if "List<petstore.v1.defaultapi.CategoryProjection>" == *p.Discriminator {
    return json.Marshal(p.oneOfType401)
  }
  return nil, errors.New("No value to marshal for OneOfCategoryListApiResponseData")
}

type OneOfPetApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
  oneOfType2080 *Pet `json:"-"`
}

func NewOneOfPetApiResponseData() *OneOfPetApiResponseData {
  p := new(OneOfPetApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfPetApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfPetApiResponseData is nil"))
  }
  switch v.(type) {
    case import2.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = v.(import2.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case Pet:
      if nil == p.oneOfType2080 {p.oneOfType2080 = new(Pet)}
      *p.oneOfType2080 = v.(Pet)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType2080.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType2080.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfPetApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if p.oneOfType2080 != nil && *p.oneOfType2080.ObjectType_ == *p.Discriminator {
    return *p.oneOfType2080
  }
  return nil
}

func (p *OneOfPetApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "petstore.v1.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType2080 := new(Pet)
  if err := json.Unmarshal(b, vOneOfType2080); err == nil {
    if "petstore.v1.defaultapi.Pet" == *vOneOfType2080.ObjectType_ {
          if nil == p.oneOfType2080 {p.oneOfType2080 = new(Pet)}
      *p.oneOfType2080 = *vOneOfType2080
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType2080.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType2080.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfPetApiResponseData"))
}

func (p *OneOfPetApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if p.oneOfType2080 != nil && *p.oneOfType2080.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType2080)
  }
  return nil, errors.New("No value to marshal for OneOfPetApiResponseData")
}

type OneOfPointValue struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType2001 *string `json:"-"`
  oneOfType2002 *int `json:"-"`
}

func NewOneOfPointValue() *OneOfPointValue {
  p := new(OneOfPointValue)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfPointValue) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfPointValue is nil"))
  }
  switch v.(type) {
    case string:
      if nil == p.oneOfType2001 {p.oneOfType2001 = new(string)}
      *p.oneOfType2001 = v.(string)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "String"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "String"
    case int:
      if nil == p.oneOfType2002 {p.oneOfType2002 = new(int)}
      *p.oneOfType2002 = v.(int)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "Integer"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "Integer"
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfPointValue) GetValue() interface{} {
  if "String" == *p.Discriminator {
    return *p.oneOfType2001
  }
  if "Integer" == *p.Discriminator {
    return *p.oneOfType2002
  }
  return nil
}

func (p *OneOfPointValue) UnmarshalJSON(b []byte) error {
  vOneOfType2001 := new(string)
  if err := json.Unmarshal(b, vOneOfType2001); err == nil {
          if nil == p.oneOfType2001 {p.oneOfType2001 = new(string)}
      *p.oneOfType2001 = *vOneOfType2001
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "String"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "String"
      return nil
  }
  vOneOfType2002 := new(int)
  if err := json.Unmarshal(b, vOneOfType2002); err == nil {
          if nil == p.oneOfType2002 {p.oneOfType2002 = new(int)}
      *p.oneOfType2002 = *vOneOfType2002
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "Integer"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "Integer"
      return nil
  }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfPointValue"))
}

func (p *OneOfPointValue) MarshalJSON() ([]byte, error) {
  if "String" == *p.Discriminator {
    return json.Marshal(p.oneOfType2001)
  }
  if "Integer" == *p.Discriminator {
    return json.Marshal(p.oneOfType2002)
  }
  return nil, errors.New("No value to marshal for OneOfPointValue")
}

type OneOfPetListApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
  oneOfType2080 []Pet `json:"-"`
  oneOfType401 []PetProjection `json:"-"`
}

func NewOneOfPetListApiResponseData() *OneOfPetListApiResponseData {
  p := new(OneOfPetListApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfPetListApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfPetListApiResponseData is nil"))
  }
  switch v.(type) {
    case import2.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = v.(import2.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case []Pet:
      p.oneOfType2080 = v.([]Pet)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<petstore.v1.defaultapi.Pet>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<petstore.v1.defaultapi.Pet>"
    case []PetProjection:
      p.oneOfType401 = v.([]PetProjection)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<petstore.v1.defaultapi.PetProjection>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<petstore.v1.defaultapi.PetProjection>"
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfPetListApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if "List<petstore.v1.defaultapi.Pet>" == *p.Discriminator {
    return p.oneOfType2080
  }
  if "List<petstore.v1.defaultapi.PetProjection>" == *p.Discriminator {
    return p.oneOfType401
  }
  return nil
}

func (p *OneOfPetListApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "petstore.v1.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType2080 := new([]Pet)
  if err := json.Unmarshal(b, vOneOfType2080); err == nil {
    if len(*vOneOfType2080) == 0 || "petstore.v1.defaultapi.Pet" == *((*vOneOfType2080)[0].ObjectType_) {
      p.oneOfType2080 = *vOneOfType2080
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<petstore.v1.defaultapi.Pet>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<petstore.v1.defaultapi.Pet>"
      return nil

    }
  }
  vOneOfType401 := new([]PetProjection)
  if err := json.Unmarshal(b, vOneOfType401); err == nil {
    if len(*vOneOfType401) == 0 || "petstore.v1.defaultapi.PetProjection" == *((*vOneOfType401)[0].ObjectType_) {
      p.oneOfType401 = *vOneOfType401
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<petstore.v1.defaultapi.PetProjection>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<petstore.v1.defaultapi.PetProjection>"
      return nil

    }
  }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfPetListApiResponseData"))
}

func (p *OneOfPetListApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if "List<petstore.v1.defaultapi.Pet>" == *p.Discriminator {
    return json.Marshal(p.oneOfType2080)
  }
  if "List<petstore.v1.defaultapi.PetProjection>" == *p.Discriminator {
    return json.Marshal(p.oneOfType401)
  }
  return nil, errors.New("No value to marshal for OneOfPetListApiResponseData")
}

type OneOfCategoryApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
  oneOfType2080 *Category `json:"-"`
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
    case import2.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = v.(import2.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case Category:
      if nil == p.oneOfType2080 {p.oneOfType2080 = new(Category)}
      *p.oneOfType2080 = v.(Category)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType2080.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType2080.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfCategoryApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if p.oneOfType2080 != nil && *p.oneOfType2080.ObjectType_ == *p.Discriminator {
    return *p.oneOfType2080
  }
  return nil
}

func (p *OneOfCategoryApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "petstore.v1.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType2080 := new(Category)
  if err := json.Unmarshal(b, vOneOfType2080); err == nil {
    if "petstore.v1.defaultapi.Category" == *vOneOfType2080.ObjectType_ {
          if nil == p.oneOfType2080 {p.oneOfType2080 = new(Category)}
      *p.oneOfType2080 = *vOneOfType2080
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType2080.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType2080.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCategoryApiResponseData"))
}

func (p *OneOfCategoryApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if p.oneOfType2080 != nil && *p.oneOfType2080.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType2080)
  }
  return nil, errors.New("No value to marshal for OneOfCategoryApiResponseData")
}

type OneOfDeletePetApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
  oneOfType2080 *interface{} `json:"-"`
}

func NewOneOfDeletePetApiResponseData() *OneOfDeletePetApiResponseData {
  p := new(OneOfDeletePetApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfDeletePetApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfDeletePetApiResponseData is nil"))
  }
  if nil == v {
    if nil == p.oneOfType2080 {p.oneOfType2080 = new(interface {})}
    *p.oneOfType2080 = nil
    if nil == p.Discriminator {p.Discriminator = new(string)}
    *p.Discriminator = "EMPTY"
    if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
    *p.ObjectType_ = "EMPTY"
    return nil
  }
  switch v.(type) {
    case import2.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = v.(import2.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfDeletePetApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if "EMPTY" == *p.Discriminator {
    return *p.oneOfType2080
  }
  return nil
}

func (p *OneOfDeletePetApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType2080 := new(interface {})
  if err := json.Unmarshal(b, vOneOfType2080); err == nil {
    if nil == *vOneOfType2080 {
      if nil == p.oneOfType2080 {p.oneOfType2080 = new(interface {})}
      *p.oneOfType2080 = nil
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "EMPTY"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "EMPTY"
      return nil
    }
  }
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "petstore.v1.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeletePetApiResponseData"))
}

func (p *OneOfDeletePetApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if "EMPTY" == *p.Discriminator {
    return json.Marshal(p.oneOfType2080)
  }
  return nil, errors.New("No value to marshal for OneOfDeletePetApiResponseData")
}

type OneOfTagListApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
  oneOfType2080 []Tag `json:"-"`
}

func NewOneOfTagListApiResponseData() *OneOfTagListApiResponseData {
  p := new(OneOfTagListApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfTagListApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfTagListApiResponseData is nil"))
  }
  switch v.(type) {
    case import2.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = v.(import2.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case []Tag:
      p.oneOfType2080 = v.([]Tag)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<petstore.v1.defaultapi.Tag>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<petstore.v1.defaultapi.Tag>"
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfTagListApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if "List<petstore.v1.defaultapi.Tag>" == *p.Discriminator {
    return p.oneOfType2080
  }
  return nil
}

func (p *OneOfTagListApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "petstore.v1.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType2080 := new([]Tag)
  if err := json.Unmarshal(b, vOneOfType2080); err == nil {
    if len(*vOneOfType2080) == 0 || "petstore.v1.defaultapi.Tag" == *((*vOneOfType2080)[0].ObjectType_) {
      p.oneOfType2080 = *vOneOfType2080
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<petstore.v1.defaultapi.Tag>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<petstore.v1.defaultapi.Tag>"
      return nil

    }
  }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfTagListApiResponseData"))
}

func (p *OneOfTagListApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if "List<petstore.v1.defaultapi.Tag>" == *p.Discriminator {
    return json.Marshal(p.oneOfType2080)
  }
  return nil, errors.New("No value to marshal for OneOfTagListApiResponseData")
}

type OneOfUrlResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
  oneOfType2080 *Url `json:"-"`
}

func NewOneOfUrlResponseData() *OneOfUrlResponseData {
  p := new(OneOfUrlResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfUrlResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfUrlResponseData is nil"))
  }
  switch v.(type) {
    case import2.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = v.(import2.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case Url:
      if nil == p.oneOfType2080 {p.oneOfType2080 = new(Url)}
      *p.oneOfType2080 = v.(Url)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType2080.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType2080.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfUrlResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if p.oneOfType2080 != nil && *p.oneOfType2080.ObjectType_ == *p.Discriminator {
    return *p.oneOfType2080
  }
  return nil
}

func (p *OneOfUrlResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "petstore.v1.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType2080 := new(Url)
  if err := json.Unmarshal(b, vOneOfType2080); err == nil {
    if "petstore.v1.defaultapi.Url" == *vOneOfType2080.ObjectType_ {
          if nil == p.oneOfType2080 {p.oneOfType2080 = new(Url)}
      *p.oneOfType2080 = *vOneOfType2080
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType2080.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType2080.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUrlResponseData"))
}

func (p *OneOfUrlResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if p.oneOfType2080 != nil && *p.oneOfType2080.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType2080)
  }
  return nil, errors.New("No value to marshal for OneOfUrlResponseData")
}

