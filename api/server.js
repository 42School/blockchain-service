var		express = require('express');
var		request = require('request');
var		app = express();

const	apiContract = require('./apiContract');
const	url = 'localhost';
const	port = 3001;

request({url: 'https://api.intra.42.fr/oauth/token/', method: 'POST',
  auth: {
    user: process.env.API_42_UID,
    pass: process.env.API_42_PASS
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