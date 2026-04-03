from fastapi import APIRouter

from routers.test_router import router as test_router

api_router = APIRouter()
api_router.include_router(test_router, tags=["test"])