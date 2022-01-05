/*
 * Generated file models/prism/v4/config/config_model.go.
 *
 * Product version: 16.2.0-SNAPSHOT
 *
 * Part of the Tasks Versioned APIs
 *
 * (c) 2022 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module prism.v4.config of Tasks Versioned APIs
*/
package config
import (
  "bytes"
  import1 "github.com/nishantikagarg/sample-go-repo/tasks_go_sdk/v16/models/common/v1/config"
  import3 "github.com/nishantikagarg/sample-go-repo/tasks_go_sdk/v16/models/common/v1/response"
  "encoding/json"
  "errors"
  "fmt"
  import2 "github.com/nishantikagarg/sample-go-repo/tasks_go_sdk/v16/models/prism/v4/error"
  "time"
)

/**
Details of entity.
*/
type EntityReference struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Globally unique identifier of the entity.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  Entity type identified as 'namespace:module[:submodule]:entityType such vmm:ahv:vm'
  */
  Rel *string `json:"rel,omitempty"`
}

func NewEntityReference() *EntityReference {
  p := new(EntityReference)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "prism.v4.config.EntityReference"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a1.config.EntityReference"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Reference to the owner of the task.
*/
type OwnerReference struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Globally unique identifier of the task owner.
  */
  ExtId *string `json:"extId,omitempty"`
}

func NewOwnerReference() *OwnerReference {
  p := new(OwnerReference)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "prism.v4.config.OwnerReference"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a1.config.OwnerReference"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
The task object tracking an async operation.
*/
type Task struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  UTC date and time in RFC-3339 format when task was completed.
  */
  CompletedTime *time.Time `json:"completedTime,omitempty"`
  /**
  Additional details post operation completion to guide further actions.
  */
  CompletionDetails []import1.KVPair `json:"completionDetails,omitempty"`
  /**
  UTC date and time in RFC-3339 format when task was created.
  */
  CreatedTime *time.Time `json:"createdTime,omitempty"`
  /**
  Reference to entities associated with the task.
  */
  EntitiesAffected []EntityReference `json:"entitiesAffected,omitempty"`
  /**
  Error details in case of task failure.
  */
  ErrorMessages []import2.AppMessage `json:"errorMessages,omitempty"`
  /**
  Globally unique idntifier of a task.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  Signifies if task can be cancelled.
  */
  IsCancelable *bool `json:"isCancelable,omitempty"`
  /**
  Provides the error in the absence of well defined error message for tasks created via legacy APIs.
  */
  LegacyErrorMessage *string `json:"legacyErrorMessage,omitempty"`
  /**
  Operation being tracked by the task.
  */
  Operation *string `json:"operation,omitempty"`
  /**
  Description of the operation being tracked by the task.
  */
  OperationDescription *string `json:"operationDescription,omitempty"`
  
  OwnedBy *OwnerReference `json:"ownedBy,omitempty"`
  
  ParentTask *TaskReferenceInternal `json:"parentTask,omitempty"`
  /**
  Task progress expressed as a percentage.
  */
  ProgressPercentage *int `json:"progressPercentage,omitempty"`
  /**
  UTC date and time in RFC-3339 format when task was started.
  */
  StartedTime *time.Time `json:"startedTime,omitempty"`
  
  Status *TaskStatus `json:"status,omitempty"`
  /**
  List of steps completed as part of task.
  */
  SubSteps []TaskStep `json:"subSteps,omitempty"`
  /**
  Reference to tasks spawned as children of current task.
  */
  SubTasks []TaskReferenceInternal `json:"subTasks,omitempty"`
}

