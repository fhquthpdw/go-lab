<?php
exit("-1");

$f = $argv[1];
$yamlArr = file2yaml($f);
print_r($yamlArr);

function file2yaml($file) {
    return yaml_parse(file_get_contents($file));
}