package folders

import (
	"errors"

	"github.com/gofrs/uuid"
)

// Explanation: Using the Id field in the Folder type, the code keeps track of the last folder that it sent.
// Then, the 'pagNum' variable determines how many more folders to send back as a response.
// The code loops through the folders and obtains a number of folders after the last folder sent.
// After this, a new Id is returned so that the client has a way of obtaining more folders after the last folder sent.
func GetFolders(req *FetchFolderRequest, chunkId uuid.UUID) (*FetchFolderResponse, uuid.UUID, error) {
	if req == nil {
		return nil, chunkId, errors.New("request is undefined")
	}
	var (
		err        error
		f1         Folder
		fs         []*Folder
		ffr        *FetchFolderResponse
		newChunkId uuid.UUID
	)

	r, err := FetchFoldersByOrgID(req.OrgID, chunkId)
	if err != nil {
		return ffr, chunkId, err
	}
	pagNum := 5
	for idx, folder := range r {
		if folder.Id == chunkId {
			newChunkId = r[idx+pagNum].Id
			r = r[idx : idx+pagNum]
			break
		}
	}
	for _, v := range r {
		f1 = *v
		fs = append(fs, &f1)
	}
	ffr = &FetchFolderResponse{Folders: fs}
	return ffr, newChunkId, nil
}

func FetchFoldersByOrgID(orgID uuid.UUID, chunkId uuid.UUID) ([]*Folder, error) {

	folders := GetSampleData()
	resFolder := []*Folder{}

	for _, folder := range folders {
		if folder.OrgId == orgID {
			resFolder = append(resFolder, folder)
		}
	}
	return resFolder, nil
}
