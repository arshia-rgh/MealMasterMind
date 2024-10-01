from datetime import datetime, timedelta

import jwt

from user.app import config

SECRET_KEY = config.SECRET_KEY
ALGORITHM = config.ALGORITHM
ACCESS_TOKEN_EXPIRE_MINUTES = config.ACCESS_TOKEN_EXPIRE_MINUTES


def create_access_token(data: dict, expire_minutes: int = ACCESS_TOKEN_EXPIRE_MINUTES):
    to_encode = data.copy()
    expire = datetime.now() + timedelta(minutes=expire_minutes)

    to_encode.update({"exp": expire})

    encoded = jwt.encode(to_encode, SECRET_KEY, ALGORITHM)

    return encoded


def verify_access_token(token: str):
    try:
        payload = jwt.decode(token, SECRET_KEY, ALGORITHM)
        return payload

    except jwt.PyJWTError:
        return None
