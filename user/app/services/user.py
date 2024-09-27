from sqlalchemy.orm import Session

from user.app.models.user import User
from user.app.schemas.user import RegisterUser, ResponseUser
from user.app.utils import hash_password


def create_user(db: Session, user: RegisterUser) -> ResponseUser:
    hashed_password = hash_password.hash_password(user.password)
    db_user = User(
        first_name=user.first_name,
        last_name=user.last_name,
        username=user.username,
        email=user.email,
        password=hashed_password,
        phone_number=user.phone_number
    )

    db.add(db_user)
    db.commit()
    db.refresh(db_user)

    return ResponseUser.model_validate(db_user)
