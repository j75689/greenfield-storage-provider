package types

type readableUploadProgressType uint32

const (
	UploadProgressMetaCreated     readableUploadProgressType = 0
	UploadProgressDataUploading   readableUploadProgressType = 1
	UploadProgressDataReplicating readableUploadProgressType = 2
	UploadProgressSealing         readableUploadProgressType = 3
	UploadProgressCompleted       readableUploadProgressType = 4
	UploadProgressFailed          readableUploadProgressType = 5
)

// ToReadableDescription convects readable type to the description string
var ToReadableDescription = map[readableUploadProgressType]string{
	UploadProgressMetaCreated:     "object meta is created in the chain, but the upload has been not started yet",
	UploadProgressDataUploading:   "object payload is uploading to the primary SP",
	UploadProgressDataReplicating: "object payload is replicating to the secondary SPs in the background",
	UploadProgressSealing:         "object meta is sealing onto the chain in the background",
	UploadProgressCompleted:       "object is succeed to upload",
	UploadProgressFailed:          "something is wrong in the upload progress",
}

// StateToProgressType convents inner state to the readable type
var StateToProgressType = map[TaskState]readableUploadProgressType{
	TaskState_TASK_STATE_INIT_UNSPECIFIED:       UploadProgressDataUploading,
	TaskState_TASK_STATE_UPLOAD_OBJECT_DOING:    UploadProgressDataUploading,
	TaskState_TASK_STATE_UPLOAD_OBJECT_DONE:     UploadProgressDataReplicating,
	TaskState_TASK_STATE_UPLOAD_OBJECT_ERROR:    UploadProgressFailed,
	TaskState_TASK_STATE_ALLOC_SECONDARY_DOING:  UploadProgressDataReplicating,
	TaskState_TASK_STATE_ALLOC_SECONDARY_DONE:   UploadProgressDataReplicating,
	TaskState_TASK_STATE_ALLOC_SECONDARY_ERROR:  UploadProgressFailed,
	TaskState_TASK_STATE_REPLICATE_OBJECT_DOING: UploadProgressDataReplicating,
	TaskState_TASK_STATE_REPLICATE_OBJECT_DONE:  UploadProgressDataReplicating,
	TaskState_TASK_STATE_REPLICATE_OBJECT_ERROR: UploadProgressFailed,
	TaskState_TASK_STATE_SIGN_OBJECT_DOING:      UploadProgressSealing,
	TaskState_TASK_STATE_SIGN_OBJECT_DONE:       UploadProgressSealing,
	TaskState_TASK_STATE_SIGN_OBJECT_ERROR:      UploadProgressFailed,
	TaskState_TASK_STATE_SEAL_OBJECT_DOING:      UploadProgressSealing,
	TaskState_TASK_STATE_SEAL_OBJECT_DONE:       UploadProgressCompleted,
	TaskState_TASK_STATE_SEAL_OBJECT_ERROR:      UploadProgressFailed,
}

// StateToDescription convents state to description.
func StateToDescription(state TaskState) string {
	uploadProgressType, ok := StateToProgressType[state]
	if !ok {
		return state.String()
	}
	description, ok := ToReadableDescription[uploadProgressType]
	if !ok {
		return state.String()
	}
	return description
}
