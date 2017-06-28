package store

import (
	"errors"
	"testing"

	"io"

	"reflect"

	"github.com/SimonRichardson/formed/pkg/fs/mock_fs"
	"github.com/SimonRichardson/formed/pkg/models"
	"github.com/golang/mock/gomock"
)

func TestRealRead(t *testing.T) {
	t.Parallel()

	t.Run("read empty file", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var (
			mockStore = mock_fs.NewMockFilesystem(ctrl)
			mockFile  = mock_fs.NewMockFile(ctrl)

			path  = "path/to/file"
			store = New(mockStore, path)
		)

		// Create some mocking expectations
		mockStore.EXPECT().
			Exists(path).
			Return(true)

		mockFile.EXPECT().
			Read(gomock.Any()).
			Return(0, io.EOF)

		mockStore.EXPECT().
			Open(path).
			Return(mockFile, nil)

		mockFile.EXPECT().
			Close().
			Return(nil)

		users, err := store.Read()
		if err != nil {
			t.Fatal(err)
		}

		if expected, actual := 0, len(users); expected != actual {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	})

	t.Run("read non-empty file", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var (
			want      = models.User{"fred", "smith"}
			mockStore = mock_fs.NewMockFilesystem(ctrl)
			stubFile  = &stubFile{
				bytes: []byte(want.String()),
			}

			path  = "path/to/file"
			store = New(mockStore, path)
		)

		// Create some mocking expectations
		mockStore.EXPECT().
			Exists(path).
			Return(true)

		mockStore.EXPECT().
			Open(path).
			Return(stubFile, nil)

		users, err := store.Read()
		if err != nil {
			t.Fatal(err)
		}

		if expected, actual := 1, len(users); expected != actual {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
		if expected, actual := want, users[0]; !reflect.DeepEqual(expected, actual) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	})

	t.Run("file does not exist", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var (
			mockStore = mock_fs.NewMockFilesystem(ctrl)

			path  = "path/to/file"
			store = New(mockStore, path)
		)

		// Create some mocking expectations
		mockStore.EXPECT().
			Exists(path).
			Return(false)

		_, err := store.Read()

		if expected, actual := false, err == nil; expected != actual {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	})

	t.Run("file does not open", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var (
			mockStore = mock_fs.NewMockFilesystem(ctrl)

			path  = "path/to/file"
			store = New(mockStore, path)
		)

		// Create some mocking expectations
		mockStore.EXPECT().
			Exists(path).
			Return(true)

		mockStore.EXPECT().
			Open(path).
			Return(nil, errors.New("permissions"))

		_, err := store.Read()

		if expected, actual := false, err == nil; expected != actual {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	})

	t.Run("unable to read file", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var (
			mockStore = mock_fs.NewMockFilesystem(ctrl)
			mockFile  = mock_fs.NewMockFile(ctrl)

			path  = "path/to/file"
			store = New(mockStore, path)
		)

		// Create some mocking expectations
		mockStore.EXPECT().
			Exists(path).
			Return(true)

		mockFile.EXPECT().
			Read(gomock.Any()).
			Return(0, errors.New("parse error"))

		mockStore.EXPECT().
			Open(path).
			Return(mockFile, nil)

		mockFile.EXPECT().
			Close().
			Return(nil)

		_, err := store.Read()

		if expected, actual := false, err == nil; expected != actual {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	})
}

func TestRealWrite(t *testing.T) {
	t.Parallel()

	t.Run("write nothing for no users", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var (
			mockStore = mock_fs.NewMockFilesystem(ctrl)
			mockFile  = mock_fs.NewMockFile(ctrl)

			path  = "path/to/file"
			store = New(mockStore, path)
		)

		// Create some mocking expectations
		mockStore.EXPECT().
			Exists(path).
			Return(true)

		mockStore.EXPECT().
			Open(path).
			Return(mockFile, nil)

		mockFile.EXPECT().
			Close().
			Return(nil)

		err := store.Write([]models.User{})
		if err != nil {
			t.Fatal(err)
		}

		if expected, actual := true, err == nil; expected != actual {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	})

	t.Run("write users", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var (
			user = models.User{"fred", "smith"}
			want = user.String()

			mockStore = mock_fs.NewMockFilesystem(ctrl)
			mockFile  = mock_fs.NewMockFile(ctrl)

			path  = "path/to/file"
			store = New(mockStore, path)
		)

		// Create some mocking expectations
		mockStore.EXPECT().
			Exists(path).
			Return(true)

		mockFile.EXPECT().
			Write(gomock.Any()).
			Return(len(want)*2, nil)

		mockStore.EXPECT().
			Open(path).
			Return(mockFile, nil)

		mockFile.EXPECT().
			Close().
			Return(nil)

		err := store.Write([]models.User{
			user,
		})
		if err != nil {
			t.Fatal(err)
		}

		if expected, actual := true, err == nil; expected != actual {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	})
}

type stubFile struct {
	bytes  []byte
	called int
}

func (f *stubFile) Read(b []byte) (int, error) {
	defer func() { f.called++ }()

	if f.called == 0 {
		x := reflect.ValueOf(b)
		for k, v := range f.bytes {
			x.Index(k).Set(reflect.ValueOf(v))
		}
		return len(f.bytes), nil
	}
	return 0, io.EOF
}

func (f *stubFile) Write(b []byte) (int, error) {
	return 0, io.EOF
}

func (f *stubFile) Close() error {
	return nil
}
