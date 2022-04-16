const note = document.getElementById("text");
const shortcutKeys = {
    "s": function() {
        console.log("saved");
    },
    "b": function() {

    },
    "i": function() {

    },
    "u": function() {

    },
    "r": function() {
        var results = note.innerHTML.matchAll(/(\*{1,3})([^*]+)\1/g)
        for (var result of results) {
            console.log(result);
            note.innerHTML = replaceRange(note.innerHTML, result.index, result.index + result[0].length, genReplaceString(result[1], result[2]))
        }
    }
}
const formatCharsToHTML = {
    "**": ["<b>", "</b>"]
}

window.onkeydown = function(e) {
    if (e.ctrlKey && e.key in shortcutKeys) {
        e.preventDefault();
        shortcutKeys[e.key]();
    }
}

function replaceRange(str, start, end, replace) {
    return str.substring(0, start) + replace + str.substring(end);
}
function genReplaceString(formatChars, substr) {
    var tags = formatCharsToHTML[formatChars];
    return tags[0] + substr + tags[1];
}