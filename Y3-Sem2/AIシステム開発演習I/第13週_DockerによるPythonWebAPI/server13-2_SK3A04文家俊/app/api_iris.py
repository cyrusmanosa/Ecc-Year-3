#api_iris.py
#必要なライブラリのインポート
import joblib
import hug

@hug.get("/")
def api_example():
    return {"message": "Development of Web API!"}

# http://IPアドレス:8000/hello
@hug.get("/hello")
def hello():
    return {"id":"2220042","name":"文家俊"}

# http://IPアドレス:8000/iris?
@hug.get("/iris")
def iris(x1: float,x2: float,x3: float,x4: float):
    m = joblib.load('iris2220042.joblib')
    r1 = m.predict([[x1, x2, x3, x4]])
    r2 = joblib.load('iris_names.joblib')
    return {"result":r1,"names":r2}
