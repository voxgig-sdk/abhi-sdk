<?php
declare(strict_types=1);

// Abhi SDK utility registration

require_once __DIR__ . '/../core/UtilityType.php';
require_once __DIR__ . '/Clean.php';
require_once __DIR__ . '/Done.php';
require_once __DIR__ . '/MakeError.php';
require_once __DIR__ . '/FeatureAdd.php';
require_once __DIR__ . '/FeatureHook.php';
require_once __DIR__ . '/FeatureInit.php';
require_once __DIR__ . '/Fetcher.php';
require_once __DIR__ . '/MakeFetchDef.php';
require_once __DIR__ . '/MakeContext.php';
require_once __DIR__ . '/MakeOptions.php';
require_once __DIR__ . '/MakeRequest.php';
require_once __DIR__ . '/MakeResponse.php';
require_once __DIR__ . '/MakeResult.php';
require_once __DIR__ . '/MakePoint.php';
require_once __DIR__ . '/MakeSpec.php';
require_once __DIR__ . '/MakeUrl.php';
require_once __DIR__ . '/Param.php';
require_once __DIR__ . '/PrepareAuth.php';
require_once __DIR__ . '/PrepareBody.php';
require_once __DIR__ . '/PrepareHeaders.php';
require_once __DIR__ . '/PrepareMethod.php';
require_once __DIR__ . '/PrepareParams.php';
require_once __DIR__ . '/PreparePath.php';
require_once __DIR__ . '/PrepareQuery.php';
require_once __DIR__ . '/ResultBasic.php';
require_once __DIR__ . '/ResultBody.php';
require_once __DIR__ . '/ResultHeaders.php';
require_once __DIR__ . '/TransformRequest.php';
require_once __DIR__ . '/TransformResponse.php';

AbhiUtility::setRegistrar(function (AbhiUtility $u): void {
    $u->clean = [AbhiClean::class, 'call'];
    $u->done = [AbhiDone::class, 'call'];
    $u->make_error = [AbhiMakeError::class, 'call'];
    $u->feature_add = [AbhiFeatureAdd::class, 'call'];
    $u->feature_hook = [AbhiFeatureHook::class, 'call'];
    $u->feature_init = [AbhiFeatureInit::class, 'call'];
    $u->fetcher = [AbhiFetcher::class, 'call'];
    $u->make_fetch_def = [AbhiMakeFetchDef::class, 'call'];
    $u->make_context = [AbhiMakeContext::class, 'call'];
    $u->make_options = [AbhiMakeOptions::class, 'call'];
    $u->make_request = [AbhiMakeRequest::class, 'call'];
    $u->make_response = [AbhiMakeResponse::class, 'call'];
    $u->make_result = [AbhiMakeResult::class, 'call'];
    $u->make_point = [AbhiMakePoint::class, 'call'];
    $u->make_spec = [AbhiMakeSpec::class, 'call'];
    $u->make_url = [AbhiMakeUrl::class, 'call'];
    $u->param = [AbhiParam::class, 'call'];
    $u->prepare_auth = [AbhiPrepareAuth::class, 'call'];
    $u->prepare_body = [AbhiPrepareBody::class, 'call'];
    $u->prepare_headers = [AbhiPrepareHeaders::class, 'call'];
    $u->prepare_method = [AbhiPrepareMethod::class, 'call'];
    $u->prepare_params = [AbhiPrepareParams::class, 'call'];
    $u->prepare_path = [AbhiPreparePath::class, 'call'];
    $u->prepare_query = [AbhiPrepareQuery::class, 'call'];
    $u->result_basic = [AbhiResultBasic::class, 'call'];
    $u->result_body = [AbhiResultBody::class, 'call'];
    $u->result_headers = [AbhiResultHeaders::class, 'call'];
    $u->transform_request = [AbhiTransformRequest::class, 'call'];
    $u->transform_response = [AbhiTransformResponse::class, 'call'];
});
