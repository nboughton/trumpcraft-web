$(function () {
	$("#trumpcraft-submit").on("click", function (e) {
		var source = $("#trumpcraft-source").val()
		var words = $("#trumpcraft-words").val()

		$.getJSON("api/" + source + "/" + words, function (d) {
			$("#trumpcraft-results").text('"' + d.data + '"')
		})
	})

	$("#trumpcraft-submit").click()
})