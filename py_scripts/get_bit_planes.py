import sys

import cv2
import jsons
import numpy as np

from py_utils import save_image_with_internal_name, OUTPUT_IMAGE_FOLDER

filename = sys.argv[1]
img = cv2.imread(filename, cv2.IMREAD_GRAYSCALE)

out_files = list()
for k in range(0, 8):
    # create an image for each k bit plane
    plane = np.full((img.shape[0], img.shape[1]), 2 ** k, np.uint8)
    # execute bitwise AND operation
    res = cv2.bitwise_and(plane, img)
    # to make pixels to be only zeros and 255 (+ to avoid zero division warning)
    x = res * (255 / (res + 0.0001))
    plane_internal_name = save_image_with_internal_name(x)
    out_files.append(plane_internal_name)


class BitPlanesResponse:
    def __init__(self):
        self.bit_planes_files = list()


rsp = BitPlanesResponse()
for i in range(0, len(out_files)):
    rsp.bit_planes_files.append(OUTPUT_IMAGE_FOLDER + out_files[i])
print(jsons.dumps(rsp), end="")
