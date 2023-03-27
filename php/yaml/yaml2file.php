<?php
$source = "istio.yaml";
$dest = "dest.yaml";

yaml2file($source, $dest);
function yaml2file($source, $dest) {
    $yamlArr = yaml_parse(file_get_contents($source));
    $yamlArr["spec"]["gateways"][] = "istio-system/internet-gw-1-test";
    yaml_emit_file($dest, $yamlArr);
}