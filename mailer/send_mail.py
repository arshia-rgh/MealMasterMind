from email.mime.multipart import MIMEMultipart

from jinja2 import Environment, FileSystemLoader

import config

msg = MIMEMultipart()
msg["From"] = config.from_email
msg["To"] = config.to_email
msg["Subject"] = config.subject

env = Environment(loader=FileSystemLoader("/templates"))
template = env.get_template("email.html")
