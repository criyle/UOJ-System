<?
$database_password = 'test';
$judger_socket_port = '233';

$config = include '/var/www/uoj/app/.default-config.php';
$config['database']['password']=$database_password;
$config['judger']['socket']['port']=$judger_socket_port;
file_put_contents('/var/www/uoj/app/.config.php', "<?php\nreturn ".str_replace('\'_httpHost_\'','UOJContext::httpHost()',var_export($config, true)).";\n");
