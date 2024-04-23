package repository

import (
	"github.com/omidfth/testish"
	"github.com/omidfth/testish/internal/types/serviceNames"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

func Test_testRepo_GetFirst(t *testing.T) {
	test := testish.NewTestish(testish.NewOption(serviceNames.POSTGRESQL, 5432, "./../cmd/postgres_dump.sql"))
	defer test.Close()

	tests := []struct {
		name      string
		db        *gorm.DB
		want      string
		wantError bool
	}{
		{
			name:      "test 1",
			db:        test.GetDB(serviceNames.POSTGRESQL),
			want:      "omid",
			wantError: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &testRepo{
				db: tt.db,
			}
			assert.Equal(t, tt.want, r.GetFirst().Name)
		})
	}
}
