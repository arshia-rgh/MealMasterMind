import os

from dotenv import load_dotenv

load_dotenv()


# DB configs
class DBConfig:
    DB_USER = os.getenv("DB_USER")
    DB_PASSWORD = os.getenv("DB_PASSWORD")
    DB_HOST = os.getenv("DB_HOST")
    DB_PORT = os.getenv("DB_PORT")
    DB_NAME = os.getenv("DB_NAME")


# JWT token configs
SECRET_KEY = str(os.getenv("SECRET_KEY"))
ALGORITHM = "HS256"
ACCESS_TOKEN_EXPIRE_MINUTES = 20

# Validations regex
PHONE_NUMBER_REGEX = r"^(09|\+989)\d{9,10}$"
PASSWORD_REGEX = r"^(?=.*[A-Za-z])(?=.*\d)[A-Za-z\d]{8,}$"

# Celery configs
CELERY_BROKER = os.getenv("CELERY_BROKER")
CELERY_BACKEND = os.getenv("CELERY_BACKEND")
CELERY_RESULT_EXPIRE = 3600
CELERY_TASK_SERIALIZER = "json"
CELERY_ACCEPT_CONTENT = ["json"]
CELERY_RESULT_SERIALIZER = "json"
CELERY_TIMEZONE = "Asia/Tehran"
CELERY_ENABLE_UTC = True
