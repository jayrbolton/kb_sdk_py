"""
This allows you to run
- a local app with param validation and config passed in
- a remote app (TODO)

You can call your Main class directly, but you won't get param validation and config
"""


class RunApp:

    def __init__(self, Cls):
        """ Initialize an app runner. Pass the app's class """
        config = {}
        self.instance = Cls(config)
        print(self.instance)
