"""
This file contains tests for src/main.py
"""
import unittest
from main import echo


class TestMain(unittest.TestCase):

    def test_echo(self):
        message = "Hello world!"
        result = echo(message)
        self.assertEqual(result, message)
