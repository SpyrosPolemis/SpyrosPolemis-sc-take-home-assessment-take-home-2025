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

	folderData := []folder.Folder{
		{Name: "alpha", Paths: "alpha", OrgId: defaultOrgID},
		{Name: "bravo", Paths: "alpha.bravo", OrgId: defaultOrgID},
		{Name: "charlie", Paths: "alpha.bravo.charlie", OrgId: defaultOrgID},
		{Name: "delta", Paths: "alpha.delta", OrgId: defaultOrgID},
		{Name: "echo", Paths: "alpha.delta.echo", OrgId: defaultOrgID},
		{Name: "foxtrot", Paths: "foxtrot", OrgId: uuid.Must(uuid.NewV4())}, // Different org
		{Name: "golf", Paths: "golf", OrgId: defaultOrgID},
	}

	tests := [...]struct {
		name         string
		folders      []folder.Folder
		folderToMove string
		dst          string
		wantErr      bool
		want         []folder.Folder
	}{
		{
			name:         "Move folder with children",
			folders:      folderData,
			folderToMove: "bravo",
			dst:          "delta",
			wantErr:      false,
			want: []folder.Folder{
				{Name: "alpha", Paths: "alpha", OrgId: defaultOrgID},
				{Name: "bravo", Paths: "alpha.delta.bravo", OrgId: defaultOrgID},
				{Name: "charlie", Paths: "alpha.delta.bravo.charlie", OrgId: defaultOrgID},
				{Name: "delta", Paths: "alpha.delta", OrgId: defaultOrgID},
				{Name: "echo", Paths: "alpha.delta.echo", OrgId: defaultOrgID},
				{Name: "foxtrot", Paths: "foxtrot", OrgId: uuid.Must(uuid.NewV4())}, // Different org
				{Name: "golf", Paths: "golf", OrgId: defaultOrgID},
			},
		},
		{
			name:         "Move folder into its child",
			folders:      folderData,
			folderToMove: "bravo",
			dst:          "charlie",
			wantErr:      true,
			want:         nil,
		},
		{
			name:         "Move folder into itself",
			folders:      folderData,
			folderToMove: "bravo",
			dst:          "bravo",
			wantErr:      true,
			want:         nil,
		},
		{
			name:         "Move folder to different org",
			folders:      folderData,
			folderToMove: "bravo",
			dst:          "foxtrot",
			wantErr:      true,
			want:         nil,
		},
		{
			name:         "Move folder that does not exist",
			folders:      folderData,
			folderToMove: "invalid_folder",
			dst:          "bravo",
			wantErr:      true,
			want:         nil,
		},
		{
			name:         "Move folder into destination that does not exist",
			folders:      folderData,
			folderToMove: "bravo",
			dst:          "invalid_folder",
			wantErr:      true,
			want:         nil,
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
