$(function() {
	$("#trumpcraft-submit").on("click", function(e) {
		var source = $("#trumpcraft-source").val()
		var words = parseInt($("#trumpcraft-words").val())
		if(words > 1000) {
			words = 1000
		}

		$.getJSON("/api/" + source + "/" + words, function(d) {
			$("#trumpcraft-results").text('"' + d.Text + '"')
		})
	})

	$("#trumpcraft-submit").click()
})