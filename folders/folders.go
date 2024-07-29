package folders

import (
	"github.com/gofrs/uuid"
)

// GetAllFolders takes in a FetchFolderRequest which contains an organization ID.
// It returns an array of folder references from the organization ID.

// The function can improved by:
// - handling the nil error
// - removing the unused variables (which actually cause compilation errors)
// - removing the unnecessary dereferencing and referencing of the folders
// - using better variable names
func GetAllFolders(req *FetchFolderRequest) (*FetchFolderResponse, error) {
	var (
		err error
		f1  Folder
		fs  []*Folder
	)
	f := []Folder{}
	r, _ := FetchAllFoldersByOrgID(req.OrgID)
	for k, v := range r {
		f = append(f, *v)
	}
	var fp []*Folder
	for k1, v1 := range f {
		fp = append(fp, &v1)
	}
	var ffr *FetchFolderResponse
	ffr = &FetchFolderResponse{Folders: fp}
	return ffr, nil
}

// FetchAllFoldersByOrgID fetches folders by the organization ID
// It fetches all folders and returns those that match the input organization ID.
func FetchAllFoldersByOrgID(orgID uuid.UUID) ([]*Folder, error) {
	folders := GetSampleData()

	resFolder := []*Folder{}
	for _, folder := range folders {
		if folder.OrgId == orgID {
			resFolder = append(resFolder, folder)
		}
	}
	return resFolder, nil
}
