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

// - also could rename id to be ID or Id everywhere
func GetAllFolders(req *FetchFolderRequest) (*FetchFolderResponse, error) {
	folders, err := FetchAllFoldersByOrgID(req.OrgID)
	if err != nil {
		return nil, err
	}

	response := &FetchFolderResponse{Folders: folders}
	return response, nil
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
