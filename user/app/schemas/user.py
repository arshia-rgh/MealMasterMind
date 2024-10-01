from typing import Optional, Self

from pydantic import BaseModel, EmailStr, Field, ValidationError, model_validator

from user.app.schemas.validators import PasswordValidator, PhoneNumberValidator


class RegisterUser(PasswordValidator, PhoneNumberValidator):
    first_name: Optional[str] = None
    last_name: Optional[str] = None
    username: str
    email: EmailStr
    password: str = Field(..., min_length=8)
    phone_number: Optional[str] = None


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

    @model_validator(mode="after")
    def check_new_password_matching(self) -> Self:
        password = self.password
        confirm_password = self.confirm_password

        if password != confirm_password:
            raise ValidationError("Passwords do not match")

        return self


class RequestResetPassword(BaseModel):
    email: EmailStr
