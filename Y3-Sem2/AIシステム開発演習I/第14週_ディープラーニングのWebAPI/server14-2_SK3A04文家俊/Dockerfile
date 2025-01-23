# Dockerfile
FROM tensorflow/tensorflow:latest

EXPOSE 8000

WORKDIR /app
COPY ./requirements.txt /app
RUN pip install -r requirements.txt

CMD ["hug", "-f", "api_image.py"]
