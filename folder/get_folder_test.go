package folder_test

import (
	"reflect"
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	// "github.com/stretchr/testify/assert"
)

// feel free to change how the unit test is structured
func Test_folder_GetFoldersByOrgID(t *testing.T) {
	t.Parallel()
	tests := [...]struct {
		name    string
		orgID   uuid.UUID
		folders []folder.Folder
		want    []folder.Folder
	}{
		// TODO: your tests here
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// f := folder.NewDriver(tt.folders)
			// get := f.GetFoldersByOrgID(tt.orgID)

		})
	}
}

func Test_folder_GetAllChildFolders(t *testing.T) {
	t.Parallel()

	defaultOrgID := uuid.FromStringOrNil(folder.DefaultOrgID)

	tests := []struct {
		name       string
		orgID      uuid.UUID
		folders    []folder.Folder
		parentName string
		wantErr    bool
		want       []folder.Folder
	}{
		{
			name:  "Multiple child folders present",
			orgID: defaultOrgID,
			folders: []folder.Folder{
				{Name: "parent-folder", OrgId: defaultOrgID, Paths: "parent-folder"},
				{Name: "child1", OrgId: defaultOrgID, Paths: "parent-folder.child1"},
				{Name: "child2", OrgId: defaultOrgID, Paths: "parent-folder.child2"},
				{Name: "child3", OrgId: defaultOrgID, Paths: "parent-folder.child2.child3"},
			},
			parentName: "parent-folder",
			wantErr:    false,
			want: []folder.Folder{
				{Name: "child1", OrgId: defaultOrgID, Paths: "parent-folder.child1"},
				{Name: "child2", OrgId: defaultOrgID, Paths: "parent-folder.child2"},
				{Name: "child3", OrgId: defaultOrgID, Paths: "parent-folder.child2.child3"},
			},
		},
		{
			name:  "No child folders present",
			orgID: defaultOrgID,
			folders: []folder.Folder{
				{Name: "parent-folder", OrgId: defaultOrgID, Paths: "parent-folder"},
				{Name: "non-child1", OrgId: defaultOrgID, Paths: "non-child1"},
				{Name: "non-child2", OrgId: defaultOrgID, Paths: "non-child1.non-child2"},
				{Name: "non-child3", OrgId: defaultOrgID, Paths: "non-child1.non-child2.non-child3"},
				{Name: "non-child4", OrgId: defaultOrgID, Paths: "non-child4"},
			},
			parentName: "parent-folder",
			wantErr:    false,
			want:       []folder.Folder{},
		},
		{
			name:  "Parent folder does not exist",
			orgID: defaultOrgID,
			folders: []folder.Folder{
				{Name: "folder1", OrgId: defaultOrgID, Paths: "folder1"},
				{Name: "folder2", OrgId: defaultOrgID, Paths: "folder2"},
				{Name: "folder3", OrgId: defaultOrgID, Paths: "folder3"},
			},
			parentName: "parent-folder",
			wantErr:    true,
			want:       nil,
		},
		{
			name:  "Parent folder exists in different org",
			orgID: uuid.Must(uuid.NewV4()),
			folders: []folder.Folder{
				{Name: "parent-folder", OrgId: defaultOrgID, Paths: "folder1"},
				{Name: "folder2", OrgId: defaultOrgID, Paths: "folder2"},
			},
			parentName: "parent-folder",
			wantErr:    true,
			want:       nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)

			// Call GetAllChildFolders method
			got, err := f.GetAllChildFolders(tt.orgID, tt.parentName)

			// Check if an error is expected
			if (err != nil) != tt.wantErr || tt.wantErr {
				t.Errorf("GetAllChildFolders() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil {
				t.Errorf("GetAllChildFolders() error = %v", err)
			}

			// Compare the result with the expected folders
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllChildFolders() = %v, want %v", got, tt.want)
			}
		})
	}
}
