package controllers

import (
	"net/http"
	"project-akhir-gdsc-backend/services"
	"project-akhir-gdsc-backend/utils"
)

func ConvertImagesToPDF(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		utils.SendJSONResponse(w, http.StatusBadRequest, map[string]string{"status": "error", "message": "tidak ada file yang diunggah atau total file yang diunggah melebihi 10 MB"})
		return
	}

	files := r.MultipartForm.File["images"]

	if err := utils.ValidateFiles(files); err != nil {
		utils.SendJSONResponse(w, http.StatusBadRequest, map[string]string{"status": "error", "message": err.Error()})
		return
	}

	pdfFilePath, err := services.CreatePDF(files, r)
	if err != nil {
		utils.SendJSONResponse(w, http.StatusInternalServerError, map[string]string{"status": "error", "message": err.Error()})
		return
	}

	utils.SendJSONResponse(w, http.StatusOK, map[string]string{"status": "success", "message": "gambar berhasil dikonversi ke PDF", "file_path": pdfFilePath})
}
