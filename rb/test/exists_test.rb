# Abhi SDK exists test

require "minitest/autorun"
require_relative "../Abhi_sdk"

class ExistsTest < Minitest::Test
  def test_create_test_sdk
    testsdk = AbhiSDK.test(nil, nil)
    assert !testsdk.nil?
  end
end
