package folders

import "github.com/gofrs/uuid"

type FetchFolderRequest struct {
	OrgID uuid.UUID
}

type FetchFolderResponse struct {
	Folders []*Folder
}

type FetchFolderRequestPaginated struct {
	OrgID uuid.UUID
	Page int
	PageSize int
}

type FetchFolderResponsePaginated struct {
	Folders []*Folder
}