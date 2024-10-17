import os

smtp_server = 'smtp.gmail.com'
smtp_port = 587
smtp_user = os.getenv("MAIL_USERNAME")
smtp_password = os.getenv("MAIL_PASSWORD")

from_email = os.getenv("MAIL_USERNAME")
to_email =  # TODO: get from gRPC
subject =  # TODO: get From gRPC
