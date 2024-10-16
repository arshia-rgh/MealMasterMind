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

# Base Url
BASE_URL = "auth-service:8080"

# Cors configs
origins = ["*"]
