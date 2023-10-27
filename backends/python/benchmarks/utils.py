import timeit
from django.http import HttpResponse

def measurePerformance(func):
    def wrapper(request, *args, **kwargs):
        stmt = lambda: func(request, *args, **kwargs)
        exec_time = int(timeit.timeit(stmt, number=5)*1000)
        return HttpResponse(exec_time)
    return wrapper