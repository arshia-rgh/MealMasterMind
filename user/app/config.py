import os

from dotenv import load_dotenv

load_dotenv()


class DBConfig:
    DB_USER = os.getenv("DB_USER")
    DB_PASSWORD = os.getenv("DB_PASSWORD")
    DB_HOST = os.getenv("DB_HOST")
    DB_PORT = os.getenv("DB_PORT")
    DB_NAME = os.getenv("DB_NAME")


SECRET_KEY = str(os.getenv("SECRET_KEY"))
ALGORITHM = "HS256"
ACCESS_TOKEN_EXPIRE_MINUTES = 20
