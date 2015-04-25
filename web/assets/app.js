"use strict";

var serversocket;

$(document).ready(function(){
	let btnCall, btnReset;

	btnCall = $('#call');
	btnReset = $('#reset');

	btnCall.click(function(){
		if (btnCall.prop("disable")) {
			return;
		} else {
			doPicar();
		}
	});

	btnCall.prop("disable", true);

	btnReset.click(function(){
		let tty = $('#tty');
		tty.html("");
	});


	serversocket = new WebSocket("ws://localhost:8088/ws");

	serversocket.onopen = function(e) {
		btnCall.prop("disable", false)
	};

	serversocket.onmessage = function(e) {
		writeTTY("---\n" + e.data + "\n---\n");
	};

});

function doPicar() {
	let data_path = $("#path").val();
	let data_prefix = $("#prefix").val();
	let data_noarchiving = $("#noarchiving").is(":checked")
	let data_debug = $("#debug").is(":checked")

	let args = {
		path: data_path,
		prefix: data_prefix,
		noarchiving: data_noarchiving,
		debug: data_debug
	};

	let params = JSON.stringify(args);

	writeTTY("发送参数：\n" + params);

	serversocket.send(params);
}

function writeTTY(message) {
	let tty = $("#tty");
	let text = tty.html();
	tty.html("\n" + message + text + "\n");
}
