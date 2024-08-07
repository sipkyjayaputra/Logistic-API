package utils

type Status int32

const (
	Status_INVALID  Status = 0
	Status_REJECT   Status = -1
	Status_APPROVED Status = 1
)
