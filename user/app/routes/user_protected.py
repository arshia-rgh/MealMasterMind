from fastapi import APIRouter, Depends

from user.app.dependencies import get_current_user
from user.app.schemas.user import ResponseUser

router = APIRouter(prefix="/protected")


@router.get("/me/", response_model=ResponseUser)
async def read_current_user(current_user: ResponseUser = Depends(get_current_user)):
    return current_user
