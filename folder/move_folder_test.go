package folder_test

import (
	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	"reflect"
	"testing"
)

func Test_folder_MoveFolder(t *testing.T) {
	t.Parallel()

	defaultOrgID := uuid.FromStringOrNil(folder.DefaultOrgID)

	tests := [...]struct {
		name         string
		folders      []folder.Folder
		folderToMove string
		dst          string
		wantErr      bool
		want         []folder.Folder
	}{
		{
			name:    "Move folder without children",
			folders: []folder.Folder{},
			dst:     "parent",
			wantErr: false,
			want:    []folder.Folder{},
		},
		{
			name:    "Move folder with children",
			folders: []folder.Folder{},
			dst:     "parent",
			wantErr: false,
			want:    []folder.Folder{},
		},
		{
			name:    "Move folder into its child",
			folders: []folder.Folder{},
			dst:     "parent",
			wantErr: true,
			want:    []folder.Folder{},
		},
		{
			name:    "Move folder into itself",
			folders: []folder.Folder{},
			dst:     "parent",
			wantErr: true,
			want:    []folder.Folder{},
		},
		{
			name:    "Move folder into folder in different org",
			folders: []folder.Folder{},
			dst:     "parent",
			wantErr: true,
			want:    []folder.Folder{},
		},
		{
			name:    "Move folder that does not exist",
			folders: []folder.Folder{},
			dst:     "parent",
			wantErr: true,
			want:    []folder.Folder{},
		},
		{
			name:    "Move folder into destination that does not exist",
			folders: []folder.Folder{},
			dst:     "parent",
			wantErr: true,
			want:    []folder.Folder{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)

			// Call MoveFolders method
			got, err := f.MoveFolder(tt.folderToMove, tt.dst)

			// Check if an error is expected
			if (err != nil) != tt.wantErr || tt.wantErr {
				t.Errorf("MoveFolder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// Check if an error is present
			if err != nil {
				t.Errorf("MoveFolder() error = %v", err)
			}

			// Compare result with exected folders
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MoveFolder() got = %v, want %v", got, tt.want)
			}
		})
	}
}
