<?php

class UOJBlogEditor {
	public $type = 'blog';
	public $name;
	public $blog_url;
	public $save;
	public $cur_data = array();
	public $post_data = array();
	
	public $label_text = array(
		'title' => '标题',
		'tags' => '标签（多个标签用逗号隔开）',
		'content' => '内容',
		'view blog' => '查看博客',
		'blog visibility' => '博客可见性',
		'private' => '未公开',
		'public' => '公开'
	);
	
	public $validator = array();
	
	function __construct() {
		global $REQUIRE_LIB;
		$REQUIRE_LIB['blog-editor'] = '';
		
		$this->validator = array(
			'title' => function(&$title) {
				if ($title == '') {
					return '标题不能为空';
				}
				if (strlen($title) > 100) {
					return '标题不能超过 100 个字节';
				}
				if (HTML::escape($title) === '') {
					return '无效编码';
				}
				return '';
			},
			'content_md' => function(&$content_md) {
				if (strlen($content_md) > 1000000) {
					return '内容过长';
				}
				return '';
			},
			'tags' => function(&$tags) {
				$tags = str_replace('，', ',', $tags);
				$tags_raw = explode(',', $tags);
				if (count($tags_raw) > 10) {
					return '标签个数不能超过10';
				}
				$tags = array();
				foreach ($tags_raw as $tag) {
					$tag = trim($tag);
					if (strlen($tag) == 0) {
						continue;
					}
					if (strlen($tag) > 30) {
						return '标签 “' . HTML::escape($tag) .'” 太长';
					}
					if (in_array($tag, $tags, true)) {
						return '标签 “' . HTML::escape($tag) .'” 重复出现';
					}
					$tags[] = $tag;
				}
				return '';
			}
		);
	}
	
	public function validate($name) {
		if (!isset($_POST["{$this->name}_{$name}"])) {
			return '不能为空';
		}
		$this->post_data[$name] = $_POST["{$this->name}_{$name}"];
		$val = $this->validator[$name];
		return $val($this->post_data[$name]);
	}
	private function receivePostData() {
		$errors = array();
		foreach (array('title', 'content_md', 'tags') as $name) {
			$cur_err = $this->validate($name);
			if ($cur_err) {
				$errors[$name] = $cur_err;
			}
		}
		if ($errors) {
			die(json_encode($errors));
		}
		crsf_defend();
		
		$this->post_data['is_hidden'] = isset($_POST["{$this->name}_is_hidden"]) ? 1 : 0;
		
		$purifier = HTML::pruifier();
		
		$this->post_data['title'] = HTML::escape($this->post_data['title']);

		$url = 'http://'.UOJConfig::$data['render_server']['host'].':'.UOJConfig::$data['render_server']['port'].'/';
		
		if ($this->type == 'blog') {
			$content_md = $_POST[$this->name . '_content_md'];
			
			$option = array(
				'http' => array(
					'header' => array('Content-Type: text/plain',
						'Context-Length: ' . strlen($this->post_data['content_md'])),
					'method' => 'POST',
					'content' => $this->post_data['content_md']
				)
			);
			$context = stream_context_create($option);
			$result = @file_get_contents($url, false, $context);
			if ($result === false) {
				die(json_encode(array('content_md' => 'failed to request markdown rendering service')));
			}
			$this->post_data['content'] = $result;

			if (preg_match('/^.*<!--.*readmore.*-->.*$/m', $this->post_data['content'], $matches, PREG_OFFSET_CAPTURE)) {
				$content_less = substr($this->post_data['content'], 0, $matches[0][1]);
				$content_more = substr($this->post_data['content'], $matches[0][1] + strlen($matches[0][0]));
				$this->post_data['content'] = $purifier->purify($content_less).'<!-- readmore -->'.$purifier->purify($content_more);
			} else {
				$this->post_data['content'] = $purifier->purify($this->post_data['content']);
			}
		} else if ($this->type == 'slide') {
			$content_array = @yaml_parse($this->post_data['content_md']);
			if ($content_array === false || !is_array($content_array)) {
				die(json_encode(array('content_md' => '不合法的 YAML 格式')));
			}

			$marked = function($md) use($url, $purifier) {
				$option = array(
					'http' => array(
						'header' => array('Content-Type: text/plain',
							'Context-Length: ' . strlen($this->post_data['content_md'])),
						'method' => 'POST',
						'content' => $md
					)
				);
				$context = stream_context_create($option);
				$result = @file_get_contents($url, false, $context);
				if ($result === false) {
					die(json_encode(array('content_md' => 'failed to request markdown rendering service')));
				}
				return $purifier->purify($result);
			};
			
			$config = array();
			$this->post_data['content'] = '';
			foreach ($content_array as $slide_name => $slide_content) {
				if (is_array($slide_content) && is_array($slide_content['config'])) {
					foreach (array('theme') as $config_key) {
						if (is_string($slide_content['config'][$config_key]) && strlen($slide_content['config'][$config_key]) <= 30) {
							$config[$config_key] = $slide_content['config'][$config_key];
						}
					}
					continue;
				}
				
				$this->post_data['content'] .= '<section>';
				
				if (is_string($slide_content)) {
					$this->post_data['content'] .= $marked($slide_content);
				} elseif (is_array($slide_content)) {
					if (is_array($slide_content['children'])) {
						foreach ($slide_content['children'] as $cslide_name => $cslide_content) {
							$this->post_data['content'] .= '<section>';
							$this->post_data['content'] .= $marked($cslide_content);
							$this->post_data['content'] .= '</section>';
						}
					}
				}
				$this->post_data['content'] .= "</section>\n";
			}
			$this->post_data['content'] = json_encode($config) . "\n" . $this->post_data['content'];
		}
	}
	
	public function handleSave() {
		$save = $this->save;
		$this->receivePostData();
		$ret = $save($this->post_data);
		if (!$ret) {
			$ret = array();
		}
		
		if (isset($_POST['need_preview'])) {
			ob_start();
			if ($this->type == 'blog') {
				echoUOJPageHeader('博客预览', array('ShowPageHeader' => false, 'REQUIRE_LIB' => array('mathjax' => '', 'shjs' => '')));
				echo '<article>';
				echo $this->post_data['content'];
				echo '</article>';
				echoUOJPageFooter(array('ShowPageFooter' => false));
			} elseif ($this->type == 'slide') {
				uojIncludeView('slide', array_merge(
					UOJContext::pageConfig(),
					array(
						'PageTitle' => '幻灯片预览',
						'content' => $this->post_data['content']
					)
				));
			}
			$ret['html'] = ob_get_contents();
			ob_end_clean();
		}
		
		die(json_encode($ret));
	}
	
	public function runAtServer() {
		if (isset($_POST["save-{$this->name}"])) {
			$this->handleSave();
		}
	}
	public function printHTML() {
		uojIncludeView('blog-editor', array('editor' => $this));
	}
}
