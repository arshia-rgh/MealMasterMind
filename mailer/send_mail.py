from email.mime.multipart import MIMEMultipart

import config

msg = MIMEMultipart()
msg["From"] = config.from_email
msg["To"] = config.to_email
msg["Subject"] = config.subject
