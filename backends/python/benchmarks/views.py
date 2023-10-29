from django.http import HttpResponse
from . import utils
import os
import threading
from .models import ElectricCars
import json
import timeit
from PIL import Image
import numpy as np

# /algorithms/prime
@utils.measurePerformance
def calculatePrimeNumbers(request, *args, **kwargs):
    p = 2
    n = kwargs.get('N')
    primes = [False] * (n + 1)

    # exclude 1
    j = n - 1

    # sieve of eratosthenes
    while p * p <= n:
        if not primes[p - 1]:
            for i in range(p * p, n + 1, p):
                if not primes[i - 1]:
                    j -= 1
                primes[i - 1] = True
        p += 1

    print(f"Number of prime numbers until {n}: {j}")

# /algorithms/fast-fibonacci
@utils.measurePerformance
def fastFibonacci(request, *args, **kwargs):
    n = kwargs.get('N')
    base_matrix = [
        [1, 1],
        [1, 0]
    ]
    result_matrix = powerMatrix(base_matrix, n)

def powerMatrix(matrix, n):
    if n == 1:
        return matrix
    if n % 2 == 1:
        return matrixMult(matrix, powerMatrix(matrix, n - 1))
    root = powerMatrix(matrix, n // 2)
    return matrixMult(root, root)

def matrixMult(a, b):
    res = [[0,0],[0,0]]
    res[0][0] = a[0][0] * b[0][0] + a[0][1] * b[1][0]
    res[0][1] = a[0][0] * b[0][1] + a[0][1] * b[1][1]
    res[1][0] = a[1][0] * b[0][0] + a[1][1] * b[1][0]
    res[1][1] = a[1][0] * b[0][1] + a[1][1] * b[1][1]

    return res

# /algorithms/quicksort
@utils.measurePerformance
def quickSort(request, *args, **kwargs):
    base_array = kwargs.get('Array')
    print(divide_array(base_array)[0])

def divide_array(array):
    length = len(array)
    if length < 2:
        return array

    pivot = length // 2
    left = 0

    array[length - 1], array[pivot] = array[pivot], array[length - 1]

    for i in range(length):
        if array[i] < array[length - 1]:
            array[i], array[left] = array[left], array[i]
            left += 1

    array[left], array[length - 1] = array[length - 1], array[left]

    divide_array(array[:left])
    divide_array(array[left+1:])

    return array

# /algorithms/two-sum
@utils.measurePerformance
def twoSum(request, *args, **kwargs):
    n = kwargs.get('N')
    hashmap = {}
    arr = kwargs.get('Array')
    count = 0

    for number in arr:
        complement = n - number
        count += hashmap.get(complement, 0)
        hashmap[number] = hashmap.get(number, 0) + 1

    print(count)

# /files/read
def createAndReadFile(request, *args, **kwargs):
    # create the files first
    content = 'Gotta go fast\n' * kwargs.get("N")
    for i in range(kwargs.get("Cores")):
        with open('test{}.txt'.format(i), 'w') as file:
            file.write(content)
    time = readFile(request, Cores=kwargs.get("Cores"))

    for i in range(kwargs.get("Cores")):
        os.remove('test{}.txt'.format(i))
    
    return time

@utils.measurePerformance
def readFile(request, *args, **kwargs):
    for i in range(kwargs.get("Cores")):
        with open('test{}.txt'.format(i), 'r') as file:
            file.read()


# /files/read-parallel
def createAndReadFileParallel(request, *args, **kwargs):
    # create the files first
    content = 'Gotta go fast\n' * kwargs.get("N")
    for i in range(kwargs.get("Cores")):
        with open('test{}.txt'.format(i), 'w') as file:
            file.write(content)
    time = readFileParallel(request, Cores=kwargs.get("Cores"))

    for i in range(kwargs.get("Cores")):
        os.remove('test{}.txt'.format(i))
    
    return time

@utils.measurePerformance
def readFileParallel(request, *args, **kwargs):
    threads = []
    for i in range(kwargs.get("Cores")):
        thread = threading.Thread(target=readFileThread, args=('test{}.txt'.format(i),))
        threads.append(thread)
        thread.start()
    for thread in threads:
        thread.join()

def readFileThread(filename):
    with open(filename, 'r') as file:
        file.read()


# /files/write
@utils.measurePerformance
def writeFile(request, *args, **kwargs):
    for _ in range(kwargs.get("Cores")):
        with open('test.txt', 'w') as file:
            for _ in range(kwargs.get("N")):
                file.write('Gotta go fast\n')
        os.remove('test.txt')

# /files/write-parallel
@utils.measurePerformance
def writeFileParallel(request, *args, **kwargs):
    threads = []
    for i in range(kwargs.get('Cores')):
        thread = threading.Thread(target=writeFileThread, args=('test{}.txt'.format(i), kwargs.get('N')))
        threads.append(thread)
        thread.start()
    for thread in threads:
        thread.join()

def writeFileThread(filename, n):
    with open(filename, 'w') as file:
        for _ in range(n):
            file.write('Gotta go fast\n')
    os.remove(filename)

# /db/select
@utils.measurePerformance
def dbSelect(request, *args, **kwargs):
    count = 0
    for _ in range(kwargs.get('N')):
        count = ElectricCars.objects.filter(vehicletype="BEV", model="MODEL 3").count()

    print(count)

# /db/insert
def dbInsert(request, *args, **kwargs):
    time = insertOnly(request, N=kwargs.get('N'))
    # delete the inserts afterwards
    ElectricCars.objects.filter(vin="test", county="test").delete()
    return time

@utils.measurePerformance
def insertOnly(request, *args, **kwargs):
    for _ in range(kwargs.get('N')):
        ElectricCars.objects.create(
            vin="test",
            county="test",
            city="test",
            state="test",
            postalcode=1,
            modelyear=1,
            make="test",
            model="test",
            vehicletype="test",
            electricrange=1,
            vehiclelocation="test"
        )

# /db/update
@utils.measurePerformance
def dbUpdate(request, *args, **kwargs):
    ElectricCars.objects.filter(model="I3").update(model="I4")
    ElectricCars.objects.filter(model="I4").update(model="I3")

# /db/delete
def dbDelete(request, *args, **kwargs):
    time = 0
    for _ in range(kwargs.get('N')):
        ElectricCars.objects.create(
            vin="test",
            county="test",
            city="test",
            state="test",
            postalcode=1,
            modelyear=1,
            make="test",
            model="test",
            vehicletype="test",
            electricrange=1,
            vehiclelocation="test"
        )
        time += int(timeit.timeit(deleteOnly, number=1)*1000)
    return HttpResponse(time//5)

def deleteOnly():
    ElectricCars.objects.filter(vin="test", county="test").delete()

# /json/serialize
@utils.measurePerformance
def serialize(request, *args, **kwargs):
    for _ in range(kwargs.get('N')):
        json.dumps(kwargs.get("Object"))

# /json/deserialize
@utils.measurePerformance
def deserialize(request, *args, **kwargs):
    for _ in range(kwargs.get('N')):
        json.loads(request.body)

# /image/decode
@utils.measurePerformance
def decode(request, *args, **kwargs):
    for _ in range(kwargs.get('N')):
        np.array(Image.open('../../public/sample.jpg'))

# /image/encode
@utils.measurePerformance
def encode(request, *args, **kwargs):
    image = kwargs.get("Image")
    image.save('../../public/sample.png', "PNG")