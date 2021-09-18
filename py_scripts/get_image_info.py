import os
import sys

from PIL import Image

filename = sys.argv[1]

sample = Image.open(filename)
print(sample.format, ";"
      ,sample.size[0], ";"
      ,sample.size[1], ";"
      ,sample.mode, ";",
      os.path.getsize(filename), "Bytes", end="")
