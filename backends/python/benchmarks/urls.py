"""
URL configuration for benchmarks project.

The `urlpatterns` list routes URLs to views. For more information please see:
    https://docs.djangoproject.com/en/4.2/topics/http/urls/
Examples:
Function views
    1. Add an import:  from my_app import views
    2. Add a URL to urlpatterns:  path('', views.home, name='home')
Class-based views
    1. Add an import:  from other_app.views import Home
    2. Add a URL to urlpatterns:  path('', Home.as_view(), name='home')
Including another URLconf
    1. Import the include() function: from django.urls import include, path
    2. Add a URL to urlpatterns:  path('blog/', include('blog.urls'))
"""
from django.contrib import admin
from django.urls import path
from . import views
import multiprocessing
from PIL import Image

def initiateQuicksortArray(N):
    return [N-i for i in range(N-1,-1,-1)]

def initiateTwoSumArray(N):
    return [i+1 for i in range(0,N,1)]

def getCores():
    return multiprocessing.cpu_count()

def generateJsonObject():
    sampleString = "Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet."
    return {"a": sampleString, "b": sampleString , "c": sampleString, "d": sampleString, "e": sampleString, "f": sampleString , "g": sampleString, "h": sampleString,"i": sampleString, "j": sampleString , "k": sampleString, "l": sampleString,"m": sampleString, "n": sampleString , "o": sampleString, "p": sampleString,"q": sampleString}

def binaryImage():
    return Image.open('../../public/sample.jpg')

urlpatterns = [
    path('admin/', admin.site.urls),
    path('algorithms/prime/', views.calculatePrimeNumbers, {"N": 10000000}),
    path('algorithms/fast-fibonacci/', views.fastFibonacci, {"N": 1000000}),
    path('algorithms/quicksort/', views.quickSort, {"N": 1000000, "Array": initiateQuicksortArray(1000000)}),
    path('algorithms/two-sum/', views.twoSum, {"N": 1000000, "Array": initiateTwoSumArray(1000000)}),
    path('files/read/', views.createAndReadFile, {"N": 1000000, "Cores": getCores()}),
    path('files/read-parallel/', views.createAndReadFileParallel, {"N": 1000000, "Cores": getCores()}),
    path('files/write/', views.writeFile, {"N": 10000, "Cores": getCores()}),
    path('files/write-parallel/', views.writeFileParallel, {"N": 10000, "Cores": getCores()}),
    path('db/select/', views.dbSelect, {"N": 10}),
    path('db/insert/', views.dbInsert, {"N": 10}),
    path('db/update/', views.dbUpdate, {"N": 10}),
    path('db/delete/', views.dbDelete, {"N": 50}),
    path('json/serialize/', views.serialize, {"N": 10000, "Object": generateJsonObject()}),
    path('json/deserialize/', views.deserialize, {"N": 10000}),
    path('image/decode/', views.decode, {"N": 10}),
    path('image/encode/', views.encode, {"Image": binaryImage()}),
]
