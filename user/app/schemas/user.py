from pydantic import BaseModel


class BaseUser(BaseModel):
    username: str
    email: str

class RegisterUser(BaseUser):
    first_name: str
    last_name: str
    phone_number: str
    password: str
