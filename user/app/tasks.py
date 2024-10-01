from typing import Optional

from celery import shared_task
from fastapi_mail import ConnectionConfig, FastMail, MessageSchema

from user.app import config

conf = ConnectionConfig(
    MAIL_USERNAME=config.MAIL_USERNAME,
    MAIL_PASSWORD=config.MAIL_PASSWORD,
    MAIL_PORT=config.MAIL_PORT,
    MAIL_SERVER=config.MAIL_SERVER,
    MAIL_FROM_NAME=config.MAIL_FROM_NAME,
    MAIL_STARTTLS=config.MAIL_STARTTLS,
    MAIL_SSL_TLS=config.MAIL_SSL_TLS,
    USE_CREDENTIALS=config.USE_CREDENTIALS,
    TEMPLATE_FOLDER="user/app/email-templates",
)


@shared_task
def send_email(subject: str, recipients: list[str], body: dict, template_name: Optional[str], subtype: str = "plain"):
    if template_name is None and subtype == "html":
        raise ValueError("Template name must be provided for HTML emails")

    message = MessageSchema(
        subject=subject,
        recipients=recipients,
        body=body,
        subtype=subtype,
    )

    fm = FastMail(conf)

    fm.send_message(message, template_name=template_name)
