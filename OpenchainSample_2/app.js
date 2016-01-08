/**
 * Created by davery on 1/4/2016.
 */
var openchain = require('./api');
var TypeEnum = api_1.ChaincodeSpec.TypeEnum;
var api_1 = require("./api");
// Create a new instance of the chaincode API
var chaincode = new openchain.DevopsApi('', 'http://108.168.183.174:5000');
// Create the chaincode spec for chaincode example 2
var spec = new openchain.ChaincodeSpec();
spec.type = TypeEnum.GOLANG;
spec.chaincodeID.url = "https://hub.jazz.net/git/averyd/cc_ex02/chaincode_example02";
spec.chaincodeID.version = "0.0.1";
spec.ctorMsg.function = "init";
spec.ctorMsg.args = ["a", "200", "b", "300"];
// Wrap this spec into a deployment spec
var deploySpec = new openchain.ChaincodeDeploymentSpec();
deploySpec.chaincodeSpec = spec;
var promise = chaincode.chaincodeDeploy(spec);
promise.then(function (accepted) {
    console.log(accepted.body);
}, function (rejected) {
    console.error(rejected.body);
});
