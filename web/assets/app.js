function doPicar() {
	var data_path = $("#path").val();
	var data_prefix = $("#prefix").val();
	var data_noarchiving = $("#noarchiving").is(":checked")
	var data_debug = $("#debug").is(":checked")

	var args = {
		path: data_path,
		prefix: data_prefix,
		noarchiving: data_noarchiving,
		debug: data_debug
	};


	$.post('/call/', args, function(data){
		alert(data['message']);
	});
}