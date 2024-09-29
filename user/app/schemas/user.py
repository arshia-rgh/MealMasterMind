import re
from typing import Optional, Self

from pydantic import BaseModel, model_validator, ValidationError, EmailStr, Field, field_validator

phone_number_regex = r"^(09|\+989)\d{9,10}$"
password_regex = r"^(?=.*[A-Za-z])(?=.*\d)[A-Za-z\d]{8,}$"


class RegisterUser(BaseModel):
    first_name: Optional[str] = None
    last_name: Optional[str] = None
    username: str
    email: EmailStr
    password: str = Field(..., min_length=8)
    phone_number: Optional[str] = None

    @field_validator('password')
    def validate_password(self, password):
        if not re.match(password_regex, password):
            raise ValidationError('Password must be at least 8 characters long and contain both letters and numbers')
        return password

    @field_validator('phone_number')
    def validate_phone_number(self, phone_number):
        if phone_number and not re.match(phone_number_regex, phone_number):
            raise ValidationError('Phone number must be a valid Iranian phone number')
        return phone_number


class ResponseUser(BaseModel):
    id: int
    first_name: Optional[str] = None
    last_name: Optional[str] = None
    username: str
    email: str
    phone_number: Optional[str] = None

    class Config:
        from_attributes = True


class UpdateUser(BaseModel):
    first_name: Optional[str] = None
    last_name: Optional[str] = None
    username: Optional[str] = None
    email: Optional[EmailStr] = None
    phone_number: Optional[str] = None

    @field_validator('phone_number')
    def validate_phone_number(self, phone_number):
        if phone_number and not re.match(phone_number_regex, phone_number):
            raise ValidationError('Phone number must be a valid Iranian phone number')
        return phone_number


class ChangePassword(BaseModel):
    old_password: str
    password: str = Field(..., min_length=8)
    confirm_password: str

    @field_validator('password')
    def validate_password(self, password):
        if not re.match(password_regex, password):
            raise ValidationError('Password must be at least 8 characters long and contain both letters and numbers')
        return password

    @model_validator(mode="after")
    def check_new_password_matching(self) -> Self:
        password = self.password
        confirm_password = self.confirm_password

        if password != confirm_password:
            raise ValidationError("Passwords do not match")

        return self
