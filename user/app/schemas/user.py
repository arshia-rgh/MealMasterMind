from pydantic import BaseModel


class BaseUser(BaseModel):
    username: str
    email: str


class RegisterUser(BaseUser):
    first_name: str
    last_name: str
    phone_number: str
    password: str


class ResponseUser(BaseModel):
    id: int
    first_name: str
    last_name: str
    username: str
    email: str
    phone_number: str

    class Config:
        from_attributes = True
