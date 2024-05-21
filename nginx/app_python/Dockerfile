# docker build -t my-py-app .
FROM python:3.10-alpine
WORKDIR /app
COPY requirements.txt /app/requirements.txt
RUN pip install -r requirements.txt
COPY . /app/
EXPOSE 6000
CMD ["python3", "my-py-app.py"]
