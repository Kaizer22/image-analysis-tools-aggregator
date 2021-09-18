package main

import (
	"github.com/gorilla/mux"
	"image-analysis-tools-aggregator/cache"
	"image-analysis-tools-aggregator/logging"
	"image-analysis-tools-aggregator/rest"
	"image-analysis-tools-aggregator/utils"
	"net/http"
)
var (
	addr string
)

func main(){
	initService()

	mux := mux.NewRouter()
	mux.HandleFunc("/upload-image", rest.UploadImage).
		Methods("POST")

	mux.HandleFunc("/get-image-info", rest.GetImageFileInfo).
		Methods("GET")
	mux.HandleFunc("/get-fft", rest.GetFourierTransform).
		Methods("GET")
	mux.HandleFunc("/get-image-palette", rest.GetImagePalette).
		Methods("GET")
	mux.HandleFunc("/get-image-bit-planes", rest.GetBitPlanes).
		Methods("GET")
	mux.HandleFunc("/get-edge-detection-result", rest.GetImageFileInfo).
		Methods("GET")
	mux.HandleFunc("/get-image-segmentation-result", rest.GetFourierTransform).
		Methods("GET")
	mux.HandleFunc("/get-rgb-histogram", rest.GetRGBHistogram).
		Methods("GET")

	logging.InfoFormat("starting server at %s...", addr)
	cache.InitCache(cache.GetCacheConfig())
	http.ListenAndServe(addr, mux)
}

func initService() {
	addr = utils.GetEnv(utils.ListenAddressEnvKey, ":8080")
}


