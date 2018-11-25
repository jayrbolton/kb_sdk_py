"""
This file contains all the methods for this KBase SDK module.
"""
import kbase_module


@kbase_module.method('echo')
def echo(params):
    """Echo back the given message."""
    return params['message']
