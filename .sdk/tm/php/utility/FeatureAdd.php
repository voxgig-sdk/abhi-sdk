<?php
declare(strict_types=1);

// Abhi SDK utility: feature_add

class AbhiFeatureAdd
{
    public static function call(AbhiContext $ctx, mixed $f): void
    {
        $ctx->client->features[] = $f;
    }
}
