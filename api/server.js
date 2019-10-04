var		express = require('express');
var		app = express();
var		bodyParser = require('body-parser');
var		request = require('request');

const	apiContract = require('./apiContract');
const	url = 'localhost';
const	port = 3001;
var		accessToken = '';

request({url: 'https://api.intra.42.fr/oauth/token/', method: 'POST',
  auth: {
    user: '36366b251a0c85fbe2950f703750f2f2504c5ceecae03b99ca21ea1a3226de59',
    pass: '21505a596ec210ca46f7d1330b1c56c84133b759c2ae92d6a3bca1c1bf0be218'
  }, form: {'grant_type': 'client_credentials'}
}, (err, res) => {
	  let json = JSON.parse(res.body);
	  process.env.ACCESS_TOKEN = json.access_token;
	  console.log("Access Token:", json.access_token);
});

app.use('', apiContract);

const server = app.listen(port, function () {
	console.log(`---  Server Start in Url http://${url}:${port} ---`);
	console.log('-----------------------------------------------------');
	const all_routes = require('express-list-endpoints');
	console.log(all_routes(app));
	console.log('-----------------------------------------------------');
});