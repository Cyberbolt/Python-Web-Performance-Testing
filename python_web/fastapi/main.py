from fastapi import FastAPI

app = FastAPI()


@app.get('/')
async def main():
    return {
        'code': 1
    }
