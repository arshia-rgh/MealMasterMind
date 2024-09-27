from fastapi import APIRouter, Depends
from sqlalchemy.orm import Session

from user.app.database import get_db
from user.app.schemas.user import ResponseUser, RegisterUser
from user.app.services.user import create_user

router = APIRouter()


@router.post("/register", response_model=ResponseUser)
def register_user(user: RegisterUser, db: Session = Depends(get_db)):
    return create_user(db, user)
