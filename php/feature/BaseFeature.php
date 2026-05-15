<?php
declare(strict_types=1);

// Abhi SDK base feature

class AbhiBaseFeature
{
    public string $version;
    public string $name;
    public bool $active;

    public function __construct()
    {
        $this->version = '0.0.1';
        $this->name = 'base';
        $this->active = true;
    }

    public function get_version(): string { return $this->version; }
    public function get_name(): string { return $this->name; }
    public function get_active(): bool { return $this->active; }

    public function init(AbhiContext $ctx, array $options): void {}
    public function PostConstruct(AbhiContext $ctx): void {}
    public function PostConstructEntity(AbhiContext $ctx): void {}
    public function SetData(AbhiContext $ctx): void {}
    public function GetData(AbhiContext $ctx): void {}
    public function GetMatch(AbhiContext $ctx): void {}
    public function SetMatch(AbhiContext $ctx): void {}
    public function PrePoint(AbhiContext $ctx): void {}
    public function PreSpec(AbhiContext $ctx): void {}
    public function PreRequest(AbhiContext $ctx): void {}
    public function PreResponse(AbhiContext $ctx): void {}
    public function PreResult(AbhiContext $ctx): void {}
    public function PreDone(AbhiContext $ctx): void {}
    public function PreUnexpected(AbhiContext $ctx): void {}
}
