package py_scripts

var (
	// *** DIRECTORIES ***
	ImageCacheFolderPath  = "./tmp/"
	OutputImageFolderPath = "./output/"

	// *** SCRIPTS IO ***
	ScriptResultsSeparator = ";"
	ScriptsInterpreter = "python"

	// *** SCRIPTS ***
	GetImageInfoScriptPath = "py_scripts/get_image_info.py"
	GetFourierTransform = "py_scripts/get_fourier_transform.py"
	GetImagePalette = "py_scripts/get_color_palette.py"
	GetEdgeDetectionResult = "py_scripts/get_edge_detection_result.py"
	GetImageSegmentationResults = "py_scripts/get_image_segmentation_result.py"
	GetRGBHistogram = "py_scripts/get_rgb_histogram.py"
	GetImageBitPlanes = "py_scripts/get_bit_planes.py"
)
