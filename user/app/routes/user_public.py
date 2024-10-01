from fastapi import APIRouter, Depends, HTTPException, status
from sqlalchemy.orm import Session
from starlette.responses import JSONResponse

from user.app.database import get_db
from user.app.schemas.user import ConfirmResetPassword, LoginUser, RegisterUser, RequestResetPassword, ResponseUser
from user.app.services.user import authenticate_user, confirm_reset_password, create_user, request_reset_password

router = APIRouter()


@router.post("/register/", response_model=ResponseUser)
async def register_user(user: RegisterUser, db: Session = Depends(get_db)) -> JSONResponse | ResponseUser:
    return create_user(db, user)


@router.post("/login/")
async def login_user(login_data: LoginUser, db: Session = Depends(get_db)) -> JSONResponse:
    token = authenticate_user(db, login_data)

    if not token:
        raise HTTPException(status_code=status.HTTP_401_UNAUTHORIZED, detail="invalid username or password")

    return JSONResponse(status_code=status.HTTP_200_OK, content=token)


@router.post("/request-reset-password/")
async def forget_password(email: RequestResetPassword, db: Session = Depends(get_db)) -> JSONResponse:
    return request_reset_password(db, email)


@router.post("/confirm-reset-password/{token}/")
def confirm_forget_password(
    token: str, change_password_data: ConfirmResetPassword, db: Session = Depends(get_db)
) -> JSONResponse:
    return confirm_reset_password(db, token, change_password_data)
