from fastapi import HTTPException, status
from sqlalchemy.exc import NoResultFound
from sqlalchemy.orm import Session

from user.app.models.user import User
from user.app.schemas.user import RegisterUser, ResponseUser, UpdateUser
from user.app.utils import hash_password
from user.app.utils.hash_password import verify_password
from user.app.utils.jwt import create_access_token


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


def authenticate_user(db: Session, username: str, password: str):
    try:
        db_user = db.query(User).filter(User.username == username).one()

    except NoResultFound:
        return None

    if not verify_password(password, db_user.password):
        return None

    access_token = create_access_token({"sub": db_user.username})

    return {"access_token": access_token, "token-type": "bearer"}


def update_user(db: Session, updated_user: UpdateUser, current_user: ResponseUser) -> ResponseUser:
    db_user = db.query(User).filter(User.id == current_user.id).first()

    if not db_user:
        raise HTTPException(
            status_code=status.HTTP_404_NOT_FOUND,
            detail="Useer not found"
        )

    db_user.first_name = updated_user.first_name
    db_user.last_name = updated_user.last_name
    db_user.username = updated_user.username
    db_user.email = updated_user.email
    db_user.phone_number = updated_user.phone_number
    db_user.password = hash_password.hash_password(updated_user.password)

    db.commit()
    db.refresh(db_user)

    return ResponseUser.model_validate(db_user)
