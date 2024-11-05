from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class AuthReq(_message.Message):
    __slots__ = ("token",)
    TOKEN_FIELD_NUMBER: _ClassVar[int]
    token: str
    def __init__(self, token: _Optional[str] = ...) -> None: ...

class AuthRes(_message.Message):
    __slots__ = ("userID", "userEmail")
    USERID_FIELD_NUMBER: _ClassVar[int]
    USEREMAIL_FIELD_NUMBER: _ClassVar[int]
    userID: int
    userEmail: str
    def __init__(self, userID: _Optional[int] = ..., userEmail: _Optional[str] = ...) -> None: ...
