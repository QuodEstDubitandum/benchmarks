from . import utils

@utils.measurePerformance
def calculatePrimeNumbers(request):
    for i in range(0,100):
        i+=1


@utils.measurePerformance
def fastFibonacci(request):
    for i in range(0,100):
        i+=1


@utils.measurePerformance
def quickSort(request):
    for i in range(0,100):
        i+=1


@utils.measurePerformance
def twoSum(request):
    for i in range(0,100):
        i+=1
