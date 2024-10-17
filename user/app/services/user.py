from typing import Optional

from app import config
from app.models.user import User
from app.schemas.user import (
    ChangePassword,
    ConfirmResetPassword,
    LoginUser,
    RegisterUser,
    RequestResetPassword,
    ResponseUser,
    UpdateUser,
)
from app.utils import hash_password
from app.utils.hash_password import verify_password
from app.utils.jwt import create_access_token, verify_access_token
from event.publish import publish_message
from fastapi import HTTPException, status
from fastapi.responses import JSONResponse
from sqlalchemy.exc import NoResultFound
from sqlalchemy.orm import Session


def create_user(db: Session, user: RegisterUser) -> JSONResponse | ResponseUser:
    hashed_password = hash_password.hash_password(user.password)
    db_user = User(
        first_name=user.first_name,
        last_name=user.last_name,
        username=user.username,
        email=user.email,
        password=hashed_password,
        phone_number=user.phone_number,
    )

    db.add(db_user)
    try:
        db.commit()
    except Exception as e:
        return JSONResponse(status_code=status.HTTP_400_BAD_REQUEST, content={"error": str(e)})

    db.refresh(db_user)

    return ResponseUser.model_validate(db_user)


def authenticate_user(db: Session, login_data: LoginUser) -> Optional[dict]:
    try:
        db_user = db.query(User).filter(User.username == login_data.username).one()

    except NoResultFound:
        return None

    if not verify_password(login_data.password, db_user.password):
        return None

    access_token = create_access_token({"sub": db_user.username})

    return {"access_token": access_token, "token-type": "bearer"}


def update_user(db: Session, updated_user: UpdateUser, current_user: ResponseUser) -> ResponseUser:
    db_user = db.query(User).filter(User.id == current_user.id).first()

    if not db_user:
        raise HTTPException(status_code=status.HTTP_404_NOT_FOUND, detail="Useer not found")

    db_user.first_name = updated_user.first_name or db_user.first_name
    db_user.last_name = updated_user.last_name or db_user.last_name
    db_user.username = updated_user.username or db_user.username
    db_user.email = updated_user.email or db_user.email
    db_user.phone_number = updated_user.phone_number or db_user.phone_number

    db.commit()
    db.refresh(db_user)

    return ResponseUser.model_validate(db_user)


def delete_user(db: Session, current_user: ResponseUser) -> JSONResponse:
    db_user = db.query(User).filter(User.id == current_user.id).first()

    if not db_user:
        raise HTTPException(status_code=status.HTTP_404_NOT_FOUND, detail="Useer not found")

    db.delete(db_user)
    db.commit()

    return JSONResponse(status_code=status.HTTP_200_OK, content={"message": "User deleted successfully."})


def change_password(db: Session, updated_data: ChangePassword, current_user: ResponseUser) -> JSONResponse:
    db_user = db.query(User).filter(User.id == current_user.id).first()

    if not db_user:
        raise HTTPException(status_code=status.HTTP_404_NOT_FOUND, detail="Useer not found")

    if not verify_password(updated_data.old_password, db_user.password):
        return JSONResponse(status_code=status.HTTP_400_BAD_REQUEST, content={"message": "old password is not correct"})

    db_user.password = hash_password.hash_password(updated_data.password)

    db.commit()
    db.refresh(db_user)

    return JSONResponse(status_code=status.HTTP_200_OK, content={"message": "password changed successfully"})


def request_reset_password(db: Session, email: RequestResetPassword) -> JSONResponse:
    db_user = db.query(User).filter(User.email == email.email).first()

    if not db_user:
        raise HTTPException(status_code=status.HTTP_404_NOT_FOUND, detail="User not found")

    reset_token = create_access_token(
        {"sub": email.email}, expire_minutes=config.ACCESS_TOKEN_EXPIRE_MINUTES_FOR_RESET_PASSWORD
    )
    resset_link = f"http://{config.BASE_URL}/api/confirm-reset-password/{reset_token}/"

    publish_message("send-mail", {"email": email.email, "link": resset_link})

    return JSONResponse(status_code=status.HTTP_200_OK, content={"message": "Password reset link sent to your email"})


def confirm_reset_password(db: Session, token: str, change_password_data: ConfirmResetPassword) -> JSONResponse:
    decoded_token = verify_access_token(token)

    email = decoded_token.get("sub")

    db_user = db.query(User).filter(User.email == email).first()

    db_user.password = hash_password.hash_password(change_password_data.password)

    db.commit()
    db.refresh(db_user)

    return JSONResponse(status_code=status.HTTP_200_OK, content={"message": "password changed successfully"})
