"""KBase Software Development Kit

Usage:
  kb-sdk (-h | --help)
  kb-sdk init <name> [<directory>]
  kb-sdk dev
  kb-sdk test [<module.Class.method>]

Commands:
  init     Initialize a new SDK app
  dev      Run the development server inside an app
  test     Run tests for an app

Options:
  -h --help    Show this screen.
  --version    Show version.
"""

from docopt import docopt
import kb_sdk.cli.init as init
import kb_sdk.cli.dev as dev
import kb_sdk.cli.test as test


def main():
    args = docopt(__doc__, version='0.0.1', help=True)
    if args['init']:
        executor = init
    elif args['test']:
        executor = test
    elif args['dev']:
        executor = dev
    print('exec:', executor.execute(args))
