# Generated by Django 2.2.18 on 2021-03-24 21:03

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('wells', '0125_alter_fieldsprovided_water_1658'),
    ]

    operations = [
        migrations.AlterField(
            model_name='activitysubmission',
            name='transmissivity',
            field=models.DecimalField(max_digits=8,
                                      decimal_places=6,
                                      blank=True,
                                      null=True,
                                      verbose_name='Transmissivity')
        ),
        migrations.AlterField(
            model_name='well',
            name='transmissivity',
            field=models.DecimalField(max_digits=8,
                                      decimal_places=6,
                                      blank=True,
                                      null=True,
                                      verbose_name='Transmissivity')
        ),
        migrations.AlterField(
            model_name='hydraulicproperty',
            name='transmissivity',
            field=models.DecimalField(max_digits=8,
                                      decimal_places=6,
                                      blank=True,
                                      null=True,
                                      verbose_name='Transmissivity')
        )
    ]
