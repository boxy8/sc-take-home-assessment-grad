package folders_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folders"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_GetAllFolders(t *testing.T) {
	t.Run("ValidOrganizationId", func(t *testing.T) {
		orgID, err := uuid.FromString(folders.DefaultOrgID)
		assert.Nil(t, err)

		req := &folders.FetchFolderRequest{OrgID: orgID}
		folders, err := folders.GetAllFolders(req)
		assert.Nil(t, err)
		assert.NotNil(t, folders)

		for _, folder := range folders.Folders {
			assert.Equal(t, orgID, folder.OrgId)
		}
	})

	t.Run("InvalidOrganizationId", func(t *testing.T) {
		orgID := uuid.Nil
		folders, err := folders.GetAllFolders(&folders.FetchFolderRequest{OrgID: orgID})
		assert.Nil(t, err)
		assert.Empty(t, folders.Folders)
	})
}
