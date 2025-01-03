package validators

import (
	"errors"
	"mime/multipart"
)

type CreateApplication struct {
	Logo *multipart.FileHeader `json:"-" form:"logo"`
	Name string                `validate:"required|maxLen:100" form:"name" json:"name" label:"Application Name"`
}

func (self CreateApplication) ValidateFile() error {
	if self.Logo == nil {
		return nil
	}

	const maxFileSize = 2 * 1024 * 1024
	if self.Logo.Size > maxFileSize {
		return errors.New("file size must not exceed 2 MB")
	}

	validTypes := []string{"image/png", "image/jpeg"}
	for _, t := range validTypes {
		if self.Logo.Header.Get("Content-Type") == t {
			return nil
		}
	}

	return errors.New("file must be a PNG or JPEG image")
}
