import os
import smtplib
from email.mime.multipart import MIMEMultipart
from email.mime.text import MIMEText

from jinja2 import Environment, FileSystemLoader

smtp_server = 'smtp.gmail.com'
smtp_port = 587
smtp_user = os.getenv("MAIL_USERNAME")
smtp_password = os.getenv("MAIL_PASSWORD")

from_email = os.getenv("MAIL_USERNAME")


def send_mail(message: dict):
    msg = MIMEMultipart()
    msg["From"] = from_email
    msg["To"] = message["email"]
    msg["Subject"] = message["subject"]

    env = Environment(loader=FileSystemLoader("./templates"))
    template = env.get_template("email.html")

    context = {
        "link": message["link"],
        "subject": message["subject"],
        "email": message["email"],
    }

    html_content = template.render(context)

    html_part = MIMEText(html_content, "html")
    msg.attach(html_part)

    with smtplib.SMTP(smtp_server, smtp_port) as server:
        server.starttls()
        server.login(smtp_user, smtp_password)
        server.sendmail(from_email, message["email"], msg.as_string())
