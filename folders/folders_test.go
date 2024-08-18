package folders_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folders"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

type Response interface {
}

func Test_GetAllFolders(t *testing.T) {

	t.Run("invalid id", func(t *testing.T) {
		req := &folders.FetchFolderRequest{
			OrgID: uuid.FromStringOrNil("test"),
		}
		result, _ := folders.GetAllFolders(req)
		assert.Empty(t, result.Folders)
	})
	t.Run("nil input", func(t *testing.T) {
		_, err := folders.GetAllFolders(nil)
		assert.EqualError(t, err, "request is undefined")
	})
	t.Run("normal test", func(t *testing.T) {
		req := &folders.FetchFolderRequest{
			OrgID: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
		}
		result, _ := folders.GetAllFolders(req)
		assert.Len(t, result.Folders, 666)
	})
	t.Run("empty id", func(t *testing.T) {
		req := &folders.FetchFolderRequest{
			OrgID: uuid.FromStringOrNil(""),
		}
		result, _ := folders.GetAllFolders(req)
		assert.Empty(t, result.Folders)
	})
}
