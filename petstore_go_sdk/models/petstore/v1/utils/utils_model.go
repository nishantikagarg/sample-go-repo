/*
 * Generated file models/petstore/v1/utils/utils_model.go.
 *
 * Product version: 1.0.0-SNAPSHOT
 *
 * Part of the Petstore API project
 *
 * (c) 2022 Nutanix Inc.  All rights reserved
 *
 */

/*
  Utility models for petstore
*/
package utils
import (
  "bytes"
  import1 "github.com/nishantikagarg/sample-go-repo/petstore_go_sdk/v1/models/common/v1/config"
  import3 "github.com/nishantikagarg/sample-go-repo/petstore_go_sdk/v1/models/common/v1/response"
  "encoding/json"
  "errors"
  "fmt"
  import2 "github.com/nishantikagarg/sample-go-repo/petstore_go_sdk/v1/models/petstore/v1/error"
)


type Alpha int

const(
  ALPHA_FOO Alpha = 0
  ALPHA_BAR Alpha = 1
  ALPHA_BAZ Alpha = 2
  ALPHA_QUX Alpha = 3
  ALPHA_FU_BAR Alpha = 4
  ALPHA_ROGUE_ONE Alpha = 5
  ALPHA_TRANSIENT Alpha = 6
  ALPHA_UNKNOWN Alpha = 7
)

// returns the name of the enum given an ordinal number
func (e *Alpha) name(index int) string {
  names := [...]string {
    "FOO",
    "BAR",
    "BAZ",
    "QUX",
    "FU:BAR",
    "ROGUE-ONE",
    "transient",
    "$UNKNOWN",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *Alpha) index(name string) Alpha {
  names := [...]string {
    "FOO",
    "BAR",
    "BAZ",
    "QUX",
    "FU:BAR",
    "ROGUE-ONE",
    "transient",
    "$UNKNOWN",
  }
  for idx := range names {
    if names[idx] == name {
      return Alpha(idx)
    }
  }
  return ALPHA_UNKNOWN
}

func (e *Alpha) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for Alpha:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *Alpha) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e Alpha) Ref() *Alpha {
  return &e
}



type Bravo struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  Alpha *Alpha `json:"alpha,omitempty"`
  
  Id *int `json:"id,omitempty"`
}

func NewBravo() *Bravo {
  p := new(Bravo)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "petstore.v1.utils.Bravo"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "petstore.v1.r0.a1.utils.Bravo"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}





type Charlie struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  Alpha *Alpha `json:"alpha,omitempty"`
  
  Flag *bool `json:"flag,omitempty"`
  
  Id *int `json:"id,omitempty"`
  
  Severity *import1.MessageSeverity `json:"severity,omitempty"`
  
  Tag *string `json:"tag,omitempty"`
}

func NewCharlie() *Charlie {
  p := new(Charlie)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "petstore.v1.utils.Charlie"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "petstore.v1.r0.a1.utils.Charlie"}
  p.UnknownFields_ = map[string]interface{}{}


  // set value for Alpha
  p.Alpha = new(Alpha)
  *p.Alpha = ALPHA_QUX
  
  return p
}




/**
REST Response for all response codes in api path /petstore/v1.0.a1/utils/echo Post operation
*/
type CreateEchoResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfCreateEchoResponseData `json:"data,omitempty"`
  
  Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewCreateEchoResponse() *CreateEchoResponse {
  p := new(CreateEchoResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "petstore.v1.utils.CreateEchoResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "petstore.v1.r0.a1.utils.CreateEchoResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *CreateEchoResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *CreateEchoResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfCreateEchoResponseData()
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




type Delta struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  Alpha *Alpha `json:"alpha,omitempty"`
  
  Bravo *Bravo `json:"bravo,omitempty"`
  
  Charlie *Charlie `json:"charlie,omitempty"`
  
  Severity *import1.MessageSeverity `json:"severity,omitempty"`
}

func NewDelta() *Delta {
  p := new(Delta)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "petstore.v1.utils.Delta"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "petstore.v1.r0.a1.utils.Delta"}
  p.UnknownFields_ = map[string]interface{}{}


  // set value for Alpha
  p.Alpha = new(Alpha)
  *p.Alpha = ALPHA_FOO
    // set value for Bravo

  p.Bravo = NewBravo()
  p.Bravo.Id = new(int)
  *p.Bravo.Id = 0
  p.Bravo.Alpha = new(Alpha)
  *p.Bravo.Alpha = ALPHA_FOO
  // set value for Charlie

  p.Charlie = NewCharlie()
  p.Charlie.Id = new(int)
  *p.Charlie.Id = 1
  p.Charlie.Tag = new(string)
  *p.Charlie.Tag = "wun"
  p.Charlie.Flag = new(bool)
  *p.Charlie.Flag = false
  p.Charlie.Alpha = new(Alpha)
  *p.Charlie.Alpha = ALPHA_FOO

  return p
}





type Echo struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  Name *string `json:"name,omitempty"`
  
  Value *string `json:"value,omitempty"`
}

