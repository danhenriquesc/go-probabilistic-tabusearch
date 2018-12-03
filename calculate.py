import sys
from math import sqrt, pow

# problem = sys.argv[1]
# size = int(sys.argv[2])

# print 'Problem: {}'.format(problem)
# print 'Size: {}'.format(size)

# coords = {}


# def calculate_distances(a, b):
#     return sqrt(pow(a[0] - b[0], 2) + pow(a[1] - b[1], 2))


# with open('instances/{}.tsp.txt'.format(problem)) as file:
#     lines = file.readlines()

#     for line in lines:
#         i, x, y = line.split()

#         coords[i] = (float(x), float(y))

# print(coords)


# to_evaluate = "1 28 6 12 9 5 26 29 3 2 20 10 4 15 18 17 14 22 11 19 25 7 23 27 8 24 16 13 21".split()

# i = 0
# fitness = 0
# while i < size - 1:
#     fitness += calculate_distances(coords[to_evaluate[i]], coords[to_evaluate[i + 1]])
#     i += 1

# fitness += calculate_distances(coords[to_evaluate[size - 1]], coords[to_evaluate[0]])

# print(fitness)

distances = []
with open('bays.txt') as file:
    lines = file.readlines()

    for line in lines:
    	distances.append(line.split())


to_evaluate = list(map(int, "19 11 22 17 14 18 15 4 10 20 2 21 5 29 3 26 9 12 6 28 1 24 13 16 27 8 23 7 25".split()))
size = len(to_evaluate)
i = 0
fitness = 0

while i < size - 1:
	try:
		fitness += int(distances[to_evaluate[i] - 1][to_evaluate[i + 1] - 1])
		i += 1
	except e:
		import ipdb; ipdb.set_trace()

fitness += int(distances[to_evaluate[size - 1] - 1][to_evaluate[0] - 1])

print(fitness)