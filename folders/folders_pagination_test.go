package folders_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folders"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_GetAllFoldersPagination(t *testing.T) {
	t.Run("ValidPagination", func(t *testing.T) {
		orgID, err := uuid.FromString(folders.DefaultOrgID)
		assert.Nil(t, err)

		req := &folders.FetchFolderRequestPaginated{orgID, 1, 10}
		folders, err := folders.GetAllFoldersPaginated(req)
		assert.Nil(t, err)
		assert.NotNil(t, folders)
		assert.Equal(t, 10, len(folders.Folders))

		for _, folder := range folders.Folders {
			assert.Equal(t, orgID, folder.OrgId)
		}
	})

	t.Run("InvalidPageNumber", func(t *testing.T) {
		orgID, err := uuid.FromString(folders.DefaultOrgID)
		assert.Nil(t, err)

		req := &folders.FetchFolderRequestPaginated{orgID, 0, 10}
		_, err = folders.GetAllFoldersPaginated(req)
		assert.NotNil(t, err)
	})

	t.Run("InvalidPageSize", func(t *testing.T) {
		orgID, err := uuid.FromString(folders.DefaultOrgID)
		assert.Nil(t, err)

		req := &folders.FetchFolderRequestPaginated{orgID, 1, 0}
		_, err = folders.GetAllFoldersPaginated(req)
		assert.NotNil(t, err)
	})
}