import sys
import cv2

import numpy as np
from numpy import fft

from py_utils import save_image_with_internal_name

filename = sys.argv[1]

img = cv2.imread(filename, cv2.IMREAD_GRAYSCALE)
# 1. get a 2d discrete Fourier Transform of an image
# 2. shift zero-frequency component to the center of a result
# 3. get absolute values of complex numbers
# 4. reduce dynamic range using gamma correction
#    (instead of logarithmic transform to reduce noise and to
#    identify the predominant frequency components)
#                   s = c * r ^ γ
gamma = 0.09
img_fourier = abs(fft.fftshift(fft.fft2(img))) ** gamma
# displaying the result on an 8-bit grayscale representation
min_val = np.min(img_fourier)
max_val = np.max(img_fourier)
f_range = max_val - min_val
normalized_fourier_image = (img_fourier - min_val) / f_range * 255

print(save_image_with_internal_name(normalized_fourier_image), end="")
