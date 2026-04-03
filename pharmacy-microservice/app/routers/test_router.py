from flask import Blueprint, jsonify

test_router = Blueprint("test", __name__)

@test_router.get("/")
def test():
    return jsonify({"message": "Pharmacy Microservice is running with Python and Flask"})