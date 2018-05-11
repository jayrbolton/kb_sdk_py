"""
Execute `kb-sdk dev`
(this is called from ./cli.py)

Run a flask-based simple development server
"""

import kb_sdk.status_server.server as server


def execute(args, config, env):
    server.app.config['CONFIG'] = config
    server.app.run()
    server.app.logger.info('hi!')
