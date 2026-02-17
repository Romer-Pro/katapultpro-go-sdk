package katapultpro

import (
	"github.com/romer-pro/katapultpro-go-sdk/v3/jobs"
	"github.com/romer-pro/katapultpro-go-sdk/v3/photos"
)

// Job status and path enums (re-exported from jobs).
type (
	JobStatus = jobs.JobStatus
	JobPath   = jobs.JobPath
)

const (
	JobStatusActive   = jobs.JobStatusActive
	JobStatusArchived = jobs.JobStatusArchived
)

const (
	JobPathName          = jobs.JobPathName
	JobPathJobCreator    = jobs.JobPathJobCreator
	JobPathJobOwner      = jobs.JobPathJobOwner
	JobPathProjectFolder = jobs.JobPathProjectFolder
	JobPathProjectID     = jobs.JobPathProjectID
	JobPathStatus        = jobs.JobPathStatus
	JobPathDone          = jobs.JobPathDone
	JobPathMapStyles     = jobs.JobPathMapStyles
	JobPathMetadata      = jobs.JobPathMetadata
	JobPathSharing       = jobs.JobPathSharing
)

// Photo association enums (re-exported from photos).
type (
	PhotoAssociationQuery  = photos.PhotoAssociationQuery
	PhotoAssociationValue  = photos.PhotoAssociationValue
)

const (
	PhotoAssociationQueryMain = photos.PhotoAssociationQueryMain
	PhotoAssociationQueryTrue = photos.PhotoAssociationQueryTrue
)

var (
	PhotoAssociationMain = photos.PhotoAssociationMain
	PhotoAssociationTrue = photos.PhotoAssociationTrue
)

// PtrPhotoAssociationMain returns a pointer suitable for AssociatePhotoRequest.AssociationValue ("main").
func PtrPhotoAssociationMain() *PhotoAssociationValue {
	return photos.PtrPhotoAssociationMain()
}

// PtrPhotoAssociationTrue returns a pointer suitable for AssociatePhotoRequest.AssociationValue (true).
func PtrPhotoAssociationTrue() *PhotoAssociationValue {
	return photos.PtrPhotoAssociationTrue()
}
