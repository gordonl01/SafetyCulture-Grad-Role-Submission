package folders

import (
	"errors"

	"github.com/gofrs/uuid"
)

// Obtains all folders with the same OrgID as the the request.
// Loops through returned folders, dereferences each value and appends the values to a new slice of Folders, f.
// Loops through f and creates a new pointer to each value so that the original values are not changed. Each pointer is added to a new slice of pointers that reference a Folder.
// Then, a FetchFolderResponse is created and returned.
// Suggestion 1: Do some error handling by not omitting the error returned by FetchAllFoldersByOrgID and adding some code to check for errors
// Suggestion 2: Use the uninitialised variables in the beginning of the function instead of creating new variables to save memory. Instead of 'fp', we can use 'fs'.
// Suggestion 3: Instead of creating new pointers in two different loops, they can be combined to remove the need for the extra 'f' variable. Thus, saving memory.
// Suggestion 4: Omit the index in the loop as they are not being used.
// Suggestion 5: Declare 'ffr' variable inside var block.
func GetAllFolders(req *FetchFolderRequest) (*FetchFolderResponse, error) {
	if req == nil {
		return nil, errors.New("request is undefined")
	}
	var (
		err error
		f1  Folder
		fs  []*Folder
		ffr *FetchFolderResponse
	)

	r, err := FetchAllFoldersByOrgID(req.OrgID)
	if err != nil {
		return ffr, err
	}
	for _, v := range r {
		f1 = *v
		fs = append(fs, &f1)
	}
	ffr = &FetchFolderResponse{Folders: fs}
	return ffr, nil
}

// Reads data from sample.json and returns a slice of Folders
// The code then loops through the folders variable and appends a Folder to resFolder only if the Folder OrgId matches the orgID given in the function parameter
// Then, it returns the result.
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
