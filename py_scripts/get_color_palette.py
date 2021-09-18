import json
import sys
from collections import namedtuple
import random

import jsons as jsons
from PIL import Image

# Determination of dominant colors by the k-means method
# source https://habr.com/ru/post/156045/

# point - 3D point in RGB space (coords)
# with associated number of uses in the original image (count)
# and number of dimensions (n=3)
Point = namedtuple('Point', ('coords', 'n', 'count'))
# group of points with a center in a 3D RGB space
Cluster = namedtuple('Cluster', ('points', 'center', 'n'))


def get_points(img):
    points = []
    w, h = img.size
    # get all colors (points of 3D RGB space) from the image with numbers of uses
    for count, color in img.getcolors(w * h):
        points.append(Point(color, 3, count))
    return points


# transform rgb tuple to HEX representation
# rtoh = lambda rgb: '#%s' % ''.join(('%02x' % p for p in rgb))


# main method of an algorithm
def colorz(filename, n=10):
    img = Image.open(filename)
    # image reduction
    img.thumbnail((150, 150))

    points = get_points(img)
    clusters = kmeans(points, n, 1)

    # transform cluster centers into integers
    rgbs = [[int(cc) for cc in c.center.coords] for c in clusters]
    return rgbs


# euclidean distance between two points in n-dimensional (RGB) space
# (sqrt operation deleted in optimisation purposes)
def euclidean(p1, p2):
    return sum([
        (p1.coords[i] - p2.coords[i]) ** 2 for i in range(p1.n)
    ])


def calculate_center(points, n):
    vals = [0.0 for i in range(n)]
    plen = 0
    # iterate over points of cluster
    for p in points:
        # count the number of points in cluster
        plen += p.count
        for i in range(n):
            vals[i] += (p.coords[i] * p.count)
    # return new point with mean values of cluster's points coordinates as coordinates
    return Point([(v / plen) for v in vals], n, 1)


def kmeans(points, k, min_diff):
    # get initial k clusters from color points in RGB-space with random centers
    clusters = [Cluster([p], p, p.n) for p in random.sample(points, k)]

    while 1:
        # init lists of points in clusters
        plists = [[] for _ in range(k)]

        for p in points:
            smallest_distance = float('Inf')
            # for every color point calculate affiliation to the cluster based on euclidean distance
            # smaller distance -> closer to the center of the cluster
            for i in range(k):
                distance = euclidean(p, clusters[i].center)
                if distance < smallest_distance:
                    smallest_distance = distance
                    idx = i
            plists[idx].append(p)

        diff = 0
        # iterate over clusters and recalculate cluster center
        # while the distance between new and old cluster center
        # will be less then min_diff
        for i in range(k):
            old = clusters[i]
            center = calculate_center(plists[i], old.n)
            new = Cluster(plists[i], center, old.n)
            clusters[i] = new
            diff = euclidean(old.center, new.center)

        if diff < min_diff:
            break

    return clusters


class PaletteResponse:
    def __init__(self):
        self.rgb_colors = list()


filename = sys.argv[1]
colNum = sys.argv[2]

rsp = PaletteResponse()
rsp.rgb_colors = colorz(filename, int(colNum))
print(jsons.dumps(rsp), end="")
