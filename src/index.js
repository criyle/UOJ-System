import express from 'express'
import marked from 'marked';
import bodyParser from 'body-parser';

marked.setOptions({
	getLangClass: function(lang) {
		lang = lang.toLowerCase();
		switch (lang) {
			case 'c': return 'c';
			case 'c++': return 'cpp';
			case 'pascal': return 'pascal';
			default: return lang;
		}
	},
	getElementClass: function(tok) {
		switch (tok.type) {
			case 'list_item_start':
				return 'fragment';
			case 'loose_item_start':
				return 'fragment';
			default:
				return null;
		}
	}
})

var app = express()
app.use(bodyParser.text())

app.post('/', function (req, res) {
  res.send(marked(req.body))
})

const server = app.listen('3456', () => {
  console.log('Server is listening on port 3456')
})
