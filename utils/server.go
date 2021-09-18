package utils

import (
	"encoding/json"
	"errors"
	"github.com/golang/snappy"
	"github.com/google/uuid"
	"github.com/gorilla/schema"
	"image-analysis-tools-aggregator/logging"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var (
	SupportedFiletypes = []string{"jpg", "jpeg", "png", "gif", "bmp", "tif"}
)

// Decoder: request parameters decoder
var Decoder = schema.NewDecoder()

func RespondWithOutputImage(w http.ResponseWriter, status int, filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Compressed", "false")
	w.Write(fileBytes)

	err = file.Close()
	if err != nil {
		logging.InfoFormat("cannot close file %s: %s", filename, err)
	}
	err = os.Remove(filename)
	if err != nil {
		logging.InfoFormat("cannot delete file %s: %s", filename, err)
	}
}
func RespondWithMultipart(w http.ResponseWriter, status int, filenames []string ) {
	mw := multipart.NewWriter(w)
	w.WriteHeader(status)
	w.Header().Set("Content-Type", mw.FormDataContentType())
	rawFiles := make([][]byte, 0,len(filenames))
	for _, filename := range filenames {
		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			panic(err)
		}
		rawFiles = append(rawFiles, fileBytes)
		file.Close()
	}
	for i, value := range rawFiles {
		fw, err := mw.CreateFormFile("file" + strconv.Itoa(i),filenames[i])
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if _, err := fw.Write(value); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = os.Remove(filenames[i])
		if err != nil {
			logging.InfoFormat("cannot delete file %s: %s", filenames[i], err)
		}
	}
	if err := mw.Close(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func RespondWithJson(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		logging.ErrorFormat("Unexpected error while marshalling:", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Compressed", "false")
	w.WriteHeader(status)
	_, _ = w.Write(response)
}

func RespondWithCompressedJson(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		logging.ErrorFormat("Unexpected error while marshalling:", err)
	}
	response = snappy.Encode(response, response)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Compressed", "true")
	w.WriteHeader(status)
	_, _ = w.Write(response)
}

func GetInternalFilename(filename string) (internalFilename string, err error) {
	buf := strings.Split(filename, ".")
	format := buf[len(buf) - 1]
	if _, contains := StringInArray(format, SupportedFiletypes); !contains {
		return "", errors.New("file cannot be processed, the image was expected")
	}
	uid, err := uuid.NewUUID()
	if err != nil {
		logging.ErrorFormat("cannot generate internal filename!")
		internalFilename = filename
	} else {
		internalFilename = uid.String()
	}
	return internalFilename + "." + format, nil
}
