from fastapi import APIRouter

router = APIRouter()

@router.get("/")
async def test():
    return {"message": "IA Integration Microservice is running with Python and FastAPI"}