from fastapi import APIRouter, Depends, HTTPException, status
from fastapi.security import OAuth2PasswordRequestForm
from sqlalchemy.orm import Session

from user.app.database import get_db
from user.app.schemas.user import RegisterUser, RequestResetPassword, ResponseUser
from user.app.services.user import authenticate_user, create_user, request_reset_password

router = APIRouter()


@router.post("/register/", response_model=ResponseUser)
async def register_user(user: RegisterUser, db: Session = Depends(get_db)):
    return create_user(db, user)


@router.post("/login/")
async def login_user(form_data: OAuth2PasswordRequestForm = Depends(), db: Session = Depends(get_db)):
    token = authenticate_user(db, form_data.username, form_data.password)

    if not token:
        raise HTTPException(status_code=status.HTTP_401_UNAUTHORIZED, detail="invalid username or password")

    return token


@router.post("/request-reset-password/")
async def forget_password(email: RequestResetPassword, db: Session = Depends(get_db)):
    return request_reset_password(db, email)


@router.post("/confirm-reset-password/{token}/")
def confirm_forget_password():
    pass
