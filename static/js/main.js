const bingUrlBase = "https://www.bing.com"

$(document).ready(function() {
	$.getJSON("/bing", function(result) {
		var imageUrl = bingUrlBase + result.images[0].url;
		var copyright = result.images[0].copyright;
		var copyrightLink = result.images[0].copyrightlink;

		$("body").css("background", "url('" + imageUrl + "') no-repeat center center fixed");
		
		$("#copyright a").attr("href", copyrightLink);
		$("#copyright a").html(copyright);
		$("#copyright").removeClass("d-none");
	});

	$.getJSON("/altproto", function(result) {
		$.getJSON(result.address + "/json", function(result2) {
			$("#alt-version").html(result2.version);
			$("#alt-ip").html(result2.ip);
			$("#alt-reverse").html(result2.reverse);

			$("#alt").removeClass("d-none");
		});
	});

	$("body").on("click", ".clipboard", function() {
		var value = $(this).text();
		navigator.clipboard.writeText(value);
		alert("Copied value to clipboard");
	});
});