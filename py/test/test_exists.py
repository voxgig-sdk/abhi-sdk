# ProjectName SDK exists test

import pytest
from abhi_sdk import AbhiSDK


class TestExists:

    def test_should_create_test_sdk(self):
        testsdk = AbhiSDK.test(None, None)
        assert testsdk is not None
