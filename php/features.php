<?php
declare(strict_types=1);

// Abhi SDK feature factory

require_once __DIR__ . '/feature/BaseFeature.php';
require_once __DIR__ . '/feature/TestFeature.php';


class AbhiFeatures
{
    public static function make_feature(string $name)
    {
        switch ($name) {
            case "base":
                return new AbhiBaseFeature();
            case "test":
                return new AbhiTestFeature();
            default:
                return new AbhiBaseFeature();
        }
    }
}
