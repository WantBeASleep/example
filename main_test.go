package trash_test

import (
	"testing"
	"trash"
	"os"
	"trash/internal/lib"
	mocklib "trash/mocks/lib"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDoSmtNotMock(t *testing.T) {
	srv := lib.NewService()
	content := "1"

	color, err := trash.DoSmt(srv, content)
	assert.NoError(t, err)
	assert.Equal(t, lib.Red, color)
}

func TestDoSmtWithMock(t *testing.T) {
	content := "123"

	// fileSrv := &mocklib.File{}
	fileSrv := mocklib.NewFile(t)

	fileSrv.EXPECT().FileColor(mock.MatchedBy(func (f *os.File) bool {
		openF, err := os.Open(f.Name())
		assert.NoError(t, err)
		stat, err := openF.Stat()
		assert.NoError(t, err)

		return assert.Equal(t, int64(3), stat.Size())
	})).Return(lib.Red, nil)

	srv := mocklib.NewService(t)
	srv.EXPECT().MakeFromFile().Return(fileSrv)

	color, err := trash.DoSmt(srv, content)
	assert.NoError(t, err)

	assert.Equal(t, lib.Red, color)
	// fileSrv.AssertExpectations(t)
}