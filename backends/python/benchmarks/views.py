from . import utils

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