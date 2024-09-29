from fastapi import APIRouter, Depends
from sqlalchemy.orm import Session

from user.app.database import get_db
from user.app.dependencies import get_current_user
from user.app.schemas.user import ResponseUser, UpdateUser, ChangePassword
from user.app.services.user import update_user, delete_user, change_password

router = APIRouter(prefix="/protected")


@router.get("/me/", response_model=ResponseUser)
async def read_current_user(current_user: ResponseUser = Depends(get_current_user)):
    return current_user


@router.put("/me/", response_model=ResponseUser)
async def update_current_user(updated_user: UpdateUser, db: Session = Depends(get_db),
                              current_user: ResponseUser = Depends(get_current_user)):
    return update_user(db, updated_user, current_user)


@router.delete("/me/")
async def delete_current_user(db: Session = Depends(get_db), current_user: ResponseUser = Depends(get_current_user)):
    return delete_user(db, current_user)


@router.post("/me/change-password/")
async def change_current_user_password(updated_data: ChangePassword, db: Session = Depends(get_db),
                                       current_user: ResponseUser = Depends(get_current_user)):
    return change_password(db, updated_data, current_user)
