from fastapi import APIRouter, Depends, HTTPException, status
from fastapi.security import OAuth2PasswordRequestForm
from sqlalchemy.orm import Session

from user.app.database import get_db
from user.app.schemas.user import ResponseUser, RegisterUser
from user.app.services.user import create_user, authenticate_user

router = APIRouter()


@router.post("/register/", response_model=ResponseUser)
async def register_user(user: RegisterUser, db: Session = Depends(get_db)):
    return create_user(db, user)


@router.post("/login/")
async def login_user(form_data: OAuth2PasswordRequestForm = Depends(), db: Session = Depends(get_db)):
    token = authenticate_user(db, form_data.username, form_data.password)

    if not token:
        raise HTTPException(
            status_code=status.HTTP_401_UNAUTHORIZED,
            detail="invalid username or password"
        )

    return token
