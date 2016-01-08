/**
 * Created by davery on 1/4/2016.
 */
import openchain = require('./api');
import TypeEnum = openchain.ChaincodeSpec.TypeEnum;
import ChaincodeID = openchain.ChaincodeID;
import ChaincodeMessage = openchain.ChaincodeMessage;

// Create a new instance of the chaincode API
var chaincode = new openchain.DevopsApi('', 'http://108.168.183.174:5000');

// Create the chaincode spec for chaincode example 2
var spec = new openchain.ChaincodeSpec();
spec.type = TypeEnum.GOLANG;
spec.chaincodeID = new ChaincodeID();
spec.chaincodeID.path = "https://hub.jazz.net/git/averyd/cc_ex02/chaincode_example02";
spec.ctorMsg = new ChaincodeMessage();
spec.ctorMsg.function = "init";
spec.ctorMsg.args = ["a", "200", "b", "300"];

// Wrap this spec into a deployment spec
var deploySpec = new openchain.ChaincodeDeploymentSpec();
deploySpec.chaincodeSpec = spec;

console.log("Deploying chaincode");
var promise = chaincode.chaincodeDeploy(spec);

promise.then(function(accepted) {
    console.log(accepted.body);
    spec.chaincodeID = new ChaincodeID();
    spec.chaincodeID.name = accepted.body.name;
    spec.ctorMsg = new ChaincodeMessage();
    spec.ctorMsg.function = "invoke";
    spec.ctorMsg.args = ["a", "b", "33"];
    var invokeSpec = new openchain.ChaincodeInvocationSpec();
    invokeSpec.chaincodeSpec = spec;
    return chaincode.chaincodeInvoke(invokeSpec);
}, function(rejected) {
    console.error(rejected.body);

}).then(function(invoked) {
    console.log(invoked.body);
    spec.ctorMsg = new ChaincodeMessage();
    spec.ctorMsg.function = "query";
    spec.ctorMsg.args = ["a"];
    var invokeSpec = new openchain.ChaincodeInvocationSpec();
    invokeSpec.chaincodeSpec = spec;
    return chaincode.chaincodeQuery(invokeSpec);
}, function(notInvoked) {
    console.error(notInvoked.body);
}).then(function(queried) {
    console.log(queried.body);
}, function(notQueried) {
    console.error(notQueried.body);
});

