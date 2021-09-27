package rest

import (
	"image-analysis-tools-aggregator/logging"
	"image-analysis-tools-aggregator/model/response"
	"image-analysis-tools-aggregator/py_scripts"
	"image-analysis-tools-aggregator/utils"
	"io"
	"net/http"
	"os"
)

func UploadImage(w http.ResponseWriter, r *http.Request) {
	logging.InfoFormat("got upload image request")
	//parse the multipart form in the request
	err := r.ParseMultipartForm(6000000)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//get a ref to the parsed multipart form
	m := r.MultipartForm
	//get the *fileheaders
	files := m.File["file"]
	for i := range files {
		//for each fileheader, get a handle to the actual file
		file, err := files[i].Open()
		defer file.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		internalFilename, err := utils.GetInternalFilename(files[i].Filename)
		if err != nil {
			utils.RespondWithJson(w, http.StatusBadRequest,
				response.AbstractResponse{Payload: err.Error()})
			return
		}
		//create destination file making sure the path is writeable.
		dst, err := os.Create(py_scripts.ImageCacheFolderPath + internalFilename)
		defer dst.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		//copy the uploaded file to the destination file
		if _, err := io.Copy(dst, file); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		utils.RespondWithJson(w, http.StatusOK, response.UploadImageResponse{InternalFileName: internalFilename})
	}
	//display success message.
	logging.InfoFormat("file uploaded successfully")
}

//func ShutdownCacheManagerInstance(w http.ResponseWriter, r *http.Request) {
//	for _, channel := range cache.InstanceControlChannels {
//		channel <- true
//		break
//	}
//}
//
//func AddCacheManagerInstance(w http.ResponseWriter, r *http.Request) {
//
//}
