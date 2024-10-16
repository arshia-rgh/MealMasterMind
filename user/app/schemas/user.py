from typing import Optional

from app.schemas.validators import PasswordMatchingValidator, PasswordValidator, PhoneNumberValidator
from pydantic import BaseModel, EmailStr, Field


class RegisterUser(PasswordValidator, PhoneNumberValidator):
    first_name: Optional[str] = None
    last_name: Optional[str] = None
    username: str
    email: EmailStr
    password: str = Field(..., min_length=8)
    phone_number: Optional[str] = None


class LoginUser(PasswordValidator):
    username: str
    password: str = Field(..., min_length=8)


class ResponseUser(BaseModel):
    id: int
    first_name: Optional[str] = None
    last_name: Optional[str] = None
    username: str
    email: EmailStr
    phone_number: Optional[str] = None

    class Config:
        from_attributes = True


class UpdateUser(PhoneNumberValidator):
    first_name: Optional[str] = None
    last_name: Optional[str] = None
    username: Optional[str] = None
    email: Optional[EmailStr] = None
    phone_number: Optional[str] = None


class ChangePassword(PasswordValidator):
    old_password: str
    password: str = Field(..., min_length=8)
    confirm_password: str


class RequestResetPassword(BaseModel):
    email: EmailStr


class ConfirmResetPassword(PasswordMatchingValidator):
    password: str = Field(..., min_length=8)
    confirm_password: str
