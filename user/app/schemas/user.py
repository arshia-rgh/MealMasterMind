from typing import Optional, Any, Self

from pydantic import BaseModel, model_validator, ValidationError


class RegisterUser(BaseModel):
    first_name: Optional[str] = None
    last_name: Optional[str] = None
    username: str
    email: str
    password: str
    phone_number: Optional[str] = None


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
    email: Optional[str] = None
    phone_number: Optional[str] = None


class ChangePassword(BaseModel):
    old_password: str
    new_password: str
    confirm_password: str


    @model_validator(mode="after")
    def validate_passwords(self) -> Self:
        new_password = self.new_password
        confirm_password = self.confirm_password

        if new_password is not None and confirm_password is not None and new_password != confirm_password:
            raise ValidationError("Passwords do not match")

        return self
