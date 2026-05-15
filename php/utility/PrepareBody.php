<?php
declare(strict_types=1);

// Abhi SDK utility: prepare_body

class AbhiPrepareBody
{
    public static function call(AbhiContext $ctx): mixed
    {
        if ($ctx->op->input === 'data') {
            return ($ctx->utility->transform_request)($ctx);
        }
        return null;
    }
}
