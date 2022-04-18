const note = document.getElementById("text");
const shortcutKeys = {
    "s": function() {
        console.log("saved");
    }
}

window.onkeydown = function(e) {
    if (e.ctrlKey && e.key in shortcutKeys) {
        e.preventDefault();
        shortcutKeys[e.key]();
    }
}