from fastapi import APIRouter

router = APIRouter()

@router.get("/")
async def test():
    return {"message": "Medical Record Microservice is running with Python and FastAPI"}