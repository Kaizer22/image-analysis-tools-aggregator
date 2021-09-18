package rest

import (
	"encoding/json"
	"image-analysis-tools-aggregator/logging"
	"image-analysis-tools-aggregator/model/request"
	"image-analysis-tools-aggregator/model/response"
	"image-analysis-tools-aggregator/py_scripts"
	"image-analysis-tools-aggregator/utils"
	"net/http"
	"os/exec"
	"strconv"
	"strings"
)

func GetImageFileInfo(w http.ResponseWriter, r *http.Request) {
	var requestParams request.CommonImageAnalysisRequest
	var resp response.GetFileInfoResponse
	if err := utils.Decoder.Decode(&requestParams, r.URL.Query()); err != nil {
		logging.ErrorFormat("failed to parse parameters %s: %s", r.URL.Query(), err)
		utils.RespondWithJson(w, http.StatusBadRequest, err)
		return
	} else {
		logging.InfoFormat( "filter: %+v", requestParams)
	}
	cmd := exec.Command(py_scripts.ScriptsInterpreter,  py_scripts.GetImageInfoScriptPath,
		py_scripts.ImageCacheFolderPath+ requestParams.InternalFileName)
	stdOutStdErr, err := cmd.CombinedOutput()
	logging.Info(stdOutStdErr)
	output := strings.Split(strings.TrimSpace(string(stdOutStdErr)), py_scripts.ScriptResultsSeparator)
	resp = response.GetFileInfoResponse{
		Width:       output[1],
		Height:      output[2],
		Filetype:    output[0],
		ColorSchema: output[3],
		Size:        output[4],
	}
	if err != nil {
		logging.Error(err)
	}
	utils.RespondWithJson(w, http.StatusOK, resp)
}

func GetFourierTransform(w http.ResponseWriter, r *http.Request) {
	var requestParams request.CommonImageAnalysisRequest
	if err := utils.Decoder.Decode(&requestParams, r.URL.Query()); err != nil {
		logging.ErrorFormat("failed to parse parameters %s: %s", r.URL.Query(), err)
		utils.RespondWithJson(w, http.StatusBadRequest, err)
		return
	} else {
		logging.InfoFormat( "filter: %+v", requestParams)
	}
	cmd := exec.Command(py_scripts.ScriptsInterpreter,  py_scripts.GetFourierTransform,
		py_scripts.ImageCacheFolderPath + requestParams.InternalFileName)
	stdOutStdErr, err := cmd.CombinedOutput()
	logging.Info(string(stdOutStdErr))
	if err != nil {
		logging.InfoFormat("error during python script execution: %s", err)
	}
	utils.RespondWithOutputImage(w, http.StatusOK, py_scripts.OutputImageFolderPath + string(stdOutStdErr))
}

func GetImagePalette(w http.ResponseWriter, r *http.Request) {
	var requestParams request.GetImagePaletteRequest
	if err := utils.Decoder.Decode(&requestParams, r.URL.Query()); err != nil {
		logging.ErrorFormat("failed to parse parameters %s: %s", r.URL.Query(), err)
		utils.RespondWithJson(w, http.StatusBadRequest, err)
		return
	} else {
		logging.InfoFormat( "filter: %+v", requestParams)
	}
	colors := strconv.Itoa(requestParams.ColorsCount)
	cmd := exec.Command(py_scripts.ScriptsInterpreter,  py_scripts.GetImagePalette,
		py_scripts.ImageCacheFolderPath + requestParams.InternalFileName, colors)
	stdOutStdErr, err := cmd.CombinedOutput()
	logging.Info(string(stdOutStdErr))
	if err != nil {
		logging.InfoFormat("error during python script execution: %s", err)
	}
	var resp response.GetImagePaletteResponse
	err = json.Unmarshal(stdOutStdErr, &resp)
	if err != nil {
		utils.RespondWithJson(w, http.StatusInternalServerError, response.AbstractResponse{
			Payload: err.Error(),
		})
	}
	utils.RespondWithJson(w, http.StatusOK, resp)
}

func GetBitPlanes(w http.ResponseWriter, r *http.Request){
	var requestParams request.CommonImageAnalysisRequest
	if err := utils.Decoder.Decode(&requestParams, r.URL.Query()); err != nil {
		logging.ErrorFormat("failed to parse parameters %s: %s", r.URL.Query(), err)
		utils.RespondWithJson(w, http.StatusBadRequest, err)
		return
	} else {
		logging.InfoFormat( "filter: %+v", requestParams)
	}
	cmd := exec.Command(py_scripts.ScriptsInterpreter,  py_scripts.GetImageBitPlanes,
		py_scripts.ImageCacheFolderPath + requestParams.InternalFileName)
	stdOutStdErr, err := cmd.CombinedOutput()
	logging.Info(string(stdOutStdErr))
	if err != nil {
		logging.InfoFormat("error during python script execution: %s", err)
	}
	var resp response.GetBitPlanesResponse
	err = json.Unmarshal(stdOutStdErr, &resp)
	if err != nil {
		utils.RespondWithJson(w, http.StatusInternalServerError, response.AbstractResponse{
			Payload: err.Error(),
		})
	}
	utils.RespondWithMultipart(w, http.StatusOK, resp.BitPlanesFiles)
}

func GetRGBHistogram(w http.ResponseWriter, r *http.Request) {
	var requestParams request.CommonImageAnalysisRequest
	if err := utils.Decoder.Decode(&requestParams, r.URL.Query()); err != nil {
		logging.ErrorFormat("failed to parse parameters %s: %s", r.URL.Query(), err)
		utils.RespondWithJson(w, http.StatusBadRequest, err)
		return
	} else {
		logging.InfoFormat( "filter: %+v", requestParams)
	}
	cmd := exec.Command(py_scripts.ScriptsInterpreter,  py_scripts.GetRGBHistogram,
		py_scripts.ImageCacheFolderPath + requestParams.InternalFileName)
	stdOutStdErr, err := cmd.CombinedOutput()
	logging.Info(string(stdOutStdErr))
	if err != nil {
		logging.InfoFormat("error during python script execution: %s", err)
	}
	var resp response.GetRGBHistogramResponse
	err = json.Unmarshal(stdOutStdErr, &resp)
	if err != nil {
		utils.RespondWithJson(w, http.StatusInternalServerError, response.AbstractResponse{
			Payload: err.Error(),
		})
	}
	utils.RespondWithJson(w, http.StatusOK, resp)
}

