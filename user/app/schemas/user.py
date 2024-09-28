from typing import Optional

from pydantic import BaseModel


class RegisterUser(BaseModel):
    first_name: Optional[str] = None
    last_name: Optional[str] = None
    username: str
    email: str
    password: str
    phone_number: Optional[str] = None


class ResponseUser(BaseModel):
    id: int
    first_name: str
    last_name: str
    username: str
    email: str
    phone_number: str

    class Config:
        from_attributes = True


class UpdateUser(BaseModel):
    first_name: Optional[str] = None
    last_name: Optional[str] = None
    username: Optional[str] = None
    email: Optional[str] = None
    phone_number: Optional[str] = None
    password: Optional[str] = None
