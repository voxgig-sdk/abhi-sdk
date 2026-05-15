<?php
declare(strict_types=1);

// Abhi SDK utility: result_body

class AbhiResultBody
{
    public static function call(AbhiContext $ctx): ?AbhiResult
    {
        $response = $ctx->response;
        $result = $ctx->result;
        if ($result && $response && $response->json_func && $response->body) {
            $result->body = ($response->json_func)();
        }
        return $result;
    }
}
