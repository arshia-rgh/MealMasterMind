import smtplib
from email.mime.multipart import MIMEMultipart
from email.mime.text import MIMEText

from jinja2 import Environment, FileSystemLoader

import config

msg = MIMEMultipart()
msg["From"] = config.from_email
msg["To"] = config.to_email
msg["Subject"] = config.subject

env = Environment(loader=FileSystemLoader("/templates"))
template = env.get_template("email.html")

context = {
    "link":  # TODO: Get from gRPC
    "subject": # TODO: Get from gRPC
    "email": # TODO: Get from gRPC
}

html_content = template.render(context)

html_part = MIMEText(html_content, "html")
msg.attach(html_part)

with smtplib.SMTP(config.smtp_server, config.smtp_port) as server:
    server.starttls()
    server.login(config.smtp_user, config.smtp_password)
    server.sendmail(config.from_email, config.to_email, msg.as_string())
