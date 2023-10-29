# This is an auto-generated Django model module.
# You'll have to do the following manually to clean this up:
#   * Rearrange models' order
#   * Make sure each model has one field with primary_key=True
#   * Make sure each ForeignKey and OneToOneField has `on_delete` set to the desired behavior
#   * Remove `managed = False` lines if you wish to allow Django to create, modify, and delete the table
# Feel free to rename the models, but don't rename db_table values or field names.
from django.db import models


class ElectricCars(models.Model):
    vin = models.TextField(blank=True, null=True)
    county = models.TextField(blank=True, null=True)
    city = models.TextField(blank=True, null=True)
    state = models.TextField(blank=True, null=True)
    postalcode = models.IntegerField(db_column='postalCode', blank=True, null=True)  # Field name made lowercase.
    modelyear = models.IntegerField(db_column='modelYear', blank=True, null=True)  # Field name made lowercase.
    make = models.TextField(blank=True, null=True)
    model = models.TextField(blank=True, null=True)
    vehicletype = models.TextField(db_column='vehicleType', blank=True, null=True)  # Field name made lowercase.
    electricrange = models.IntegerField(db_column='electricRange', blank=True, null=True)  # Field name made lowercase.
    vehiclelocation = models.TextField(db_column='vehicleLocation', blank=True, null=False, primary_key=True)  # Field name made lowercase.

    class Meta:
        managed = False
        db_table = 'electric_cars'