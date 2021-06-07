<?php
return [
	'profile' => [
		'oj-name' => 'Universal Online Judge',
		'oj-name-short' => 'UOJ',
		'administrator' => 'root',
		'admin-email' => 'admin@local_uoj.ac',
		'QQ-group' => '',
		'ICP-license' => ''
	],
	'database' => [
		'database' => 'app_uoj233',
		'username' => 'root',
		'password' => 'test',
		'host' => 'uoj_mysql'
	],
	'web' => [
		'domain' => null,
		'main' => [
			'protocol' => 'http',
			'host' => 'localhost',
			'port' => 4567
		],
		'blog' => [
			'protocol' => 'http',
			'host' => 'localhost',
			'port' => 4567
		]
	],
	'render_server' => [
		'host' => 'uoj_render_server',
		'port' => 3456
	],
	'security' => [
		'user' => [
			'client_salt' => 'salt0'
		],
		'cookie' => [
			'checksum_salt' => ['salt1', 'salt2', 'salt3']
		],
	],
	'mail' => [
		'noreply' => [
			'username' => 'noreply@local_uoj.ac',
			'password' => '_mail_noreply_password_',
			'host' => 'smtp.local_uoj.ac',
			'secure' => 'tls',
			'port' => 587
		]
	],
	'judger' => [
		'socket' => [
			'port' => '233',
			'password' => '_judger_socket_password_'
		]
	],
	'switch' => [
		'web-analytics' => false,
		'blog-domain-mode' => 3
	]
];
