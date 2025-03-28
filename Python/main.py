import os
import sys
import requests

API_KEY = os.environ.get("API_KEY")
API_URL = "https://fontys.cloud-builders.nl/api/v1/send-image"

def main():
    if API_KEY is None:
        print("Environment variable API_KEY is not set")
        sys.exit(1)
    if len(sys.argv) < 4:
        print("Usage: python main.py <x> <y> <image_path>")
        sys.exit(1)

    try:
        x: int = int(sys.argv[1])
        y: int = int(sys.argv[2])
    except ValueError:
        print("x and y must be an integer")
        sys.exit(1)

    image_path: str = sys.argv[3]

    headers = {
        "X-Api-Key": API_KEY
    }
    url = f"{API_URL}?x={x}&y={y}"
    with open(file=image_path, mode="rb") as image_file:
        image = image_file.read()

    requests.post(
        url = url,
        headers = headers,
        files = {
            "image": (image_path, image, "image/jpeg")
        },
        timeout=30,
    )

if __name__ == "__main__":
    main()
