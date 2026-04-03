from flask import Blueprint
from routers.test_router import test_router

api_router = Blueprint("api", __name__)
api_router.register_blueprint(test_router)