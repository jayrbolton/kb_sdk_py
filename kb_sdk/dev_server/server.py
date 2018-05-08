"""
SDK development server with Flask

Flask docs: http://flask.pocoo.org/docs/1.0
"""

from flask import Flask

app = Flask(__name__)


@app.route('/')
def root():
    return 'Hola mundo!'