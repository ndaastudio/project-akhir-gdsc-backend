package services

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/jung-kurt/gofpdf"
)

func CreatePDF(files []*multipart.FileHeader, r *http.Request) (string, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")

	for _, file := range files {
		filePath := filepath.Join("results", file.Filename)
		src, err := file.Open()
		if err != nil {
			return "", fmt.Errorf("gagal membuka file: %v", err)
		}
		defer src.Close()

		dst, err := createFile(filePath)
		if err != nil {
			return "", fmt.Errorf("gagal membuat file: %v", err)
		}
		defer dst.Close()

		if _, err := io.Copy(dst, src); err != nil {
			return "", fmt.Errorf("gagal menyimpan file: %v", err)
		}

		pdf.AddPage()
		pdf.Image(filePath, 10, 10, 190, 0, false, "", 0, "")
	}

	pdfFilePath := filepath.Join("results", "converted.pdf")
	err := pdf.OutputFileAndClose(pdfFilePath)
	if err != nil {
		return "", fmt.Errorf("gagal membuat PDF: %v", err)
	}

	return pdfFilePath, nil
}

func createFile(filePath string) (*os.File, error) {
	dst, err := os.Create(filePath)
	if err != nil {
		return nil, err
	}
	return dst, nil
}
