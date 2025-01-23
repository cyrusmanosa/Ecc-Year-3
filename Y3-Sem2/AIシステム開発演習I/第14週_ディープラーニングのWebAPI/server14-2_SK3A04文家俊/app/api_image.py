#api_sklearn.py
import hug
from PIL import Image
import io
#その他必要なライブラリをインポートする
from tensorflow.keras.models import load_model
import joblib

# tf.keras.models.load_model()を使用してモデルを読み込む
import tensorflow as tf

@hug.get("/")
def api_root():
    return {"message":"Development of Web API!"}

# http://IPアドレス:8000/hello
@hug.get("/hello")
def hello():
    return {"id":"2220042","name":"ブンカシュン"}

# http://IPアドレス:8000/catsdogs にアクセスすると推定結果を返す
# POSTを使用して画像を受け取る
@hug.post("/catsdogs")
def api_catsdogs(file1):
    img_bytes = io.BytesIO(file1)
    img = Image.open(img_bytes)
    img = img.resize((224,224))
    m = load_model('catsdogs2220042.keras')

    # 画像を配列に変換
    img_array = tf.keras.preprocessing.image.img_to_array(img)

    # 軸の追加（1枚の画像のため）
    img_array = img_array[tf.newaxis, ...]

    # スケーリング（0～255 → 0～1）
    img_array = img_array/255.0

    # 予測
    r1 = m.predict(img_array)
    r2 = joblib.load('animal.joblib')
    pred = m.predict(img_array).round().astype(int)
    ans = list(r2)[pred[0][0]]
    if ans == 'cats':
        # 結果を返す
        return {"results": r1[0][0], "category": {"cat": 1, "dog": 0}}
    else:
        return {"results": r1[0][0], "category": {"cat": 0, "dog": 1}}
