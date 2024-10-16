import os

# DB configs
DB_USER = os.getenv("POSTGRES_USER")
DB_PASSWORD = os.getenv("POSTGRES_PASSWORD")
DB_HOST = os.getenv("POSTGRES_HOST")
DB_PORT = os.getenv("POSTGRES_PORT")
DB_NAME = os.getenv("POSTGRES_DB")

# JWT token configs
SECRET_KEY = str(os.getenv("SECRET_KEY"))
ALGORITHM = "HS256"
ACCESS_TOKEN_EXPIRE_MINUTES = 20
ACCESS_TOKEN_EXPIRE_MINUTES_FOR_RESET_PASSWORD = 5

# Validations regex
PHONE_NUMBER_REGEX = r"^(09|\+989)\d{9,10}$"
PASSWORD_REGEX = r"^(?=.*[A-Za-z])(?=.*\d)[A-Za-z\d]{8,}$"

# Celery configs
CELERY_BROKER = os.getenv("CELERY_BROKER_URL")
CELERY_BACKEND = os.getenv("CELERY_BACKEND")
CELERY_RESULT_EXPIRE = 3600
CELERY_TASK_SERIALIZER = "json"
CELERY_ACCEPT_CONTENT = ["json"]
CELERY_RESULT_SERIALIZER = "json"
CELERY_TIMEZONE = "Asia/Tehran"
CELERY_ENABLE_UTC = True

# Mail configs
MAIL_USERNAME = os.getenv("MAIL_USERNAME")
MAIL_PASSWORD = os.getenv("MAIL_PASSWORD")
MAIL_FROM = os.getenv("MAIL_FROM")
MAIL_PORT = 587
MAIL_SERVER = "smtp.gmail.com"
MAIL_FROM_NAME = os.getenv("MAIL_FROM_NAME")
MAIL_STARTTLS = True
MAIL_SSL_TLS = False
USE_CREDENTIALS = True

# Base Url
BASE_URL = "auth-service:8080"

# Cors configs
origins = ["*"]
