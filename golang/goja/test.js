"use strict";

function booleanToggleSensor() {
	console.log("sensor(" + this.uid + "): " + this.value);

	if (this.value === false)
		this.value = true;
	else
		this.value = false;
}

function integerIncrementSensor() {
	this.value++;
}

function integerSawtoothSensor() {
	if (this.up == null)
		this.up = true;

	if (this.up)
		this.value++;
	else
		this.value--;

	if (this.value == 4)
		this.up = false;
	else if (this.value == -4)
		this.up = true;
}

sensor.new(1000, "tst1-bool", "bool",   booleanToggleSensor,    false);
sensor.new(1001, "tst2-bool", "bool",   booleanToggleSensor,    true);
sensor.new(1002, "tst3-incr", "number", integerIncrementSensor, 0);
sensor.new(1003, "tst3-sawt", "number", integerSawtoothSensor,  0);

Request(function(msg) {
	console.log("anonymous cb");
	console.log("msg", msg);
});

Request("device:ping", function() {
	console.log("device:ping");
});

Request("config:get", 3, function(param) {
	console.log(JSON.stringify(param));
});

Request("config:set", 3, function(param) {
	console.log(JSON.stringify(param));
});

Request("config:reset", 3, function(param) {
	console.log(JSON.stringify(param));
	console.log(this.SayHello());
});

Request("sensor:info", function() {
	console.log("sensor:info");
});

Request("sensor:set", 3, function(param) {
	console.log(JSON.stringify(param));
});

Request("simulator:set", "sensor", function(param) {
	
});

//for (var i = 0; i < 20; i++) {
//	sensor.update(1003);
//}
