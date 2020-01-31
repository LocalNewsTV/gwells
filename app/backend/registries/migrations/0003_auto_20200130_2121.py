# Generated by Django 2.2.9 on 2020-01-30 21:21

from django.db import migrations
from django.core.exceptions import ObjectDoesNotExist

NS_CERT_GUID = 'e6c6ab53-7ae0-4549-ad2c-d45fce8afc7f'
NB_CERT_GUID = '625af649-28e9-4f79-9154-b7079e0c56ee'

def add_more_certs(apps, schema_editor):
    CertifyingAuthorityCode = apps.get_model('registries', 'CertifyingAuthorityCode')
    AccreditedCertificateCode = apps.get_model('registries', 'AccreditedCertificateCode')
    ActivityCode = apps.get_model('registries', 'ActivityCode')

    # Attempt to get the existing DRILL AND PUMP ActivityCode from the DB if it
    # exists. If it doesn't (normally because of a fresh DB has been created)
    # then create it for the AccreditedCertificateCode() call below.
    drill_code = None
    try:
        drill_code = ActivityCode.objects.get(registries_activity_code='DRILL')
    except ObjectDoesNotExist:
        drill_code = ActivityCode.objects.create(
            pk='DRILL',
            create_user='DATALOAD_USER',
            create_date='2018-01-01T08:00:00Z',
            update_user='DATALOAD_USER',
            update_date='2018-01-01T08:00:00Z',
            description='Well Driller',
            display_order=2,
            effective_date='2018-05-25T18:20:35.936Z',
            expiry_date='9999-12-31T23:59:59Z'
        )

    pump_code = None
    try:
        pump_code = ActivityCode.objects.get(registries_activity_code='PUMP')
    except ObjectDoesNotExist:
        pump_code = ActivityCode.objects.create(
            pk='PUMP',
            create_user='DATALOAD_USER',
            create_date='2018-01-01T08:00:00Z',
            update_user='DATALOAD_USER',
            update_date='2018-01-01T08:00:00Z',
            description='Pump Installer',
            display_order=4,
            effective_date='2018-05-25T18:20:35.936Z',
            expiry_date='9999-12-31T23:59:59Z'
        )


    new_brunswick = CertifyingAuthorityCode(
        create_user='DATALOAD_USER',
        update_user='DATALOAD_USER',
        cert_auth_code='NB',
        description='Province of New Brunswick',
        effective_date='1970-01-01T08:00:00Z',
        expiry_date='9999-12-31T23:59:59Z'
    )
    new_brunswick.save()

    nova_scotia = CertifyingAuthorityCode(
        create_user='DATALOAD_USER',
        update_user='DATALOAD_USER',
        cert_auth_code='NS',
        description='Province of Nova Scotia',
        effective_date='1970-01-01T08:00:00Z',
        expiry_date='9999-12-31T23:59:59Z'
    )
    nova_scotia.save()

    nb_driller_permit = AccreditedCertificateCode(
        acc_cert_guid=NB_CERT_GUID,
        create_user='DATALOAD_USER',
        update_user='DATALOAD_USER',
        cert_auth=new_brunswick,
        registries_activity=drill_code,
        name='Well Driller’s Permit',
        description='',
        effective_date='2020-01-01T08:00:00Z',
        expiry_date='9999-12-31T23:59:59Z'
    )
    nb_driller_permit.save()

    ns_pump_cert = AccreditedCertificateCode(
        acc_cert_guid=NS_CERT_GUID,
        create_user='DATALOAD_USER',
        update_user='DATALOAD_USER',
        cert_auth=nova_scotia,
        registries_activity=pump_code,
        name='Certificate of Qualification Pump Installer Class I',
        description='',
        effective_date='2020-01-01T08:00:00Z',
        expiry_date='9999-12-31T23:59:59Z'
    )
    ns_pump_cert.save()

def remove_nb_ns_certs(apps, schema_editor):
    CertifyingAuthorityCode = apps.get_model('registries', 'CertifyingAuthorityCode')
    AccreditedCertificateCode = apps.get_model('registries', 'AccreditedCertificateCode')

    AccreditedCertificateCode.objects.get(acc_cert_guid=NB_CERT_GUID).delete()
    AccreditedCertificateCode.objects.get(acc_cert_guid=NB_CERT_GUID).delete()

    CertifyingAuthorityCode.objects.get(cert_auth_code='NB').delete()
    CertifyingAuthorityCode.objects.get(cert_auth_code='NS').delete()

class Migration(migrations.Migration):

    dependencies = [
        ('registries', '0002_auto_20200120_1617'),
    ]

    operations = [
        migrations.RunPython(add_more_certs, remove_nb_ns_certs),
    ]