func NewTask() *Task {
  p := new(Task)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "prism.v4.config.Task"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a1.config.Task"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST Response for all response codes in api path /prism/v4.0.a1/config/tasks/{taskExtId}/$actions/cancel Post operation
*/
type TaskCancelResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfTaskCancelResponseData `json:"data,omitempty"`
  
  Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewTaskCancelResponse() *TaskCancelResponse {
  p := new(TaskCancelResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "prism.v4.config.TaskCancelResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a1.config.TaskCancelResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *TaskCancelResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *TaskCancelResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfTaskCancelResponseData()
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
REST Response for all response codes in api path /prism/v4.0.a1/config/tasks/{taskExtId} Get operation
*/
type TaskGetResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfTaskGetResponseData `json:"data,omitempty"`
  
  Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewTaskGetResponse() *TaskGetResponse {
  p := new(TaskGetResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "prism.v4.config.TaskGetResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a1.config.TaskGetResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *TaskGetResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *TaskGetResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfTaskGetResponseData()
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
Reference to a task tracking the async operation.
*/
type TaskReference struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Globally unique idntifier of a task.
  */
  ExtId *string `json:"extId,omitempty"`
}

func NewTaskReference() *TaskReference {
  p := new(TaskReference)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "prism.v4.config.TaskReference"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a1.config.TaskReference"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Reference to task associated with current task.
*/
type TaskReferenceInternal struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Glocally unique identifier of the task.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  The URL at which the entity described by this link can be accessed.
  */
  Href *string `json:"href,omitempty"`
  /**
  A name that identifies the relationship of this link to the object that is returned by the URL.  The special value of "self" identifies the URL for the object.
  */
  Rel *string `json:"rel,omitempty"`
}

func NewTaskReferenceInternal() *TaskReferenceInternal {
  p := new(TaskReferenceInternal)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "prism.v4.config.TaskReferenceInternal"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a1.config.TaskReferenceInternal"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Status of the task.
*/
type TaskStatus int

const(
  TASKSTATUS_UNKNOWN TaskStatus = 0
  TASKSTATUS_REDACTED TaskStatus = 1
  TASKSTATUS_QUEUED TaskStatus = 2
  TASKSTATUS_RUNNING TaskStatus = 3
  TASKSTATUS_CANCELING TaskStatus = 4
  TASKSTATUS_SUCCEEDED TaskStatus = 5
  TASKSTATUS_FAILED TaskStatus = 6
  TASKSTATUS_CANCELED TaskStatus = 7
)

// returns the name of the enum given an ordinal number
func (e *TaskStatus) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "QUEUED",
    "RUNNING",
    "CANCELING",
    "SUCCEEDED",
    "FAILED",
    "CANCELED",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *TaskStatus) index(name string) TaskStatus {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "QUEUED",
    "RUNNING",
    "CANCELING",
    "SUCCEEDED",
    "FAILED",
    "CANCELED",
  }
  for idx := range names {
    if names[idx] == name {
      return TaskStatus(idx)
    }
  }
  return TASKSTATUS_UNKNOWN
}

func (e *TaskStatus) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for TaskStatus:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *TaskStatus) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e TaskStatus) Ref() *TaskStatus {
  return &e
}


/**
A single step in the task.
*/
type TaskStep struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Message describing the step completed of the task.
  */
  Name *string `json:"name,omitempty"`
}

func NewTaskStep() *TaskStep {
  p := new(TaskStep)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "prism.v4.config.TaskStep"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a1.config.TaskStep"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}



type OneOfTaskCancelResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
  oneOfType2001 *interface{} `json:"-"`
}

func NewOneOfTaskCancelResponseData() *OneOfTaskCancelResponseData {
  p := new(OneOfTaskCancelResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfTaskCancelResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfTaskCancelResponseData is nil"))
  }
  if nil == v {
    if nil == p.oneOfType2001 {p.oneOfType2001 = new(interface {})}
    *p.oneOfType2001 = nil
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

func (p *OneOfTaskCancelResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if "EMPTY" == *p.Discriminator {
    return *p.oneOfType2001
  }
  return nil
}

func (p *OneOfTaskCancelResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType2001 := new(interface {})
  if err := json.Unmarshal(b, vOneOfType2001); err == nil {
    if nil == *vOneOfType2001 {
      if nil == p.oneOfType2001 {p.oneOfType2001 = new(interface {})}
      *p.oneOfType2001 = nil
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "EMPTY"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "EMPTY"
      return nil
    }
  }
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "prism.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfTaskCancelResponseData"))
}

func (p *OneOfTaskCancelResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if "EMPTY" == *p.Discriminator {
    return json.Marshal(p.oneOfType2001)
  }
  return nil, errors.New("No value to marshal for OneOfTaskCancelResponseData")
}

type OneOfTaskGetResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType2001 *Task `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfTaskGetResponseData() *OneOfTaskGetResponseData {
  p := new(OneOfTaskGetResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfTaskGetResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfTaskGetResponseData is nil"))
  }
  switch v.(type) {
    case Task:
      if nil == p.oneOfType2001 {p.oneOfType2001 = new(Task)}
      *p.oneOfType2001 = v.(Task)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType2001.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType2001.ObjectType_
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

func (p *OneOfTaskGetResponseData) GetValue() interface{} {
  if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
    return *p.oneOfType2001
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfTaskGetResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType2001 := new(Task)
  if err := json.Unmarshal(b, vOneOfType2001); err == nil {
    if "prism.v4.config.Task" == *vOneOfType2001.ObjectType_ {
          if nil == p.oneOfType2001 {p.oneOfType2001 = new(Task)}
      *p.oneOfType2001 = *vOneOfType2001
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType2001.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType2001.ObjectType_
      return nil
    }
    }
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "prism.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfTaskGetResponseData"))
}

func (p *OneOfTaskGetResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType2001)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfTaskGetResponseData")
}

