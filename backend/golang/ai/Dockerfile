FROM python:3.9-slim-buster
WORKDIR /
RUN apt-get update && apt-get install -y \
    build-essential \
    pkg-config \
    libhdf5-dev \
    && rm -rf /var/lib/apt/lists/*
COPY requirements.txt requirements.txt
RUN pip install -U pip
RUN pip install --no-cache-dir -r requirements.txt
COPY . .
EXPOSE 8080
CMD python3 server.py