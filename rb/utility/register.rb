# Abhi SDK utility registration
require_relative '../core/utility_type'
require_relative 'clean'
require_relative 'done'
require_relative 'make_error'
require_relative 'feature_add'
require_relative 'feature_hook'
require_relative 'feature_init'
require_relative 'fetcher'
require_relative 'make_fetch_def'
require_relative 'make_context'
require_relative 'make_options'
require_relative 'make_request'
require_relative 'make_response'
require_relative 'make_result'
require_relative 'make_point'
require_relative 'make_spec'
require_relative 'make_url'
require_relative 'param'
require_relative 'prepare_auth'
require_relative 'prepare_body'
require_relative 'prepare_headers'
require_relative 'prepare_method'
require_relative 'prepare_params'
require_relative 'prepare_path'
require_relative 'prepare_query'
require_relative 'result_basic'
require_relative 'result_body'
require_relative 'result_headers'
require_relative 'transform_request'
require_relative 'transform_response'

AbhiUtility.registrar = ->(u) {
  u.clean = AbhiUtilities::Clean
  u.done = AbhiUtilities::Done
  u.make_error = AbhiUtilities::MakeError
  u.feature_add = AbhiUtilities::FeatureAdd
  u.feature_hook = AbhiUtilities::FeatureHook
  u.feature_init = AbhiUtilities::FeatureInit
  u.fetcher = AbhiUtilities::Fetcher
  u.make_fetch_def = AbhiUtilities::MakeFetchDef
  u.make_context = AbhiUtilities::MakeContext
  u.make_options = AbhiUtilities::MakeOptions
  u.make_request = AbhiUtilities::MakeRequest
  u.make_response = AbhiUtilities::MakeResponse
  u.make_result = AbhiUtilities::MakeResult
  u.make_point = AbhiUtilities::MakePoint
  u.make_spec = AbhiUtilities::MakeSpec
  u.make_url = AbhiUtilities::MakeUrl
  u.param = AbhiUtilities::Param
  u.prepare_auth = AbhiUtilities::PrepareAuth
  u.prepare_body = AbhiUtilities::PrepareBody
  u.prepare_headers = AbhiUtilities::PrepareHeaders
  u.prepare_method = AbhiUtilities::PrepareMethod
  u.prepare_params = AbhiUtilities::PrepareParams
  u.prepare_path = AbhiUtilities::PreparePath
  u.prepare_query = AbhiUtilities::PrepareQuery
  u.result_basic = AbhiUtilities::ResultBasic
  u.result_body = AbhiUtilities::ResultBody
  u.result_headers = AbhiUtilities::ResultHeaders
  u.transform_request = AbhiUtilities::TransformRequest
  u.transform_response = AbhiUtilities::TransformResponse
}
