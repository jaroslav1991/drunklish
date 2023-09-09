package handlers

import (
	"drunklish/internal/pkg/httputils"
	"fmt"
	"go.uber.org/zap"
	"io"
	"log"
	"net/http"
	"os"
)

func CreateUploadHandler(logger *zap.Logger, service WordService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "unexpect method", http.StatusMethodNotAllowed)
			return
		}

		token := r.Header.Get("Token")

		if err := r.ParseMultipartForm(10 << 20); err != nil {
			log.Println(err)
			httputils.WriteErrorResponse(logger, w, r, fmt.Errorf("%v", err))
			return
		}
		file, header, err := r.FormFile("fileToRead")
		if err != nil {
			log.Println(err)
			httputils.WriteErrorResponse(logger, w, r, fmt.Errorf("%v", err))
			return
		}

		defer file.Close()

		dst, err := os.Create(header.Filename)
		if err != nil {
			log.Println(err)
			httputils.WriteErrorResponse(logger, w, r, fmt.Errorf("%v", err))
			return
		}

		if _, err = io.Copy(dst, file); err != nil {
			log.Println(err)
			httputils.WriteErrorResponse(logger, w, r, fmt.Errorf("%v", err))
			return
		}

		records, err := httputils.TryCSV(header.Filename)
		if err != nil {
			log.Println(err)
			httputils.WriteErrorResponse(logger, w, r, fmt.Errorf("%v", err))
			return
		}

		words, err := httputils.Parse(records)
		if err != nil {
			log.Println(err)
			httputils.WriteErrorResponse(logger, w, r, fmt.Errorf("%v", err))
			return
		}

		fromUpload, err := service.CreateListFromUpload(words, token)
		if err != nil {
			log.Println(err)
			httputils.WriteErrorResponse(logger, w, r, fmt.Errorf("%v", err))
			return
		}

		httputils.WriteSuccessResponse(logger, w, r, http.StatusOK, fromUpload)
	}
}
