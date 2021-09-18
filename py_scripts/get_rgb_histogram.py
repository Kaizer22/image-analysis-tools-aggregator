import sys

import cv2
import jsons
import numpy as np

filename = sys.argv[1]
image = cv2.imread(filename)

red = np.histogram(image[:, :, 0], bins=np.arange(256), density=True)
green = np.histogram(image[:, :, 1], bins=np.arange(256), density=True)
blue = np.histogram(image[:, :, 2], bins=np.arange(256), density=True)

class HistogramResponse:
    def __init__(self):
        self.red_hist = list()
        self.green_hist = list()
        self.blue_hist = list()

resp = HistogramResponse()
resp.red_hist = red[0]
resp.green_hist = green[0]
resp.blue_hist = blue[0]
print(jsons.dumps(resp))