# Generated by Django 2.2.1 on 2019-05-06 19:59

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('wells', '0078_merge_20190506_1922'),
    ]

    operations = [
        migrations.AlterField(
            model_name='activitysubmission',
            name='final_casing_stick_up',
            field=models.DecimalField(blank=True, decimal_places=3, max_digits=6, null=True, verbose_name='Final Casing Stick Up'),
        ),
        migrations.AlterField(
            model_name='activitysubmission',
            name='water_supply_system_name',
            field=models.CharField(blank=True, max_length=80, null=True, verbose_name='Water Supply System Name'),
        ),
        migrations.AlterField(
            model_name='activitysubmission',
            name='water_supply_system_well_name',
            field=models.CharField(blank=True, max_length=80, null=True, verbose_name='Water Supply System Well Name'),
        ),
        migrations.AlterField(
            model_name='licencedstatuscode',
            name='description',
            field=models.CharField(max_length=255, verbose_name='Licence Status'),
        ),
        migrations.AlterField(
            model_name='well',
            name='owner_city',
            field=models.CharField(blank=True, max_length=100, null=True, verbose_name='Town/City'),
        ),
    ]
