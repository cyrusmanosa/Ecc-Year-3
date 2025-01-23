#api_simple.py
import hug

@hug.get("/")
def api_example():
    return {"message": "Development of Web API!"}
@hug.get("/hello")
def information():
    return {"id":"2220042","name":"文家俊"}

# http://IPaddr:8000/hello
