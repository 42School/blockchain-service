var		express = require('express');
var		request = require('request');
var		dotenv = require('dotenv');
var		cors = require('cors')
var		app = express();

const	mongoose = require('mongoose');

mongoose.connect('mongodb://localhost:27017/graduate_marvin', {
  useNewUrlParser: true,
  useUnifiedTopology: true
});

var db = mongoose.connection;
db.on('error', console.error.bind(console, 'Erreur lors de la connexion'));
db.once('open', function (){
    console.log("Connexion à la base OK");
});

dotenv.config();

const	apiContract = require('./apiContract');
const	apiSign = require('./apiSign');
const	url = process.env.URL;
const	port = process.env.PORT;

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

app.use(cors());
app.use('', apiContract);
app.use('/sign', apiSign);

const server = app.listen(port, function () {
	console.log(`---  Server Start in Url http://${url}:${port} ---`);
	console.log('-----------------------------------------------------');
	const all_routes = require('express-list-endpoints');
	console.log(all_routes(app));
	console.log('-----------------------------------------------------');
});