import uuid

import cv2

OUTPUT_IMAGE_FOLDER = "./output/"
OUTPUT_FILE_TYPE = ".jpg"


def save_image_with_internal_name(img):
    internal_name = str(uuid.uuid1()) + OUTPUT_FILE_TYPE
    cv2.imwrite(OUTPUT_IMAGE_FOLDER + internal_name, img)
    return internal_name
