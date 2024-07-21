package utils

import (
	"fmt"
	"mime/multipart"
	"path/filepath"
	"strings"
)

func ValidateFiles(files []*multipart.FileHeader) error {
	for _, file := range files {
		if file.Size > 2*1024*1024 {
			return fmt.Errorf("ukuran file %s melebihi 2 MB", file.Filename)
		}

		validFormats := []string{".jpg", ".jpeg", ".png"}
		ext := strings.ToLower(filepath.Ext(file.Filename))
		isValidFormat := false
		for _, format := range validFormats {
			if ext == format {
				isValidFormat = true
				break
			}
		}
		if !isValidFormat {
			return fmt.Errorf("format file %s tidak didukung", file.Filename)
		}
	}
	return nil
}
