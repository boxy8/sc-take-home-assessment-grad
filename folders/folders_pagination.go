package folders

import (
	"errors"

	"github.com/gofrs/uuid"
)

// My pagination takes in a page number and page size.
// GetAllFoldersPaginated returns a list of folders from that page number
// according to the page size.
// It returns an error if the page number is less than 1, page size is less
// than 1, or if the page is out of bounds.
func GetAllFoldersPaginated(req *FetchFolderRequestPaginated) (*FetchFolderResponsePaginated, error) {
	page := req.Page
	pageSize := req.PageSize

	if page < 1 {
		return nil, errors.New("page must be greater than 0")
	}

	if pageSize < 1 {
		return nil, errors.New("page size must be greater than 0")
	}

	folders, err := FetchAllFoldersByOrgIDPaginated(req.OrgID, page, pageSize)
	if err != nil {
		return nil, err
	}

	response := &FetchFolderResponsePaginated{Folders: folders}
	return response, nil
}


func FetchAllFoldersByOrgIDPaginated(orgID uuid.UUID, page int, pageSize int) ([]*Folder, error) {
	folders := GetSampleData()

	resFolder := []*Folder{}
	for _, folder := range folders {
		if folder.OrgId == orgID {
			resFolder = append(resFolder, folder)
		}
	}

	totalFolders := len(resFolder)
	startIndex := (page - 1) * pageSize
	endIndex := page * pageSize

	if startIndex >= totalFolders {
		return nil, errors.New("page out of bounds")
	}

	if endIndex > totalFolders {
		endIndex = totalFolders
	}

	resFolder = resFolder[startIndex:endIndex]

	return resFolder, nil
}