from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class MailRequestResetLink(_message.Message):
    __slots__ = ("email", "subject", "link")
    EMAIL_FIELD_NUMBER: _ClassVar[int]
    SUBJECT_FIELD_NUMBER: _ClassVar[int]
    LINK_FIELD_NUMBER: _ClassVar[int]
    email: str
    subject: str
    link: str
    def __init__(self, email: _Optional[str] = ..., subject: _Optional[str] = ..., link: _Optional[str] = ...) -> None: ...

class MailResponseResetLink(_message.Message):
    __slots__ = ("result",)
    RESULT_FIELD_NUMBER: _ClassVar[int]
    result: str
    def __init__(self, result: _Optional[str] = ...) -> None: ...
