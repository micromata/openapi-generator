# coding: utf-8

"""
    OpenAPI Petstore

    This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\  # noqa: E501

    The version of the OpenAPI document: 1.0.0
    Generated by: https://openapi-generator.tech
"""


from __future__ import absolute_import
import sys
import unittest

import petstore_api
try:
    from petstore_api.model import category
except ImportError:
    category = sys.modules[
        'petstore_api.model.category']
try:
    from petstore_api.model import tag
except ImportError:
    tag = sys.modules[
        'petstore_api.model.tag']
from petstore_api.model.pet import Pet


class TestPet(unittest.TestCase):
    """Pet unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def testPet(self):
        """Test Pet"""
        # FIXME: construct object with mandatory attributes with example values
        # model = Pet()  # noqa: E501
        pass


if __name__ == '__main__':
    unittest.main()
