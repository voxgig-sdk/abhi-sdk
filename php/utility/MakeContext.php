<?php
declare(strict_types=1);

// Abhi SDK utility: make_context

require_once __DIR__ . '/../core/Context.php';

class AbhiMakeContext
{
    public static function call(array $ctxmap, ?AbhiContext $basectx): AbhiContext
    {
        return new AbhiContext($ctxmap, $basectx);
    }
}
