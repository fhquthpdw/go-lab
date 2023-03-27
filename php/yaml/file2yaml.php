<?php
$yamlArr = file2yaml("istio.yaml");
print_r($yamlArr);

function file2yaml($file) {
    return yaml_parse(file_get_contents($file));
}