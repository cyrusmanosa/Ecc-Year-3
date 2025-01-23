# Dockerfile for api_simple
FROM python:3.9.20

EXPOSE 8000

WORKDIR /app
COPY ./requirements.txt ./app
RUN pip install -r requirements.txt

CMD ["hug", "-f", "api_simple.py"]
