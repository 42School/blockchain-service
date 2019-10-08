const	mongoose = require('mongoose');

let		intraLinksSchema = mongoose.Schema({
    login: String,
    graduateId: String,
    isForgot: Boolean
});

var intraLink = mongoose.model('intraLink', intraLinksSchema);

module.exports = {
	intraLink: intraLink
}