func NewEcho() *Echo {
  p := new(Echo)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "petstore.v1.utils.Echo"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "petstore.v1.r0.a1.utils.Echo"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST Response for all response codes in api path /petstore/v1.0.a1/utils/namespaces/{id} Get operation
*/
type MessageApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfMessageApiResponseData `json:"data,omitempty"`
  
  Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewMessageApiResponse() *MessageApiResponse {
  p := new(MessageApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "petstore.v1.utils.MessageApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "petstore.v1.r0.a1.utils.MessageApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *MessageApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *MessageApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfMessageApiResponseData()
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
REST Response for all response codes in api path /petstore/v1.0.a1/utils/testUtils Post operation
*/
type TestUtilsResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfTestUtilsResponseData `json:"data,omitempty"`
  
  Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewTestUtilsResponse() *TestUtilsResponse {
  p := new(TestUtilsResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "petstore.v1.utils.TestUtilsResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "petstore.v1.r0.a1.utils.TestUtilsResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *TestUtilsResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *TestUtilsResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfTestUtilsResponseData()
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


type OneOfCreateEchoResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
  oneOfType2080 *Echo `json:"-"`
}

func NewOneOfCreateEchoResponseData() *OneOfCreateEchoResponseData {
  p := new(OneOfCreateEchoResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfCreateEchoResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfCreateEchoResponseData is nil"))
  }
  switch v.(type) {
    case import2.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = v.(import2.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case Echo:
      if nil == p.oneOfType2080 {p.oneOfType2080 = new(Echo)}
      *p.oneOfType2080 = v.(Echo)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType2080.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType2080.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfCreateEchoResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if p.oneOfType2080 != nil && *p.oneOfType2080.ObjectType_ == *p.Discriminator {
    return *p.oneOfType2080
  }
  return nil
}

func (p *OneOfCreateEchoResponseData) UnmarshalJSON(b []byte) error {
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
  vOneOfType2080 := new(Echo)
  if err := json.Unmarshal(b, vOneOfType2080); err == nil {
    if "petstore.v1.utils.Echo" == *vOneOfType2080.ObjectType_ {
          if nil == p.oneOfType2080 {p.oneOfType2080 = new(Echo)}
      *p.oneOfType2080 = *vOneOfType2080
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType2080.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType2080.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateEchoResponseData"))
}

func (p *OneOfCreateEchoResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if p.oneOfType2080 != nil && *p.oneOfType2080.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType2080)
  }
  return nil, errors.New("No value to marshal for OneOfCreateEchoResponseData")
}

type OneOfTestUtilsResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
  oneOfType2080 *Bravo `json:"-"`
}

func NewOneOfTestUtilsResponseData() *OneOfTestUtilsResponseData {
  p := new(OneOfTestUtilsResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfTestUtilsResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfTestUtilsResponseData is nil"))
  }
  switch v.(type) {
    case import2.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = v.(import2.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case Bravo:
      if nil == p.oneOfType2080 {p.oneOfType2080 = new(Bravo)}
      *p.oneOfType2080 = v.(Bravo)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType2080.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType2080.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfTestUtilsResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if p.oneOfType2080 != nil && *p.oneOfType2080.ObjectType_ == *p.Discriminator {
    return *p.oneOfType2080
  }
  return nil
}

func (p *OneOfTestUtilsResponseData) UnmarshalJSON(b []byte) error {
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
  vOneOfType2080 := new(Bravo)
  if err := json.Unmarshal(b, vOneOfType2080); err == nil {
    if "petstore.v1.utils.Bravo" == *vOneOfType2080.ObjectType_ {
          if nil == p.oneOfType2080 {p.oneOfType2080 = new(Bravo)}
      *p.oneOfType2080 = *vOneOfType2080
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType2080.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType2080.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfTestUtilsResponseData"))
}

func (p *OneOfTestUtilsResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if p.oneOfType2080 != nil && *p.oneOfType2080.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType2080)
  }
  return nil, errors.New("No value to marshal for OneOfTestUtilsResponseData")
}

type OneOfMessageApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType2080 *import1.Message `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfMessageApiResponseData() *OneOfMessageApiResponseData {
  p := new(OneOfMessageApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfMessageApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfMessageApiResponseData is nil"))
  }
  switch v.(type) {
    case import1.Message:
      if nil == p.oneOfType2080 {p.oneOfType2080 = new(import1.Message)}
      *p.oneOfType2080 = v.(import1.Message)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType2080.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType2080.ObjectType_
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

func (p *OneOfMessageApiResponseData) GetValue() interface{} {
  if p.oneOfType2080 != nil && *p.oneOfType2080.ObjectType_ == *p.Discriminator {
    return *p.oneOfType2080
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfMessageApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType2080 := new(import1.Message)
  if err := json.Unmarshal(b, vOneOfType2080); err == nil {
    if "common.v1.config.Message" == *vOneOfType2080.ObjectType_ {
          if nil == p.oneOfType2080 {p.oneOfType2080 = new(import1.Message)}
      *p.oneOfType2080 = *vOneOfType2080
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType2080.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType2080.ObjectType_
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
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfMessageApiResponseData"))
}

func (p *OneOfMessageApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType2080 != nil && *p.oneOfType2080.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType2080)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfMessageApiResponseData")
}

