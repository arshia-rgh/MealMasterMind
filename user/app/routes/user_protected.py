from app.db.database import get_db
from app.dependencies import get_current_user
from app.schemas.user import ChangePassword, ResponseUser, UpdateUser
from app.services.user import change_password, delete_user, update_user
from fastapi import APIRouter, Depends
from sqlalchemy.orm import Session
from starlette.responses import JSONResponse

router = APIRouter(prefix="/protected")


@router.get("/me/", response_model=ResponseUser)
async def read_current_user(current_user: ResponseUser = Depends(get_current_user)) -> ResponseUser:
    return current_user


@router.put("/me/", response_model=ResponseUser)
async def update_current_user(
    updated_user: UpdateUser, db: Session = Depends(get_db), current_user: ResponseUser = Depends(get_current_user)
) -> ResponseUser:
    return await update_user(db, updated_user, current_user)


@router.delete("/me/")
async def delete_current_user(
    db: Session = Depends(get_db), current_user: ResponseUser = Depends(get_current_user)
) -> JSONResponse:
    return await delete_user(db, current_user)


@router.post("/me/change-password/")
async def change_current_user_password(
    updated_data: ChangePassword, db: Session = Depends(get_db), current_user: ResponseUser = Depends(get_current_user)
) -> JSONResponse:
    return await change_password(db, updated_data, current_user)
