window.onbeforeunload = function() {
	return "YO CHILL"
}

function updateNote(newContent, noteID, sessionKey) {
	//prep request
	var request = new XMLHttpRequest();
	request.open("POST", window.location.host + "/save", true);
	request.setRequestHeader("Content-Type", "application/json");
	request.send(JSON.stringify({
		"session-key": sessionKey,
		"note-id": noteID,
		"new-content": newContent
	}));
}