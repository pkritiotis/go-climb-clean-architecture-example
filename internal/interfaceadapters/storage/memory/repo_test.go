package memory

import (
	"github.com/google/uuid"
	"github.com/pkritiotis/go-climb-clean-architecture-example/internal/domain/crag"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewRepo(t *testing.T) {
	tests := []struct {
		name string
		want crag.Repository
	}{
		{
			name: "Should create an inmemory memory",
			want: Repo{
				crags: make(map[string]crag.Crag),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewRepo()
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_inMemoryRepo_AddCrag(t *testing.T) {
	mockUUID := uuid.MustParse("3e204a57-4449-4c74-8227-77934cf25322")

	type fields struct {
		crags map[string]crag.Crag
	}
	type args struct {
		crag crag.Crag
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "should add crag",
			fields: fields{
				crags: make(map[string]crag.Crag),
			},
			args: args{
				crag: crag.Crag{ID: mockUUID},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Repo{
				crags: tt.fields.crags,
			}
			err := m.Add(tt.args.crag)
			assert.Equal(t, tt.wantErr, err != nil)
			c, _ := m.GetByID(mockUUID)
			assert.Equal(t, *c, tt.args.crag)
		})
	}
}

func Test_inMemoryRepo_DeleteCrag(t *testing.T) {
	mockUUID := uuid.MustParse("3e204a57-4449-4c74-8227-77934cf25322")

	type fields struct {
		crags map[string]crag.Crag
	}
	type args struct {
		id uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "crag exists - should delete existing crag",
			fields: fields{
				crags: func() map[string]crag.Crag {
					mp := make(map[string]crag.Crag)
					mp[mockUUID.String()] = crag.Crag{ID: mockUUID}
					return mp
				}(),
			},
			args:    args{id: mockUUID},
			wantErr: false,
		},
		{
			name: "crag does not exist - should return error",
			fields: fields{
				crags: make(map[string]crag.Crag),
			},
			args:    args{id: mockUUID},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Repo{
				crags: tt.fields.crags,
			}
			err := m.Delete(tt.args.id)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func Test_inMemoryRepo_GetCrag(t *testing.T) {
	mockUUID := uuid.MustParse("3e204a57-4449-4c74-8227-77934cf25322")

	type fields struct {
		crags map[string]crag.Crag
	}
	type args struct {
		id uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *crag.Crag
		wantErr bool
	}{
		{
			name: "crag exists, should return crag",
			fields: fields{
				crags: func() map[string]crag.Crag {
					mp := make(map[string]crag.Crag)
					mp[mockUUID.String()] = crag.Crag{ID: mockUUID}
					return mp
				}(),
			},
			args: args{
				id: mockUUID,
			},
			want:    &crag.Crag{ID: mockUUID},
			wantErr: false,
		},
		{
			name: "crag does not exists, should return nil",
			fields: fields{
				crags: make(map[string]crag.Crag),
			},
			args: args{
				id: mockUUID,
			},
			want:    (*crag.Crag)(nil),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Repo{
				crags: tt.fields.crags,
			}
			got, err := m.GetByID(tt.args.id)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_inMemoryRepo_GetCrags(t *testing.T) {
	mockUUID1 := uuid.MustParse("3e204a57-4449-4c74-8227-77934cf25322")
	mockUUID2 := uuid.MustParse("4e204a57-4449-4c74-8227-77934cf25322")

	type fields struct {
		crags map[string]crag.Crag
	}
	tests := []struct {
		name    string
		fields  fields
		want    []crag.Crag
		wantErr bool
	}{
		{
			name: "should return 2 crags",
			fields: fields{
				crags: func() map[string]crag.Crag {
					mp := make(map[string]crag.Crag)
					mp[mockUUID1.String()] = crag.Crag{ID: mockUUID1}
					mp[mockUUID2.String()] = crag.Crag{ID: mockUUID2}
					return mp
				}(),
			},
			want:    []crag.Crag{{ID: mockUUID1}, {ID: mockUUID2}},
			wantErr: false,
		},
		{
			name: "should return 0 crags",
			fields: fields{
				crags: make(map[string]crag.Crag),
			},
			want:    ([]crag.Crag)(nil),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Repo{
				crags: tt.fields.crags,
			}
			got, err := m.GetAll()
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_inMemoryRepo_UpdateCrag(t *testing.T) {
	mockUUID := uuid.MustParse("3e204a57-4449-4c74-8227-77934cf25322")

	type fields struct {
		crags map[string]crag.Crag
	}
	type args struct {
		crag crag.Crag
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "crag exists - should update crag",
			fields: fields{
				crags: func() map[string]crag.Crag {
					mp := make(map[string]crag.Crag)
					mp[mockUUID.String()] = crag.Crag{ID: mockUUID, Name: "old"}
					return mp
				}(),
			},
			args: args{
				crag: crag.Crag{ID: mockUUID, Name: "new"},
			},
			wantErr: false,
		},
		{
			name: "crag does not exist - should return error",
			fields: fields{
				crags: make(map[string]crag.Crag),
			},
			args: args{
				crag: crag.Crag{ID: mockUUID, Name: "new"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Repo{
				crags: tt.fields.crags,
			}
			err := m.Update(tt.args.crag)
			assert.Equal(t, tt.wantErr, err != nil)
			if err != nil {
				c, _ := m.GetByID(mockUUID)
				assert.Equal(t, tt.args.crag.Name, c.Name)
			}
		})
	}
}
