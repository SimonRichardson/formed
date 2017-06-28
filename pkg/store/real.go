package store

import (
	"encoding/csv"

	"github.com/SimonRichardson/formed/pkg/fs"
	"github.com/SimonRichardson/formed/pkg/models"
	"github.com/pkg/errors"
)

type realStore struct {
	fsys fs.Filesystem
	path string
}

// New creates a default store with the correct dependencies.
func New(fsys fs.Filesystem, path string) Store {
	return &realStore{
		fsys: fsys,
		path: path,
	}
}

// Read reads all the user models from the storage, or it returns an error
// if there issue.
func (r *realStore) Read() ([]models.User, error) {
	if !r.fsys.Exists(r.path) {
		return nil, errors.Errorf("no file found at %q", r.path)
	}

	file, err := r.fsys.Open(r.path)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to open file at %q", r.path)
	}
	defer file.Close()

	// Create a reader to read all the lines of the file.
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, errors.Wrapf(err, "unable to read file at %q", r.path)
	}

	// Parse the records to a user model
	users := make([]models.User, len(records))
	for k, v := range records {
		user := &models.User{}
		if err := user.Unmarshal(v); err != nil {
			return nil, errors.Wrapf(err, "unable to parse user at %q for index %d", r.path, k)
		}

		users[k] = *user
	}

	return users, nil
}

// Write, writes users to the underlying storage.
func (r *realStore) Write(users []models.User) error {
	var (
		file fs.File
		err  error
	)
	if !r.fsys.Exists(r.path) {
		file, err = r.fsys.Create(r.path)
		if err != nil {
			return errors.Wrapf(err, "unable to create file at %q", r.path)
		}
	} else {
		file, err = r.fsys.Open(r.path)
		if err != nil {
			return errors.Wrapf(err, "unable to open file at %q", r.path)
		}
	}
	defer file.Close()

	// Marshal all the users to records
	records := make([][]string, len(users))
	for k, v := range users {
		fields, err := v.Marshal()
		if err != nil {
			return errors.Wrapf(err, "unable to marshal user at index %d", k)
		}
		records[k] = fields
	}

	// Write the csv to the file
	writer := csv.NewWriter(file)
	if err := writer.WriteAll(records); err != nil {
		return errors.Wrapf(err, "unable to write file at %q", r.path)
	}

	return nil
}
