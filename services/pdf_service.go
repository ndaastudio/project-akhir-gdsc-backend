package services

import (
	"fmt"
	"image"
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

		imgFile, err := os.Open(filePath)
		if err != nil {
			return "", fmt.Errorf("gagal membuka file gambar: %v", err)
		}
		defer imgFile.Close()

		img, _, err := image.DecodeConfig(imgFile)
		if err != nil {
			return "", fmt.Errorf("gagal mendekode gambar: %v", err)
		}

		const pageWidth, pageHeight = 210.0, 297.0
		imgHeight := float64(img.Height) * 25.4 / 72.0

		numPages := int(imgHeight / pageHeight)
		if imgHeight > float64(numPages)*pageHeight {
			numPages++
		}

		for i := 0; i < numPages; i++ {
			pdf.AddPage()
			yOffset := float64(i) * pageHeight
			pdf.ImageOptions(filePath, 0, -yOffset, pageWidth, imgHeight, false, gofpdf.ImageOptions{ReadDpi: true}, 0, "")
		}
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
