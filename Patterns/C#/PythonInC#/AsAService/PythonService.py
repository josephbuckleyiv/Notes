from flask import Flask

app = Flask(__name__)

@app.route('/python/calculation', methods=['GET'])
def calculation():
    return 9

if __name__ == '__main__':
    app.run(port=8000)
