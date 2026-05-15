<?php
declare(strict_types=1);

// Abhi SDK utility: result_headers

class AbhiResultHeaders
{
    public static function call(AbhiContext $ctx): ?AbhiResult
    {
        $response = $ctx->response;
        $result = $ctx->result;
        if ($result) {
            if ($response && is_array($response->headers)) {
                $result->headers = $response->headers;
            } else {
                $result->headers = [];
            }
        }
        return $result;
    }
}
