import re

from pydantic import BaseModel, ValidationError, field_validator

from user.app import config

PASSWORD_REGEX = config.PASSWORD_REGEX
PHONE_NUMBER_REGEX = config.PHONE_NUMBER_REGEX


class PasswordValidator(BaseModel):
    @field_validator("password", check_fields=False)
    @classmethod
    def validate_password(cls, password):
        if not re.match(PASSWORD_REGEX, password):
            raise ValidationError("Password must be at least 8 characters long and contain both letters and numbers")
        return password


class PhoneNumberValidator(BaseModel):
    @field_validator("phone_number", check_fields=False)
    @classmethod
    def validate_phone_number(cls, phone_number):
        if phone_number and not re.match(PHONE_NUMBER_REGEX, phone_number):
            raise ValidationError("Phone number must be a valid Iranian phone number")
        return phone_number


class PasswordMatchingValidator(BaseModel):
    @model_validator(mode="after")
    def check_new_password_matching(self) -> Self:
        password = self.password
        confirm_password = self.confirm_password

        if password != confirm_password:
            raise ValidationError("Passwords do not match")

        return self
