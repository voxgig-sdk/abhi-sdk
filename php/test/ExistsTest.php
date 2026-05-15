<?php
declare(strict_types=1);

// Abhi SDK exists test

require_once __DIR__ . '/../abhi_sdk.php';

use PHPUnit\Framework\TestCase;

class ExistsTest extends TestCase
{
    public function test_create_test_sdk(): void
    {
        $testsdk = AbhiSDK::test(null, null);
        $this->assertNotNull($testsdk);
    }
}
