#api_keras.py
import hug
from tensorflow.keras.models import load_model
import numpy as np
import joblib

@hug.get("/")
def api_root():
    return {"message":"Development of Web API!"}

# http://IPアドレス:8000/hello
@hug.get("/hello")
def hello():
    return {"id":"2220042","name":"ブンカシュン"}

# http://IPアドレス:8000/iris?
@hug.get("/iris")
def iris(x1: float,x2: float,x3: float,x4: float):
    m = load_model('iris2220042.keras')
    r1 = m.predict([[x1, x2, x3, x4]])
    r2 = joblib.load('iris_names.joblib')
    return {"result":r1,"names":r2}
