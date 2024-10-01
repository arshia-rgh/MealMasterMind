from celery import shared_task
from fastapi_mail import ConnectionConfig

from user.app import config

conf = ConnectionConfig(
    MAIL_USERNAME=config.MAIL_USERNAME,
    MAIL_PASSWORD=config.MAIL_PASSWORD,
    MAIL_PORT=config.MAIL_PORT,
    MAIL_SERVER=config.MAIL_SERVER,
    MAIL_FROM_NAME=config.MAIL_FROM_NAME,
    MAIL_TLS=config.MAIL_TLS,
    MAIL_SSL=config.MAIL_SSL,
    USE_CREDENTIALS=config.USE_CREDENTIALS,
)


@shared_task
def send_email():
    pass
