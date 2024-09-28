from sqlalchemy import Column, Integer, String

from user.app.database import base


class User(base):
    __tablename__ = "users"

    id = Column(Integer, primary_key=True, autoincrement=True)
    first_name = Column(String, nullable=True)
    last_name = Column(String, nullable=True)
    username = Column(String, unique=True)
    email = Column(String, unique=True)
    password = Column(String)
    phone_number = Column(String, unique=True)
