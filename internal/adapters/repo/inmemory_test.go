package repo

import (
	"github.com/google/uuid"
	"github.com/pkritiotis/go-clean/internal/app/services"
	"github.com/pkritiotis/go-clean/internal/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewInMemory(t *testing.T) {
	tests := []struct {
		name string
		want services.CragRepository
	}{
		{
			name: "Should create an inmemory repo",
			want: inMemoryRepo{
				crags: make(map[string]domain.Crag),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewInMemory()
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_inMemoryRepo_AddCrag(t *testing.T) {
	mockUUID := uuid.MustParse("3e204a57-4449-4c74-8227-77934cf25322")

	type fields struct {
		crags map[string]domain.Crag
	}
	type args struct {
		crag domain.Crag
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
				crags: make(map[string]domain.Crag),
			},
			args: args{
				crag: domain.Crag{ID: mockUUID},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := inMemoryRepo{
				crags: tt.fields.crags,
			}
			err := m.AddCrag(tt.args.crag)
			assert.Equal(t, tt.wantErr, err != nil)
			c, _ := m.GetCrag(mockUUID)
			assert.Equal(t, *c, tt.args.crag)
		})
	}
}

func Test_inMemoryRepo_DeleteCrag(t *testing.T) {
	mockUUID := uuid.MustParse("3e204a57-4449-4c74-8227-77934cf25322")

	type fields struct {
		crags map[string]domain.Crag
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
				crags: func() map[string]domain.Crag {
					mp := make(map[string]domain.Crag)
					mp[mockUUID.String()] = domain.Crag{ID: mockUUID}
					return mp
				}(),
			},
			args:    args{id: mockUUID},
			wantErr: false,
		},
		{
			name: "crag does not exist - should return error",
			fields: fields{
				crags: make(map[string]domain.Crag),
			},
			args:    args{id: mockUUID},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := inMemoryRepo{
				crags: tt.fields.crags,
			}
			err := m.DeleteCrag(tt.args.id)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func Test_inMemoryRepo_GetCrag(t *testing.T) {
	mockUUID := uuid.MustParse("3e204a57-4449-4c74-8227-77934cf25322")

	type fields struct {
		crags map[string]domain.Crag
	}
	type args struct {
		id uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.Crag
		wantErr bool
	}{
		{
			name: "crag exists, should return crag",
			fields: fields{
				crags: func() map[string]domain.Crag {
					mp := make(map[string]domain.Crag)
					mp[mockUUID.String()] = domain.Crag{ID: mockUUID}
					return mp
				}(),
			},
			args: args{
				id: mockUUID,
			},
			want:    &domain.Crag{ID: mockUUID},
			wantErr: false,
		},
		{
			name: "crag does not exists, should return nil",
			fields: fields{
				crags: make(map[string]domain.Crag),
			},
			args: args{
				id: mockUUID,
			},
			want:    (*domain.Crag)(nil),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := inMemoryRepo{
				crags: tt.fields.crags,
			}
			got, err := m.GetCrag(tt.args.id)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_inMemoryRepo_GetCrags(t *testing.T) {
	mockUUID1 := uuid.MustParse("3e204a57-4449-4c74-8227-77934cf25322")
	mockUUID2 := uuid.MustParse("4e204a57-4449-4c74-8227-77934cf25322")

	type fields struct {
		crags map[string]domain.Crag
	}
	tests := []struct {
		name    string
		fields  fields
		want    []domain.Crag
		wantErr bool
	}{
		{
			name: "should return 2 crags",
			fields: fields{
				crags: func() map[string]domain.Crag {
					mp := make(map[string]domain.Crag)
					mp[mockUUID1.String()] = domain.Crag{ID: mockUUID1}
					mp[mockUUID2.String()] = domain.Crag{ID: mockUUID2}
					return mp
				}(),
			},
			want:    []domain.Crag{{ID: mockUUID1}, {ID: mockUUID2}},
			wantErr: false,
		},
		{
			name: "should return 0 crags",
			fields: fields{
				crags: make(map[string]domain.Crag),
			},
			want:    ([]domain.Crag)(nil),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := inMemoryRepo{
				crags: tt.fields.crags,
			}
			got, err := m.GetCrags()
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_inMemoryRepo_UpdateCrag(t *testing.T) {
	mockUUID := uuid.MustParse("3e204a57-4449-4c74-8227-77934cf25322")

	type fields struct {
		crags map[string]domain.Crag
	}
	type args struct {
		crag domain.Crag
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
				crags: func() map[string]domain.Crag {
					mp := make(map[string]domain.Crag)
					mp[mockUUID.String()] = domain.Crag{ID: mockUUID, Name: "old"}
					return mp
				}(),
			},
			args: args{
				crag: domain.Crag{ID: mockUUID, Name: "new"},
			},
			wantErr: false,
		},
		{
			name: "crag does not exist - should return error",
			fields: fields{
				crags: make(map[string]domain.Crag),
			},
			args: args{
				crag: domain.Crag{ID: mockUUID, Name: "new"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := inMemoryRepo{
				crags: tt.fields.crags,
			}
			err := m.UpdateCrag(tt.args.crag)
			assert.Equal(t, tt.wantErr, err != nil)
			if err != nil {
				c, _ := m.GetCrag(mockUUID)
				assert.Equal(t, tt.args.crag.Name, c.Name)
			}
		})
	}
}
