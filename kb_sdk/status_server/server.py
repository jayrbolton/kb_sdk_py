"""
SDK development server with Flask

Flask docs: http://flask.pocoo.org/docs/1.0
"""

import logging
from flask import Flask, render_template
import os
from dotenv import load_dotenv

from kb_sdk.logger import logger, main_handler

load_dotenv(dotenv_path='./.env', override=True)

app = Flask(__name__)
app.logger.handlers = [main_handler]
app.config['ENV'] = 'development'
app.debug = True
level = os.getenv('LOG_LEVEL', logging.INFO)
app.logger.setLevel(50)


@app.route('/')
def root():
    app.logger.critical('hi?')
    app.logger.info('hi!')
    config = app.config['CONFIG']
    env = os.environ
    template_data = {
        'config': config,
        'env': env,
        'registered': True
    }
    return render_template('index.html', **template_data)
